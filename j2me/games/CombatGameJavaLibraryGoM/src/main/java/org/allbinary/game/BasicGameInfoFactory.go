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

type BasicGameInfoFactory struct {
    CurrentReleaseIndex int
    CurrentDestroyIndex int
    Total int
}

var BasicGameInfoFactory_SINGLETON *BasicGameInfoFactory = NewBasicGameInfoFactory()

func (this *BasicGameInfoFactory) initFields() {
    
    
    
}


func BasicGameInfoFactory_GetInstance() *BasicGameInfoFactory {
        return BasicGameInfoFactory_SINGLETON
}


func NewBasicGameInfoFactory() *BasicGameInfoFactory {
    this := &BasicGameInfoFactory{}
    this.initFields()
    return this
}


func (this *BasicGameInfoFactory) Init(total int) {
        this.CurrentReleaseIndex = 0
        this.CurrentDestroyIndex = 0
        this.Total = total
}


func (this *BasicGameInfoFactory) IsPeaked() bool {
        if this.CurrentDestroyIndex >= this.Total {
            return true
        } else {
            return false
        }
}


func (this *BasicGameInfoFactory) HowManyLeft() int {
        return this.Total - this.CurrentDestroyIndex
}

