import json
import ujson
import avro
import timeit
import os

from avro.datafile import DataFileWriter, DataFileReader
from avro.io import DatumWriter, DatumReader

event_str = '{"scan_rate": 100, "node_id": "factory.unit.source", "value": 376.9000000000172, "type": "VariantType.Double", "timestamp": "2018-02-25T13:52:33.580957"}'
event = json.loads(event_str)
n = 1*500*1000

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

with open(json_filename, "w") as json_f:
    def write_ujson():
        json_f.write(ujson.dumps(event) + "\n")

    json_t = timeit.timeit(write_ujson, number=n)
    print("python,ujson,encoding,{},{:.2f}".format(n, json_t))


with open(json_filename, "r") as json_f:
    def read_ujson():
        s = json_f.readline()
        json.loads(s)
    json_wt = timeit.timeit(read_ujson, number=n)
    print("python,ujson,decoding,{},{:.2f}".format(n, json_wt))

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

print("AVRO output size: {:,} bytes".format(os.path.getsize(avro_filename)))
