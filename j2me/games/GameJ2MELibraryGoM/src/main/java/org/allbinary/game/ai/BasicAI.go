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

type BasicAI struct {
    ArtificialIntelligence
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    GameKeyEventFactory *GameKeyEventFactory
    OwnerLayerInterface IAllBinaryLayer
    GameInput IGameInput
    LastKey int
}

// Implements GameKeyEventSourceInterface

// The Java class BasicAI has subclasses so its variables are this interface
type IBasicAI interface {
    IArtificialIntelligence
    BasicAI_() *BasicAI
    ProcessAI(allBinaryLayerManager IAllBinaryLayerManager)
    ProcessKeyAI(key int)
    SetLastKey(lastKey int)
    GetLastKey() int
    GetOwnerLayerInterface() IAllBinaryLayer
    GetGameInput() IGameInput
    GetName() string
    ToString() string
    GetSourceId() int
}

func (this *BasicAI) BasicAI_() *BasicAI {
    return this
}

var BasicAI_AI_VISITOR int = SmallIntegerSingletonFactory_GetInstance().GetAt(2)
var BasicAI_ID int = SmallIntegerSingletonFactory_GetInstance().GetAt(1)

func (this *BasicAI) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.GameKeyEventFactory = GameKeyEventFactory_GetInstance()
    
    
    this.LastKey = -1
}


func NewBasicAI(ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *BasicAI {
    this := &BasicAI{}
    this.ArtificialIntelligence = *NewArtificialIntelligence()
    this.initFields()
        this.OwnerLayerInterface = ownerLayerInterface
        this.GameInput = gameInput
    return this
}


// throws Exception
func (this *BasicAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *BasicAI) ProcessKeyAI(key int) {
        if key != -1 {
            this.GameInput.Add(this.GameKeyEventFactory.GetInstanceForKey(this, key))
        }
}


func (this *BasicAI) SetLastKey(lastKey int) {
        this.LastKey = lastKey
}


func (this *BasicAI) GetLastKey() int {
        return this.LastKey
}


func (this *BasicAI) GetOwnerLayerInterface() IAllBinaryLayer {
        return this.OwnerLayerInterface
}


func (this *BasicAI) GetGameInput() IGameInput {
        return this.GameInput
}


func (this *BasicAI) GetName() string {
        return reflect.TypeOf(this).String()
}


func (this *BasicAI) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        stringBuffer.Append(this.GetName())
        stringBuffer.Append(" LastKey: ")
        stringBuffer.Append(CanvasUtil_GetKeyName(this.GetLastKey()))
        return fmt.Sprint(stringBuffer)
}

func (this *BasicAI) String() string {
    return this.ToString()
}


func (this *BasicAI) GetSourceId() int {
        return 1
}

