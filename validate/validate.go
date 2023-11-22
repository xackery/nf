package validate

import (
	"fmt"

	"github.com/xackery/nf/raw"
)

// Against validates a value against a struct
func Against(val interface{}) error {
	switch e := val.(type) {
	case *raw.AARankEffect:
		return validateAARankEffect(e)
	case *raw.AAAbility:
		return validateAAAbility(e)
	default:
		return fmt.Errorf("unknown type: %T", val)
	}

}

func validateAARankEffect(val *raw.AARankEffect) error {
	return nil
}

func validateAAAbility(val *raw.AAAbility) error {
	if val.ID == 0 {
		return fmt.Errorf("id is required")
	}
	if val.Name == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}
