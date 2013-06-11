package main

import (
	"fmt"
)

type Spoon struct {
	Id int
}

func TryPickUp(s *Spoon) bool {
	return true
}

func PickUp(s *Spoon) bool {
	return true
}

func PutDown(s *Spoon) bool {
	return true
}

type Philosoph struct {
	Id int
	Left *Spoon
	Right *Spoon
	Consumed int
}

func InitialzePhilosoph(p *Philosoph, l *Spoon, r *Spoon)  {
}

func main() {
	var spoons [5] Spoon
	var philosophs [5] Philosoph
	
	for i := 0; i < 5; i++ {
		spoons[i] = Spoon { Id:i }
		philosophs[i] = Philosoph { Id: i }
		
		if (i == 5)
			IntializePhilosoph(&philosophs[i], &spoons[i], &spoons[0])
		else
			IntializePhilosoph(&philosophs[i], &spoons[i], &spoons[i + 1])
	}
	
	for i := 0; i < 5; i++ {
		fmt.Printf("\nSpoon: %+v", spoons[i].Id)
	}
	
	for i := 0; i < 5; i++ {
		//p := &philosophs[i]
		//fmt.Printf("\nLeft spoon: %+v  > Philosoph: %+v < Right spoon: %+v", p.Left.Id, p.Id, p.Right.Id)
	}
}