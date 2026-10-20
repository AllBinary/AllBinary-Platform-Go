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

type BasicTouchInputFactory struct {
    LogUtil *LogUtil
    SPECIAL_BUTTON_SIX *TouchButtonInput
    SPECIAL_BUTTON_FIVE *TouchButtonInput
    SPECIAL_BUTTON_FOUR *TouchButtonInput
    SPECIAL_BUTTON_THREE *TouchButtonInput
    SPECIAL_BUTTON_TWO *TouchButtonInput
    SPECIAL_BUTTON_ONE *TouchButtonInput
    UP *TouchButtonInput
    LEFT *TouchButtonInput
    RIGHT *TouchButtonInput
    DOWN *TouchButtonInput
    NONE *TouchButtonInput
    Initialized bool
    List IBasicArrayList
}

var BasicTouchInputFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *BasicTouchInputFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.SPECIAL_BUTTON_SIX = NewTouchButtonInput(InputFactory_GetInstance().MAX - 30, "Button 6")
    this.SPECIAL_BUTTON_FIVE = NewTouchButtonInput(InputFactory_GetInstance().MAX - 31, "Button 5")
    this.SPECIAL_BUTTON_FOUR = NewTouchButtonInput(InputFactory_GetInstance().MAX - 32, "Button 4")
    this.SPECIAL_BUTTON_THREE = NewTouchButtonInput(InputFactory_GetInstance().MAX - 33, "Button 3")
    this.SPECIAL_BUTTON_TWO = NewTouchButtonInput(InputFactory_GetInstance().MAX - 34, "Button 2")
    this.SPECIAL_BUTTON_ONE = NewTouchButtonInput(InputFactory_GetInstance().MAX - 35, "Button 1")
    this.UP = NewTouchButtonInput(InputFactory_GetInstance().MAX - 36, "Up Button")
    this.LEFT = NewTouchButtonInput(InputFactory_GetInstance().MAX - 37, "Left Button")
    this.RIGHT = NewTouchButtonInput(InputFactory_GetInstance().MAX - 38, "Right Button")
    this.DOWN = NewTouchButtonInput(InputFactory_GetInstance().MAX - 39, "Down Button")
    this.NONE = NewTouchButtonInput(InputFactory_GetInstance().MAX - 40, "No Button")
    this.Initialized = false
    this.List = NewBasicArrayListD()
}

func NewBasicTouchInputFactory() *BasicTouchInputFactory {
    this := &BasicTouchInputFactory{}
    this.initFields()
    return this
}


func BasicTouchInputFactory_GetInstance() *BasicTouchInputFactory {
        if BasicTouchInputFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            BasicTouchInputFactory_Instance = NewBasicTouchInputFactory()
        }
        return any(BasicTouchInputFactory_Instance).(*BasicTouchInputFactory)
}


// synchronized
func (this *BasicTouchInputFactory) Init(inputToGameKeyMapping *InputToGameKeyMapping) {
        if !this.Initialized {
            this.Initialized = true
            this.List.Add(this.UP)
            this.List.Add(this.LEFT)
            this.List.Add(this.RIGHT)
            this.List.Add(this.DOWN)
            this.List.Add(this.SPECIAL_BUTTON_TWO)
            this.List.Add(this.SPECIAL_BUTTON_FOUR)
            this.List.Add(this.SPECIAL_BUTTON_THREE)
            this.List.Add(this.SPECIAL_BUTTON_ONE)
            this.List.Add(this.SPECIAL_BUTTON_FIVE)
            this.List.Add(this.SPECIAL_BUTTON_SIX)
            this.UpdateAllList(this.List, inputToGameKeyMapping)
        }
}


func (this *BasicTouchInputFactory) UpdateAll(inputToGameKeyMapping *InputToGameKeyMapping) {
        this.UpdateAllList(this.GetList(), inputToGameKeyMapping)
}


func (this *BasicTouchInputFactory) UpdateAllList(list IBasicArrayList, inputToGameKeyMapping *InputToGameKeyMapping) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START).Appendint(list.Size())), this, "updateAll")
        for index := int(list.Size() - 1); index >= 0; index-- {
            var touchButtonInput *TouchButtonInput = any(list.Get(index)).(*TouchButtonInput); _ = touchButtonInput
            touchButtonInput.Update(inputToGameKeyMapping)
        }
}


func (this *BasicTouchInputFactory) GetList() IBasicArrayList {
        return this.List
}

