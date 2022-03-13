package main

import (
	"fmt"
	"github.com/lpuig/selfavoidingpath/zone"
	"math/rand"
	"time"
)

const (
	side int = 8
)

var Attempt int
var MaxDepth int

func main() {
	z := zone.New(side)

	//currentPos := zone.Pos{
	//	rand.Intn(len(z)),
	//	rand.Intn(len(z)),
	//}

	currentPos := zone.Pos{}
	t := time.Now()
	rand.Seed(t.Unix())

	SelfAvoidingPath(z, currentPos, 1)

	fmt.Println(z.String())
	fmt.Printf("result in %d attempts, took %s\n", Attempt, time.Since(t).String())

}

func SelfAvoidingPath(z zone.Zone, p zone.Pos, val int) bool {
	Attempt++
	if val > MaxDepth {
		MaxDepth = val
	}
	if Attempt%100000000 == 0 {
		fmt.Printf("Attempt %d, depth %d\n", Attempt, MaxDepth)
	}
	z.SetPos(p, val)
	if val == z.Size*z.Size {
		return true
	}
	possibleDirs := z.PossibleDirections(p)
	if len(possibleDirs) == 0 {
		//fmt.Printf("fail at %d\n", val)
		//fmt.Println(z.String())
		return false
	}
	// one or more possible dirs ... scramble them
	//dirOrder := possibleDirs
	dirOrder := make([]zone.Direction, len(possibleDirs))
	for i := 0; i < len(dirOrder); i++ {
		choice := rand.Intn(len(possibleDirs))
		dirOrder[i] = possibleDirs[choice]
		possibleDirs = append(possibleDirs[:choice], possibleDirs[choice+1:]...)
	}
	// and test each of one
	for _, direction := range dirOrder {
		nextPos, _ := z.Neighbor(p, direction)
		if SelfAvoidingPath(z, nextPos, val+1) {
			return true
		}
		// reverse this test dir and proceed
		z.SetPos(nextPos, 0)
	}
	return false
}
