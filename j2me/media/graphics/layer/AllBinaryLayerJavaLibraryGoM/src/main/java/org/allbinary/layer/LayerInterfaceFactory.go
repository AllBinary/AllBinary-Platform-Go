/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
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

type LayerInterfaceFactory struct {
    LogUtil *LogUtil
    List IBasicArrayList
    HASHTABLE_LABEL string
}

var LayerInterfaceFactory_SINGLETON *LayerInterfaceFactory = NewLayerInterfaceFactory()

func (this *LayerInterfaceFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.List = BasicArrayListUtil_GetInstance().GetImmutableInstance()
    this.HASHTABLE_LABEL = "Hashtable: "
}

func NewLayerInterfaceFactory() *LayerInterfaceFactory {
    this := &LayerInterfaceFactory{}
    this.initFields()
    return this
}


func LayerInterfaceFactory_GetInstance() *LayerInterfaceFactory {
        return LayerInterfaceFactory_SINGLETON
}


func (this *LayerInterfaceFactory) Init() {
        this.List = NewBasicArrayListD()
}


// throws Exception
func (this *LayerInterfaceFactory) GetNexInstance(hashtable *ABHashtable[any, any], x int, y int, z int) IAllBinaryLayer {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.HASHTABLE_LABEL).Append(StringUtil_GetInstance().ToString(hashtable))), this, commonStrings.GET_INSTANCE)
        var typeInteger int = any(hashtable.Get(any(Layer_ID).(any))).(int); _ = typeInteger
        var layerInterfaceFactoryInterface LayerInterfaceFactoryInterface = any(this.List.BasicArrayList_().ObjectArray[int(typeInteger) - 1]).(LayerInterfaceFactoryInterface); _ = layerInterfaceFactoryInterface
        return layerInterfaceFactoryInterface.GetNextInstance(hashtable, x, y, z)
}


func (this *LayerInterfaceFactory) Add(layerInterfaceFactoryInterface LayerInterfaceFactoryInterface) {
        this.List.Add(layerInterfaceFactoryInterface)
}

