package Seeder

import (
	budaya "github.com/wkloose/Buddya/Seeder/Budaya"
	"github.com/wkloose/Buddya/Seeder/prompt"
)

func SeedAll() {
	budaya.SeedBudayaBatch1()
	budaya.SeedBudayaBatch2()
	budaya.SeedBudayaBatch3()
	prompt.SeedPrompt()
	SeedPlaces()
}
