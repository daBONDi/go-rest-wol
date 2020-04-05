//All Global Constants and Program Default Values
package main

//DefaultHTTPPort defines the Default HTTP Port to listen
const DefaultHTTPPort int = 8080

//DefaultComputerFilePath define the Default File Path to search for the Computer List
const DefaultComputerFilePath string = "computer.csv"

//DefaultComputerFilePathEnvironmentName define the Name of the Enviroment Variable where we should look for a Path
const DefaultComputerFilePathEnvironmentName string = "WOLFILE"

//DefaultHTTPPortEnvironmentVariableName define the Name of the Environment Variable what TCP Port we should use for the Webserver
const DefaultHTTPPortEnvironmentVariableName string = "WOLHTTPPORT"

//DefaultComputerUsernameEnvironmentName define the Name of the Enviroment Variable where we should look for a Username
const DefaultComputerUsernameEnvironmentName string = "WOLUSER"

//UsernameForShutdown define the Username on the Remote computer
const UsernameForShutdown string = "user"

//DefaultComputerPasswordEnvironmentName define the Name of the Enviroment Variable where we should look for a Username
const DefaultComputerPasswordEnvironmentName string = "WOLPWD"

//PasswordForShutdown define the Username on the Remote computer
const PasswordForShutdown string = "password"