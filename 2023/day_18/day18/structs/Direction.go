package structs

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func (d Direction) String() string {
	return [...]string{"UP", "RIGHT", "DOWN", "LEFT"}[d]
}

func Direction_from_char(_direction string) Direction {
	switch _direction {
	case "U":
		return UP
	case "R":
		return RIGHT
	case "D":
		return DOWN
	case "L":
		return LEFT
	}
	return UP
}
