package message

import (
	"errors"
	"fmt"
	"log"
)

type result []string

func (r *result) Add(err error, in ...string) {
	if err != nil {
		switch {
		case errors.Is(err, errHostNotMath):
			return
		default:
			*r = append(*r, fmt.Sprintf("Error | ,%v", err))
			log.Println(err)
			return
		}
	}
	*r = append(*r, in...)
}
