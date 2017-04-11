package rancher

import (
	"regexp"
	"testing"
)

func bucketNameTest(t *testing.T) {
	projectName := "uploader-project-name"
	testBucketName := "s3-test-bucket"

	uploaderWithoutBucketName := &S3Uploader{
		Bucket: "",
	}
	uploaderWithBucketName := S3Uploader{
		Bucket: testBucketName,
	}

	if uploaderWithBucketName.bucketName(projectName) != testBucketName {
		t.Fail()
	}

	bucketRegex := regexp.MustCompile(projectName + "-[a-zA-Z0-9]{12}")
	if !bucketRegex.MatchString(uploaderWithoutBucketName.bucketName(projectName)) {
		t.Fail()
	}

}
