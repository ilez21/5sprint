package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // Количество метров в километре.
	minInH                     = 60   // Количество минут в часе.
	stepLengthCoefficient      = 0.45 // Коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // Коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("время должно быть положительным")
	}
	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / 60
	resCalories := walkingCaloriesCoefficient * calories
	return resCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("время должно быть положительным")
	}

	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	hour := duration.Hours()
	averageSpeed := distance / hour
	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	step := stepLengthCoefficient * height
	distances := (float64(steps) * step) / mInKm
	return distances
}
