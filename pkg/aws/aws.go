package aws

import (
	"embed"
	"encoding/base64"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Embed the user data script
//
//go:embed ovpn_userdata.sh
var userDataFS embed.FS

// initAWSSession initializes an AWS session with the given credentials.
func InitAWSSession(accessKey, secretKey, region string) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		fmt.Println("Error initializing AWS session:", err)
		os.Exit(1)
	}
	return sess
}

func CreateSecurityGroupAndRule(sess *session.Session, groupName, description, ip string) (string, error) {
	svc := ec2.New(sess)

	// Create security group
	createRes, err := svc.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(groupName),
		Description: aws.String(description),
	})
	if err != nil {
		return "", err
	}

	// Get the ID of the security group
	groupId := *createRes.GroupId

	// Create security group rule
	_, err = svc.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(groupId),
		IpPermissions: []*ec2.IpPermission{
			{
				IpProtocol: aws.String("tcp"),
				FromPort:   aws.Int64(22),
				ToPort:     aws.Int64(22),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp:      aws.String(ip + "/32"),
						Description: aws.String("SSH access"),
					},
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	return groupId, nil
}

// launchEC2Instance launches an EC2 instance with the given parameters.
func LaunchEC2Instance(sess *session.Session, instanceType, region, keyPairName string) (string, error) {
	svc := ec2.New(sess)

	userdata, err := PrepareUserDataScript()
	if err != nil {
		return "", fmt.Errorf("error preparing user data script: %v", err)
	}

	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String("ami-0944e91aed79c721c"), // Replace with actual AMI ID
		InstanceType: aws.String(instanceType),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
		UserData:     aws.String(userdata),
		KeyName:      aws.String(keyPairName), // Include the key pair name
		// Add additional parameters as necessary
	})
	if err != nil {
		return "", fmt.Errorf("error launching instance: %v", err)
	}

	instanceID := *runResult.Instances[0].InstanceId
	return instanceID, nil
}

// prepareUserDataScript reads the embedded user data script.
func PrepareUserDataScript() (string, error) {
	data, err := fs.ReadFile(userDataFS, "ovpn_userdata.sh")
	if err != nil {
		return "", err
	}
	encodedData := base64.StdEncoding.EncodeToString(data)
	return string(encodedData), nil
}

func CreateSSHKeyPair(sess *session.Session, keyPairName string) (string, error) {
	svc := ec2.New(sess)

	input := &ec2.CreateKeyPairInput{
		KeyName: aws.String(keyPairName),
	}

	result, err := svc.CreateKeyPair(input)
	if err != nil {
		return "", err
	}

	return *result.KeyMaterial, nil // Private key material
}

func SavePrivateKeyToFile(privateKey, filename string) error {
	// Construct the path to save the key in the current working directory
	currentDir, _ := os.Getwd()
	fullPath := filepath.Join(currentDir, filename)

	// Save the private key to the file
	return ioutil.WriteFile(fullPath, []byte(privateKey), 0600)
}

func KeyPairExists(svc *ec2.EC2, keyPairName string) (bool, error) {
	input := &ec2.DescribeKeyPairsInput{
		KeyNames: []*string{
			aws.String(keyPairName),
		},
	}

	_, err := svc.DescribeKeyPairs(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "InvalidKeyPair.NotFound":
				return false, nil
			default:
				return false, aerr
			}
		} else {
			return false, err
		}
	}

	return true, nil
}
