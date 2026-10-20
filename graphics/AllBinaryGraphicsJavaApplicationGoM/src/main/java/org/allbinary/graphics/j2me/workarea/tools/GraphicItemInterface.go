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

type GraphicItemInterface interface {
    GetName() string
    SetColor(color *Color)
    GetTreeNode() *MutableTreeNode
    GetPointsInterface() *Points
    SetPointsInterface(points *Points)
    IsValid() bool
    IsActive() bool
    Deactivate()
    Activate()
    SetFulcrumPoint(point IGPoint)
    Translate(x int, y int)
    SetRotate(theta float64)
    AddRotate(theta float64)
    SetAngle(angle float64)
    AddAngle(angle float64)
    GetAngle() float64
    Paint(g IGraphics, canvasAngle float64, dimension *IntegerDimension, x int, y int)
    ToDom(canvasDom *CanvasDom) *Node
    RemovePoint() IGPoint
    AddPoint(point IGPoint)
    MouseClicked(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int)
    MouseEntered(mouseEvent IMouseEvent)
    MouseExited(mouseEvent IMouseEvent)
    MousePressed(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int)
    MouseReleased(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int)
    MouseDragged(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int)
    MouseMoved(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int)
    KeyPressed(keyEvent IKeyEvent)
    KeyReleased(keyEvent IKeyEvent)
    KeyTyped(keyEvent IKeyEvent)
    Duplicate() GraphicItemInterface
}


