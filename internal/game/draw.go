package game

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	green        = "\x1b[32m"
	yellow       = "\x1b[33m"
	brightYellow = "\x1b[93m"
	cyan         = "\x1b[36m"
	red          = "\x1b[31m"
	reset        = "\x1b[0m"
	bold         = "\x1b[1m"
	boldOff      = "\x1b[22m"
)

func (g *Game) draw() {
	fmt.Print("\033[H\033[2J")

	for i := 0; i < g.width+2; i++ {
		fmt.Printf("%s#", green)
	}
	fmt.Printf("\r\n")

	for y := 0; y < g.height; y++ {
		fmt.Print("#")

		if g.running {
			fmt.Print(yellow)
		} else {
			fmt.Print(red)
		}

		for x := 0; x < g.width; x++ {

			os.Stdout.Write([]byte{byte(g.getPixel(Point{X: x, Y: y}))})
		}

		fmt.Printf("%s#\r\n", green)
	}

	for i := 0; i < g.width+2; i++ {
		fmt.Print("#")
	}

	fmt.Printf("\r\n%sCount: %s%d\r\n", cyan, brightYellow, g.counter)

	fmt.Printf("\r\n")

}

func (g *Game) getPixel(pixel Point) rune {

	if pixel == g.food {
		return '$'
	}

	for _, value := range g.snake.body {
		if pixel == value {
			if pixel == g.snake.body[0] {
				if g.running {
					return 'O'
				}
				return 'X'
			}
			return 'o'
		}
	}

	return ' '
}

func info1() {
	fmt.Printf("%s\nWelcome to Snake Game. This is a simple version of snake in %s%sGo%s\n", green, bold, cyan, boldOff)
	fmt.Printf("%s\nATTENTION: Only W, A, S, D are used to control the snake\n", green)
	fmt.Printf("Press ` to exit\n")

	fmt.Print("Please to pick certain game mode: Classic or Hardcore (c/h):")
}

func info2() {
	fmt.Printf("\n%s%sStarting", bold, red)

	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}

	fmt.Print(strings.Repeat("\n", 2))

	fmt.Printf(`%s   ██████╗  ██████╗ 
  ██╔════╝ ██╔═══██╗
  ██║  ███╗██║   ██║
  ██║   ██║██║   ██║
  ╚██████╔╝╚██████╔╝
   ╚═════╝  ╚═════╝`, cyan)

	fmt.Printf("%s\n", reset)
	time.Sleep(time.Second)
}
