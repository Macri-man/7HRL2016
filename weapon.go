package main

type weapon struct {
	name   string
	damage int
	symbol rune
}

func (w weapon) getName() string {
	return w.name
}

func (w weapon) Attack() int {
	return w.damage
}

func (w weapon) Draw() rune {
	return w.symbol
}

func newWeapon(newName string, newDamage int, newSymbol rune) weapon {
	return weapon{name: newName, damage: newDamage, symbol: newSymbol}
}

func newClassWeapon(name string) weapon {
	switch name {
	case "Sword":
		return newWeapon(name, 5, 'S')
	case "Axe":
		return newWeapon(name, 5, 'S')
	case "Hell":
		return newWeapon(name, 1, 'S')
	case "YOLO":
		return newWeapon(name, 0, 'S')
	case "Wand":
		return newWeapon(name, 5, 'S')
	}
	return newWeapon(name, 8, 'l')
}
