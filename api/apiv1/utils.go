package apiv1

import (
	"bufio"
	"os"
	"fmt"
)

var InterfacePath = "/etc/network/interfaces"
var InterfaceBack = "/etc/network/interfaces.bak"

var templateForReplacement = "        address "
var templateForInterface = []string{
	"# Configure Loopback",
	"auto lo",
	"iface lo inet loopback",
	"",
	"# Configure eth0",
	"auto eth0",
	"iface eth0 inet static",
	"        address 192.168.1.253",
	"        netmask 255.255.255.0",
	"        network 192.168.1.0",
	"        broadcast 192.168.255.255",
	"        gateway 192.168.1.1",
}

func generateTemplateBackup(desip string) error {
	templateForInterface[7] = templateForReplacement + desip
	file, err := os.OpenFile(InterfaceBack, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
	   return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range templateForInterface {
	   fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func modifyInterface() error {
	return os.Rename(InterfaceBack, InterfacePath)
}