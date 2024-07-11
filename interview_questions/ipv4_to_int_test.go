package main

import "testing"

func TestBasic(t *testing.T) {
	res, err := ipv4ToInt("192.0.2.235")
	if err != nil {
		t.Error("should not return err")
	}
	want := uint32(3221226219)
	if res != want {
		t.Errorf("want: %v, got: %v", want, res)
	}
}

func TestStringWithWhitespace(t *testing.T) {
	res, err := ipv4ToInt("192.0.2 .235")
	if err != nil {
		t.Fatalf("should not return err: %v", err)
	}
	want := uint32(3221226219)
	if res != want {
		t.Errorf("want: %v, got: %v", want, res)
	}
}

func TestInvalidAddr(t *testing.T) {
	_, err := ipv4ToInt("192.0.2 .xyz")
	if err == nil {
		t.Fatal("should return err")
	}
}

func TestInvalidAddrV2(t *testing.T) {
	_, err := ipv4ToInt("123.456.789.000")
	if err == nil {
		t.Fatal("should return err")
	}
}
