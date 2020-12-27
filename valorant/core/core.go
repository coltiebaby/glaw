package core

type LN struct {
	ArAE string `json:"ar-AE"`
	DeDE string `json:"de-DE"`
	EnGB string `json:"en-GB"`
	EnUS string `json:"en-US"`
	EsES string `json:"es-ES"`
	EsMX string `json:"es-MX"`
	FrFR string `json:"fr-FR"`
	IDID string `json:"id-ID"`
	ItIT string `json:"it-IT"`
	JaJP string `json:"ja-JP"`
	KoKR string `json:"ko-KR"`
	PlPL string `json:"pl-PL"`
	PtBR string `json:"pt-BR"`
	RuRU string `json:"ru-RU"`
	ThTH string `json:"th-TH"`
	TrTR string `json:"tr-TR"`
	ViVN string `json:"vi-VN"`
	ZhCN string `json:"zh-CN"`
	ZhTW string `json:"zh-TW"`
}

type Item struct {
	Name           string `json:"name"`
	LocalizedNames LN     `json:"localizedNames"`
	Id             string `json:"id"`
	AssetName      string `json:"assetName"`
	AssetPath      string `json:"assetPath"`
}

type Act struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	LocalizedNames LN     `json:"localizedNames"`
	IsActive       bool   `json:"isActive"`
}

type Content struct {
	Version      string `json:"version"`
	Characters   []Item `json:"characters"`
	Maps         []Item `json:"maps"`
	Chromas      []Item `json:"chromas"`
	Skins        []Item `json:"skins"`
	SkinLevels   []Item `json:"skinLevels"`
	Equips       []Item `json:"equips"`
	GameModes    []Item `json:"gameModes"`
	Sprays       []Item `json:"sprays"`
	SprayLevels  []Item `json:"sprayLevels"`
	Charms       []Item `json:"charms"`
	CharmLevels  []Item `json:"charmLevels"`
	PlayerCards  []Item `json:"playerCards"`
	PlayerTitles []Item `json:"playerTitles"`
	Acts         []Act  `json:"acts"`
}
