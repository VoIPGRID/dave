package main

import "testing"

func TestNextBranch(t *testing.T) {
	next := nextBranch("release-", "release-2.87.0", "release-2.89.0", "release-2.90.0")
	if next != "release-2.91.0" {
		t.Errorf("expected \"release-2.91.0\", got %q", next)
	}
	next = nextBranch("release-", "garbage", "junk", "release-2.90.0", "something", "release-2.87.0", "release-2.89.0")
	if next != "release-2.91.0" {
		t.Errorf("expected \"release-2.91.0\", got %q", next)
	}
}
