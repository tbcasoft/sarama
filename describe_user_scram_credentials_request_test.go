package sarama

import (
	"testing"

	"go.uber.org/goleak"
)

var (
	emptyDescribeUserScramCredentialsRequest = []byte{
		1, 0, // empty tagged fields
	}
	userDescribeUserScramCredentialsRequest = []byte{
		2,                            // DescribeUsers array, Array length 1
		7,                            // User name length 6
		'r', 'a', 'n', 'd', 'o', 'm', // User name
		0, 0, // empty tagged fields
	}
)

func TestDescribeUserScramCredentialsRequest(t *testing.T) {
	t.Cleanup(func() { goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick") })
	request := &DescribeUserScramCredentialsRequest{
		Version:       0,
		DescribeUsers: []DescribeUserScramCredentialsRequestUser{},
	}
	testRequest(t, "no users", request, emptyDescribeUserScramCredentialsRequest)

	request.DescribeUsers = []DescribeUserScramCredentialsRequestUser{
		{
			Name: "random",
		},
	}
	testRequest(t, "single user", request, userDescribeUserScramCredentialsRequest)
}
