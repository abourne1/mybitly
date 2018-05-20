package lib

import (
	"testing"
)

func TestGetURLSuffix(t *testing.T) {
	tests := []struct{
		URI string
		Expected string
	}{
		{
			URI: "",
			Expected: "",
		},
		{
			URI: "/",
			Expected: "",
		},
		{
			URI: "/test-url",
			Expected: "",
		},
		{
			URI: "/test-url/with-more-url?and=params",
			Expected: "/with-more-url?and=params",
		},
	}

	for _, test := range tests {
		uriSuffix := GetURISuffix(test.URI)
		if uriSuffix != test.Expected{
			t.Errorf("URI suffix did not match expected. %v != %v", uriSuffix, test.Expected)
		}
	}
}