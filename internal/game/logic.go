package game

import (
	"math/rand"
)

func (g *Game) update() {
	if len(g.inputBuffer) > 0 {
		g.snake.dir = g.inputBuffer[0]
		g.inputBuffer = g.inputBuffer[1:]
	}

	willEat := g.checkFood()

	g.snake.move(willEat)

	if willEat {
		g.getFood()
		g.getCounter()
	}

	if g.checkColussion() {
		g.running = false
	}

}

func (s *Snake) move(grow bool) {

	newHead := s.body[0]

	switch s.dir {
	case Up:
		newHead.Y--
	case Down:
		newHead.Y++
	case Right:
		newHead.X++
	case Left:
		newHead.X--
	}

	if grow {
		s.body = append([]Point{newHead}, s.body...)
	} else {
		s.body = append([]Point{newHead}, s.body[:len(s.body)-1]...)
	}

}

func (g *Game) checkFood() (grow bool) {

	if g.snake.body[0] == g.food {
		grow = true
	} else {
		grow = false
	}

	return
}

func (g *Game) getFood() {

	for {
		newFood := Point{
			X: rand.Intn(g.width),
			Y: rand.Intn(g.height),
		}

		flag := true
		for _, value := range g.snake.body {
			if newFood == value {
				flag = false
				break
			}
		}

		if flag {
			g.food = newFood
			break
		}
	}
}

func (g *Game) checkColussion() bool {

	head := g.snake.body[0]

	if head.X < 0 || head.X > g.width-1 || head.Y < 0 || head.Y > g.height-1 {
		return true
	}

	for _, value := range g.snake.body[1:] {
		if head == value {
			return true
		}
	}

	return false
}

func (g *Game) getCounter() {
	g.counter++
}
