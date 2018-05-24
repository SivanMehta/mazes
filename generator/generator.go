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
  CellSizeX int
  CellSizeY int
  Method string
  PathColor string
  WallColor string
}

func GenerateMaze(args *Args) {
	canvas := svg.New(os.Stdout)
	canvas.Startview(args.PixelsX, args.PixelsY, 0, 0, args.PixelsX, args.PixelsY)
  drawCells(args, canvas)
	canvas.End()
}
