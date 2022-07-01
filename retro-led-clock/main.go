package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

/*
0	1	2	3	4	5	6	7		b	s	e	sft	  sft%n	sft%2n>n	  s		  e
A	B	C	D	E	F	G	H		0	0	7	16		0		F		sft%n	  n
B	C	D	E	F	G	H			0	1	7	17		1		F		sft%n	  n
C	D	E	F	G	H				0	2	7	18		2		F		sft%n  	  n
D	E	F	G	H					0	3	7	19		3		F		sft%n	  n
E	F	G	H						0	4	7	20		4		F		sft%n	  n
F	G	H							0	5	7	21		5		F		sft%n	  n
G	H								0	6	7	22		6		F		sft%n	  n
H									0	7	7	23		7		F		sft%n	  n
							A		7	0	0	24		0		F		  0		sft%n
						A	B		6	0	1	25		1		T		  0		sft%n
					A	B	C		5	0	2	26		2		T		  0		sft%n
				A	B	C	D		4	0	3	27		3		T		  0		sft%n
			A	B	C	D	E		3	0	4	28		4		T		  0		sft%n
		A	B	C	D	E	F		2	0	5	29		5		T		  0		sft%n
	A	B	C	D	E	F	G		1	0	6	30		6		T		  0		sft%n
A	B	C	D	E	F	G	H		0	0	7	31		7		T		  0		sft%n
*/

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

		n := len(clock)
		var s, e int // s: clock start; e: clock end

		if shift%(2*n) < n { // inverted triangle
			s, e = shift%n, n
		} else { // triangle
			s, e = 0, shift%n
		}

		for line := range clock[0] {

			for j := 0; j < n-e; j++ {
				fmt.Print("   ")
			}
			for i := s; i < e; i++ {
				fmt.Printf("%v ", clock[i][line])
			}
			fmt.Println()
		}
		time.Sleep(time.Second * 2)
	}
}
