package main

import "github.com/smartfood-capstone/backend/cmd"

func main() {
	cli := cmd.New()
	err := cli.Execute()
	if err != nil {
		panic(err)
	}
}
