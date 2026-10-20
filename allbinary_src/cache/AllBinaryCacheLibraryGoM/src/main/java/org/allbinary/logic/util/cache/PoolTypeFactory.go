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

type PoolTypeFactory struct {
    VECTOR_POOL *PoolType
}

var PoolTypeFactory_Instance *PoolTypeFactory = NewPoolTypeFactory()

func (this *PoolTypeFactory) initFields() {
    this.VECTOR_POOL = NewPoolType(CacheTypeFactory_GetInstance().NO_CACHE)
}

func NewPoolTypeFactory() *PoolTypeFactory {
    this := &PoolTypeFactory{}
    this.initFields()
    return this
}


func PoolTypeFactory_GetInstance() *PoolTypeFactory {
        return PoolTypeFactory_Instance
}

