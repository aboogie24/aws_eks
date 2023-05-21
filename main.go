package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucketV2(ctx, "my-bucket", nil)
		if err != nil {
			return err
		}

		// Create an S3 Bucket object
		_, err = s3.NewBucketObjectv2(ctx, "testing", &s3.BucketObjectv2Args{
			Key:    pulumi.String("test"),
			Bucket: bucket.ID(),
			Source: pulumi.NewFileAsset("./index.html"),
		})
		if err != nil {
			return err
		}
		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())

		return nil
	})

}
