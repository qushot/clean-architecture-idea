package task

import (
	"fmt"
)

type Name string

func (r Name) validate() error {
	if n := len(r); n > 5 {
		return fmt.Errorf("タスク名は5文字以内にしてください (n=%d)", n)
	}
	return nil
}
