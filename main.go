package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

const (
	Termwidth  = 80
	Termheight = 24
)

func debug(rest ...interface{}) {
	fmt.Fprintln(os.Stderr, rest...)
}

func main() {
	fmt.Printf("Hello, world.\n")

	//initialize Termbox
	termbox.Init()
	defer termbox.Close()

	//Create new Game
	d := newDungeon(Termwidth, Termheight)
	//draw gameboard
	d.DrawDungeon()
	//display player messages

	//Game lpop
	gameover := false
	//	fmt.Fprintf(os.Stderr, "%d %d\n", (d.player.x - d.level.rooms[0].x), ((d.player.y + 1) - (d.level.rooms[0].y)))
	//	fmt.Fprintf(os.Stderr, "%d %d\n", d.level.rooms[0].x, d.level.rooms[0].y)
	//	fmt.Fprintf(os.Stderr, "%d %d\n", d.level.rooms[0].width, d.level.rooms[0].height)

	for !gameover {

		//update board
		//draw gameboard
		//display player messages
		d.DrawDungeon()
		event := termbox.PollEvent()
		switch event.Ch {
		case 'w':
			//fmt.Fprintf(os.Stderr, "%d %d", d.player.x, d.player.y+1)
			//	fmt.Fprintf(os.Stderr, "%d %d", d.level.rooms[0].x, d.level.rooms[0].y)

			/*	if d.level.rooms[d.player.roomID].AT((d.player.x - d.level.rooms[0].x), ((d.player.y + 1) - (d.level.rooms[0].y))).Walkable() {
				debug(d.player.x-1, d.player.y)

			}*/
			d.player.move(d.player.x, d.player.y-1)
		case 'a':
			//	debug(d.player.x-1, d.player.y)
			//	fmt.Fprintf(os.Stderr, "%d %d", d.player.x-1, d.player.y)
			//	fmt.Fprintf(os.Stderr, "%d %d", d.level.rooms[0].x, d.level.rooms[0].y)
			/*if d.level.rooms[d.player.roomID].AT(d.player.x-1, d.player.y).Walkable() {
				debug(d.player.x-1, d.player.y)

			}*/
			d.player.move(d.player.x-1, d.player.y)
		case 's':
			//		fmt.Fprintf(os.Stderr, "%d %d", d.player.x, d.player.y-1)
			//	fmt.Fprintf(os.Stderr, "%d %d", d.level.rooms[0].x, d.level.rooms[0].y)

			/*if d.level.rooms[d.player.roomID].AT(d.player.x, d.player.y-1).Walkable() {
				debug(d.player.x-1, d.player.y)

			}*/
			d.player.move(d.player.x, d.player.y+1)
		case 'd':
			//	fmt.Fprintf(os.Stderr, "%d %d", d.player.x+1, d.player.y)
			//	fmt.Fprintf(os.Stderr, "%d %d", d.level.rooms[0].x, d.level.rooms[0].y)

			/*if d.level.rooms[d.player.roomID].AT(d.player.x+1, d.player.y).Walkable() {
				debug(d.player.x-1, d.player.y)

			}*/
			d.player.move(d.player.x+1, d.player.y)
		case 'q':
			gameover = !gameover
		}
	}

}
