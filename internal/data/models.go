package data

import (
	"database/sql"
	"errors"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Create a Models struct which wraps the MovieModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
// 這種模式的好處之一是，從 API 處理程序的角度來看，在電影表上執行操作的程式碼將非常清晰且可讀
// ex: app.models.Movies.Insert(...)
// 通用結構也易於擴展。當我們將來創建更多資料庫模型時，我們所要做的就是將它們包含在 Models 結構中，
// 它們將自動可供我們的 API 處理程序使用。
type Models struct {
	// Set the Movies field to be an interface containing the methods that both the
	// 'real' model and mock model need to support.
	Movies interface {
		Insert(movie *Movie) error
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
	}
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}

// Create a helper function which returns a Models instance containing the mock models
// only.
func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}
