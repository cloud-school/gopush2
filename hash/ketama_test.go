package hash

import (
	"testing"
)

func TestKetama(t *testing.T) {
	k := NewKetama(15, 255)
	n := k.Node("Terry-Mao332")
	if n != "node15" {
		t.Error("Terry-Mao332 must hit node13")
	}

	n = k.Node("Terry-Mao2")
	if n != "node13" {
		t.Error("Terry-Mao2 must hit node13")
	}

	n = k.Node("Terry-Mao3")
	if n != "node5" {
		t.Error("Terry-Mao3 must hit node13")
	}

	n = k.Node("Terry-Mao4")
	if n != "node2" {
		t.Error("Terry-Mao4 must hit node13")
	}

	n = k.Node("Terry-Mao5")
	if n != "node6" {
		t.Error("Terry-Mao5 must hit node13")
	}
}
