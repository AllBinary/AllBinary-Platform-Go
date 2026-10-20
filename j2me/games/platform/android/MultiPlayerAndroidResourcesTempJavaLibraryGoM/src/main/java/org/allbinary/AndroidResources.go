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

type AndroidResources struct {
    Drawable *drawable
    Id *id
    Layout *layout
    Raw *raw
    String *string
}

var AndroidResources_Instance *AndroidResources = NewAndroidResources()

func (this *AndroidResources) initFields() {
    this.Drawable = Newdrawable()
    this.Id = Newid()
    this.Layout = Newlayout()
    this.Raw = Newraw()
    this.String = Newstring()
}

func NewAndroidResources() *AndroidResources {
    this := &AndroidResources{}
    this.initFields()
    return this
}


func AndroidResources_GetInstance() *AndroidResources {
        return AndroidResources_Instance
}

