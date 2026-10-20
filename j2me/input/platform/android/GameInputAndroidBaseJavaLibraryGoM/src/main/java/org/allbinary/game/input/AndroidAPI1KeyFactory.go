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

type AndroidAPI1KeyFactory struct {
    AndroidKeyFactory
}

var AndroidAPI1KeyFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT


func AndroidAPI1KeyFactory_GetInstance() *AndroidAPI1KeyFactory {
        if AndroidAPI1KeyFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            AndroidAPI1KeyFactory_Instance = NewAndroidAPI1KeyFactory()
        }
        return any(AndroidAPI1KeyFactory_Instance).(*AndroidAPI1KeyFactory)
}


func NewAndroidAPI1KeyFactory() *AndroidAPI1KeyFactory {
    this := &AndroidAPI1KeyFactory{}
    this.AndroidKeyFactory = *NewAndroidKeyFactory()
    return this
}

