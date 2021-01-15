package main

import "fmt"

func main() {
	// map[key type]value type
	m := map[string]int{
		"Jason": 25,
		"Fan":   31,
	}
	fmt.Println(m)
	fmt.Println(m["Jason"])

	fmt.Println("==================")

	fmt.Println(m["Wendy"])
	v, ok := m["Wendy"]
	fmt.Println(v, ok)

	fmt.Println("==================")

	name := "Wendy"
	if v, ok := m[name]; ok {
		fmt.Println("key exist!")
		fmt.Println(v)
	} else {
		fmt.Println("key not exist!")
		m[name] = 999
		fmt.Println("add key")
	}

	fmt.Println("==================")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("==================")

	if v, ok := m[name]; ok {
		fmt.Printf("Delete %v: %v\n", name, v)
		delete(m, name)
	}

	fmt.Println("==================")

	for k, v := range m {
		fmt.Println(k, v)
	}
}
