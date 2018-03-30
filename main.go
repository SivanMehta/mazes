package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "mazes"
  app.Usage = "maze a maze in SVG"

  // flags
  var cellDimensions string
  var actualDimensions string

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

  app.Action = func(c *cli.Context) error {
    fmt.Printf("made a %s maze that is %s square pixels\n", cellDimensions, actualDimensions)
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
