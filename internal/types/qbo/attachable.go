// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// AttachableResponse represents the QuickBooks attachable response envelope.
type AttachableResponse struct {
	Attachable Attachable `json:"Attachable"`
	Time       time.Time  `json:"time"`
}

// AttachableDeleteResponse represents the QuickBooks deleted attachable response envelope.
type AttachableDeleteResponse struct {
	Attachable DeletedEntity `json:"Attachable"`
	Time       time.Time     `json:"time"`
}

// AttachableQueryResponse represents a query response for attachables.
type AttachableQueryResponse struct {
	QueryResponse AttachableQueryResult `json:"QueryResponse"`
	Time          time.Time             `json:"time"`
}

// AttachableQueryResult represents attachable query results.
type AttachableQueryResult struct {
	Attachable    []Attachable `json:"Attachable,omitempty"`
	StartPosition int          `json:"startPosition,omitempty"`
	MaxResults    int          `json:"maxResults,omitempty"`
}

// AttachableUploadResponse represents upload results for attachables.
type AttachableUploadResponse struct {
	AttachableResponse []AttachableUploadResult `json:"AttachableResponse,omitempty"`
	Time               time.Time                `json:"time"`
}

// AttachableUploadResult represents a single uploaded attachable result.
type AttachableUploadResult struct {
	Attachable Attachable `json:"Attachable"`
}

// Attachable represents a QuickBooks attachable object.
type Attachable struct {
	ID                       string          `json:"Id"`
	SyncToken                string          `json:"SyncToken,omitempty"`
	FileName                 string          `json:"FileName,omitempty"`
	Note                     string          `json:"Note,omitempty"`
	FileAccessURI            string          `json:"FileAccessUri,omitempty"`
	Size                     float64         `json:"Size,omitempty"`
	ThumbnailFileAccessURI   string          `json:"ThumbnailFileAccessUri,omitempty"`
	TempDownloadURI          string          `json:"TempDownloadUri,omitempty"`
	ThumbnailTempDownloadURI string          `json:"ThumbnailTempDownloadUri,omitempty"`
	Category                 string          `json:"Category,omitempty"`
	ContentType              string          `json:"ContentType,omitempty"`
	PlaceName                string          `json:"PlaceName,omitempty"`
	AttachableRef            []AttachableRef `json:"AttachableRef,omitempty"`
	Long                     string          `json:"Long,omitempty"`
	Tag                      string          `json:"Tag,omitempty"`
	Lat                      string          `json:"Lat,omitempty"`
	MetaData                 *MetaData       `json:"MetaData,omitempty"`
	Domain                   string          `json:"domain,omitempty"`
	Sparse                   *bool           `json:"sparse,omitempty"`
}

// CreateAttachableRequest represents the documented create attachable payload.
type CreateAttachableRequest struct {
	Note          string          `json:"Note,omitempty"`
	FileName      string          `json:"FileName,omitempty"`
	AttachableRef []AttachableRef `json:"AttachableRef,omitempty"`
	Lat           string          `json:"Lat,omitempty"`
	Long          string          `json:"Long,omitempty"`
	Tag           string          `json:"Tag,omitempty"`
	PlaceName     string          `json:"PlaceName,omitempty"`
	ContentType   string          `json:"ContentType,omitempty"`
}

// AttachableRef represents a reference from an attachable to another object.
type AttachableRef struct {
	IncludeOnSend *bool                `json:"IncludeOnSend,omitempty"`
	LineInfo      string               `json:"LineInfo,omitempty"`
	NoRefOnly     *bool                `json:"NoRefOnly,omitempty"`
	CustomField   []CustomField        `json:"CustomField,omitempty"`
	Inactive      *bool                `json:"Inactive,omitempty"`
	EntityRef     *AttachableEntityRef `json:"EntityRef,omitempty"`
}

// AttachableEntityRef represents the linked object reference for an attachable.
type AttachableEntityRef struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
}
