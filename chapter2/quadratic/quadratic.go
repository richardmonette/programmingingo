package main

import (
	"net/http"
	"log"
	"fmt"
	"math"
	"strconv"
)

const form = `<form action="/" method="POST">
<label for="numbers">Solves equations in the form of ax&sup2; + bx + c</label><br />
<input type="text" name="a" size="10">
<input type="text" name="b" size="10">
<input type="text" name="c" size="10"><br />
<input type="submit" value="Calculate">
</form>`

const pageTop = "<html>"
const anError = "anError"
const pageBottom = "</html>"

func solve(a float64, b float64, c float64) (float64, float64) {
	
	discriminant := complex64(math.Pow(b, 2) - 4 * a * c)
	x1 := (-b + math.Sqrt(discriminant)) / 2 * a
	x2 := (-b - math.Sqrt(discriminant)) / 2 * a

	return x1, x2
}

func formatSolution(solution1 float64, solution2 float64) string {
	return "here is the solution"
}

func processRequest(request *http.Request) (float64, float64, float64, bool) {
	//var a, b, c float64
	//var aExists, bExists, cExists bool

	aSlice, aExists := request.Form["a"]
	bSlice, bExists := request.Form["b"]
	cSlice, cExists := request.Form["c"]

	a, _ := strconv.ParseFloat(aSlice[0], 64)
	b, _ := strconv.ParseFloat(bSlice[0], 64)
	c, _ := strconv.ParseFloat(cSlice[0], 64)

	/*	text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // no data first time form is shown
	}
	return numbers, "", true*/
	if aExists && bExists && cExists {
		return a, b, c, true
	}

	return a, b, c, false
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if a, b, c, ok := processRequest(request); ok {
			solution1, solution2 := solve(a, b, c)
			fmt.Fprint(writer, formatSolution(solution1, solution2))
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

