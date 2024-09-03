package main

import (
	"Chiffre_de_Vigenere/internal/dialog"
	"os"
	"os/exec"
	"runtime"
)

var clearOS map[string]func()

func main() {
	clearTerminal()

	dialog.Dialog()
}

func clearTerminal() {
	if value, ok := clearOS[runtime.GOOS]; ok {
		value()
	}
}

func init() {
	clearOS = make(map[string]func())
	clearOS["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearOS["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearOS["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
