package main

import "testing"

func TestGfMultiply(t *testing.T) {
	testCases := []struct {
		name            string
		a, b, primitive byte
		expected        byte
	}{
		{"A(x)=(x2+x+1), B(x)=(x2+1), P(x)=x3+x+1", byte(0b111), byte(0b101), byte(0b1011), byte(0b110)},
		{"A(x)=(x3+x2+1), B(x)=(x2+x), P(x)=x4+x+1", byte(0b1101), byte(0b110), byte(0b10011), byte(0b1000)},
		{"A(x)=(x2+1), B(x)=(x3+x2+1), P(x)=x4+x+1", byte(0b101), byte(0b1101), byte(0b10011), byte(0b1100)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := gfMultiply(tc.a, tc.b, tc.primitive)
			if result != tc.expected {
				t.Errorf("Encrypt of (%d, %d, %d) = %d;", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
