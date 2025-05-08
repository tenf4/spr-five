package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("incorrect steps value\n")
	}
	if weight <= 0 {
		return 0, errors.New("incorrect weight value\n")
	}
	if height <= 0 {
		return 0, errors.New("incorrect height value\n")
	}
	mean := MeanSpeed(steps, height, duration)
	durationMins := duration.Minutes()
	spentCalories := (weight * mean * durationMins) / minInH
	return spentCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("incorrect steps value\n")
	}
	if weight <= 0 {
		return 0, errors.New("incorrect weight value\n")
	}
	if height <= 0 {
		return 0, errors.New("incorrect height value\n")
	}
	if duration <= 0 {
		return 0, errors.New("zero or negative duration\n")
	}
	mean := MeanSpeed(steps, height, duration)
	durationMins := duration.Minutes()
	spentCalories := (weight * mean * durationMins) / minInH
	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}
	if height <= 0 {
		return 0
	}
	distanceKm := Distance(steps, height)
	mean := distanceKm / duration.Hours()
	return mean
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm
	return distanceKm
}
