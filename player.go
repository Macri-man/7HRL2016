package main

type Player struct {
	hp     int
	maxHP  int
	weapon weapon
	x, y   int
	roomID int
}

func newPlayer(x, y int) Player {
	return Player{
		hp:     15,
		maxHP:  15,
		weapon: newWeapon("UNARMED", 5, '?'),
		x:      x,
		y:      y,
		roomID: 1,
	}
}

func (p Player) getName(name string) string {
	return name
}

func (p Player) Draw() rune {
	return '@'
}

func (p Player) Position() (int, int) {
	return p.x, p.y
}

func (p *Player) Attack() int {
	return p.weapon.Attack()
}

func (p *Player) Defend(amount int) {
	p.hp -= amount
}

func (p Player) getHP() int {
	return p.hp
}

func (p *Player) move(x, y int) {
	p.x = x
	p.y = y
}

/*
func (p *player) useItem(item int, name string) string {
	switch item {
	case cell:
		p.hp = p.maxHP
		return "Cosume Potion"
	case cell:
		p.hp += 5
		return "Cosume food"
	case cell:
		p.weapon = newClassWeapon(name)
		return "Equip Weapon"
	}
}*/

/*
type playerActions interface {
	getName() string
	getAttackType() string
	Attack() int
	Defend() int
	UseItem() bool
	Draw() rune
	getHP() int
*/
