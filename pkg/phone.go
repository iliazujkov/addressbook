package pkg

import (
	"regexp"
)

func PhoneNormalize(phone string) (normalizedPhone string, err error) {
	re := regexp.MustCompile("[^0-9]")
	normalizedPhone = re.ReplaceAllString(phone, "")

	if len(normalizedPhone) > 10 {
		normalizedPhone = "7" + normalizedPhone[len(normalizedPhone)-10:]
	}

	return normalizedPhone, nil
}
