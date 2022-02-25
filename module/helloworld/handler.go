package helloworld

import (
	"fmt"
)

func (mo *HellWorld) say(name string) (r string, err error) {
	return fmt.Sprintf("hi %v", name), nil
}
