package service

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"
)

// UploadARN using for tests
const UploadARN = "wefere3f3f33gv3fre3f3f3f3f3v34v3v43v433v34v43v34"

// CreateFakeServer creates a fake server for tests
func CreateFakeServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	return server
}

// MockClient mock for AWS Device Farm API
type MockClient struct {
	devicefarmiface.DeviceFarmAPI
	Failed                  bool
	UploadTest              bool
	AWSFail                 bool
	PartlyUnavailableDevice bool
	FullUnavailableDevices  bool
	FakeServer              *httptest.Server
}

func (c *MockClient) CreateUpload(*devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error) {
	var res *devicefarm.CreateUploadOutput
	if c.UploadTest {
		res = &devicefarm.CreateUploadOutput{
			Upload: &devicefarm.Upload{
				Arn:  aws.String(UploadARN),
				Url:  aws.String(c.FakeServer.URL),
				Type: aws.String(devicefarm.UploadTypeIosApp),
			},
		}
		return res, nil
	}
	if c.Failed {
		res = &devicefarm.CreateUploadOutput{
			Upload: &devicefarm.Upload{
				Arn:  aws.String(""),
				Url:  aws.String(""),
				Type: aws.String(devicefarm.UploadTypeIosApp),
			},
		}
	} else {
		res = &devicefarm.CreateUploadOutput{
			Upload: &devicefarm.Upload{
				Arn:  aws.String("111"),
				Url:  aws.String("localhost"),
				Type: aws.String(devicefarm.UploadTypeAndroidApp),
			},
		}
	}
	return res, nil
}

func (c *MockClient) GetRun(*devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error) {
	var res *devicefarm.GetRunOutput
	if c.Failed {
		res = &devicefarm.GetRunOutput{
			Run: &devicefarm.Run{
				Status: aws.String(""),
				Result: aws.String(""),
			},
		}
		if c.UploadTest {
			res.Run.Status = aws.String(devicefarm.ExecutionStatusCompleted)
		}
	} else {
		res = &devicefarm.GetRunOutput{
			Run: &devicefarm.Run{
				Status: aws.String(devicefarm.ExecutionStatusCompleted),
				Result: aws.String(devicefarm.ExecutionResultPassed),
			},
		}
	}
	return res, nil
}

func (c *MockClient) GetUpload(input *devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error) {
	var res *devicefarm.GetUploadOutput
	if *input.Arn == UploadARN {
		res = &devicefarm.GetUploadOutput{
			Upload: &devicefarm.Upload{
				Status: aws.String(devicefarm.UploadStatusSucceeded),
			},
		}
		return res, nil
	}
	if c.Failed {
		res = &devicefarm.GetUploadOutput{
			Upload: &devicefarm.Upload{
				Status: aws.String(""),
			},
		}
	} else {
		res = &devicefarm.GetUploadOutput{
			Upload: &devicefarm.Upload{
				Status: aws.String(devicefarm.UploadStatusSucceeded),
			},
		}
	}
	return res, nil
}

func (c *MockClient) ListArtifacts(input *devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error) {
	var res *devicefarm.ListArtifactsOutput
	if strings.Contains(*input.Arn, "fail") {
		res = &devicefarm.ListArtifactsOutput{
			Artifacts: []*devicefarm.Artifact{
				{
					Arn:  aws.String(""),
					Type: aws.String(devicefarm.ArtifactTypeUnknown),
				},
			},
		}
	} else {
		res = &devicefarm.ListArtifactsOutput{
			Artifacts: []*devicefarm.Artifact{
				{
					Arn: aws.String(""),
				},
			},
		}
	}
	return res, nil
}

func (c *MockClient) ListDevicePools(*devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error) {
	var res *devicefarm.ListDevicePoolsOutput
	if c.Failed {
		res = &devicefarm.ListDevicePoolsOutput{
			DevicePools: []*devicefarm.DevicePool{},
		}
	} else {
		res = &devicefarm.ListDevicePoolsOutput{
			DevicePools: []*devicefarm.DevicePool{
				{
					Name: aws.String("a"),
					Arn:  aws.String("b"),
				},
				{
					Name: aws.String("111"),
					Arn:  aws.String("test"),
				},
			},
		}
	}
	return res, nil
}

