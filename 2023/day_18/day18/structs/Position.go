package structs

import "fmt"

type Position struct {
	X int
	Y int
}

func (p *Position) String() string {
	return "(" + fmt.Sprintf("%d", p.X) + ", " + fmt.Sprintf("%d", p.Y) + ")"
}
