package glidercore

import (
	"github.com/djmgit/go-iptables/iptables"
	"paraglider/glider/models"
)



func CreateTargetForLb(TargetBackend models.TargetBackendHolder) error {

	err := iptables.AppendUnique()
}
