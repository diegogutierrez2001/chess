package chess

import (
	"fmt"
	"errors"
	"strconv"
)

type Square struct {
	Rank, File Int
}

func (s Square) InBounds() bool {
	return (1 <= s.Rank && s.Rank <= 8) && (1 <= s.File && s.File <= 8)
}

func (s Square) ToString() (string, error) {

	// check that square is in bounds
	if (s.InBounds() == false)
		return "", errors.New("Out of bounds.")

	// get appropriate letter for square file
	switch (s.File) {
	case 1:	return "A" + strconv.Itoa(s.Rank), nil
	case 2: return "B" + strconv.Itoa(s.Rank), nil
	case 3:	return "C" + strconv.Itoa(s.Rank), nil
	case 4: return "D" + strconv.Itoa(s.Rank), nil
	case 5:	return "E" + strconv.Itoa(s.Rank), nil
	case 6: return "F" + strconv.Itoa(s.Rank), nil
	case 7:	return "G" + strconv.Itoa(s.Rank), nil
	case 8: return "H" + strconv.Itoa(s.Rank), nil
	}
}

func Test() {
	Square testSquare{1, 1}
	fmt.Println(testSquare.ToString())
}

