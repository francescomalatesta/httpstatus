package statuscodes

import (
	"testing"
)

func TestIsInList(t *testing.T) {
	allowedCodes := GetAllowedCodes()

	if allowedCodes.IsInList(200) == false {
		t.Error("Expected to find 200 in allowed codes list")
	}
}

func TestIsNotInList(t *testing.T) {
	allowedCodes := GetAllowedCodes()

	if allowedCodes.IsInList(589) == true {
		t.Error("Expected to not find 589 in allowed codes list")
	}
}
