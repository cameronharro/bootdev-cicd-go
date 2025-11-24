package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAuthHeader(t *testing.T) {
	type TestCase struct {
		name       string
		authHeader string
		wantErr    bool
		wantMatch  bool
		matchValue string
	}

	testCases := []TestCase{
		{
			name:       "Simple",
			authHeader: "ApiKey test123",
			wantErr:    false,
			wantMatch:  true,
			matchValue: "test123",
		},
		{
			name:       "One field",
			authHeader: "test123",
			wantErr:    true,
			wantMatch:  false,
		},
		{
			name:       "Bad Declaration",
			authHeader: "ApiKey: test123",
			wantErr:    true,
			wantMatch:  false,
		},
		{
			name:       "No Header",
			authHeader: "",
			wantErr:    true,
			wantMatch:  false,
		},
		{
			name:       "Extra fields",
			authHeader: "ApiKey one two",
			wantErr:    false,
			wantMatch:  true,
			matchValue: "one",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			headers := http.Header{}
			headers.Add("Authorization", testCase.authHeader)
			key, err := auth.GetAPIKey(headers)
			if (err != nil) != testCase.wantErr {
				t.Errorf("GetAPIKey() err: %v, wantErr: %t", err, testCase.wantErr)
			}
			if testCase.wantMatch && (key != testCase.matchValue) {
				t.Errorf("GetAPIKey() got: %v, expected: %v", key, testCase.matchValue)
			}
		})
	}
}
