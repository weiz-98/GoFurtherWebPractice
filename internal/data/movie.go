package data

import (
	"GoFurtherWebPractice/internal/validator" // New import
	"database/sql"
	"errors"
	"time"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type MovieModel struct {
	DB *sql.DB
}

// Create a Models struct which wraps the MovieModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
// 這種模式的好處之一是，從 API 處理程序的角度來看，在電影表上執行操作的程式碼將非常清晰且可讀
// ex: app.models.Movies.Insert(...)
// 通用結構也易於擴展。當我們將來創建更多資料庫模型時，我們所要做的就是將它們包含在 Models 結構中，
// 它們將自動可供我們的 API 處理程序使用。
type Models struct {
	Movies MovieModel
}

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

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}

// Add a placeholder method for inserting a new record in the movies table.
func (m MovieModel) Insert(movie *Movie) error {
	return nil
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
