package mass

import "time"

type Ask struct {
	Description   string
	Abstract      string
	Name          string
	Location      string
	Contact       string
	Key           string
	ID            string
	TimeToLive    time.Duration
	DateRequested time.Time
	DateFulfilled time.Time
}

type Offer struct {
	Description string
	Abstract    string
	Name        string
	Location    string
	Contact     string
	Key         string
	ID          string
	TimeToLive  time.Duration
}
