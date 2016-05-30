package utils

import (
	"time"
	"strconv"
)

//处理文章时间
func SinceTime(RawTime time.Time) string {
	times := time.Since(RawTime).Hours()
	if times < 1{
		times = time.Since(RawTime).Minutes()
		x := int64(times)
		return strconv.FormatInt(x,10) + "分钟前"
	}
	if times > 24{
		x := int64(times) / 24
		return strconv.FormatInt(x,10) + "天前"
	}
	x := int64(times)
	return strconv.FormatInt(x,10) + "小时前"

}