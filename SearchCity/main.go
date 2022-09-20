package searchcity

import "strings"

var cities = []string{"Paris", "Budapest", "Skopje", "Rotterdam", "Valencia", "Vancouver", "Amsterdam", "Vienna", "Sydney", "New York City", "London", "Bangkok", "Hong Kong", "Dubai", "Rome", "Istanbul"}

func SearchCity(input string) (result []string) {
	if len(input) < 2 {
		return
	}
	for _, v := range cities {
		if strings.HasPrefix(strings.ToLower(v), strings.ToLower(input)) {
			result = append(result, v)
		}
	}
	return
}
