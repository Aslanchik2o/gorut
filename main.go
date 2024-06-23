package main

import (
	"fmt"
	"time"
)

var action = []string{
	"logged in",
	"logged out",
	"create record",
	"delate record",
	"update record",

}

type logItems struct {
	action   string
	timestep time.Time
}
type User struct {
	id    int
	email string
	logs [] logItems
}

func (u User)  getActiviteInfo()string{
	out := fmt.Sprintf("ID: %d | email: %s\nActivity Log:\n", u.id ,u.email)
	for i, item := range u.logs{
		out += fmt.Sprintf("%d. [%s] at %s\n, ", i, item.action, item.timestep)

	}
return out
}
func main() {
	u := User{
		id:1,
		email: "gagagabugaga@gmail.com",
		logs: []logItems{
			{action[0], time.Now()},
			{action[3], time.Now()},
			{action[0], time.Now()},
			{action[2], time.Now()},
			{action[0], time.Now()},
			{action[1], time.Now()},
		},
	
	}
	fmt.Println(u.getActiviteInfo())

}