package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	Personal personaldata.Personal
}

func (ds DaySteps) Weight() float64 {
	return ds.Personal.Weight
}

func (ds DaySteps) Height() float64 {
	return ds.Personal.Height
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return errors.New("количество элементов не равна 2")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return errors.New("ошибка преобразования")
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	times, err := time.ParseDuration(slice[1])
	if err != nil {
		return errors.New("время должно быть положительным")
	}
	if times <= 0 {
		return errors.New("время должно быть положительным")
	}
	ds.Steps = steps
	ds.Duration = times
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height())
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight(), ds.Height(), ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %v", err)
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}

func (ds DaySteps) Print() {
	ds.Personal.Print()
}
