package service

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

type mockClient struct {
	Failed bool
}

func (c *mockClient) CreateDevicePool(*devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error) {
	return nil, nil
}
func (c *mockClient) CreateDevicePoolWithContext(aws.Context, *devicefarm.CreateDevicePoolInput, ...request.Option) (*devicefarm.CreateDevicePoolOutput, error) {
	return nil, nil
}
func (c *mockClient) CreateDevicePoolRequest(*devicefarm.CreateDevicePoolInput) (*request.Request, *devicefarm.CreateDevicePoolOutput) {
	return nil, nil
}

func (c *mockClient) CreateNetworkProfile(*devicefarm.CreateNetworkProfileInput) (*devicefarm.CreateNetworkProfileOutput, error) {
	return nil, nil
}
func (c *mockClient) CreateNetworkProfileWithContext(aws.Context, *devicefarm.CreateNetworkProfileInput, ...request.Option) (*devicefarm.CreateNetworkProfileOutput, error) {
	return nil, nil
}
func (c *mockClient) CreateNetworkProfileRequest(*devicefarm.CreateNetworkProfileInput) (*request.Request, *devicefarm.CreateNetworkProfileOutput) {
	return nil, nil
}

func (c *mockClient) CreateProject(*devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error) {
	return nil, nil
}
func (c *mockClient) CreateProjectWithContext(aws.Context, *devicefarm.CreateProjectInput, ...request.Option) (*devicefarm.CreateProjectOutput, error) {
	return nil, nil
}
func (c *mockClient) CreateProjectRequest(*devicefarm.CreateProjectInput) (*request.Request, *devicefarm.CreateProjectOutput) {
	return nil, nil
}

