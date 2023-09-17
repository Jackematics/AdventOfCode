package part_2_light_configurer

type Light struct {
	brightness int
}

func NewLight() *Light {
	return &Light{brightness: 0}
}

func (l *Light) GetBrightness() int {
	return l.brightness
}

func (l *Light) Toggle() {
	l.brightness += 2
}

func (l *Light) TurnOn() {
	l.brightness += 1
}

func (l *Light) TurnOff() {
	if l.brightness > 0 {
		l.brightness -= 1
	}
}
