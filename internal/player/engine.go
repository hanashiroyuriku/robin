package player

import (
	"fmt"
	"time"
)

func PlayLysrics(text string, charDelay float64) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(time.Duration(charDelay * float64(time.Second)))
	}
	fmt.Println()
}