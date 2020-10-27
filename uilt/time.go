package uilt

import "time"

const TIME_FORMAT_ONE  = "2006年01月02日 15:04:05"
const TIME_FORMAT_TWO  = "2006/01/02 15:04:05"
const TIME_FORMAT_THREE  = "2006.01.02 15:04:05"
const TIME_FORMAT_FOUR  = "2006-01-02 15:04:05"
/**
根据所需要的格式，生成相应的格式的当前日期
 */

func  Time_1(format string)string{
	return time.Now().Format(format)

}
func Time_2(sec int64,nsec int64,format string) string{
	return time.Unix(sec,nsec).Format(format)
}