package types

import (
	"errors"
	"fmt"

	"github.com/VadymNovoselskyi/GoCourse/notes"
)

func PrintAnything(value any) {
	switch valueType := value.(type) {
	case string:
		fmt.Printf("%v is a str: %v", value, valueType)
	default:
		fmt.Printf("%v is not a str: %v", value, valueType)
	}
}

func ToInt(value any) (int, error) {
	intValue, ok := value.(int)

	if ok {
		return intValue, nil
	}
	return -1, errors.New("value passed in is not an int")
}

func IsNote(value any) bool {
	_, ok := value.(notes.Note)
	return ok
}

func Add[T int | float64 | string](lhs T, rhs T) T {
	return lhs + rhs
}
