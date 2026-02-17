package valueobject

import (
	"testing"
)

func TestCep_Validate(t *testing.T) {
	// Test valid CEP
	validCep, err := NewCep("12345678")
	if err != nil {
		t.Fatalf("expected no error for valid CEP, got %v", err)
	}
	if validCep.value != "12345678" {
		t.Fatalf("expected CEP to be '12345678', got '%s'", validCep)
	}

	// Test invalid CEP (too short)
	_, err = NewCep("1234567")
	if err == nil || err.Error() != "Invalid CEP format" {
		t.Fatalf("expected 'Invalid CEP format' error for too short CEP, got %v", err)
	}

	// Test invalid CEP (too long)
	_, err = NewCep("123456789")
	if err == nil || err.Error() != "Invalid CEP format" {
		t.Fatalf("expected 'Invalid CEP format' error for too long CEP, got %v", err)
	}

	// Test invalid CEP (non-numeric)
	_, err = NewCep("abcdefgh")
	if err == nil || err.Error() != "Invalid CEP format" {
		t.Fatalf("expected 'Invalid CEP format' error for non-numeric CEP, got %v", err)
	}
}
