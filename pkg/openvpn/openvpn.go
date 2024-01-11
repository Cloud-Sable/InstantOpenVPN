package openvpn

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

// executeSSHCommand executes a given command on the remote server via SSH.
func executeSSHCommand(instanceIP, privateKeyPath, command string) (string, error) {
	key, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return "", fmt.Errorf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "ec2-user",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Note: This is insecure; for production, use a proper host key callback
	}

	conn, err := ssh.Dial("tcp", net.JoinHostPort(instanceIP, "22"), config)
	if err != nil {
		return "", fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	return string(output), err
}
