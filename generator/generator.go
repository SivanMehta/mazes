package generator

import (
  "os"
  "github.com/ajstarks/svgo"
)

type Args struct {
  Width int
  Height int
  PixelsX int
  PixelsY int
  CellHeightX int
  CellHeightY int
}

func GenerateMaze(args *Args) {
  width := args.PixelsX
	height := args.PixelsY
	canvas := svg.New(os.Stdout)
	canvas.Startview(width, height, 0, 0, width, height)
	canvas.End()
}
