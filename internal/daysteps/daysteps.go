package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("ожидаются два параметра: шаги и продолжительность")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось разобрать количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("не удалось разобрать продолжительность: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	), nil
}
