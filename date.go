package datetimeutil

import "time"

type DTF string

const (
	F_Layout                 DTF = "2006-01-02T15:04:05.000Z"
	F_YYYY                   DTF = "2006"
	F_YYYYMM                 DTF = "200601"
	F_YYYYMMDD               DTF = "20060102"
	F_YYYYMMDDhh             DTF = "2006010215"
	F_YYYYMMDDhhmm           DTF = "200601021504"
	F_YYYYMMDDhhmmss         DTF = "20060102150405"
	F_YYYYMM_slash           DTF = "2006/01"
	F_YYYYMMDD_slash         DTF = "2006/01/02"
	F_YYYYMMDDhhmm_slash     DTF = "2006/01/02 15:04"
	F_YYYYMMDDhhmmss_slash   DTF = "2006/01/02 15:04:05"
	F_YYYYMM_hyphen          DTF = "2006-01"
	F_YYYYMMDD_hyphen        DTF = "2006-01-02"
	F_YYYYMMDDhhmm_hyphen    DTF = "2006-01-02 15:04"
	F_YYYYMMDDhhmmss_hyphen  DTF = "2006-01/02 15:04:05"
	F_YYYY_chinese           DTF = "2006年"
	F_YYYYMM_chinese         DTF = "2006年01月"
	F_YYYYMMDD_chinese       DTF = "2006年01月02日"
	F_YYYYMMDDhhmm_chinese   DTF = "2006年01月02日 15:04"
	F_YYYYMMDDhhmmss_chinese DTF = "2006年01月02日 15:04:05"
)

func (f DTF) String() string {
	return string(f)
}

// return YYYYMMDD
func GetDate() string {
	return time.Now().Format(F_YYYYMMDD.String())
}

// return f format
func GetDateF(f DTF) string {
	return time.Now().Format(f.String())
}

// @param f: input date format
// @return YYYYMMDD
func ParseDate(f DTF, date string) (string, error) {
	t, err := time.Parse(f.String(), date)
	if err != nil {
		return "", err
	}
	return t.Format(F_YYYYMMDD.String()), nil
}

// return time.Time
func ParseDate2Time(f DTF, date string) (time.Time, error) {
	return time.Parse(f.String(), date)
}

// @param f: input date format
// @return f2 format
func ParseDateF(f DTF, date string, f2 DTF) (string, error) {
	t, err := time.Parse(f.String(), date)
	if err != nil {
		return "", err
	}
	return t.Format(f2.String()), nil
}

// @param f: input date format
// @return unix timestamp
func ParseDate2Unix(f DTF, date string) (int64, error) {
	t, err := time.Parse(f.String(), date)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// @param f: output date format
func ParseUnix2Date(t int64, f DTF) string {
	return time.Unix(t, 0).Format(f.String())
}

// return time.Time
func ParseUnix2Time(t int64) time.Time {
	return time.Unix(t, 0)
}

func AddDays(f DTF, date string, days int) (string, error) {
	t, err := time.Parse(f.String(), date)
	if err != nil {
		return "", err
	}
	t = t.Add(time.Duration(days*86400) * time.Second)
	return t.Format(f.String()), nil
}

// return time.Time
func AddDays2Time(f DTF, date string, days int) (time.Time, error) {
	t, err := time.Parse(f.String(), date)
	if err != nil {
		return t, err
	}
	t = t.Add(time.Duration(days*86400) * time.Second)
	return t, nil
}
