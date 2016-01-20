package goonix

import (
	"testing"
	"time"
)

// test local telnet port
func TestPortAvailable(t *testing.T) {
	n := Network{}
	retval, err := n.CheckPort("localhost", 22, 2*time.Second)
	if !retval {
		t.Error("Expected true, got false")
	}
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

}

// test non-existant host/port
func TestFailedPortAvailable(t *testing.T) {
	n := Network{}
	retval, err := n.CheckPort("fuuuuuuu.com", 99, 2*time.Second)
	if retval {
		t.Fail()
	}
	if err == nil {
		t.Fail()
	}
}
