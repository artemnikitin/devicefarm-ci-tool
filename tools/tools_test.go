package tools

import "testing"

func TestStringEndWith(t *testing.T) {
	cases := []struct {
		src, sub string
		res      bool
	}{
		{"dfd/", "/", true},
		{"dfd/", "d/", true},
		{"dfd/", "sdssd/", false},
		{"dfd/", "f", false},
		{"dfd", "/", false},
		{"dfd/", "x/", false},
		{"", "d/", false},
		{"dfd/", "", true},
		{"", "", true},
	}

	for _, v := range cases {
		result := StringEndsWith(v.src, v.sub)
		if result != v.res {
			t.Errorf("For string: %s end with: %s, actual: %v, expected: %v", v.src, v.sub, result, v.res)
		}
	}
}

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
