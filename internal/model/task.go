package model

import (
	"time"
)

type Task struct {
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	StartedAt   time.Time     `json:"started_at"`
	FinishedAt  time.Time     `json:"finished_at"`
	IsFinished  bool          `json:"is_finished"`
	UserId      int           `json:"user_id"`
	Duration    time.Duration `json:"-"`
	DurationStr string        `json:"duration"`
}

func (t *Task) GetDuration() time.Duration {
	dur := t.FinishedAt.Sub(t.StartedAt)
	return dur
}

type TaskPeriod struct {
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	UserId int       `json:"user_id"`
}

func (p *TaskPeriod) SetTime(fromStr string, toStr string) error {
	from, err := time.Parse(time.DateOnly, fromStr)
	if err != nil {
		return err
	}
	to, err := time.Parse(time.DateOnly, toStr)
	if err != nil {
		return err
	}
	p.From = from
	p.To = to

	return nil
}
