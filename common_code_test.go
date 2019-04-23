package bilierrorcode

import "testing"

func TestCommonCodeNotMatch(t *testing.T) {
	result := getCommonCodeDetail(0)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}
}

func TestCommonCode(t *testing.T) {
	result := getCommonCodeDetail(-1)
	if result.Message != "AppKeyInvalid" {
		t.Errorf("error message mismatch: expected %s, but got %s", "AppKeyInvalid", result.Message)
	}
}
