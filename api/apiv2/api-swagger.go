package apiv2

// swagger:model Request
type RequestTemplate struct {
	// setting index for data
	Index 		string 		`json:"index,omitempty"`

	// certain value to set
	//
	// required: true
	Value 		string 		`json:"value,omitempty"`

	// some certain notes
	Note		[]string    `json:"note,omitempty`
}

// swagger:response Response
type ResponseTemplate struct {
	// response body
	// in: body
	Body struct {
		// returned message
		//
		// Required: true
		Msg 		string
	}
}