package flight

import (
	"context"
	"testing"
)

// Run tests using `encore test`, which compiles the Encore app and then runs `go test`.
// It supports all the same flags that the `go test` command does.
// You automatically get tracing for tests in the local dev dash: http://localhost:9400
// Learn more: https://encore.dev/docs/develop/testing
func TestCalculate(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		params := CalculateParams{
			Flights: [][]string{
				{"IND", "EWR"},
				{"SFO", "ALT"},
				{"GSO", "IND"},
				{"ALT", "GSO"},
			},
		}
		want := []string{"SFO", "EWR"}
		resp, err := Calculate(ctx, params)
		if err != nil {
			t.Fatal(err)
		}
		if resp.Response[0] != want[0] || resp.Response[1] != want[1] {
			t.Errorf("got %q, want %q", resp.Response, want)
		}
	})
	t.Run("failure", func(t *testing.T) {
		params := CalculateParams{
			Flights: [][]string{
				{"IND", "EWR"},
				{"SFO", "ALT"},
				{"GSO", "***"},
				{"ALT", "GSO"},
			},
		}
		_, err := Calculate(ctx, params)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}
