package client

import (
	"testing"
)

func TestGetSupplyById(t *testing.T) {
	s, e := GetSupplyById(1)
	if e != nil {
		t.Errorf("%v", e)
	}
	t.Logf("SUPPLY FOR ID 1: %v  Type: %T", s, s)
}

func TestGetAllSupplies(t *testing.T) {
	s, e := GetAllSupplies()
	if e != nil {
		t.Errorf("%v", e)
	}
	t.Logf("ALL SUPPLIES: %v %T %p", s, s, s)

}
