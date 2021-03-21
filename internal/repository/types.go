package repository

// Computer represents a Computer Object
type Computer struct {
	Name               string `csv:"name"`
	Mac                string `csv:"mac"`
	BroadcastIPAddress string `csv:"ip"`
}

// ComputerList contains all Computers who we can use to work with
var ComputerList []Computer
