# Reconstruct Itinerary

> NeetCode 150 · Advanced Graphs — an Eulerian path via Hierholzer: consume departures smallest-first, finalize an airport only when its tickets are spent, read the route backwards. Go, stdlib only, tested.

## Problem

Use every flight ticket exactly once starting from JFK, returning the lexically smallest full itinerary. Using each *edge* once (tickets), not each vertex, makes this an Eulerian path, not a Hamiltonian search.

## Approach

Naive greedy — always fly to the smallest destination — walks into dead ends (the JFK→KUL trap, pinned by a test). Hierholzer's post-order insight fixes it without backtracking: recurse into departures (smallest first, via sorted lists and a cursor), and append an airport to the route only when it has **no unused tickets left** — dead ends get finalized deepest-first and end up at the route's *tail* once the whole thing is reversed. Edges are consumed exactly once, so the run is linear in tickets; the lexical sort is what makes the answer canonical rather than just valid.

**Complexity:** O(E log E) for sorting departures, O(E) walk.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
