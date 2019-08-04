package matches

type Position struct {
	Y int `json:"y"`
	X int `json:"x"`
}

type ParticipantFrame struct {
	TotalGold           int      `json:"totalGold"`
	TeamScore           int      `json:"teamScore"`
	ParticipantID       int      `json:"participantId"`
	Level               int      `json:"level"`
	CurrentGold         int      `json:"currentGold"`
	MinionsKilled       int      `json:"minionsKilled"`
	DominionScore       int      `json:"dominionScore"`
	Position            Position `json:"position"`
	Xp                  int      `json:"xp"`
	JungleMinionsKilled int      `json:"jungleMinionsKilled"`
}

type Event struct {
	Timestamp               int      `json:"timestamp"`
	Type                    string   `json:"type"`
	CreatorID               int      `json:"creatorId,omitempty"`
	WardType                string   `json:"wardType,omitempty"`
	SkillSlot               int      `json:"skillSlot,omitempty"`
	LevelUpType             string   `json:"levelUpType,omitempty"`
	ParticipantID           int      `json:"participantId,omitempty"`
	ItemID                  int      `json:"itemId,omitempty"`
	KillerID                int      `json:"killerId,omitempty"`
	BuildingType            string   `json:"buildingType,omitempty"`
	TowerType               string   `json:"towerType,omitempty"`
	TeamID                  int      `json:"teamId,omitempty"`
	AssistingParticipantIds []int    `json:"assistingParticipantIds,omitempty"`
	Position                Position `json:"position,omitempty"`
	LaneType                string   `json:"laneType,omitempty"`
	VictimID                int      `json:"victimId,omitempty"`
}

type Frame struct {
	Timestamp         int                         `json:"timestamp"`
	ParticipantFrames map[string]ParticipantFrame `json:"participantFrames"`
	Events            []Event                     `json:"events"`
}

type Timeline struct {
	Frames        []Frame `json:"frames"`
	FrameInterval int     `json:"frameInterval"`
}

func (m Timeline) Filter(by ...string) {
	for i, frame := range m.Frames {
		events := filter(frame.Events, by)
		frame.Events = events
		m.Frames[i] = frame
	}
}
