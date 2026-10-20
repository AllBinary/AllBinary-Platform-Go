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

type MakeCountedPartsSingletonArrayFactory struct {
    LogUtil *LogUtil
    List IBasicArrayList
}

var MakeCountedPartsSingletonArrayFactory_Instance *MakeCountedPartsSingletonArrayFactory = NewMakeCountedPartsSingletonArrayFactory()

func (this *MakeCountedPartsSingletonArrayFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.List = BasicArrayListUtil_GetInstance().GetImmutableInstance()
}

func NewMakeCountedPartsSingletonArrayFactory() *MakeCountedPartsSingletonArrayFactory {
    this := &MakeCountedPartsSingletonArrayFactory{}
    this.initFields()
    return this
}


func MakeCountedPartsSingletonArrayFactory_GetInstance() *MakeCountedPartsSingletonArrayFactory {
        return MakeCountedPartsSingletonArrayFactory_Instance
}


// throws Exception
func (this *MakeCountedPartsSingletonArrayFactory) GetUpdated(partInterfaceArray []PartInterface) []PartInterface {
        this.List = NewBasicArrayListD()
        for index := int(0); index < len(partInterfaceArray); index++ {
            this.List.Add(partInterfaceArray[index])
        }
        var countedBasicArrayList IBasicArrayList = CountedPickedUpLayerInterfaceFactoryPool_GetInstance().GetList(); _ = countedBasicArrayList
        var layerInterfaceFactoryInterface ICountedPickedUpLayerInterfaceFactory; _ = layerInterfaceFactoryInterface
        var size int = countedBasicArrayList.Size(); _ = size
        for index := int(0); index < size; index++ {
            layerInterfaceFactoryInterface = any(countedBasicArrayList.BasicArrayList_().ObjectArray[index]).(ICountedPickedUpLayerInterfaceFactory)
            this.List.Add(NewCountedLayerInterfaceFactoryPart(0, layerInterfaceFactoryInterface))
        }
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Total Parts: ")
        stringBuffer.Appendint(this.List.Size())
        stringBuffer.Append(" Counted: ")
        stringBuffer.Appendint(countedBasicArrayList.Size())
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, commonStrings.GET_INSTANCE)
        var newPartInterfaceArray []PartInterface = make([]PartInterface, this.List.Size()); _ = newPartInterfaceArray
        var newSize int = len(newPartInterfaceArray); _ = newSize
        for index := int(0); index < newSize; index++ {
            newPartInterfaceArray[index] = any(this.List.Get(index)).(PartInterface)
        }
        return newPartInterfaceArray
}

