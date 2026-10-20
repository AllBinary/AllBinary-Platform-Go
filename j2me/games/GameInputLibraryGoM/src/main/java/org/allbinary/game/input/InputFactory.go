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

type InputFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    KEY_CODE_LABEL string
    DEVICE_ID_LABEL string
    MAX int
    InputIntegerArray []IInput
    NO_INPUT IInput
}

var InputFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *InputFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.KEY_CODE_LABEL = "KeyCode: "
    this.DEVICE_ID_LABEL = "DeviceId: "
    this.MAX = SmallIntegerSingletonFactory_GetInstance().GetMin()
    this.InputIntegerArray = make([]IInput, this.MAX)
    this.NO_INPUT = NewInput(0, this.CommonStrings.UNKNOWN)
}


func InputFactory_GetInstance() *InputFactory {
        if InputFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            InputFactory_Instance = NewInputFactory()
        }
        return any(InputFactory_Instance).(*InputFactory)
}


func NewInputFactory() *InputFactory {
    this := &InputFactory{}
    this.initFields()
        var size int = len(this.InputIntegerArray); _ = size
        for index := int(0); index < size; index++ {
            this.InputIntegerArray[index] = this.NO_INPUT
        }
    return this
}


func (this *InputFactory) Add(id int, input IInput) {
        this.InputIntegerArray[id] = input
}


func (this *InputFactory) GetInstanceById(id int) IInput {
        if id < 0 {
            id = -id
        }
        if id > len(this.InputIntegerArray) {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Warning id: ").Appendint(id)), this, this.CommonStrings.GET_INSTANCE)
            return this.NO_INPUT
        }
        return this.InputIntegerArray[id]
}

