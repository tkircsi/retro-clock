package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

type placeholder [5]string

func main() {

	var digits [10]placeholder = [10]placeholder{
		{
			"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		},
		{
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"  █",
			"███",
		},
		{
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		},
		{
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		},
		{
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		},
		{
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		},
		{
			"███",
			"█ █",
			"███",
			"  █",
			"███",
		},
	}

	colon := [2]placeholder{
		{
			"   ",
			" ▒ ",
			"   ",
			" ▒ ",
			"   ",
		},
		{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
	}

	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon[sec%2],
			digits[min/10], digits[min%10],
			colon[sec%2],
			digits[sec/10], digits[sec%10],
		}

		// fmt.Printf("%d:%d:%d\n", hour, min, sec)

		for i := range digits[0] {
			for _, d := range clock {
				fmt.Printf("%s ", d[i])
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}

}
