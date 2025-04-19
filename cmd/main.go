package main

import (
	"fmt"
	"log"
	"os"

	"missile-go/internal/missile"
)

func main() {
	missileDevice := missile.NewMissileDevice()
	defer missileDevice.Close()

	if err := missileDevice.Open(); err != nil {
		log.Fatalf("Error opening missile device: %v", err)
	}

	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "--network", "-n":
			missile.RunNetworkMode(missileDevice)
		case "--console", "-c":
			missile.RunConsoleMode(missileDevice)
		case "--display", "-d":
			missile.NewMissileDisplay(*missileDevice).Run()
		default:
			fmt.Println("Unknown mode. Use --network, --display or --console.")
		}
	} else {
		fmt.Println("No mode specified. Use --network, --display or --console.")
	}
}
