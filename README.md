# AWS OpenVPN Setup in Go

## This is a work in progress and is currently incomplete.  Use at your own risk!

## Project Overview

This project is a Go application that automates the setup of an OpenVPN server on AWS EC2. It handles AWS credentials, launches an EC2 instance, sets up OpenVPN, and generates client configurations.

## Features

- Checks for AWS credentials in environment variables or prompts the user for them.
- Reads and writes user preferences, including AWS region, instance type, and username, to a configuration file (`userprefs.cfg`).
- Launches an EC2 instance with Amazon Linux 2 and sets up OpenVPN using user data scripts.
- Generates and saves an SSH key pair for secure access to the EC2 instance.
- (Optional) Additional OpenVPN configuration and management tasks.

## Prerequisites

- Go 1.16 or later
- AWS account with necessary permissions (EC2, IAM)
- AWS CLI installed and configured (optional for environment variable setup)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-repository/InstantOpenVPN.git
    cd InstantOpenVPN
    ```

2. Build the application:

    ```bash
    go build cmd/openvpnaws/instantovpn.go
    ```

## Usage

1. Run the application:

    ```bash
    ./instantovpn
    ```

2. Follow the on-screen prompts to enter AWS credentials (if not set in environment variables), select AWS region, instance type, and username.

3. The application will create an EC2 instance with OpenVPN installed and configured. SSH key pair will be saved in the working directory.

4. Client configuration files and additional setup details can be managed as per project extension.

## Configuration File (`userprefs.cfg`)

The `userprefs.cfg` file stores user preferences. It's structured as key-value pairs:

instance_type=t3.micro
region=us-west-2
username=user


## Contributing

Contributions to this project are welcome! Please fork the repository and submit a pull request with your enhancements.

## License

MIT License

Copyright (c) [2024] [Cloud Sable]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
