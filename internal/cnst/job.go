package cnst

//go:generate enumer --type=Job --linecomment --extramethod --output=job_enum.go
type Job uint

const (
	Job_Demo Job = iota + 1 // DEMO
)
