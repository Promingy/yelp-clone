package validation

func ValidatePassword(p string) map[string]string {
	errs := make(map[string]string)

	if len(p) < 8 {
		errs["password"] = "Password must be at least 8 characters"
		return errs
	}
	if len(p) > 72 {
		errs["password"] = "Password must be at no more than 72 characters"
		return errs
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range p {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case char == '!' || char == '@' || char == '#' || char == '$' || char == '%' || char == '^' || char == '&' || char == '*':
			hasSpecial = true
		}
	}

	if !hasUpper {
		errs["password_uppercase"] = "Password must contain at least one uppercase letter"
	}
	if !hasLower {
		errs["password_lowercase"] = "Password must contain at least one lowercase letter"
	}
	if !hasNumber {
		errs["password_number"] = "Password must contain at least one number"
	}
	if !hasSpecial {
		errs["password_special"] = "Password must contain at least special character (!@#$%^&*)"
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
