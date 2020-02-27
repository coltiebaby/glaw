package glaw

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

	resp, err := c.Do(req.NewHttpRequestWithCtx(ctx))
	if err != nil {
		return ci, err
	}

	err = ProcessRequest(resp, &ci)
	return ci, err
}

type ChampionInfo struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}
