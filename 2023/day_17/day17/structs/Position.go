package structs

type Position struct {
	X int
	Y int
}

func (p *Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}
