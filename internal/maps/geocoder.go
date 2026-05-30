package maps

import "context"

// GeocodeResult is the result of a geocode or reverse-geocode operation.
type GeocodeResult struct {
	// Coordinates is the geographic position.
	Coordinates LatLong
	// NormalizedAddress is the canonical address from the provider.
	NormalizedAddress NormalizedAddress
	// Provider identifies the backend that produced this result (e.g. "google").
	Provider string
	// Confidence is a 0-1 score indicating how confident the provider is in this result.
	Confidence float64
}

// Geocoder geocodes addresses and reverse-geocodes coordinates.
type Geocoder interface {
	// Geocode converts an address into coordinates and a normalized address.
	Geocode(ctx context.Context, input AddressInput) (GeocodeResult, error)
	// ReverseGeocode converts coordinates into an approximate address.
	ReverseGeocode(ctx context.Context, latlng LatLong) (GeocodeResult, error)
}
