package matches

// Timeline utilities

// Get the only events you would like to have
// events := filter(events, CHAMPION_KILL, WARD_KILL)
func filter(events []Event, by []string) []Event {
	filters := make(map[string]struct{})
	var empty struct{}

	for _, b := range by {
		filters[b] = empty
	}

	var sorted []Event

	for _, v := range events {
		if _, ok := filters[v.Type]; ok {
			sorted = append(sorted, v)
		}
	}

	return sorted
}

// Timeline Event Filters
const (
	CHAMPION_KILL      string = "CHAMPION_KILL"
	WARD_PLACED        string = "WARD_PLACED"
	WARD_KILL          string = "WARD_KILL"
	BUILDING_KILL      string = "BUILDING_KILL"
	ELITE_MONSTER_KILL string = "ELITE_MONSTER_KILL"
	ITEM_PURCHASED     string = "ITEM_PURCHASED"
	ITEM_SOLD          string = "ITEM_SOLD"
	ITEM_DESTROYED     string = "ITEM_DESTROYED"
	ITEM_UNDO          string = "ITEM_UNDO"
	SKILL_LEVEL_UP     string = "SKILL_LEVEL_UP"
	ASCENDED_EVENT     string = "ASCENDED_EVENT"
	CAPTURE_POINT      string = "CAPTURE_POINT"
	PORO_KING_SUMMON   string = "PORO_KING_SUMMON"
)
