package apiv1

type RequestData struct {
	// setting index for data
	Index 		string 		`json:"index,omitempty"`

	// certain value to set
	//
	// required: true
	Value 		string 		`json:"value,omitempty"`

	Values 		[]string	`json:"values,omitempty"`

	// some certain notes
	Note		[]string    `json:"note,omitempty`
}