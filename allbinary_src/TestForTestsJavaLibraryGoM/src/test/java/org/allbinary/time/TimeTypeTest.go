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

type TimeTypeTest struct {
}


func NewTimeTypeTest() *TimeTypeTest {
    this := &TimeTypeTest{}
    return this
}


func TimeTypeTest_MillisForHourLocal(hour int) int64 {
        var cal *Calendar = Calendar_GetInstance(); _ = cal
        cal.Set(Calendar_HOUR_OF_DAY, hour)
        cal.Set(Calendar_MINUTE, 0)
        cal.Set(Calendar_SECOND, 0)
        cal.Set(Calendar_MILLISECOND, 0)
        return cal.GetTimeInMillis()
}


func (this *TimeTypeTest) IsNight_returnsTrue_for_hours_less_than_6() {
        /* Otherwise - statement - ForEachStmt */
}


func (this *TimeTypeTest) IsNight_returnsTrue_for_hours_greater_than_18() {
        /* Otherwise - statement - ForEachStmt */
}


func (this *TimeTypeTest) IsNight_returnsFalse_for_boundary_hours_6_and_18() {
        /* Otherwise - statement - ForEachStmt */
}


func (this *TimeTypeTest) GetNightOrDay_returns_expected_label() {
        var nightUtil *TimeTypeUtil = NewTestableTimeTypeUtil(0); _ = nightUtil
        Assertions_AssertEquals(nightUtil.NIGHT, nightUtil.GetNightOrDay())
        var dayUtil *TimeTypeUtil = NewTestableTimeTypeUtil(12); _ = dayUtil
        Assertions_AssertEquals(dayUtil.DAY, dayUtil.GetNightOrDay())
}


func (this *TimeTypeTest) GetHourOfDay_long_converts_epoch_to_local_hour() {
        var util *TimeTypeUtil = NewTimeTypeUtil(); _ = util
        /* Otherwise - statement - ForEachStmt */
}


func (this *TimeTypeTest) IsNight_long_expected_ranges() {
        var util *TimeTypeUtil = NewTimeTypeUtil(); _ = util
        /* Otherwise - statement - ForEachStmt */
        /* Otherwise - statement - ForEachStmt */
}


func (this *TimeTypeTest) GetNightOrDay_long_consistency() {
        var util *TimeTypeUtil = NewTimeTypeUtil(); _ = util
        var nightMillis int64 = TimeTypeTest_MillisForHourLocal(0); _ = nightMillis
        var dayMillis int64 = TimeTypeTest_MillisForHourLocal(12); _ = dayMillis
        Assertions_AssertEquals(util.NIGHT, util.GetNightOrDay(nightMillis))
        Assertions_AssertEquals(util.DAY, util.GetNightOrDay(dayMillis))
}


func (this *TimeTypeTest) Boundaries_day_for_long_based_api() {
        var util *TimeTypeUtil = NewTimeTypeUtil(); _ = util
        /* Otherwise - statement - ForEachStmt */
}


func (this *TimeTypeTest) Constants_accessible_and_stable() {
        var util *TimeTypeUtil = NewTimeTypeUtil(); _ = util
        Assertions_AssertEquals("Night", util.NIGHT)
        Assertions_AssertEquals("Day", util.DAY)
}


func (this *TimeTypeTest) GetInstance_returns_same_singleton() {
        var a *TimeTypeUtil = TimeTypeUtil_GetInstance(); _ = a
        var b *TimeTypeUtil = TimeTypeUtil_GetInstance(); _ = b
        Assertions_AssertSame(a, b)
}

type TestableTimeTypeUtil struct {
    TimeTypeUtil
    Hour int
}


func (this *TestableTimeTypeUtil) initFields() {
        
}
    

    func NewTimeTypeTest(hour int) *TimeTypeTest {
        this := &TimeTypeTest{}
        this.TimeTypeUtil = *NewTimeTypeUtil()
        this.initFields()
            this.Hour = hour
        return this
    }
    

    func (this *TimeTypeTest) GetHourOfDay() int {
            return this.Hour
    }

