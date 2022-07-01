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
		hr, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]led{
			digits[hr/10], digits[hr%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			n := len(clock)

			// s, e := shift%n, n

			// if shift%(n*2) >= n { // as if the clock is double of its length.
			// 	s, e = 0, s // n != e
			// }
			var s, e int

			if shift%(2*n) < n {
				s, e = shift%n, n
			} else {
				s, e = 0, shift%n
			}

			for j := 0; j < n-e; j++ {
				fmt.Print("   ")
			}

			for i := s; i < e; i++ { // from 's' to 'e'
				fmt.Printf("%v ", clock[i][line])
			}
			fmt.Println()
		}
		time.Sleep(time.Second * 2)
	}
}
