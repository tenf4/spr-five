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

func (ds *DaySteps) Parse(datastring string) (err error) {
	separatedData := strings.Split(datastring, ",")
	if len(separatedData) != 2 {
		return errors.New("incorrect data format - amount of paramaters doesn't equal 2")
	}
	steps, err := strconv.Atoi(separatedData[0])
	if err != nil {
		return err
	}
	ds.Steps = steps

	if steps <= 0 {
		return errors.New("incorrect steps value")
	}
	duration, err := time.ParseDuration(separatedData[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("incorrect duration value")
	}
	ds.Duration = duration
	return err
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", fmt.Errorf("zero or negative duration of training")
	}
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("calory calculation error: %w", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCalories), err
}
