// Assignment 4: The Local Replace
//
// Goal: Create a local module my-logger. Use replace in go.mod
//       to redirect dependency resolution to the local folder.
//
// Instructions:
// 1. Create a folder: my-logger/
// 2. Inside my-logger: go mod init github.com/yourname/my-logger
// 3. Create my-logger/logger.go with a Log function
// 4. In main go.mod, add:
//    replace github.com/yourname/my-logger => ./my-logger
// 5. Import and use your local logger
//
// This is ESSENTIAL for microservices development!
//
// TODO: Implement your solution below

package main

func main() {
	// Your code here
	// Example: mylogger.Log("Hello from local module!")
}

