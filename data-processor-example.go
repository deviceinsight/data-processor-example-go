package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type datapoint struct {
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
	Timestamp string      `json:"timestamp"`
	DataType  string      `json:"dataType"`
}

type event struct {
	EventKey string `json:"eventKey"`
	//Priority  *int
	//Timestamp *string
}

type geoPosition struct {
	Latitude  float32
	Longitude float32
	Timestamp string
}

type threshold struct {
	High  *float64
	Low   *float64
	Event *event `json:"event"`
}

type filter struct {
	Keys      []string `json:"keys"`
	keysMap   map[string]struct{}
	Threshold *threshold `json:"threshold"`
}

type config struct {
	Filter filter `json:"filter"`
}

func (f filter) ContainsKey(s string) bool {
	_, ok := f.keysMap[s]
	return ok
}

func writeStdout(d interface{}) {
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	os.Stdout.WriteString("\n")
}

func toDatapoint(b []byte) (datapoint, error) {
	var d datapoint
	err := json.Unmarshal(b, &d)
	return d, err
}

func toEvent(b []byte) (event, error) {
	var e event
	err := json.Unmarshal(b, &e)
	return e, err
}

func toGeoPosition(b []byte) (geoPosition, error) {
	var g geoPosition
	err := json.Unmarshal(b, &g)
	return g, err
}

func toConfig(b []byte) (config, error) {
	var c config
	err := json.Unmarshal(b, &c)
	c.Filter.keysMap = make(map[string]struct{})
	for _, key := range c.Filter.Keys {
		var s struct{}
		c.Filter.keysMap[key] = s
	}
	return c, err
}

var conf config

func checkHighLimit(t threshold, v float64) {
	if high := conf.Filter.Threshold.High; high != nil {
		if v > *high {
			if e := conf.Filter.Threshold.Event; e != nil {
				writeStdout(e)
			}
		}
	}
}

func checkLowLimit(t threshold, v float64) {
	if low := conf.Filter.Threshold.High; low != nil {
		if v < *low {
			if e := conf.Filter.Threshold.Event; e != nil {
				writeStdout(e)
			}
		}
	}
}

func processDatapoint(d *datapoint) {
	if conf.Filter.ContainsKey(d.Key) {
		if t := conf.Filter.Threshold; t != nil {
			v := d.Value.(float64)
			checkHighLimit(*t, v)
			checkLowLimit(*t, v)
		}
	}
	writeStdout(d)
}

func readStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	b := scanner.Bytes()
	fmt.Println("first line:", string(b))

	var err error
	conf, err = toConfig(b)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		//fmt.Println("Filter type:", conf.Filter.Type)

	}
	for scanner.Scan() {
		b := scanner.Bytes()
		d, err := toDatapoint(b)
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println("key:", d.Key)
			processDatapoint(&d)
		}
	}
}

func main() {
	go readStdin()
	for {
	}
}
