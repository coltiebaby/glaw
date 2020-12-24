package core

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
	TeamSize      int            `json:"teamSize"`
	GameType      core.GameType  `json:"pickType"`
	Map           core.Map       `json:"mapType"`
	SpectatorType core.Spectator `json:"spectatorType"`

	// Optional
	SummonerIds []string `json:"allowedSummonerIds"`
	Metadata    string   `json:"metadata"`
}
