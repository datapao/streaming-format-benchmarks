package main

import "encoding/json"
import "time"
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

import "github.com/linkedin/goavro"
import (
	"github.com/datapao/streaming-format-benchmarks/avro"
	"github.com/xeipuuv/gojsonschema"
)

type Event struct {
	ScanRate  int     `json:"scan_rate"`
	NodeId    string  `json:"node_id"`
	Value     float32 `json:"value"`
	Type      string  `json:"type"`
	TimeStamp string  `json:"timestamp"`
}

var N = 1000000
var eventStrStr = "{\"scan_rate\": 100, \"node_id\": \"factory.unit.source\", \"value\": 376.9000000000172, \"type\": \"VariantType.Double\", \"timestamp\": \"2018-02-25T13:52:33.580957\"}"
var eventStr = []byte(eventStrStr)
var BINARY []byte
var jsonSchema = gojsonschema.NewStringLoader(`{"type": "object", "properties": {"scan_rate": {"type": "number"}, "node_id": {"type": "string"}, "value": {"type": "number"}, "type": {"type": "string"}, "timestamp": {"type": "string"}}}`)

type trackFN func()

func timeTrack(name string, fn trackFN) {
	start := time.Now()
	for i := 0; i < N; i++ {
		fn()
	}
	elapsed := time.Since(start)
	fmt.Printf("%v,%.2f\n", name, float64(elapsed)/1e9)
}

func main() {
	// Event object of JSON
	var jsonEvent Event
	json.Unmarshal(eventStr, &jsonEvent)

	// Event object for AVRO
	avroEvent := avro.AVROEvent{
		Scan_rate: 100,
		Node_id:   "factory.unit.source",
		Value:     376.9000000000172,
		Type:      "VariantType.Double",
		Timestamp: "2018-02-25T13:52:33.580957",
	}

	// Event object for goavro
	var eventMap map[string]interface{}
	eventMap = make(map[string]interface{})
	eventMap["scan_rate"] = 100
	eventMap["node_id"] = "factory.unit.source"
	eventMap["value"] = 376.9000000000172
	eventMap["type"] = "VariantType.Double"
	eventMap["timestamp"] = "2018-02-25T13:52:33.580957"


	s := gojsonschema.NewStringLoader(eventStrStr)
	_, err := gojsonschema.Validate(jsonSchema, s)
	if err != nil {
		log.Fatal(err)
	}

	// Encode JSON
	f, err := os.Create("output.jsonlines")
	if err != nil {
		log.Fatal(err)
	}
	timeTrack(fmt.Sprintf("go,json,encoding,%v", N), encodeJSON(f, &jsonEvent))
	f.Close()

	// Decode JSON
	f, err = os.Open("output.jsonlines")
	scanner := bufio.NewScanner(f)
	timeTrack(fmt.Sprintf("go,json,decoding,%v", N), decodeJSON(scanner))
	f.Close()

	// Encode EasyJSON
	f, err = os.Create("output.jsonlines")
	if err != nil {
		log.Fatal(err)
	}
	timeTrack(fmt.Sprintf("go,easyjson,encoding,%v", N), encodeEasyJSON(f, jsonEvent))
	f.Close()

	// Decode EasyJSON
	f, err = os.Open("output.jsonlines")
	scanner = bufio.NewScanner(f)
	timeTrack(fmt.Sprintf("go,easyjson,decoding,%v", N), decodeEasyJSON(scanner))
	f.Close()

	// Decode EasyJSON w/ Schema
	f, err = os.Open("output.jsonlines")
	scanner = bufio.NewScanner(f)
	timeTrack(fmt.Sprintf("go,easyjson-jsonschema,decoding,%v", N), decodeEasyJSONWithSchema(scanner))
	f.Close()

	// Encode AVRO Codegen
	f, err = os.Create("events.avro")
	if err != nil {
		log.Fatal(err)
	}
	timeTrack(fmt.Sprintf("go,avro-codegen,encoding,%v", N), encodeAVROCodeGen(f, &avroEvent))
	f.Close()

	// Decode AVRO Codegen
	f, err = os.Open("events.avro")
	if err != nil {
		log.Fatal(err)
	}
	timeTrack(fmt.Sprintf("go,avro-codegen,decoding,%v", N), decodeAVROCodeGen(f))
	f.Close()

	// Encode AVRO nogen
	var codec *goavro.Codec
	codec, err = goavro.NewCodec(`
        {"namespace": "datapao.benchmark",
		 "type": "record",
		 "name": "AVROEvent",
		 "fields": [
			 {"name": "scan_rate", "type": "int"},
			 {"name": "node_id",  "type": "string"},
			 {"name": "value",  "type": "double"},
			 {"name": "type",  "type": "string"},
			 {"name": "timestamp",  "type": "string"}
		 ]
		}`)
	if err != nil {
		fmt.Println(err)
	}

	f, err = os.Create("events.goavro.avro")
	if err != nil {
		log.Fatal(err)
	}
	timeTrack(fmt.Sprintf("go,avro,encoding,%v", N), encodeAVRO(f, codec, eventMap))
	f.Close()

	timeTrack(fmt.Sprintf("go,avro,decoding,%v", N), decodeAVRO(codec))
}


func decodeAVRO(codec *goavro.Codec) trackFN {
	return func() {
		_, _, err := codec.NativeFromBinary(BINARY)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func encodeAVRO(file *os.File, codec *goavro.Codec, eventMap map[string]interface{}) trackFN {
	return func() {
		binary, err := codec.BinaryFromNative(nil, eventMap)
		if err != nil {
			fmt.Println(err)
		}
		file.Write(binary)
		if BINARY == nil {
			BINARY = binary
		}
	}
}

func encodeJSON(file *os.File, event *Event) trackFN {
	return func() {
		jsonStr, err := json.Marshal(event)
		if err != nil {
			log.Fatal(err)
		}
		file.Write(jsonStr)
		file.Write([]byte{'\n'})
	}
}

func decodeJSON(scanner *bufio.Scanner) trackFN {
	return func() {
		for scanner.Scan() {
			str := scanner.Text()
			var event Event
			err := json.Unmarshal([]byte(str), &event)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func decodeEasyJSON(scanner *bufio.Scanner) trackFN {
	return func() {
		for scanner.Scan() {
			str := scanner.Text()
			event := new(Event)
			err := event.UnmarshalJSON([]byte(str))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func encodeEasyJSON(file *os.File, event Event) trackFN {
	return func() {
		jsonStr, err := event.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
		file.Write(jsonStr)
		file.Write([]byte{'\n'})
	}
}

func decodeEasyJSONWithSchema(scanner *bufio.Scanner) trackFN {
	return func() {
		for scanner.Scan() {
			str := scanner.Text()
			event := new(Event)
			err := event.UnmarshalJSON([]byte(str))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func encodeAVROCodeGen(file *os.File, event *avro.AVROEvent) trackFN {
	return func() {
		err := event.Serialize(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func decodeAVROCodeGen(file *os.File) trackFN {
	return func() {
		_, err := avro.DeserializeAVROEvent(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}
