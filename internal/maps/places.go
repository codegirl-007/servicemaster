package maps

import "context"

// Place represents detailed information about a place.
type Place struct {
	// ID is the provider-specific place identifier.
	ID string
	// Name is the human-readable place name.
	Name string
	// FormattedAddress is the full address as a single string.
	FormattedAddress string
	// Coordinates is the geographic location of the place.
	Coordinates LatLong
	// PhoneNumber is the place's contact phone number.
	PhoneNumber string
	// Website is the place's website URL.
	Website string
	// Rating is the average user rating on a 0-5 scale.
	Rating float64
	// Types are the place categories (e.g. "restaurant", "plumber").
	Types []string
	// Provider identifies the backend that returned this place.
	Provider string
}

// PlacePrediction represents an autocomplete prediction from a place search.
type PlacePrediction struct {
	// PlaceID is the provider-specific identifier for the predicted place.
	PlaceID string
	// Description is the full human-readable description of the prediction.
	Description string
	// MainText is the primary text of the prediction (e.g. business name).
	MainText string
	// SecondaryText is the secondary text (e.g. city, state).
	SecondaryText string
	// Types are the predicted place categories.
	Types []string
}

// PlacesSearcher searches for places and retrieves place details.
type PlacesSearcher interface {
	// Search returns autocomplete predictions for the given query string.
	Search(ctx context.Context, query string) ([]PlacePrediction, error)
	// PlaceDetails returns detailed information for a specific place ID.
	PlaceDetails(ctx context.Context, placeID string) (Place, error)
}
