package utils

import (
	"api-gateway/pkg/constant"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient is an interface for making HTTP requests
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	// DefaultClient is the default HTTP client
	DefaultClient HTTPClient = &http.Client{
		Timeout: 30 * time.Second,
	}
)

// ForwardRequest forwards an HTTP request to the specified URL
func ForwardRequest(ctx context.Context, method, url string, headers http.Header, body []byte) (*http.Response, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Copy headers
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Send request
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return resp, respBody, nil
}

// MakeRequest is a generic function to make HTTP requests
func MakeRequest(ctx context.Context, method, url string, headers map[string]string, requestBody interface{}, responseBody interface{}) error {
	var bodyBytes []byte
	var err error

	// Marshal request body if it exists
	if requestBody != nil {
		bodyBytes, err = json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set(constant.HeaderContentType, constant.ContentTypeJSON)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send request
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Unmarshal response if a response structure was provided
	if responseBody != nil {
		if err := json.NewDecoder(resp.Body).Decode(responseBody); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// ExtractRequestID extracts the request ID from headers
func ExtractRequestID(headers http.Header) string {
	requestID := headers.Get(constant.HeaderXRequestID)
	if requestID == "" {
		requestID = headers.Get(constant.HeaderXCorrelationID)
	}
	return requestID
}

// GetContentType returns the content type without charset
func GetContentType(header http.Header) string {
	contentType := header.Get(constant.HeaderContentType)
	for i, char := range contentType {
		if char == ';' {
			return contentType[:i]
		}
	}
	return contentType
}

// SetRequestID sets a request ID in the headers if one doesn't exist
func SetRequestID(headers http.Header, requestID string) {
	if headers.Get(constant.HeaderXRequestID) == "" {
		headers.Set(constant.HeaderXRequestID, requestID)
	}
}

// IsJSONResponse checks if the response is JSON
func IsJSONResponse(header http.Header) bool {
	return GetContentType(header) == constant.ContentTypeJSON
}

// CreateClientWithTimeout creates an HTTP client with timeout
func CreateClientWithTimeout(timeout time.Duration) HTTPClient {
	return &http.Client{
		Timeout: timeout,
	}
}
