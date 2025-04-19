package missile

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func RunNetworkMode(missileDevice *MissileDevice) {
	addr := ":8080"
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatalf("Failed to start UDP server: %v", err)
	}
	defer conn.Close()

	fmt.Printf("Listening for UDP commands on %s...\n", addr)
	buf := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("Error reading from UDP: %v", err)
			continue
		}

		cmd := strings.TrimSpace(string(buf[:n]))
		fmt.Printf("Received command: %s\n", cmd)

		switch strings.ToLower(cmd) {
		case "w", "up":
			missileDevice.Move(UP)
		case "x", "down":
			missileDevice.Move(DOWN)
		case "a", "left":
			missileDevice.Move(LEFT)
		case "d", "right":
			missileDevice.Move(RIGHT)
		case "f", "space":
			missileDevice.Move(FIRE)
		case "s":
			missileDevice.Move(STOP)
		case "q":
			missileDevice.Move(LEFTUP)
		case "e":
			missileDevice.Move(RIGHTUP)
		case "z":
			missileDevice.Move(LEFTDOWN)
		case "c":
			missileDevice.Move(RIGHTDOWN)
		case "r":
			for i := 0; i < 3; i++ {
				missileDevice.Move(FIRE)
			}
		case "v":
			missileDevice.Move(FIRE)
		case "esc":
			fmt.Println("Exiting network mode...")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
