package maps

import (
	"context"
	"time"
)

// RouteRequest contains parameters for calculating a route.
type RouteRequest struct {
	// Origin is the starting point.
	Origin LatLong
	// Destination is the end point.
	Destination LatLong
	// Waypoints are intermediate stops.
	Waypoints []LatLong
	// DepartureTime is the desired departure time. Zero value means "now".
	DepartureTime time.Time
}

// RouteLeg represents a single leg of a route between two waypoints.
type RouteLeg struct {
	// DistanceMeters is the leg distance in meters.
	DistanceMeters float64
	// Duration is the leg travel duration.
	Duration time.Duration
	// StartLocation is the leg origin coordinate.
	StartLocation LatLong
	// EndLocation is the leg destination coordinate.
	EndLocation LatLong
	// Polyline is the encoded polyline string for this leg.
	Polyline string
}

// RouteResult is the result of a route calculation.
type RouteResult struct {
	// DistanceMeters is the total route distance in meters.
	DistanceMeters float64
	// Duration is the total travel duration without traffic.
	Duration time.Duration
	// DurationInTraffic is the total travel duration accounting for traffic.
	DurationInTraffic time.Duration
	// Polyline is the encoded polyline for the full route.
	Polyline string
	// Legs is the list of route legs between consecutive waypoints.
	Legs []RouteLeg
	// Provider identifies the backend that produced this result.
	Provider string
}

// RouteMatrixRequest contains parameters for a distance matrix calculation.
type RouteMatrixRequest struct {
	// Origins is the list of origin coordinates.
	Origins []LatLong
	// Destinations is the list of destination coordinates.
	Destinations []LatLong
	// DepartureTime is the desired departure time. Zero value means "now".
	DepartureTime time.Time
}

// RouteMatrixElement represents a single origin-destination pair in a matrix result.
type RouteMatrixElement struct {
	// OriginIndex is the index into the origins list.
	OriginIndex int
	// DestinationIndex is the index into the destinations list.
	DestinationIndex int
	// DistanceMeters is the distance between the origin and destination.
	DistanceMeters float64
	// Duration is the travel time between the origin and destination.
	Duration time.Duration
	// Status is "ok" or an error description explaining why this element failed.
	Status string
}

// RouteMatrixResult is the result of a distance matrix calculation.
type RouteMatrixResult struct {
	// Elements is the flattened list of matrix entries.
	Elements []RouteMatrixElement
	// Provider identifies the backend that produced this result.
	Provider string
}

// Router calculates routes and distance matrices.
type Router interface {
	// Route calculates a route between an origin and destination.
	Route(ctx context.Context, req RouteRequest) (RouteResult, error)
	// RouteMatrix calculates a distance-duration matrix between origins and destinations.
	RouteMatrix(ctx context.Context, req RouteMatrixRequest) (RouteMatrixResult, error)
}
