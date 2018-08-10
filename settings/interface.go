package settings

import (
	"bufio"
	"os"
	"fmt"
)

var InterfacePath = "/etc/network/interfaces"
var InterfaceBack = "/etc/network/interface.bak"

var templateForReplacement = []string {
	"        address ",
	"        netmask ",
	"        network ",
	"        broadcast ",
	"        gateway ",
}

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

func replaceValue(des string, i int) {
	if des!="" {
		templateForInterface[i+7] = templateForReplacement[i] + des
	}
}

func GenerateTemplateBackup(args ...string) error {
	for i, v := range args {
		replaceValue(v, i)
	}

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

func ModifyInterface() error {
	return os.Rename(InterfaceBack, InterfacePath)
}