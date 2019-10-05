package main

import (
    "fmt"
    "github.com/vishvananda/netlink"
    "golang.zx2c4.com/wireguard/wgctrl"
    "golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {
    la := netlink.NewLinkAttrs()
    la.Name = "wg1"
    mydev := &netlink.GenericLink{LinkAttrs: la, LinkType: "wireguard"}
    err := netlink.LinkAdd(mydev)
    if err != nil  {
        fmt.Printf("could not add %s: %v\n", la.Name, err)
    }
    // TODO 
    // - Generate private key
    // - Generate public key
    // - setup device
    // - ip addr add
    // - ip link up
    // - ip route
    // - setup peer
    err = netlink.LinkDel(mydev)
    if err != nil  {
        fmt.Printf("could not del %s: %v\n", la.Name, err)
    }
}
