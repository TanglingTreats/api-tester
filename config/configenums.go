package config

import (
	"encoding/json"
	"fmt"
	"strings"
)

// --- Step Type enum ---
type StepType int

const (
	Sequential StepType = iota + 1
	Concurrent
)

var (
	StepTypeName = map[StepType]string{
		Sequential: "sequential", Concurrent: "concurrent",
	}
	StepTypeValue = map[string]StepType{
		"sequential": Sequential, "concurrent": Concurrent,
	}
)

func parseStepType(s string) (StepType, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := StepTypeValue[s]
	if !ok {
		return StepType(0), fmt.Errorf("%q is not a valid step type", s)
	}
	return StepType(value), nil
}

func (st StepType) String() string {
	return StepTypeName[st]
}

func (st StepType) MarshalJSON() ([]byte, error) {
	return json.Marshal(st.String())
}

func (st *StepType) UnmarshalJSON(data []byte) (err error) {
	var stepType string
	if err := json.Unmarshal(data, &stepType); err != nil {
		return err
	}
	if *st, err = parseStepType(stepType); err != nil {
		return err
	}
	return nil
}

// --- --- ---
// --- Measurement ---
type Measurement int

const (
	ResponseTime Measurement = iota
	BlockingDuration
)

var (
	MeasurementName = map[Measurement]string{
		ResponseTime: "response time", BlockingDuration: "blocking duration",
	}

	MeasurementValue = map[string]Measurement{
		"response time": ResponseTime, "blocking duration": BlockingDuration,
	}
)

func parseMeasurement(s string) (Measurement, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := StepTypeValue[s]
	if !ok {
		return Measurement(0), fmt.Errorf("%q is not a valid step type", s)
	}
	return Measurement(value), nil
}

func (m Measurement) String() string {
	return MeasurementName[m]
}

func (m Measurement) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *Measurement) UnmarshalJSON(data []byte) (err error) {
	var measurement string
	if err := json.Unmarshal(data, &measurement); err != nil {
		return err
	}
	if *m, err = parseMeasurement(measurement); err != nil {
		return err
	}
	return nil
}

// --- --- ---

type TestType string

const (
	Sequence TestType = "sequence"
	Parallel TestType = "parallel"
)

type HttpMethod string

const (
	GET  HttpMethod = "GET"
	POST HttpMethod = "POST"
	PUT  HttpMethod = "PUT"
)

type Step struct {
	Name      string      `json:"name"`
	Endpoint  string      `json:"endpoint"`
	Method    HttpMethod  `json:"method"`
	StepType  StepType    `json:"stepType"`
	TestType  TestType    `json:"testType"`
	Load      int         `json:"load"`
	DependsOn []string    `json:"dependsOn"`
	Measure   Measurement `json:"measure"`
	Threshold string      `json:"threshold"`
}

type Config struct {
	Version int    `json:"version"`
	URL     string `json:"url"`
	Steps   []Step `json:"steps"`
}
