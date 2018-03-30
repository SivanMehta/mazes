package generator

import (
  "log"
  "math/rand"
  "strconv"

  "github.com/ajstarks/svgo"
)

type Maze struct {
  Height int
  Width int
  Maze []int
}

func (this *Maze) IncrementCell(x int, y int) {
  this.Maze[ x * this.Width + y ] = 1
}

func (this *Maze) InBounds(x, y int) bool {
  return (0 <= x) && (x < this.Width) && (0 <= y) && (y < this.Height)
}

func (this *Maze) EmptyAt(x, y int) bool {
  cell := x * this.Width + y
  return this.Maze[cell] == 0
}

func (this *Maze) Print() {
  out := "\n"
  for row := 0; row < this.Width; row ++ {
    for col := 0; col < this.Height; col ++ {
      out += strconv.Itoa(this.Maze[row * this.Width + col])
    }
    out += "\n"
  }
  log.Print(out)
}

func drawCells(args *Args, canvas *svg.SVG) {
  width := args.Width
  height := args.Height
  cells := randomCells(args)
  for row := 0; row < height; row ++ {
    for col := 0; col < width; col ++ {
      color := "fill:"
      if cells[row * width + col] {
        color += "black"
      } else {
        color += "white"
      }
      canvas.Rect(
        col * args.CellSizeX,
        row * args.CellSizeY,
        args.CellSizeX,
        args.CellSizeY,
        color,
      )
    }
  }
}

func randomCells(args *Args) []bool {
  width := args.Width
  height := args.Height
  cells := make([]bool, width * height)

  for row := 0; row < height; row ++ {
    for col := 0; col < width; col ++ {
      cells[row * width + col] = rand.Intn(2) > 0
    }
  }

  return cells
}

var (
  left = [2]int{ -1, 0 }
  right = [2]int{ 1, 0 }
  up = [2]int{ 0, -1 }
  down = [2]int{ 0, 1 }
)

func generateCells(args *Args) []int {
  height := args.Height
  width := args.Width
  path := newStack(height * width)
  cells := make([]int, height * width)
  maze := Maze{ Maze: cells, Height: height, Width: width}
  directions := [4][2]int{ up, down, left, right }

  path.Push(Cell{ X: width / 2, Y: height / 2})

  log.Println(maze, directions)

  return cells
}
