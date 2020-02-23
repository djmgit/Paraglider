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
	TImeOut int
}
