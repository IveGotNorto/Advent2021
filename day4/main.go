package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type BoardPiece struct {
	piece  string
	called bool
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	game := strings.Split(string(data), "\n\n")
	numbers := strings.Split(game[0], ",")
	game = game[1:]

	var boards = make(map[int]([5][5]BoardPiece))
	//fmt.Printf("numbers: %v\n", numbers)
	for i, board := range game {
		var tmp [5][5]BoardPiece
		rows := strings.Split(board, "\n")
		for j, row := range rows {
			pieces := strings.Split(row, " ")
			k := 0
			for _, piece := range pieces {
				if piece != "" {
					tmp[j][k] = BoardPiece{piece, false}
					k += 1
				}
			}
		}
		boards[i] = tmp
	}

	score := 0
	for _, play := range numbers {
		for k := 0; k < len(boards); k++ {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if boards[k][i][j].piece == play {
						if entry, ok := boards[k]; ok {
							inBoard := entry[i][j]
							inBoard.called = true
							entry[i][j] = inBoard
							boards[k] = entry
						}
					}
				}
			}
			if checkWin(boards[k]) {
				score = getScore(boards[k])
				buff, err := strconv.Atoi(play)
				score = score * buff
				if err != nil {
					fmt.Println(err)
				}
				goto WINNER
			}
		}
	}
WINNER:
	fmt.Printf("Winning score is %v\n", score)

}

func checkWin(board [5][5]BoardPiece) bool {
	winner := false
	count := 1

	for i := 0; i < 5; i++ {
		count = 1
		if board[i][0].called {
			for j := 1; j < 5; j++ {
				if !board[i][j].called {
					break
				}
				count++
			}
		}
		if count == 5 {
			winner = true
			break
		}
	}

	if !winner {
		for i := 0; i < 5; i++ {
			count = 1
			if board[0][i].called {
				for j := 1; j < 5; j++ {
					if !board[j][i].called {
						break
					}
					count++
				}
			}
			if count == 5 {
				winner = true
				break
			}
		}
	}

	if !winner {
		for i := 0; i < 5; i++ {
			count = 1
			if board[i][i].called {
				count++
			} else {
				break
			}
		}
		if count == 5 {
			winner = true
		}
	}

	if !winner {
		for i := 0; i < 5; i++ {
			count = 1
			if board[i][4-i].called {
				count++
			} else {
				break
			}
		}
		if count == 5 {
			winner = true
		}
	}

	return winner
}

func getScore(board [5][5]BoardPiece) int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board[i][j].called {
				buff, err := strconv.Atoi(board[i][j].piece)
				if err != nil {
					fmt.Println(err)
				}
				score += buff
			}
		}
	}
	return score
}
