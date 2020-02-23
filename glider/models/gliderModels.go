package models

type TargetBackend struct {
	LbName string
	BackendName string
	BackendIP string
	BackendPort int
	LbIp string
	LbPort int
}

type IptableConstants struct {
	Prerouting string
	Postrouting string
	nat string
	dnat string
	snat string
}
