package main

// Computer represents a Computer Object for WOL
type Computer struct {
	Name               string `csv:"name"`
	Mac                string `csv:"mac"`
	BroadcastIPAddress string `csv:"ip"`
}
