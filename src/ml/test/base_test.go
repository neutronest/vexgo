package base_test

import (
	"../base"
	"testing"
)

func TestGetRandf(t *testing.T) {

	if f_random := base.GetRandf(10, 20); f_random > 10 && f_random < 20 {

		t.Log("GetRandf pass!")
	} else {
		t.Error("Bad!")
	}
}

func TestGetRandi(t *testing.T) {

}

func TestGetRandn(t *testing.T) {

}
