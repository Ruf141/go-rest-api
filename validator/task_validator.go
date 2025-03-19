package validator

import "go-rest-api/model"

type ITaskValidator interface {
	TaskValidate(task model.Task)error
}

type TaskValidator struct {}

