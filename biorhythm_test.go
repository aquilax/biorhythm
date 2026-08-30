package biorhythm

import (
	"math"
	"testing"
	"time"
)

func almostEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func TestBiorhythmValues(t *testing.T) {
	birthDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Date(2000, 1, 6, 12, 0, 0, 0, time.UTC)
	b := New(birthDate, now)

	days := now.Sub(birthDate).Hours() / 24

	expectedPhysical := math.Sin(2 * math.Pi * days / 23)
	expectedEmotional := math.Sin(2 * math.Pi * days / 28)
	expectedIntellectual := math.Sin(2 * math.Pi * days / 33)

	if !almostEqual(b.Physical(), expectedPhysical, 1e-12) {
		t.Fatalf("physical mismatch: got %v want %v", b.Physical(), expectedPhysical)
	}
	if !almostEqual(b.Emotional(), expectedEmotional, 1e-12) {
		t.Fatalf("emotional mismatch: got %v want %v", b.Emotional(), expectedEmotional)
	}
	if !almostEqual(b.Intellectual(), expectedIntellectual, 1e-12) {
		t.Fatalf("intellectual mismatch: got %v want %v", b.Intellectual(), expectedIntellectual)
	}
}

func TestConvenienceFunctions(t *testing.T) {
	birthDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Date(2000, 1, 7, 0, 0, 0, 0, time.UTC)

	if !almostEqual(Physical(birthDate, now), New(birthDate, now).Physical(), 1e-12) {
		t.Fatalf("physical convenience mismatch")
	}
	if !almostEqual(Emotional(birthDate, now), New(birthDate, now).Emotional(), 1e-12) {
		t.Fatalf("emotional convenience mismatch")
	}
	if !almostEqual(Intellectual(birthDate, now), New(birthDate, now).Intellectual(), 1e-12) {
		t.Fatalf("intellectual convenience mismatch")
	}
}

func TestNowIsBeforeBirthDate(t *testing.T) {
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	now := birthDate.Add(-time.Hour * 24) // One day before birth date

	b := New(birthDate, now)
	expectedPhysical := -0.269796771157
	if !almostEqual(b.Physical(), expectedPhysical, 1e-12) {
		t.Fatalf("physical mismatch: got %v want %v", b.Physical(), expectedPhysical)
	}
}
