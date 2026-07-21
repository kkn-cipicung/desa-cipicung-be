package utils

import (
	"encoding/json"
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

func TestPublicErrorDetailHidesInternalServerErrors(t *testing.T) {
	err := json.Unmarshal([]byte(`{"hamlet_one":"wrong"}`), &struct {
		HamletOne *int64 `json:"hamlet_one"`
	}{})

	got := PublicErrorDetail(http.StatusInternalServerError, err)
	if got != "" {
		t.Fatalf("PublicErrorDetail() = %q, want empty string", got)
	}
}
