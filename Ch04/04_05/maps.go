package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)

	fmt.Println(states)

	states["WA"] = "Washinton"
	states["OR"] = "Oregon"
	states["CA"] = "Califonia"
	fmt.Println(states)
	fmt.Println(states["CA"])

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	fmt.Println("\nSorted")
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
