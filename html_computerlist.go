package main

import (
	"fmt"
)

// GetComputerListTable Return a HTML Table with the current Computers
func GetComputerListTable(List []Computer) string {
	var result string
	result = result + fmt.Sprintln("<table>")
	for _, c := range List {
		result = result + fmt.Sprintf("<tr><td><a href=\"/?computer=%s\">%v</a></td><td>%v</td><td>%v</td></tr>\n", c.Name, c.Name, c.Mac, c.BroadcastIPAddress)
	}
	return result + fmt.Sprintln("</table>")
}
