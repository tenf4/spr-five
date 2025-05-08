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

func (t *Training) Parse(datastring string) (err error) {
	separatedData := strings.Split(datastring, ",")
	if len(separatedData) != 3 {
		return errors.New("incorrect data format - amount of paramaters doesn't equal 3")
	}
	t.TrainingType = separatedData[1]
	steps, err := strconv.Atoi(separatedData[0])
	if err != nil {
		return err
	}
	t.Steps = steps

	if steps <= 0 {
		return errors.New("incorrect steps value")
	}
	duration, err := time.ParseDuration(separatedData[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("incorrect steps value")
	}
	t.Duration = duration
	return err
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var spentCalories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	default:
		return "", errors.New("неизвестный вид тренировки")
	}
	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanSpeed,
		spentCalories,
	), err

}
