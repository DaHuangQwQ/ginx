package openapi

// openAPIResponse describes a response error in the OpenAPI spec.
type openAPIResponse struct {
	Response
	Code        int
	Description string
}

// when setting custom response types on routes
type Response struct {
	// content-type of the response i.e application/json
	ContentTypes []string
	// user provided type
	Type any
}
