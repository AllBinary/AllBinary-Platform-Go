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

type MotionGestureToMotionGestureActionAssociation struct {
    List IBasicArrayList
    MotionGestureAction ICompleteMotionGestureInputEvent
}


func (this *MotionGestureToMotionGestureActionAssociation) initFields() {
    
    
}


func NewMotionGestureToMotionGestureActionAssociation(list IBasicArrayList, commandAction ICompleteMotionGestureInputEvent) *MotionGestureToMotionGestureActionAssociation {
    this := &MotionGestureToMotionGestureActionAssociation{}
    this.initFields()
        this.List = list
        this.MotionGestureAction = commandAction
    return this
}


func (this *MotionGestureToMotionGestureActionAssociation) GetMotionGesture() IBasicArrayList {
        return this.List
}


func (this *MotionGestureToMotionGestureActionAssociation) SetMotionGestures(list IBasicArrayList) {
        this.List = list
}


func (this *MotionGestureToMotionGestureActionAssociation) GetCommandAction() ICompleteMotionGestureInputEvent {
        return this.MotionGestureAction
}


func (this *MotionGestureToMotionGestureActionAssociation) SetMotionGestureAction(commandAction ICompleteMotionGestureInputEvent) {
        this.MotionGestureAction = commandAction
}


func (this *MotionGestureToMotionGestureActionAssociation) IsMotionGestureArrayEquals(aList IBasicArrayList) bool {
        var size int = aList.Size(); _ = size
        if this.List.Size() != size {
            return false
        }
        for i := int(size - 1); i >= 0; i-- {
            if aList.BasicArrayList_().ObjectArray[i] != this.List.BasicArrayList_().ObjectArray[i] {
                return false
            }
        }
        return true
}

