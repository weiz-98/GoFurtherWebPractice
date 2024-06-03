package data

import (
	"GoFurtherWebPractice/internal/validator" // New import
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type MovieModel struct {
	DB *sql.DB
}

type MockMovieModel struct{}

type Movie struct {
	// Runtime and Genres fields in the output if and only if they are empty.
	// Notice that the leading comma is still required.
	ID        int64     `json:"id"`                        // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`                         // Timestamp for when the movie is added to our database Use the - directive
	Title     string    `json:"title"`                     // Movie title
	Year      int32     `json:"year    ,omitempty"`        // Movie release year
	Runtime   Runtime   `json:"runtime ,omitempty,string"` // Movie runtime (in minutes) to be represented as a JSON string
	Genres    []string  `json:"genres  ,omitempty"`        // Slice of genres for the movie (romance, comedy, etc.)
	Version   int32     `json:"version"`                   // The version number starts at 1 and will be incremented each
	// time the movie information is updated

	// Use the Runtime type instead of int32. Note that the omitempty directive will
	// still work on this: if the Runtime field has the underlying value 0, then it will
	// be considered empty and omitted -- and the MarshalJSON() method we just made
	// won't be called at all.
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	// Use the Check() method to execute our validation checks. This will add the
	// provided key and error message to the errors map if the check does not evaluate
	// to true. For example, in the first line here we "check that the title is not
	// equal to the empty string". In the second, we "check that the length of the title
	// is less than or equal to 500 bytes" and so on.
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	// Note that we're using the Unique helper in the line below to check that all
	// values in the input.Genres slice are unique.
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}

// The Insert() method accepts a pointer to a movie struct, which should contain the
// data for the new record.
func (m MovieModel) Insert(movie *Movie) error {
	// Define the SQL query for inserting a new record in // the system-generated data.
	query := `
	INSERT INTO movies (title, year, runtime, genres) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, version`
	// Create an args slice containing the values for the placeholder parameters from
	// the movie struct. Declaring this slice immediately next to our SQL query helps to
	// make it nice and clear *what values are being used where* in the query.
	args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	// Use the QueryRow() method to execute the SQL query on our connection pool
	// passing in the args slice as a variadic parameter and scanning the system-
	// generated id, created_at and version values into the movie struct.
	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

// Add a placeholder method for fetching a specific record from the movies table.
func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

// Add a placeholder method for updating a specific record in the movies table.
func (m MovieModel) Update(movie *Movie) error {
	return nil
}

// Add a placeholder method for deleting a specific record from the movies table.
func (m MovieModel) Delete(id int64) error {
	return nil
}

func (m MockMovieModel) Insert(movie *Movie) error {
	return nil // Mock the action...
}
func (m MockMovieModel) Get(id int64) (*Movie, error) {
	return nil, nil // Mock the action...
}
func (m MockMovieModel) Update(movie *Movie) error {
	return nil // Mock the action...
}
func (m MockMovieModel) Delete(id int64) error {
	return nil // Mock the action...
}
