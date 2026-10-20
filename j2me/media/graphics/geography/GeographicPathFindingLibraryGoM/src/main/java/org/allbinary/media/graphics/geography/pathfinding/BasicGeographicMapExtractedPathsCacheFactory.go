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

type BasicGeographicMapExtractedPathsCacheFactory struct {
    BaseBasicArrayListCache
    MaxConcurrentPaths int
    MaxPaths int
}

var BasicGeographicMapExtractedPathsCacheFactory_Instance *BasicGeographicMapExtractedPathsCacheFactory = NewBasicGeographicMapExtractedPathsCacheFactory()

func (this *BasicGeographicMapExtractedPathsCacheFactory) initFields() {
    
    
}


func BasicGeographicMapExtractedPathsCacheFactory_GetInstance() *BasicGeographicMapExtractedPathsCacheFactory {
        return BasicGeographicMapExtractedPathsCacheFactory_Instance
}


func NewBasicGeographicMapExtractedPathsCacheFactory() *BasicGeographicMapExtractedPathsCacheFactory {
    this := &BasicGeographicMapExtractedPathsCacheFactory{}
    this.BaseBasicArrayListCache = *NewBaseBasicArrayListCache()
    this.initFields()
    return this
}


func (this *BasicGeographicMapExtractedPathsCacheFactory) Init(maxConcurrentPaths int, maxPaths int) {
        if maxConcurrentPaths != this.MaxConcurrentPaths {
            for index := int(maxConcurrentPaths); index >= this.MaxConcurrentPaths; index-- {
                this.List.Add(NewBasicArrayListS(maxPaths))
            }
            this.MaxConcurrentPaths = maxConcurrentPaths
        }
        this.MaxPaths = maxPaths
}


func (this *BasicGeographicMapExtractedPathsCacheFactory) GetPaths() IBasicArrayList {
        if this.List.Size() > 0 {
            return any(this.List.RemoveAt(0)).(IBasicArrayList)
        } else {
            return NewBasicArrayListS(this.MaxPaths)
        }
}


func (this *BasicGeographicMapExtractedPathsCacheFactory) Release(list IBasicArrayList) {
        for index := int(list.Size() - 1); index >= 0; index-- {
            BasicGeographicMapExtractedPathCacheFactory_GetInstance().Release(any(list.Get(index)).(IBasicArrayList))
        }
        list.Clear()
        this.List.Add(list)
}

