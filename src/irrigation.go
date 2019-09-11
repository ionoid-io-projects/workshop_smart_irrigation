package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	pin := rpio.Pin(21)
	pinR := rpio.Pin(11)

	pin.Input()
	pinR.Output()
	pin.PullUp()

	fmt.Println("Searching for water...")

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		pinR.Low()
		os.Exit(0)
	}()

	for {
		//		fmt.Println(pin.Read())
		if pin.Read() == 1 { // check if event occured
			fmt.Println("Water not detected....")
			time.Sleep(time.Second * 2)
			pinR.Toggle()
		} else {
			fmt.Println("Water detected....")
			pinR.Low()
		}
	}
}
