package main

import (
	"fmt"
	"os"

	"github.com/Cloud-Sable/InstantOpenVPN/pkg/aws"
	"github.com/Cloud-Sable/InstantOpenVPN/pkg/config"
	"github.com/Cloud-Sable/InstantOpenVPN/pkg/prompt"
)

func checkAWSEnv() (string, string, bool) {
	accessKey, akExists := os.LookupEnv("AWS_ACCESS_KEY_ID")
	secretKey, skExists := os.LookupEnv("AWS_SECRET_ACCESS_KEY")

	// Both keys must exist to return true
	return accessKey, secretKey, akExists && skExists
}

func main() {
	// Check and/or prompt for AWS credentials
	accessKey, secretKey, credsExist := checkAWSEnv()
	if !credsExist {
		accessKey, secretKey = prompt.PromptAWSCredentials()
	}

	// Initialize AWS session
	region := "us-west-2" // Default region, can be dynamically set or read from config
	sess := aws.InitAWSSession(accessKey, secretKey, region)

	// Read user preferences from config file
	prefs, err := config.ReadUserPrefs("userprefs.cfg")
	if err != nil {
		fmt.Println("No existing userprefs.cfg file or error reading it. Proceeding with defaults.")
		prefs = config.UserPrefs{InstanceType: "t3.micro", Region: region, Username: "defaultUser"} // Set default preferences
	}

	// Prompt for instance type, region, and username
	prefs.InstanceType = prompt.PromptForInstanceType(prefs.InstanceType)
	prefs.Region = prompt.PromptForRegion(prefs.Region)
	prefs.Username = prompt.PromptForUsername(prefs.Username)
	// prefs.IP, err = getPublicIP()

	// Write user preferences back to config file
	err = config.WriteUserPrefs("userprefs.cfg", prefs)
	if err != nil {
		fmt.Println("Error writing to userprefs.cfg:", err)
		os.Exit(1)
	}

	// Create SSH key pair
	keyPairName := fmt.Sprintf("%s-%s-keypair", prefs.Username, prefs.InstanceType)
	privateKey, err := aws.CreateSSHKeyPair(sess, keyPairName)
	if err != nil {
		fmt.Println("Error creating SSH key pair:", err)
		os.Exit(1)
	}

	// Save private key to file
	err = aws.SavePrivateKeyToFile(privateKey, keyPairName+".pem")
	if err != nil {
		fmt.Println("Error saving private key:", err)
		os.Exit(1)
	}

	// Launch EC2 instance with OpenVPN setup
	instanceID, err := aws.LaunchEC2Instance(sess, prefs.InstanceType, prefs.Region, keyPairName)
	if err != nil {
		fmt.Println("Error launching EC2 instance:", err)
		os.Exit(1)
	}

	fmt.Println("Launched EC2 Instance:", instanceID)

	// Additional OpenVPN configuration and client setup
	// This is where you'd use executeSSHCommand from openvpn.go
	// For example:
	// setupOpenVPN(instanceID, prefs.Username, keyPairName+".pem")
	// ...
}
