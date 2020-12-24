package league

import (
	"context"

	"github.com/coltiebaby/glaw"
)

type CodeRequest struct {
	SummonerId string
	Region     glaw.Region
}

func (c *Client) ThirdPartyCode(ctx context.Context, cr CodeRequest) (code string, err error) {
	uri := fmt.Sprintf(`third-party-code/by-summoner/%s`, cr.SummonerId)
	req := NewRequest("GET", "platform", uri, cr.Region, glaw.V4)

	err = c.Do(ctx, req, &code)
	return code, err
}
