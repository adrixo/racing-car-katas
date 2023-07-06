package tpms

type Alarm interface {
	Check() bool
}

type Sensor interface {
	popNextPressurePsiValue() int
}

type alarm struct {
	lowPressureThreshold  int
	highPressureThreshold int
	sensor                Sensor
	alarmOn               bool
}

func (a *alarm) Check() bool {
	p := a.sensor.popNextPressurePsiValue()

	if p < a.lowPressureThreshold || a.highPressureThreshold < p {
		a.alarmOn = true
	}

	return a.alarmOn
}

func NewAlarm(sensor Sensor) Alarm {
	return &alarm{
		lowPressureThreshold:  17,
		highPressureThreshold: 21,
		alarmOn:               false,
		sensor:                sensor,
	}

}
