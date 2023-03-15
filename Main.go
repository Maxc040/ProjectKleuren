package main

import (
	"fmt"
	"os"
)

func main() {
	// Maak een map aan waarin de teksten opgeslagen worden voor de bijbehorende kleuren
	texts := map[string]string{
		"blauw": "Blauw zoals de lucht.",
		"rood":  "Rood met passie.",
		"geel":  "Geel als de stralen van de zon.",
		"groen": "Groen van de natuur.",
		"paars": "Zo paars als een aubergine.",
		"zwart": "Zwart als de gemiddelde PSV supporter.",
		"roze":  "Roze als de vriendenboekje van mijn zusje.",
	}

	// Lees de opgegeven kleur van de gebruiker
	fmt.Println("Geef je lievelingskleur op:")
	var color string
	_, err := fmt.Scanln(&color)
	if err != nil {
		fmt.Println("Fout bij het lezen van de kleur:", err)
		os.Exit(1) // Sluit de applicatie af met een exit code 1 bij een fout
	}

	// Zoek de tekst op voor de opgegeven kleur
	text, found := texts[color]
	if !found {
		fmt.Println("De opgegeven kleur is niet bekend.")
		os.Exit(1) // Sluit de applicatie af met een exit code 1 bij een fout
	}

	// Schrijf de tekst naar een bestand
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Fout bij het aanmaken van output bestand:", err)
		os.Exit(1) // Sluit de applicatie af met een exit code 1 bij een fout
	}
	defer file.Close() // Sluit het bestand altijd aan het einde van de functie
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Fout bij het schrijven naar output bestand:", err)
		os.Exit(1) // Sluit de applicatie af met een exit code 1 bij een fout
	}

	fmt.Println("De tekst is opgeslagen in output.txt")
}
