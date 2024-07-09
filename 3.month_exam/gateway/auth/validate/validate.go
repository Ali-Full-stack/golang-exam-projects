package validate

import (
	"errors"
	"regexp"
)


func ValidateInfo(name , email, phone, card_number,  home_address string, balance float32)([]error){
	var errors []error

	if err := ValidateString(name); err != nil {
		errors = append(errors, err)
	}
	if err := ValidateEmail(email); err != nil {
		errors = append(errors, err)
	}
	if err := ValidatePhone(phone); err != nil {
		errors = append(errors, err)
	}
	if err := ValidateBankCard(card_number); err != nil {
		errors = append(errors, err)
	}
	
	if err := ValidateHome(home_address); err != nil {
		errors = append(errors, err)
	}
	if err := ValidateBalance(balance); err != nil {
		errors = append(errors, err)
	}
	return errors
}



func ValidateString(name string) error {
	if name == "" {
		return errors.New("field cannot be empty")
	}
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char == ' ')) {
			return errors.New("it should be only  letters")
		}
	}
	return nil
}

func ValidateEmail(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(email) {
		return errors.New("invalid email address")
	}
	return nil
}

func ValidatePhone(phone string) error {
	phoneRegex := `^\+\d{3}\s\d{2}\s\d{3}-\d{2}-\d{2}$`
	if !regexp.MustCompile(phoneRegex).MatchString(phone) {
		return errors.New("invalid phone number format")
	}
	return nil
}

func ValidateBankCard(bankCard string) error {
	if len(bankCard) != 16 {
		return errors.New("bank card number must be 16 digits long")
	}
	return nil
}

func ValidateBalance(balance float32) error {
	if balance < 0 {
		return errors.New("balance cannot be negative")
	}
	return nil
}

func ValidateHome(street string) error {
	if street == "" {
		return errors.New("field cannot be empty")
	}
	hasNumber := false
	for _, char := range street {
		if char >= '0' && char <= '9' {
			hasNumber = true
			break
		}
	}
	if !hasNumber {
		return errors.New("street/home_number  must contain at least one number")
	}
	return nil
}
