package env_test

import (
	"os"
	"testing"

	"github.com/edmarfelipe/go-ci/internal/env"
)

func TestLoad(t *testing.T) {
	testCases := []struct {
		name     string
		env      map[string]string
		expected *env.Env
	}{
		{
			name: "default values",
			env:  map[string]string{},
			expected: &env.Env{
				ServiceAddr: ":8080",
			},
		},
		{
			name: "custom values",
			env: map[string]string{
				"SERVICE_PORT": ":9090",
			},
			expected: &env.Env{
				ServiceAddr: ":9090",
			},
		},
	}

	for _, tc := range testCases {
		os.Clearenv()
		for k, v := range tc.env {
			os.Setenv(k, v)
		}

		env, err := env.Load()
		if err != nil {
			t.Fatalf("test %s: unexpected error: %v", tc.name, err)
		}

		if env.ServiceAddr != tc.expected.ServiceAddr {
			t.Errorf("test %s: expected ServiceAddr %q, got %q", tc.name, tc.expected.ServiceAddr, env.ServiceAddr)
		}
	}
}
