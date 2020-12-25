package core

import (
	"encoding/json"
	"strconv"
	"time"
)

// include url:port (ex http://example.com/callback:80, https://securedexample.com/callback:443)
type TournamentProviderRegistration struct {
	Url    string `json:"url"`
	Region string `json:"region"`
}

type TournamentRegistration struct {
	ProviderId int    `json:"providerId"`
	Name       string `json:"name"`
}

type TournamentCodeRegistration struct {
	TeamSize      int       `json:"teamSize"`
	GameType      GameType  `json:"pickType"`
	Map           Map       `json:"mapType"`
	SpectatorType Spectator `json:"spectatorType"`

	// Optional
	SummonerIds []string `json:"allowedSummonerIds"`
	Metadata    string   `json:"metadata"`
}

type TournamentEvents struct {
	Events []TournamentEvent `json:"eventList"`
}

type event struct {
	SummonerId string `json:"summonerId"`
	EventType  string `json:"eventType"`
	Timestamp  string `json:"timestamp"`
}
type TournamentEvent struct {
	SummonerId string    `json:"summonerId"`
	EventType  string    `json:"eventType"`
	Timestamp  time.Time `json:"timestamp"`
}

func (te *TournamentEvent) UnmarshalJSON(b []byte) error {
	var e event
	if err := json.Unmarshal(b, &e); err != nil {
		return err
	}

	ts, err := strconv.ParseInt(e.Timestamp, 10, 64)
	if err != nil {
		return err
	}

	t := TournamentEvent{
		SummonerId: e.SummonerId,
		EventType:  e.EventType,
		Timestamp:  time.Unix(ts, 0),
	}

	*te = t

	return nil
}
