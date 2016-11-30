package main

import (
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-todo/todo"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.Llongfile | log.LstdFlags)

	config := "./app.ini"

	if len(os.Args) > 1 {
		config = os.Args[1]
	}

	a := todo.TodoApp{}

	err := app.Load(&a, config)

	if err != nil {
		log.Panicln(err)
	}

	app.Obtain(&a)

	app.Handle(&a, &app.InitTask{})

	app.Handle(&a, &todo.TodoCreateTask{})

	kk.DispatchMain()

	app.Recycle(&a)

}
