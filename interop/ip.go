package interop

import (
	"net"
)
import log "github.com/sirupsen/logrus"

func GetIPs(theInterface string) []net.IP {
	ifaces, _ := net.Interfaces()
	// handle err
	var ips []net.IP

	if theInterface != "" {
		log.Debugf("The interface given is: %s", theInterface)
	} else {
		log.Debug("No interface name provided, searching for any with an internal IP")
	}
	for _, i := range ifaces {
		addrs, _ := i.Addrs()

		if (theInterface == "") || i.Name == theInterface {
			log.Debugf("Checking IP address on interface: %s", i.Name)
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				/*if ip.To4() != nil {
					if ip.String() != "127.0.0.1" {
						log.Debugf("Found IP: %s", ip.String())
						ips = append(ips, ip.String())
					}
				}
				*/				
				if ip.To16() != nil {
					if ((ip.String() != "::1") && (ip.String() != "127.0.0.1")) {
						log.Debugf("Found IP: %s", ip.String())
						ips = append(ips, ip)

					}
				}

			}
		}
	}




	if len(ips) == 0 {
		log.Fatal("No IPs found")
	} else {
		if len(ips) > 1 {
			log.Debugf("More than one IP address found, please run with --interface to specify which interface to use")
		} else {
			log.Debugf("Just one IP found: %s", ips[0].String)
		}
	}
	return ips
}

func GetGlobalIPs(Ips []net.IP) []net.IP{
	var ips []net.IP

	for _, ip := range Ips {
		if (ip.IsGlobalUnicast()) {
			ips = append(ips, ip) 		
		}
	}

	return ips
}
