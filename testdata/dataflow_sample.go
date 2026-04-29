package testdata

import (
	"encoding/json"
	"io"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var review map[string]interface{}
	json.Unmarshal(body, &review)
	name := review["name"]
	query := "SELECT * FROM users WHERE name = " + name.(string)
	_ = query
}

func SimpleAssignment() string {
	x := "hello"
	y := x
	return y
}

func MultiAssignment() {
	a, b := 1, 2
	_ = a
	_ = b
}

func FieldAccess() {
	r := &http.Request{}
	host := r.Host
	_ = host
}

func ChainedFieldAccess(r *http.Request) string {
	return r.URL.Path
}
