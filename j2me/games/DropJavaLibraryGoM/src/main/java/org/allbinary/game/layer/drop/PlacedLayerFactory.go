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

type PlacedLayerFactory struct {
    List IBasicArrayList
    BasicArrayListUtil *BasicArrayListUtil
}

var PlacedLayerFactory_Instance *PlacedLayerFactory = NewPlacedLayerFactory()

func (this *PlacedLayerFactory) initFields() {
    this.List = NewBasicArrayListD()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
}


func NewPlacedLayerFactory() *PlacedLayerFactory {
    this := &PlacedLayerFactory{}
    this.initFields()
    return this
}


func (this *PlacedLayerFactory) Clear() {
        this.List.Clear()
}


func PlacedLayerFactory_GetInstance() *PlacedLayerFactory {
        return PlacedLayerFactory_Instance
}


func (this *PlacedLayerFactory) GetSize() int {
        return this.List.Size()
}


func (this *PlacedLayerFactory) GetRandomInstance() AllBinaryLayerFactoryInterface {
        return any(this.BasicArrayListUtil.GetRandom(this.List)).(AllBinaryLayerFactoryInterface)
}


func (this *PlacedLayerFactory) Add(layerInterfaceFactoryInterface AllBinaryLayerFactoryInterface) {
        this.List.Add(layerInterfaceFactoryInterface)
}

