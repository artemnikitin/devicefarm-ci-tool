package service

import (
	"testing"

	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

func TestBla(t *testing.T) {
	config := aws.NewConfig()
	config.WithCredentials(credentials.NewEnvCredentials())
	config.WithRegion("us-west-2")
	ses, _ := session.NewSession(config)
	client := devicefarm.New(ses)

	/*svc := &DeviceFarmRun{
		Client:  client,
		Config:  &model.RunConfig{},
		Project: "uikit",
	}*/

	runArn := "arn:aws:devicefarm:us-west-2:091463382595:run:497a048d-6056-4685-9b96-be51ae4d3ba5/4135368c-1b3e-4e1a-a484-1152d1db9330"
	jobs := getListOfJobsForRun(client, runArn)
	fmt.Println("Number of jobs:", len(jobs))

	for _, v := range jobs {
		suites := getListOfSuitesForJob(client, *v.Arn)
		fmt.Println("Number of suites:", len(suites))

		for _, k := range suites {
			tests := getListOfTestForSuite(client, *k.Arn)
			for i := range tests {
				//fmt.Println(tests[i].String())
				if *tests[i].Result == devicefarm.ExecutionResultPassed && *tests[i].Name != "Setup Test" && *tests[i].Name != "Teardown Test" {
					fmt.Println("Number of tests:", len(tests))
				}
			}
		}

	}

	fmt.Println("--------------------")

	runArn = "arn:aws:devicefarm:us-west-2:091463382595:run:497a048d-6056-4685-9b96-be51ae4d3ba5/5ca82af1-18b8-4f78-b179-f4e060912808"
	jobs = getListOfJobsForRun(client, runArn)
	fmt.Println("Number of jobs:", len(jobs))

	for _, v := range jobs {
		suites := getListOfSuitesForJob(client, *v.Arn)
		fmt.Println("Number of suites:", len(suites))

		for _, k := range suites {
			tests := getListOfTestForSuite(client, *k.Arn)
			for i := range tests {
				if *tests[i].Result == devicefarm.ExecutionResultPassed && *tests[i].Name != "Setup Test" && *tests[i].Name != "Teardown Test" {
					fmt.Println("Number of tests:", len(tests))
				}
			}
		}

	}

	fmt.Println("--------------------")

	runArn = "arn:aws:devicefarm:us-west-2:091463382595:run:497a048d-6056-4685-9b96-be51ae4d3ba5/a707824a-9c13-4dcb-99a1-d65fcce62e79"
	jobs = getListOfJobsForRun(client, runArn)
	fmt.Println("Number of jobs:", len(jobs))

	for _, v := range jobs {
		suites := getListOfSuitesForJob(client, *v.Arn)
		fmt.Println("Number of suites:", len(suites))

		for _, k := range suites {
			tests := getListOfTestForSuite(client, *k.Arn)
			for i := range tests {
				if *tests[i].Result == devicefarm.ExecutionResultPassed && *tests[i].Name != "Setup Test" && *tests[i].Name != "Teardown Test" {
					fmt.Println("Number of tests:", len(tests))
				}
			}
		}

	}

}
