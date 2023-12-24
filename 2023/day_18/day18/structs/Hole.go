package structs

import "fmt"

type Hole struct {
	Start  *Position
	End    *Position
	Length int
	Color  string
}

func (h *Hole) String() string {
	return h.Color + " " + h.Start.String() + " " + h.End.String() + " " + fmt.Sprintf("%d", h.Length)
}
