package main

import (
	"testing"
	c "github.com/xytan0056/depwithgen/.gen/client"
	// a "github.com/xytan0056/depwithgen/api"
)

func TestSomeTest(t *testing.T) {
	t.Logf("testing on a test")
	t.Log(c.Client("test"))
	// a.API()
}