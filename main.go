package main

import (
	"github.com/smartfood-capstone/backend/internal/app"
)

func main() {
	// cli := cmd.New()
	// err := cli.Execute()

	cmd := app.NewStartCmd()
	cmd.Start()

	// if err != nil {
	// 	panic(err)
	// }
}
