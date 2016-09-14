package tools

import "testing"

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
