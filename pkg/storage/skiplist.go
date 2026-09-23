package storage

import (
	"math/rand"
	"sync"
)

const MaxLevel = 16
const Probability = 0.5

type Node struct {
	key     string
	value   []byte
	forward []*Node
}

type SkipList struct {
	mu     sync.RWMutex
	head   *Node
	level  int
	length int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head: &Node{
			forward: make([]*Node, MaxLevel),
		},
		level: 1,
	}
}

func (s *SkipList) Get(key string) ([]byte, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	current := s.head
	for i := s.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
	}

	current = current.forward[0]
	if current != nil && current.key == key {
		return current.value, true
	}
	return nil, false
}

func (s *SkipList) Put(key string, value []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	update := make([]*Node, MaxLevel)
	current := s.head

	for i := s.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}

	current = current.forward[0]
	if current != nil && current.key == key {
		current.value = value
		return
	}

	newLevel := 1
	for rand.Float64() < Probability && newLevel < MaxLevel {
		newLevel++
	}

	if newLevel > s.level {
		for i := s.level; i < newLevel; i++ {
			update[i] = s.head
		}
		s.level = newLevel
	}

	newNode := &Node{
		key:     key,
		value:   value,
		forward: make([]*Node, newLevel),
	}

	for i := 0; i < newLevel; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
	s.length++
}
