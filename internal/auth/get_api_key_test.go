package auth

import (
	"cmp"
	"log"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		got  http.Header
		want string
	}{
		"Ideal Case": {
			got:  http.Header{"Authorization": {"ApiKey 12345"}},
			want: "12345",
		},
		"Missing Key": {
			got:  http.Header{},
			want: "1234",
		},
	}

	for name, tc := range tests {
		key, _ := GetAPIKey(tc.got)
		if cmp.Compare(key, tc.want) != 0 {
			log.Fatalf("Test Case : %v failed. Expected %#v but got %#v", name, tc.want, key)
		}
	}
}
