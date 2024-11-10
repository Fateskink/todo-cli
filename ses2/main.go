package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go prompter()

	for {
		fmt.Print("App is running...")

		select {
		case res := <-sig:
			fmt.Printf("\n%s signal recieved\nGoood bye!", res)
			signal.Stop(sig)
			os.Exit(0)
		}
	}
}

func prompter() {
	fmt.Print("\n>> ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Printf("<- %s\n", line)
		fmt.Print(">> ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(0)
	}
}
