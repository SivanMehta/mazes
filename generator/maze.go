package generator

import (
  "log"
  "math/rand"

  "github.com/ajstarks/svgo"
)

type Maze struct {
  Height int
  Width int
  Maze []bool
}

func (this *Maze) IncrementCell(x int, y int) {
  this.Maze[ x * this.Width + y ] = true
}

func (this *Maze) InBounds(x, y int) bool {
  return (0 <= x) && (x < this.Width) && (0 <= y) && (y < this.Height)
}

func (this *Maze) EmptyAt(x, y int) bool {
  cell := x * this.Width + y
  return !this.Maze[cell]
}

func (this *Maze) Print() {
  out := "\n"
  for row := 0; row < this.Width; row ++ {
    for col := 0; col < this.Height; col ++ {
      if this.Maze[row * this.Width + col] {
        out += "1"
      } else {
        out += "0"
      }
    }
    out += "\n"
  }
  log.Print(out)
}

func drawCells(args *Args, canvas *svg.SVG) {
  width := args.Width
  height := args.Height
  var cells []bool
  switch method := args.Method; method {
    case "backtracking":
      cells = backtracking(args)
    case "random":
      cells = completelyRandom(args)
    default:
      cells = completelyRandom(args)
  }
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

func completelyRandom(args *Args) []bool {
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

func backtracking(args *Args) []bool {
  height := args.Height
  width := args.Width
  path := newStack(height * width)
  cells := make([]bool, height * width)
  maze := Maze{ Maze: cells, Height: height, Width: width}
  directions := [4][2]int{ up, down, left, right }

  // there is more branching if we "seed" in the center
  path.Push(Cell{ X: width / 2, Y: height / 2})

  /**
   * Add current cell to stack
   * while stack is not empty
   * check a random direction
   * if direction is good, fill that cell and add to stack
   * if all directions are not good, pop that cell from the stack as it is a dead end
   */
  for path.Len() > 0 {
    // record that we visited this spot
    curCel := path.Peek()
    maze.IncrementCell(curCel.X, curCel.Y)

    // shuffle around directions
    for i := range directions {
      j := rand.Intn(i + 1)
      directions[i], directions[j] = directions[j], directions[i]
    }

    deadEnd := true
    for _, dir := range directions {
      destination := Cell{ X: curCel.X + dir[0] * 2, Y: curCel.Y + dir[1] * 2}

      if maze.InBounds(destination.X, destination.Y) && maze.EmptyAt(destination.X, destination.Y) {
        onTheWay := Cell{ X: curCel.X + dir[0], Y: curCel.Y + dir[1]}
        maze.IncrementCell(onTheWay.X, onTheWay.Y)
        maze.IncrementCell(destination.X, destination.Y)
        path.Push(destination)

        deadEnd = false
        break
      }
    }

    if(deadEnd) {
      _ = path.Pop()
    }
  }

  return maze.Maze
}
