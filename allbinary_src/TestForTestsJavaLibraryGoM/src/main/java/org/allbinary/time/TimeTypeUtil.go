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

type TimeTypeUtil struct {
    NIGHT string
    DAY string
}

var TimeTypeUtil_Instance *TimeTypeUtil = NewTimeTypeUtil()

func (this *TimeTypeUtil) initFields() {
    this.NIGHT = "Night"
    this.DAY = "Day"
}

func NewTimeTypeUtil() *TimeTypeUtil {
    this := &TimeTypeUtil{}
    this.initFields()
    return this
}


func TimeTypeUtil_GetInstance() *TimeTypeUtil {
        return TimeTypeUtil_Instance
}


func (this *TimeTypeUtil) IsNight_1() bool {
        var hoursOfDay int = this.GetHourOfDay(); _ = hoursOfDay
        return hoursOfDay < 6 || hoursOfDay > 18
}

// Java overloads of TimeTypeUtil.isNight
func (this *TimeTypeUtil) IsNight(args ...any) bool {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.IsNight_1()
        }
        if len(args) == 1 && Java_IsArg[int64](args[0], pass) {
            return this.IsNight_2(Java_Arg[int64](args[0]))
        }
    }
    panic("No Java overload of TimeTypeUtil.isNight for the arguments")
}


func (this *TimeTypeUtil) GetHourOfDay_1() int {
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        return calendar.Get(Calendar_HOUR_OF_DAY)
}

// Java overloads of TimeTypeUtil.getHourOfDay
func (this *TimeTypeUtil) GetHourOfDay(args ...any) int {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.GetHourOfDay_1()
        }
        if len(args) == 1 && Java_IsArg[int64](args[0], pass) {
            return this.GetHourOfDay_2(Java_Arg[int64](args[0]))
        }
    }
    panic("No Java overload of TimeTypeUtil.getHourOfDay for the arguments")
}


func (this *TimeTypeUtil) GetNightOrDay_1() string {
        return func() string { if this.IsNight() { return this.NIGHT }; return this.DAY }()
}

// Java overloads of TimeTypeUtil.getNightOrDay
func (this *TimeTypeUtil) GetNightOrDay(args ...any) string {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.GetNightOrDay_1()
        }
        if len(args) == 1 && Java_IsArg[int64](args[0], pass) {
            return this.GetNightOrDay_2(Java_Arg[int64](args[0]))
        }
    }
    panic("No Java overload of TimeTypeUtil.getNightOrDay for the arguments")
}


func (this *TimeTypeUtil) GetHourOfDay_2(timeInMillis int64) int {
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        calendar.SetTimeInMillis(timeInMillis)
        return calendar.Get(Calendar_HOUR_OF_DAY)
}


func (this *TimeTypeUtil) IsNight_2(timeInMillis int64) bool {
        var hoursOfDay int = this.GetHourOfDay(timeInMillis); _ = hoursOfDay
        return hoursOfDay < 6 || hoursOfDay > 18
}


func (this *TimeTypeUtil) GetNightOrDay_2(timeInMillis int64) string {
        return func() string { if this.IsNight(timeInMillis) { return this.NIGHT }; return this.DAY }()
}

