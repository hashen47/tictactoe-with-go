package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Player int

const (
    playerX Player = 0
    playerO Player = 1
)

type State int

const (
    Normal State = 0
    Win State = 1
    Draw State = 2
)

type Game struct {
    isPlay bool
    positionInvalidMsg string
    board [9]string
    player Player
    prevPositions []int
    gameState State
}

func (g *Game) showBoard() {
    fmt.Printf("             %v|%v|%v\n", g.board[0], g.board[1], g.board[2])
    fmt.Printf("             %v|%v|%v\n", g.board[3], g.board[4], g.board[5])
    fmt.Printf("             %v|%v|%v\n", g.board[6], g.board[7], g.board[8])
}

func (g *Game) showTurn() {
    fmt.Print(`
          Turn : ` + g.getSymbol() + ` 
    `)
}

func (g *Game) gameOverMsg() {
    state := "DRAW"
    winner := "----"

    if g.gameState == Win {
        state = "WIN"
        winner = g.getSymbol()
    }

    fmt.Print(`
        ############### 
        |  Game Over  |
        ############### 

          State : ` + state + `
         Winner : ` + winner + ` 
    `)
}

func (g *Game) clearTerminal() {
    cmdText := "clear"

    if runtime.GOOS == "windows" {
        cmdText = "cls"
    }

    cmd := exec.Command(cmdText)
    cmd.Stdout = os.Stdout
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

func (g *Game) showGameStateMsg() {
    g.clearTerminal()
    switch {
    case g.gameState == Win || g.gameState == Draw:
        g.gameOverMsg()
    default:
        fmt.Println(`
        ############### 
        | Tic Tac Toe |
        ############### 
        `)
        g.showBoard()
        g.showTurn()
    }
}

func (g *Game) playAgainOrExit() {
    scan := bufio.NewScanner(os.Stdin)
    fmt.Print("\nDo you want to play again? ")
    scan.Scan()

    input := []rune(scan.Text())

    g.isPlay = true
    if len(input) == 0 {
        g.isPlay = false
    } else if strings.ToUpper(string(input[0])) != "Y" {
        g.isPlay = false
    }
}

func (g *Game) switchPlayer() {
    if g.positionInvalidMsg != "" {
        return
    }

    if g.player == playerX {
        g.player = playerO
    } else {
        g.player = playerX
    }
}

func (g *Game) getPositionFromUser() int {
    scan := bufio.NewScanner(os.Stdin)
    fmt.Print("       Pos : ")
    scan.Scan()
    input := []rune(scan.Text())

    if string(input) == "" {
        g.positionInvalidMsg = "position is required"
        return 0
    }

    pos, _ := strconv.Atoi(string(input))
    g.positionInvalidMsg = ""

    return pos
}

func (g *Game) getSymbol() string {
    if g.player == playerX {
        return "X"
    }
    return "O"
}

func (g *Game) setGameState() {
    symbol := g.getSymbol()
    winPositions := [][]int{
        {0, 1, 2},
        {3, 4, 5},
        {6, 7, 8},
        {0, 3, 6},
        {1, 4, 7},
        {2, 5, 8},
        {0, 4, 8},
        {2, 4, 6},
    }

    for _, winPos := range winPositions {
        if symbol == g.board[winPos[0]] && g.board[winPos[0]] == g.board[winPos[1]] && g.board[winPos[1]] == g.board[winPos[2]] {
            g.gameState = Win
            return
        }
    }

    if len(g.prevPositions) == 9 {
        g.gameState = Draw
        return
    }
}

func (g *Game) setPosition() {
    pos := g.getPositionFromUser()

    if g.positionInvalidMsg != "" {
        return
    }

    if pos < 1 || pos > 9 {
        g.positionInvalidMsg = "Invalid input"
        return
    }

    isAlreadyIncludes := false
    for _, p := range g.prevPositions {
        if p == pos - 1 {
            isAlreadyIncludes = true
            break
        }
    }

    if isAlreadyIncludes {
        g.positionInvalidMsg = "This position is already choose, try another position"
        return
    }

    g.prevPositions = append(g.prevPositions, pos - 1)
    g.board[pos - 1] = g.getSymbol()
    g.positionInvalidMsg = ""
}

func (g *Game) reset() {
    g.board = [9]string{"_", "_", "_", "_", "_", "_", "_", "_", "_"}
    g.prevPositions = make([]int, 0, 9)
    g.player = playerX
    g.gameState = Normal
    g.positionInvalidMsg = ""
}

func Run() {
    g := Game{}
    for {
        g.reset()
        for {
            g.showGameStateMsg()
            g.setPosition()
            g.setGameState()
            if g.gameState != Normal {
                g.showGameStateMsg()
                g.playAgainOrExit()
                break
            }
            g.switchPlayer()
        }

        if !g.isPlay {
            fmt.Println("Byeee..")
            break
        }
    }
}
