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

type ToastUtil struct {
    LogUtil *LogUtil
}

var ToastUtil_Instance *ToastUtil = NewToastUtil()

func (this *ToastUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewToastUtil() *ToastUtil {
    this := &ToastUtil{}
    this.initFields()
    return this
}


func ToastUtil_GetInstance() *ToastUtil {
        return ToastUtil_Instance
}


func (this *ToastUtil) MakeToast(view *View, context *Context, message string, time int) {
        /* Java local class: move it out of the function
type MakeToast struct {
    String string
}

// Implements Runnable


func (this *MakeToast) initFields() {
            
}
        

        func NewToastUtil(string_ string) *ToastUtil {
            this := &ToastUtil{}
            this.initFields()
                this.String = string_
            return this
        }
        

        func (this *ToastUtil) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                    logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                        }
                    }()
                    Toast_MakeText(this.Context, this.String, this.Time).Show()
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/
        view.Post(NewMakeToast(message))
}

