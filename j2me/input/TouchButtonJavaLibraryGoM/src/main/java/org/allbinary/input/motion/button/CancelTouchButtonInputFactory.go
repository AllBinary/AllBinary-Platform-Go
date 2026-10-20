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

type CancelTouchButtonInputFactory struct {
    CancelInputArray []*TouchButtonInput
}

var CancelTouchButtonInputFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *CancelTouchButtonInputFactory) initFields() {
    this.CancelInputArray = make([]*TouchButtonInput, InputFactory_GetInstance().MAX)
}


func CancelTouchButtonInputFactory_GetInstance() *CancelTouchButtonInputFactory {
        if CancelTouchButtonInputFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            CancelTouchButtonInputFactory_Instance = NewCancelTouchButtonInputFactory()
        }
        return any(CancelTouchButtonInputFactory_Instance).(*CancelTouchButtonInputFactory)
}


func NewCancelTouchButtonInputFactory() *CancelTouchButtonInputFactory {
    this := &CancelTouchButtonInputFactory{}
    this.initFields()
        var basicTouchInputFactory *BasicTouchInputFactory = BasicTouchInputFactory_GetInstance(); _ = basicTouchInputFactory
        for index := int(len(this.CancelInputArray) - 1); index >= 0; index-- {
            this.CancelInputArray[index] = basicTouchInputFactory.NONE
        }
        this.CancelInputArray[basicTouchInputFactory.UP.GetSourceId()] = basicTouchInputFactory.DOWN
        this.CancelInputArray[basicTouchInputFactory.DOWN.GetSourceId()] = basicTouchInputFactory.UP
        this.CancelInputArray[basicTouchInputFactory.LEFT.GetSourceId()] = basicTouchInputFactory.RIGHT
        this.CancelInputArray[basicTouchInputFactory.RIGHT.GetSourceId()] = basicTouchInputFactory.LEFT
    return this
}


func (this *CancelTouchButtonInputFactory) GetCancel(touchButtonInput *TouchButtonInput) *TouchButtonInput {
        return this.CancelInputArray[touchButtonInput.GetSourceId()]
}

