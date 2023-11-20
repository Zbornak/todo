/*
A CLI todo list in Go
Author: zbornak
Date: 20/11/2023
*/

package main

import (
	"cli/todo"
	"fmt"
	"os"
	"strings"
)

const todoFilename = ".todo.json"

func main() {
	l := &todo.List{}
	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
