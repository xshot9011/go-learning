package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	// TODO Move init value to yaml file -> phrase
	// TODO Logging level enable
	// TODO Move calculation process to function
	const InitialDiscount float64 = 60000.0

	salary := 45000.0
	numberOfMouth := 11

	totalIncome := salary * float64(numberOfMouth)
	personalDiscount := getPersonalDiscount(totalIncome)
	netIncome := totalIncome - (personalDiscount + InitialDiscount)

	fmt.Printf("Total income is: %v\n", totalIncome)
	fmt.Printf("Net income is %v - (%v + %v) = %v\n", totalIncome, personalDiscount, InitialDiscount, netIncome)
	total_tax := getTotalTax(netIncome)
	fmt.Printf("Total tax is %v (%.2v%% of your income)\n", total_tax, total_tax/totalIncome)
	fmt.Printf("%v", trace())
}

func getPersonalDiscount(income float64) float64 {
	DiscountRatio := 0.5
	MaxDiscount := 100000.0

	totalDiscount := DiscountRatio * income

	if totalDiscount > MaxDiscount {
		return MaxDiscount
	}
	return totalDiscount
}

func getTotalTax(netIncome float64) float64 {
	// TODO Rewrite this Function
	Ratio := []float64{0.05, 0.1, 0.15, 0.20, 0.25, 0.3, 0.35}
	Step := []float64{150000, 300000, 500000, 750000, 1000000, 2000000, 5000000}

	tax := 0.0
	taxStep := 0.0

	if netIncome <= float64(Step[0]) {
		return 0
	}
	fmt.Println(" ============ STEP INFO ============")
	for index := 0; index < len(Ratio)-1; index++ {
		if netIncome < Step[index] {
			break
		}
		taxStep = math.Min((netIncome-Step[index])*Ratio[index], (Step[index+1]-Step[index])*Ratio[index])
		tax += taxStep
		fmt.Printf("STEP %2v%% (%7v-%7v) cost: %6v\n", Ratio[index]*100, Step[index], Step[index+1], taxStep)
	}

	if netIncome <= float64(Step[len(Step)-1]) {
		return tax
	}
	taxStep = (netIncome - Step[len(Step)-1]) * Ratio[len(Ratio)-1]
	fmt.Printf("STEP %2v%% (%7v-%7v) cost: %6v\n", Ratio[len(Ratio)-1]*100, Step[len(Step)-2], Step[len(Step)-1], taxStep)

	fmt.Println("Function")
	return tax + taxStep
}

func trace() string {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s:%d %s\n", file, line, f.Name())
}
