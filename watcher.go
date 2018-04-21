package socks5

import (
	"fmt"
)

func watch(ctx *AuthContext, writtenCh chan int64) {
	for i := 0; i < 2; i++ {
		w := float64(<-writtenCh) / 1024

		if u, ok := ctx.Payload["Username"]; ok {
			fmt.Printf("[WATCH] %v: %0.4f KB\n", u, w)
		} else {
			fmt.Printf("[WATCH] %0.4f KB\n", w)
		}
	}
}
