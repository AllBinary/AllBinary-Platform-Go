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

type ActivityFractureControllerUtilFactory struct {
}


func NewActivityFractureControllerUtilFactory() *ActivityFractureControllerUtilFactory {
    this := &ActivityFractureControllerUtilFactory{}
    return this
}


func ActivityFractureControllerUtilFactory_GetInstance() IAndroidKeyFactory {
        var SDK_VERSION int = AndroidInfoFactory_GetInstance().GetVersion(); _ = SDK_VERSION
        if SDK_VERSION <= 8 {
            return AndroidAPI1KeyFactory_GetInstance()
        } else if SDK_VERSION <= 11 {
            return AndroidAPI9KeyFactory_GetInstance()
        } else {
            return AndroidAPI12KeyFactory_GetAPI12Instance()
        }
}

