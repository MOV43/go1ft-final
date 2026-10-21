package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчётов.
const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

func WalkingSpentCalories(
	steps int,
	weight, height float64,
	duration time.Duration,
) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(
	steps int,
	weight, height float64,
	duration time.Duration,
) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше нуля")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше нуля")
	}

	speed := MeanSpeed(steps, height, duration)
	calories := weight * speed * duration.Minutes() / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
