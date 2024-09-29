package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Storage options for the runner instance.
// Experimental.
type StorageOptions struct {
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for `volumeType`: `EbsDeviceVolumeType.IO1`
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	// Default: - none, required for `EbsDeviceVolumeType.IO1`
	//
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The throughput that the volume supports, in MiB/s Takes a minimum of 125 and maximum of 1000.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-throughput
	//
	// Default: - 125 MiB/s. Only valid on gp3 volumes.
	//
	// Experimental.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	// Default: `EbsDeviceVolumeType.GP2`
	//
	// Experimental.
	VolumeType awsec2.EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
}

