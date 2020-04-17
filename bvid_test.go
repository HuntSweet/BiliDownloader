package main

import "testing"

func testgetCids(t *testing.T)  {
	_,err := getCids("1qk4y197bB")
	if err != nil{
		t.Error("sa")
	}
}

func TestWriteCounter_Write(t *testing.T) {
	t.Run("sa",testgetCids)
}