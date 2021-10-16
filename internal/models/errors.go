package models

import (
	"errors"
)

var ErrConflictInsert error = errors.New("conflict")
var ErrUrlSetToDel error = errors.New("conflict")
