package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"
)

func TestPublicErrorDetailHidesJSONUnmarshalTypeInternals(t *testing.T) {
	type payload struct {
		HamletOne *int64 `json:"hamlet_one"`
	}

	var input payload
	err := json.Unmarshal([]byte(`{"hamlet_one":"wrong"}`), &input)
	if err == nil {
		t.Fatal("expected json unmarshal error")
	}

	got := PublicErrorDetail(http.StatusBadRequest, err)
	want := "hamlet_one must be a number"
	if got != want {
		t.Fatalf("PublicErrorDetail() = %q, want %q", got, want)
	}
}

func TestPublicErrorDetailIncludesInternalServerErrors(t *testing.T) {
	err := errors.New("database error detail")

	got := PublicErrorDetail(http.StatusInternalServerError, err)
	if got != "database error detail" {
		t.Fatalf("PublicErrorDetail() = %q, want %q", got, "database error detail")
	}
}
