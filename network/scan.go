package network

import (
    // "fmt"
    // "log"
    //"os"

    // "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)

type DeviceAddress struct {
	IPAddress string
	SubnetMask string
}

type NetworkDevice struct {
	Name string
	Description string
	Addresses []DeviceAddress
}

type Scanner struct {
}

func (s* Scanner) GetNetworkDevices() (*[]NetworkDevice, error) {
	result := []NetworkDevice{}

	devices,err := pcap.FindAllDevs()
	if err != nil {
		return nil, err
	}

	for _, device := range devices {
		addresses := []DeviceAddress{}

		for _, address := range device.Addresses {
			addresses = append(addresses, DeviceAddress{
				IPAddress: address.IP.String(),
				SubnetMask: address.Netmask.String(),
			})
		}

		nwd := NetworkDevice{
			Name: device.Name,
			Description: device.Description,
			Addresses: addresses,
		}

		result = append(result, nwd)
	}

	return &result, nil
}

/*
func main() {
    // Find all network devices on the system
    devices, err := pcap.FindAllDevs()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Devices found:")
    for _, device := range devices {
        fmt.Println("\nName: ", device.Name)
        fmt.Println("Description: ", device.Description)
        for _, address := range device.Addresses {
            fmt.Println("- IP address: ", address.IP)
            fmt.Println("- Subnet mask: ", address.Netmask)
        }
    }

    // Select the first device found
    deviceName := devices[0].Name
    fmt.Printf("\nStarting capture on device: %s\n", deviceName)

    // Open the device for capturing
    handle, err := pcap.OpenLive(deviceName, 1024, false, pcap.BlockForever)
    if err != nil {
        log.Fatal(err)
    }
    defer handle.Close()

    // Set up packet capture
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        fmt.Println(packet)
    }
}
*/

