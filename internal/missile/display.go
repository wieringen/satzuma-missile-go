package missile

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type MissileDisplay struct {
	missileDevice MissileDevice
}

func NewMissileDisplay(missileDevice MissileDevice) *MissileDisplay {
	return &MissileDisplay{missileDevice: missileDevice}
}

func (md *MissileDisplay) Run() {
	for {
		md.clearScreen()
		md.displayInstructions()
		input := md.getUserInput()
		md.handleInput(input)
	}
}

func (md *MissileDisplay) clearScreen() {
	cmd := exec.Command("clear") // Use "cls" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (md *MissileDisplay) displayInstructions() {
	fmt.Println("MISSILE COMMAND!")
	fmt.Println("Use the following keys to move:")
	fmt.Println("Q   W   E")
	fmt.Println("  \\ | /")
	fmt.Println("A -(S)- D")
	fmt.Println("  / | \\")
	fmt.Println("Z   X   C")
	fmt.Println("F to fire all guns once")
	fmt.Println("R for rapid sequential fire of all missiles")
	fmt.Println("S to Stop")
	fmt.Println("ESC to Exit")
	fmt.Println("CAUTION: AIM AWAY FROM FACE")
}

func (md *MissileDisplay) getUserInput() string {
	var input string
	fmt.Print("Enter command: ")
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func (md *MissileDisplay) handleInput(input string) {
	switch input {
	case "w", "up":
		md.missileDevice.Move(UP)
	case "x", "down":
		md.missileDevice.Move(DOWN)
	case "a", "left":
		md.missileDevice.Move(LEFT)
	case "d", "right":
		md.missileDevice.Move(RIGHT)
	case "f", "space":
		md.missileDevice.Move(FIRE)
	case "s":
		md.missileDevice.Move(STOP)
	case "q":
		md.missileDevice.Move(LEFTUP)
	case "e":
		md.missileDevice.Move(RIGHTUP)
	case "z":
		md.missileDevice.Move(LEFTDOWN)
	case "c":
		md.missileDevice.Move(RIGHTDOWN)
	case "r":
		for i := 0; i < 3; i++ {
			md.missileDevice.Move(FIRE)
			time.Sleep(500 * time.Millisecond)
		}
	case "esc":
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
	}
}
