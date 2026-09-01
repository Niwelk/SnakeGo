package game

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Snake struct {
	body  []Point
	dir   Direction
	speed time.Duration
}

type Point struct {
	X int
	Y int
}

type Game struct {
	snake       Snake
	food        Point
	width       int
	height      int
	running     bool
	counter     int
	inputBuffer []Direction
	end         chan struct{}
}

func newGameClassic() Game {
	return Game{
		snake: Snake{
			body:  []Point{{X: 3, Y: 5}},
			dir:   Right,
			speed: 750 * time.Millisecond,
		},
		food:        Point{X: 3, Y: 2},
		width:       18,
		height:      7,
		running:     true,
		counter:     0,
		inputBuffer: []Direction{Right},
		end:         make(chan struct{}),
	}
}

func newGameHardcore() Game {
	return Game{
		snake: Snake{
			body:  []Point{{X: 3, Y: 5}},
			dir:   Right,
			speed: 180 * time.Millisecond,
		},
		food:        Point{X: 3, Y: 2},
		width:       18,
		height:      7,
		running:     true,
		counter:     0,
		inputBuffer: []Direction{Right},
		end:         make(chan struct{}),
	}
}

func selectorMode() Game {
	var signMode [1]byte

	_, err := os.Stdin.Read(signMode[:])

	if err != nil {
		log.Print(err)
		return newGameClassic()
	}

	switch signMode[0] {
	case 'c', 'C':
		return newGameClassic()
	case 'h', 'H':
		return newGameHardcore()
	default:
		fmt.Print("Something went wrong..")
		return newGameClassic()
	}
}

func Create() (game Game) {
	info1()
	game = selectorMode()
	info2()

	return
}

func (g *Game) Run() {
	fmt.Print("\033[?1049h\033[?25l")

	defer fmt.Print("\033[?25h\033[?1049l")

	go g.handleInput()

	for g.running {
		g.update()
		g.draw()
		time.Sleep(g.snake.speed)
	}

	<-g.end

	fmt.Print(bold, red)
	fmt.Print("Game Over")

	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
}
