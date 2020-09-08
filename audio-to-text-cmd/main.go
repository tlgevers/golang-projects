package main

import (
	"fmt"
	speech "github.com/tlgevers/golang-projects/audio-to-text/pkg/speech"
	"github.com/urfave/cli/v2"
	"os"
)

const (
// file1   = "gs://audio-conversions/clpp.flac"
// output1 = "./output.txt"
)

func main() {
	var input string
	var output string
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Usage:       "Sets the name of the file to be converted from cloud storage",
				Destination: &input,
			},
			&cli.StringFlag{
				Name:        "output",
				Usage:       "Sets the name of the output file of where the conversion to text is written",
				Destination: &output,
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				input = c.Args().Get(0)
			}
			if input != "" {
				fmt.Println("input", input)
				fmt.Println("output", output)
				speech.Convert(input, output)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
