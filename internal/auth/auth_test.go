package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type pair struct{ key, value string }
	setupInput := func(args ...pair) http.Header {
		header := make(http.Header)
		for _, arg := range args {
			header.Add(arg.key, arg.value)
		}
		return header
	}

	tests := map[string]struct {
		input  http.Header
		expect string
		hasErr bool
	}{
		"apikey auth": {input: setupInput(pair{key: "Authorization", value: "ApiKey KeyValueHere"}), expect: "KeyValueHere", hasErr: false},
		"basic auth":  {input: setupInput(pair{key: "Authorization", value: "Basic Password"}), expect: "", hasErr: true},
		"no auth":     {input: setupInput(), expect: "", hasErr: true},
		"missing key": {input: setupInput(pair{key: "Authorization", value: "ApiKey"}), expect: "", hasErr: true},
	}

	for name, test := range tests {
		actual, err := GetAPIKey(test.input)
		if actual != test.expect {
			t.Fatalf("%s: expected %v, but received %v", name, test.expect, actual)
		}
		if (err != nil && !test.hasErr) || (err == nil && test.hasErr) {
			t.Fatalf("%s: expected error occurrence to be %v, but received '%v'", name, test.hasErr, err)
		}
	}
}
