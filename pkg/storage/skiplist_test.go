package storage

import (
	"bytes"
	"testing"
)

func TestSkipListBasicOperations(t *testing.T) {
	sl := NewSkipList()
	sl.Put("cluster_name", []byte("prod-east-1"))
	sl.Put("max_nodes", []byte("5"))

	val, found := sl.Get("cluster_name")
	if !found || !bytes.Equal(val, []byte("prod-east-1")) {
		t.Fatalf("expected prod-east-1, got %s", string(val))
	}

	_, foundMissing := sl.Get("non_existent")
	if foundMissing {
		t.Fatal("expected non_existent key to return false")
	}
}
