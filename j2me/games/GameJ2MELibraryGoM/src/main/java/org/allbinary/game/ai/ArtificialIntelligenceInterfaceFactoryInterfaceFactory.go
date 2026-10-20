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

type ArtificialIntelligenceInterfaceFactoryInterfaceFactory struct {
    List IBasicArrayList
}

var ArtificialIntelligenceInterfaceFactoryInterfaceFactory_SINGLETON *ArtificialIntelligenceInterfaceFactoryInterfaceFactory = NewArtificialIntelligenceInterfaceFactoryInterfaceFactory()

func (this *ArtificialIntelligenceInterfaceFactoryInterfaceFactory) initFields() {
    this.List = BasicArrayListUtil_GetInstance().GetImmutableInstance()
}

func NewArtificialIntelligenceInterfaceFactoryInterfaceFactory() *ArtificialIntelligenceInterfaceFactoryInterfaceFactory {
    this := &ArtificialIntelligenceInterfaceFactoryInterfaceFactory{}
    this.initFields()
    return this
}


func (this *ArtificialIntelligenceInterfaceFactoryInterfaceFactory) Clear() {
        this.List = NewBasicArrayListD()
}


func ArtificialIntelligenceInterfaceFactoryInterfaceFactory_GetInstance() *ArtificialIntelligenceInterfaceFactoryInterfaceFactory {
        return ArtificialIntelligenceInterfaceFactoryInterfaceFactory_SINGLETON
}


func (this *ArtificialIntelligenceInterfaceFactoryInterfaceFactory) GetInstanceFromCollection(hashtable *ABHashtable[any, any]) ArtificialIntelligenceInterfaceFactoryInterface {
        var typeInteger int = any(hashtable.Get(any(BasicAI_ID).(any))).(int); _ = typeInteger
        var type_ int = int(typeInteger); _ = type_
        var artificialIntelligenceInterfaceFactoryInterface ArtificialIntelligenceInterfaceFactoryInterface = any(this.List.BasicArrayList_().ObjectArray[type_ - 1]).(ArtificialIntelligenceInterfaceFactoryInterface); _ = artificialIntelligenceInterfaceFactoryInterface
        return artificialIntelligenceInterfaceFactoryInterface
}


func (this *ArtificialIntelligenceInterfaceFactoryInterfaceFactory) Add(artificialIntelligenceInterfaceFactoryInterface ArtificialIntelligenceInterfaceFactoryInterface) {
        this.List.Add(artificialIntelligenceInterfaceFactoryInterface)
}

