package canary

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"mime/quotedprintable"
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
	return strings.Contains(input, url.QueryEscape(canary))
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
	utf16Canary := ""
	for _, r := range canary {
		utf16Canary += fmt.Sprintf("%c%c", r>>8&0xFF, r&0xFF)
	}
	return strings.Contains(input, utf16Canary)
}

func checkQuotedPrintable(input, canary string) bool {
	var qpBuffer bytes.Buffer
	qpWriter := quotedprintable.NewWriter(&qpBuffer)
	_, _ = qpWriter.Write([]byte(canary))
	_ = qpWriter.Close()
	return strings.Contains(input, qpBuffer.String())
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
		checkQuotedPrintable(input, canary) ||
		checkGzip(input, canary)
}
