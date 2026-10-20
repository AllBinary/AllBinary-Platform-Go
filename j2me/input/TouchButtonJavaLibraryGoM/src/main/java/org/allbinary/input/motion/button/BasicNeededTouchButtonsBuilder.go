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

type BasicNeededTouchButtonsBuilder struct {
    BaseTouchInput
    LogUtil *LogUtil
}


func (this *BasicNeededTouchButtonsBuilder) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewBasicNeededTouchButtonsBuilder() *BasicNeededTouchButtonsBuilder {
    this := &BasicNeededTouchButtonsBuilder{}
    this.BaseTouchInput = *NewBaseTouchInput()
    this.initFields()
    return this
}


func (this *BasicNeededTouchButtonsBuilder) GetList() IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.GET_LIST, e)
            tryControl1, tryResult1 = 1, BasicArrayListUtil_GetInstance().GetImmutableInstance(); return
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
            var list IBasicArrayList = NewBasicArrayListD(); _ = list
            var touchButtonLocationHelper *TouchButtonLocationHelper = NewTouchButtonLocationHelper(); _ = touchButtonLocationHelper
            var basicTouchButtonCellPositionFactory *BasicTouchButtonCellPositionFactory = NewBasicTouchButtonCellPositionFactory(); _ = basicTouchButtonCellPositionFactory
            var commonButtons *CommonButtons = CommonButtons_GetInstance(); _ = commonButtons
            var animationInterface IAnimation = NullAnimationFactory_GetFactoryInstance().GetInstance(0); _ = animationInterface
            var UP ITouchButton = NewTouchButton(BasicTouchInputFactory_GetInstance().UP, animationInterface, animationInterface, commonButtons.NORMAL_BUTTON, basicTouchButtonCellPositionFactory.SECOND_FROM_BOTTOM_LEFT, touchButtonLocationHelper.GetColumnsRemainderHalf(), touchButtonLocationHelper.GetRowsRemainderHalf()); _ = UP
            var DOWN ITouchButton = NewTouchButton(BasicTouchInputFactory_GetInstance().DOWN, animationInterface, animationInterface, commonButtons.NORMAL_BUTTON, basicTouchButtonCellPositionFactory.BOTTOM_LEFT, touchButtonLocationHelper.GetColumnsRemainderHalf(), touchButtonLocationHelper.GetRowsRemainderHalf()); _ = DOWN
            var LEFT ITouchButton = NewTouchButton(BasicTouchInputFactory_GetInstance().LEFT, animationInterface, animationInterface, commonButtons.NORMAL_BUTTON, basicTouchButtonCellPositionFactory.BOTTOM_SECOND_FROM_RIGHT, touchButtonLocationHelper.GetColumnsRemainderHalf(), touchButtonLocationHelper.GetRowsRemainderHalf()); _ = LEFT
            var RIGHT ITouchButton = NewTouchButton(BasicTouchInputFactory_GetInstance().RIGHT, animationInterface, animationInterface, commonButtons.NORMAL_BUTTON, basicTouchButtonCellPositionFactory.BOTTOM_RIGHT, touchButtonLocationHelper.GetColumnsRemainderHalf(), touchButtonLocationHelper.GetRowsRemainderHalf()); _ = RIGHT
            list.Add(UP)
            list.Add(DOWN)
            list.Add(LEFT)
            list.Add(RIGHT)
            tryControl1, tryResult1 = 1, list; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

