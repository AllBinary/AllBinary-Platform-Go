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

type SelectionTool struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CanvasJPanel *CanvasJPanel
    Active bool
    ControlIdDown bool
    RectColor *Color
    StartPoint IGPoint
    EndPoint IGPoint
}

// Implements GraphicItemInterface

var SelectionTool_NAME string = "Selection Tool"

func (this *SelectionTool) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    this.Active = true
    
    this.RectColor = NewColor(BasicColorFactory_GetInstance().RED.IntValue())
    this.StartPoint = PointFactory_GetInstance().ZERO_ZERO
    this.EndPoint = this.StartPoint
}


func NewSelectionTool(canvasJPanel *CanvasJPanel) *SelectionTool {
    this := &SelectionTool{}
    this.initFields()
        this.CanvasJPanel = canvasJPanel
    return this
}


func (this *SelectionTool) SetColor(color *Color) {
}


func (this *SelectionTool) GetTreeNode() *MutableTreeNode {
        return nil
}


func (this *SelectionTool) GetPointsInterface() *Points {
        return nil
}


func (this *SelectionTool) SetPointsInterface(points *Points) {
}


func (this *SelectionTool) IsValid() bool {
        return true
}


func (this *SelectionTool) IsActive() bool {
        return this.Active
}


func (this *SelectionTool) Deactivate() {
}


func (this *SelectionTool) Activate() {
}


func (this *SelectionTool) SetFulcrumPoint(point IGPoint) {
}


// throws Exception
func (this *SelectionTool) Translate(x int, y int) {
}


func (this *SelectionTool) SetRotate(theta float64) {
}


func (this *SelectionTool) AddRotate(theta float64) {
}


func (this *SelectionTool) SetAngle(angle float64) {
}


func (this *SelectionTool) AddAngle(angle float64) {
}


func (this *SelectionTool) GetAngle() float64 {
        return 0.0
}


func (this *SelectionTool) MakeSelection() {
        var xRect int = this.GetXRect(); _ = xRect
        var xRect2 int = max(this.EndPoint.GetX() - this.StartPoint.GetX(), -(this.EndPoint.GetX() - this.StartPoint.GetX())) + xRect; _ = xRect2
        var yRect int = this.GetYRect(); _ = yRect
        var yRect2 int = max(this.EndPoint.GetY() - this.StartPoint.GetY(), -(this.EndPoint.GetY() - this.StartPoint.GetY())) + yRect; _ = yRect2
        var hashMap *HashMap = this.CanvasJPanel.GetGraphicItemHashMap(); _ = hashMap
        var graphicItemArray []any = hashMap.KeySet().ToArray(); _ = graphicItemArray
        var size int = len(graphicItemArray); _ = size
        for index := int(0); index < size; index++ {
            var item GraphicItemInterface = any(hashMap.Get(graphicItemArray[index])).(GraphicItemInterface); _ = item
            if item.GetName() == LinesGraphicItem_GetStaticName() {
                if RectangleCollisionPointUtil_AllPointsInsideCell(xRect, yRect, xRect2, yRect2, item.GetPointsInterface().GetPoints(), this.CanvasJPanel.GetXPixelsPerWorkAreaPixel(), this.CanvasJPanel.GetYPixelsPerWorkAreaPixel()) {
                    item.SetColor(Color_YELLOW)
                    item.Activate()
                } else if !this.ControlIdDown {
                    item.SetColor(Color_WHITE)
                    item.Deactivate()
                }
            }
        }
}


func (this *SelectionTool) GetXRect() int {
        var xDiff int = this.EndPoint.GetX() - this.StartPoint.GetX(); _ = xDiff
        var xRect int = this.StartPoint.GetX(); _ = xRect
        if xDiff < 0 {
            xRect = this.EndPoint.GetX()
        }
        return xRect
}


func (this *SelectionTool) GetYRect() int {
        var yDiff int = this.EndPoint.GetY() - this.StartPoint.GetY(); _ = yDiff
        var yRect int = this.StartPoint.GetY(); _ = yRect
        if yDiff < 0 {
            yRect = this.EndPoint.GetY()
        }
        return yRect
}


func (this *SelectionTool) Paint(g IGraphics, canvasAngle float64, dimension *IntegerDimension, x int, y int) {
        g.SetColor(this.RectColor)
        var xRect int = this.GetXRect(); _ = xRect
        var width int = max(this.EndPoint.GetX() - this.StartPoint.GetX(), -(this.EndPoint.GetX() - this.StartPoint.GetX())); _ = width
        var yRect int = this.GetYRect(); _ = yRect
        var height int = max(this.EndPoint.GetY() - this.StartPoint.GetY(), -(this.EndPoint.GetY() - this.StartPoint.GetY())); _ = height
        g.DrawRect(xRect, yRect, width, height)
}


// throws Exception
func (this *SelectionTool) ToDom(canvasDom *CanvasDom) *Node {
        return nil
}


func (this *SelectionTool) RemovePoint() IGPoint {
        return nil
}


func (this *SelectionTool) AddPoint(point IGPoint) {
}


func (this *SelectionTool) GetName() string {
        return SelectionTool_NAME
}


func SelectionTool_GetStaticName() string {
        return SelectionTool_NAME
}


func (this *SelectionTool) MouseClicked(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int) {
}


func (this *SelectionTool) MouseMoved(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int) {
}


func (this *SelectionTool) MouseEntered(mouseEvent IMouseEvent) {
}


func (this *SelectionTool) MouseExited(mouseEvent IMouseEvent) {
}


func (this *SelectionTool) MousePressed(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int) {
        this.LogUtil.PutF(this.CommonStrings.START, this, MouseStrings_GetInstance().MOUSE_PRESSED)
        this.StartPoint = PointFactory_GetInstance().CreateXY(mouseEvent.GetX(), mouseEvent.GetY())
        this.EndPoint = this.StartPoint
}


func (this *SelectionTool) MouseReleased(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int) {
        this.LogUtil.PutF(this.CommonStrings.START, this, MouseStrings_GetInstance().MOUSE_RELEASED)
        this.EndPoint = PointFactory_GetInstance().CreateXY(mouseEvent.GetX(), mouseEvent.GetY())
        this.Active = false
        this.MakeSelection()
}


func (this *SelectionTool) MouseDragged(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int) {
        this.EndPoint = PointFactory_GetInstance().CreateXY(mouseEvent.GetX(), mouseEvent.GetY())
}


func (this *SelectionTool) KeyPressed(keyEvent IKeyEvent) {
        if keyEvent.GetKeyCode() == keyEvent.KeyEvent_().VK_CONTROL {
            this.ControlIdDown = true
        }
}


func (this *SelectionTool) KeyReleased(keyEvent IKeyEvent) {
        if keyEvent.GetKeyCode() == keyEvent.KeyEvent_().VK_CONTROL {
            this.ControlIdDown = false
        }
}


func (this *SelectionTool) KeyTyped(keyEvent IKeyEvent) {
}


// throws Exception
func (this *SelectionTool) Duplicate() GraphicItemInterface {
        return nil
}

