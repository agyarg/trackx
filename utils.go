package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearTerminal clears the terminal screen
func ClearTerminal() {
	switch runtime.GOOS {
	case "linux", "darwin": // Linux or macOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported platform. Cannot clear terminal.")
	}
}
