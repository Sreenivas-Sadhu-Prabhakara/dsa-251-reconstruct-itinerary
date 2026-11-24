package main

import (
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	got := FindItinerary([][2]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}})
	want := []string{"JFK", "MUC", "LHR", "SFO", "SJC"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("itinerary = %v, want %v", got, want)
	}

	// Lexical preference can lead into a dead end that must be deferred:
	// greedy-first JFK->ATL works here, but the algorithm must handle both.
	got = FindItinerary([][2]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}})
	want = []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("itinerary = %v, want %v", got, want)
	}

	// The famous trap: smallest next hop (KUL) is a dead end; NRT must come first.
	got = FindItinerary([][2]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}})
	want = []string{"JFK", "NRT", "JFK", "KUL"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("dead-end case = %v, want %v", got, want)
	}

	// Single ticket.
	got = FindItinerary([][2]string{{"JFK", "AAA"}})
	if !reflect.DeepEqual(got, []string{"JFK", "AAA"}) {
		t.Errorf("single = %v", got)
	}
}
