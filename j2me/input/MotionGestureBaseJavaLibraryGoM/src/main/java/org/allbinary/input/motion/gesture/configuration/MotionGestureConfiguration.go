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

type MotionGestureConfiguration struct {
    DiagonalMotionGestureAllowed bool
    DiagonalTolerance int
    MinimumMotionGesture int
    Button string
    Color *BasicColor
    ExecutingActions bool
    MouseButtonMask int
    ActiveCommands IBasicArrayList
}


func (this *MotionGestureConfiguration) initFields() {
    this.DiagonalMotionGestureAllowed = true
    this.DiagonalTolerance = 7
    this.MinimumMotionGesture = 9
    this.Button = StringUtil_GetInstance().EMPTY_STRING
    this.Color = BasicColorFactory_GetInstance().BLUE
    this.ExecutingActions = true
    this.MouseButtonMask = 0
    this.ActiveCommands = NewBasicArrayListD()
}


func NewMotionGestureConfiguration() *MotionGestureConfiguration {
    this := &MotionGestureConfiguration{}
    this.initFields()
    return this
}


func (this *MotionGestureConfiguration) AddMotionGestureAction(list IBasicArrayList, commandAction ICompleteMotionGestureInputEvent) {
        this.ActiveCommands.Add(NewMotionGestureToMotionGestureActionAssociation(list, commandAction))
}


func (this *MotionGestureConfiguration) ClearMotionGestureAction() {
        this.ActiveCommands.Clear()
}


func (this *MotionGestureConfiguration) GetAssociateCommandActionsList() IBasicArrayList {
        return this.ActiveCommands
}


func (this *MotionGestureConfiguration) IsDiagonalMotionGestureAllowed() bool {
        return this.DiagonalMotionGestureAllowed
}


func (this *MotionGestureConfiguration) IsExecutingActionsAllowed() bool {
        return this.ExecutingActions
}


func (this *MotionGestureConfiguration) SetExecutingActionsAllowed(executingActionAllowed bool) {
        this.ExecutingActions = executingActionAllowed
}


func (this *MotionGestureConfiguration) SetDiagonalMotionGestureAllowed(diagonalMotionGestureAllowed bool) {
        this.DiagonalMotionGestureAllowed = diagonalMotionGestureAllowed
}


func (this *MotionGestureConfiguration) GetDiagonalTolerance() int {
        return this.DiagonalTolerance
}


func (this *MotionGestureConfiguration) SetDiagonalTolerance(diagonalTolerance int) {
        this.DiagonalTolerance = diagonalTolerance
}


func (this *MotionGestureConfiguration) GetMinimumMotionGesture() int {
        return this.MinimumMotionGesture
}


func (this *MotionGestureConfiguration) SetMinimumMotionGesture(minimumMotionGesture int) {
        this.MinimumMotionGesture = minimumMotionGesture
}


func (this *MotionGestureConfiguration) GetButton() string {
        return this.Button
}


func (this *MotionGestureConfiguration) SetButton(button string) {
        this.Button = button
}


func (this *MotionGestureConfiguration) GetColor() *BasicColor {
        return this.Color
}


func (this *MotionGestureConfiguration) SetColor(color *BasicColor) {
        this.Color = color
}


func (this *MotionGestureConfiguration) GetMouseButtonMask() int {
        return this.MouseButtonMask
}


func (this *MotionGestureConfiguration) SetMouseButtonMask(mouseButtonMask int) {
        this.MouseButtonMask = mouseButtonMask
}

