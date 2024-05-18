package main

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetFavicon(t *testing.T) {
	t.Run("Favicon found in Redis", func(t *testing.T) {
		server := miniredis.RunT(t)
		defer server.Close()

		// Set the Redis address
		os.Setenv("FAVICON_REDIS_SERVER", server.Addr())
		rdb := redis.NewClient(&redis.Options{Addr: server.Addr()})

		// Mock the storage of a favicon
		err := rdb.Set(ctx, "example.com", "FAKE_FAVICON_CONTENT", 0).Err()
		if err != nil {
			t.Fatalf("Failed to set favicon in Redis: %v", err)
		}

		// Create a new HTTP request
		req, err := http.NewRequest("GET", "/images/example.com", nil)
		if err != nil {
			t.Fatalf("Failed to create HTTP request: %v", err)
		}

		// Create a new HTTP response recorder
		rr := httptest.NewRecorder()

		// Call the GetFavicon handler function
		GetFavicon(rr, req)

		// Verify that the favicon was returned in the HTTP response
		assert.Equal(t, "FAKE_FAVICON_CONTENT", rr.Body.String())
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Favicon not found in Redis", func(t *testing.T) {
		// Create a new miniredis server
		server := miniredis.RunT(t)
		defer server.Close()

		os.Setenv("FAVICON_REDIS_SERVER", server.Addr())

		// Create a new HTTP request
		req, err := http.NewRequest("GET", "/images/example.com", nil)
		if err != nil {
			t.Fatalf("Failed to create HTTP request: %v", err)
		}

		// Create a new HTTP response recorder
		rr := httptest.NewRecorder()

		// Call the GetFavicon handler function
		GetFavicon(rr, req)

		// Verify that a 404 Not Found response was returned
		assert.Equal(t, "Image not found", rr.Body.String())
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
