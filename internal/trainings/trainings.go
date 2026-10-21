package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New(
			"ожидаются три параметра: шаги, тип тренировки и продолжительность",
		)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось разобрать количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("не удалось разобрать продолжительность: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	), nil
}
