package apiv2

type RequestData struct {
	// setting index for data
	Index 		string 		`json:"index,omitempty"`

	// certain value to set
	//
	// required: true
	Value 		string 		`json:"value,omitempty"`

	// some certain notes
	Note		[]string    `json:"note,omitempty`
}