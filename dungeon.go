package main

import termbox "github.com/nsf/termbox-go"

type Dungeon struct {
	level  GameMap
	player Player
}

func (d Dungeon) DrawDungeon() {
	for index := 0; index < d.level.numberOfRooms; index++ {
		d.level.rooms[index].DrawRoom()
	}

	termbox.SetCell(d.player.x, d.player.y, d.player.Draw(), termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func newDungeon(w, h int) Dungeon {
	return Dungeon{level: newGame(w, h, 1), player: newPlayer(w/2, h/2)}
}
