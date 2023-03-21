package datetimeutil_test

import (
	"testing"
	"time"

	datetimeutil "github.com/realjf/datetime"
)

func TestTime(t *testing.T) {
	cases := map[string]struct {
		p func(string)
	}{
		"GetTime": {
			p: func(name string) {
				t.Logf("%s: %s", name, datetimeutil.GetTime())
			},
		},
		"GetTimeF": {
			p: func(name string) {
				t.Logf("%s: %s", name, datetimeutil.GetTimeF(datetimeutil.F_hhmmss_chinese))
			},
		},
		"GetTimeFrom": {
			p: func(name string) {
				tm, err := datetimeutil.GetTimeFromF(datetimeutil.F_YYYYMMDDhhmmss, "20230321040506", datetimeutil.F_hhmmss_chinese)
				if err != nil {
					t.Errorf("%s: %v", name, err)
				}
				t.Logf("%s: %s", name, tm)
			},
		},
		"AddDuration": {
			p: func(name string) {
				tm, err := datetimeutil.AddDuration(datetimeutil.F_hhmmss_colon, "18:01:01", time.Duration(30*time.Minute))
				if err != nil {
					t.Errorf("%s: %v", name, err)
				}
				t.Logf("%s: %s", name, tm)
			},
		},
	}

	for name, ts := range cases {
		t.Run(name, func(t *testing.T) {
			ts.p(name)
		})
	}
}
