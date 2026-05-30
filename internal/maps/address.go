package maps

// LatLong represents a geographic coordinate pair.
type LatLong struct {
	// Lat is the latitude in decimal degrees.
	Lat float64
	// Longitude is the longitude in decimal degrees.
	Longitude float64
}

// AddressInput represents a free-form address to be geocoded.
type AddressInput struct {
	// Street is the street address line (e.g. "123 Main St").
	Street string
	// City is the city name.
	City string
	// State is the state or region code (e.g. "CA").
	State string
	// PostalCode is the postal or ZIP code.
	PostalCode string
	// Country is the full country name (e.g. "United States").
	Country string
}

// NormalizedAddress represents a provider-canonicalized address.
type NormalizedAddress struct {
	// FormattedAddress is the full address formatted as a single string.
	FormattedAddress string
	// Street is the street address line.
	Street string
	// City is the city name.
	City string
	// State is the state or region code.
	State string
	// PostalCode is the postal or ZIP code.
	PostalCode string
	// Country is the full country name.
	Country string
	// CountryCode is the ISO 3166-1 alpha-2 country code (e.g. "US").
	CountryCode string
}
