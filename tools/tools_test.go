package tools

import (
	"testing"
)

func TestGetFileName(t *testing.T) {
	cases := []struct{ path, res string }{
		{"dfd.kkk", "dfd.kkk"},
		{"/dfdfd/www/dfd.kkk", "dfd.kkk"},
		{"", ""},
	}

	for _, v := range cases {
		result := GetFileName(v.path)
		if result != v.res {
			t.Errorf("For path: %s actual filename: %s, expected: %s", v.path, result, v.res)
		}
	}
}

func TestGenerateReportURL(t *testing.T) {
	cases := []struct{ arn, res string }{
		{"", ""},
		{
			"arn:aws:devicefarm:us-west-2:000000000:run:34t334-6056-4685-5y55-3h54h4g3r/f3f3f3-3130-09b7-acbw-6c23a92ab391",
			"https://us-west-2.console.aws.amazon.com/devicefarm/home?region=us-west-2#/projects/34t334-6056-4685-5y55-3h54h4g3r/runs/f3f3f3-3130-09b7-acbw-6c23a92ab391",
		},
		{
			"arn:aws:devicefarm:us-west-2:000000000:run:/f3f3f3-3130-09b7-acbw-6c23a92ab391",
			"https://us-west-2.console.aws.amazon.com/devicefarm/home?region=us-west-2#/projects//runs/f3f3f3-3130-09b7-acbw-6c23a92ab391",
		},
		{
			"arn:aws:devicefarm:us-west-2:000000000:run:34t334-6056-4685-5y55-3h54h4g3r/",
			"https://us-west-2.console.aws.amazon.com/devicefarm/home?region=us-west-2#/projects/34t334-6056-4685-5y55-3h54h4g3r/runs/",
		},
		{
			"arn:aws:devicefarm:us-west-2:000000000:run:/",
			"https://us-west-2.console.aws.amazon.com/devicefarm/home?region=us-west-2#/projects//runs/",
		},
		{
			"arn:aws:devicefarm:us-west-2:000000000",
			"",
		},
	}

	for _, v := range cases {
		result := GenerateReportURL(v.arn)
		if result != v.res {
			t.Errorf("For ARN: %s\n actual URL: %s,\n expected: %s \n", v.arn, result, v.res)
		}
	}
}

func TestRandom(t *testing.T) {
	i := Random(1, 10)
	if i < 1 || i > 10 {
		t.Error("Random int should be in given range")
	}
}
