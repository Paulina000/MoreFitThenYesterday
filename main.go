package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Userer interface {
	CreateConstantFeatures()
	CraeteChangableFeaturesScale()
	CreateChangableFeaturesKaliper()
}

type ConstantFeatures struct {
	gender string
	age    int
	height int
}

type ChangableFeaturesScale struct {
	weight   float32
	fat      float32
	water    float32
	musseles float32
}

type ChangableFeaturesKaliper struct {
	backFat  float32
	hipFat   float32
	armFat   float32
	musseles float32
}

func findGender() string {

	reader := bufio.NewReader(os.Stdin)
	for {
		var gender string
		fmt.Println("Give me your gender:")
		gender, _ = reader.ReadString('\n')
		gender = strings.TrimSpace(gender)

		if gender == "woman" || gender == "w" {
			return "woman"
		} else if gender == "man" || gender == "m" {
			return "man"
		} else {
			fmt.Println("Choose woman or man")
		}
	}

}

func findAge() float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Ask for a birth year
		fmt.Println("Give me a birth year:")
		year, _ := reader.ReadString('\n')
		year = year[:len(year)-1] // Remove the newline character
		// Check if the input is a valid number
		year = strings.TrimSpace(year)
		if _, err := strconv.Atoi(year); err != nil {
			fmt.Println("That's not a valid number. Please try again.")
			continue
		}
		// Convert year to int for further checks
		yearInt, _ := strconv.Atoi(year)

		// Check if the year is within a reasonable range
		if yearInt < 1925 || yearInt > 2012 {
			fmt.Println("That's not a reasonable date. Please try again.")
			continue
		}

		// Calculate and return the age if the year is valid
		age := 2025 - yearInt
		ageFloat := float64(age)

		return ageFloat
	}
}

func findMeasureFat() float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Ask for a birth year
		fmt.Println("Give me fat measures in mm:")
		measureFatmm, _ := reader.ReadString('\n')
		measureFatmm = measureFatmm[:len(measureFatmm)-1] // Remove the newline character
		// Check if the input is a valid number
		measureFatmm = strings.TrimSpace(measureFatmm)
		if _, err := strconv.ParseFloat(measureFatmm, 64); err != nil {
			fmt.Println("That's not a valid number. Please try again.")
			continue
		}
		// Convert year to int for further checks
		bellyFatmmInt, _ := strconv.ParseFloat(measureFatmm, 64)

		// Check if the year is within a reasonable range
		if bellyFatmmInt < 5 || bellyFatmmInt > 200 {
			fmt.Println("That's not a reasonable measure. Please try again.")
			continue
		}

		// Calculate and return the age if the year is valid
		return bellyFatmmInt
	}
}

func CreateConstantFeatures() {
	gender := findGender()
	fmt.Println("gender: ", gender)
	age := findAge()
	fmt.Println("age: ", age)
	// backFat := findMeasureFat()

	// fmt.Println("backFat: ", backFat)

	// armFat := findMeasureFat()
	// fmt.Println("armFat: ", armFat)

	bellyFat := findMeasureFat()
	fmt.Println("findBellyFat: ", bellyFat)
	bodyFatWoman := bellyFat*1.223 - (0.0174 * bellyFat * bellyFat) + 0.124*(age-6.07)
	bodyFatMan := bellyFat*1.378 - (0.0174 * bellyFat * bellyFat) + 0.213*(age-5.84)
	fmt.Println("bodyFatWoman : ", bodyFatWoman)
	fmt.Println("bodyFatMan : ", bodyFatMan)
}
func main() {
	CreateConstantFeatures()
}
