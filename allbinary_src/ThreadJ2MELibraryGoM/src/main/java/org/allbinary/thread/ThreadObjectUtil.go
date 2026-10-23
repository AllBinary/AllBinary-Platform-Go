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

type ThreadObjectUtil struct {
    NULL_PRIORITY_RUNNABLE PriorityRunnable
}

var ThreadObjectUtil_Instance *ThreadObjectUtil = NewThreadObjectUtil()

func (this *ThreadObjectUtil) initFields() {
    this.NULL_PRIORITY_RUNNABLE = NewNullPriorityRunnable()
}

func NewThreadObjectUtil() *ThreadObjectUtil {
    this := &ThreadObjectUtil{}
    this.initFields()
    return this
}


func ThreadObjectUtil_GetInstance() *ThreadObjectUtil {
        return ThreadObjectUtil_Instance
}


func (this *ThreadObjectUtil) NotifyObject(anyType any) {
        Java_Notify(anyType)
}


// throws Exception
func (this *ThreadObjectUtil) WaitObject(anyType any) {
        Java_Wait(anyType)
}


// throws Exception
func (this *ThreadObjectUtil) WaitObjectFor(anyType any, delay int64) {
        TsUtil_GetInstance().WaitFor(anyType, delay)
}


func (this *ThreadObjectUtil) ProcessThread(runnable Runnable) {
        var thread IThread = NewThread(runnable); _ = thread
        thread.Start()
}

