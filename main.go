package main

import (
  "strings"
  "strconv"

  "log"
  "os"
  "errors"

  "github.com/urfave/cli"
  "github.com/SivanMehta/mazes/generator"
)

var (
  cellDimensions string
  actualDimensions string
)

/**
 * Validate the arguments, generating the args along the way
 * @type {[type]}
 */
func checkArgs(c *cli.Context) error {
  cd := strings.Split(cellDimensions, "x")
  ad := strings.Split(actualDimensions, "x")
  var dims [4]int
  if(len(cd) != 2) {
    return errors.New("You must provide 2 cell dimensions")
  }

  if(len(ad) != 2) {
    return errors.New("You must provide 2 actual dimensions")
  }

  allDims := append(cd, ad...)
  for i, dim := range allDims {
    value, err := strconv.Atoi(dim)
    if err != nil || value < 1 {
      return errors.New("Dimension must be positive numbers")
    }
    dims[i] = value
  }

  args := generator.Args{
    Width: dims[0],
    Height: dims[1],
    PixelsX: dims[2],
    PixelsY: dims[3],
    PathColor: "#FFFFFF",
    WallColor: "#000000",
    Method: "backtracking",
  }

  args.CellSizeX = args.PixelsX / args.Width
  args.CellSizeY = args.PixelsY / args.Height

  if(args.CellSizeX * args.Width != args.PixelsX || args.CellSizeY * args.Height != args.PixelsY) {
    return errors.New("Actual dimensions must be positive multiples of cell dimensions")
   }

  // all is good and parse, let's fire off the generation
  generator.GenerateMaze(&args)
  return nil
}

func main() {
  app := cli.NewApp()
  app.Name = "mazes"
  app.Usage = "create amazing mazes in SVG"

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name:        "cell-dimensions, cd",
      Value:       "100x100",
      Usage:       "dimensions of the maze, in cells",
      Destination: &cellDimensions,
    },
    cli.StringFlag{
      Name:        "actual-dimensions, ad",
      Value:       "1000x1000",
      Usage:       "dimensions of the maze, in pixels",
      Destination: &actualDimensions,
    },
  }

  app.Action = checkArgs

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
