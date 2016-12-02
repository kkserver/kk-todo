package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoService struct {
	app.Service
	Create *TodoCreateTask
	Set    *TodoSetTask
	Get    *TodoTask
	Remove *TodoRemoveTask
	Query  *TodoQueryTask
}

func (S *TodoService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *TodoService) HandleTodoCreateTask(a *TodoApp, task *TodoCreateTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoService) HandleTodoSetTask(a *TodoApp, task *TodoSetTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoService) HandleTodoTask(a *TodoApp, task *TodoTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoService) HandleTodoRemoveTask(a *TodoApp, task *TodoRemoveTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoService) HandleTodoQueryTask(a *TodoApp, task *TodoQueryTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}
