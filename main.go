package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func LineMoveUp(n int) {
	fmt.Printf("\033[%dA", n)
}

func HideCursor() {
	fmt.Printf("\033[?25l")
}

func ShowCursor() {
	fmt.Printf("\033[?25h")
}

// https://en.wikipedia.org/wiki/ANSI_escape_code
// CSI (Control Sequence Introducer) sequences
func main() {
	for true {
		out, err := exec.Command("nvidia-smi").Output()
		if err != nil {
			log.Fatal(err)
		}

		HideCursor()
		scanner := bufio.NewScanner(bytes.NewReader(out))
		line := 0
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
			line++
			if line > 12 {
				break
			}
		}
		LineMoveUp(line)
		time.Sleep(150 * time.Millisecond)
		ShowCursor()
	}
}
