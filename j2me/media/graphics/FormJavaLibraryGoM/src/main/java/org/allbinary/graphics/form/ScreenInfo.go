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

type ScreenInfo struct {
    SMALL_WIDTH int
    SMALL_HEIGHT int
    MEDIUM_WIDTH int
    MEDIUM_HEIGHT int
}

var ScreenInfo_Instance *ScreenInfo = NewScreenInfo()

func (this *ScreenInfo) initFields() {
    this.SMALL_WIDTH = 320
    this.SMALL_HEIGHT = 240
    this.MEDIUM_WIDTH = 800
    this.MEDIUM_HEIGHT = 600
}

func NewScreenInfo() *ScreenInfo {
    this := &ScreenInfo{}
    this.initFields()
    return this
}


func ScreenInfo_GetInstance() *ScreenInfo {
        return ScreenInfo_Instance
}

