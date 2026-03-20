package light

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/sstallion/go-hid"
)

const (
	productID = 0xc900

	MinBrightness = 0x14
	MaxBrightness = 0xfa

	MinTemperature = 2700
	MaxTemperature = 6500
)

type litraGlow struct {
	dev *hid.Device

	isOn        bool
	brightness  int
	temperature int
}

func NewLitraGlow() (Litra, error) {
	if err := hid.Init(); err != nil {
		return nil, fmt.Errorf("failed to init HID: %w", err)
	}
	dev, err := hid.OpenFirst(vendorID, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to open Litra Glow (is the device connected?): %w", err)
	}
	return &litraGlow{
		dev:         dev,
		brightness:  20,
		temperature: MinTemperature,
		isOn:        false,
	}, nil
}

func (l *litraGlow) Close() error {
	return l.dev.Close()
}

func (l *litraGlow) write(data []byte) error {
	padding := make([]byte, 20-len(data))
	_, err := l.dev.Write(append(data, padding...))
	return err
}

func (l *litraGlow) TurnOn() error {
	if err := l.write([]byte{0x11, 0xff, 0x04, 0x1c, 0x01}); err != nil {
		return fmt.Errorf("turn on: %w", err)
	}
	l.isOn = true
	return nil
}

func (l *litraGlow) TurnOff() error {
	if err := l.write([]byte{0x11, 0xff, 0x04, 0x1c, 0x00}); err != nil {
		return fmt.Errorf("turn off: %w", err)
	}
	l.isOn = false
	return nil
}

func (l *litraGlow) BrightnessSet(amount int) error {
	if amount < 0 || amount > 100 {
		return fmt.Errorf("brightness must be between 0 and 100, got %d", amount)
	}
	normalized := MinBrightness + int(math.Round(float64(amount)/100.0*float64(MaxBrightness-MinBrightness)))
	if err := l.write([]byte{0x11, 0xff, 0x04, 0x4c, 0x00, byte(normalized)}); err != nil {
		return fmt.Errorf("set brightness: %w", err)
	}
	l.brightness = amount
	return nil
}

func (l *litraGlow) BrightnessIncrease(amount int) error {
	return l.BrightnessSet(l.brightness + amount)
}

func (l *litraGlow) BrightnessDecrease(amount int) error {
	return l.BrightnessSet(l.brightness - amount)
}

func (l *litraGlow) TemperatureSet(amount int) error {
	if amount < MinTemperature || amount > MaxTemperature {
		return fmt.Errorf("temperature must be between %d and %d K, got %d", MinTemperature, MaxTemperature, amount)
	}
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(amount))
	if err := l.write([]byte{0x11, 0xff, 0x04, 0x9c, b[0], b[1]}); err != nil {
		return fmt.Errorf("set temperature: %w", err)
	}
	l.temperature = amount
	return nil
}

func (l *litraGlow) TemperatureIncrease(amount int) error {
	return l.TemperatureSet(l.temperature + amount)
}

func (l *litraGlow) TemperatureDecrease(amount int) error {
	return l.TemperatureSet(l.temperature - amount)
}
