package bilierrorcode

import "testing"

func TestOpenPlatformSiteNotMatch(t *testing.T) {
	result := getOpenPlatformSiteDetail(-1)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}

	result = getOpenPlatformSiteDetail(1999999)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}

	result = getOpenPlatformSiteDetail(3000000)
	if result.Message != "" {
		t.Errorf("error message mismatch: expected %s, but got %s", "EMPTY_STRING", result.Message)
	}
}

func TestOpenPlatformSiteCode(t *testing.T) {
	result := getOpenPlatformSiteDetail(2000000)
	if result.Message != "TicketUnKnown" {
		t.Errorf("error message mismatch: expected %s, but got %s", "TicketUnKnown", result.Message)
	}
}
