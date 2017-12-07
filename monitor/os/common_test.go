package os

import (
	"testing"
)

func Test_Gethostname(t *testing.T) {
	t.Log(GetHostname())
}