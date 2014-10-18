package main

import (
	"fmt"
	"log"
	"net"

	"github.com/milosgajdos83/tenus"
)

func main() {
	veth, err := tenus.NewVethPairWithOptions("myveth01", tenus.VethOptions{PeerName: "myveth02"})
	if err != nil {
		log.Fatal(err)
	}

	vethHostIp, vethHostIpNet, err := net.ParseCIDR("10.0.41.2/16")
	if err != nil {
		log.Fatal(err)
	}

	if err := veth.SetLinkIp(vethHostIp, vethHostIpNet); err != nil {
		fmt.Println(err)
	}

	if err := veth.SetLinkMTU(1600); err != nil {
		fmt.Println(err)
	}
}
