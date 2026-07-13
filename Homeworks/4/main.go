package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	blackBackRank = []string{"♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜"}
	whiteBackRank = []string{"♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖"}
)

const (
	lightSquare = "\033[47;30m"
	darkSquare  = "\033[48;2;208;135;50;97m"
	resetColor  = "\033[0m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	size := readBoardSize(reader)
	whitePlayer := readName(reader, "Введите имя игрока белыми", "белый")
	blackPlayer := readName(reader, "Введите имя игрока чёрными", "чёрный")

	board := makeBoard(size)
	placePieces(board)
	printBoard(board, whitePlayer, blackPlayer)
}

func readBoardSize(reader *bufio.Reader) int {
	for {
		fmt.Print("Введите размерность поля (не меньше 2) [8]: ")
		line, _ := reader.ReadString('\n')
		value := strings.TrimSpace(line)
		if value == "" {
			return 8
		}

		size, err := strconv.Atoi(value)
		if err == nil && size >= 2 {
			return size
		}

		fmt.Println("Размерность должна быть целым числом не меньше 2.")
	}
}

func readName(reader *bufio.Reader, prompt, defaultName string) string {
	fmt.Printf("%s [%s]: ", prompt, defaultName)
	line, _ := reader.ReadString('\n')
	name := strings.TrimSpace(line)
	if name == "" {
		return defaultName
	}
	return name
}

func makeBoard(size int) [][]string {
	board := make([][]string, size)
	for row := range board {
		board[row] = make([]string, size)
	}
	return board
}

func placePieces(board [][]string) {
	size := len(board)

	placeBackRank(board[0], blackBackRank)
	placeBackRank(board[size-1], whiteBackRank)

	placePawns(board, 1, "♟")
	placePawns(board, size-2, "♙")
}

func placeBackRank(row, pieces []string) {
	for col := range row {
		row[col] = pieces[col%len(pieces)]
	}
}

func placePawns(board [][]string, row int, pawn string) {
	for col := range board[row] {
		if board[row][col] == "" {
			board[row][col] = pawn
		}
	}
}

func printBoard(board [][]string, whitePlayer, blackPlayer string) {
	size := len(board)
	rowNumberWidth := len(strconv.Itoa(size))

	fmt.Printf("\n%*s ", rowNumberWidth, "")
	for col := 0; col < size; col++ {
		fmt.Printf(" %-2s", columnName(col))
	}
	fmt.Println()

	for row := 0; row < size; row++ {
		fmt.Printf("%*d ", rowNumberWidth, row+1)
		for col := 0; col < size; col++ {
			piece := board[row][col]
			if piece == "" {
				piece = " "
			}

			color := lightSquare
			if (row+col)%2 != 0 {
				color = darkSquare
			}
			fmt.Printf("%s %-2s%s", color, piece, resetColor)
		}

		switch row {
		case 0:
			fmt.Printf("  %s (чёрные)", blackPlayer)
		case size - 1:
			fmt.Printf("  %s (белые)", whitePlayer)
		}
		fmt.Println()
	}
}

func columnName(column int) string {
	name := ""
	for column >= 0 {
		name = string(rune('A'+column%26)) + name
		column = column/26 - 1
	}
	return name
}
