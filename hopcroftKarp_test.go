package hopcroftKarp

import (
	"testing"
)

func TestNewHK(t *testing.T) {
	g1 := &Vertex{"1"}
	g3 := &Vertex{"3"}
	g6 := &Vertex{"6"}
	g8 := &Vertex{"8"}
	graph := map[*Vertex][]*Vertex{&Vertex{"a"}: {g1, g3},
		&Vertex{"c"}: {g1, g3},
		&Vertex{"d"}: {g3, g6},
		&Vertex{"h"}: {g8},
	}
	expect1 := map[string]string{"1": "a", "a": "1", "3": "c", "c": "3", "6": "d", "d": "6", "8": "h", "h": "8"}
	expect2 := map[string]string{"3": "a", "a": "3", "1": "c", "c": "1", "6": "d", "d": "6", "8": "h", "h": "8"}

	hk := NewHK(graph)
	h := hk.MaximumMatching()
	flag := true
	for k, v := range h {
		t.Log(k, v)
		if expect1[k.vertex] != v.vertex {
			flag = false
			break
		}
	}
	if !flag {
		for k, v := range h {
			if expect2[k.vertex] != v.vertex {
				flag = false
				break
			}
		}
	}

	if !flag {
		t.Fail()
	}
}
