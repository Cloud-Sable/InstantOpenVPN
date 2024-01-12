# OpenVPN AWS Setup

This project is a Go application that automates the setup of an OpenVPN server on an AWS EC2 instance.

## Description

The application performs the following steps:

1. Gets the user's public IP.
2. Writes user preferences to a configuration file.
3. Creates an SSH key pair.
4. Saves the private key to a file.
5. Creates a security group.
6. Launches an EC2 instance with OpenVPN setup.
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
    go build -o InstantOpenVPN cmd/openvpnaws/main.go
    ```

## Usage

1. Run the application:

    ```bash
    ./InstantOpenVPN
    ```

2. Follow the on-screen prompts to enter AWS credentials (if not set in environment variables), select AWS region, instance type, and username.

3. The application will create an EC2 instance with OpenVPN installed and configured. SSH key pair will be saved in the working directory.

4. Import-ready client configuration files will be downloaded to the current working directory.

## Configuration File (`userprefs.cfg`)

The `userprefs.cfg` file stores user preferences. It's structured as key-value pairs:
```
instance_type=t3.micro
region=us-west-2
username=user
ip=1.2.3.4  
```
(ip is the user's public IP to be added to the instance allowlist)
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
