package provider

import (
	"fmt"

	"github.com/lspaccatrosi16/go-cli-tools/aws"
	"github.com/lspaccatrosi16/go-cli-tools/credential"
	"github.com/lspaccatrosi16/go-cli-tools/gcloud"
	"github.com/lspaccatrosi16/go-cli-tools/input"
	"github.com/lspaccatrosi16/go-cli-tools/storage"
)

func GetProvider(cred credential.Credential, bucket string) (storage.StorageProvider, error) {
	opts := []input.SelectOption{
		{Name: "AWS S3", Value: "s3"},
		{Name: "Firebase Cloud Storage", Value: "firebase"},
	}

	sel, err := input.GetSelection("Choose Provider", opts)

	if err != nil {
		return nil, err
	}

	switch sel {
	case "s3":
		region := input.GetInput("AWS Region (default: eu-west-2)")
		if region == "" {
			region = "eu-west-2"
		}

		awsCred := aws.GetCredentialFromValue(cred.Key, cred.Secret)

		config, err := aws.GetConfigWithCredential(awsCred, region)
		if err != nil {
			return nil, err
		}

		provider := aws.NewBucket(*config, bucket)
		return &provider, nil

	case "firebase":
		gcloud.RegisterServiceAccount([]byte(cred.Secret))
		client, err := gcloud.NewGStorage()
		if err != nil {
			return nil, err
		}

		provider, err := client.GetBucket(bucket)

		if err != nil {
			return nil, err
		}

		return provider, nil
	default:
		return nil, fmt.Errorf("unknown provider code %s", sel)
	}

}
