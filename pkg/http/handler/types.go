package handler

// WakeUpResponseObject Datastructure for holding information for the WakeUpResult
type WakeUpResponseObject struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	ErrorObject error  `json:"error"`
}
