package main

import (
	"fmt"
  "os"
  "os/exec"
  "strings"
  "strconv"
)

const BoardSize = 15

type Gomoku struct {
	Board []int
}

func main() {
  gomock_control()
}

func gomock_control() {
  stone := map[bool]string{
    true : "black",
    false : "white",
  }
  game_status := draw_board(nil)
  turn := true
  game_end := false
  for game_end == false {
    fmt.Printf("turn %s (Input x,y) : ", stone[turn])
    var inputstr string
    _, err := fmt.Scanf("%s", &inputstr)
    if err != nil {
        fmt.Println(err)
    }
    slice := strings.Split(inputstr, ",")
    x, _ := strconv.Atoi(slice[0])
    y, _ := strconv.Atoi(slice[1])
    pos := x + (y * BoardSize)
    if len(game_status) < pos || game_status[pos] == 9 || game_status[pos] == 10 {
      continue
    }
    if turn {
      game_status[pos] = 9
      game_end = check_end(game_status, 9)
    } else {
      game_status[pos] = 10
      game_end = check_end(game_status, 10)
    }
    draw_board(game_status)
    if game_end {
      fmt.Printf("WIN %s ! ", stone[turn])
    }
    turn = !turn
  }
}

func check_end(status []int, turn int) bool {
  for row := 0; row < BoardSize; row++ {
    for col := 0; col < BoardSize - 5; col++ {
      win := true
      for i := 0; i < 5; i++ {
        if status[row + col * BoardSize + i] != turn {
          win = false
        }
      }
      if win { return true }
    }
  }
  for col := 0; col < BoardSize - 5; col++ {
    for row := 0; row < BoardSize; row++ {
      win := true
      for i := 0; i < 5; i++ {
        if status[row + ((col + i) * BoardSize)] != turn {
          win = false
        }
      }
      if win { return true }
    }
  }
  for col := 0; col < BoardSize - 5; col++ {
    for row := 0; row < BoardSize - 5; row++ {
      win := true
      for i := 0; i < 5; i++ {
        if status[(row + i) + ((col + i) * BoardSize)] != turn {
          win = false
        }
      }
      if win { return true }
    }
  }
  for col := 0; col < BoardSize - 5; col++ {
    for row := BoardSize - 1; row >= 4; row-- {
      win := true
      for i := 0; i < 5; i++ {
        if status[(row - i) + ((col + i) * BoardSize)] != turn {
          win = false
        }
      }
      if win { return true }
    }
  }

  return false
}

func draw_board(status []int) []int {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
  gomoku := new(Gomoku)
  if status != nil {
    gomoku.Board = status
  } else {
    gomoku.Board = make([]int, BoardSize * BoardSize)
    for i := 0; i < BoardSize * BoardSize; i++ {
      if i == 0 {
		    gomoku.Board[i] = 1
      } else if i == BoardSize - 1 {
		    gomoku.Board[i] = 3
      } else if i < BoardSize {
		    gomoku.Board[i] = 2
      } else if i == BoardSize * BoardSize - 1 {
		    gomoku.Board[i] = 8
      } else if i == BoardSize * (BoardSize - 1) {
		    gomoku.Board[i] = 6
      } else if i % BoardSize == 0 {
		    gomoku.Board[i] = 4
      } else if i % BoardSize == BoardSize - 1 {
		    gomoku.Board[i] = 5
      } else if i > BoardSize * (BoardSize - 1) {
		    gomoku.Board[i] = 7
      } else {
		    gomoku.Board[i] = 0
      }
	  }
  }
  board := map[int]string{
    0: "\u254B", // ╋
    1: "\u250F", // ┏
    2: "\u2533", // ┳
    3: "\u2513", // ┓
    4: "\u2523", // ┣
    5: "\u252B", // ┫
    6: "\u2517", // ┗	
    7: "\u253B", // ┻
    8: "\u251B", // ┛
    9: "●",
    10: "○",
  }
  for i := range gomoku.Board {
    if i > 0 && i % BoardSize == 0 {
      fmt.Print("\n")
    }
    fmt.Print(board[gomoku.Board[i]])
  }
  fmt.Print("\n")
  return gomoku.Board
}
