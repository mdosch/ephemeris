package ephemeris

import (
	"strings"
	"testing"
)

// TestDemoSite - Test we can load/parse our demo-site
func TestDemoSite(t *testing.T) {

	x, err := New("_demo/data", "_demo/comments", "https://steve.kemp.example.com/")
	if err != nil {
		t.Fatalf("error creating site: %s", err.Error())
	}

	//
	// Blog entries
	//
	url := "https://steve.kemp.example.com/"

	entries := x.Entries()

	//
	// ensure each entry has the prefix
	//
	for _, ent := range entries {
		if !strings.HasPrefix(ent.Link, url) {
			t.Errorf("blog entry doesn't have our url-prefix")
		}
	}
}

// TestDemoSiteSorting - Ensure that we find the most recent posts first
func TestDemoSiteSorting(t *testing.T) {

	x, err := New("_demo/data", "_demo/comments", "https://steve.kemp.example.com/")
	if err != nil {
		t.Fatalf("error creating site: %s", err.Error())
	}

	recent := x.Recent(10)

	//
	// First post is the most recent
	//
	expected := "2019-10-12 09:12:00 +0000 UTC"

	//
	// Compare the date, completely.
	//
	if recent[0].Date.String() != expected {
		t.Fatalf("Most recent post had date '%s' not '%s'", recent[0].Date, expected)
	}

	//
	// Compare the fields
	//
	b := recent[0]
	if b.Year() != "2019" {
		t.Fatalf("Most recent date had wrong year")
	}
	if b.MonthName() != "October" {
		t.Fatalf("Most recent date had month")
	}
	if b.MonthNumber() != "10" {
		t.Fatalf("Most recent date had month")
	}
}

// TestMissingDirectories tests that missing paths fail, as expected
func TestMissingDirectories(t *testing.T) {

	// Empty path to posts & comments is OK
	_, err := New("", "", "https://example.com/")
	if err != nil {
		t.Fatalf("error creating site: %s", err.Error())
	}

	// Bogus path to entries is an error
	_, err = New("/path/bogus:c:/", "", "https://example.com")
	if err == nil {
		t.Fatalf("Bogus site-directory should have made an error, but didn't")
	}

	// Bogus path to comments is an error
	_, err = New("", "/path/bogus:c:/", "https://example.com")
	if err == nil {
		t.Fatalf("Bogus site-directory should have made an error, but didn't")
	}

}
