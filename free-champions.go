package glaw

import (
	"context"
)

type ChampionRotationsRequest struct {
	Region Region
}

func (c *Client) ChampionRotations(ctx context.Context, fcr ChampionRotationsRequest) (ci ChampionInfo, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `platform`,
		Version: V3,
		Region:  fcr.Region,
		Uri:     `champion-rotations`,
	}

	resp, err := c.Do(ctx, req)
	if err != nil {
		return ci, err
	}

	err = ProcessResponse(resp, &ci)
	return ci, err
}

type ChampionInfo struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}
