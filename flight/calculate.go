// Package flight takes a bunch of flights (including intermediate ones) and calculates resulting flight.
package flight

import (
	"context"
	"encore.dev/beta/errs"
	"fmt"
)

type CalculateParams struct {
	Flights [][]string `json:"flights"` // All flights including intermediate ones
}

type CalculateResponse struct {
	Response []string `json:"response"` // All flights including intermediate ones
}

// Calculate calculates resulting flight path (source->destination) from the input p.
// Returns error in case there is no single exact path to be constructed from the input p.
//
//encore:api public method=POST path=/calculate
func Calculate(_ context.Context, p CalculateParams) (*CalculateResponse, error) {
	// parse & validate inputs
	type Flight struct {
		Source      string
		Destination string
	}
	flights := make([]Flight, 0, len(p.Flights))
	for _, flight := range p.Flights {
		if len(flight) != 2 {
			return nil, &errs.Error{
				Code:    errs.InvalidArgument,
				Message: fmt.Sprintf("each flight must contain exactly two items"),
			}
		}
		if flight[0] == "" || flight[1] == "" {
			return nil, &errs.Error{
				Code:    errs.InvalidArgument,
				Message: fmt.Sprintf("flight source or destination must not be an empty string"),
			}
		}
		flights = append(flights, Flight{
			Source:      flight[0],
			Destination: flight[1],
		})
	}
	sources := make(map[string]struct{})      // a set of inferred flight sources
	destinations := make(map[string]struct{}) // a set of inferred flight destinations
	for _, flight := range flights {
		if _, ok := sources[flight.Source]; ok {
			return nil, &errs.Error{
				Code:    errs.InvalidArgument,
				Message: fmt.Sprintf("duplicate flight source: %s", flight.Source),
			}
		}
		sources[flight.Source] = struct{}{}
		if _, ok := destinations[flight.Destination]; ok {
			return nil, &errs.Error{
				Code:    errs.InvalidArgument,
				Message: fmt.Sprintf("duplicate flight destination: %s", flight.Destination),
			}
		}
		destinations[flight.Destination] = struct{}{}
	}

	// to validate that there is exactly one valid flight path provided we have to "execute" it below
	var (
		start, end           string // starting and final airport codes
		startFound, endFound bool
	)
	for _, flight := range flights {
		_, ok := sources[flight.Destination]
		if !ok {
			if endFound {
				return nil, &errs.Error{
					Code:    errs.InvalidArgument,
					Message: fmt.Sprintf("flight must have exactly 1 ending point, got at least 2: %s, %s", end, flight.Destination),
				}
			}
			endFound = true
			end = flight.Destination
		}
		delete(sources, flight.Destination) // means flight.Destination is not a final airport
		_, ok = destinations[flight.Source]
		if !ok {
			if startFound {
				return nil, &errs.Error{
					Code:    errs.InvalidArgument,
					Message: fmt.Sprintf("flight must have exactly 1 starting point, got at least 2: %s, %s", start, flight.Source),
				}
			}
			startFound = true
			start = flight.Source
		}
		delete(destinations, flight.Source) // means flight.Source is not a starting airport
	}
	if len(sources) != 1 {
		return nil, &errs.Error{
			Code:    errs.InvalidArgument,
			Message: fmt.Sprintf("flight must have exactly 1 ending point, got %d", len(sources)),
		}
	}
	if len(destinations) != 1 {
		return nil, &errs.Error{
			Code:    errs.InvalidArgument,
			Message: fmt.Sprintf("flight must have exactly 1 starting point, got %d", len(destinations)),
		}
	}

	return &CalculateResponse{
		Response: []string{start, end},
	}, nil
}
