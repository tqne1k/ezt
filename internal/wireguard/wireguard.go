package wireguard

import (
	"eztrust/bootstrap"
	"net"
	"os/exec"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func GeneratePrivateKey() string {
	// Generate a new private key
	privateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		bootstrap.Logger.Err(err).Msg("Failed to generate wireguard private key")
		return ""
	}
	return privateKey.String()
}

func GeneratePublicKey(privateKey string) string {
	// Parse the private key and generate the public key
	key, err := wgtypes.ParseKey(privateKey)
	if err != nil {
		bootstrap.Logger.Err(err).Msg("Failed to parse wireguard private key")
		return ""
	}
	return key.PublicKey().String()
}

func CreateTunnel(interfaceName string, strPrivateKey string, listenPort int, allowNetwork string) error {
	env := bootstrap.NewEnv()
	bootstrap.Logger.Info().Msg("Creating wireguard tunnel")
	client, err := wgctrl.New()
	if err != nil {
		bootstrap.Logger.Err(err).Msg("Failed to create wireguard client")
		return err
	}
	defer client.Close()

	// Parse the private key and generate the public key
	privateKey, err := wgtypes.ParseKey(strPrivateKey)
	if err != nil {
		bootstrap.Logger.Err(err).Msg("Failed to parse wireguard private key")
		return err
	}
	publicKey := privateKey.PublicKey()

	// Convert the allowed network to IPNet
	_, allowedNetwork, err := net.ParseCIDR(allowNetwork)
	if err != nil {
		bootstrap.Logger.Err(err).Msg("Failed to parse allowed network")
		return err
	}

	// Generate the configuration
	wgInterface := wgtypes.Config{
		PrivateKey: &privateKey,
		ListenPort: &listenPort,
		Peers: []wgtypes.PeerConfig{
			{
				PublicKey: publicKey,
				AllowedIPs: []net.IPNet{
					{
						IP:   allowedNetwork.IP,
						Mask: net.CIDRMask(env.Netcidr, 32),
					},
				},
			},
		},
	}

	// Create interface before configuring it
	err = CreateInterface(interfaceName)
	if err != nil {
		return err
	}

	// Configure the interface
	err = client.ConfigureDevice(interfaceName, wgInterface)
	if err != nil {
		bootstrap.Logger.Err(err).Msg("Failed to configure wireguard interface")
		return err
	}

	return nil
}

func CreateInterface(interfaceName string) error {
	cmd := exec.Command("ip", "link", "add", "dev", interfaceName, "type", "wireguard")
	err := cmd.Run()
	if err != nil {
		bootstrap.Logger.Err(err).Msg("Failed to create wireguard interface")
		return err
	}
	return nil
}

func DeleteInterface(interfaceName string) error {
	cmd := exec.Command("sudo", "ip", "link", "delete", "dev", interfaceName)
	err := cmd.Run()
	if err != nil {
		bootstrap.Logger.Err(err).Msg("Failed to delete wireguard interface")
		return err
	}
	return nil
}

func GetPeerStatusRealtime(interfaceName string) (string, error) {
	cmd := exec.Command("wg", "show", interfaceName)
	out, err := cmd.Output()
	if err != nil {
		bootstrap.Logger.Err(err).Msg("Failed to get wireguard peer status")
		return "", err
	}
	return string(out), nil
}
