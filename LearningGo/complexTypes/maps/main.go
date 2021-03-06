package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println("states", states)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println("states", states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println("states", states)

	states["NY"] = "New York"
	fmt.Println("states", states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	//Sorting is not given in maps so it could take any order
	sort.Strings(keys)
	for j := range keys {
		fmt.Printf("%v: %v\n", keys[j], states[keys[j]])
	}
}
