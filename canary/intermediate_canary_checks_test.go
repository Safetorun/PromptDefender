package canary

import (
	"testing"
)

func TestCheckDirectMatch(t *testing.T) {
	if !checkDirectMatch("hello world", "world") {
		t.Errorf("Failed direct match test")
	}
	if checkDirectMatch("hello world", "notworld") {
		t.Errorf("False positive in direct match test")
	}
}

func TestCheckBase64(t *testing.T) {
	if !checkBase64("hello d29ybGQ=", "world") {
		t.Errorf("Failed base64 match test")
	}
}

func TestCheckURLEncode(t *testing.T) {
	if !checkURLEncode("hello%20world", "hello world") {
		t.Errorf("Failed URL encode match test")
	}
}

func TestCheckHTMLEntityEncode(t *testing.T) {
	if !checkHTMLEntityEncode("hello &amp; world", "&") {
		t.Errorf("Failed HTML entity encode match test")
	}
}

func TestCheckHex(t *testing.T) {
	if !checkHex("hello 776f726c64", "world") {
		t.Errorf("Failed hex match test")
	}
}

func TestCheckUTF16(t *testing.T) {
	// The string "hello" encoded in UTF-16LE hex is 680065006c006c006f00
	if !checkUTF16("hello \x68\x00\x65\x00\x6c\x00\x6c\x00\x6f\x00", "hello") {
		t.Errorf("Failed UTF-16 match test")
	}
}

func TestCheckGzip(t *testing.T) {
	// Hard to test gzip content as it contains headers that might vary.
	// This test just ensures that a compressed input does not falsely match an uncompressed canary.
	if checkGzip("some gzip compressed string", "world") {
		t.Errorf("False positive in gzip match test")
	}
}

func TestCheckForCanary(t *testing.T) {
	input := "776f726c64 hello%20world hello &amp; world hello=20world d29ybGQ= hello \x68\x00\x65\x00\x6c\x00\x6c\x00\x6f\x00"
	canary := "world"
	if !checkForCanary(input, canary) {
		t.Errorf("Failed canary check")
	}

	input = "no match here"
	if checkForCanary(input, canary) {
		t.Errorf("False positive in canary check")
	}
}
