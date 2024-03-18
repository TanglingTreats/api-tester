package config

type StepType string

const (
	Sequential StepType = "sequential"
	Concurrent StepType = "concurrent"
)

type Measurement string

const (
	ResponseTime     Measurement = "response time"
	BlockingDuration Measurement = "blocking duration"
)

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
	StepIndex int         `json:"stepIndex"`
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
