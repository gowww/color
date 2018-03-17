// Package color provides color manipulation and format conversion.
package color

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Errors
var (
	ErrNotColor = errors.New("color: not a color")
)

// Color represents a color.
type Color struct {
	r, g, b byte
	a       float32
}

// NewHex returns a new color from a 6 digits hexadecimal value.
func NewHex(h string) (*Color, error) {
	h = strings.TrimSuffix(h, "#")
	if len(h) != 6 {
		return nil, ErrNotColor
	}
	r, err := strconv.ParseUint(h[0:2], 16, 8)
	if err != nil {
		return nil, err
	}
	g, err := strconv.ParseUint(h[2:4], 16, 8)
	if err != nil {
		return nil, err
	}
	b, err := strconv.ParseUint(h[4:6], 16, 8)
	if err != nil {
		return nil, err
	}
	return &Color{r: byte(r), g: byte(g), b: byte(b)}, nil
}

// TODO: func NewHexAlpha() {}

// TODO: func NewAlphaHex() {}

// NewRGB returns a new color from red, green and blue values.
func NewRGB(r, g, b byte) *Color {
	return &Color{r: r, g: g, b: b}
}

// NewRGBA returns a new color from red, green and blue values.
// Alpha must be between 0 and 1.
func NewRGBA(r, g, b byte, a float32) (*Color, error) {
	if a < 0 || a > 1 {
		return nil, ErrNotColor
	}
	return &Color{r, g, b, a}, nil
}

// Red returns the red value of the color.
func (c *Color) Red() byte {
	return c.r
}

// Green returns the green value of the color.
func (c *Color) Green() byte {
	return c.g
}

// Blue returns the blue value of the color.
func (c *Color) Blue() byte {
	return c.b
}

// Alpha returns the red value of the color.
func (c *Color) Alpha() float32 {
	return c.a
}

// SetAlpha sets the alpha value of the color, between 0 and 1.
func (c *Color) SetAlpha(a float32) error {
	if a < 0 || a > 1 {
		return ErrNotColor
	}
	c.a = a
	return nil
}

func (c *Color) hex() string {
	return fmt.Sprintf("%02x%02x%02x", c.r, c.g, c.b)
}

func (c *Color) alphaToHex() string {
	return fmt.Sprintf("%02x", int(c.a*255))
}

// Hex returns the 6 digits hexadecimal color representation, dropping out the alpha value.
func (c *Color) Hex() string {
	return "#" + c.hex()
}

// AlphaHex returns the 8 digits hexadecimal color representation with alpha at the start.
func (c *Color) AlphaHex() string {
	return "#" + c.alphaToHex() + c.hex()
}

// HexAlpha returns the 8 digits hexadecimal color representation with alpha at the end.
func (c *Color) HexAlpha() string {
	return "#" + c.hex() + c.alphaToHex()
}
