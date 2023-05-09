package validatepincode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidatePincode(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    bool
	}{
		//Case 1
		{
			description: "Case 1.1 : Input must be greater than or equal to 6 characters.",
			input:       "17283",
			expected:    false,
		},
		{
			description: "Case 1.2 : input must be greater than or equal to 6 characters.",
			input:       "172839",
			expected:    true,
		},
		//Case 2
		{
			description: "Case 2.1 : input must prevent more than 2 duplicate numbers.",
			input:       "118822",
			expected:    false,
		},
		{
			description: "Case 2.2 : input must prevent more than 2 duplicate numbers.",
			input:       "111762",
			expected:    true,
		},
		//Case 3
		{
			description: "Case 3.1 : input must prevent more than 2 consecutive numbers.",
			input:       "123743",
			expected:    false,
		},
		{
			description: "Case 3.2 : input must prevent more than 2 consecutive numbers.",
			input:       "321895",
			expected:    false,
		},
		{
			description: "Case 3.3 : input must prevent more than 2 consecutive numbers.",
			input:       "124578",
			expected:    true,
		},
		//Case 4
		{
			description: "Case 4.1 : input must prevent more than 2 duplicate numbers.",
			input:       "112233",
			expected:    false,
		},
		{
			description: "Case 4.2 : input must prevent more than 2 duplicate numbers.",
			input:       "882211",
			expected:    false,
		},
		{
			description: "Case 4.3 : input must prevent more than 2 duplicate numbers.",
			input:       "887712",
			expected:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := IsValidPinCode(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
