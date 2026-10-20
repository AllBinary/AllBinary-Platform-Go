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


func (this *TimeTypeUtil) IsNight() bool {
        var hoursOfDay int = this.GetHourOfDay(); _ = hoursOfDay
        return hoursOfDay < 6 || hoursOfDay > 18
}


func (this *TimeTypeUtil) GetHourOfDay() int {
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        return calendar.Get(Calendar_HOUR_OF_DAY)
}


func (this *TimeTypeUtil) GetNightOrDay() string {
        return func() string { if this.IsNight() { return this.NIGHT }; return this.DAY }()
}


// Java overloaded method: Go needs a different name
func (this *TimeTypeUtil) GetHourOfDay(timeInMillis int64) int {
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        calendar.SetTimeInMillis(timeInMillis)
        return calendar.Get(Calendar_HOUR_OF_DAY)
}


// Java overloaded method: Go needs a different name
func (this *TimeTypeUtil) IsNight(timeInMillis int64) bool {
        var hoursOfDay int = this.GetHourOfDay(timeInMillis); _ = hoursOfDay
        return hoursOfDay < 6 || hoursOfDay > 18
}


// Java overloaded method: Go needs a different name
func (this *TimeTypeUtil) GetNightOrDay(timeInMillis int64) string {
        return func() string { if this.IsNight(timeInMillis) { return this.NIGHT }; return this.DAY }()
}

