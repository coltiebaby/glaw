package matches

type Delta struct {
	Zero10 float64 `json:"0-10"`
	One020 float64 `json:"10-20"`
}

type ParticipantTimeline struct {
	Lane                    string `json:"lane"`
	ParticipantID           int    `json:"participantId"`
	Role                    string `json:"role"`
	GoldPerMinDeltas        Delta  `json:"goldPerMinDeltas"`
	CreepsPerMinDeltas      Delta  `json:"creepsPerMinDeltas"`
	XpPerMinDeltas          Delta  `json:"xpPerMinDeltas"`
	DamageTakenPerMinDeltas Delta  `json:"damageTakenPerMinDeltas"`
}

type Participant struct {
	Timeline                  ParticipantTimeline `json:"timeline"`
	Spell1ID                  int                 `json:"spell1Id"`
	ParticipantID             int                 `json:"participantId"`
	Spell2ID                  int                 `json:"spell2Id"`
	TeamID                    int                 `json:"teamId"`
	Stats                     Stat                `json:"stats"`
	ChampionID                int                 `json:"championId"`
	HighestAchievedSeasonTier string              `json:"highestAchievedSeasonTier,omitempty"`
}

type Stat struct {
	FirstBloodAssist               bool `json:"firstBloodAssist"`
	VisionScore                    int  `json:"visionScore"`
	MagicDamageDealtToChampions    int  `json:"magicDamageDealtToChampions"`
	LargestMultiKill               int  `json:"largestMultiKill"`
	TotalTimeCrowdControlDealt     int  `json:"totalTimeCrowdControlDealt"`
	LongestTimeSpentLiving         int  `json:"longestTimeSpentLiving"`
	Perk1Var1                      int  `json:"perk1Var1"`
	Perk1Var3                      int  `json:"perk1Var3"`
	Perk1Var2                      int  `json:"perk1Var2"`
	TripleKills                    int  `json:"tripleKills"`
	Perk5                          int  `json:"perk5"`
	Perk4                          int  `json:"perk4"`
	PlayerScore9                   int  `json:"playerScore9"`
	PlayerScore8                   int  `json:"playerScore8"`
	Kills                          int  `json:"kills"`
	PlayerScore1                   int  `json:"playerScore1"`
	PlayerScore0                   int  `json:"playerScore0"`
	PlayerScore3                   int  `json:"playerScore3"`
	PlayerScore2                   int  `json:"playerScore2"`
	PlayerScore5                   int  `json:"playerScore5"`
	PlayerScore4                   int  `json:"playerScore4"`
	PlayerScore7                   int  `json:"playerScore7"`
	PlayerScore6                   int  `json:"playerScore6"`
	Perk5Var1                      int  `json:"perk5Var1"`
	Perk5Var3                      int  `json:"perk5Var3"`
	Perk5Var2                      int  `json:"perk5Var2"`
	TotalScoreRank                 int  `json:"totalScoreRank"`
	NeutralMinionsKilled           int  `json:"neutralMinionsKilled"`
	StatPerk1                      int  `json:"statPerk1"`
	StatPerk0                      int  `json:"statPerk0"`
	DamageDealtToTurrets           int  `json:"damageDealtToTurrets"`
	PhysicalDamageDealtToChampions int  `json:"physicalDamageDealtToChampions"`
	DamageDealtToObjectives        int  `json:"damageDealtToObjectives"`
	Perk2Var2                      int  `json:"perk2Var2"`
	Perk2Var3                      int  `json:"perk2Var3"`
	TotalUnitsHealed               int  `json:"totalUnitsHealed"`
	Perk2Var1                      int  `json:"perk2Var1"`
	Perk4Var1                      int  `json:"perk4Var1"`
	TotalDamageTaken               int  `json:"totalDamageTaken"`
	Perk4Var3                      int  `json:"perk4Var3"`
	LargestCriticalStrike          int  `json:"largestCriticalStrike"`
	LargestKillingSpree            int  `json:"largestKillingSpree"`
	QuadraKills                    int  `json:"quadraKills"`
	MagicDamageDealt               int  `json:"magicDamageDealt"`
	Item2                          int  `json:"item2"`
	Item3                          int  `json:"item3"`
	Item0                          int  `json:"item0"`
	Item1                          int  `json:"item1"`
	Item6                          int  `json:"item6"`
	Item4                          int  `json:"item4"`
	Item5                          int  `json:"item5"`
	Perk1                          int  `json:"perk1"`
	Perk0                          int  `json:"perk0"`
	Perk3                          int  `json:"perk3"`
	Perk2                          int  `json:"perk2"`
	Perk3Var3                      int  `json:"perk3Var3"`
	Perk3Var2                      int  `json:"perk3Var2"`
	Perk3Var1                      int  `json:"perk3Var1"`
	DamageSelfMitigated            int  `json:"damageSelfMitigated"`
	MagicalDamageTaken             int  `json:"magicalDamageTaken"`
	Perk0Var2                      int  `json:"perk0Var2"`
	FirstInhibitorKill             bool `json:"firstInhibitorKill"`
	TrueDamageTaken                int  `json:"trueDamageTaken"`
	Assists                        int  `json:"assists"`
	Perk4Var2                      int  `json:"perk4Var2"`
	GoldSpent                      int  `json:"goldSpent"`
	TrueDamageDealt                int  `json:"trueDamageDealt"`
	ParticipantID                  int  `json:"participantId"`
	PhysicalDamageDealt            int  `json:"physicalDamageDealt"`
	SightWardsBoughtInGame         int  `json:"sightWardsBoughtInGame"`
	TotalDamageDealtToChampions    int  `json:"totalDamageDealtToChampions"`
	PhysicalDamageTaken            int  `json:"physicalDamageTaken"`
	TotalPlayerScore               int  `json:"totalPlayerScore"`
	Win                            bool `json:"win"`
	ObjectivePlayerScore           int  `json:"objectivePlayerScore"`
	TotalDamageDealt               int  `json:"totalDamageDealt"`
	Deaths                         int  `json:"deaths"`
	PerkPrimaryStyle               int  `json:"perkPrimaryStyle"`
	PerkSubStyle                   int  `json:"perkSubStyle"`
	TurretKills                    int  `json:"turretKills"`
	FirstBloodKill                 bool `json:"firstBloodKill"`
	TrueDamageDealtToChampions     int  `json:"trueDamageDealtToChampions"`
	GoldEarned                     int  `json:"goldEarned"`
	KillingSprees                  int  `json:"killingSprees"`
	UnrealKills                    int  `json:"unrealKills"`
	FirstTowerAssist               bool `json:"firstTowerAssist"`
	FirstTowerKill                 bool `json:"firstTowerKill"`
	ChampLevel                     int  `json:"champLevel"`
	DoubleKills                    int  `json:"doubleKills"`
	InhibitorKills                 int  `json:"inhibitorKills"`
	FirstInhibitorAssist           bool `json:"firstInhibitorAssist"`
	Perk0Var1                      int  `json:"perk0Var1"`
	CombatPlayerScore              int  `json:"combatPlayerScore"`
	Perk0Var3                      int  `json:"perk0Var3"`
	VisionWardsBoughtInGame        int  `json:"visionWardsBoughtInGame"`
	PentaKills                     int  `json:"pentaKills"`
	TotalHeal                      int  `json:"totalHeal"`
	TotalMinionsKilled             int  `json:"totalMinionsKilled"`
	TimeCCingOthers                int  `json:"timeCCingOthers"`
	StatPerk2                      int  `json:"statPerk2"`
}
