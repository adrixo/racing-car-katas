package tpms

import "testing"

type alarmTest struct {
	it       string
	expected bool
	actual   func() bool
}

func TestAlarmHappyPath(t *testing.T) {
	t.Run("the value is inside thresholds", func(t *testing.T) {
		want := true
		sensor := NewSensor()
		a := NewAlarm(sensor)
		got := a.Check()
		if got != want {
			t.Errorf("got %t want %t \n", got, want)
		}
	})
}

type sensorSpy struct{}

func (s sensorSpy) popNextPressurePsiValue() int {
	return 17
}

func TestAlarmWrongPSI(t *testing.T) {
	t.Run("the value is minor than the threshold", func(t *testing.T) {
		want := false
		sensor := &sensorSpy{}
		a := NewAlarm(sensor)

		got := a.Check()
		if got != want {
			t.Errorf("got %t want %t \n", got, want)
		}
	})
}
