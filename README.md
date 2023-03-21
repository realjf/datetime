# datetimeutil

datetime util in go(日期时间工具集)

### Example

```go

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
```

the follow is output:

```sh
20230321
2023/03/21
2023-03-21
2023年03月21日
2023/03/21 23:09
2023-03-21 23:09
2023年03月21日 23:09
2023/03/03 -> 2023/03/07
2023-03-03 -> 1677801600
2023-03-03 -> 2023-03-03 00:00:00 +0000 UTC
23:09:10
230910
23时09分10秒
16:04:04 -> 16时04分04秒
16:04:04 -- +1 hour --> 17:04:04
20230325 distance to 20230321: 5 days
from 20230325 to 20230321: 5 days
```
