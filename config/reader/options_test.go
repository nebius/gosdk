package reader_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nebius/gosdk/config"
	"github.com/nebius/gosdk/config/reader"
)

//nolint:paralleltest // Subtests share process environment variables.
func TestEnvironmentOverrides(t *testing.T) {
	t.Setenv("NEBIUS_PROFILE", "environment")
	t.Setenv("NEBIUS_IAM_TOKEN", "environment-token")
	t.Setenv("NEBIUS_ENDPOINT", "environment.example.test:443")
	t.Setenv("CUSTOM_PROFILE", "custom")
	t.Setenv("CUSTOM_TOKEN", "custom-token")
	t.Setenv("CUSTOM_ENDPOINT", "custom.example.test:443")

	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("profile-token"), 0600))

	testCases := []struct {
		name         string
		options      []config.Option
		wantProfile  string
		wantEndpoint string
		wantToken    string
	}{
		{
			name:         "default environment variables",
			wantProfile:  "environment",
			wantEndpoint: "environment.example.test:443",
			wantToken:    "environment-token",
		},
		{
			name:         "without environment variables",
			options:      []config.Option{reader.WithoutEnvVars()},
			wantProfile:  "profile",
			wantEndpoint: "profile.example.test:443",
			wantToken:    "profile-token",
		},
		{
			name: "custom environment variables",
			options: []config.Option{
				reader.WithProfileEnv("CUSTOM_PROFILE"),
				reader.WithTokenEnv("CUSTOM_TOKEN"),
				reader.WithEndpointEnv("CUSTOM_ENDPOINT"),
			},
			wantProfile:  "custom",
			wantEndpoint: "custom.example.test:443",
			wantToken:    "custom-token",
		},
		{
			name: "without custom environment variables",
			options: []config.Option{
				reader.WithProfileEnv("CUSTOM_PROFILE"),
				reader.WithTokenEnv("CUSTOM_TOKEN"),
				reader.WithEndpointEnv("CUSTOM_ENDPOINT"),
				reader.WithoutEnvVars(),
			},
			wantProfile:  "profile",
			wantEndpoint: "profile.example.test:443",
			wantToken:    "profile-token",
		},
		{
			name: "explicit profile and endpoint remain enabled",
			options: []config.Option{
				reader.WithProfileName("explicit"),
				reader.WithCustomEndpoint("explicit.example.test:443"),
				reader.WithoutEnvVars(),
			},
			wantProfile:  "explicit",
			wantEndpoint: "explicit.example.test:443",
			wantToken:    "profile-token",
		},
		{
			name: "later options enable environment variables",
			options: []config.Option{
				reader.WithoutEnvVars(),
				reader.WithProfileEnv("CUSTOM_PROFILE"),
				reader.WithTokenEnv("CUSTOM_TOKEN"),
				reader.WithEndpointEnv("CUSTOM_ENDPOINT"),
			},
			wantProfile:  "custom",
			wantEndpoint: "custom.example.test:443",
			wantToken:    "custom-token",
		},
		{
			name: "disable only profile and token overrides",
			options: []config.Option{
				reader.WithProfileEnv(""),
				reader.WithTokenEnv(""),
			},
			wantProfile:  "profile",
			wantEndpoint: "environment.example.test:443",
			wantToken:    "profile-token",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			profiles := make(config.ProfilesConfig)
			for _, name := range []string{"profile", "environment", "custom", "explicit"} {
				profiles[name] = &config.Profile{
					Name:      name,
					Endpoint:  "profile.example.test:443",
					TokenFile: tokenPath,
				}
			}
			options := []config.Option{
				reader.WithPreloadedConfig(&config.Config{
					Default:  "profile",
					Profiles: profiles,
				}),
			}
			options = append(options, tc.options...)
			cfg := reader.NewConfigReader(options...)
			require.NoError(t, cfg.LoadIfNeeded(t.Context()))
			require.Equal(t, tc.wantProfile, cfg.CurrentProfileName())
			require.Equal(t, tc.wantEndpoint, cfg.Endpoint())
			tokener, err := cfg.GetCredentials(t.Context())
			require.NoError(t, err)
			token, err := tokener.BearerToken(t.Context())
			require.NoError(t, err)
			require.Equal(t, tc.wantToken, token.Token)
		})
	}
}
