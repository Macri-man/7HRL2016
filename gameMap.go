package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type mapcell int

const (
	cellempty mapcell = iota
	cellwall
	celldoorh
	celldoorw
)

func (char mapcell) DrawCell() rune {
	switch char {
	case cellempty:
		return '.'
	case cellwall:
		return '#'
	case celldoorh:
		return '-'
	case celldoorw:
		return '|'
	}
	return '!'
}

func (char mapcell) Walkable() bool {
	walkable := false
	if char == cellempty || char == celldoorh || char == celldoorw {
		walkable = true
	}
	return walkable
}

type Door struct {
	x, y int
}

func newDoor(w, h int) *Door {
	if w == 0 {
		return &Door{w, rand.Intn(h)}
	} else {
		return &Door{rand.Intn(h), h}
	}
	return &Door{w, h}
}

type Room struct {
	width, height int
	x, y          int
	ID            int
	contents      []mapcell
	items         []Items
	doors         []Door
}

func newRoom(w, h, newx, newy, newID int) Room {
	room := Room{width: w, height: h, x: newx, y: newy, ID: newID, contents: make([]mapcell, h*w)}

	for index := 0; index < w*h; index++ {
		room.contents[index] = cellempty
	}

	for index := 0; index < w; index++ {
		room.SET(index, 0, cellwall)
		room.SET(index, h-1, cellwall)
	}

	for index := 0; index < h; index++ {
		room.SET(0, index, cellwall)
		room.SET(w-1, index, cellwall)
	}

	return room
}

func (r Room) AT(x, y int) mapcell {
	return r.contents[x+r.width*y]
}

func (r Room) ATPlayer(x, y int) mapcell {
	return r.contents[x+r.width*y]
}

func (r *Room) SET(x, y int, char mapcell) {
	r.contents[x+r.width*y] = char
}

func (r Room) DrawRoom() {
	for indexx := 0; indexx < r.width; indexx++ {
		for indexy := 0; indexy < r.height; indexy++ {
			termbox.SetCell(r.x+indexx, r.y+indexy, r.AT(indexx, indexy).DrawCell(), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

type GameMap struct {
	numberOfRooms int
	rooms         []Room
}

/*
func makeRoom(roomW,roomH,x,y,index int, Room *rooms) Room {
  if(index > 0){
    return newRoom(roomWidth, roomHeight, Termw/2, Termh/2, index)
  }
}
*/
/*
func makeHallway(roomW, roomH, x, y, index int, Room *rooms) {
	if index > 0 {
		&rooms = append(&rooms, newRoom(roomWidth, roomHeight, roomW/2, roomH/2, index))
	}
}

func makeRoom(roomW, roomH, x, y, index int, Room *rooms) {
	if index > 0 {
		&rooms = append(&rooms, newRoom(roomWidth, roomHeight, roomW/2, roomH/2, index))
	}
}
*/

func newGame(Termw, Termh, NOR int) GameMap {
	game := GameMap{numberOfRooms: NOR, rooms: []Room{}}
	//numRooms := NOR
	//doors := []Door{
	rand.Seed(time.Now().Unix())
	roomHeight := rand.Intn(15) + 9
	roomWidth := rand.Intn(15) + 9
	firstroom := newRoom(roomWidth, roomHeight, (Termw/2)-(roomWidth/2), (Termh/2)-(roomHeight/2), 0)
	game.rooms = append(game.rooms, firstroom)
	/*for index := 1;  < count; ++ {
	    roomHeight = rand.Intn(6) + 4
	  	roomWidth = rand.Intn(6) + 4
	    makeRoom(&game.rooms,roomHeight,roomWidth,roomHeight/2,roomWidth/2,index)
	  }*/

	//game.rooms = append(game.rooms, newRoom(roomWidth, roomHeight, Termw/2, Termh/2, index))
	//}
	return game
}
