package gears

import "testing"

func TestStrSliceDiff(t *testing.T) {
	s1 := []string{"1", "5", "2", "3"}
	s2 := []string{"4", "2", "3"}

	want := []string{"1", "5"}
	got2 := StrSliceDiff2(s1, s2)
	got := StrSliceDiff(s1, s2)
	if got[0] != want[0] || got[1] != want[1] || got2[0] != want[0] || got2[1] != want[1] {
		t.Errorf("\nGot: %v, Want %v\n", got, want)
	}
}
