package todo

import (
	"fmt"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	Value "github.com/kkserver/kk-lib/kk/value"
	"reflect"
	"strings"
	"time"
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

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	var v = Todo{}

	if task.Pid != 0 {

		var p = TodoTask{}
		p.Id = task.Pid
		app.Handle(a, &p)

		if p.Result.Todo == nil {
			task.Result.Errno = ERROR_TODO_NOT_FOUND
			task.Result.Errmsg = fmt.Sprintf("Not found parent todo pid:%d", task.Pid)
			return nil
		}

		v.Pid = task.Pid
		v.Path = fmt.Sprintf("%s/%d/", p.Result.Todo.Path, task.Pid)
	} else {
		v.Pid = 0
		v.Path = "/"
	}

	v.Uid = task.Uid
	v.Title = task.Title
	v.Ctime = time.Now().Unix()
	v.Mtime = v.Ctime

	_, err = kk.DBInsert(db, &a.TodoTable, a.DB.Prefix, &v)

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Todo = &v

	return nil
}

func (S *TodoService) HandleTodoSetTask(a *TodoApp, task *TodoSetTask) error {

	if task.Id == 0 {
		task.Result.Errno = ERROR_TODO_NOT_FOUND_ID
		task.Result.Errmsg = "Not found id"
		return nil
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	var v = Todo{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, &a.TodoTable, a.DB.Prefix, " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {

		err := scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		if task.Pid != nil {

			var pid = Value.IntValue(reflect.ValueOf(task.Pid), 0)

			if pid != 0 {

				var p = TodoTask{}
				p.Id = pid
				app.Handle(a, &p)

				if p.Result.Todo == nil {
					task.Result.Errno = ERROR_TODO_NOT_FOUND
					task.Result.Errmsg = fmt.Sprintf("Not found parent todo pid:%d", pid)
					return nil
				}

				if v.Pid != 0 && strings.HasPrefix(p.Result.Todo.Path, v.Path) {
					task.Result.Errno = ERROR_TODO_PATH
					task.Result.Errmsg = fmt.Sprintf("Not move to sub todo pid:%d", pid)
					return nil
				}

				v.Pid = pid
				v.Path = fmt.Sprintf("%s/%d/", p.Result.Todo.Path, pid)

			} else {
				v.Pid = 0
				v.Path = "/"
			}
		}

		if task.Uid != nil {
			v.Uid = Value.IntValue(reflect.ValueOf(task.Uid), 0)
		}

		if task.Title != nil {
			v.Title = Value.StringValue(reflect.ValueOf(task.Title), "")
		}

		v.Mtime = time.Now().Unix()

		_, err = kk.DBUpdate(db, &a.TodoTable, a.DB.Prefix, &v)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		task.Result.Todo = &v

	} else {
		task.Result.Errno = ERROR_TODO_NOT_FOUND
		task.Result.Errmsg = "Not found todo"
		return nil
	}

	return nil
}

func (S *TodoService) HandleTodoTask(a *TodoApp, task *TodoTask) error {

	if task.Id == 0 {
		task.Result.Errno = ERROR_TODO_NOT_FOUND_ID
		task.Result.Errmsg = "Not found id"
		return nil
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	var v = Todo{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, &a.TodoTable, a.DB.Prefix, " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {

		err := scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		task.Result.Todo = &v

	} else {
		task.Result.Errno = ERROR_TODO_NOT_FOUND
		task.Result.Errmsg = "Not found todo"
		return nil
	}

	return nil
}

func (S *TodoService) HandleTodoRemoveTask(a *TodoApp, task *TodoRemoveTask) error {

	if task.Id == 0 {
		task.Result.Errno = ERROR_TODO_NOT_FOUND_ID
		task.Result.Errmsg = "Not found id"
		return nil
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	var v = Todo{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, &a.TodoTable, a.DB.Prefix, " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {

		err := scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		_, err = kk.DBDelete(db, &a.TodoTable, a.DB.Prefix, " WHERE path LIKE ? OR id=?", fmt.Sprintf("%s/%d/%%", v.Path, v.Id), v.Id)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

	} else {
		task.Result.Errno = ERROR_TODO_NOT_FOUND
		task.Result.Errmsg = "Not found todo"
		return nil
	}
	return nil
}

func (S *TodoService) HandleTodoQueryTask(a *TodoApp, task *TodoQueryTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	var todos = []Todo{}

	var args = []interface{}{}

	var sql = " WHERE 1"

	if task.Id != 0 {
		sql = sql + " AND id=?"
		args = append(args, task.Id)
	} else {

		if task.Pid != nil {
			sql = sql + " AND id=?"
			args = append(args, Value.IntValue(reflect.ValueOf(task.Pid), 0))
		}

		if task.Path != nil {
			sql = sql + " AND path LIKE ?"
			args = append(args, fmt.Sprintf("%s%%", Value.StringValue(reflect.ValueOf(task.Pid), "")))
		}

		var pageIndex = task.PageIndex
		var pageSize = task.PageSize

		if pageIndex < 1 {
			pageIndex = 1
		}

		if pageSize < 1 {
			pageSize = 10
		}

		if task.Counter {
			var counter = TodoQueryCounter{}
			counter.PageIndex = pageIndex
			counter.PageSize = pageSize
			counter.PageSize, err = kk.DBQueryCount(db, &a.TodoTable, a.DB.Prefix, sql, args...)
			if err != nil {
				task.Result.Errno = ERROR_TODO
				task.Result.Errmsg = err.Error()
				return nil
			}
		}

		if task.OrderBy == "asc" {
			sql = sql + " ORDER BY id ASC"
		} else {
			sql = sql + " ORDER BY id DESC"
		}

		sql = fmt.Sprintf("%s LIMIT %d,%d", sql, (pageIndex-1)*pageSize, pageSize)

		var v = Todo{}
		var scanner = kk.NewDBScaner(&v)

		rows, err := kk.DBQuery(db, &a.TodoTable, a.DB.Prefix, sql, args...)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		defer rows.Close()

		for rows.Next() {

			err = scanner.Scan(rows)

			if err != nil {
				task.Result.Errno = ERROR_TODO
				task.Result.Errmsg = err.Error()
				return nil
			}

			todos = append(todos, v)
		}
	}

	task.Result.Todos = todos

	return nil
}
