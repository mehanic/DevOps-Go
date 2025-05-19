package main

import (
	"fmt"
	"os"
	"os/exec"
//	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: sudo go run wpa2_wifi_connect.go <ssid> <pass>")
		return
	}

	ssid := os.Args[1]
	password := os.Args[2]

	// Disable wlan0
	if err := exec.Command("ifdown", "wlan0").Run(); err != nil {
		fmt.Printf("Failed to bring down wlan0: %v\n", err)
		return
	}

	// Create /etc/network/interfaces
	interfacesFile := "/etc/network/interfaces"
	if err := os.Remove(interfacesFile); err != nil && !os.IsNotExist(err) {
		fmt.Printf("Failed to remove %s: %v\n", interfacesFile, err)
		return
	}

	// Create and write to the interfaces file
	file, err := os.Create(interfacesFile)
	if err != nil {
		fmt.Printf("Failed to create %s: %v\n", interfacesFile, err)
		return
	}
	defer file.Close()

	interfaceConfig := []string{
		"auto lo",
		"iface lo inet loopback",
		"iface eth0 inet dhcp",
		"allow-hotplug wlan0",
		"iface wlan0 inet manual",
		"wpa-roam /etc/wpa_supplicant/wpa_supplicant.conf",
		"iface default inet dhcp",
	}
	for _, line := range interfaceConfig {
		if _, err := file.WriteString(line + "\n"); err != nil {
			fmt.Printf("Failed to write to %s: %v\n", interfacesFile, err)
			return
		}
	}

	// Create /etc/wpa_supplicant/wpa_supplicant.conf
	supplicantFile := "/etc/wpa_supplicant/wpa_supplicant.conf"
	if err := os.Remove(supplicantFile); err != nil && !os.IsNotExist(err) {
		fmt.Printf("Failed to remove %s: %v\n", supplicantFile, err)
		return
	}

	// Create and write to the supplicant file
	file, err = os.Create(supplicantFile)
	if err != nil {
		fmt.Printf("Failed to create %s: %v\n", supplicantFile, err)
		return
	}
	defer file.Close()

	supplicantConfig := []string{
		"ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev",
		"update_config=1",
	}

	for _, line := range supplicantConfig {
		if _, err := file.WriteString(line + "\n"); err != nil {
			fmt.Printf("Failed to write to %s: %v\n", supplicantFile, err)
			return
		}
	}

	// Append the WPA passphrase to the supplicant file
	passphraseCmd := exec.Command("wpa_passphrase", ssid, password)
	output, err := passphraseCmd.Output()
	if err != nil {
		fmt.Printf("Failed to run wpa_passphrase: %v\n", err)
		return
	}

	// Write the passphrase output to the supplicant file
	if _, err := file.Write(output); err != nil {
		fmt.Printf("Failed to write WPA passphrase to %s: %v\n", supplicantFile, err)
		return
	}

	// Bring wlan0 back up
	if err := exec.Command("ifup", "wlan0").Run(); err != nil {
		fmt.Printf("Failed to bring up wlan0: %v\n", err)
		return
	}

	fmt.Println("Wi-Fi configuration applied successfully.")
}

