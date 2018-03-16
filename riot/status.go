package riot

import (
    "encoding/json"
)

type Status struct {
	Name      string `json:"name"`
	RegionTag string `json:"region_tag"`
	Hostname  string `json:"hostname"`
    Services []StatusServices  `json:"services"`
	Slug    string   `json:"slug"`
	Locales []string `json:"locales"`
}

type StatusServices  struct {
    Status    string        `json:"status"`
    Incidents []interface{} `json:"incidents"`
    Name      string        `json:"name"`
    Slug      string        `json:"slug"`
}


func GetStatus() (Status) {
    // https://na1.api.riotgames.com/lol/status/v3/shard-data
    var status Status

    rr := &RiotRequest {
        Type: "status",
        Uri:  "shard-data",
    }

    resp := rr.GetData()
    json.Unmarshal(resp, &status)

    return status
}
