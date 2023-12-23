package structs

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func (d Direction) Opposite() Direction {
	return (d + 2) % 4
}

func (d Direction) Turn_left(current_direction Direction) Direction {
	return (current_direction + 3) % 4
}

func (d Direction) Turn_right(current_direction Direction) Direction {
	return (current_direction + 1) % 4
}

func (d Direction) To_String() string {
	switch d {
	case UP:
		return "UP"
	case RIGHT:
		return "RIGHT"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	}
	return ""
}
