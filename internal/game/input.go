package game

import (
	"log"
	"os"

	"golang.org/x/term"
)

func (g *Game) handleInput() {
	defer close(g.end)

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Printf("ERROR: trying to change terminal mode\t%v", err)
		return
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var data [1]byte

	for g.running {
		_, err := os.Stdin.Read(data[:])

		if err != nil {
			log.Printf("ERROR: trying to read indormation from console\t%v", err)
			break
		}

		lastDir := g.snake.dir

		if len(g.inputBuffer) > 0 {
			lastDir = g.inputBuffer[len(g.inputBuffer)-1]
		}

		if len(g.inputBuffer) >= 2 {
			continue
		}

		switch data[0] {
		case 'w', 'W':
			if lastDir != Down && lastDir != Up {
				g.inputBuffer = append(g.inputBuffer, Up)
			}
		case 'a', 'A':
			if lastDir != Right && lastDir != Left {
				g.inputBuffer = append(g.inputBuffer, Left)
			}
		case 'd', 'D':
			if lastDir != Left && lastDir != Right {
				g.inputBuffer = append(g.inputBuffer, Right)
			}
		case 's', 'S':
			if lastDir != Up && lastDir != Down {
				g.inputBuffer = append(g.inputBuffer, Down)
			}
		case '`':
			g.running = false
			return
		}
	}

}
