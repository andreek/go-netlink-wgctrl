package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {

	// setup wireguard generic link
	la := netlink.NewLinkAttrs()
	la.Name = "wg1"
	wgdev := &netlink.GenericLink{LinkAttrs: la, LinkType: "wireguard"}
	err := netlink.LinkAdd(wgdev)
	if err != nil  {
		fmt.Printf("could not add %s: %v\n", la.Name, err)
	}

	// generate private key
	privateKey, perr := wgtypes.GeneratePrivateKey()
	if perr != nil  {
		fmt.Printf("could not generate private key %s: %v\n", privateKey, err)
	}

	// print out public key
	publicKey := privateKey.PublicKey()
	fmt.Printf("public key: %s\n", publicKey.String())

	// TODO 
	// - setup device
	// - ip addr add
	// - ip link up
	// - ip route
	// - setup peer
	err = netlink.LinkDel(wgdev)
	if err != nil  {
		fmt.Printf("could not del %s: %v\n", la.Name, err)
	}
}
