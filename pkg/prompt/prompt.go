package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptAWSCredentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter AWS Access Key ID: ")
	accessKey, _ := reader.ReadString('\n')
	accessKey = strings.TrimSpace(accessKey)

	fmt.Print("Enter AWS Secret Access Key: ")
	secretKey, _ := reader.ReadString('\n')
	secretKey = strings.TrimSpace(secretKey)

	return accessKey, secretKey
}

// PromptForInstanceType prompts the user to input the instance type.
func PromptForInstanceType(defaultType string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter instance type [default: %s]: ", defaultType)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultType
	}
	return input
}

// PromptForRegion prompts the user to input the AWS region.
func PromptForRegion(defaultRegion string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter AWS region [default: %s]: ", defaultRegion)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultRegion
	}
	return input
}

// PromptForUsernames prompts the user to input a username or comma-separated usernames.
func PromptForUsername(defaultUsername string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter username [default: %s]: ", defaultUsername)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultUsername
	}
	return input
}
