package register

import (
	"testing"
)

func Test_Register(t *testing.T) {
	err := Register("../conf/app.conf")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("ok")
}