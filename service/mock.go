package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

type mockClient struct {
	Failed bool
}

func (c *mockClient) CreateDevicePoolRequest(*devicefarm.CreateDevicePoolInput) (*request.Request, *devicefarm.CreateDevicePoolOutput) {
	return nil, nil
}

func (c *mockClient) CreateDevicePool(*devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error) {
	return nil, nil
}

func (c *mockClient) CreateProjectRequest(*devicefarm.CreateProjectInput) (*request.Request, *devicefarm.CreateProjectOutput) {
	return nil, nil
}

func (c *mockClient) CreateProject(*devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error) {
	return nil, nil
}

func (c *mockClient) CreateUploadRequest(*devicefarm.CreateUploadInput) (*request.Request, *devicefarm.CreateUploadOutput) {
	return nil, nil
}

func (c *mockClient) CreateUpload(*devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error) {
	var res *devicefarm.CreateUploadOutput
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

func (c *mockClient) DeleteDevicePoolRequest(*devicefarm.DeleteDevicePoolInput) (*request.Request, *devicefarm.DeleteDevicePoolOutput) {
	return nil, nil
}

func (c *mockClient) DeleteDevicePool(*devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error) {
	return nil, nil
}

func (c *mockClient) DeleteProjectRequest(*devicefarm.DeleteProjectInput) (*request.Request, *devicefarm.DeleteProjectOutput) {
	return nil, nil
}

func (c *mockClient) DeleteProject(*devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error) {
	return nil, nil
}

func (c *mockClient) DeleteRunRequest(*devicefarm.DeleteRunInput) (*request.Request, *devicefarm.DeleteRunOutput) {
	return nil, nil
}

func (c *mockClient) DeleteRun(*devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error) {
	return nil, nil
}

func (c *mockClient) DeleteUploadRequest(*devicefarm.DeleteUploadInput) (*request.Request, *devicefarm.DeleteUploadOutput) {
	return nil, nil
}

func (c *mockClient) DeleteUpload(*devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error) {
	return nil, nil
}

func (c *mockClient) GetAccountSettingsRequest(*devicefarm.GetAccountSettingsInput) (*request.Request, *devicefarm.GetAccountSettingsOutput) {
	return nil, nil
}

func (c *mockClient) GetAccountSettings(*devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error) {
	return nil, nil
}

func (c *mockClient) GetDeviceRequest(*devicefarm.GetDeviceInput) (*request.Request, *devicefarm.GetDeviceOutput) {
	return nil, nil
}

func (c *mockClient) GetDevice(*devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error) {
	return nil, nil
}

func (c *mockClient) GetDevicePoolRequest(*devicefarm.GetDevicePoolInput) (*request.Request, *devicefarm.GetDevicePoolOutput) {
	return nil, nil
}

func (c *mockClient) GetDevicePool(*devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error) {
	return nil, nil
}

func (c *mockClient) GetDevicePoolCompatibilityRequest(*devicefarm.GetDevicePoolCompatibilityInput) (*request.Request, *devicefarm.GetDevicePoolCompatibilityOutput) {
	return nil, nil
}

func (c *mockClient) GetDevicePoolCompatibility(*devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
	return nil, nil
}

func (c *mockClient) GetJobRequest(*devicefarm.GetJobInput) (*request.Request, *devicefarm.GetJobOutput) {
	return nil, nil
}

func (c *mockClient) GetJob(*devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error) {
	return nil, nil
}

func (c *mockClient) GetProjectRequest(*devicefarm.GetProjectInput) (*request.Request, *devicefarm.GetProjectOutput) {
	return nil, nil
}

func (c *mockClient) GetProject(*devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error) {
	return nil, nil
}

func (c *mockClient) GetRunRequest(*devicefarm.GetRunInput) (*request.Request, *devicefarm.GetRunOutput) {
	return nil, nil
}

func (c *mockClient) GetRun(*devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error) {
	var res *devicefarm.GetRunOutput
	if c.Failed {
		res = &devicefarm.GetRunOutput{
			Run: &devicefarm.Run{
				Status: aws.String(""),
				Result: aws.String(""),
			},
		}
	} else {
		res = &devicefarm.GetRunOutput{
			Run: &devicefarm.Run{
				Status: aws.String(devicefarm.ExecutionStatusCompleted),
				Result: aws.String("ok"),
			},
		}
	}
	return res, nil
}

func (c *mockClient) GetSuiteRequest(*devicefarm.GetSuiteInput) (*request.Request, *devicefarm.GetSuiteOutput) {
	return nil, nil
}

func (c *mockClient) GetSuite(*devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error) {
	return nil, nil
}

func (c *mockClient) GetTestRequest(*devicefarm.GetTestInput) (*request.Request, *devicefarm.GetTestOutput) {
	return nil, nil
}

func (c *mockClient) GetTest(*devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error) {
	return nil, nil
}

func (c *mockClient) GetUploadRequest(*devicefarm.GetUploadInput) (*request.Request, *devicefarm.GetUploadOutput) {
	return nil, nil
}

func (c *mockClient) GetUpload(*devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error) {
	var res *devicefarm.GetUploadOutput
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

func (c *mockClient) ListArtifactsRequest(*devicefarm.ListArtifactsInput) (*request.Request, *devicefarm.ListArtifactsOutput) {
	return nil, nil
}

func (c *mockClient) ListArtifacts(*devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error) {
	return nil, nil
}

func (c *mockClient) ListArtifactsPages(*devicefarm.ListArtifactsInput, func(*devicefarm.ListArtifactsOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ListDevicePoolsRequest(*devicefarm.ListDevicePoolsInput) (*request.Request, *devicefarm.ListDevicePoolsOutput) {
	return nil, nil
}

func (c *mockClient) ListDevicePools(*devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error) {
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

func (c *mockClient) ListDevicePoolsPages(*devicefarm.ListDevicePoolsInput, func(*devicefarm.ListDevicePoolsOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ListDevicesRequest(*devicefarm.ListDevicesInput) (*request.Request, *devicefarm.ListDevicesOutput) {
	return nil, nil
}

func (c *mockClient) ListDevices(*devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error) {
	return nil, nil
}

func (c *mockClient) ListDevicesPages(*devicefarm.ListDevicesInput, func(*devicefarm.ListDevicesOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ListJobsRequest(*devicefarm.ListJobsInput) (*request.Request, *devicefarm.ListJobsOutput) {
	return nil, nil
}

func (c *mockClient) ListJobs(*devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error) {
	return nil, nil
}

func (c *mockClient) ListJobsPages(*devicefarm.ListJobsInput, func(*devicefarm.ListJobsOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ListProjectsRequest(*devicefarm.ListProjectsInput) (*request.Request, *devicefarm.ListProjectsOutput) {
	return nil, nil
}

func (c *mockClient) ListProjects(*devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error) {
	var res *devicefarm.ListProjectsOutput
	if c.Failed {
		res = &devicefarm.ListProjectsOutput{}
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

func (c *mockClient) ListProjectsPages(*devicefarm.ListProjectsInput, func(*devicefarm.ListProjectsOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ListRunsRequest(*devicefarm.ListRunsInput) (*request.Request, *devicefarm.ListRunsOutput) {
	return nil, nil
}

func (c *mockClient) ListRuns(*devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error) {
	return nil, nil
}

func (c *mockClient) ListRunsPages(*devicefarm.ListRunsInput, func(*devicefarm.ListRunsOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ListSamplesRequest(*devicefarm.ListSamplesInput) (*request.Request, *devicefarm.ListSamplesOutput) {
	return nil, nil
}

func (c *mockClient) ListSamples(*devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error) {
	return nil, nil
}

func (c *mockClient) ListSamplesPages(*devicefarm.ListSamplesInput, func(*devicefarm.ListSamplesOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ListSuitesRequest(*devicefarm.ListSuitesInput) (*request.Request, *devicefarm.ListSuitesOutput) {
	return nil, nil
}

func (c *mockClient) ListSuites(*devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error) {
	return nil, nil
}

func (c *mockClient) ListSuitesPages(*devicefarm.ListSuitesInput, func(*devicefarm.ListSuitesOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ListTestsRequest(*devicefarm.ListTestsInput) (*request.Request, *devicefarm.ListTestsOutput) {
	return nil, nil
}

func (c *mockClient) ListTests(*devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error) {
	return nil, nil
}

func (c *mockClient) ListTestsPages(*devicefarm.ListTestsInput, func(*devicefarm.ListTestsOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ListUniqueProblemsRequest(*devicefarm.ListUniqueProblemsInput) (*request.Request, *devicefarm.ListUniqueProblemsOutput) {
	return nil, nil
}

func (c *mockClient) ListUniqueProblems(*devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error) {
	return nil, nil
}

func (c *mockClient) ListUniqueProblemsPages(*devicefarm.ListUniqueProblemsInput, func(*devicefarm.ListUniqueProblemsOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ListUploadsRequest(*devicefarm.ListUploadsInput) (*request.Request, *devicefarm.ListUploadsOutput) {
	return nil, nil
}

func (c *mockClient) ListUploads(*devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error) {
	return nil, nil
}

func (c *mockClient) ListUploadsPages(*devicefarm.ListUploadsInput, func(*devicefarm.ListUploadsOutput, bool) bool) error {
	return nil
}

func (c *mockClient) ScheduleRunRequest(*devicefarm.ScheduleRunInput) (*request.Request, *devicefarm.ScheduleRunOutput) {
	return nil, nil
}

func (c *mockClient) ScheduleRun(*devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error) {
	var res *devicefarm.ScheduleRunOutput
	if c.Failed {
		res = &devicefarm.ScheduleRunOutput{
			Run: &devicefarm.Run{
				Arn:    aws.String(""),
				Status: aws.String(""),
			},
		}
	} else {
		res = &devicefarm.ScheduleRunOutput{
			Run: &devicefarm.Run{
				Arn:    aws.String("123"),
				Status: aws.String("OK"),
			},
		}
	}
	return res, nil
}

func (c *mockClient) UpdateDevicePoolRequest(*devicefarm.UpdateDevicePoolInput) (*request.Request, *devicefarm.UpdateDevicePoolOutput) {
	return nil, nil
}

func (c *mockClient) UpdateDevicePool(*devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error) {
	return nil, nil
}

func (c *mockClient) UpdateProjectRequest(*devicefarm.UpdateProjectInput) (*request.Request, *devicefarm.UpdateProjectOutput) {
	return nil, nil
}

func (c *mockClient) UpdateProject(*devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error) {
	return nil, nil
}
