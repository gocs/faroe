package main

import (
	"crypto/subtle"
	"mime"
	"net/http"
	"strings"
)

func verifyRequestSecret(secret []byte, r *http.Request) bool {
	if len(secret) == 0 {
		return true
	}
	authorizationHeader, ok := r.Header["Authorization"]
	if !ok {
		return false
	}
	return subtle.ConstantTimeCompare(secret, []byte(authorizationHeader[0])) == 1
}

func verifyJSONContentTypeHeader(r *http.Request) bool {
	contentType, ok := r.Header["Content-Type"]
	if !ok {
		return true
	}
	mediatype, _, err := mime.ParseMediaType(contentType[0])
	if err != nil {
		return false
	}
	return mediatype == "application/json" || mediatype == "text/plain"
}

func verifyJSONAcceptHeader(r *http.Request) bool {
	accept, ok := r.Header["Accept"]
	if !ok {
		return true
	}
	entries := strings.Split(accept[0], ",")
	for _, entry := range entries {
		entry = strings.TrimSpace(entry)
		parts := strings.Split(entry, ";")
		mediaType := strings.TrimSpace(parts[0])
		if mediaType == "*/*" || mediaType == "application/*" || mediaType == "application/json" {
			return true
		}
	}
	return false
}

func parseJSONOrTextAcceptHeader(r *http.Request) (ContentType, bool) {
	accept, ok := r.Header["Accept"]
	if !ok {
		return ContentTypeJSON, true
	}
	entries := strings.Split(accept[0], ",")
	for _, entry := range entries {
		entry = strings.TrimSpace(entry)
		parts := strings.Split(entry, ";")
		mediaType := strings.TrimSpace(parts[0])
		if mediaType == "*/*" || mediaType == "application/*" || mediaType == "application/json" {
			return ContentTypeJSON, true
		}
		if mediaType == "text/plain" {
			return ContentTypePlainText, true
		}
	}
	return ContentTypeJSON, false
}

type ContentType = int

const (
	ContentTypeJSON ContentType = iota
	ContentTypePlainText
)
