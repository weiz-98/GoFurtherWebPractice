package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Define an error that our UnmarshalJSON() method can return if we're unable to parse
// or convert the JSON string successfully.
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// Declare a custom Runtime type, which has the underlying type int32 (the same as our
// Movie struct field).
type Runtime int32

// 專門為運行時欄位建立自訂類型，並在此自訂類型上實作 MarshalJSON() 方法。
// Implement a MarshalJSON() method on the Runtime type so that it satisfies the
// json.Marshaler interface. This should return the JSON-encoded value for the movie
// runtime (in our case, it will return a string in the format "<runtime> mins").
func (r Runtime) MarshalJSON() ([]byte, error) {
	// Generate a string containing the movie runtime in the required format.
	jsonValue := fmt.Sprintf("%d mins", r)
	// Use the strconv.Quote() function on the string to wrap it in double quotes. It
	// needs to be surrounded by double quotes in order to be a valid *JSON string*.
	quotedJSONValue := strconv.Quote(jsonValue)
	// Convert the quoted string value to a byte slice and return it.
	return []byte(quotedJSONValue), nil
}

// 為何使用 value receiver (r Runtime)?
// The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values,
// but pointer methods can only be invoked on pointers.

// 我們有一個自訂的運行時類型，我們可以隨時隨地使用它。
// 但有一個缺點。重要的是要注意，在將程式碼與其他套件整合時，使用自訂類型有時會很尷尬，
// 可能需要執行類型轉換以將自訂類型變更為其他套件理解和接受的值或從其他套件中理解和接受的值變更為自訂類型。

// Implement a UnmarshalJSON() method on the Runtime type so that it satisfies the
// json.Unmarshaler interface. IMPORTANT: Because UnmarshalJSON() needs to modify the // receiver (our Runtime type), we must use a pointer receiver for this to work
// correctly. Otherwise, we will only be modifying a copy (which is then discarded when // this method returns).
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// We expect that the incoming JSON value will be a string in the format
	// "<runtime> mins", and the first thing we need to do is remove the surrounding // double-quotes from this string. If we can't unquote it, then we return the
	// ErrInvalidRuntimeFormat error.
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	// Split the string to isolate the part containing the number.
	parts := strings.Split(unquotedJSONValue, " ")
	// Sanity check the parts of the string to make sure it was in the expected format. // If it isn't, we return the ErrInvalidRuntimeFormat error again.
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
	// Otherwise, parse the string containing the number into an int32. Again, if this // fails return the ErrInvalidRuntimeFormat error.
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	// Convert the int32 to a Runtime type and assign this to the receiver. Note that we // use the * operator to deference the receiver (which is a pointer to a Runtime
	// type) in order to set the underlying value of the pointer.
	*r = Runtime(i)
	return nil
}
