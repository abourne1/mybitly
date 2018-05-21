package lib

import (
	"errors"
	"math"
	"strings"
)

const (
	sixtyTwo float64 = 62
)

var (
	baseConversionChars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
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
	
	if uuid <= 0 || uuid >= 56800235584 { // 56800235584 is the max number that can be converted to base62 with 6 characters
		return nil, errors.New("Cannot convert non-positive integers")
	}
	if base <= 0 || base > 62 {
		return nil, errors.New("Base must be in (0,62]")
	}

	idx := 0
	newBaseChars := []byte{'a','a','a','a','a','a'}
	dividend := uuid
	remainder := int64(0)
	// each time the base goes into the dividend add character in new base equal to the value of the remainder
	for dividend > 0 && idx < 6 {
		remainder = int64(math.Mod(float64(dividend), float64(base)))
		dividend = dividend / base
		// change least significant characters first
		newBaseChars[5-idx] = baseConversionChars[remainder]
		idx += 1
	}

	newBaseStr := string(newBaseChars)
	return &newBaseStr, nil
}