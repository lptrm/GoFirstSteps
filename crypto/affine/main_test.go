package main

import "testing"

func TestGCD(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positive numbers", 48, 18, 6},
		{"One number is zero", 0, 5, 5},
		{"Both numbers are zero", 0, 0, 0},
		{"Negative numbers", -48, -18, 6}, // Nur wenn Ihre Funktion negative Zahlen unterstützt
		{"One number is a multiple of the other", 3, 9, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := gcd(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("GCD of (%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
func TestModInverse(t *testing.T) {
	testCases := []struct {
		name     string
		a, m     int
		expected int
	}{
		{"Positive numbers", 3, 26, 9},
		{"Negative numbers", -3, 26, 17},
		{"One number is zero", 0, 5, -1},
		{"Both numbers are zero", 0, 0, -1},
		{"No inverse", 2, 26, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := modInverse(tc.a, tc.m)
			if result != tc.expected {
				t.Errorf("ModInverse of (%d, %d) = %d; want %d", tc.a, tc.m, result, tc.expected)
			}
		})
	}
}
func TestEncrypt(t *testing.T) {
	testCases := []struct {
		name     string
		a, b, x  int
		expected int
	}{
		{"Positive numbers", 9, 13, 0, 13},
		{"Negative numbers", -9, 13, 0, 13},
		{"One number is zero", 9, 0, 0, 0},
		{"Both numbers are zero", 0, 0, 0, 0},
		{"Encrypt number one", 9, 13, 1, 22},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := encrypt(tc.a, tc.b, tc.x)
			if result != tc.expected {
				t.Errorf("Encrypt of (%d, %d, %d) = %d; want %d", tc.a, tc.b, tc.x, result, tc.expected)
			}
		})
	}
}
func TestDecrypt(t *testing.T) {
	testCases := []struct {
		name     string
		a, b, y  int
		expected int
	}{
		{"Positive numbers", 9, 13, 0, 13},
		{"Negative numbers", -9, 13, 0, 13},
		{"One number is zero", 9, 0, 0, 0},
		{"Both numbers are zero", 0, 0, 0, -1},
		{"Decrypt number one", 9, 13, 22, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := decrypt(tc.a, tc.b, tc.y)
			if result != tc.expected {
				t.Errorf("Decrypt of (%d, %d, %d) = %d; want %d", tc.a, tc.b, tc.y, result, tc.expected)
			}
		})
	}
}
func TestAbs(t *testing.T) {
	testCases := []struct {
		name     string
		i        int
		expected int
	}{
		{"Positive numbers", 5, 5},
		{"Negative numbers", -5, 5},
		{"Zero", 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := abs(tc.i)
			if result != tc.expected {
				t.Errorf("Abs of (%d) = %d; want %d", tc.i, result, tc.expected)
			}
		})
	}
}
func TestAffineCipher(t *testing.T) {
	testCases := []struct {
		name, text string
		a, b       int
	}{
		{"Test 1", "ATTACK", 9, 13},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, symbol := range tc.text {
				result := decrypt(tc.a, tc.b, encrypt(tc.a, tc.b, int(symbol-65)))
				if result != int(symbol-65) {
					t.Errorf("Encrypt and Decrypt of (%d, %d, %d) = %d; want %d", tc.a, tc.b, int(symbol-65), result, int(symbol-65))
				}
			}
		})
	}
}
