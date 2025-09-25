package Seeder

import (
	budaya "github.com/wkloose/Buddya/Seeder/Budaya"
	"github.com/wkloose/Buddya/Seeder/Prompt"
)

func SeedAll() {
	budaya.SeedBudayaBatch1()
	budaya.SeedBudayaBatch2()
	budaya.SeedBudayaBatch3()
	Prompt.SeedPrompt()
	SeedPlaces()
}
