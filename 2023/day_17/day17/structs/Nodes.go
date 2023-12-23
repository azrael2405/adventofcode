package structs

import (
	"helper"
	"strconv"
)

type Node struct {
	Position   Position
	Base_value int
}

func (n *Node) Init(x, y int, str_value string) {
	n.Position = Position{x, y}
	value, ok := strconv.Atoi(str_value)
	helper.Check_error(ok)
	n.Base_value = value
}
