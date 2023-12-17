package structs

type direction int

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

func (dir direction) get_opposite() direction {
	return direction((int(dir) + 2) % 4)
}
