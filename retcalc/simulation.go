/*

   simulation.go

   This file is part of retirement_calculator.

   retirement_calculator is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   retirement_calculator is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with retirement_calculator.  If not, see <http://www.gnu.org/licenses/>.
*/

package retcalc

import (
	"math/rand"
	"time"
)

type Sim []float64

func (s Sim) GrowthFactor(startYear int) float64 {
	fac := 1.0
	for i := startYear; i < len(s); i++ {
		fac *= (1 + s[i])
	}
	return fac
}

func (s Sim) GrowthFactorWithTaxes(startYear int, eff_tax float64) float64 {
	fac := 1.0
	for i := startYear; i < len(s); i++ {
		fac *= (1 + s[i]*eff_tax)
	}
	return fac
}

func Simulation(mean, stdev float64, sample_size int) []float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	sim := make([]float64, sample_size, sample_size)
	for i := 0; i < sample_size; i++ {
		sim[i] = rand.NormFloat64()*stdev + mean
	}
	return sim
}
