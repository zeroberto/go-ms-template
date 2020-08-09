package tool

import (
	"testing"

	"github.com/zeroberto/go-ms-template/tool"
)

func TestContains(t *testing.T) {
	expected := true
	values := []interface{}{"A", "B", "C"}
	target := "A"

	got := tool.Contains(target, values)

	if expected != got {
		t.Errorf("Contains() failed, expected %v, got %v", expected, got)
	}
}

func TestContainsFailure(t *testing.T) {
	expected := false
	values := []interface{}{"A", "B", "C"}
	target := "D"

	got := tool.Contains(target, values)

	if expected != got {
		t.Errorf("Contains() failed, expected %v, got %v", expected, got)
	}
}

func TestContainsString(t *testing.T) {
	expected := true
	values := []string{"A", "B", "C"}
	target := "A"

	got := tool.ContainsString(target, values)

	if expected != got {
		t.Errorf("Contains() failed, expected %v, got %v", expected, got)
	}
}

func TestContainsStringFailure(t *testing.T) {
	expected := false
	values := []string{"A", "B", "C"}
	target := "D"

	got := tool.ContainsString(target, values)

	if expected != got {
		t.Errorf("Contains() failed, expected %v, got %v", expected, got)
	}
}

func TestContainsKey(t *testing.T) {
	expected := true
	values := map[interface{}]interface{}{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	targetKey := "A"

	got := tool.ContainsKey(targetKey, values)

	if expected != got {
		t.Errorf("ContainsKey() failed, expected %v, got %v", expected, got)
	}
}

func TestContainsKeyFailure(t *testing.T) {
	expected := false
	values := map[interface{}]interface{}{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	targetKey := "D"

	got := tool.ContainsKey(targetKey, values)

	if expected != got {
		t.Errorf("ContainsKey() failed, expected %v, got %v", expected, got)
	}
}

func TestContainsStringKey(t *testing.T) {
	expected := true
	values := map[string]interface{}{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	targetKey := "A"

	got := tool.ContainsStringKey(targetKey, values)

	if expected != got {
		t.Errorf("ContainsKey() failed, expected %v, got %v", expected, got)
	}
}

func TestContainsStringKeyFailure(t *testing.T) {
	expected := false
	values := map[string]interface{}{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	targetKey := "D"

	got := tool.ContainsStringKey(targetKey, values)

	if expected != got {
		t.Errorf("ContainsKey() failed, expected %v, got %v", expected, got)
	}
}
