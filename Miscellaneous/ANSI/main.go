package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	fmt.Print("\033[?25l") // make cursor invisible
}

func exit() {
	fmt.Print("\033[?25h") // make cursor visible
}

func clearScreen() {
	fmt.Print("\033[2J")
}

func printAt(s string, x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
	fmt.Println(s)
}

func setColor(f, b int) {
	fmt.Printf("\033[38;5;%dm", f) // foreground color
	fmt.Printf("\033[48;5;%dm", b) // background color
}

func main() {
	defer exit()

	clearScreen()

	printAt("Cat", 2, 1)

	setColor(0, 150)

	printAt("Dog", 4, 5)

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
