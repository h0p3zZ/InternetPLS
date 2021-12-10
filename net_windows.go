//go:build windows

package main

import (
	"net"
)

func getDialInterfaceAddress() net.Addr {
	// addrs, _ := winipcfg.GetAdaptersAddresses(winipcfg.AddressFamily(2), winipcfg.GAAFlagIncludeAll)

	// for _, addr := range addrs {
	// 	if addr.DNSSuffix() == "htl.grieskirchen.local" {
	// 		ifaces, _ := net.Interfaces()
	// 		for _, i := range ifaces {
	// 			if addr.IfIndex == uint32(i.Index) {
	// 				localAddr, _ := i.Addrs()
	// 				return localAddr[0]
	// 			}
	// 		}
	// 	}
	// }
	return nil
}
