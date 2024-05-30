package data

import "time"

type Movie struct {
	// Runtime and Genres fields in the output if and only if they are empty.
	// Notice that the leading comma is still required.
	ID        int64     `json:"id"`                        // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`                         // Timestamp for when the movie is added to our database Use the - directive
	Title     string    `json:"title"`                     // Movie title
	Year      int32     `json:"year    ,omitempty"`        // Movie release year
	Runtime   int32     `json:"runtime ,omitempty,string"` // Movie runtime (in minutes) to be represented as a JSON string
	Genres    []string  `json:"genres  ,omitempty"`        // Slice of genres for the movie (romance, comedy, etc.)
	Version   int32     `json:"version"`                   // The version number starts at 1 and will be incremented each
	// time the movie information is updated
}
