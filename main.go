package main

import (
	"fmt"
	"go-functions/simplemath"
	"net/http"
	"strings"
)

func main() {
	fmt.Printf("p1 + p2: %.2f\n", simplemath.Add(3, 5))
	fmt.Printf("p1 - p2: %.2f\n", simplemath.Subtract(3, 5))
	fmt.Printf("p1 * p2: %.2f\n", simplemath.Multiply(3, 5))
	division, err := simplemath.Divide(3, 5)
	if err == nil {
		fmt.Printf("p1 / p2: %.2f\n", division)
	} else {
		fmt.Printf("error dividing p1 by p2: %s", err.Error())
	}

	division, err = simplemath.Divide(3, 0)
	if err == nil {
		fmt.Printf("p1 / p2: %.2f\n", division)
	} else {
		fmt.Printf("error dividing p1 by p2: %s\n", err.Error())
	}

	fmt.Printf("Total sum of 1,2,3,4: %.2f\n", simplemath.Sum(1, 2, 3, 4))

	sv := simplemath.NewSemanticVersion(1, 0, 1)
	fmt.Printf("Semantic version: %s\n", sv.String())
	sv.IncrementMajor()
	sv.IncrementMinor()
	sv.IncrementPatch()
	fmt.Printf("Incremented Semantic version: %s\n", sv.String())

	var tripper http.RoundTripper = &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "www.pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)

	// anonymous function
	f := func(name string) {
		fmt.Printf(" My name is: %s\n", name)
	}

	f("Claudio")
	f("Maria")

	// returning a func from a func
	f1 := mathExpression(AddExpr)
	f2 := mathExpression(MultiplyExpr)
	fmt.Printfw("Summing 1 and 5: %.2f\n", f1(1, 5))
	fmt.Printf("Multiplying 3 by 4: %.2f\n", f2(3, 4))
}

type RoundTripCounter struct {
	count int
}

func (counter *RoundTripCounter) RoundTrip(r *http.Request) (*http.Response, error) {
	counter.count += 1
	return nil, nil
}

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func mathExpression(expr MathExpr) func(f1, f2 float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		return func(f1, f2 float64) float64 {
			return 0
		}
	}
}
