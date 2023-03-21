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

	fmt.Println(datetimeutil.GetTime())
	fmt.Println(datetimeutil.GetTimeF(datetimeutil.F_hhmmss))
	fmt.Println(datetimeutil.GetTimeF(datetimeutil.F_hhmmss_chinese))
	t2, _ := datetimeutil.GetTimeFromF(datetimeutil.F_hhmmss_colon, "16:04:04", datetimeutil.F_hhmmss_chinese)
	fmt.Printf("16:04:04 -> %s\n", t2)
	t3, _ := datetimeutil.AddDuration(datetimeutil.F_hhmmss_colon, "16:04:04", time.Hour*1)
	fmt.Printf("16:04:04 -- +1 hour --> %s\n", t3)
}
