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

type StupidTimer struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    NullUtil *NullUtil
    ThreadObjectUtil *ThreadObjectUtil
}


func (this *StupidTimer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.NullUtil = NullUtil_GetInstance()
    this.ThreadObjectUtil = ThreadObjectUtil_GetInstance()
}

func NewStupidTimer() *StupidTimer {
    this := &StupidTimer{}
    this.initFields()
    return this
}


func (this *StupidTimer) VisitBool(visitorInterface IVisitor) bool {
        var result bool = any(visitorInterface.Visit(this.NullUtil.NULL_OBJECT)).(bool); _ = result
        var result2 bool = result; _ = result2
        return result2
}


// throws Exception
func (this *StupidTimer) Visit(visitorInterface IVisitor, timeDelayHelper ITimeDelayHelper) {
        var tookTooLong bool = false; _ = tookTooLong
        var WAITING_FOR string = "Waiting for: "; _ = WAITING_FOR
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(WAITING_FOR).Append(StringUtil_GetInstance().ToString(visitorInterface))), this, this.CommonStrings.VISIT)
        var index int = 0; _ = index
        for this.VisitBool(visitorInterface) {
            if index % 10 == 0 {
                PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(WAITING_FOR).Appendint(index)), this, this.CommonStrings.VISIT)
            }
            index++
            { // synchronized (this)
                this.ThreadObjectUtil.WaitObjectFor(this, 1800)
            }
            if timeDelayHelper.IsTimeTNT() {
                tookTooLong = true
                break
            }
        }
        if tookTooLong {
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.VISIT, errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("Took Too Long: ").Append(StringUtil_GetInstance().ToString(visitorInterface))))))
        } else {
            PreLogUtil_Put(fmt.Sprint(timeDelayHelper), this, this.CommonStrings.VISIT)
        }
}


// throws Exception
func (this *StupidTimer) StopWaiting() {
        { // synchronized (this)
            this.ThreadObjectUtil.NotifyObject(this)
        }
}
