//ding shi renwu zhixing config from renwuconfigjson

package main

import (
	//"encoding/gob"
	//"flag"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

/*
cat /etc/resolv.conf
nameserver 61.144.56.100
nameserver 61.144.56.101
*/

const (
	dns1       = "61.144.56.100"
	dns2       = "61.144.56.101"
	configName = "eth.json"
)

type EthAddr struct {
	Ip      string
	Gateway string
	Dns     string
	Ether   string
}

func (eth *EthAddr) new() *EthAddr {
	eth.Dns = "192.168.1.1"
	eth.Gateway = "192.168.1.1"
	eth.Ether = "eth0"
	eth.Ip = "192.168.1.222"

	fmt.Println("new")
	return eth
}

func (eth *EthAddr) Read() {

	fmt.Println("Read")
}

func (eth *EthAddr) ReadFromConfig() {
	b1, _ := ioutil.ReadFile(configName)
	json.Unmarshal(b1, eth)
	fmt.Println(*eth)
	fmt.Println("ReadFromConfig")
}

func (eth *EthAddr) Write() {

	fmt.Println("Write")
}

func (eth *EthAddr) Save() {
	f1, _ := os.OpenFile(configName, syscall.O_CREAT|syscall.O_RDWR, os.ModePerm)
	defer f1.Close()
	b1, _ := json.Marshal(eth)
	f1.Write(b1)
	fmt.Println("Save")
}

func (eth *EthAddr) Config() {

	fmt.Println("Config")
}

func (cli *EthAddr) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  read -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  readFromConfig - Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  write  -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  save   - Lists all addresses from the wallet file")
	fmt.Println("  config - Lists all addresses from the wallet file")
}

func (cli *EthAddr) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func main() {

	var renwu = new(EthAddr)
	renwu.new()

	renwu.validateArgs()
	// ReadCmd := flag.NewFlagSet("Read", flag.ExitOnError)
	// ReadFromConfigCmd := flag.NewFlagSet("ReadFromConfig", flag.ExitOnError)
	// WriteCmd := flag.NewFlagSet("Write", flag.ExitOnError)
	// listaddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)
	// SaveCmd := flag.NewFlagSet("Save", flag.ExitOnError)
	// ConfigCmd := flag.NewFlagSet("Config", flag.ExitOnError)
	switch os.Args[1] {
	case "read":
		renwu.Read()
		break

	case "readFromConfig":
		renwu.ReadFromConfig()
		break
	case "write":
		renwu.Write()
		break
	case "save":
		renwu.Save()
		break
	case "config":
		renwu.Config()
		break
	}
}
