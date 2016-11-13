package main

type itemType int

const (
	potion itemType = iota
	food
	dirt
	empty
)

func (char itemType) DrawCell() rune {
	switch char {
	case potion:
		return 'P'
	case food:
		return 'F'
	case dirt:
		return 'D'
	case empty:
		return '.'
	}
	return '!'
}

type Items struct {
	name string
	item itemType

	effect func() (string, int)
}

func (i Items) DrawItem() rune {
	return i.item.DrawCell()
}

func (i Items) getName() string {
	return i.name
}

func newItem(item itemType) Items {
	switch item {
	case potion:
		return Items{
			name:   "Potion",
			item:   potion,
			effect: func() (string, int) { return "Consume Potion", 6 },
		}
	case food:
		return Items{
			name:   "Food",
			item:   food,
			effect: func() (string, int) { return "Consume Food", 3 },
		}
	case dirt:
		return Items{
			name:   "Dirt",
			item:   dirt,
			effect: func() (string, int) { return "Eat Dirt", 0 },
		}
	}
	return Items{
		name:   "Empty",
		item:   empty,
		effect: func() (string, int) { return "Nothing Here", -1 },
	}
}
