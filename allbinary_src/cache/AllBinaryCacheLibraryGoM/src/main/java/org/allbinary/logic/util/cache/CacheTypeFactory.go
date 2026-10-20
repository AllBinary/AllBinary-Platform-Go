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

type CacheTypeFactory struct {
    NO_CACHE *CacheType
    CACHE *CacheType
    SET *CacheType
    CACHE_VECTOR *CacheType
}

var CacheTypeFactory_Instance *CacheTypeFactory = NewCacheTypeFactory()

func (this *CacheTypeFactory) initFields() {
    this.NO_CACHE = NewCacheType(StringUtil_GetInstance().EMPTY_STRING, StringUtil_GetInstance().EMPTY_STRING)
    this.CACHE = NewCacheType(StringUtil_GetInstance().EMPTY_STRING, StringUtil_GetInstance().EMPTY_STRING)
    this.SET = NewCacheType(StringUtil_GetInstance().EMPTY_STRING, "Set")
    this.CACHE_VECTOR = NewCacheType(StringUtil_GetInstance().EMPTY_STRING, StringUtil_GetInstance().EMPTY_STRING)
}

func NewCacheTypeFactory() *CacheTypeFactory {
    this := &CacheTypeFactory{}
    this.initFields()
    return this
}


func CacheTypeFactory_GetInstance() *CacheTypeFactory {
        return CacheTypeFactory_Instance
}

