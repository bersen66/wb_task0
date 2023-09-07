package generators

import (
	"math/rand"
	"time"
)

const (
	DIGITS = "0123456789"
	ENGLET = "aqwertyuiopasdfghjklzxcvbnmQAZWSXEDCRFVTGBYHNUJMIKOLP"
	SPACES = "        "
)

func RandomString(length int, alphabet string) string {
	selectSet := []rune(alphabet)
	var result []rune

	var i = 0
	for ; i < length; i++ {
		result = append(result, selectSet[rand.Int31n(int32(len(alphabet)))])
	}

	return string(result)
}

func RandomPhoneNumber() string {
	return "+" + RandomString(11, DIGITS)
}

var currencies = []string{"RUB", "USD", "EUR", "CNY", "ILS"}

func RandomCurrency() string {
	return currencies[rand.Int31n(int32(len(currencies)))]
}

func randomDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func RandomDate() string {
	date := randomDate()
	return date.Format("2006-01-02T15:04:05Z")
}

var locales = []string{"en", "ru"}

func RandomLocale() string {
	return locales[rand.Int31n(int32(len(locales)))]
}
