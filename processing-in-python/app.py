import json
import os
import timeit
import sys

import avro
import fastavro
import ujson
from avro.datafile import DataFileWriter, DataFileReader
from avro.io import DatumWriter, DatumReader
from jsonschema import validate

json_schema = {
    "type" : "object",
    "properties": {
        "scan_rate": {"type": "number"},
        "node_id":  {"type": "string"},
        "value": {"type": "number"},
        "type": {"type": "string"},
        "timestamp": {"type": "string"}
    }
}

event_str = '{"scan_rate": 100, "node_id": "factory.unit.source", "value": 376.9000000000172, "type": "VariantType.Double", "timestamp": "2018-02-25T13:52:33.580957"}'
event = json.loads(event_str)
n = 1*1000*1000

print("Starting benchmark...")
print("Sample event: {}".format(event_str))

json_filename = "benchmark_output.jsonlines"
avro_filename = "benchmark_output.avro"

with open(json_filename, "w") as json_f:
    def write_json():
        json_f.write(json.dumps(event) + "\n")

    json_t = timeit.timeit(write_json, number=n)
    print("python,json,encoding,{},{:.2f}".format(n, json_t))

with open(json_filename, "r") as json_f:
    def read_json():
        s = json_f.readline()
        json.loads(s)
    json_wt = timeit.timeit(read_json, number=n)
    print("python,json,decoding,{},{:.2f}".format(n, json_wt))

with open(json_filename, "r") as json_f:
    def read_json_schema():
        s = json_f.readline()
        j = json.loads(s)
        validate(j, json_schema)
    json_wt = timeit.timeit(read_json_schema, number=n)
    print("python,json-schema,decoding,{},{:.2f}".format(n, json_wt))

with open(json_filename, "r") as json_f:
    def read_json_manualschema():
        s = json_f.readline()
        j = json.loads(s)
        assert("scan_rate" in j)
        assert (type(j["scan_rate"]) is int)
        assert ("node_id" in j)
        assert (type(j["node_id"]) is str)
        assert ("value" in j)
        assert (type(j["value"]) is float)
        assert ("type" in j)
        assert (type(j["type"]) is str)
        assert ("timestamp" in j)
        assert (type(j["timestamp"]) is str)

    json_wt = timeit.timeit(read_json_manualschema, number=n)
    print("python,json-manualschema,decoding,{},{:.2f}".format(n, json_wt))

with open(json_filename, "w") as json_f:
    def write_ujson():
        json_f.write(ujson.dumps(event) + "\n")

    json_t = timeit.timeit(write_ujson, number=n)
    print("python,ujson,encoding,{},{:.2f}".format(n, json_t))

with open(json_filename, "r") as json_f:
    def read_ujson():
        s = json_f.readline()
        ujson.loads(s)
    json_wt = timeit.timeit(read_ujson, number=n)
    print("python,ujson,decoding,{},{:.2f}".format(n, json_wt))

with open(json_filename, "r") as json_f:
    def read_ujson_schema():
        s = json_f.readline()
        j = ujson.loads(s)
        validate(j, json_schema)
    json_wt = timeit.timeit(read_ujson_schema, number=n)
    print("python,ujson-schema,decoding,{},{:.2f}".format(n, json_wt))

with open(json_filename, "r") as json_f:
    def read_ujson_manualschema():
        s = json_f.readline()
        j = ujson.loads(s)
        assert("scan_rate" in j)
        assert (type(j["scan_rate"]) is int)
        assert ("node_id" in j)
        assert (type(j["node_id"]) is str)
        assert ("value" in j)
        assert (type(j["value"]) is float)
        assert ("type" in j)
        assert (type(j["type"]) is str)
        assert ("timestamp" in j)
        assert (type(j["timestamp"]) is str)

    json_wt = timeit.timeit(read_ujson_manualschema, number=n)
    print("python,ujson-manualschema,decoding,{},{:.2f}".format(n, json_wt))


print("JSON output size: {:,} bytes".format(os.path.getsize(json_filename)))



with open(avro_filename, "wb") as avro_f:
    def write_avro():
        avro_writer.append(event)
    avro_schema = avro.schema.Parse(open("event.avsc").read())
    avro_writer = DataFileWriter(avro_f, DatumWriter(), avro_schema)
    avro_t = timeit.timeit(write_avro, number=n)

    print("python,avro,encoding,{},{:.2f}".format(n, avro_t))
    avro_writer.close()

with open(avro_filename, "rb") as avro_f:
    def read_avro():
        u = reader.__iter__().__next__()
    reader = DataFileReader(avro_f, DatumReader())
    avro_wt = timeit.timeit(read_avro, number=n)
    print("python,avro,decoding,{},{:.2f}".format(n, avro_wt))
    reader.close()


with open(avro_filename, "wb") as avro_f:
    def write_avro():
        event_generator=(event for _ in range(n))
        fastavro.writer(avro_f,avro_schema,event_generator)
    avro_schema = fastavro.schema.load_schema("event.avsc")
    avro_t = timeit.timeit(write_avro, number=1)
    print("python,fastavro,encoding,{},{:.2f}".format(n, avro_t))

with open(avro_filename, "rb") as avro_f:
    def read_avro():
        u = reader.__iter__().__next__()
    reader = fastavro.reader(avro_f)
    avro_wt = timeit.timeit(read_avro, number=n)
    print("python,fastavro,decoding,{},{:.2f}".format(n, avro_wt))

print("AVRO output size: {:,} bytes".format(os.path.getsize(avro_filename)))
