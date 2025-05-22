package game

import (
	"fmt"
)

type Game struct {
	board      [][]byte
	size       byte
	player     byte
	winner     byte
	playersMap map[byte]string
}

func (g *Game) Init(size byte) {
	g.size = size
	g.board = make([][]byte, g.size)
	for i := range g.board {
		g.board[i] = make([]byte, g.size)
	}
	g.player = 1
	g.winner = 0

	g.playersMap = map[byte]string{
		0: " ",
		1: "X",
		2: "0",
	}

	var err error

	for g.winner == 0 {
		CallClear()
		if err != nil {
			fmt.Println(err)
			err = nil
		}
		fmt.Println()
		g.PrintBoard()
		fmt.Println("Ход ", g.playersMap[g.player])

		var row byte
		var col byte

		fmt.Print("Строка: ")
		fmt.Scan(&row)
		fmt.Print("Колонка: ")
		fmt.Scan(&col)
		fmt.Println()

		err = g.MakeMove(row-1, col-1)
	}
	CallClear()
	g.PrintBoard()
	fmt.Println("┌───────────┐")
	fmt.Println("│ПОБЕДИЛ:", g.playersMap[g.winner], "│")
	fmt.Println("└───────────┘")
}

func (g *Game) ValidateMove(x byte, y byte) error {
	if x >= g.size || y >= g.size {
		return InvalidMove{}
	}
	if g.board[x][y] != 0 {
		return InvalidMove{}
	}
	return nil
}

func (g *Game) PrintBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(g.playersMap[g.board[i][j]])
			if j < 2 {
				fmt.Print("│")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("─┼─┼─")
		}
	}
}

func (g *Game) MakeMove(x byte, y byte) error {
	err := g.ValidateMove(x, y)

	if err != nil {
		return err
	}

	g.board[x][y] = g.player

	g.CheckWinner()

	g.player = g.player%2 + 1
	return nil
}

func (g *Game) CheckWinner() {
	var j byte

	// Check row win
	for _, row := range g.board {
		rowWin := true

		for j = 0; j < g.size; j++ {
			if row[j] != g.player {
				rowWin = false
				break
			}
		}

		if rowWin {
			g.winner = g.player
			return
		}
	}

	// Check column win
	for i := range g.board {
		colWin := true

		for j = 0; j < g.size; j++ {
			if g.board[j][i] != g.player {
				colWin = false
				break
			}
		}

		if colWin {
			g.winner = g.player
			return
		}
	}

	// Check main diagonal win
	diagWin := true

	for j = 0; j < g.size; j++ {
		if g.board[j][j] != g.player {
			diagWin = false
		}
	}

	if diagWin {
		g.winner = g.player
		return
	}

	// Check secondary diagonal win
	diagWin = true

	for j = 0; j < g.size; j++ {
		if g.board[g.size-j-1][j] != g.player {
			diagWin = false
		}
	}

	if diagWin {
		g.winner = g.player
		return
	}
}
