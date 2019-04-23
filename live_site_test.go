package bilierrorcode

import "testing"

func TestLiveSiteNotMatch(t *testing.T) {
	result := getLiveSiteDetail(-1)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}

	result = getLiveSiteDetail(2000000)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}

	result = getLiveSiteDetail(999999)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}
}

func TestLiveSiteCode(t *testing.T) {
	result := getLiveSiteDetail(1000000)
	if result.Message != "CoinNotEnough" {
		t.Errorf("error message mismatch: expected %s, but got %s", "CoinNotEnough", result.Message)
	}
}
