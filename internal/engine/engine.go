package engine

import (
	"time"
)

type Engine struct {
	TargetText []rune
	InputText  []rune
	StartTime  time.Time
	EndTime    time.Time
	Active     bool
	Finished   bool
}

func NewEngine(text string) *Engine {
	tr := []rune(text)
	return &Engine{
		TargetText: tr,
		InputText:  make([]rune, 0, len(tr)),
	}
}

func (e *Engine) Start() {
	e.StartTime = time.Now()
	e.Active = true
}

func (e *Engine) Push(r rune) {
	if e.Finished {
		return
	}
	if !e.Active {
		e.Start()
	}
	if len(e.InputText) < len(e.TargetText) {
		e.InputText = append(e.InputText, r)
	}
	if len(e.InputText) == len(e.TargetText) {
		e.EndTime = time.Now()
		e.Active = false
		e.Finished = true
	}
}

func (e *Engine) Pop() {
	if e.Finished || len(e.InputText) == 0 {
		return
	}
	e.InputText = e.InputText[:len(e.InputText)-1]
}

func (e *Engine) WPM() float64 {
	if !e.Active && !e.Finished {
		return 0
	}
	now := time.Now()
	if e.Finished {
		now = e.EndTime
	}
	duration := now.Sub(e.StartTime).Minutes()
	if duration <= 0 {
		return 0
	}
	// WPM = (chars / 5) / minutes
	return float64(len(e.InputText)) / 5.0 / duration
}

func (e *Engine) Accuracy() float64 {
	if len(e.InputText) == 0 {
		return 100
	}
	correct := 0
	for i, r := range e.InputText {
		if r == e.TargetText[i] {
			correct++
		}
	}
	return float64(correct) / float64(len(e.InputText)) * 100
}
