package generator

import (
  "os"
  "github.com/ajstarks/svgo"
  "math/rand"
)

type Args struct {
  Width int
  Height int
  PixelsX int
  PixelsY int
  CellSizeX int
  CellSizeY int
  PathColor string
  WallColor string
}

func GenerateMaze(args *Args) {
  width := args.Width
	height := args.Height
	canvas := svg.New(os.Stdout)
	canvas.Startview(args.PixelsX, args.PixelsY, 0, 0, args.PixelsX, args.PixelsY)
  for row := 0; row < height; row ++ {
    for col := 0; col < width; col ++ {
      var color string
      if(rand.Intn(2) > 0) {
        color = "fill:#FFF;"
      } else {
        color = "fill:#000;"
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
	canvas.End()
}
