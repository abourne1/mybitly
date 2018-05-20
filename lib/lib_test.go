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
		if uriSuffix != test.Expected {
			t.Errorf("URI suffix did not match expected. %v != %v", uriSuffix, test.Expected)
		}
	}
}

func TestConvertBase(t *testing.T) {
	tests := []struct{
		UUID int64
		Exp string
		HasErr bool
		Base int64
	}{
		{
			UUID: 1,
			Base: 63,
			HasErr: true,
		},
		{
			UUID: 0,
			Base: 62,
			HasErr: true,
		},
		{
			UUID: 1,
			Exp: "b",
			Base: 62,
		},
		{
			UUID: 62,
			Exp: "ba",
			Base: 62,
		},
		{
			UUID: 63,
			Exp: "bb",
			Base: 62,
		},
	}

	for _, test := range tests {
		base62Str, err := ConvertBase(test.UUID, test.Base)
		if test.HasErr && err == nil {
			t.Errorf("Error expected")
		}
		if err != nil {
			if !test.HasErr {
				t.Errorf(err.Error())
			}
			continue
		}
		if *base62Str != test.Exp {
			t.Errorf("Base62 conversion did not match expected. %v != %v", *base62Str, test.Exp)
		}
	}
}