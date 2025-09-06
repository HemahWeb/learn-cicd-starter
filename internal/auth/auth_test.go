package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input   http.Header
		want    string
		wantErr bool
	}{
		"valid api key": {
			input:   http.Header{"Authorization": []string{"ApiKey test-key"}},
			want:    "test-key",
			wantErr: false,
		},
		"no auth header": {
			input:   http.Header{},
			want:    "",
			wantErr: true,
		},
		"malformed header": {
			input:   http.Header{"Authorization": []string{"Bearer token"}},
			want:    "",
			wantErr: true,
		},
		"empty auth header": {
			input:   http.Header{"Authorization": []string{""}},
			want:    "",
			wantErr: true,
		},
		"incomplete api key": {
			input:   http.Header{"Authorization": []string{"ApiKey"}},
			want:    "",
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)

			// Check error expectation
			if tc.wantErr && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Check return value
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
