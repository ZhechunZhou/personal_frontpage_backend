package common

import (
	"github.com/aymanbagabas/fss3"
)

var config fss3.Config

func Builds3Config(accessKeyID string, secretAccessKey string, region string, bucket string) {
	config = fss3.Config{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		Endpoint:        bucket + ".s3." + region + ".amazonaws.com",
		BucketName:      "assets",
		UseSSL:          true,
		Region:          region,
		DirFileName:     "",
		Umask:           0,
	}
}

func connectAws() (*fss3.FSS3, error) {
	sess, err := fss3.NewFSS3(config)
	return sess, err
}

func UploadDoc(address string, data []byte) error {
	s3, err := connectAws()

	if err != nil {
		return err
	}
	err = s3.WriteFile(address, data, 0666)
	return err
}

func DownloadDoc(address string) ([]byte, error) {
	s3, err := connectAws()
	if err != nil {
		return nil, err
	}

	file, err := s3.ReadFile(address)
	return file, err
}

func UploadEncryptedDoc(address string, data []byte) (string, error) {
	encrypted, hashcode, err := EncryptFile(data)
	if err != nil {
		return "", err
	}
	if err = UploadDoc(address, encrypted); err != nil {
		return "", err
	}
	return hashcode, nil
}

func DownloadEncryptedDoc(address string, hashcode string) ([]byte, error) {
	encrypted, err := DownloadDoc(address)
	if err != nil {
		return nil, err
	}
	data, err := DecryptFile(encrypted, hashcode)
	if err != nil {
		return nil, err
	}
	return data, nil
}
