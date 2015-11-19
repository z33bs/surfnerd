package wavewatch

import (
	"fmt"
	"testing"
)

func TestWaveWatchFetch(t *testing.T) {
	riLocation := &Location{41.336872, 288.635294}
	forecast := FetchWaveWatchData(riLocation)
	if forecast == nil {
		t.FailNow()
	}

	fmt.Println(forecast.forecastData[0].Time)
}

// func TestWaveWatchParse(t *testing.T) {
//  fileData, err := ioutil.ReadFile("resources/east_coast_model_example")
//  if err != nil {
//      t.FailNow()
//  }

//  modelData := parseRawWaveWatchData(fileData)
//  if modelData == nil {
//      t.FailNow()
//  }

//  fmt.Println(modelData["time"][0])
// }
