package lib

import (
	"errors"
	"math"
	"strings"
)

const (
	sixtyTwo float64 = 62
	baseConversionChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

// GetURISuffix returns the portion of the URL that comes after the short link
// TODO: replace logic with regex match
func GetURISuffix(uri string) string {
	if len(uri) == 0 {
		return ""
	}

	cleanURI := uri[1:] // remove initial backslash from uri
	i := strings.Index(cleanURI, "/")
	if i == -1 {
		// URL does not contain suffix
		return ""
	}
	return cleanURI[i+1:]
}

// ConvertBase converts an integer to base 62
// TODO: update function to return fixed-width base 62 numbers
func ConvertBase(uuid int64, base int64) (*string, error) {
	if uuid <= 0 {
		return nil, errors.New("Cannot convert non-positive integers")
	}
	if base <= 0 || base > 62 {
		return nil, errors.New("Base must be in (0,62]")
	}

	newBaseStr := ""
	dividend := uuid
	remainder := int64(0)
	// for each time the base goes into the dividend
	// add a digit to the new base string 
	// whose value is equal to the remainder of dividend/base
	for dividend > 0 {
		remainder = int64(math.Mod(float64(dividend), float64(base)))
		dividend = dividend / base
		newBaseStr = baseConversionChars[remainder:remainder+1] + newBaseStr
	}

	return &newBaseStr, nil
}