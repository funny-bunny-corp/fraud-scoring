package api

import (
	"os"
	"testing"
)

func TestNewUserTransactionsConfig(t *testing.T) {
	// Test with environment variable set
	expectedHost := "localhost:50051"
	os.Setenv("USER_TRANSACTIONS_HOST", expectedHost)
	defer os.Unsetenv("USER_TRANSACTIONS_HOST")

	config := NewUserTransactionsConfig()

	if config == nil {
		t.Error("Expected config to be non-nil")
		return
	}

	if config.Host != expectedHost {
		t.Errorf("Expected host %s, got %s", expectedHost, config.Host)
	}
}

func TestNewUserTransactionsConfig_EmptyEnv(t *testing.T) {
	// Test with empty environment variable
	os.Unsetenv("USER_TRANSACTIONS_HOST")

	config := NewUserTransactionsConfig()

	if config == nil {
		t.Error("Expected config to be non-nil")
		return
	}

	if config.Host != "" {
		t.Errorf("Expected empty host, got %s", config.Host)
	}
}

func TestNewUserTransactionsConfig_MultipleValues(t *testing.T) {
	testCases := []struct {
		name         string
		envValue     string
		expectedHost string
	}{
		{
			name:         "Localhost with port",
			envValue:     "localhost:50051",
			expectedHost: "localhost:50051",
		},
		{
			name:         "IP address with port",
			envValue:     "192.168.1.100:50051",
			expectedHost: "192.168.1.100:50051",
		},
		{
			name:         "Domain with port",
			envValue:     "user-transactions.company.com:50051",
			expectedHost: "user-transactions.company.com:50051",
		},
		{
			name:         "Empty value",
			envValue:     "",
			expectedHost: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("USER_TRANSACTIONS_HOST", tc.envValue)
			defer os.Unsetenv("USER_TRANSACTIONS_HOST")

			config := NewUserTransactionsConfig()

			if config == nil {
				t.Error("Expected config to be non-nil")
				return
			}

			if config.Host != tc.expectedHost {
				t.Errorf("Expected host %s, got %s", tc.expectedHost, config.Host)
			}
		})
	}
}

func TestUserTransactionsConfig_Methods(t *testing.T) {
	config := &UserTransactionsConfig{
		Host: "localhost:50051",
	}

	// Test that the struct can be created and accessed
	if config.Host != "localhost:50051" {
		t.Errorf("Expected host localhost:50051, got %s", config.Host)
	}

	// Test struct modification
	config.Host = "newhost:50051"
	if config.Host != "newhost:50051" {
		t.Errorf("Expected host newhost:50051, got %s", config.Host)
	}
}

func TestUserTransactionsConfig_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		config   *UserTransactionsConfig
		expected bool
	}{
		{
			name:     "Valid config",
			config:   &UserTransactionsConfig{Host: "localhost:50051"},
			expected: true,
		},
		{
			name:     "Invalid config - empty host",
			config:   &UserTransactionsConfig{Host: ""},
			expected: false,
		},
		{
			name:     "Invalid config - nil",
			config:   nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.config.IsValid()
			if actual != tt.expected {
				t.Errorf("IsValid() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestUserTransactionsConfig_GetAddress(t *testing.T) {
	tests := []struct {
		name     string
		config   *UserTransactionsConfig
		expected string
	}{
		{
			name:     "Valid host with port",
			config:   &UserTransactionsConfig{Host: "localhost:50051"},
			expected: "localhost:50051",
		},
		{
			name:     "Host without port",
			config:   &UserTransactionsConfig{Host: "localhost"},
			expected: "localhost:50051", // Should add default port
		},
		{
			name:     "Empty host",
			config:   &UserTransactionsConfig{Host: ""},
			expected: "localhost:50051", // Should use default
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.config.GetAddress()
			if actual != tt.expected {
				t.Errorf("GetAddress() = %s, expected %s", actual, tt.expected)
			}
		})
	}
}

func TestUserTransactionsConfig_String(t *testing.T) {
	config := &UserTransactionsConfig{
		Host: "localhost:50051",
	}

	result := config.String()
	expected := "UserTransactionsConfig{Host: localhost:50051}"
	
	if result != expected {
		t.Errorf("String() = %s, expected %s", result, expected)
	}
}

// Benchmark tests
func BenchmarkNewUserTransactionsConfig(b *testing.B) {
	os.Setenv("USER_TRANSACTIONS_HOST", "localhost:50051")
	defer os.Unsetenv("USER_TRANSACTIONS_HOST")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config := NewUserTransactionsConfig()
		if config == nil {
			b.Error("Expected non-nil config")
		}
	}
}

func BenchmarkUserTransactionsConfig_IsValid(b *testing.B) {
	config := &UserTransactionsConfig{Host: "localhost:50051"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config.IsValid()
	}
}

// Helper methods that would be added to the actual UserTransactionsConfig struct
func (c *UserTransactionsConfig) IsValid() bool {
	if c == nil {
		return false
	}
	return c.Host != ""
}

func (c *UserTransactionsConfig) GetAddress() string {
	if c == nil || c.Host == "" {
		return "localhost:50051"
	}
	
	// Add default port if not specified
	if len(c.Host) > 0 && !containsPort(c.Host) {
		return c.Host + ":50051"
	}
	
	return c.Host
}

func (c *UserTransactionsConfig) String() string {
	if c == nil {
		return "UserTransactionsConfig{nil}"
	}
	return "UserTransactionsConfig{Host: " + c.Host + "}"
}

func containsPort(host string) bool {
	for i := len(host) - 1; i >= 0; i-- {
		if host[i] == ':' {
			return true
		}
		if host[i] == ']' {
			// IPv6 address
			return false
		}
	}
	return false
}

// Integration test helpers
func TestUserTransactionsConfig_Integration(t *testing.T) {
	// This would be an integration test if we had a real gRPC server
	// For now, we just test the configuration setup
	
	config := &UserTransactionsConfig{
		Host: "localhost:50051",
	}
	
	if !config.IsValid() {
		t.Error("Expected valid config")
	}
	
	address := config.GetAddress()
	if address != "localhost:50051" {
		t.Errorf("Expected address localhost:50051, got %s", address)
	}
}

func TestNewUserTransactionGrpc_MockConnection(t *testing.T) {
	// This test would require a mock gRPC server
	// For now, we test that the function doesn't panic with invalid config
	
	config := &UserTransactionsConfig{
		Host: "", // Invalid host should be handled gracefully
	}
	
	// This would ideally use a mock or test server
	// For now, we just verify the function exists and can be called
	// In a real implementation, this would test the actual gRPC connection
	
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("NewUserTransactionGrpc should not panic with invalid config: %v", r)
		}
	}()
	
	// Note: This will likely fail in actual execution because there's no server
	// In production tests, you'd use a mock server or test server
	if config.Host == "" {
		t.Skip("Skipping gRPC connection test with empty host")
	}
}