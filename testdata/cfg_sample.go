package testdata

import (
	"io"
	"net/http"
)

func IfElse(x int) string {
	if x > 0 {
		return "positive"
	} else {
		return "non-positive"
	}
}

func EarlyReturn(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func ForLoop(items []string) string {
	var result string
	for _, item := range items {
		result = item
	}
	return result
}

func SwitchCase(op string) int {
	var result int
	switch op {
	case "add":
		result = 1
	case "sub":
		result = 2
	default:
		result = 0
	}
	return result
}

func NestedIfInFor(items []string) int {
	count := 0
	for _, item := range items {
		if len(item) > 3 {
			count++
		}
	}
	return count
}

func HandleValidated(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad", 400)
		return
	}
	if len(body) == 0 {
		http.Error(w, "empty", 400)
		return
	}
	_ = body
}

func PanicTerminates(x int) {
	if x < 0 {
		panic("negative")
	}
	_ = x
}

func EmptyFunction() {
}

func LinearFunction() int {
	x := 1
	y := x + 2
	return y
}

// Edge case: if-else if without else
func ElseIfNoElse(x int) int {
	result := 0
	if x > 10 {
		result = 1
	} else if x > 5 {
		result = 2
	}
	return result
}

// Edge case: loop with break
func LoopWithBreak(items []string) string {
	for _, item := range items {
		if item == "stop" {
			break
		}
	}
	return "done"
}

// Edge case: loop with continue
func LoopWithContinue(items []string) int {
	count := 0
	for _, item := range items {
		if item == "" {
			continue
		}
		count++
	}
	return count
}

// Edge case: single return function
func SingleReturn() int {
	return 42
}
