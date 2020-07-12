package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := &cli.App{
		Name:    "hl",
		Usage:   "get header list.",
		Version: "0.1.0",
		Action: func(c *cli.Context) error {
			f, err := os.Open("test.csv")
			if err != nil {
				log.Fatal(err)
			}

			r := csv.NewReader(f)

			header, err := r.Read()
			fmt.Println(header)
			return nil

		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
