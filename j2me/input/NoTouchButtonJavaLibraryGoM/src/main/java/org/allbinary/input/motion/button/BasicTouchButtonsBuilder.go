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

type BasicTouchButtonsBuilder struct {
    BaseTouchInput
    LogUtil *LogUtil
}

var BasicTouchButtonsBuilder_SINGLETON *BasicTouchButtonsBuilder = NewBasicTouchButtonsBuilder()

func (this *BasicTouchButtonsBuilder) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewBasicTouchButtonsBuilder() *BasicTouchButtonsBuilder {
    this := &BasicTouchButtonsBuilder{}
    this.BaseTouchInput = *NewBaseTouchInput()
    this.initFields()
    return this
}


func BasicTouchButtonsBuilder_GetInstance() IBaseTouchInput {
        return BasicTouchButtonsBuilder_SINGLETON
}


func (this *BasicTouchButtonsBuilder) GetList() IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.GET_LIST, e)
            tryControl1, tryResult1 = 1, BasicArrayListUtil_GetInstance().GetImmutableInstance(); return
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.GET_LIST)
            var list IBasicArrayList = NewBasicArrayListD(); _ = list
            tryControl1, tryResult1 = 1, list; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