func (c *mockClient) CreateRemoteAccessSession(*devicefarm.CreateRemoteAccessSessionInput) (*devicefarm.CreateRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *mockClient) CreateRemoteAccessSessionWithContext(aws.Context, *devicefarm.CreateRemoteAccessSessionInput, ...request.Option) (*devicefarm.CreateRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *mockClient) CreateRemoteAccessSessionRequest(*devicefarm.CreateRemoteAccessSessionInput) (*request.Request, *devicefarm.CreateRemoteAccessSessionOutput) {
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
func (c *mockClient) CreateUploadWithContext(aws.Context, *devicefarm.CreateUploadInput, ...request.Option) (*devicefarm.CreateUploadOutput, error) {
	return nil, nil
}
func (c *mockClient) CreateUploadRequest(*devicefarm.CreateUploadInput) (*request.Request, *devicefarm.CreateUploadOutput) {
	return nil, nil
}

func (c *mockClient) DeleteDevicePool(*devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteDevicePoolWithContext(aws.Context, *devicefarm.DeleteDevicePoolInput, ...request.Option) (*devicefarm.DeleteDevicePoolOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteDevicePoolRequest(*devicefarm.DeleteDevicePoolInput) (*request.Request, *devicefarm.DeleteDevicePoolOutput) {
	return nil, nil
}

func (c *mockClient) DeleteNetworkProfile(*devicefarm.DeleteNetworkProfileInput) (*devicefarm.DeleteNetworkProfileOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteNetworkProfileWithContext(aws.Context, *devicefarm.DeleteNetworkProfileInput, ...request.Option) (*devicefarm.DeleteNetworkProfileOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteNetworkProfileRequest(*devicefarm.DeleteNetworkProfileInput) (*request.Request, *devicefarm.DeleteNetworkProfileOutput) {
	return nil, nil
}

func (c *mockClient) DeleteProject(*devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteProjectWithContext(aws.Context, *devicefarm.DeleteProjectInput, ...request.Option) (*devicefarm.DeleteProjectOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteProjectRequest(*devicefarm.DeleteProjectInput) (*request.Request, *devicefarm.DeleteProjectOutput) {
	return nil, nil
}

func (c *mockClient) DeleteRemoteAccessSession(*devicefarm.DeleteRemoteAccessSessionInput) (*devicefarm.DeleteRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteRemoteAccessSessionWithContext(aws.Context, *devicefarm.DeleteRemoteAccessSessionInput, ...request.Option) (*devicefarm.DeleteRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteRemoteAccessSessionRequest(*devicefarm.DeleteRemoteAccessSessionInput) (*request.Request, *devicefarm.DeleteRemoteAccessSessionOutput) {
	return nil, nil
}

func (c *mockClient) DeleteRun(*devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteRunWithContext(aws.Context, *devicefarm.DeleteRunInput, ...request.Option) (*devicefarm.DeleteRunOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteRunRequest(*devicefarm.DeleteRunInput) (*request.Request, *devicefarm.DeleteRunOutput) {
	return nil, nil
}

func (c *mockClient) DeleteUpload(*devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteUploadWithContext(aws.Context, *devicefarm.DeleteUploadInput, ...request.Option) (*devicefarm.DeleteUploadOutput, error) {
	return nil, nil
}
func (c *mockClient) DeleteUploadRequest(*devicefarm.DeleteUploadInput) (*request.Request, *devicefarm.DeleteUploadOutput) {
	return nil, nil
}

func (c *mockClient) GetAccountSettings(*devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error) {
	return nil, nil
}
func (c *mockClient) GetAccountSettingsWithContext(aws.Context, *devicefarm.GetAccountSettingsInput, ...request.Option) (*devicefarm.GetAccountSettingsOutput, error) {
	return nil, nil
}
func (c *mockClient) GetAccountSettingsRequest(*devicefarm.GetAccountSettingsInput) (*request.Request, *devicefarm.GetAccountSettingsOutput) {
	return nil, nil
}

func (c *mockClient) GetDevice(*devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error) {
	return nil, nil
}
func (c *mockClient) GetDeviceWithContext(aws.Context, *devicefarm.GetDeviceInput, ...request.Option) (*devicefarm.GetDeviceOutput, error) {
	return nil, nil
}
func (c *mockClient) GetDeviceRequest(*devicefarm.GetDeviceInput) (*request.Request, *devicefarm.GetDeviceOutput) {
	return nil, nil
}

func (c *mockClient) GetDevicePool(*devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error) {
	return nil, nil
}
func (c *mockClient) GetDevicePoolWithContext(aws.Context, *devicefarm.GetDevicePoolInput, ...request.Option) (*devicefarm.GetDevicePoolOutput, error) {
	return nil, nil
}
func (c *mockClient) GetDevicePoolRequest(*devicefarm.GetDevicePoolInput) (*request.Request, *devicefarm.GetDevicePoolOutput) {
	return nil, nil
}

func (c *mockClient) GetDevicePoolCompatibility(*devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
	return nil, nil
}
func (c *mockClient) GetDevicePoolCompatibilityWithContext(aws.Context, *devicefarm.GetDevicePoolCompatibilityInput, ...request.Option) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
	return nil, nil
}
func (c *mockClient) GetDevicePoolCompatibilityRequest(*devicefarm.GetDevicePoolCompatibilityInput) (*request.Request, *devicefarm.GetDevicePoolCompatibilityOutput) {
	return nil, nil
}

func (c *mockClient) GetJob(*devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error) {
	return nil, nil
}
func (c *mockClient) GetJobWithContext(aws.Context, *devicefarm.GetJobInput, ...request.Option) (*devicefarm.GetJobOutput, error) {
	return nil, nil
}
func (c *mockClient) GetJobRequest(*devicefarm.GetJobInput) (*request.Request, *devicefarm.GetJobOutput) {
	return nil, nil
}

func (c *mockClient) GetNetworkProfile(*devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error) {
	return nil, nil
}
func (c *mockClient) GetNetworkProfileWithContext(aws.Context, *devicefarm.GetNetworkProfileInput, ...request.Option) (*devicefarm.GetNetworkProfileOutput, error) {
	return nil, nil
}
func (c *mockClient) GetNetworkProfileRequest(*devicefarm.GetNetworkProfileInput) (*request.Request, *devicefarm.GetNetworkProfileOutput) {
	return nil, nil
}

func (c *mockClient) GetOfferingStatus(*devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error) {
	return nil, nil
}
func (c *mockClient) GetOfferingStatusWithContext(aws.Context, *devicefarm.GetOfferingStatusInput, ...request.Option) (*devicefarm.GetOfferingStatusOutput, error) {
	return nil, nil
}
func (c *mockClient) GetOfferingStatusRequest(*devicefarm.GetOfferingStatusInput) (*request.Request, *devicefarm.GetOfferingStatusOutput) {
	return nil, nil
}

func (c *mockClient) GetOfferingStatusPages(*devicefarm.GetOfferingStatusInput, func(*devicefarm.GetOfferingStatusOutput, bool) bool) error {
	return nil
}
func (c *mockClient) GetOfferingStatusPagesWithContext(aws.Context, *devicefarm.GetOfferingStatusInput, func(*devicefarm.GetOfferingStatusOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) GetProject(*devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error) {
	return nil, nil
}
func (c *mockClient) GetProjectWithContext(aws.Context, *devicefarm.GetProjectInput, ...request.Option) (*devicefarm.GetProjectOutput, error) {
	return nil, nil
}
func (c *mockClient) GetProjectRequest(*devicefarm.GetProjectInput) (*request.Request, *devicefarm.GetProjectOutput) {
	return nil, nil
}

func (c *mockClient) GetRemoteAccessSession(*devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *mockClient) GetRemoteAccessSessionWithContext(aws.Context, *devicefarm.GetRemoteAccessSessionInput, ...request.Option) (*devicefarm.GetRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *mockClient) GetRemoteAccessSessionRequest(*devicefarm.GetRemoteAccessSessionInput) (*request.Request, *devicefarm.GetRemoteAccessSessionOutput) {
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
func (c *mockClient) GetRunWithContext(aws.Context, *devicefarm.GetRunInput, ...request.Option) (*devicefarm.GetRunOutput, error) {
	return nil, nil
}
func (c *mockClient) GetRunRequest(*devicefarm.GetRunInput) (*request.Request, *devicefarm.GetRunOutput) {
	return nil, nil
}

func (c *mockClient) GetSuite(*devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error) {
	return nil, nil
}
func (c *mockClient) GetSuiteWithContext(aws.Context, *devicefarm.GetSuiteInput, ...request.Option) (*devicefarm.GetSuiteOutput, error) {
	return nil, nil
}
func (c *mockClient) GetSuiteRequest(*devicefarm.GetSuiteInput) (*request.Request, *devicefarm.GetSuiteOutput) {
	return nil, nil
}

func (c *mockClient) GetTest(*devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error) {
	return nil, nil
}
func (c *mockClient) GetTestWithContext(aws.Context, *devicefarm.GetTestInput, ...request.Option) (*devicefarm.GetTestOutput, error) {
	return nil, nil
}
func (c *mockClient) GetTestRequest(*devicefarm.GetTestInput) (*request.Request, *devicefarm.GetTestOutput) {
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
func (c *mockClient) GetUploadWithContext(aws.Context, *devicefarm.GetUploadInput, ...request.Option) (*devicefarm.GetUploadOutput, error) {
	return nil, nil
}
func (c *mockClient) GetUploadRequest(*devicefarm.GetUploadInput) (*request.Request, *devicefarm.GetUploadOutput) {
	return nil, nil
}

func (c *mockClient) InstallToRemoteAccessSession(*devicefarm.InstallToRemoteAccessSessionInput) (*devicefarm.InstallToRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *mockClient) InstallToRemoteAccessSessionWithContext(aws.Context, *devicefarm.InstallToRemoteAccessSessionInput, ...request.Option) (*devicefarm.InstallToRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *mockClient) InstallToRemoteAccessSessionRequest(*devicefarm.InstallToRemoteAccessSessionInput) (*request.Request, *devicefarm.InstallToRemoteAccessSessionOutput) {
	return nil, nil
}

func (c *mockClient) ListArtifacts(input *devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error) {
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
func (c *mockClient) ListArtifactsWithContext(aws.Context, *devicefarm.ListArtifactsInput, ...request.Option) (*devicefarm.ListArtifactsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListArtifactsRequest(*devicefarm.ListArtifactsInput) (*request.Request, *devicefarm.ListArtifactsOutput) {
	return nil, nil
}

func (c *mockClient) ListArtifactsPages(*devicefarm.ListArtifactsInput, func(*devicefarm.ListArtifactsOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListArtifactsPagesWithContext(aws.Context, *devicefarm.ListArtifactsInput, func(*devicefarm.ListArtifactsOutput, bool) bool, ...request.Option) error {
	return nil
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
func (c *mockClient) ListDevicePoolsWithContext(aws.Context, *devicefarm.ListDevicePoolsInput, ...request.Option) (*devicefarm.ListDevicePoolsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListDevicePoolsRequest(*devicefarm.ListDevicePoolsInput) (*request.Request, *devicefarm.ListDevicePoolsOutput) {
	return nil, nil
}

func (c *mockClient) ListDevicePoolsPages(*devicefarm.ListDevicePoolsInput, func(*devicefarm.ListDevicePoolsOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListDevicePoolsPagesWithContext(aws.Context, *devicefarm.ListDevicePoolsInput, func(*devicefarm.ListDevicePoolsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) ListDevices(*devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error) {
	return nil, nil
}
func (c *mockClient) ListDevicesWithContext(aws.Context, *devicefarm.ListDevicesInput, ...request.Option) (*devicefarm.ListDevicesOutput, error) {
	return nil, nil
}
func (c *mockClient) ListDevicesRequest(*devicefarm.ListDevicesInput) (*request.Request, *devicefarm.ListDevicesOutput) {
	return nil, nil
}

func (c *mockClient) ListDevicesPages(*devicefarm.ListDevicesInput, func(*devicefarm.ListDevicesOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListDevicesPagesWithContext(aws.Context, *devicefarm.ListDevicesInput, func(*devicefarm.ListDevicesOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) ListJobs(*devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error) {
	res := &devicefarm.ListJobsOutput{
		Jobs: []*devicefarm.Job{
			{
				Arn:  aws.String(""),
				Name: aws.String(""),
				Device: &devicefarm.Device{
					Platform: aws.String(""),
					Os:       aws.String(""),
				},
			},
		},
	}
	return res, nil
}
func (c *mockClient) ListJobsWithContext(aws.Context, *devicefarm.ListJobsInput, ...request.Option) (*devicefarm.ListJobsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListJobsRequest(*devicefarm.ListJobsInput) (*request.Request, *devicefarm.ListJobsOutput) {
	return nil, nil
}

func (c *mockClient) ListJobsPages(*devicefarm.ListJobsInput, func(*devicefarm.ListJobsOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListJobsPagesWithContext(aws.Context, *devicefarm.ListJobsInput, func(*devicefarm.ListJobsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) ListNetworkProfiles(*devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error) {
	return nil, nil
}
func (c *mockClient) ListNetworkProfilesWithContext(aws.Context, *devicefarm.ListNetworkProfilesInput, ...request.Option) (*devicefarm.ListNetworkProfilesOutput, error) {
	return nil, nil
}
func (c *mockClient) ListNetworkProfilesRequest(*devicefarm.ListNetworkProfilesInput) (*request.Request, *devicefarm.ListNetworkProfilesOutput) {
	return nil, nil
}

func (c *mockClient) ListOfferingPromotions(*devicefarm.ListOfferingPromotionsInput) (*devicefarm.ListOfferingPromotionsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListOfferingPromotionsWithContext(aws.Context, *devicefarm.ListOfferingPromotionsInput, ...request.Option) (*devicefarm.ListOfferingPromotionsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListOfferingPromotionsRequest(*devicefarm.ListOfferingPromotionsInput) (*request.Request, *devicefarm.ListOfferingPromotionsOutput) {
	return nil, nil
}

func (c *mockClient) ListOfferingTransactions(*devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListOfferingTransactionsWithContext(aws.Context, *devicefarm.ListOfferingTransactionsInput, ...request.Option) (*devicefarm.ListOfferingTransactionsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListOfferingTransactionsRequest(*devicefarm.ListOfferingTransactionsInput) (*request.Request, *devicefarm.ListOfferingTransactionsOutput) {
	return nil, nil
}

func (c *mockClient) ListOfferingTransactionsPages(*devicefarm.ListOfferingTransactionsInput, func(*devicefarm.ListOfferingTransactionsOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListOfferingTransactionsPagesWithContext(aws.Context, *devicefarm.ListOfferingTransactionsInput, func(*devicefarm.ListOfferingTransactionsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) ListOfferings(*devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListOfferingsWithContext(aws.Context, *devicefarm.ListOfferingsInput, ...request.Option) (*devicefarm.ListOfferingsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListOfferingsRequest(*devicefarm.ListOfferingsInput) (*request.Request, *devicefarm.ListOfferingsOutput) {
	return nil, nil
}

func (c *mockClient) ListOfferingsPages(*devicefarm.ListOfferingsInput, func(*devicefarm.ListOfferingsOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListOfferingsPagesWithContext(aws.Context, *devicefarm.ListOfferingsInput, func(*devicefarm.ListOfferingsOutput, bool) bool, ...request.Option) error {
	return nil
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
func (c *mockClient) ListProjectsWithContext(aws.Context, *devicefarm.ListProjectsInput, ...request.Option) (*devicefarm.ListProjectsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListProjectsRequest(*devicefarm.ListProjectsInput) (*request.Request, *devicefarm.ListProjectsOutput) {
	return nil, nil
}

func (c *mockClient) ListProjectsPages(*devicefarm.ListProjectsInput, func(*devicefarm.ListProjectsOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListProjectsPagesWithContext(aws.Context, *devicefarm.ListProjectsInput, func(*devicefarm.ListProjectsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) ListRemoteAccessSessions(*devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListRemoteAccessSessionsWithContext(aws.Context, *devicefarm.ListRemoteAccessSessionsInput, ...request.Option) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListRemoteAccessSessionsRequest(*devicefarm.ListRemoteAccessSessionsInput) (*request.Request, *devicefarm.ListRemoteAccessSessionsOutput) {
	return nil, nil
}

func (c *mockClient) ListRuns(*devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListRunsWithContext(aws.Context, *devicefarm.ListRunsInput, ...request.Option) (*devicefarm.ListRunsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListRunsRequest(*devicefarm.ListRunsInput) (*request.Request, *devicefarm.ListRunsOutput) {
	return nil, nil
}

func (c *mockClient) ListRunsPages(*devicefarm.ListRunsInput, func(*devicefarm.ListRunsOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListRunsPagesWithContext(aws.Context, *devicefarm.ListRunsInput, func(*devicefarm.ListRunsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) ListSamples(*devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error) {
	return nil, nil
}
func (c *mockClient) ListSamplesWithContext(aws.Context, *devicefarm.ListSamplesInput, ...request.Option) (*devicefarm.ListSamplesOutput, error) {
	return nil, nil
}
func (c *mockClient) ListSamplesRequest(*devicefarm.ListSamplesInput) (*request.Request, *devicefarm.ListSamplesOutput) {
	return nil, nil
}

func (c *mockClient) ListSamplesPages(*devicefarm.ListSamplesInput, func(*devicefarm.ListSamplesOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListSamplesPagesWithContext(aws.Context, *devicefarm.ListSamplesInput, func(*devicefarm.ListSamplesOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) ListSuites(*devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error) {
	var res *devicefarm.ListSuitesOutput
	if c.Failed {
		res = &devicefarm.ListSuitesOutput{
			Suites: []*devicefarm.Suite{
				{
					Arn:    aws.String("fail"),
					Result: aws.String(devicefarm.ExecutionResultFailed),
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
func (c *mockClient) ListSuitesWithContext(aws.Context, *devicefarm.ListSuitesInput, ...request.Option) (*devicefarm.ListSuitesOutput, error) {
	return nil, nil
}
func (c *mockClient) ListSuitesRequest(*devicefarm.ListSuitesInput) (*request.Request, *devicefarm.ListSuitesOutput) {
	return nil, nil
}

func (c *mockClient) ListSuitesPages(*devicefarm.ListSuitesInput, func(*devicefarm.ListSuitesOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListSuitesPagesWithContext(aws.Context, *devicefarm.ListSuitesInput, func(*devicefarm.ListSuitesOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) ListTests(input *devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error) {
	var res *devicefarm.ListTestsOutput
	if *input.Arn == "fail" {
		res = &devicefarm.ListTestsOutput{
			Tests: []*devicefarm.Test{
				{
					Arn:     aws.String("fail 1"),
					Message: aws.String("Fail :("),
					Result:  aws.String(devicefarm.ExecutionResultFailed),
				},
				{
					Arn:     aws.String("fail 2"),
					Message: aws.String("Fail :("),
					Result:  aws.String(devicefarm.ExecutionResultFailed),
				},
			},
		}
	} else {
		res = &devicefarm.ListTestsOutput{
			Tests: []*devicefarm.Test{
				{
					Arn: aws.String(""),
				},
			},
		}
	}
	return res, nil
}
func (c *mockClient) ListTestsWithContext(aws.Context, *devicefarm.ListTestsInput, ...request.Option) (*devicefarm.ListTestsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListTestsRequest(*devicefarm.ListTestsInput) (*request.Request, *devicefarm.ListTestsOutput) {
	return nil, nil
}

func (c *mockClient) ListTestsPages(*devicefarm.ListTestsInput, func(*devicefarm.ListTestsOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListTestsPagesWithContext(aws.Context, *devicefarm.ListTestsInput, func(*devicefarm.ListTestsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) ListUniqueProblems(*devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListUniqueProblemsWithContext(aws.Context, *devicefarm.ListUniqueProblemsInput, ...request.Option) (*devicefarm.ListUniqueProblemsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListUniqueProblemsRequest(*devicefarm.ListUniqueProblemsInput) (*request.Request, *devicefarm.ListUniqueProblemsOutput) {
	return nil, nil
}

func (c *mockClient) ListUniqueProblemsPages(*devicefarm.ListUniqueProblemsInput, func(*devicefarm.ListUniqueProblemsOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListUniqueProblemsPagesWithContext(aws.Context, *devicefarm.ListUniqueProblemsInput, func(*devicefarm.ListUniqueProblemsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) ListUploads(*devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListUploadsWithContext(aws.Context, *devicefarm.ListUploadsInput, ...request.Option) (*devicefarm.ListUploadsOutput, error) {
	return nil, nil
}
func (c *mockClient) ListUploadsRequest(*devicefarm.ListUploadsInput) (*request.Request, *devicefarm.ListUploadsOutput) {
	return nil, nil
}

func (c *mockClient) ListUploadsPages(*devicefarm.ListUploadsInput, func(*devicefarm.ListUploadsOutput, bool) bool) error {
	return nil
}
func (c *mockClient) ListUploadsPagesWithContext(aws.Context, *devicefarm.ListUploadsInput, func(*devicefarm.ListUploadsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *mockClient) PurchaseOffering(*devicefarm.PurchaseOfferingInput) (*devicefarm.PurchaseOfferingOutput, error) {
	return nil, nil
}
func (c *mockClient) PurchaseOfferingWithContext(aws.Context, *devicefarm.PurchaseOfferingInput, ...request.Option) (*devicefarm.PurchaseOfferingOutput, error) {
	return nil, nil
}
func (c *mockClient) PurchaseOfferingRequest(*devicefarm.PurchaseOfferingInput) (*request.Request, *devicefarm.PurchaseOfferingOutput) {
	return nil, nil
}

func (c *mockClient) RenewOffering(*devicefarm.RenewOfferingInput) (*devicefarm.RenewOfferingOutput, error) {
	return nil, nil
}
func (c *mockClient) RenewOfferingWithContext(aws.Context, *devicefarm.RenewOfferingInput, ...request.Option) (*devicefarm.RenewOfferingOutput, error) {
	return nil, nil
}
func (c *mockClient) RenewOfferingRequest(*devicefarm.RenewOfferingInput) (*request.Request, *devicefarm.RenewOfferingOutput) {
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
func (c *mockClient) ScheduleRunWithContext(aws.Context, *devicefarm.ScheduleRunInput, ...request.Option) (*devicefarm.ScheduleRunOutput, error) {
	return nil, nil
}
func (c *mockClient) ScheduleRunRequest(*devicefarm.ScheduleRunInput) (*request.Request, *devicefarm.ScheduleRunOutput) {
	return nil, nil
}

func (c *mockClient) StopRemoteAccessSession(*devicefarm.StopRemoteAccessSessionInput) (*devicefarm.StopRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *mockClient) StopRemoteAccessSessionWithContext(aws.Context, *devicefarm.StopRemoteAccessSessionInput, ...request.Option) (*devicefarm.StopRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *mockClient) StopRemoteAccessSessionRequest(*devicefarm.StopRemoteAccessSessionInput) (*request.Request, *devicefarm.StopRemoteAccessSessionOutput) {
	return nil, nil
}

func (c *mockClient) StopRun(*devicefarm.StopRunInput) (*devicefarm.StopRunOutput, error) {
	return nil, nil
}
func (c *mockClient) StopRunWithContext(aws.Context, *devicefarm.StopRunInput, ...request.Option) (*devicefarm.StopRunOutput, error) {
	return nil, nil
}
func (c *mockClient) StopRunRequest(*devicefarm.StopRunInput) (*request.Request, *devicefarm.StopRunOutput) {
	return nil, nil
}

func (c *mockClient) UpdateDevicePool(*devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error) {
	return nil, nil
}
func (c *mockClient) UpdateDevicePoolWithContext(aws.Context, *devicefarm.UpdateDevicePoolInput, ...request.Option) (*devicefarm.UpdateDevicePoolOutput, error) {
	return nil, nil
}
func (c *mockClient) UpdateDevicePoolRequest(*devicefarm.UpdateDevicePoolInput) (*request.Request, *devicefarm.UpdateDevicePoolOutput) {
	return nil, nil
}

func (c *mockClient) UpdateNetworkProfile(*devicefarm.UpdateNetworkProfileInput) (*devicefarm.UpdateNetworkProfileOutput, error) {
	return nil, nil
}
func (c *mockClient) UpdateNetworkProfileWithContext(aws.Context, *devicefarm.UpdateNetworkProfileInput, ...request.Option) (*devicefarm.UpdateNetworkProfileOutput, error) {
	return nil, nil
}
func (c *mockClient) UpdateNetworkProfileRequest(*devicefarm.UpdateNetworkProfileInput) (*request.Request, *devicefarm.UpdateNetworkProfileOutput) {
	return nil, nil
}

func (c *mockClient) UpdateProject(*devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error) {
	return nil, nil
}
func (c *mockClient) UpdateProjectWithContext(aws.Context, *devicefarm.UpdateProjectInput, ...request.Option) (*devicefarm.UpdateProjectOutput, error) {
	return nil, nil
}
func (c *mockClient) UpdateProjectRequest(*devicefarm.UpdateProjectInput) (*request.Request, *devicefarm.UpdateProjectOutput) {
	return nil, nil
}
