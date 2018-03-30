package main

import (
  "log"
  "strings"
  "strconv"

  "os"
  "errors"

  "github.com/urfave/cli"
  "github.com/SivanMehta/mazes/generator"
)

var (
  cellDimensions string
  actualDimensions string
  args [4]int
)

func checkArgs(c *cli.Context) error {
  cd := strings.Split(cellDimensions, "x")
  ad := strings.Split(actualDimensions, "x")

  if(len(cd) != 2) {
    return errors.New("You must provide 2 cell dimensions")
  }

  if(len(ad) != 2) {
    return errors.New("You must provide 2 actual dimensions")
  }

  allDims := append(cd, ad...)
  for i, dim := range allDims {
    value, err := strconv.Atoi(dim)
    if err != nil {
      return errors.New("Invalid dimensions")
    }
    args[i] = value
  }

  // all is good and parse, let's fire off the generation
  generator.GenerateMaze(args)
  return nil
}

func main() {
  app := cli.NewApp()
  app.Name = "mazes"
  app.Usage = "maze a maze in SVG"

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
