package arithmetic

import "testing"

func TestAdd(t *testing.T) {
	service := NewService()

	result := service.Add(20, 5)

	if result != 25 {
		t.Fatalf("expected 25, got %v", result)
	}
}

func TestDivideByZero(t *testing.T) {
	service := NewService()

	_, err := service.Divide(20, 0)

	if err == nil {
		t.Fatal("expected division by zero error")
	}
}
