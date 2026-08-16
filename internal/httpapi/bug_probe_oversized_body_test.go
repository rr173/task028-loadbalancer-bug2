package httpapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOversizedJSONBodyIsRejected(t *testing.T) {
	body := `{"id":"a","address":"a:80","weight":1}` + strings.Repeat(" ", 1<<20) + "x"
	req := httptest.NewRequest(http.MethodPost, "/nodes", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	New().Handler().ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("oversized body status=%d, want %d; body=%s", rr.Code, http.StatusBadRequest, fmt.Sprint(rr.Body.String()))
	}
}
