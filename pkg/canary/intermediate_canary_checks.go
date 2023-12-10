package canary

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/url"
	"strings"
)

func checkDirectMatch(input, canary string) bool {
	return strings.Contains(input, canary)
}

func checkBase64(input, canary string) bool {
	return strings.Contains(input, base64.StdEncoding.EncodeToString([]byte(canary)))
}

func checkURLEncode(input, canary string) bool {
	decodedInput, err := url.QueryUnescape(input)
	if err != nil {
		return false
	}
	return strings.Contains(decodedInput, canary)
}

func checkHTMLEntityEncode(input, canary string) bool {
	htmlEntities := map[string]string{"&": "&amp;", "<": "&lt;", ">": "&gt;", "\"": "&quot;", "'": "&#39;"}
	escapedCanary := canary
	for k, v := range htmlEntities {
		escapedCanary = strings.ReplaceAll(escapedCanary, k, v)
	}
	return strings.Contains(input, escapedCanary)
}

func checkHex(input, canary string) bool {
	return strings.Contains(input, hex.EncodeToString([]byte(canary)))
}

func checkUTF16(input, canary string) bool {
	utf16Encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()

	canaryBytes, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(canary)), utf16Encoder))

	return bytes.Contains([]byte(input), canaryBytes)
}

func checkGzip(input, canary string) bool {
	var gzBuffer bytes.Buffer
	gzWriter := gzip.NewWriter(&gzBuffer)
	_, _ = gzWriter.Write([]byte(canary))
	_ = gzWriter.Close()
	return strings.Contains(input, gzBuffer.String())
}

func checkForCanary(input, canary string) bool {
	return checkDirectMatch(input, canary) ||
		checkBase64(input, canary) ||
		checkURLEncode(input, canary) ||
		checkHTMLEntityEncode(input, canary) ||
		checkHex(input, canary) ||
		checkUTF16(input, canary) ||
		checkGzip(input, canary)
}
