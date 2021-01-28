package logic

import "time"

func Auth(name string) string {
	if name == "cool" {
		time.Sleep(3 * time.Second)
		return "OK"
	}
	return "Wrong user name"
}