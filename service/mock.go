package service

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

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
	Failed     bool
	UploadTest bool
	AWSFail    bool
	FakeServer *httptest.Server
}

func (c *MockClient) CreateDevicePool(*devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error) {
	return nil, nil
}
func (c *MockClient) CreateDevicePoolWithContext(aws.Context, *devicefarm.CreateDevicePoolInput, ...request.Option) (*devicefarm.CreateDevicePoolOutput, error) {
	return nil, nil
}
func (c *MockClient) CreateDevicePoolRequest(*devicefarm.CreateDevicePoolInput) (*request.Request, *devicefarm.CreateDevicePoolOutput) {
	return nil, nil
}

func (c *MockClient) CreateNetworkProfile(*devicefarm.CreateNetworkProfileInput) (*devicefarm.CreateNetworkProfileOutput, error) {
	return nil, nil
}
func (c *MockClient) CreateNetworkProfileWithContext(aws.Context, *devicefarm.CreateNetworkProfileInput, ...request.Option) (*devicefarm.CreateNetworkProfileOutput, error) {
	return nil, nil
}
func (c *MockClient) CreateNetworkProfileRequest(*devicefarm.CreateNetworkProfileInput) (*request.Request, *devicefarm.CreateNetworkProfileOutput) {
	return nil, nil
}

func (c *MockClient) CreateProject(*devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error) {
	return nil, nil
}
func (c *MockClient) CreateProjectWithContext(aws.Context, *devicefarm.CreateProjectInput, ...request.Option) (*devicefarm.CreateProjectOutput, error) {
	return nil, nil
}
func (c *MockClient) CreateProjectRequest(*devicefarm.CreateProjectInput) (*request.Request, *devicefarm.CreateProjectOutput) {
	return nil, nil
}

