package trainings

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	// TODO: добавить поля
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return errors.New("количество элементов не равна 3")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return errors.New("ошибка преобразования")
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	times, err := time.ParseDuration(slice[2])
	if err != nil {
		return errors.New("время должно быть положительным")
	}
	if times <= 0 {
		return errors.New("время должно быть положительным")
	}
	t.Duration = times
	t.Steps = steps
	t.TrainingType = slice[1]
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %v", err)
	}

	durationHours := t.Duration.Hours()

	info := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		durationHours,
		distance,
		speed,
		calories,
	)

	return info, nil
}
