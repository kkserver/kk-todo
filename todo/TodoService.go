package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoService struct {
	app.Service
	Create  *TodoCreateTask
	Set     *TodoSetTask
	Get     *TodoTask
	Cancel  *TodoCancelTask
	Checkin *TodoCheckinTask
	Finish  *TodoFinishTask
	In      *TodoInTask
	Working *TodoWorkingTask
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

func (S *TodoService) HandleTodoCancelTask(a *TodoApp, task *TodoCancelTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoService) HandleTodoCheckinTask(a *TodoApp, task *TodoCheckinTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoService) HandleTodoFinishTask(a *TodoApp, task *TodoFinishTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoService) HandleTodoInTask(a *TodoApp, task *TodoInTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}

func (S *TodoService) HandleTodoWorkingTask(a *TodoApp, task *TodoWorkingTask) error {

	task.Result.Errno = ERROR_NOIMPLEMENTS
	task.Result.Errmsg = "no implements"

	return nil
}
