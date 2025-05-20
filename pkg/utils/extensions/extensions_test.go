package extensions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidatorValidate(t *testing.T) {
	validator := NewValidator([]string{".go"}, nil)
	require.Equal(t, Regular, validator.ValidatePath("main.go"), "could not validate correct data with extensions")
	require.Equal(t, Skip, validator.ValidatePath("main.php"), "could not validate correct data with wrong extension")

	validator = NewValidator(nil, []string{".php"})
	require.Equal(t, Skip, validator.ValidatePath("main.php"), "could not validate correct data with deny list extension")
	require.Equal(t, Regular, validator.ValidatePath("main.go"), "could not validate correct data with no custom extensions")

	validator = NewValidator([]string{"png"}, nil)
	require.Equal(t, Media, validator.ValidatePath("main.png"), "could not validate correct data with default denylist bypass")

	validator = NewValidator(nil, nil)
	require.Equal(t, Media, validator.ValidatePath("main.png"), "could not validate correct data with default denylist bypass")

}
