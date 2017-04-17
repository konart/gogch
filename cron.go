package main

import "fmt"

type Cron struct {
	questions  []string
	available  map[int][]string
	enabled    map[int]string
	done       bool
	questionId int
}

func (c *Cron) Execute() {

}

func (c *Cron) set(q, x int) {
	c.enabled[q] = c.available[q][x]
}

func (c *Cron) ListOptions() {
	for i, cron := range c.available[c.questionId] {
		fmt.Printf("%d: %s\n", i, cron)
	}
}

func (c *Cron) RequestAddInfo() {
	fmt.Println(c.questions[c.questionId])
}

func (c *Cron) ReadAnswer() {
	var x int
	fmt.Scanln(&x)
	c.set(c.questionId, x)
	c.questionId += 1
	if c.questionId == len(c.questions) {
		c.Done()
	}
}

func (c *Cron) Init() {
	c.questions = []string{"What type of Cron will you need?"}
	c.available = map[int][]string{0: {"CronModule15", "CronModele", "CronModule3", "CronModule12", "CronModule24"}}
	c.enabled = map[int]string{}
	c.done = false
}

func (c *Cron) isDone() bool {
	return c.done
}

func (c *Cron) Done() {
	c.done = true
}
