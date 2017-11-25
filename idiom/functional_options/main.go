package main

import (
	"fmt"

	"./file"
)

func main() {
	err := file.New("./empty.txt")
	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	err = file.New("./file.txt", file.UID(1), file.Contents("Lorem Ipsum Dolor Amet"))
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}
