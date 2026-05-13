package main

import (
	"os"
	"strings"
	"time"
)

var defaultAllowedOrigins = []string{"http://localhost:3001"}

func getAllowedOrigins() []string {
	raw := strings.TrimSpace(os.Getenv("ALLOWED_ORIGINS"))
	if raw == "" {
		return defaultAllowedOrigins
	}

	origins := make([]string, 0)
	for _, entry := range strings.Split(raw, ",") {
		origin := strings.TrimSpace(entry)
		if origin != "" {
			origins = append(origins, origin)
		}
	}
	if len(origins) == 0 {
		return defaultAllowedOrigins
	}

	return origins
}

// getPositiveTimeout returns the configured timeout in seconds, but ensures a
// sensible default if the provided value is non-positive.
func getPositiveTimeout(envKey string, defaultSeconds int) time.Duration {
	seconds := getEnvAsInt(envKey, defaultSeconds)
	if seconds <= 0 {
		seconds = defaultSeconds
	}
	return time.Duration(seconds) * time.Second
}

// Timeout helpers (configurable via env vars)
func getRequestTimeout() time.Duration  { return getPositiveTimeout("REQUEST_TIMEOUT_SECONDS", 60) }
func getAITimeout() time.Duration       { return getPositiveTimeout("AI_REQUEST_TIMEOUT_SECONDS", 30) }
func getVerifierTimeout() time.Duration { return getPositiveTimeout("VERIFIER_TIMEOUT_SECONDS", 2) }
func getHealthCheckTimeout() time.Duration {
	return getPositiveTimeout("HEALTH_CHECK_TIMEOUT_SECONDS", 2)
}
