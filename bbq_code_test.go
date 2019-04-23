package bilierrorcode

import "testing"

func TestBBQCodeNotMatch(t *testing.T) {
	result := getBBQCodeDetail(-1)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}

	result = getBBQCodeDetail(4999999)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}

	result = getBBQCodeDetail(6000000)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}
}

func TestBBQCode(t *testing.T) {
	result := getBBQCodeDetail(3001016)
	if result.Message != "CheckInviteCodeErr" {
		t.Errorf("error message mismatch: expected %s, but got %s", "CheckInviteCodeErr", result.Message)
	}
}
