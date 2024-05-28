package schema

import "time"

type Presece struct {
	id          int
	name        string
	presence_at time.Time
}
