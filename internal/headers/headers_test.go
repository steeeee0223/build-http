package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderParse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)

	h1, ok := headers.Get("Host")
	assert.True(t, ok)
	assert.Equal(t, "localhost:42069", h1)

	h2, ok := headers.Get("MissingKey")
	assert.False(t, ok)
	assert.Equal(t, "", h2)

	assert.Equal(t, 25, n)
	assert.True(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("       Host : localhost:42069       \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Invalid character in header name
	headers = NewHeaders()
	data = []byte("H©st: localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Header with multiple values
	headers = NewHeaders()
	data = []byte("Host: localhost:42069\r\nHost: example.com\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	assert.NotNil(t, headers)
	h3, ok := headers.Get("Host")
	assert.True(t, ok)
	assert.Equal(t, "localhost:42069,example.com", h3)
	assert.True(t, done)
}
