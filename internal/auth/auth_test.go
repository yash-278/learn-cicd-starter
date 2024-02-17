package auth

import (
	"reflect"
	"testing"
)

func TestGetAPIKeyNoAuthHeader(t *testing.T) {

	got, err := GetAPIKey(map[string][]string{"Authorization": {"ApiKey a"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "a"

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
