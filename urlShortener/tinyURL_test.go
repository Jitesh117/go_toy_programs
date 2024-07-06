package main

import (
	"net/url"
	"strings"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	longURL := "https://www.example.com"
	shortURL := generateShortURL(longURL)

	if !strings.HasPrefix(shortURL, baseURL) {
		t.Errorf("Expected short URL to start with %s, got %s", baseURL, shortURL)
	}

	if len(shortURL) != len(baseURL)+shortLength {
		t.Errorf("Expected short URL length to be %d, got %d", len(baseURL)+shortLength, len(shortURL))
	}
}

func TestGenerateShortURLUniqueness(t *testing.T) {
	longURL1 := "https://www.example.com/page1"
	longURL2 := "https://www.example.com/page2"

	shortURL1 := generateShortURL(longURL1)
	shortURL2 := generateShortURL(longURL2)

	if shortURL1 == shortURL2 {
		t.Errorf("Expected different short URLs for different long URLs, got %s and %s", shortURL1, shortURL2)
	}
}

func TestGenerateShortURLEmpty(t *testing.T) {
	longURL := ""
	shortURL := generateShortURL(longURL)

	if !strings.HasPrefix(shortURL, baseURL) {
		t.Errorf("Expected short URL to start with %s, got %s", baseURL, shortURL)
	}

	if len(shortURL) != len(baseURL)+shortLength {
		t.Errorf("Expected short URL length to be %d, got %d", len(baseURL)+shortLength, len(shortURL))
	}
}

func TestInvalidURL(t *testing.T) {
	invalidURL := "htp:/invalid-url"
	_, err := url.ParseRequestURI(invalidURL)

	if err == nil {
		t.Errorf("Expected error for invalid URL, got nil")
	}
}

func TestValidURL(t *testing.T) {
	validURL := "https://www.valid-url.com"
	parsedURL, err := url.ParseRequestURI(validURL)

	if err != nil {
		t.Errorf("Expected no error for valid URL, got %v", err)
	}

	if parsedURL.String() != validURL {
		t.Errorf("Expected parsed URL to be %s, got %s", validURL, parsedURL.String())
	}
}

func TestShortURLFormat(t *testing.T) {
	longURL := "https://www.example.com"
	shortURL := generateShortURL(longURL)

	if !strings.HasPrefix(shortURL, baseURL) {
		t.Errorf("Expected short URL to start with %s, got %s", baseURL, shortURL)
	}

	shortCode := strings.TrimPrefix(shortURL, baseURL)
	if len(shortCode) != shortLength {
		t.Errorf("Expected short code length to be %d, got %d", shortLength, len(shortCode))
	}
}

func TestURLWithQuery(t *testing.T) {
	longURL := "https://www.example.com?page=1&sort=asc"
	shortURL := generateShortURL(longURL)

	if !strings.HasPrefix(shortURL, baseURL) {
		t.Errorf("Expected short URL to start with %s, got %s", baseURL, shortURL)
	}

	if len(shortURL) != len(baseURL)+shortLength {
		t.Errorf("Expected short URL length to be %d, got %d", len(baseURL)+shortLength, len(shortURL))
	}
}
