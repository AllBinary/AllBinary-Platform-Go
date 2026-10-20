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

type BasicGeographicMapExtractedPathCacheFactory struct {
    BaseBasicArrayListCache
    MaxConcurrentPaths int
    MaxPaths int
}

var BasicGeographicMapExtractedPathCacheFactory_Instance *BasicGeographicMapExtractedPathCacheFactory = NewBasicGeographicMapExtractedPathCacheFactory()

func (this *BasicGeographicMapExtractedPathCacheFactory) initFields() {
    
    
}


func BasicGeographicMapExtractedPathCacheFactory_GetInstance() *BasicGeographicMapExtractedPathCacheFactory {
        return BasicGeographicMapExtractedPathCacheFactory_Instance
}


func NewBasicGeographicMapExtractedPathCacheFactory() *BasicGeographicMapExtractedPathCacheFactory {
    this := &BasicGeographicMapExtractedPathCacheFactory{}
    this.BaseBasicArrayListCache = *NewBaseBasicArrayListCache()
    this.initFields()
    return this
}


func (this *BasicGeographicMapExtractedPathCacheFactory) Init(maxConcurrentPaths int, maxPaths int) {
        if maxConcurrentPaths != this.MaxConcurrentPaths {
            for index := int(maxConcurrentPaths); index >= this.MaxConcurrentPaths; index-- {
                this.List.Add(NewBasicArrayListS(maxPaths))
            }
            this.MaxConcurrentPaths = maxConcurrentPaths
        }
        this.MaxPaths = maxPaths
}


func (this *BasicGeographicMapExtractedPathCacheFactory) GetPath() IBasicArrayList {
        if this.List.Size() > 0 {
            return any(this.List.RemoveAt(0)).(IBasicArrayList)
        } else {
            return NewBasicArrayListS(this.MaxPaths)
        }
}


func (this *BasicGeographicMapExtractedPathCacheFactory) Release(list IBasicArrayList) {
        list.Clear()
        this.List.Add(list)
}

