// Package ephemeris holds some minimal code to create a blog.
//
// A blog is largely made up of blog-posts, which are parsed from
// a series of text-files.
//
// Each post will have a small header to include tags, date, title,
// and will be transformed into a simple site.
//
package ephemeris

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Ephemeris holds our site structure.
//
// There are only a few settings for the blog, which are the obvious
// ones - a path pointing to the blog-posts, a URL-prefix for use in
// generation of the output files, and a list of comment files.
type Ephemeris struct {
	// Root is the source of our posts.
	Root string

	// BlogEntries holds the entries we've found
	BlogEntries []BlogEntry

	// CommentFiles holds the filenames of comments we've found.
	CommentFiles []string

	// Prefix is the absolute URL prefix for the blog
	Prefix string
}

// New creates a new site object.
func New(directory string, commentPath string, prefix string) (*Ephemeris, error) {

	// Create object
	x := &Ephemeris{Root: directory, Prefix: prefix}

	// If the comment-path is set we'll load comments
	if commentPath != "" {

		// Now we can find comments - by reading the given
		// directory and adding each of them.
		comments, err := ioutil.ReadDir(commentPath)
		if err != nil {
			return x, err
		}

		// Sort the comments, since we want to show them upon
		// entries in the oldest->newest order.
		sort.Slice(comments, func(i, j int) bool {
			return comments[i].ModTime().Before(comments[j].ModTime())
		})

		// Save the (complete) path to each comment-file in our
		// object, now they're sorted.
		for _, f := range comments {

			// By appending
			x.CommentFiles = append(x.CommentFiles, filepath.Join(commentPath, f.Name()))
		}
	}

	var err error

	//
	// Find the blog-posts, recursively.
	//
	if directory != "" {
		err = filepath.Walk(directory,
			func(path string, info os.FileInfo, err error) error {

				// Error?  Then we're done.
				if err != nil {
					return err
				}

				// Ignore non-text files.
				if !strings.HasSuffix(path, ".txt") {
					return nil
				}

				// Parse the blog-post from the file.
				out, err := NewBlogEntry(path, x)
				if err != nil {
					return fmt.Errorf("failed to parse %s - %s", path, err.Error())
				}

				// Store the result.
				x.BlogEntries = append(x.BlogEntries, out)

				// Continue walking.
				return nil
			})
	}

	// Return the entries we found.
	return x, err
}

// Entries returns the blog-entries contained within a site.  Note that
// the input directory is searched recursively for files matching the
// pattern "*.txt" - this allows you to create entries in sub-directories
// if you wish.
//
// The entries are returned in a random-order, and contain a complete
// copy of all the text in the entries.  This means that there is a reasonable
// amount of memory overhead here.
func (e *Ephemeris) Entries() []BlogEntry {
	return e.BlogEntries
}

// Recent returns the data about the most recent N entries from the
// site
func (e *Ephemeris) Recent(count int) []BlogEntry {

	// The return-value
	var recent []BlogEntry

	// Sort the list of posts by date.
	sort.Slice(e.BlogEntries, func(i, j int) bool {
		a := e.BlogEntries[i].Date
		b := e.BlogEntries[j].Date
		return a.Before(b)
	})

	// We want to include at-max `count` posts.
	//
	// But of course if this is a new blog there might
	// be fewer than that present.  Terminate early in
	// that case.
	c := 0
	for c < len(e.BlogEntries) && c < count {
		ent := e.BlogEntries[len(e.BlogEntries)-1-c]
		recent = append(recent, ent)
		c++
	}

	// All done
	return recent
}
