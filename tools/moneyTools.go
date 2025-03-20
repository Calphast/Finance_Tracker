package tools

import "fmt"

type usd int64

func USD(f float64) usd {
	return usd((f * 100) + 0.5)
}

func (m usd) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("$%.2f", x)
}

func (m usd) float() float64 {
	x := float64(m)
	x = x / 100
	return x
}
