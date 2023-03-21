package datetimeutil

import "time"

const (
	F_hhmmss         DTF = "150405"
	F_hh_colon       DTF = "15"
	F_hhmm_colon     DTF = "15:04"
	F_hhmmss_colon   DTF = "15:04:05"
	F_hh_chinese     DTF = "15时"
	F_hhmm_chinese   DTF = "15时04分"
	F_hhmmss_chinese DTF = "15时04分05秒"
)

// return 15:04:05
func GetTime() string {
	return time.Now().Format(F_hhmmss_colon.String())
}

// return f format
func GetTimeF(f DTF) string {
	return time.Now().Format(f.String())
}

// return f2 format
func GetTimeFromF(f DTF, date string, f2 DTF) (string, error) {
	t, err := time.Parse(f.String(), date)
	if err != nil {
		return "", err
	}
	return t.Format(f2.String()), nil
}

// return f format
func AddDuration(f DTF, datetime string, d time.Duration) (string, error) {
	t, err := time.Parse(f.String(), datetime)
	if err != nil {
		return "", err
	}
	return t.Add(d).Format(f.String()), nil
}
