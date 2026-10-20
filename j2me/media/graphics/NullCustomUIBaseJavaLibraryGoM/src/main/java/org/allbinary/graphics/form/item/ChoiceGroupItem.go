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

type ChoiceGroupItem struct {
}

// The Java class ChoiceGroupItem has subclasses so its variables are this interface
type IChoiceGroupItem interface {
    ChoiceGroupItem_() *ChoiceGroupItem
}

func (this *ChoiceGroupItem) ChoiceGroupItem_() *ChoiceGroupItem {
    return this
}


func NewChoiceGroupItem() *ChoiceGroupItem {
    this := &ChoiceGroupItem{}
    return this
}

