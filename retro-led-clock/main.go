package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	for shift := 0; ; shift++ {
		screen.Clear()
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]led{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			n := len(clock)

			s, e := shift%n, n

			if shift%(n*2) >= n { // as if the clock is double of its length.
				s, e = 0, s // n != e
			}

			for j := 0; j < n-e; j++ {
				fmt.Print("   ")
			}

			for i := s; i < e; i++ { // from 's' to 'e'
				next := clock[i][line]
				fmt.Print(next, " ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second * 2)
	}
}
