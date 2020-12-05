package netutils

type TargetType string

const (
	Undefined TargetType = ""
	Port      TargetType = "port"
	NamedPipe TargetType = "named_pipe"
)

type ConnectionTarget struct {
	TargetType TargetType
	Port       uint16
	NamedPipe  string
}
