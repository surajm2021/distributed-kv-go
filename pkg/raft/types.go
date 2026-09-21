package raft

import (
	"sync"
	"time"
)

type Role string

const (
	RoleFollower  Role = "Follower"
	RoleCandidate Role = "Candidate"
	RoleLeader    Role = "Leader"
)

type LogEntry struct {
	Index uint64
	Term  uint64
	Key   string
	Value []byte
}

type NodeState struct {
	mu          sync.RWMutex
	CurrentTerm uint64
	VotedFor    string
	Role        Role
	LastHeartbeat time.Time
}

func NewNodeState() *NodeState {
	return &NodeState{
		CurrentTerm: 0,
		Role:        RoleFollower,
		LastHeartbeat: time.Now(),
	}
}
