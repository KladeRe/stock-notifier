package config

import (
	"testing"
	"reflect"
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

func TestDecodeJSONConfig(t *testing.T) {
	var want = []StockConfig{{Symbol: "BMW", Value: 16.43, Buy: true, Notification: "11:00"}, 
		{Symbol: "AAPL", Value: 92.65, Buy: false, Notification: "19:00"}}

	var data = []uint8{91,10,32,32,32,32,123,10,32,32,32,32,32,32,32,32,34,115,121,109,98,111,108,34,58,32,34,66,77,87,34,44,
		10,32,32,32,32,32,32,32,32,34,118,97,108,117,101,34,58,32,49,54,46,52,51,44,10,32,32,32,32,32,32,32,32,34,98,117,121,34,58,
		32,116,114,117,101,44,10,32,32,32,32,32,32,32,32,34,110,111,116,105,102,105,99,97,116,105,111,110,34,58,32,34,49,49,58,48,48,
		34,10,32,32,32,32,125,44,10,32,32,32,32,123,10,32,32,32,32,32,32,32,32,34,115,121,109,98,111,108,34,58,32,34,65,65,80,76,34,44,
		10,32,32,32,32,32,32,32,32,34,118,97,108,117,101,34,58,32,57,50,46,54,53,44,10,32,32,32,32,32,32,32,32,34,98,117,121,34,58,32,102,
		97,108,115,101,44,10,32,32,32,32,32,32,32,32,34,110,111,116,105,102,105,99,97,116,105,111,110,34,58,32,34,49,57,58,48,48,34,10,32,32,32,32,125,10,93}
	msg, err := DecodeJSONConfig(data)

	if (reflect.DeepEqual(want, msg) == false || err != nil) {
		t.Fatalf(`DecodeJSONConfig(./files/config_test.json) = %v, %v, want match for %v, nil`, msg, err, want)
	}
}

func TestFaultyDecodeJSONConfig(t *testing.T) {
	var want = []StockConfig{{Symbol: "BMW", Value: 16.43, Buy: true, Notification: "11:00"}, 
		{Symbol: "AAPL", Value: 92.65, Buy: false, Notification: "19:00"}}

	var data = []uint8{1, 5, 7, 1, 54, 34, 65, 98}
	msg, err := DecodeJSONConfig(data)

	if (reflect.DeepEqual(want, msg) == true || err == nil) {
		t.Fatalf(`DecodeJSONConfig(./files/config_test.json) = %v, %v, want match for %v, nil`, msg, err, want)
	}
}



