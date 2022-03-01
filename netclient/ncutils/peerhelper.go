package ncutils

import (
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// Returns the peers for a given wireguard interface or an
// empty array and a wgctrl defined error.
func GetPeers(iface string) ([]wgtypes.Peer, error) {
	// Create a client for communicating with wireguard
	c, err := wgctrl.New()
	if err != nil {
		return []wgtypes.Peer{}, err
	}
	// Get the wireguard network interface if it exists
	device, err := c.Device(iface)
	if err != nil {
		return []wgtypes.Peer{}, err
	}
	return device.Peers, nil
}
