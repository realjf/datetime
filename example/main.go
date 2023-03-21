package main

import (
	"fmt"
	"time"

	"github.com/realjf/datetimeutil"
)

func main() {

	fmt.Println(datetimeutil.GetDate())
	fmt.Println(datetimeutil.GetDateF(datetimeutil.F_YYYYMMDD_slash))
	fmt.Println(datetimeutil.GetDateF(datetimeutil.F_YYYYMMDD_hyphen))
	fmt.Println(datetimeutil.GetDateF(datetimeutil.F_YYYYMMDD_chinese))
	fmt.Println(datetimeutil.GetDateF(datetimeutil.F_YYYYMMDDhhmm_slash))
	fmt.Println(datetimeutil.GetDateF(datetimeutil.F_YYYYMMDDhhmm_hyphen))
	fmt.Println(datetimeutil.GetDateF(datetimeutil.F_YYYYMMDDhhmm_chinese))
	t1, _ := datetimeutil.AddDays(datetimeutil.F_YYYYMMDD_slash, "2023/03/03", 4)
	fmt.Printf("2023/03/03 -> %s\n", t1)
	u1, _ := datetimeutil.ParseDate2Unix(datetimeutil.F_YYYYMMDD_hyphen, "2023-03-03")
	fmt.Printf("2023-03-03 -> %d\n", u1)
	t2, _ := datetimeutil.ParseDate2Time(datetimeutil.F_YYYYMMDD, "20230303")
	fmt.Printf("2023-03-03 -> %v\n", t2)
	fmt.Println(datetimeutil.GetTime())
	fmt.Println(datetimeutil.GetTimeF(datetimeutil.F_hhmmss))
	fmt.Println(datetimeutil.GetTimeF(datetimeutil.F_hhmmss_chinese))
	t4, _ := datetimeutil.GetTimeFromF(datetimeutil.F_hhmmss_colon, "16:04:04", datetimeutil.F_hhmmss_chinese)
	fmt.Printf("16:04:04 -> %s\n", t4)
	t5, _ := datetimeutil.AddDuration(datetimeutil.F_hhmmss_colon, "16:04:04", time.Hour*1)
	fmt.Printf("16:04:04 -- +1 hour --> %s\n", t5)
	t6, _ := datetimeutil.GetPosFromF(datetimeutil.F_YYYYMMDD, "20230321", "20230325", time.Duration(1*24)*time.Hour)
	fmt.Printf("20230325 distance to 20230321: %d days\n", t6+1)
	t7, _ := datetimeutil.Count(datetimeutil.F_YYYYMMDD, "20230321", "20230325", time.Duration(1*24)*time.Hour)
	fmt.Printf("from 20230325 to 20230321: %d days\n", t7)
}
