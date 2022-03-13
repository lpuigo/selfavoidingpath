package zone

import "fmt"

type Direction int

const (
	DirRight Direction = iota
	DirUp
	DirLeft
	DirDown
)

type Pos struct {
	X int
	Y int
}

func (p Pos) String() string {
	return fmt.Sprintf("[%2d, %2d]", p.X, p.Y)
}
