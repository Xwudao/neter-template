package cron

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Cron struct {
	logger      *zap.SugaredLogger
	cron        *cron.Cron
	minuteFunc  map[int][]func()
	secondsFunc map[int][]func()
	hourFunc    map[int][]func()

	//Run once a year, midnight, Jan. 1st
	yearlyFunc []func()
	//Run once a month, midnight, first of month
	monthlyFunc []func()
	//Run once a week, midnight between Sat/Sun
	weeklyFunc []func()
	//Run once a day, midnight
	dailyFunc []func()
	//Run once an hour, beginning of hour
	hourlyFunc []func()
}

func (c *Cron) Close() error {
	c.logger.Infof("Closing cron")
	c.cron.Stop()
	return nil
}

func (c *Cron) AddSecondsFunc(s int, f ...func()) {
	c.secondsFunc[s] = append(c.secondsFunc[s], f...)
}
func (c *Cron) AddMinuteFunc(m int, f ...func()) {
	c.minuteFunc[m] = append(c.minuteFunc[m], f...)
}
func (c *Cron) AddHourly(f func()) {
	c.hourlyFunc = append(c.hourlyFunc, f)
}
func (c *Cron) AddDaily(f func()) {
	c.dailyFunc = append(c.dailyFunc, f)
}
func (c *Cron) AddMonthly(f func()) {
	c.monthlyFunc = append(c.monthlyFunc, f)
}
func (c *Cron) AddYearly(f func()) {
	c.yearlyFunc = append(c.yearlyFunc, f)
}
func (c *Cron) Run() {
	c.logger.Infof("Starting cron")
	c.cron.Start()
}
func NewCron(logger *zap.SugaredLogger) *Cron {
	c := &Cron{
		logger:      logger,
		cron:        cron.New(),
		minuteFunc:  make(map[int][]func()),
		secondsFunc: make(map[int][]func()),
		hourFunc:    make(map[int][]func()),
	}

	return c
}
