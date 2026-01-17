package events

import (
	"strings"
	"testing"
)

func TestShowEventTerminal(t *testing.T) {
	events := []*Event{
		{
			Timestamp: "1900-01-01T12:00:00Z",
			Type:      "note",
			User:      "user1",
			Cwd:       "/tmp",
			Id:        1,
			Content:   "first",
		},
		{
			Timestamp: "1900-01-01T12:00:00Z",
			Type:      "command",
			User:      "user2",
			Cwd:       "/home",
			Id:        2,
			Content:   "second",
		},
	}

	result := ShowEventTerminal(*events[0])
	if result == "" {
		t.Error("ShowEventTerminal returned empty string")
	}
	if !strings.Contains(result, "user1") {
		t.Error("Result should contain username")
	}
	if !strings.Contains(result, "first") {
		t.Error("Result should contain event content")
	}

	result = ShowEventsTerminal(events)
	if !strings.Contains(result, "first") {
		t.Error("Result should contain first event content")
	}
	if !strings.Contains(result, "second") {
		t.Error("Result should contain second event content")
	}
}
