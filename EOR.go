package main

import (
	"fmt"
	"sort"
)

func main() {
	// Just initilizing the variables

	var api, viscosity, oil_sat, perm, temp, depth, m1, m2, m3, m4, m5, m6, m7, m8 float64
	var formation, thickness, composition int
	m1, m2, m3, m4, m5, m6, m7, m8 = 0, 0, 0, 0, 0, 0, 0, 0

	// Getting the input from the user

	fmt.Print("API: ")
	fmt.Scan(&api)
	fmt.Print("Viscosity(cp): ")
	fmt.Scan(&viscosity)
	fmt.Print("Oil Saturation: ")
	fmt.Scan(&oil_sat)
	fmt.Print("Permeability(mD): ")
	fmt.Scan(&perm)
	fmt.Print("Temperature(F): ")
	fmt.Scan(&temp)
	fmt.Print("Depth(ft): ")
	fmt.Scan(&depth)
	fmt.Print("Choose the formation type: (1)Sandstone (2)Carbonate (3)High porosity Sandstone: ")
	fmt.Scan(&formation)
	fmt.Print("Choose theQ thickness: (1) < 20 ft (2) > 20 ft no dip (3) > 20 ft with dip: ")
	fmt.Scan(&thickness)
	fmt.Print("Choose the composition: (1) High %C1-C7 (2) High %C2-C7 (3) High %C5-C12: ")
	fmt.Scan(&composition)
	fmt.Print("\n")
	fmt.Print("\n")

	// Calculating the percentage of screening for each method

	// Method 1, Nitrogen and Flue gas

	if api > 35 {
		if api > 48 {
			m1 = 1
		} else {
			m1 = 0.5
		}
	}

	if viscosity < 0.4 {
		if viscosity < 0.2 {
			m1 = m1 + 1
		} else {
			m1 = m1 + 0.5
		}
	}

	if composition == 1 {
		m1 = m1 + 1
	}

	if composition == 2 {
		m1 = m1 + 0.5
	}

	if oil_sat > 0.4 {
		if oil_sat > 0.75 {
			m1 = m1 + 1
		} else {
			m1 = m1 + 0.5
		}
	}

	if thickness == 1 || thickness == 3 {
		m1 = m1 + 1
	}

	if depth > 6000 {
		m1 = m1 + 1
	}

	m1 = m1 / 6

	// Method 2, Hydrocarbon

	if api > 23 {
		if api > 41 {
			m2 = 1
		} else {
			m2 = 0.5
		}
	}

	if viscosity < 3 {
		if viscosity < 0.5 {
			m2 = m2 + 1
		} else {
			m2 = m2 + 0.5
		}
	}

	if composition == 2 {
		m2 = m2 + 1
	}
	if composition == 1 {
		m2 = m2 + 0.5
	}

	if oil_sat > 0.3 {
		if oil_sat > 0.8 {
			m2 = m2 + 1
		} else {
			m2 = m2 + 0.5
		}
	}

	if thickness == 1 || thickness == 3 {
		m2 = m2 + 1
	}

	if depth > 4000 {
		m2 = m2 + 1
	}

	m2 = m2 / 6

	// Method 3, CO2

	if api > 22 {
		if api > 36 {
			m3 = 1
		} else {
			m3 = 0.5
		}
	}

	if viscosity < 10 {
		if viscosity < 1.5 {
			m3 = m3 + 1
		} else {
			m3 = m3 + 0.5
		}
	}

	if composition == 3 {
		m3 = m3 + 1
	}

	if oil_sat > 0.2 {
		if oil_sat > 0.55 {
			m3 = m3 + 1
		} else {
			m3 = m3 + 0.5
		}
	}

	if depth > 2500 {
		m3 = m3 + 1
	}

	m3 = m3 / 6

	// Method 4, immiscible gasses

	if api > 12 {
		m4 = 1
	}

	if viscosity < 600 {
		m4 = m4 + 1
	}

	if oil_sat > 0.35 {
		if oil_sat > 0.7 {
			m4 = m4 + 1
		} else {
			m4 = m4 + 0.5
		}
	}

	if depth > 1800 {
		m4 = m4 + 1
	}

	if thickness == 1 || thickness == 3 {
		m4 = m4 + 1
	}

	m4 = m4 / 5

	// Method 5, Micellar / Polymer, ASP, and Alkaline Flooding

	if api > 20 {
		if api > 35 {
			m5 = 1
		} else {
			m5 = 0.5
		}
	}

	if viscosity < 35 {
		if viscosity < 13 {
			m5 = m5 + 1
		} else {
			m5 = m5 + 0.5
		}
	}

	if composition < 3 {
		m5 = m5 + 1
	}
	if oil_sat > 0.35 {
		if oil_sat > 0.53 {
			m5 = m5 + 1
		} else {
			m5 = m5 + 0.5
		}
	}

	if formation == 1 || formation == 3 {
		m5 = m5 + 1
	}

	if perm > 10 {
		if perm > 450 {
			m5 = m5 + 1
		} else {
			m5 = m5 + 0.5
		}
	}

	if depth < 9000 {
		if depth < 3500 {
			m5 = m5 + 1
		} else {
			m5 = m5 + 0.5
		}
	}

	if temp > 200 {
		m5 = m5 + 1
	}

	m5 = m5 / 8

	// Method 6, Polymer Flooding

	if api > 15 || api < 40 {
		m6 = 1
	}

	if viscosity < 150 || viscosity > 10 {
		m6 = m6 + 1
	}

	if oil_sat > 0.7 {
		if oil_sat > 0.8 {
			m6 = m6 + 1
		} else {
			m6 = m6 + 0.5
		}
	}

	if formation == 1 || formation == 3 {
		m6 = m6 + 1
	}

	if perm > 10 {
		if perm > 800 {
			m6 = m6 + 1
		} else {
			m6 = m6 + 0.5
		}
	}

	if depth < 9000 {
		m6 = m6 + 1
	}

	if temp > 200 {
		m6 = m6 + 1
	}

	m6 = m6 / 7

	// Method 7, Combustion

	if api > 10 {
		if api > 16 {
			m7 = 1
		} else {
			m7 = 0.5
		}
	}

	if viscosity < 5000 {
		if viscosity < 1200 {
			m7 = m7 + 1
		} else {
			m7 = m7 + 0.5
		}
	}

	if oil_sat > 0.5 {
		if oil_sat > 0.72 {
			m7 = m7 + 1
		} else {
			m7 = m7 + 0.5
		}
	}

	if formation == 3 {
		m7 = m7 + 1
	}

	if perm > 50 {
		m7 = m7 + 1
	}

	if depth < 11500 {
		if depth < 3500 {
			m7 = m7 + 1
		} else {
			m7 = m7 + 0.5
		}
	}
	if temp > 100 {
		m7 = m7 + 1
	}

	m7 = m7 / 7

	// Method 8, Steam

	if api > 8 {
		if api > 13.5 {
			m8 = 1
		} else {
			m8 = 0.5
		}
	}
	if viscosity < 200000 {
		if viscosity < 4700 {
			m8 = m8 + 1
		} else {
			m8 = m8 + 0.5
		}
	}

	if oil_sat > 0.4 {
		if oil_sat > 0.66 {
			m8 = m8 + 1
		} else {
			m8 = m8 + 0.5
		}
	}

	if formation == 3 {
		m8 = m8 + 1
	}

	if perm > 200 {
		m8 = m8 + 1
	}

	if depth < 4500 {
		m8 = m8 + 1
	}

	m8 = m8 / 6

	// Sorting the results
	EOR := []struct {
		name  string
		score float64
	}{
		{"Nitrogen and Flue gas", m1},
		{"Hydrocarbon", m2},
		{"CO2", m3},
		{"Immiscible gasses", m4},
		{"Micellar / Polymer, ASP, and Alkaline Flooding", m5},
		{"Polymer Flooding", m6},
		{"Combustion", m7},
		{"Steam", m8},
	}

	sort.SliceStable(EOR, func(i, j int) bool {
		return EOR[i].score > EOR[j].score
	})

	// Printing the results
	for _, method := range EOR {
		fmt.Printf("%s: %f\n", method.name, method.score)
	}
}
