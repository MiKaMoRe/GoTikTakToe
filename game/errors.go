package game

import "fmt"

type InvalidMove Game

func (e InvalidMove) Error() string {
	return fmt.Sprint("Недопустимый ход")
}
