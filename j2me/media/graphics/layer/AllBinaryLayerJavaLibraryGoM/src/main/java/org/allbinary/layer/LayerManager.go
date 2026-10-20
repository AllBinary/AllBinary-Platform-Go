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

type LayerManager struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    SystemWrapper *ABSystemWrapper
    TsUtil *TsUtil
    LayerManagerLogging ILayerManagerLoggingBase
    List IBasicArrayList
    AnyType *SynchObject
}

// The Java class LayerManager has subclasses so its variables are this interface
type ILayerManager interface {
    LayerManager_() *LayerManager
    Contains(layerInterface IAllBinaryLayer) bool
    Update(layerInterface IAllBinaryLayer)
    Insert(layerInterface IAllBinaryLayer)
    Append(layerInterface IAllBinaryLayer)
    AppendAt(layerInterface IAllBinaryLayer, index int)
    Remove(layerInterface IAllBinaryLayer)
    GetLayerAt(index int) ILayer
    GetSize() int
    Cleanup()
    Paint(g IGraphics, x int, y int)
}

func (this *LayerManager) LayerManager_() *LayerManager {
    return this
}


func (this *LayerManager) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.SystemWrapper = ABSystemWrapper_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
    
    this.List = NewBasicArrayListD()
    this.AnyType = NewSynchObject()
}


func NewLayerManager(layerManagerLogging ILayerManagerLoggingBase) *LayerManager {
    this := &LayerManager{}
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.LayerManagerLogging = layerManagerLogging
    return this
}


func (this *LayerManager) Contains(layerInterface IAllBinaryLayer) bool {
        return this.List.Contains(layerInterface)
}


// throws Exception
func (this *LayerManager) Update(layerInterface IAllBinaryLayer) {
        var had bool = this.List.Remove(layerInterface); _ = had
        if had {
            var nextLayerInterface IAllBinaryLayer; _ = nextLayerInterface
            var size int = this.List.Size(); _ = size
            for index := int(0); index < size; index++ {
                nextLayerInterface = any(this.List.Get(index)).(IAllBinaryLayer)
                if layerInterface.GetZP() > nextLayerInterface.GetZP() {
                    this.List.AddAt(index, layerInterface)
                    return
                }
            }
            this.List.Add(layerInterface)
        }
}


// throws Exception
func (this *LayerManager) Insert(layerInterface IAllBinaryLayer) {
        var nextLayerInterface IAllBinaryLayer; _ = nextLayerInterface
        var size int = this.List.Size(); _ = size
        for index := int(0); index < size; index++ {
            nextLayerInterface = any(this.List.Get(index)).(IAllBinaryLayer)
            if layerInterface.GetZP() > nextLayerInterface.GetZP() {
                this.AppendAt(layerInterface, index)
                return
            }
        }
        this.Append(layerInterface)
}


// throws Exception
func (this *LayerManager) Append(layerInterface IAllBinaryLayer) {
        this.LayerManagerLogging.Append(layerInterface)
        this.List.Add(layerInterface)
}


// throws Exception
func (this *LayerManager) AppendAt(layerInterface IAllBinaryLayer, index int) {
        this.LayerManagerLogging.AppendAt(layerInterface, index)
        this.List.AddAt(index, layerInterface)
}


// throws Exception
func (this *LayerManager) Remove(layerInterface IAllBinaryLayer) {
        { // synchronized (this.AnyType)
            this.LayerManagerLogging.Remove(layerInterface)
            var result bool = this.List.Remove(layerInterface); _ = result
            this.LayerManagerLogging.RemoveResult(this, layerInterface, result)
        }
}


func (this *LayerManager) GetLayerAt(index int) ILayer {
        return any(this.List.BasicArrayList_().ObjectArray[index]).(ILayer)
}


func (this *LayerManager) GetSize() int {
        return this.List.Size()
}


// throws Exception
func (this *LayerManager) Cleanup() {
        { // synchronized (this.AnyType)
            this.List.Clear()
            this.LayerManagerLogging.Clear()
            this.TsUtil.Gc()
            this.TsUtil.Gc()
        }
}


func (this *LayerManager) Paint(g IGraphics, x int, y int) {
        { // synchronized (this.AnyType)
            var comp ILayer; _ = comp
            for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
                comp = any(this.List.BasicArrayList_().ObjectArray[index]).(ILayer)
                if comp != nil && comp.IsVisible() {
                    comp.Paint(g)
                }
            }
        }
}

