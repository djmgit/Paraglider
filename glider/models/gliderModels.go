package models
import "gopkg.in/yaml.v2"

type TargetBackendHolder struct {
	//LbName string
	//BackendName string
	BackendIP string
	BackendPort int
	LbIP string
	LbPort int
	LbPrivateIP string
}

type IptableConstantsHolder struct {
	Prerouting string
	Postrouting string
	Nat string
	Dnat string
	Snat string
	Tcp string
}

type TcpHealthCheckHolder struct {
	TargetIP string
	TargetPort int
	Threshold int
	TimeOut int
}

type HTTPHealthCheckHolder struct {
	TargetIP string
	TargetPort int
	Threshold int
	TimeOut int
	HealthCheckPath string
	PassMessage string
}

type BackendMemberHolder struct {
	Name string
	Members []string
}

type FrontendHolder struct {
	Name string
	Bind string
	PrivateIP string
	Backends []string
}

type Config struct {
	Frontend FrontendHolder
}

func (c *Config) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}
