package models

type TargetBackendHolder struct {
	LbName string
	BackendName string
	BackendIP string
	BackendPort int
	LbIp string
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
	bind string
	backend string
}

type config struct {

}
