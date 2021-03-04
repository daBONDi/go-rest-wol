package main

//WakeUpResponseObject Datastructure for holding information for the WakeUpResult
type WakeUpResponseObject struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	ErrorObject error  `json:"error"`
}

//PingResponseObject Datastructure for holding information for the PingResult
type PingResponseObject struct {
	IsOnLine    bool  `json:"isonline"`
	Message     string `json:"message"`
	ErrorObject error `json:"error"`
}

// Computer represents a Computer Object
type Computer struct {
	Name               string `csv:"name"`
	Mac                string `csv:"mac"`
	IPAddress          string `csv:"ip"`
	BroadcastIPAddress string `csv:"broadcast"`
}
