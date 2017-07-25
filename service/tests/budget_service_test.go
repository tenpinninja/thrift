package tests

import (
	"testing"
)

func TestCreateBudget(t *testing.T) {
	var testID = 1
	var actualID int

	if actualID != testID {
		t.Errorf("TestCreateBudget: Expected budget id of %d, but recieved a budget id of %d", testID, actualID)
	}
}
