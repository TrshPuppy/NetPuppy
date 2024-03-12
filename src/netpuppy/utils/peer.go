package utils

import (
	"fmt"
)

// Depending on user flags given, create this peer's type:
type Peer struct {
	ConnectionType string
	RPort          int
	LPort          int
	Address        string
	Connection     Socket
	Shell          bool
	ShellProcess   BashShell
	//ShellProcess *RealShell
	ReportTo string
}

// Create the peer struct which will guide the rest of the logic for NetPuppy:
func CreatePeer(port int, address string, listen bool, shell bool) *Peer {
	var thisPeer *Peer // POINTER: the functions which initialize thisPeer are returning addresses to the instances of Peer they create

	// If listen is true, we are the Offense peer (server):
	if listen { // Offense peer
		thisPeer = getOffense(port, shell)
	} else { // else, we are the Connect-back peer
		thisPeer = getConnectBack(port, address, shell)
	}

	fmt.Printf("Address of peer in peer.go %p\n", thisPeer)
	return thisPeer
}

// The Offense-type peer is the server and will listen on the any address:
func getOffense(port int, shell bool) *Peer {
	// Create offense-type instance of Peer:
	var offensePeer *Peer = &Peer{LPort: port, Address: "0.0.0.0", ConnectionType: "offense", Shell: shell}

	fmt.Printf("Address of peer in get peer peer.go: %p\n", offensePeer)
	return offensePeer // POINTER: return the address of the offensePeer instance (instead of a copy)
}

// The Connect-back peer is the client & will connect to the given IP address & port:
func getConnectBack(port int, address string, shell bool) *Peer {
	// Create connect-back instance of Peer:
	var connectBackPeer *Peer = &Peer{RPort: port, Address: address, ConnectionType: "connect-back", Shell: shell}

	fmt.Printf("Address of peer in getpeer peer.go: %p\n", connectBackPeer)
	//pointerToPeer := &connectBackPeer
	return connectBackPeer // POINTER: return the address of the connectBackPeer instance (instead of copy)
}
