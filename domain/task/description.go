package task

import (
	"fmt"
)

type Description string

func (r Description) validate() error {
	if n := len(r); n > 10 {
		return fmt.Errorf("タスク詳細は10文字以内にしてください (n=%d)", n)
	}
	return nil
}
