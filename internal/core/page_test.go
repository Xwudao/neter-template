package core

import (
	"encoding/json"
	"testing"
)

func TestPageJSON(t *testing.T) {
	page := NewPage([]string{"first", "second"}, 2)

	data, err := json.Marshal(page)
	if err != nil {
		t.Fatalf("marshal page: %v", err)
	}

	if got, want := string(data), `{"list":["first","second"],"total":2}`; got != want {
		t.Fatalf("unexpected JSON: got %s, want %s", got, want)
	}
}
