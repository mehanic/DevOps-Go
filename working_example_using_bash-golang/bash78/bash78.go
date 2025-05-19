package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: sudo ./wpa2-wifi-connect <ssid> <pass>")
	}

	ssid := os.Args[1]
	pass := os.Args[2]

	// Disable wlan0 interface
	fmt.Println("Shutting down wlan0...")
	if err := exec.Command("ifdown", "wlan0").Run(); err != nil {
		log.Fatalf("Failed to ifdown wlan0: %v", err)
	}

	// Rebuild /etc/network/interfaces
	fmt.Println("Rebuilding /etc/network/interfaces...")
	if err := rebuildInterfaces(); err != nil {
		log.Fatalf("Failed to rebuild /etc/network/interfaces: %v", err)
	}

	// Rebuild /etc/wpa_supplicant/wpa_supplicant.conf
	fmt.Println("Rebuilding /etc/wpa_supplicant/wpa_supplicant.conf...")
	if err := rebuildWpaSupplicant(ssid, pass); err != nil {
		log.Fatalf("Failed to rebuild /etc/wpa_supplicant/wpa_supplicant.conf: %v", err)
	}

	// Enable wlan0 interface
	fmt.Println("Bringing up wlan0...")
	if err := exec.Command("ifup", "wlan0").Run(); err != nil {
		log.Fatalf("Failed to ifup wlan0: %v", err)
	}

	fmt.Println("WiFi connection setup complete.")
}

func rebuildInterfaces() error {
	interfacesFile := "/etc/network/interfaces"

	// Remove the file if it exists, and create a new one
	if err := os.Remove(interfacesFile); err != nil && !os.IsNotExist(err) {
		return err
	}

	file, err := os.Create(interfacesFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the interface configuration
	_, err = file.WriteString(
		`auto lo
iface lo inet loopback
iface eth0 inet dhcp
allow-hotplug wlan0
iface wlan0 inet manual
wpa-roam /etc/wpa_supplicant/wpa_supplicant.conf
iface default inet dhcp
`)
	return err
}

func rebuildWpaSupplicant(ssid, pass string) error {
	supplicantFile := "/etc/wpa_supplicant/wpa_supplicant.conf"

	// Remove the file if it exists, and create a new one
	if err := os.Remove(supplicantFile); err != nil && !os.IsNotExist(err) {
		return err
	}

	file, err := os.Create(supplicantFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the basic WPA supplicant config
	_, err = file.WriteString(
		`ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
`)
	if err != nil {
		return err
	}

	// Append the WPA passphrase (using wpa_passphrase system command)
	cmd := exec.Command("wpa_passphrase", ssid, pass)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to run wpa_passphrase: %v", err)
	}

	// Append the passphrase output to the file
	_, err = file.Write(output)
	return err
}

