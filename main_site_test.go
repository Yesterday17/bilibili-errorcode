package bilierrorcode

import "testing"

func TestMainSiteNotMatch(t *testing.T) {
	result := getMainSiteDetail(-1)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}

	result = getMainSiteDetail(990001)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}

	result = getMainSiteDetail(0)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}
}

func TestMainSiteCode(t *testing.T) {
	result := getMainSiteDetail(10001)
	if result.Message != "ArchiveTypeForbidAdd" {
		t.Errorf("error message mismatch: expected %s, but got %s", "ArchiveTypeForbidAdd", result.Message)
	}
}
