package tpms

import (
	"testing"
)

func newSensorSpy(mockValue int) Sensor {
	return &sensor{
		offset: func() int {
			return 0
		},
		samplePressure: func() int {
			return mockValue
		},
	}
}

func TestAlarmHappyPath(t *testing.T) {
	t.Run("the value is inside thresholds", func(t *testing.T) {
		want := true
		sensor := newSensorSpy(0)
		a := NewAlarm(sensor)
		got := a.Check()
		if got != want {
			t.Errorf("got %t want %t \n", got, want)
		}
	})
	t.Run("the value is below thresholds", func(t *testing.T) {
		want := true
		sensor := newSensorSpy(30)
		a := NewAlarm(sensor)
		got := a.Check()
		if got != want {
			t.Errorf("got %t want %t \n", got, want)
		}
	})
}

func TestAlarmWrongPSI(t *testing.T) {
	t.Run("the value is minor than the threshold", func(t *testing.T) {
		want := false
		sensor := newSensorSpy(17)
		a := NewAlarm(sensor)

		got := a.Check()
		if got != want {
			t.Errorf("got %t want %t \n", got, want)
		}
	})
}
