package main

import (
	"fmt"
	"time"
)

// Author commit author
type Author struct {
	name string
	email string
	time time.Time
}

func (a *Author) ToString() string {
	return fmt.Sprintf("%s <%s> %s", a.name, a.email, a.time.String())
}