package main

import (
	"math/rand"
	"time"
)

func (p *Pokemon) Caught() bool {
	minProbability := 5
	maxProbability := 95
	maxBaseExperience := 400
	normalizedExperience := float64(p.BaseExperience) / float64(maxBaseExperience)
	catchProbability := float64(maxProbability) - float64(normalizedExperience*float64(maxProbability-minProbability))

	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	catchRate := randomGenerator.Intn(100) + 1

	return float64(catchRate) <= catchProbability
}
