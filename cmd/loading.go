package cmd

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration, done chan bool) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	for {
		select {
		case <-done:
			fmt.Print("\r") // 清除动画
			return
		default:
			for _, frame := range frames {
				fmt.Printf("\r%s Thinking...", frame)
				time.Sleep(delay)
			}
		}
	}
}
