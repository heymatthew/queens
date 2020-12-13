package main

import "fmt"

type Queen struct {
	x, y int
}

func (q Queen) attackY() int {
	return q.x - q.y
}

func (q Queen) attackX() int {
	return q.x + q.y
}

type Board []Queen

func (b Board) PrintSquare(x, y int) rune {
	for _, q := range b {
		if q.x == x && q.y == y {
			return '👑'
		}
	}
	if (x+y)%2 == 0 {
		return '⬜'
	}
	return '⬛'
}

func (b Board) String() string {
	size := len(b)
	var squares [size][size]bool
	for _, q := range b {
		squares[q.x][q.y] = true
	}

	var buff []rune
	for x := range b {
		for y := range b {
			buff = append(buff, b.PrintSquare(x, y))
		}
		buff = append(buff, '\n')
	}
	return string(buff)
}

func (b Board) Valid() bool {
	if b == nil {
		return true
	}

	for pos, queen := range b {
		for _, opponent := range b[pos+1:] {
			if queen.x == opponent.x {
				return false
			}
			if queen.y == opponent.y {
				return false
			}
			if queen.attackY() == opponent.attackY() {
				return false
			}
			if queen.attackX() == opponent.attackX() {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(Board{
		Queen{0, 0},
		Queen{1, 1},
		Queen{2, 2},
		Queen{3, 2},
		Queen{4, 2},
		Queen{5, 2},
		Queen{5, 2},
		Queen{5, 2},
		Queen{5, 2},
		Queen{5, 2},
		Queen{5, 2},
	})
}
