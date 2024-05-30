package data

import (
	"fmt"
	"strconv"
)

// Declare a custom Runtime type, which has the underlying type int32 (the same as our // Movie struct field).
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
