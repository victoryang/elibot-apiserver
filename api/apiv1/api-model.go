package apiv1

type RequestData struct {
	// setting index for data
	Index 		int 		`json:"index,omitempty"`

	// certain value to set
	//
	// required: true
	Value 		string 		`json:"value,omitempty"`

	Values 		[]string	`json:"values,omitempty"`
}

type RequestDataForIP struct {
	Address 		string 			`json:address, omitempty`
	Netmask 		string 			`json:netmask, omitempty`
	Network 		string 			`json:network, omitempty`
	Broadcast 		string 			`json:broadcast, omitempty`
	Gateway 		string 			`json:gateway, omitempty`
}

type RequestDataForCommandArgs struct {
	Args 		[]string		`json:args, omitempty`
}