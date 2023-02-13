package sys

import (
	"log"
	"testing"
)

type User struct {
}

func (u *User) UserMethod() {
	fullname := CurrentFuncFullName()
	name := CurrentFuncName()
	log.Println(fullname, name) // 2022/05/06 08:16:43 github.com/pescox/go-kit/exp.(*User).UserMethod UserMethod
}

func PkgMethod() {
	fullname := CurrentFuncFullName()
	name := CurrentFuncName()
	log.Println(fullname, name) // 2022/05/06 08:16:43 github.com/pescox/go-kit/exp.PkgMethod PkgMethod
}

func TestFuncName(t *testing.T) {
	PkgMethod()
	u := &User{}
	u.UserMethod()
	t.Log("done")
}
