package datetimeutil_test

import (
	"testing"
	"time"

	"github.com/realjf/datetimeutil"
)

func TestDateTime(t *testing.T) {
	cases := map[string]struct {
		p func(string)
	}{
		"GetDate": {
			p: func(name string) {
				t.Logf("%s: %s", name, datetimeutil.GetDate())
			},
		},
		"GetDateF": {
			p: func(name string) {
				t.Logf("%s: %s", name, datetimeutil.GetDateF(datetimeutil.F_YYYYMMDD_chinese))
			},
		},
		"AddDays": {
			p: func(name string) {
				tm, err := datetimeutil.AddDays(datetimeutil.F_YYYYMMDD, "20230321", 3)
				if err != nil {
					t.Errorf("%s: %v", name, err)
				}
				t.Logf("%s: %s", name, tm)
			},
		},
		"GetPosFromF": {
			p: func(name string) {
				tm, err := datetimeutil.GetPosFromF(datetimeutil.F_YYYYMMDD, "20230321", "20230325", time.Duration(1*24)*time.Hour)
				if err != nil {
					t.Errorf("%s: %v", name, err)
				}
				t.Logf("%s: %d", name, tm)
			},
		},
	}

	for name, ts := range cases {
		t.Run(name, func(t *testing.T) {
			ts.p(name)
		})
	}
}
