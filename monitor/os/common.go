package os

import (
	"os"
)

func GetHostname() string {
	hostname,err := os.Hostname()
	if err != nil {
		return err.Error()
	}
	return hostname
}