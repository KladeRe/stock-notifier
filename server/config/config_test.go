package config

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	want := "[\n    {\n        \"symbol\": \"BMW\",\n        \"value\": 16.43,\n        \"buy\": true,\n        \"notification\": \"11:00\"\n    },\n    " +
	"{\n        \"symbol\": \"AAPL\",\n        \"value\": 92.65,\n        \"buy\": false,\n        \"notification\": \"19:00\"\n    }\n]"
	msg, err := ReadFile("./files/config_test.json")

	if (string(msg) != string(want) || err != nil) {
		t.Fatalf(`readFile(./files/config_test.json) = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}

func TestEmptyReadFile(t *testing.T) {
	want := ""
	msg, err := ReadFile("./files/empty_config_test.json")

	if (string(msg) != want || err != nil) {
		t.Fatalf(`readFile(./files/config_test.json) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestWrongPathReadFile(t *testing.T) {
	want := ""
	msg, err := ReadFile("./files/doesnt_exist.json")

	if (string(msg) != want || err == nil) {
		t.Fatalf(`readFile(./files/doesnt_exist.json) = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}



