package main

import (
	"fmt"
	"math"
	"strings"
	// "strings"
)

type Person struct {
	first_name string
	last_name  string
	address    string
}

// func hash(p Person) string {
// 	combine_names := p.first_name + p.last_name + p.address
// 	no_spaces_cn := strings.ReplaceAll(combine_names, " ", "")
// 	lower_case_cn := strings.ToLower(no_spaces_cn)
// 	return lower_case_cn
// }

func sumString(s string) int {
	sum := 0
	for _, char := range s {
		sum += int(char)
	}
	return sum
}

func hash(p Person, a map[interface{}]interface{}) float64 {
	first := strings.ToLower(string(p.first_name[0]))
	i := a[first].(int)
	total := 0
	total += sumString(p.first_name)
	total += sumString(p.last_name)
	total += sumString(p.address)
	total += total * i
	return math.Mod(float64(total), 100)
}

func main() {
	alphabet := map[interface{}]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
	fmt.Println(alphabet)

	personA := Person{"Trevor", "Gevers", "114 Dorothy Street"}
	personB := Person{"Jessica", "Gevers", "114 Dorothy Street"}
	personC := Person{"Adalynn", "Gevers", "114 Dorothy Street"}
	personD := Person{"Julia", "Gevers", "114 Dorothy Street"}
	personE := Person{"Hudson", "Gevers", "114 Dorothy Street"}
	// names := []Person{personA, personB}
	// fmt.Println(names)
	// for _, person := range names {
	// 	print(fmt.Printf("%v", person))
	// }
	fmt.Println("hash for ", personA.first_name, hash(personA, alphabet))
	fmt.Println("hash for ", personB.first_name, hash(personB, alphabet))
	fmt.Println("hash for ", personC.first_name, hash(personC, alphabet))
	fmt.Println("hash for ", personD.first_name, hash(personD, alphabet))
	fmt.Println("hash for ", personE.first_name, hash(personE, alphabet))
	// var a = 5
	// fmt.Printf("var a: %v", &a)
	fmt.Println("")
}
