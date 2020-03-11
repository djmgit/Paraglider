package glidercore

import (
	"github.com/djmgit/go-iptables/iptables"
	"paraglider/glider/models"
	"time"
)

var iptableConstants = models.IptableConstantsHolder{
	Prerouting: "PREROUTING",
	Postrouting: "POSTROUTING",
	Nat: "nat",
	Dnat: "DNAT",
	Snat: "SNAT",
	Tcp: "tcp",
}

func CreateTargetForLb(TargetBackend models.TargetBackendHolder) error {

	ipt, err := iptables.New()
	if err != nil {
		return err
	}
	err = ipt.AppendUnique(iptableConstants.Nat, iptableConstants.Prerouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.LbIP, "--dport", string(TargetBackend.LbPort), "-j", iptableConstants.Dnat, "--to-destination", TargetBackend.BackendIP + ":" + string(TargetBackend.BackendPort))

	if err != nil {
		return err
	}

	err = ipt.AppendUnique(iptableConstants.Nat, iptableConstants.Postrouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.BackendIP, "--dport", string(TargetBackend.BackendPort), "-j", iptableConstants.Snat, "--to-source", TargetBackend.LbPrivateIP)

	if err != nil {

		_ = ipt.Delete(iptableConstants.Nat, iptableConstants.Prerouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.LbIP, "--dport", string(TargetBackend.LbPort), "-j", iptableConstants.Dnat, "--to-destination", TargetBackend.BackendIP + ":" + string(TargetBackend.BackendPort))

		return err
	}

	return nil
}

func RemoveTargetForLb(TargetBackend models.TargetBackendHolder) error {

	ipt, err := iptables.New()
	if err != nil {
		return err
	}

	err = ipt.Delete(iptableConstants.Nat, iptableConstants.Prerouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.LbIP, "--dport", string(TargetBackend.LbPort), "-j", iptableConstants.Dnat, "--to-destination", TargetBackend.BackendIP + ":" + string(TargetBackend.BackendPort))

	if err != nil {
		return  err
	}

	time.Sleep(5 * time.Second)

	err = ipt.Delete(iptableConstants.Nat, iptableConstants.Postrouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.BackendIP, "--dport", string(TargetBackend.BackendPort), "-j", iptableConstants.Snat, "--to-source", TargetBackend.LbPrivateIP)

	if err != nil {
		return err
	}

	return nil
}
