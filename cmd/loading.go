package cmd

import (
	"fmt"
	"time"
)

var isDone bool

func spinner(delay time.Duration, done chan bool) {
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	frames := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\r")
			return
		default:
			if isDone {
				fmt.Print("\r")
				done <- true
				return
			}
			fmt.Printf("\r%s Thinking...", frames[i])
			i = (i + 1) % len(frames)
			time.Sleep(delay)
		}
	}
}