func (c *MockClient) ListJobs(input *devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error) {
	if strings.HasSuffix(*input.Arn, "22222222-2222-2222-2222-222222222222") {
		return &devicefarm.ListJobsOutput{
			Jobs: []*devicefarm.Job{
				{
					Arn:  aws.String(""),
					Name: aws.String(""),
					Device: &devicefarm.Device{
						Platform: aws.String(""),
						Os:       aws.String(""),
					},
					Result: aws.String(devicefarm.ExecutionResultFailed),
				},
			},
		}, nil
	}

	res := &devicefarm.ListJobsOutput{
		Jobs: []*devicefarm.Job{
			{
				Arn:  aws.String(""),
				Name: aws.String(""),
				Device: &devicefarm.Device{
					Platform: aws.String(""),
					Os:       aws.String(""),
				},
				Result: aws.String(devicefarm.ExecutionResultPassed),
			},
		},
	}

	if c.PartlyUnavailableDevice {
		res.Jobs = append(res.Jobs, &devicefarm.Job{
			Arn:  aws.String(""),
			Name: aws.String(""),
			Device: &devicefarm.Device{
				Platform: aws.String(""),
				Os:       aws.String(""),
			},
			Result: aws.String(devicefarm.ExecutionResultErrored),
		})
	}

	if c.FullUnavailableDevices {
		res = &devicefarm.ListJobsOutput{
			Jobs: []*devicefarm.Job{
				{
					Arn:  aws.String(""),
					Name: aws.String(""),
					Device: &devicefarm.Device{
						Platform: aws.String(""),
						Os:       aws.String(""),
					},
					Result: aws.String(devicefarm.ExecutionResultErrored),
				},
			},
		}
	}

	return res, nil
}

func (c *MockClient) ListProjects(*devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error) {
	var res *devicefarm.ListProjectsOutput
	if c.Failed {
		res = &devicefarm.ListProjectsOutput{}
		if c.UploadTest {
			res.Projects = []*devicefarm.Project{
				{
					Arn:  aws.String("qwerty"),
					Name: aws.String("test"),
				},
			}
		}
	} else {
		res = &devicefarm.ListProjectsOutput{
			Projects: []*devicefarm.Project{
				{
					Arn:  aws.String("111"),
					Name: aws.String("111"),
				},
				{
					Arn:  aws.String("qwerty"),
					Name: aws.String("test"),
				},
			},
		}
	}
	return res, nil
}

func (c *MockClient) ListSuites(*devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error) {
	var res *devicefarm.ListSuitesOutput
	if c.Failed && !c.AWSFail {
		res = &devicefarm.ListSuitesOutput{
			Suites: []*devicefarm.Suite{
				{
					Arn:    aws.String("fail"),
					Result: aws.String(devicefarm.ExecutionResultFailed),
				},
			},
		}
	} else if c.AWSFail {
		res = &devicefarm.ListSuitesOutput{
			Suites: []*devicefarm.Suite{
				{
					Arn:    aws.String("AWSFail"),
					Result: aws.String(devicefarm.ExecutionResultSkipped),
				},
			},
		}
	} else {
		res = &devicefarm.ListSuitesOutput{
			Suites: []*devicefarm.Suite{
				{
					Arn:    aws.String(""),
					Result: aws.String(devicefarm.ExecutionResultPassed),
				},
			},
		}
	}
	return res, nil
}

func (c *MockClient) ListTests(input *devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error) {
	var res *devicefarm.ListTestsOutput
	if *input.Arn == "fail" {
		res = &devicefarm.ListTestsOutput{
			Tests: []*devicefarm.Test{
				{
					Arn:     aws.String("fail 1"),
					Message: aws.String("Fail :("),
					Name:    aws.String("Setup Test"),
					Result:  aws.String(devicefarm.ExecutionResultFailed),
				},
				{
					Arn:     aws.String("fail 2"),
					Message: aws.String("Fail :("),
					Name:    aws.String("Teardown Test"),
					Result:  aws.String(devicefarm.ExecutionResultFailed),
				},
				{
					Arn:     aws.String("fail 2"),
					Message: aws.String("Fail :("),
					Name:    aws.String(""),
					Result:  aws.String(devicefarm.ExecutionResultPassed),
				},
			},
		}
	} else if *input.Arn == "AWSFail" {
		res = &devicefarm.ListTestsOutput{
			Tests: []*devicefarm.Test{
				{
					Arn: aws.String("AWSFail"),
				},
			},
		}
	} else {
		res = &devicefarm.ListTestsOutput{
			Tests: []*devicefarm.Test{
				{
					Arn:     aws.String("fail 1"),
					Message: aws.String("Fail :("),
					Name:    aws.String("Setup Test"),
					Result:  aws.String(devicefarm.ExecutionResultFailed),
				},
				{
					Arn:     aws.String("fail 2"),
					Message: aws.String("Fail :("),
					Name:    aws.String("Teardown Test"),
					Result:  aws.String(devicefarm.ExecutionResultFailed),
				},
				{
					Arn:     aws.String("fail 2"),
					Message: aws.String("Fail :("),
					Name:    aws.String(""),
					Result:  aws.String(devicefarm.ExecutionResultPassed),
				},
			},
		}
	}
	return res, nil
}

func (c *MockClient) ScheduleRun(*devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error) {
	var res *devicefarm.ScheduleRunOutput
	if c.Failed {
		res = &devicefarm.ScheduleRunOutput{
			Run: &devicefarm.Run{
				Arn:    aws.String(""),
				Status: aws.String(""),
			},
		}
		if c.UploadTest {
			res.Run.Status = aws.String(devicefarm.ExecutionStatusScheduling)
		}
	} else {
		res = &devicefarm.ScheduleRunOutput{
			Run: &devicefarm.Run{
				Arn:    aws.String("123"),
				Status: aws.String(devicefarm.ExecutionStatusScheduling),
			},
		}
	}
	return res, nil
}
