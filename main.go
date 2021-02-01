package main

import (
	"fmt"
	"os"
	"runtime"
	"github.com/losanni/todo/api"
)

func help() {
	fmt.Print(` Use the following comands. Wrap "todonames" and "texttoadd" in quotes

List todos:
	todo

Add todo:
	todo add todoname

Read todo:
	todo read todoname

Update todo description:
	todo edit todoname texttoadd

Remove todo:
	todo remove todoname

`)
}

func main() {

	if runtime.GOOS == "windows" {
		fmt.Println("RIP Windows user")
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		api.Alltodos()
	} else {
		if os.Args[1] == "help" {
			help()
		} else if os.Args[1] == "add" {
			api.Createtodo(os.Args[2])
		} else if os.Args[1] == "remove" {
			api.Removetodo(os.Args[2])
		} else if os.Args[1] == "read" {
			api.Readtodo(os.Args[2])
		} else if os.Args[1] == "edit" {
			api.Updatetodo(os.Args[2], os.Args[3]) // 2 is the file and 3 is text to write
		} else {
			help()
		}
	}

}

// Api

//api.Alltodos() lists all things todo
//api.Createtodo() creates a todo
//api.Readtodo() reads info about one todo
//api.Removetodo() deletes a todo
//api.Updatetodo() writes to the todo file
