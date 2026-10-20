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

type NullCommandListener struct {
}

// Implements CommandListener

// The Java class NullCommandListener has subclasses so its variables are this interface
type INullCommandListener interface {
    NullCommandListener_() *NullCommandListener
    CommandAction(c ICommand, d IDisplayable)
}

func (this *NullCommandListener) NullCommandListener_() *NullCommandListener {
    return this
}

var NullCommandListener_NULL_COMMAND_LISTENER INullCommandListener = NewNullCommandListener()

func NewNullCommandListener() *NullCommandListener {
    this := &NullCommandListener{}
    return this
}


func (this *NullCommandListener) CommandAction(c ICommand, d IDisplayable) {
}

