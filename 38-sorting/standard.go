package main

import (
	"fmt"
	"sort"
)

// Simple Slices
func ascStdSort(unsorted []int) {
	sort.Slice(unsorted, func(i, j int) bool {
		return unsorted[i] < unsorted[j]
	})

	fmt.Println("Ascending order:", unsorted)
}

func desStdSort(unsorted []int) {
	sort.Slice(unsorted, func(i, j int) bool {
		return unsorted[i] > unsorted[j]
	})

	fmt.Println("Descending order:", unsorted)
}

func ascStdIntSort(unsorted []int) {
	// much faster than sort.Slice
	sort.Ints(unsorted)

	fmt.Println("Ascending order:", unsorted)
}

// Binary Search
func binarySearchStd(sorted []int, target int) int {
	// much faster than sort.Search
	return sort.Search(len(sorted), func(i int) bool {
		return sorted[i] >= target
	})
	//return sort.SearchInts(sorted, target)
}

func binaryFindStd(sorted []int, target int) (int, bool) {
	// Comparor must return 0 if target is found
	return sort.Find(len(sorted), func(i int) int {
		return target - sorted[i]
	})
}

// Person Slices of Structs
type Person struct {
	Name string
	Age  int
}

// Method 1
func sortByAge(p []Person) {
	sort.Slice(p, func(i, j int) bool {
		return p[i].Age < p[j].Age
	})
	fmt.Println("sorted by age people: ", p)
}

// ByAge Implement the sort.Interface interface (Method 2)
type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func ascStdSortPeople(p []Person) {
	sort.Sort(ByAge(p))

	fmt.Println("Ascending order:", p)
}

// Sort Keys
type earthMass float64
type au float64

type Planet struct {
	name     string
	mass     earthMass
	distance au
}

// By is the type of "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *Planet) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *planetSorter) Len() int {
	return len(s.planets)
}

// Swap is part of sort.Interface.
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

func sortPlanetsExample() {
	var planets = []Planet{
		{"Mercury", 0.055, 0.4},
		{"Venus", 0.815, 0.7},
		{"Earth", 1.0, 1.0},
		{"Mars", 0.107, 1.5},
	}
	// Closures that order the Planet structure.
	name := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}
	mass := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}
	distance := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}
	decreasingDistance := func(p1, p2 *Planet) bool {
		return distance(p2, p1)
	}

	// Sort the planets by the various criteria.
	By(name).Sort(planets)
	fmt.Println("By name:", planets)

	By(mass).Sort(planets)
	fmt.Println("By mass:", planets)

	By(distance).Sort(planets)
	fmt.Println("By distance:", planets)

	By(decreasingDistance).Sort(planets)
	fmt.Println("By decreasing distance:", planets)
}
