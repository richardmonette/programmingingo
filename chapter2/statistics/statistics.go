package main

import (
	"sort"
	"net/http"
	"log"
	"fmt"
	"strings"
	"strconv"
	"math"
)

const form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`

const pageTop = "<html>"
const anError = "anError"
const pageBottom = "</html>"

type statistics struct {
	numbers []float64
	mean float64
	median float64
	mode []float64
	stddev float64
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func mode(numbers []float64) []float64 {
	var m map[float64]int
	m = make(map[float64]int)

	for _, x := range numbers {
		m[x]++
	}

	var highest int

	for _, value := range m {
		if value > highest {
			highest = value
		}
	}

	var retVal []float64

	for key, value := range m {
		if value == highest {
			retVal = append(retVal, key)
		}
	}

	return retVal
}

func stddev(numbers []float64) float64 {
	n := float64(len(numbers))
	mean := sum(numbers) / float64(len(numbers))
	var total float64
	for _, x := range numbers {
		total += math.Pow(x - mean, 2)
	}
	return math.Sqrt(total / (n-1))
}

func mean(numbers []float64) float64 {
	return sum(numbers) / float64(len(numbers))
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = mean(numbers)
	stats.median = median(numbers)
	stats.mode = mode(numbers)
	stats.stddev = stddev(numbers)
	return stats
}

func formatStats(stats statistics) string { 
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
<tr><td>Mode</td><td>%v</td></tr>
<tr><td>Std. Dev.</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.mode, stats.stddev)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
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
	return numbers, "", true
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
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
