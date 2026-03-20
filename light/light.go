package light

const (
	vendorID = 0x046d
)

type Litra interface {
	TurnOn() error
	TurnOff() error

	BrightnessSet(amount int) error
	BrightnessIncrease(amount int) error
	BrightnessDecrease(amount int) error

	TemperatureSet(temp int) error
	TemperatureIncrease(amount int) error
	TemperatureDecrease(amount int) error

	Close() error
}
