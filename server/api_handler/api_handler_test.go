package api_handler

import(
	"testing"
)

func TestDecodeResponseJSON(t *testing.T) {
	input := []uint8{123,10,32,32,32,32,34,71,108,111,98,97,108,32,81,117,111,
		116,101,34,58,32,123,10,32,32,32,32,32,32,32,32,34,48,49,46,32,115,121,109,
		98,111,108,34,58,32,34,73,66,77,34,44,10,32,32,32,32,32,32,32,32,34,48,50,46,32,
		111,112,101,110,34,58,32,34,49,57,54,46,55,57,48,48,34,44,10,32,32,32,32,32,32,32,
		32,34,48,51,46,32,104,105,103,104,34,58,32,34,49,57,55,46,51,56,48,48,34,44,10,32,32,
		32,32,32,32,32,32,34,48,52,46,32,108,111,119,34,58,32,34,49,57,52,46,51,57,48,48,34,44,
		10,32,32,32,32,32,32,32,32,34,48,53,46,32,112,114,105,99,101,34,58,32,34,49,57,54,46,49,48,
		48,48,34,44,10,32,32,32,32,32,32,32,32,34,48,54,46,32,118,111,108,117,109,101,34,58,32,34,50,51,50,
		49,57,54,49,34,44,10,32,32,32,32,32,32,32,32,34,48,55,46,32,108,97,116,101,115,116,32,116,114,97,100,105,
		110,103,32,100,97,121,34,58,32,34,50,48,50,52,45,48,56,45,50,51,34,44,10,32,32,32,32,32,32,32,32,34,48,56,46,32,
		112,114,101,118,105,111,117,115,32,99,108,111,115,101,34,58,32,34,49,57,53,46,57,54,48,48,34,44,10,32,32,32,32,32,32,
		32,32,34,48,57,46,32,99,104,97,110,103,101,34,58,32,34,48,46,49,52,48,48,34,44,10,32,32,32,32,32,32,32,32,34,49,48,46,32,
		99,104,97,110,103,101,32,112,101,114,99,101,110,116,34,58,32,34,48,46,48,55,49,52,37,34,10,32,32,32,32,125,10,125}
	
	want := Quote{}
	want.Global_Quote = Global_Quote{}
	want.Global_Quote.Symbol = "IBM"
	want.Global_Quote.Open = "196.7900"
	want.Global_Quote.High = "197.3800"
	want.Global_Quote.Low = "194.3900"
	want.Global_Quote.Price = "196.1000"
	want.Global_Quote.Volume = "2321961"
	want.Global_Quote.Ltd = "2024-08-23"
	want.Global_Quote.Previous_close = "195.9600"
	want.Global_Quote.Change = "0.1400"
	want.Global_Quote.Change_percent = "0.0714%"

	msg, err := DecodeResponseJSON(input)

	if (msg != want || err != nil) {
		t.Fatalf(`DecodeResponseJSON(good input) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDecodeResponseJSONEmpty(t *testing.T) {
	msg, err := DecodeResponseJSON([]uint8{})

	if (err == nil) {
		t.Fatalf(`DecodeResponseJSON(bad input) = wanted unexpected end of JSON input, got %q, nil`, msg)
	}
}

func TestSymbolSearchRealSymbol(t *testing.T) {
	msg, err := SymbolSearch("IBM", "demo")

	if (err != nil) {
		t.Fatalf(`SymbolSearch("Test") = wanted no error, got %q, %v`, msg, err)
	}
}

func TestCheckDecodedJSONReal(t *testing.T) {
	keyword := "Test"

	input := Quote{}
	input.Global_Quote = Global_Quote{}
	input.Global_Quote.Symbol = keyword
	input.Global_Quote.Open = "196.7900"
	input.Global_Quote.High = "197.3800"
	input.Global_Quote.Low = "194.3900"
	input.Global_Quote.Price = "196.1000"
	input.Global_Quote.Volume = "2321961"
	input.Global_Quote.Ltd = "2024-08-23"
	input.Global_Quote.Previous_close = "195.9600"
	input.Global_Quote.Change = "0.1400"
	input.Global_Quote.Change_percent = "0.0714%"


	msg, err := CheckDecodedJSON(input, keyword)

	if (msg != input || err != nil) {
		t.Fatalf(`CheckDecodedJSON(good input, "Test") = %q, %v, want match for %#q, nil`, msg, err, input)
	}
}

func TestCheckDecodedJSONFake(t *testing.T) {
	input := Quote{}

	msg, err := CheckDecodedJSON(input, "Test")

	if (msg != input || err == nil) {
		t.Fatalf(`CheckDecodedJSON(bad input, "Test") = %q, %v, want match for %#q, nil`, msg, err, input)
	}



}

