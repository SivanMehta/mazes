package generator

import (
  // "log"
  // "math/rand"
  "github.com/ajstarks/svgo"
)

type Maze struct {
  Maze []int
}

func (this *Maze) IncrementCell(x int, y int) {
  this.Maze[ x * mazeSize + y ] = 1
}

func (this *Maze) InBounds(x, y int) bool {
  return (0 <= x) && (x < mazeSize) && (0 <= y) && (y < mazeSize)
}

func (this *Maze) EmptyAt(x, y int) bool {
  cell := x * mazeSize + y
  return this.Maze[cell] == 0
}

func (this *Maze) Print() {
  out := "\n"
  for row := 0; row < mazeSize; row ++ {
    for col := 0; col < mazeSize; col ++ {
      out += strconv.Itoa(this.Maze[row * mazeSize + col])
    }
    out += "\n"
  }
  log.Print(out)
}

func drawCells(args *Args, canvas *svg.SVG) {
  // uses a maze algorithm to generate a solveable maze
}
