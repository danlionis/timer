package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	start := time.Now()
	defer finished(start)
	c := make(chan os.Signal)

	// If the programm crashes or gets cancled stop the timer
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		finished(start)
		os.Exit(1)
	}()

	cmd.Run()
}

func finished(start time.Time) {
	end := time.Since(start)
	fmt.Printf("\n\ntime\t%s\n", end)
}

func printHelp() {
	fmt.Fprintf(os.Stderr, "\nUsage: %s <command>\n\n", os.Args[0])
}
