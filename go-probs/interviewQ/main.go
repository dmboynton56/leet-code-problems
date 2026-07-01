package main

import "fmt"

/*
	You are tasked with managing configuration settings.

	Each configuration value can be one of the following simple Golang datatypes... bool, int, string.

	Each configuration value is associated with a key.  Keys are the simple Golang datatype string.

	For example, the "username" configuration key has a value of "Roger Rabbit".

	What would you use to store configuration settings so inserting, updating, and finding individual
	configuration settings is fast and simple?

	You will be storing the following settings...

	| Key      | Value            | Datatype |
	| -------- | ---------------  | -------- |
	| username | "Roger Rabbit"   | string   |
	| password | "IsBeingFramed!" | string   |
	| use_vpn  | true             | bool     |
	| timeout  | 17               | int      |

	Write code that stores those settings using your chosen method for managing configuration settings.
*/

func main() {
	config := map[string]any{
		"username": "Roger Rabbit",
		"password": "IsBeingFramed!",
		"use_vpn":  true,
		"timeout":  17,
	}
	for key, value := range config {
		fmt.Printf("%s: %v\n", key, value)
	}
}
