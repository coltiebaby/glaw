package glaw

import (
	"fmt"
	"strings"
)

type Region int

func (r Region) Base() string {
	return fmt.Sprintf(`%s.api.riotgames.com`, strings.ToLower(RegionsPlatform[r]))
}

func (r Region) String() string {
	return Regions[r]
}

const (
	REGION_NA Region = iota + 1
	REGION_BR
	REGION_EUNE
	REGION_EUW
	REGION_JP
	REGION_KR
	REGION_LAN
	REGION_LAS
	REGION_OCE
	REGION_TR
	REGION_RU
	REGION_PBE
	REGION_AMERICAS
	REGION_ASIA
	REGION_EUROPE
)

var Regions = map[Region]string{
	REGION_NA:       `NA`,
	REGION_BR:       `BR`,
	REGION_EUNE:     `EUNE`,
	REGION_EUW:      `EUW`,
	REGION_JP:       `JP`,
	REGION_KR:       `KR`,
	REGION_LAN:      `LAN`,
	REGION_LAS:      `LAS`,
	REGION_OCE:      `OCE`,
	REGION_TR:       `TR`,
	REGION_RU:       `RU`,
	REGION_PBE:      `PBE`,
	REGION_AMERICAS: "americas",
	REGION_ASIA:     "asia",
	REGION_EUROPE:   "europe",
}

var RegionsPlatform = map[Region]string{
	REGION_NA:       `NA1`,
	REGION_BR:       `BR1`,
	REGION_EUNE:     `EUN1`,
	REGION_EUW:      `EUW1`,
	REGION_JP:       `JP1`,
	REGION_KR:       `KR`,
	REGION_LAN:      `LA1`,
	REGION_LAS:      `LA2`,
	REGION_OCE:      `OC1`,
	REGION_TR:       `TR1`,
	REGION_RU:       `RU`,
	REGION_PBE:      `PBE1`,
	REGION_AMERICAS: "americas",
	REGION_ASIA:     "asia",
	REGION_EUROPE:   "europe",
}
