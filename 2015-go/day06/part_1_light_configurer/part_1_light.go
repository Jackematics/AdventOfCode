package part_1_light_configurer

type Light struct {
	on bool
}

func NewLight() *Light {
	return &Light{on: false}
}

func (l *Light) isOn() bool {
	return l.on
}

func (l *Light) Toggle() {
	l.on = !l.on
}

func (l *Light) TurnOn() {
	l.on = true
}

func (l *Light) TurnOff() {
	l.on = false
}
