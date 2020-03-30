package shapes

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	length float64
}

func (r *Rectangle) SetLength(length float64) error {
	if length < 0 {
		return fmt.Errorf("invalid length: %f", length)
	}
	r.length = length
	return nil
}

func (r *Rectangle) Length() float64 {
	return r.length
}

func (r *Rectangle) SetWidth(width float64) error {
	if width < 0 {
		return fmt.Errorf("invalid width: %f", width)
	}
	r.width = width
	return nil
}

func (r *Rectangle) Width() float64 {
	return r.width
}
