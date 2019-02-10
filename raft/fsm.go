package main

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/hashicorp/raft"
)

type fsm struct {
	mutex      sync.Mutex
	stateValue int
}

type event struct {
	Type  string
	Value int
}

func (f *fsm) Apply(logEntry *raft.Log) interface{} {
	var e event
	if err := json.Unmarshal(logEntry.Data, &e); err != nil {
		panic("Failed unmarshalling Raft log entry. This is a bug.")
	}

	switch e.Type {
	case "set":
		f.mutex.Lock()
		defer f.mutex.Unlock()
		f.stateValue = e.Value
	default:
		panic(fmt.Sprintf("Unrecognized event type in Raft log entry: %s. This is a bug.", e.Type))
	}
	return nil
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	return &fsmSnapshot{stateValue: f.stateValue}, nil
}

func (f *fsm) Restore(serialized io.ReadCloser) error {
	var snapshot fsmSnapshot
	if err := json.NewDecoder(serialized).Decode(&snapshot); err != nil {
		return err
	}
	f.stateValue = snapshot.stateValue
	return nil
}
