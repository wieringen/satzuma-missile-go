package missile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunConsoleMode(missileDevice *MissileDevice) {
	fmt.Println("Console Mode: Enter commands (w/a/s/d/f/r/z/c/q/e/v/esc):")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch strings.ToLower(input) {
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
			fmt.Println("Exiting console mode...")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
