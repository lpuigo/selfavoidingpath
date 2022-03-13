package zone

import (
	"errors"
	"fmt"
	"strings"
)

type Zone struct {
	Vals []int
	Size int
}

func New(size int) Zone {
	z := Zone{
		Vals: make([]int, size*size),
		Size: size,
	}
	return z
}

func (z Zone) String() string {
	builder := strings.Builder{}
	index := 0
	for row := 0; row < z.Size; row++ {
		for col := 0; col < z.Size; col++ {
			builder.WriteString(fmt.Sprintf("%2d ", z.Vals[index]))
			index++
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (z *Zone) SetPos(p Pos, val int) {
	z.Vals[p.Y*z.Size+p.X] = val
}

func (z Zone) GetPos(p Pos) int {
	return z.Vals[p.Y*z.Size+p.X]
}

func (z Zone) Neighbor(p Pos, dir Direction) (Pos, error) {
	switch dir {
	case DirRight:
		if p.X == z.Size-1 {
			return p, errors.New(fmt.Sprintf("can not go right from %s", p.String()))
		}
		return Pos{
			X: p.X + 1,
			Y: p.Y,
		}, nil
	case DirLeft:
		if p.X == 0 {
			return p, errors.New(fmt.Sprintf("can not go left from %s", p.String()))
		}
		return Pos{
			X: p.X - 1,
			Y: p.Y,
		}, nil
	case DirDown:
		if p.Y == z.Size-1 {
			return p, errors.New(fmt.Sprintf("can not go down from %s", p.String()))
		}
		return Pos{
			X: p.X,
			Y: p.Y + 1,
		}, nil
	case DirUp:
		if p.Y == 0 {
			return p, errors.New(fmt.Sprintf("can not go up from %s", p.String()))
		}
		return Pos{
			X: p.X,
			Y: p.Y - 1,
		}, nil
	}
	return Pos{}, errors.New(fmt.Sprintf("unkwnown direction %d", int(dir)))
}

func (z Zone) PossibleDirections(p Pos) []Direction {
	res := []Direction{}
	if p.X < z.Size-1 && z.GetPos(Pos{p.X + 1, p.Y}) == 0 {
		res = append(res, DirRight)
	}
	if p.Y > 0 && z.GetPos(Pos{p.X, p.Y - 1}) == 0 {
		res = append(res, DirUp)
	}
	if p.X > 0 && z.GetPos(Pos{p.X - 1, p.Y}) == 0 {
		res = append(res, DirLeft)
	}
	if p.Y < z.Size-1 && z.GetPos(Pos{p.X, p.Y + 1}) == 0 {
		res = append(res, DirDown)
	}
	return res
}
