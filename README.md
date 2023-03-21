# datetime

datetime util in go(日期时间工具集)

### Example

```go

import (
 "fmt"
 "time"

 datetimeutil "github.com/realjf/datetime"
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
```

the follow is output:

```sh
20230321
2023/03/21
2023-03-21
2023年03月21日
2023/03/21 18:20
2023-03-21 18:20
2023年03月21日 18:20
2023/03/03 -> 2023/03/07
2023-03-03 -> 1677801600
18:20:39
18:20:39
18时20分39秒
16:04:04 -> 16时04分04秒
16:04:04 -- +1 hour --> 17:04:04
```
