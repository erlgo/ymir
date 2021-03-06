package model

import "time"

// TWork and TWorkInstance saved as config map

// TWork is a test Work
type TWork struct {
	JobName  string
	WorkID   string
	Replicas int
	Result   []TResult
	Instance []TWorkInstance
	Created  time.Time
}

// TWorkInstance in a single node....
type TWorkInstance struct {
	JobName    string
	WorkID     string
	InstanceID string
	NodeName   string
	Created    time.Time
	Results    []TResult
}

// TResult ...
type TResult struct {
	Name       string
	Start      time.Time
	End        time.Time
	Count      int64
	Sum        time.Duration
	Max        time.Duration
	Avg        time.Duration
	Min        time.Duration
	CodeMap    map[int]int64 // code counter
	Percentile []Bucket
	Buckets    []Bucket
	Lats       []float64
}

// Bucket ...
type Bucket struct {
	Count int
	Cost  time.Duration
}

// TReqResult ...
type TReqResult struct {
	Name string
	Time time.Time
	Cost time.Duration
	Code int
}
