package config

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type UserPrefs struct {
	InstanceType string
	Region       string
	Username     string
	IP           string
}

func GetPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

// readUserPrefs reads user preferences from the specified file.
func ReadUserPrefs(filePath string) (UserPrefs, error) {
	prefs := UserPrefs{}
	file, err := os.Open(filePath)
	if err != nil {
		return prefs, err // File not found or other errors
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

			switch key {
			case "instance_type":
				prefs.InstanceType = value
			case "region":
				prefs.Region = value
			case "username":
				prefs.Username = value
			}
		}
	}

	return prefs, scanner.Err()
}

// writeUserPrefs writes the user preferences to the specified file.
func WriteUserPrefs(filePath string, prefs UserPrefs) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("instance_type=%s\n", prefs.InstanceType))
	if err != nil {
		return err
	}
	_, err = writer.WriteString(fmt.Sprintf("region=%s\n", prefs.Region))
	if err != nil {
		return err
	}
	_, err = writer.WriteString(fmt.Sprintf("username=%s\n", prefs.Username))
	if err != nil {
		return err
	}
	_, err = writer.WriteString(fmt.Sprintf("ip=%s\n", prefs.IP)) // Add this block
	if err != nil {
		return err
	}

	return writer.Flush()
}
