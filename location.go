package surfnerd

import (
	"math"
)

// Container holding location information.
type Location struct {
	Latitude     float64 `xml:"lat,attr"`
	Longitude    float64 `xml:"lon,attr"`
	Elevation    float64 `xml:"elev,attr"`
	LocationName string  `xml:"name,attr"`
}

// Get an adjusted longitude that will be + or - 180 degrees
func (l Location) AdjustedLongitude() float64 {
	if l.Longitude > 180 {
		return l.Longitude - 360.0
	} else {
		return l.Longitude
	}
}

// Get an adjusted latitude that will be + or - 85
func (l Location) AdjustedLatitude() float64 {
	if l.Latitude > 85 {
		return l.Latitude - 360.0
	} else {
		return l.Latitude
	}
}

// Get the lat and long components of the distance between two locations
func (l Location) ComponentDistanceTo(otherLoc Location) (latDist, lonDist float64) {
	latDist = math.Abs(l.Latitude - otherLoc.Latitude)
	lonDist = math.Abs(l.Longitude - otherLoc.Longitude)
	return
}

// Get the estimated distance vector between two location points
func (l Location) DistanceTo(otherLoc Location) float64 {
	latDist, lonDist := l.ComponentDistanceTo(otherLoc)
	return math.Sqrt(math.Pow(latDist, 2) + math.Pow(lonDist, 2))
}

// Create a new Location object from a given latitude and longitude pair
// The latitude must be in degress N
// The longitude must be in degrees E
// If the values are out of range nil is returned
func NewLocationForLatLong(lat, lon float64) Location {
	return Location{Latitude: lat, Longitude: lon}
}
