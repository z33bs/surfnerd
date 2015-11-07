package wavewatch

type EastCoastModel struct {
}

func (e *EastCoastModel) Name() string {
	return "multi_1.at_10m"
}

func (e *EastCoastModel) Description() string {
	return "Multi-grid wave model: US East Coast 10 arc-min grid"
}

func (e *EastCoastModel) BottomLeftCoord() *Location {
	return &Location{0.00, 260.00}
}

func (e *EastCoastModel) TopRightCoord() *Location {
	return &Location{55.00011, 310.00011}
}

func (e *EastCoastModel) LocationResolution() float64 {
	return 0.167
}

func (e *EastCoastModel) ContainsLocation(loc *Location) bool {
	if loc.Latitude > e.BottomLeftCoord().Latitude && loc.Latitude < e.TopRightCoord().Latitude {
		if loc.Longitude > e.BottomLeftCoord().Longitude && loc.Longitude < e.TopRightCoord().Longitude {
			return true
		}
	}
	return false
}

func (e *EastCoastModel) TimeResolution() float64 {
	return 0.125
}
