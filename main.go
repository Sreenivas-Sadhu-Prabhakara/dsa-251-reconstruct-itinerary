package main

import (
	"fmt"
	"sort"
)

// FindItinerary orders all tickets into one continuous trip from "JFK",
// choosing the lexically smallest itinerary. Hierholzer's algorithm for an
// Eulerian path: walk greedily consuming edges, and append each airport to
// the route only when it has no unused departures left (then reverse).
func FindItinerary(tickets [][2]string) []string {
	adj := map[string][]string{}
	for _, t := range tickets {
		adj[t[0]] = append(adj[t[0]], t[1])
	}
	for from := range adj {
		sort.Strings(adj[from]) // smallest destination consumed first
	}
	used := map[string]int{} // cursor into each airport's sorted departures
	route := []string{}
	var visit func(airport string)
	visit = func(airport string) {
		for used[airport] < len(adj[airport]) {
			next := adj[airport][used[airport]]
			used[airport]++
			visit(next)
		}
		route = append(route, airport) // dead end or exhausted: finalize
	}
	visit("JFK")
	for l, r := 0, len(route)-1; l < r; l, r = l+1, r-1 {
		route[l], route[r] = route[r], route[l]
	}
	return route
}

func main() {
	fmt.Println(FindItinerary([][2]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
	fmt.Println(FindItinerary([][2]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
}
