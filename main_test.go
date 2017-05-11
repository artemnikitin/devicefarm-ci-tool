package main

import (
	"testing"

	"github.com/artemnikitin/devicefarm-ci-tool/model"
	"github.com/artemnikitin/devicefarm-ci-tool/service"
)

func TestRunJob(t *testing.T) {
	cases := []struct {
		name    string
		project string
		file    string
		config  *model.RunConfig
		failed  []*model.FailedTest
		pass    bool
	}{
		{
			name:    "Success run",
			project: "test",
			file:    "main_test.go",
			config:  &model.RunConfig{},
			failed:  []*model.FailedTest{},
			pass:    true,
		},
		{
			name:    "Run with failed tests",
			project: "test",
			file:    "main_test.go",
			config:  &model.RunConfig{},
			failed:  []*model.FailedTest{{}, {}},
			pass:    false,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			client := &service.MockClient{}

			client.FakeServer = service.CreateFakeServer()
			defer client.FakeServer.Close()

			*project = v.project
			*appPath = v.file
			*wait = true
			client.Failed = !v.pass
			client.UploadTest = true

			failed, pass := runJob(client, v.config)
			if pass != v.pass {
				t.Errorf("Test: %s\nTest result expected: %t, actual: %t", v.name, v.pass, pass)
			}
			if len(failed) != len(v.failed) {
				t.Errorf("Test: %s\nNumber of failed tests expected: %d, actual: %d", v.name, len(v.failed), len(failed))
			}
		})
	}
}
