// Package biorhythm computes three pseudo-scientific cycle values from a birth date.
package biorhythm

import (
	"math"
	"time"
)

const (
	// PhysicalCycle is the length of the physical cycle in days.
	PhysicalCycle = 23
	// EmotionalCycle is the length of the emotional cycle in days.
	EmotionalCycle = 28
	// IntellectualCycle is the length of the intellectual cycle in days.
	IntellectualCycle = 33
)

// Biorhythm stores the birth date and the reference date used for calculations.
type Biorhythm struct {
	days float64
}

// New creates a biorhythm calculator for a birth date and optional reference time.
// When no reference time is provided, the current local time is used.
func New(birthDate time.Time, now time.Time) *Biorhythm {
	days := now.Sub(birthDate).Hours() / 24

	return &Biorhythm{
		days: days,
	}
}

// Physical returns the physical cycle value for the configured date.
func (b Biorhythm) Physical() float64 {
	return b.value(PhysicalCycle)
}

// Emotional returns the emotional cycle value for the configured date.
func (b Biorhythm) Emotional() float64 {
	return b.value(EmotionalCycle)
}

// Intellectual returns the intellectual cycle value for the configured date.
func (b Biorhythm) Intellectual() float64 {
	return b.value(IntellectualCycle)
}

// Physical calculates the physical cycle value for a birth date using the given time.
func Physical(birtDate time.Time, now time.Time) float64 {
	return New(birtDate, now).Physical()
}

// Emotional calculates the emotional cycle value for a birth date using the given time.
func Emotional(birtDate time.Time, now time.Time) float64 {
	return New(birtDate, now).Emotional()
}

// Intellectual calculates the intellectual cycle value for a birth date using the given time.
func Intellectual(birtDate time.Time, now time.Time) float64 {
	return New(birtDate, now).Intellectual()
}

func (b Biorhythm) value(period float64) float64 {
	if b.days == 0 {
		return 0
	}

	return math.Sin(2 * math.Pi * b.days / period)
}
