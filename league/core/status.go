package core

type Platform struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Locales      []string `json:"locales"`
	Maintenances []Status `json:"maintenances"`
	Incidents    []Status `json:"incidents"`
}

type Status struct {
	Id                 int       `json:"id"`
	Maintenance_status string    `json:"maintenance_status"`
	Incident_severity  string    `json:"incident_severity"`
	Titles             []Content `json:"titles"`
	Updates            []Update  `json:"updates"`
	Created            string    `json:"created_at"`
	Archive            string    `json:"archive_at"`
	Updated            string    `json:"updated_at"`
	Platforms          []string  `json:"platforms"`
}

type Content struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

type Update struct {
	Id                int       `json:"id"`
	Author            string    `json:"author"`
	Publish           bool      `json:"publish"`
	Publish_locations []string  `json:"publish_locations"`
	Translations      []Content `json:"translations"`
	Created           string    `json:"created_at"`
	Updated           string    `json:"updated_at"`
}
