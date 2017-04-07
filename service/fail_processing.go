package service

import (
	"sync"

	"github.com/artemnikitin/devicefarm-ci-tool/model"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

func populateResult(tests chan *model.FailedTest, client *devicefarm.DeviceFarm) []*model.FailedTest {
	var m sync.Mutex
	var wg sync.WaitGroup
	var result []*model.FailedTest

	wg.Add(len(tests))
	for v := range tests {
		go func(v *model.FailedTest) {
			artifacts := getArtifactsForTest(client, v.ARN)

			for i := 0; i < len(artifacts); i++ {
				if *artifacts[i].Type == devicefarm.ArtifactTypeDeviceLog {
					v.LogURL = *artifacts[i].Url
				}
				if *artifacts[i].Type == "VIDEO" {
					v.VideoURL = *artifacts[i].Url
				}
			}

			m.Lock()
			result = append(result, v)
			m.Unlock()

			wg.Done()
		}(v)
	}
	wg.Wait()

	return result
}

func getListOfFailedTestsFromSuite(client *devicefarm.DeviceFarm, suitesArn []string, device string, os string) chan *model.FailedTest {
	testch := make(chan *model.FailedTest, 100000)
	var wg sync.WaitGroup
	wg.Add(len(suitesArn))
	for i := 0; i < len(suitesArn); i++ {
		go func(i int, testch chan *model.FailedTest) {
			tests := getListOfTestForSuite(client, suitesArn[i])

			for i := 0; i < len(tests); i++ {
				if *tests[i].Result == devicefarm.ExecutionResultFailed {
					test := &model.FailedTest{
						ARN:     *tests[i].Arn,
						Message: *tests[i].Message,
						Device:  device,
						OS:      os,
					}
					testch <- test
				}
			}

			wg.Done()
		}(i, testch)
	}
	wg.Wait()
	close(testch)

	return testch
}

func getListOfTestArnFromSuite(suites []*devicefarm.Suite) []string {
	var suitesArn []string
	for i := 0; i < len(suites); i++ {
		if *suites[i].Result == devicefarm.ExecutionResultFailed {
			suitesArn = append(suitesArn, *suites[i].Arn)
		}
	}
	return suitesArn
}
