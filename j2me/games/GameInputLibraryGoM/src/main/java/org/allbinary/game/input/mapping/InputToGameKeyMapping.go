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

type InputToGameKeyMapping struct {
    InputMapping
    GameKeyFactory *GameKeyFactory
    MappedGameKeys []*GameKey
    PlatformToGameKeyMapping []*GameKey
    NegativePlatformToGameKeyMapping []*GameKey
}

var InputToGameKeyMapping_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *InputToGameKeyMapping) initFields() {
    this.GameKeyFactory = GameKeyFactory_GetInstance()
    this.MappedGameKeys = []*GameKey{this.GameKeyFactory.UP, this.GameKeyFactory.DOWN, this.GameKeyFactory.LEFT, this.GameKeyFactory.RIGHT}
    this.PlatformToGameKeyMapping = make([]*GameKey, InputFactory_GetInstance().MAX)
    this.NegativePlatformToGameKeyMapping = make([]*GameKey, 0)
}


func InputToGameKeyMapping_GetNullInstance() *InputToGameKeyMapping {
        if InputToGameKeyMapping_Instance == NullUtil_GetInstance().NULL_OBJECT {
            InputToGameKeyMapping_Instance = NewInputToGameKeyMapping()
        }
        return any(InputToGameKeyMapping_Instance).(*InputToGameKeyMapping)
}


func NewInputToGameKeyMapping() *InputToGameKeyMapping {
    this := &InputToGameKeyMapping{}
    this.InputMapping = *NewInputMapping()
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.Clear()
    return this
}


func (this *InputToGameKeyMapping) Clear() {
        for index := int(len(this.PlatformToGameKeyMapping) - 1); index >= 0; index-- {
            this.PlatformToGameKeyMapping[index] = this.GameKeyFactory.NONE
        }
}


func (this *InputToGameKeyMapping) RemoveAll() {
        this.InputMapping.RemoveAll()
        this.Clear()
}


func (this *InputToGameKeyMapping) AddMapping(inputToGameKeyMapping *InputToGameKeyMapping) {
        for index := int(inputToGameKeyMapping.GetMaxMappable() - 1); index >= 0; index-- {
            var input IInput = inputToGameKeyMapping.GetInstance(index); _ = input
            if input != this.GameKeyFactory.NONE {
                this.AddAll(input, inputToGameKeyMapping.GetMappedInput(input))
            }
        }
}


func (this *InputToGameKeyMapping) AddAll(input IInput, list IBasicArrayList) {
        for index := int(list.Size() - 1); index >= 0; index-- {
            var mappedToInput IInput = any(list.BasicArrayList_().ObjectArray[index]).(IInput); _ = mappedToInput
            this.Add(input, mappedToInput)
        }
}


func (this *InputToGameKeyMapping) Add(input IInput, mappedToInput IInput) {
        this.InputMapping.Add(input, mappedToInput)
        this.Set(input, mappedToInput)
}


func (this *InputToGameKeyMapping) Remove(input IInput, mappedToInput IInput) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START_LABEL).Append(StringUtil_GetInstance().ToString(input)).Append(" == ").Append(StringUtil_GetInstance().ToString(mappedToInput))), this, "InputToGameKeyMapping::remove")
        this.InputMapping.Remove(input, mappedToInput)
        this.Set(input, this.GameKeyFactory.NONE)
}


func (this *InputToGameKeyMapping) Set(input IInput, mappedToInput IInput) {
        var id int = mappedToInput.GetId(); _ = id
        if id >= 0 && id < len(this.PlatformToGameKeyMapping) {
            this.PlatformToGameKeyMapping[id] = any(input).(*GameKey)
        }
}


func (this *InputToGameKeyMapping) GetMaxMappable() int {
        return len(this.PlatformToGameKeyMapping)
}


func (this *InputToGameKeyMapping) GetReverseInstance(id int) IBasicArrayList {
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var inputFactory *InputFactory = InputFactory_GetInstance(); _ = inputFactory
        for index := int(0); index < len(this.PlatformToGameKeyMapping); index++ {
            if this.PlatformToGameKeyMapping[index].GetId() == id {
                var input IInput = inputFactory.GetInstanceById(index); _ = input
                if input != nil {
                    list.Add(input)
                }
            }
        }
        return list
}


// throws Exception
func (this *InputToGameKeyMapping) InitGameKey(canvas ICanvas, gameKey *GameKey) {
        var key int = canvas.GetKeyCode(gameKey.GetId()); _ = key
        if key < 0 && -key < len(this.NegativePlatformToGameKeyMapping) {
            this.NegativePlatformToGameKeyMapping[-key] = gameKey
        } else if key > 0 {
            if this.PlatformToGameKeyMapping[key] == this.GameKeyFactory.NONE {
                this.PlatformToGameKeyMapping[key] = gameKey
            }
        } else {
            PreLogUtil_PutOE(fmt.Sprint(NewStringMaker().Append(fmt.Sprint(gameKey)).Append(CommonSeps_GetInstance().COLON_SEP).Appendint(key)), this, this.CommonStrings.INIT, errors.New(fmt.Sprint()))
        }
}


// throws Exception
func (this *InputToGameKeyMapping) GetSmallestCanvasGameKeyCode(canvas ICanvas) int {
        var nextKey int = 0; _ = nextKey
        var smallestKey int = 0; _ = smallestKey
        for index := int(len(this.MappedGameKeys)); Java_PreDecrement(&index) >= 0;  {
            nextKey = canvas.GetKeyCode(this.MappedGameKeys[index].GetId())
            if nextKey < smallestKey {
                smallestKey = nextKey
            }
        }
        return smallestKey
}


func (this *InputToGameKeyMapping) Init(canvas ICanvas) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    t := Java_ToError(recovered)
                    _ = t
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, t)
                }
            }()
            var negativePlatformToGameKeyMapping []*GameKey = make([]*GameKey, 0); _ = negativePlatformToGameKeyMapping
            var smallestKey int = this.GetSmallestCanvasGameKeyCode(canvas); _ = smallestKey
            if smallestKey < 0 {
                smallestKey--
                if smallestKey < -256 {
                    smallestKey = -256
                }
                negativePlatformToGameKeyMapping = make([]*GameKey, -smallestKey)
                for index := int(len(negativePlatformToGameKeyMapping)); Java_PreDecrement(&index) >= 0;  {
                    negativePlatformToGameKeyMapping[index] = this.GameKeyFactory.NONE
                }
                this.NegativePlatformToGameKeyMapping = negativePlatformToGameKeyMapping
            }
            for index := int(len(this.MappedGameKeys)); Java_PreDecrement(&index) >= 0;  {
                this.InitGameKey(canvas, this.MappedGameKeys[index])
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InputToGameKeyMapping) GetInstanceForCanvas(canvas ICanvas, id int) *GameKey {
        if id >= 0 && id < len(this.PlatformToGameKeyMapping) {
            return this.PlatformToGameKeyMapping[id]
        } else if id < 0 && -id < len(this.NegativePlatformToGameKeyMapping) {
            return this.NegativePlatformToGameKeyMapping[-id]
        } else {
            return this.GameKeyFactory.NONE
        }
}


func (this *InputToGameKeyMapping) GetInstance(id int) *GameKey {
        if id >= 0 && id < len(this.PlatformToGameKeyMapping) {
            return this.PlatformToGameKeyMapping[id]
        } else if id < 0 && -id < len(this.NegativePlatformToGameKeyMapping) {
            return this.NegativePlatformToGameKeyMapping[-id]
        } else {
            return this.GameKeyFactory.NONE
        }
}

