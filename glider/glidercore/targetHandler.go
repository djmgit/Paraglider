package glidercore

import (
	"github.com/djmgit/go-iptables/iptables"
	"paraglider/glider/models"
	"strconv"
)

// Struct containing list of constants for iptable rules
var iptableConstants = models.IptableConstantsHolder{
	Prerouting: "PREROUTING",
	Postrouting: "POSTROUTING",
	Nat: "nat",
	Dnat: "DNAT",
	Snat: "SNAT",
	Tcp: "tcp",
}

func CreateTargetForLb(TargetBackend models.TargetBackendHolder, roundRobinTurn int) error {

	ipt, err := iptables.New()
	if err != nil {
		return err
	}

	// Add the PREROUTING rule for the given backend
	if roundRobinTurn != 1 {
		err = ipt.AppendUnique(iptableConstants.Nat, iptableConstants.Prerouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.LbIP, "--dport", strconv.Itoa(TargetBackend.LbPort), "-m", "statistic", "--mode", "nth", "--every", strconv.Itoa(roundRobinTurn), "--packet", "0", "-j", iptableConstants.Dnat, "--to-destination", TargetBackend.BackendIP + ":" + strconv.Itoa(TargetBackend.BackendPort))
	} else {
		err = ipt.AppendUnique(iptableConstants.Nat, iptableConstants.Prerouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.LbIP, "--dport", strconv.Itoa(TargetBackend.LbPort), "-j", iptableConstants.Dnat, "--to-destination", TargetBackend.BackendIP + ":" + strconv.Itoa(TargetBackend.BackendPort))
	}

	if err != nil {
		return err
	}

	// Add POSTROUTING rule for the given target only after PREROUTING rule has been added
	err = ipt.AppendUnique(iptableConstants.Nat, iptableConstants.Postrouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.BackendIP, "--dport", strconv.Itoa(TargetBackend.BackendPort), "-j", iptableConstants.Snat, "--to-source", TargetBackend.LbPrivateIP)

	if err != nil {

		if roundRobinTurn != 1 {
			_ = ipt.Delete(iptableConstants.Nat, iptableConstants.Prerouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.LbIP, "--dport", strconv.Itoa(TargetBackend.LbPort), "-m", "statistic", "--mode", "nth", "--every", strconv.Itoa(roundRobinTurn), "--packet", "0", "-j", iptableConstants.Dnat, "--to-destination", TargetBackend.BackendIP + ":" + strconv.Itoa(TargetBackend.BackendPort))
		} else {
			_ = ipt.Delete(iptableConstants.Nat, iptableConstants.Prerouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.LbIP, "--dport", strconv.Itoa(TargetBackend.LbPort), "-j", iptableConstants.Dnat, "--to-destination", TargetBackend.BackendIP + ":" + strconv.Itoa(TargetBackend.BackendPort))
		}

		return err
	}

	return nil
}

func RemoveTargetForLb(TargetBackend models.TargetBackendHolder, roundRobinTurn int) error {

	ipt, err := iptables.New()
	if err != nil {
		return err
	}

	// Remove the PREROUTING rule
	if roundRobinTurn != 1 {
		err = ipt.Delete(iptableConstants.Nat, iptableConstants.Prerouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.LbIP, "--dport", strconv.Itoa(TargetBackend.LbPort), "-m", "statistic", "--mode", "nth", "--every", strconv.Itoa(roundRobinTurn), "--packet", "0", "-j", iptableConstants.Dnat, "--to-destination", TargetBackend.BackendIP + ":" + strconv.Itoa(TargetBackend.BackendPort))
	} else {
		err = ipt.Delete(iptableConstants.Nat, iptableConstants.Prerouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.LbIP, "--dport", strconv.Itoa(TargetBackend.LbPort), "-j", iptableConstants.Dnat, "--to-destination", TargetBackend.BackendIP + ":" + strconv.Itoa(TargetBackend.BackendPort))
	}

	if err != nil {
		return  err
	}

	// Delete the coresponding POSTROUTING rule
	err = ipt.Delete(iptableConstants.Nat, iptableConstants.Postrouting, "-p", iptableConstants.Tcp, "-d", TargetBackend.BackendIP, "--dport", strconv.Itoa(TargetBackend.BackendPort), "-j", iptableConstants.Snat, "--to-source", TargetBackend.LbPrivateIP)

	if err != nil {
		return err
	}

	return nil
}
