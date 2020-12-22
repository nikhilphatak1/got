package got

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

// NewAuthor Author constructor
func NewAuthor(name string, email string, time time.Time) *Author {
    author := Author{}
    author.name = name
    author.email = email
    author.time = time
    return &author
}

// String convert author to string
func (a *Author) String() string {
    return fmt.Sprintf("%s <%s> %s", a.name, a.email, a.time.String())
}