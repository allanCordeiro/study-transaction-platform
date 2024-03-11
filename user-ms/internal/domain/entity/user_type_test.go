package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserType(t *testing.T) {
	userType := Customer

	assert.Equal(t, "customer", userType.String())
	assert.Equal(t, uint8(1), userType.EnumIndex())

}
