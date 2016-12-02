package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoUserService struct {
	app.Service
	Join   *TodoUserJoinTask
	Remove *TodoUserRemoveTask
	Get    *TodoUserTask
	Query  *TodoQueryTask
}

func (S *TodoUserService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *TodoUserService) HandleTodoUserJoinTask(a *TodoApp, task *TodoUserJoinTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoUserService) HandleTodoUserRemoveTask(a *TodoApp, task *TodoUserRemoveTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoUserService) HandleTodoUserTask(a *TodoApp, task *TodoUserTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoUserService) HandleTodoQueryTask(a *TodoApp, task *TodoQueryTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}
