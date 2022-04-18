package clnt

import (
	"testing"
)

func TestGetOneSupply(t *testing.T) {
	for i := 1; i < 4; i++ {
		supply, err := GetSupplyById(int64(i))
		if err != nil {
			t.Errorf("%v", err)
		}
		t.Logf("One Supply: %v", supply)
	}
}

func TestGetAllSupplies(t *testing.T) {
	var supplies []Supplies
	supplies, err := GetAllSupplies()
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("All Supplies: %v", supplies)
}
