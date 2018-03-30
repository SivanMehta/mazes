package generator

import (
  "os"
  "github.com/ajstarks/svgo"
)

func GenerateMaze(args [4]int) {
  width := args[2]
	height := args[3]
	canvas := svg.New(os.Stdout)
	canvas.Startview(width, height, 0, 0, width, height)
	canvas.End()
}