func (c *MockClient) CreateRemoteAccessSession(*devicefarm.CreateRemoteAccessSessionInput) (*devicefarm.CreateRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *MockClient) CreateRemoteAccessSessionWithContext(aws.Context, *devicefarm.CreateRemoteAccessSessionInput, ...request.Option) (*devicefarm.CreateRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *MockClient) CreateRemoteAccessSessionRequest(*devicefarm.CreateRemoteAccessSessionInput) (*request.Request, *devicefarm.CreateRemoteAccessSessionOutput) {
	return nil, nil
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
func (c *MockClient) CreateUploadWithContext(aws.Context, *devicefarm.CreateUploadInput, ...request.Option) (*devicefarm.CreateUploadOutput, error) {
	return nil, nil
}
func (c *MockClient) CreateUploadRequest(*devicefarm.CreateUploadInput) (*request.Request, *devicefarm.CreateUploadOutput) {
	return nil, nil
}

func (c *MockClient) DeleteDevicePool(*devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteDevicePoolWithContext(aws.Context, *devicefarm.DeleteDevicePoolInput, ...request.Option) (*devicefarm.DeleteDevicePoolOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteDevicePoolRequest(*devicefarm.DeleteDevicePoolInput) (*request.Request, *devicefarm.DeleteDevicePoolOutput) {
	return nil, nil
}

func (c *MockClient) DeleteNetworkProfile(*devicefarm.DeleteNetworkProfileInput) (*devicefarm.DeleteNetworkProfileOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteNetworkProfileWithContext(aws.Context, *devicefarm.DeleteNetworkProfileInput, ...request.Option) (*devicefarm.DeleteNetworkProfileOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteNetworkProfileRequest(*devicefarm.DeleteNetworkProfileInput) (*request.Request, *devicefarm.DeleteNetworkProfileOutput) {
	return nil, nil
}

func (c *MockClient) DeleteProject(*devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteProjectWithContext(aws.Context, *devicefarm.DeleteProjectInput, ...request.Option) (*devicefarm.DeleteProjectOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteProjectRequest(*devicefarm.DeleteProjectInput) (*request.Request, *devicefarm.DeleteProjectOutput) {
	return nil, nil
}

func (c *MockClient) DeleteRemoteAccessSession(*devicefarm.DeleteRemoteAccessSessionInput) (*devicefarm.DeleteRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteRemoteAccessSessionWithContext(aws.Context, *devicefarm.DeleteRemoteAccessSessionInput, ...request.Option) (*devicefarm.DeleteRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteRemoteAccessSessionRequest(*devicefarm.DeleteRemoteAccessSessionInput) (*request.Request, *devicefarm.DeleteRemoteAccessSessionOutput) {
	return nil, nil
}

func (c *MockClient) DeleteRun(*devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteRunWithContext(aws.Context, *devicefarm.DeleteRunInput, ...request.Option) (*devicefarm.DeleteRunOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteRunRequest(*devicefarm.DeleteRunInput) (*request.Request, *devicefarm.DeleteRunOutput) {
	return nil, nil
}

func (c *MockClient) DeleteUpload(*devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteUploadWithContext(aws.Context, *devicefarm.DeleteUploadInput, ...request.Option) (*devicefarm.DeleteUploadOutput, error) {
	return nil, nil
}
func (c *MockClient) DeleteUploadRequest(*devicefarm.DeleteUploadInput) (*request.Request, *devicefarm.DeleteUploadOutput) {
	return nil, nil
}

func (c *MockClient) GetAccountSettings(*devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error) {
	return nil, nil
}
func (c *MockClient) GetAccountSettingsWithContext(aws.Context, *devicefarm.GetAccountSettingsInput, ...request.Option) (*devicefarm.GetAccountSettingsOutput, error) {
	return nil, nil
}
func (c *MockClient) GetAccountSettingsRequest(*devicefarm.GetAccountSettingsInput) (*request.Request, *devicefarm.GetAccountSettingsOutput) {
	return nil, nil
}

func (c *MockClient) GetDevice(*devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error) {
	return nil, nil
}
func (c *MockClient) GetDeviceWithContext(aws.Context, *devicefarm.GetDeviceInput, ...request.Option) (*devicefarm.GetDeviceOutput, error) {
	return nil, nil
}
func (c *MockClient) GetDeviceRequest(*devicefarm.GetDeviceInput) (*request.Request, *devicefarm.GetDeviceOutput) {
	return nil, nil
}

func (c *MockClient) GetDevicePool(*devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error) {
	return nil, nil
}
func (c *MockClient) GetDevicePoolWithContext(aws.Context, *devicefarm.GetDevicePoolInput, ...request.Option) (*devicefarm.GetDevicePoolOutput, error) {
	return nil, nil
}
func (c *MockClient) GetDevicePoolRequest(*devicefarm.GetDevicePoolInput) (*request.Request, *devicefarm.GetDevicePoolOutput) {
	return nil, nil
}

func (c *MockClient) GetDevicePoolCompatibility(*devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
	return nil, nil
}
func (c *MockClient) GetDevicePoolCompatibilityWithContext(aws.Context, *devicefarm.GetDevicePoolCompatibilityInput, ...request.Option) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
	return nil, nil
}
func (c *MockClient) GetDevicePoolCompatibilityRequest(*devicefarm.GetDevicePoolCompatibilityInput) (*request.Request, *devicefarm.GetDevicePoolCompatibilityOutput) {
	return nil, nil
}

func (c *MockClient) GetJob(*devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error) {
	return nil, nil
}
func (c *MockClient) GetJobWithContext(aws.Context, *devicefarm.GetJobInput, ...request.Option) (*devicefarm.GetJobOutput, error) {
	return nil, nil
}
func (c *MockClient) GetJobRequest(*devicefarm.GetJobInput) (*request.Request, *devicefarm.GetJobOutput) {
	return nil, nil
}

func (c *MockClient) GetNetworkProfile(*devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error) {
	return nil, nil
}
func (c *MockClient) GetNetworkProfileWithContext(aws.Context, *devicefarm.GetNetworkProfileInput, ...request.Option) (*devicefarm.GetNetworkProfileOutput, error) {
	return nil, nil
}
func (c *MockClient) GetNetworkProfileRequest(*devicefarm.GetNetworkProfileInput) (*request.Request, *devicefarm.GetNetworkProfileOutput) {
	return nil, nil
}

func (c *MockClient) GetOfferingStatus(*devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error) {
	return nil, nil
}
func (c *MockClient) GetOfferingStatusWithContext(aws.Context, *devicefarm.GetOfferingStatusInput, ...request.Option) (*devicefarm.GetOfferingStatusOutput, error) {
	return nil, nil
}
func (c *MockClient) GetOfferingStatusRequest(*devicefarm.GetOfferingStatusInput) (*request.Request, *devicefarm.GetOfferingStatusOutput) {
	return nil, nil
}

func (c *MockClient) GetOfferingStatusPages(*devicefarm.GetOfferingStatusInput, func(*devicefarm.GetOfferingStatusOutput, bool) bool) error {
	return nil
}
func (c *MockClient) GetOfferingStatusPagesWithContext(aws.Context, *devicefarm.GetOfferingStatusInput, func(*devicefarm.GetOfferingStatusOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) GetProject(*devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error) {
	return nil, nil
}
func (c *MockClient) GetProjectWithContext(aws.Context, *devicefarm.GetProjectInput, ...request.Option) (*devicefarm.GetProjectOutput, error) {
	return nil, nil
}
func (c *MockClient) GetProjectRequest(*devicefarm.GetProjectInput) (*request.Request, *devicefarm.GetProjectOutput) {
	return nil, nil
}

func (c *MockClient) GetRemoteAccessSession(*devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *MockClient) GetRemoteAccessSessionWithContext(aws.Context, *devicefarm.GetRemoteAccessSessionInput, ...request.Option) (*devicefarm.GetRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *MockClient) GetRemoteAccessSessionRequest(*devicefarm.GetRemoteAccessSessionInput) (*request.Request, *devicefarm.GetRemoteAccessSessionOutput) {
	return nil, nil
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
func (c *MockClient) GetRunWithContext(aws.Context, *devicefarm.GetRunInput, ...request.Option) (*devicefarm.GetRunOutput, error) {
	return nil, nil
}
func (c *MockClient) GetRunRequest(*devicefarm.GetRunInput) (*request.Request, *devicefarm.GetRunOutput) {
	return nil, nil
}

func (c *MockClient) GetSuite(*devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error) {
	return nil, nil
}
func (c *MockClient) GetSuiteWithContext(aws.Context, *devicefarm.GetSuiteInput, ...request.Option) (*devicefarm.GetSuiteOutput, error) {
	return nil, nil
}
func (c *MockClient) GetSuiteRequest(*devicefarm.GetSuiteInput) (*request.Request, *devicefarm.GetSuiteOutput) {
	return nil, nil
}

func (c *MockClient) GetTest(*devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error) {
	return nil, nil
}
func (c *MockClient) GetTestWithContext(aws.Context, *devicefarm.GetTestInput, ...request.Option) (*devicefarm.GetTestOutput, error) {
	return nil, nil
}
func (c *MockClient) GetTestRequest(*devicefarm.GetTestInput) (*request.Request, *devicefarm.GetTestOutput) {
	return nil, nil
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
func (c *MockClient) GetUploadWithContext(aws.Context, *devicefarm.GetUploadInput, ...request.Option) (*devicefarm.GetUploadOutput, error) {
	return nil, nil
}
func (c *MockClient) GetUploadRequest(*devicefarm.GetUploadInput) (*request.Request, *devicefarm.GetUploadOutput) {
	return nil, nil
}

func (c *MockClient) InstallToRemoteAccessSession(*devicefarm.InstallToRemoteAccessSessionInput) (*devicefarm.InstallToRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *MockClient) InstallToRemoteAccessSessionWithContext(aws.Context, *devicefarm.InstallToRemoteAccessSessionInput, ...request.Option) (*devicefarm.InstallToRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *MockClient) InstallToRemoteAccessSessionRequest(*devicefarm.InstallToRemoteAccessSessionInput) (*request.Request, *devicefarm.InstallToRemoteAccessSessionOutput) {
	return nil, nil
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
func (c *MockClient) ListArtifactsWithContext(aws.Context, *devicefarm.ListArtifactsInput, ...request.Option) (*devicefarm.ListArtifactsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListArtifactsRequest(*devicefarm.ListArtifactsInput) (*request.Request, *devicefarm.ListArtifactsOutput) {
	return nil, nil
}

func (c *MockClient) ListArtifactsPages(*devicefarm.ListArtifactsInput, func(*devicefarm.ListArtifactsOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListArtifactsPagesWithContext(aws.Context, *devicefarm.ListArtifactsInput, func(*devicefarm.ListArtifactsOutput, bool) bool, ...request.Option) error {
	return nil
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
func (c *MockClient) ListDevicePoolsWithContext(aws.Context, *devicefarm.ListDevicePoolsInput, ...request.Option) (*devicefarm.ListDevicePoolsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListDevicePoolsRequest(*devicefarm.ListDevicePoolsInput) (*request.Request, *devicefarm.ListDevicePoolsOutput) {
	return nil, nil
}

func (c *MockClient) ListDevicePoolsPages(*devicefarm.ListDevicePoolsInput, func(*devicefarm.ListDevicePoolsOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListDevicePoolsPagesWithContext(aws.Context, *devicefarm.ListDevicePoolsInput, func(*devicefarm.ListDevicePoolsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) ListDevices(*devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error) {
	return nil, nil
}
func (c *MockClient) ListDevicesWithContext(aws.Context, *devicefarm.ListDevicesInput, ...request.Option) (*devicefarm.ListDevicesOutput, error) {
	return nil, nil
}
func (c *MockClient) ListDevicesRequest(*devicefarm.ListDevicesInput) (*request.Request, *devicefarm.ListDevicesOutput) {
	return nil, nil
}

func (c *MockClient) ListDevicesPages(*devicefarm.ListDevicesInput, func(*devicefarm.ListDevicesOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListDevicesPagesWithContext(aws.Context, *devicefarm.ListDevicesInput, func(*devicefarm.ListDevicesOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) ListJobs(*devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error) {
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
func (c *MockClient) ListJobsWithContext(aws.Context, *devicefarm.ListJobsInput, ...request.Option) (*devicefarm.ListJobsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListJobsRequest(*devicefarm.ListJobsInput) (*request.Request, *devicefarm.ListJobsOutput) {
	return nil, nil
}

func (c *MockClient) ListJobsPages(*devicefarm.ListJobsInput, func(*devicefarm.ListJobsOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListJobsPagesWithContext(aws.Context, *devicefarm.ListJobsInput, func(*devicefarm.ListJobsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) ListNetworkProfiles(*devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error) {
	return nil, nil
}
func (c *MockClient) ListNetworkProfilesWithContext(aws.Context, *devicefarm.ListNetworkProfilesInput, ...request.Option) (*devicefarm.ListNetworkProfilesOutput, error) {
	return nil, nil
}
func (c *MockClient) ListNetworkProfilesRequest(*devicefarm.ListNetworkProfilesInput) (*request.Request, *devicefarm.ListNetworkProfilesOutput) {
	return nil, nil
}

func (c *MockClient) ListOfferingPromotions(*devicefarm.ListOfferingPromotionsInput) (*devicefarm.ListOfferingPromotionsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListOfferingPromotionsWithContext(aws.Context, *devicefarm.ListOfferingPromotionsInput, ...request.Option) (*devicefarm.ListOfferingPromotionsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListOfferingPromotionsRequest(*devicefarm.ListOfferingPromotionsInput) (*request.Request, *devicefarm.ListOfferingPromotionsOutput) {
	return nil, nil
}

func (c *MockClient) ListOfferingTransactions(*devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListOfferingTransactionsWithContext(aws.Context, *devicefarm.ListOfferingTransactionsInput, ...request.Option) (*devicefarm.ListOfferingTransactionsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListOfferingTransactionsRequest(*devicefarm.ListOfferingTransactionsInput) (*request.Request, *devicefarm.ListOfferingTransactionsOutput) {
	return nil, nil
}

func (c *MockClient) ListOfferingTransactionsPages(*devicefarm.ListOfferingTransactionsInput, func(*devicefarm.ListOfferingTransactionsOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListOfferingTransactionsPagesWithContext(aws.Context, *devicefarm.ListOfferingTransactionsInput, func(*devicefarm.ListOfferingTransactionsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) ListOfferings(*devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListOfferingsWithContext(aws.Context, *devicefarm.ListOfferingsInput, ...request.Option) (*devicefarm.ListOfferingsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListOfferingsRequest(*devicefarm.ListOfferingsInput) (*request.Request, *devicefarm.ListOfferingsOutput) {
	return nil, nil
}

func (c *MockClient) ListOfferingsPages(*devicefarm.ListOfferingsInput, func(*devicefarm.ListOfferingsOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListOfferingsPagesWithContext(aws.Context, *devicefarm.ListOfferingsInput, func(*devicefarm.ListOfferingsOutput, bool) bool, ...request.Option) error {
	return nil
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
func (c *MockClient) ListProjectsWithContext(aws.Context, *devicefarm.ListProjectsInput, ...request.Option) (*devicefarm.ListProjectsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListProjectsRequest(*devicefarm.ListProjectsInput) (*request.Request, *devicefarm.ListProjectsOutput) {
	return nil, nil
}

func (c *MockClient) ListProjectsPages(*devicefarm.ListProjectsInput, func(*devicefarm.ListProjectsOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListProjectsPagesWithContext(aws.Context, *devicefarm.ListProjectsInput, func(*devicefarm.ListProjectsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) ListRemoteAccessSessions(*devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListRemoteAccessSessionsWithContext(aws.Context, *devicefarm.ListRemoteAccessSessionsInput, ...request.Option) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListRemoteAccessSessionsRequest(*devicefarm.ListRemoteAccessSessionsInput) (*request.Request, *devicefarm.ListRemoteAccessSessionsOutput) {
	return nil, nil
}

func (c *MockClient) ListRuns(*devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListRunsWithContext(aws.Context, *devicefarm.ListRunsInput, ...request.Option) (*devicefarm.ListRunsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListRunsRequest(*devicefarm.ListRunsInput) (*request.Request, *devicefarm.ListRunsOutput) {
	return nil, nil
}

func (c *MockClient) ListRunsPages(*devicefarm.ListRunsInput, func(*devicefarm.ListRunsOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListRunsPagesWithContext(aws.Context, *devicefarm.ListRunsInput, func(*devicefarm.ListRunsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) ListSamples(*devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error) {
	return nil, nil
}
func (c *MockClient) ListSamplesWithContext(aws.Context, *devicefarm.ListSamplesInput, ...request.Option) (*devicefarm.ListSamplesOutput, error) {
	return nil, nil
}
func (c *MockClient) ListSamplesRequest(*devicefarm.ListSamplesInput) (*request.Request, *devicefarm.ListSamplesOutput) {
	return nil, nil
}

func (c *MockClient) ListSamplesPages(*devicefarm.ListSamplesInput, func(*devicefarm.ListSamplesOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListSamplesPagesWithContext(aws.Context, *devicefarm.ListSamplesInput, func(*devicefarm.ListSamplesOutput, bool) bool, ...request.Option) error {
	return nil
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
func (c *MockClient) ListSuitesWithContext(aws.Context, *devicefarm.ListSuitesInput, ...request.Option) (*devicefarm.ListSuitesOutput, error) {
	return nil, nil
}
func (c *MockClient) ListSuitesRequest(*devicefarm.ListSuitesInput) (*request.Request, *devicefarm.ListSuitesOutput) {
	return nil, nil
}

func (c *MockClient) ListSuitesPages(*devicefarm.ListSuitesInput, func(*devicefarm.ListSuitesOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListSuitesPagesWithContext(aws.Context, *devicefarm.ListSuitesInput, func(*devicefarm.ListSuitesOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) ListTests(input *devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error) {
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
					Arn: aws.String(""),
				},
			},
		}
	}
	return res, nil
}
func (c *MockClient) ListTestsWithContext(aws.Context, *devicefarm.ListTestsInput, ...request.Option) (*devicefarm.ListTestsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListTestsRequest(*devicefarm.ListTestsInput) (*request.Request, *devicefarm.ListTestsOutput) {
	return nil, nil
}

func (c *MockClient) ListTestsPages(*devicefarm.ListTestsInput, func(*devicefarm.ListTestsOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListTestsPagesWithContext(aws.Context, *devicefarm.ListTestsInput, func(*devicefarm.ListTestsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) ListUniqueProblems(*devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListUniqueProblemsWithContext(aws.Context, *devicefarm.ListUniqueProblemsInput, ...request.Option) (*devicefarm.ListUniqueProblemsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListUniqueProblemsRequest(*devicefarm.ListUniqueProblemsInput) (*request.Request, *devicefarm.ListUniqueProblemsOutput) {
	return nil, nil
}

func (c *MockClient) ListUniqueProblemsPages(*devicefarm.ListUniqueProblemsInput, func(*devicefarm.ListUniqueProblemsOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListUniqueProblemsPagesWithContext(aws.Context, *devicefarm.ListUniqueProblemsInput, func(*devicefarm.ListUniqueProblemsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) ListUploads(*devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListUploadsWithContext(aws.Context, *devicefarm.ListUploadsInput, ...request.Option) (*devicefarm.ListUploadsOutput, error) {
	return nil, nil
}
func (c *MockClient) ListUploadsRequest(*devicefarm.ListUploadsInput) (*request.Request, *devicefarm.ListUploadsOutput) {
	return nil, nil
}

func (c *MockClient) ListUploadsPages(*devicefarm.ListUploadsInput, func(*devicefarm.ListUploadsOutput, bool) bool) error {
	return nil
}
func (c *MockClient) ListUploadsPagesWithContext(aws.Context, *devicefarm.ListUploadsInput, func(*devicefarm.ListUploadsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (c *MockClient) PurchaseOffering(*devicefarm.PurchaseOfferingInput) (*devicefarm.PurchaseOfferingOutput, error) {
	return nil, nil
}
func (c *MockClient) PurchaseOfferingWithContext(aws.Context, *devicefarm.PurchaseOfferingInput, ...request.Option) (*devicefarm.PurchaseOfferingOutput, error) {
	return nil, nil
}
func (c *MockClient) PurchaseOfferingRequest(*devicefarm.PurchaseOfferingInput) (*request.Request, *devicefarm.PurchaseOfferingOutput) {
	return nil, nil
}

func (c *MockClient) RenewOffering(*devicefarm.RenewOfferingInput) (*devicefarm.RenewOfferingOutput, error) {
	return nil, nil
}
func (c *MockClient) RenewOfferingWithContext(aws.Context, *devicefarm.RenewOfferingInput, ...request.Option) (*devicefarm.RenewOfferingOutput, error) {
	return nil, nil
}
func (c *MockClient) RenewOfferingRequest(*devicefarm.RenewOfferingInput) (*request.Request, *devicefarm.RenewOfferingOutput) {
	return nil, nil
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
func (c *MockClient) ScheduleRunWithContext(aws.Context, *devicefarm.ScheduleRunInput, ...request.Option) (*devicefarm.ScheduleRunOutput, error) {
	return nil, nil
}
func (c *MockClient) ScheduleRunRequest(*devicefarm.ScheduleRunInput) (*request.Request, *devicefarm.ScheduleRunOutput) {
	return nil, nil
}

func (c *MockClient) StopRemoteAccessSession(*devicefarm.StopRemoteAccessSessionInput) (*devicefarm.StopRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *MockClient) StopRemoteAccessSessionWithContext(aws.Context, *devicefarm.StopRemoteAccessSessionInput, ...request.Option) (*devicefarm.StopRemoteAccessSessionOutput, error) {
	return nil, nil
}
func (c *MockClient) StopRemoteAccessSessionRequest(*devicefarm.StopRemoteAccessSessionInput) (*request.Request, *devicefarm.StopRemoteAccessSessionOutput) {
	return nil, nil
}

func (c *MockClient) StopRun(*devicefarm.StopRunInput) (*devicefarm.StopRunOutput, error) {
	return nil, nil
}
func (c *MockClient) StopRunWithContext(aws.Context, *devicefarm.StopRunInput, ...request.Option) (*devicefarm.StopRunOutput, error) {
	return nil, nil
}
func (c *MockClient) StopRunRequest(*devicefarm.StopRunInput) (*request.Request, *devicefarm.StopRunOutput) {
	return nil, nil
}

func (c *MockClient) UpdateDevicePool(*devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error) {
	return nil, nil
}
func (c *MockClient) UpdateDevicePoolWithContext(aws.Context, *devicefarm.UpdateDevicePoolInput, ...request.Option) (*devicefarm.UpdateDevicePoolOutput, error) {
	return nil, nil
}
func (c *MockClient) UpdateDevicePoolRequest(*devicefarm.UpdateDevicePoolInput) (*request.Request, *devicefarm.UpdateDevicePoolOutput) {
	return nil, nil
}

func (c *MockClient) UpdateNetworkProfile(*devicefarm.UpdateNetworkProfileInput) (*devicefarm.UpdateNetworkProfileOutput, error) {
	return nil, nil
}
func (c *MockClient) UpdateNetworkProfileWithContext(aws.Context, *devicefarm.UpdateNetworkProfileInput, ...request.Option) (*devicefarm.UpdateNetworkProfileOutput, error) {
	return nil, nil
}
func (c *MockClient) UpdateNetworkProfileRequest(*devicefarm.UpdateNetworkProfileInput) (*request.Request, *devicefarm.UpdateNetworkProfileOutput) {
	return nil, nil
}

func (c *MockClient) UpdateProject(*devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error) {
	return nil, nil
}
func (c *MockClient) UpdateProjectWithContext(aws.Context, *devicefarm.UpdateProjectInput, ...request.Option) (*devicefarm.UpdateProjectOutput, error) {
	return nil, nil
}
func (c *MockClient) UpdateProjectRequest(*devicefarm.UpdateProjectInput) (*request.Request, *devicefarm.UpdateProjectOutput) {
	return nil, nil
}
