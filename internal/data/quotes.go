package data

import (
	"math/rand"
	"time"
)

var quotes = []string{
	"The only way to go fast, is to go well.",
	"Clean code always looks like it was written by someone who cares.",
	"Talk is cheap. Show me the code.",
	"Programs must be written for people to read, and only incidentally for machines to execute.",
	"Measuring programming progress by lines of code is like measuring aircraft building progress by weight.",
	"The most disastrous thing that you can ever learn is your first programming language.",
	"Any fool can write code that a computer can understand. Good programmers write code that humans can understand.",
	"The computer was born to solve problems that did not exist before.",
	"Optimization is the root of all evil.",
	"Hardware is the part of a computer that you can kick; software is the part that you can only curse.",
}

func GetRandomQuote() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return quotes[r.Intn(len(quotes))]
}
