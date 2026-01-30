package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/StephenCotterrell/pokedexcli/internal/pokecache"
)

func TestListLocationsUsesCache(t *testing.T) {
	var hits int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		if r.URL.Path != "/api/v2/location-area" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"count":1,"next":null,"previous":null,"results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"}]}`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client := Client{
		httpClient: http.Client{Timeout: 2 * time.Second},
		cache:      pokecache.NewCache(time.Minute),
	}
	pageURL := server.URL + "/api/v2/location-area"

	if _, err := client.ListLocations(&pageURL); err != nil {
		t.Fatalf("first ListLocations failed: %v", err)
	}

	if _, err := client.ListLocations(&pageURL); err != nil {
		t.Fatalf("second ListLocations failed: %v", err)
	}

	if got := atomic.LoadInt32(&hits); got != 1 {
		t.Fatalf("expected 1 request, got %d", got)
	}
}
