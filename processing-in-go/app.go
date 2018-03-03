package main
import "encoding/json"
import "time"
import "log"

type Event struct {
    ScanRate   int      `json:"scan_rate"`
    NodeId    string    `json:"node_id"`
    Value 	  float32	`json:"value"`
    Type      string    `json:"type"`
    TimeStamp string    `json:"timestamp"`	
}

func main() {
	event_str := []byte("{\"scan_rate\": 100, \"node_id\": \"factory.unit.source\", \"value\": 376.9000000000172, \"type\": \"VariantType.Double\", \"timestamp\": \"2018-02-25T13:52:33.580957\"}")
	
    var totalTime time.Duration 
    totalTime = 0

	for i := 0; i < 1000000; i++ {
	    start := time.Now()
		var event Event
		json.Unmarshal(event_str, &event)
		totalTime += time.Since(start)
	}

    log.Printf("Unpacking took %s", totalTime)
}	