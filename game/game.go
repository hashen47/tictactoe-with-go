package game

import (
	"fmt"
)

type Player int

const (
    playerX Player = 0
    playerO Player = 1
)

type Game struct {
    board []string
    player Player
}

func (g *Game) showBoard() {
    fmt.Printf("%v|%v|%v\n", g.board[0], g.board[1], g.board[2])
    fmt.Printf("%v|%v|%v\n", g.board[3], g.board[4], g.board[5])
    fmt.Printf("%v|%v|%v\n", g.board[6], g.board[7], g.board[8])
}

func Run() {
    g := Game{}
    g.board = []string{"_", "_", "_", "_", "_", "_", "_", "_", "_"}
    g.player = playerX
    g.showBoard()
}
