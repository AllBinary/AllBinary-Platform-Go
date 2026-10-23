/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
// Generated Code Do Not Modify

package allbinary

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
    "time"
    "unicode"
)

var _ = errors.New
var _ = fmt.Sprint
var _ = math.Abs
var _ = rand.Float64
var _ = os.Exit
var _ = reflect.TypeOf
var _ = runtime.GC
var _ = strconv.Itoa
var _ = strings.Contains
var _ = time.Now
var _ = unicode.IsDigit

type TimeUtil struct {
}

var TimeUtil_Instance *TimeUtil = NewTimeUtil()
var TimeUtil_Calendar *Calendar = Calendar_GetInstance()

func NewTimeUtil() *TimeUtil {
    this := &TimeUtil{}
    return this
}


func TimeUtil_GetInstance() *TimeUtil {
        return TimeUtil_Instance
}


func (this *TimeUtil) SetCalendar_1(calendar *Calendar, yearString string, monthString string, dayString string, hourString string) {
        var year int = int(int(yearString)); _ = year
        var month int = int(int(monthString)); _ = month
        var day int = int(int(dayString)); _ = day
        var hour int = int(int(hourString)); _ = hour
        calendar.Set(year, month, day, hour, 0)
}

// Java overloads of TimeUtil.setCalendar
func (this *TimeUtil) SetCalendar(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 5 && Java_IsArg[*Calendar](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) && Java_IsArg[string](args[4], pass) {
            this.SetCalendar_1(Java_Arg[*Calendar](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]), Java_Arg[string](args[4]))
            return
        }
        if len(args) == 7 && Java_IsArg[*Calendar](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) && Java_IsArg[string](args[4], pass) && Java_IsArg[string](args[5], pass) && Java_IsArg[string](args[6], pass) {
            this.SetCalendar_2(Java_Arg[*Calendar](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]), Java_Arg[string](args[4]), Java_Arg[string](args[5]), Java_Arg[string](args[6]))
            return
        }
    }
    panic("No Java overload of TimeUtil.setCalendar for the arguments")
}


func (this *TimeUtil) SetCalendar_2(calendar *Calendar, yearString string, monthString string, dayString string, hourString string, minuteString string, secondString string) {
        var year int = int(int(yearString)); _ = year
        var month int = int(int(monthString)); _ = month
        var day int = int(int(dayString)); _ = day
        var hour int = int(int(hourString)); _ = hour
        var minute int = int(int(minuteString)); _ = minute
        var second int = int(int(secondString)); _ = second
        calendar.Set(year, month, day, hour, minute, second)
}


func (this *TimeUtil) SetExpirationCalendar(calendar *Calendar, yearsString string, monthsString string, daysString string, hoursString string, minutesString string, secondsString string) {
        var year int = calendar.Get(Calendar_YEAR) + int(int(yearsString)); _ = year
        var month int = calendar.Get(Calendar_MONTH) + int(int(monthsString)); _ = month
        var day int = calendar.Get(Calendar_DAY_OF_MONTH) + int(int(daysString)); _ = day
        var hour int = calendar.Get(Calendar_HOUR) + int(int(hoursString)); _ = hour
        var minute int = calendar.Get(Calendar_MINUTE) + int(int(minutesString)); _ = minute
        var second int = calendar.Get(Calendar_SECOND) + int(int(secondsString)); _ = second
        calendar.Set(year, month, day, hour, minute, second)
}


func (this *TimeUtil) GetTotalTime(yearString string, monthString string, dayString string, hourString string, minuteString string, secondString string) int64 {
        var time int64 = TimeUtil_Calendar.GetTimeInMillis(); _ = time
        this.SetExpirationCalendar(TimeUtil_Calendar, yearString, monthString, dayString, hourString, minuteString, secondString)
        var timeDelta int64 = TimeUtil_Calendar.GetTimeInMillis(); _ = timeDelta
        return timeDelta - time
}

