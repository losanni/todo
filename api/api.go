package api

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

var homedir = os.Getenv("HOME")
var tododir = homedir + "/.local/share/todo/"


func Alltodos()  {
	files, err := ioutil.ReadDir(tododir)
	if err != nil {
		fmt.Println("ERROR LISTING TODOS")
		log.Fatal(err)
	}
	fmt.Println("Todo list:")
	for _, file := range files {
		fmt.Println("	" + file.Name())
	}
}

func Readtodo(todoname string) {
	file, err := ioutil.ReadFile(tododir + todoname)
	if err != nil {
		fmt.Println("ERROR READING FILE")
		log.Fatal(err)
	}
	fmt.Println(string(file))
}

func Createtodo(todoname string)  {
	file, err := os.Create(tododir + todoname)
	if err != nil {
		fmt.Println("COULD NOT CREATE FILE")
		log.Fatal(err)
	}
	file.Close()
	fmt.Println("Created todo with the name " + todoname)
}

func Updatetodo(todoname string, texttowrite string)  {
	err := ioutil.WriteFile(tododir + todoname, []byte(texttowrite), 0777)
	if err != nil {
		fmt.Println("ERROR WRITING TO FILE")
		log.Fatal(err)
	}
}
func Removetodo(todoname string)  {
	err := os.Remove(tododir + todoname)
    if err != nil {
		fmt.Println("COULD NOT DELETE FILE")
        log.Fatal(err)
    }
    fmt.Println("Todo '" + todoname + "' removed")
}