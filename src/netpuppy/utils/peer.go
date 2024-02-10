package utils

import (
	"os/exec"
)

// Depending on input, create this peer's type:
type Peer struct {
	ConnectionType string
	RPort          int
	LPort          string
	Address        string
	Connection     Socket
	Shell          bool
	ShellProcess   *exec.Cmd
}

func CreatePeer(port int, address string, listen bool, shell bool) Peer {
	var thisPeer Peer

	if listen {
		thisPeer = getOffense(port, shell)
	} else {
		thisPeer = getConnectBack(port, address, shell)
	}
	return thisPeer
}

func getOffense(port int, shell bool) Peer {
	offensePeer := Peer{RPort: port, Address: "0.0.0.0", ConnectionType: "offense", Shell: shell}
	return offensePeer
}

func getConnectBack(port int, address string, shell bool) Peer {
	connectBackPeer := Peer{RPort: port, Address: address, ConnectionType: "connect_back", Shell: shell}
	return connectBackPeer
}