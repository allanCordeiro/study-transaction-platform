package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	scenarios := []struct {
		name     string
		value    string
		errorMsg error
	}{
		{
			name:     "Given a valid email when validate it should return ok",
			value:    "allan_1985@gmail.com",
			errorMsg: nil,
		},
		{
			name:     "Given an invalid email when validate it should return error",
			value:    "allan_1985.gmail.com",
			errorMsg: ErrInvalidEmail,
		},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			email, err := NewEmail(test.value)
			assert.Equal(t, test.errorMsg, err)
			if err == nil {
				assert.Equal(t, test.value, email.GetEmail())
			}
		})
	}
}
