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

type LinesGraphicItem struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    GameInputStrings *GameInputStrings
    TreeNode *DefaultMutableTreeNode
    PointTreeNodeVector IBasicArrayList
    CurrentMousePoint IGPoint
    FulcrumPoint IGPoint
    Active bool
    Theta float64
    Color *Color
    Points *Points
}

// Implements GraphicItemInterface

var LinesGraphicItem_RECT string = "rect"
var LinesGraphicItem_ARC string = "arc"
var LinesGraphicItem_STRING string = "string"
var LinesGraphicItem_CHAR string = "char"
var LinesGraphicItem_IMAGE string = "image"
var LinesGraphicItem_Item int = 0
var LinesGraphicItem_NAME string = "Line Tool"

func (this *LinesGraphicItem) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.GameInputStrings = GameInputStrings_GetInstance()
    
    
    
    
    this.Active = false
    this.Theta = 0.0
    this.Color = NewColor(BasicColorFactory_GetInstance().WHITE.IntValue())
    
}




// Java overloads of new LinesGraphicItem
func NewLinesGraphicItem(args ...any) *LinesGraphicItem {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewLinesGraphicItem_1()
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewLinesGraphicItem_2(Java_Arg[*Node](args[0]))
        }
    }
    panic("No Java overload of new LinesGraphicItem for the arguments")
}
func NewLinesGraphicItem_1() *LinesGraphicItem {
    this := &LinesGraphicItem{}
    this.initFields()
        this.Points = NewPoints()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
                }
            }()
            this.Init()
            this.Active = true
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func NewLinesGraphicItem_2(linesNode *Node) *LinesGraphicItem {
    this := &LinesGraphicItem{}
    this.initFields()
        this.Points = NewPoints()
        this.Init()
        this.Active = false
        var lineNodes IBasicArrayList = DomHelper_GetInstance().GetWithoutTextNodes(linesNode.GetChildNodes()); _ = lineNodes
        var numberOfLines int = lineNodes.Size(); _ = numberOfLines
        for index := int(0); index < numberOfLines; index++ {
            var lineNode *Node = any(lineNodes.Get(index)).(*Node); _ = lineNode
            var pointNodes IBasicArrayList = DomHelper_GetInstance().GetWithoutTextNodes(lineNode.GetChildNodes()); _ = pointNodes
            var pointOneNode *Node = DomHelper_GetInstance().SearchNodeList(PointsDomUtil_GetInstance().POINTONE, pointNodes); _ = pointOneNode
            this.AddPoint(pointOneNode.GetChildNodes())
            if index == numberOfLines - 1 {
                var pointTwoNode *Node = DomHelper_GetInstance().SearchNodeList(PointsDomUtil_GetInstance().POINTTWO, pointNodes); _ = pointTwoNode
                this.AddPoint(pointTwoNode.GetChildNodes())
            }
        }
        if numberOfLines == 0 {
            panic(errors.New(fmt.Sprint("Lines node does not contain a line")))
        }
    return this
}


// throws Exception
func (this *LinesGraphicItem) Init() {
        this.TreeNode = NewDefaultMutableTreeNode(PointsDomUtil_GetInstance().LINES + LinesGraphicItem_Item)
        LinesGraphicItem_Item++
        this.Points.Init()
        this.PointTreeNodeVector = NewBasicArrayListD()
        this.FulcrumPoint = PointFactory_GetInstance().CreateXY(0, 0)
}


// throws Exception
func (this *LinesGraphicItem) Translate(x int, y int) {
        StatusFactory_GetInstance().SetStatus("Translating: " + fmt.Sprint(this.Points.GetPoints()))
        var basicGraphicsPipeline IBasicGraphicsPipeline = NewBasicGraphicsPipeline(this.Points.GetPoints()); _ = basicGraphicsPipeline
        basicGraphicsPipeline.Translate(x, y)
        this.Points = NewPoints()
        this.Points.AddPoints(basicGraphicsPipeline.GetMatrix())
}


func (this *LinesGraphicItem) SetRotate(theta float64) {
        for theta > 2 * Math_PI {
            theta -= 2 * Math_PI
        }
        for theta < 0.0 {
            theta += 2 * Math_PI
        }
        this.Theta = theta
}


func (this *LinesGraphicItem) AddRotate(theta float64) {
        this.SetRotate(this.Theta + theta)
}


func (this *LinesGraphicItem) SetAngle(angle float64) {
        this.SetRotate((float64(angle) * math.Pi / 180))
}


func (this *LinesGraphicItem) AddAngle(angle float64) {
        this.SetRotate(this.Theta + (float64(angle) * math.Pi / 180))
}


func (this *LinesGraphicItem) GetTreeNode() *MutableTreeNode {
        return any(this.TreeNode).(*MutableTreeNode)
}


func (this *LinesGraphicItem) GetName() string {
        return LinesGraphicItem_NAME
}


func LinesGraphicItem_GetStaticName() string {
        return LinesGraphicItem_NAME
}


// throws Exception
// synchronized
func (this *LinesGraphicItem) AddPoint_1(pointNodes *NodeList) {
        var list IBasicArrayList = DomHelper_GetInstance().GetWithoutTextNodes(pointNodes); _ = list
        var xNode *Node = DomHelper_GetInstance().SearchNodeList(PositionStrings_GetInstance().X, list); _ = xNode
        var xTextNode *Node = xNode.GetFirstChild(); _ = xTextNode
        var xInteger int = int(xTextNode.GetNodeValue()); _ = xInteger
        var yNode *Node = DomHelper_GetInstance().SearchNodeList(PositionStrings_GetInstance().Y, list); _ = yNode
        var yTextNode *Node = yNode.GetFirstChild(); _ = yTextNode
        var yInteger int = int(yTextNode.GetNodeValue()); _ = yInteger
        var point IGPoint = PointFactory_GetInstance().CreateXY(int(xInteger), int(yInteger)); _ = point
        this.AddPoint(point)
}

// Java overloads of LinesGraphicItem.addPoint
func (this *LinesGraphicItem) AddPoint(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*NodeList](args[0], pass) {
            this.AddPoint_1(Java_Arg[*NodeList](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[IGPoint](args[0], pass) {
            this.AddPoint_2(Java_Arg[IGPoint](args[0]))
            return
        }
    }
    panic("No Java overload of LinesGraphicItem.addPoint for the arguments")
}


// synchronized
func (this *LinesGraphicItem) AddPoint_2(point IGPoint) {
        this.Points.GetPoints().Add(point)
        this.LogUtil.PutF(fmt.Sprint(point), this, "addPoint")
        this.PointTreeNodeVector.Add(NewDefaultMutableTreeNode(fmt.Sprint(point)))
        this.TreeNode.Add(any(this.PointTreeNodeVector.Get(this.PointTreeNodeVector.Size() - 1)).(*DefaultMutableTreeNode))
}


// synchronized
func (this *LinesGraphicItem) RemovePoint() IGPoint {
        var point IGPoint = nil; _ = point
        if this.Points.GetSize() > 0 {
            var lastPoint int = this.Points.GetSize() - 1; _ = lastPoint
            point = any(this.Points.GetPoints().RemoveAt(lastPoint)).(IGPoint)
            var index int = this.PointTreeNodeVector.Size() - 1; _ = index
            if index > 0 {
                this.TreeNode.Remove(any(this.PointTreeNodeVector.Get(index)).(*DefaultMutableTreeNode))
                this.PointTreeNodeVector.Remove(index)
            }
        }
        return point
}


func (this *LinesGraphicItem) IsActive() bool {
        return this.Active
}


func (this *LinesGraphicItem) Deactivate() {
        this.Active = false
}


func (this *LinesGraphicItem) Activate() {
        this.Active = true
}


// throws Exception
func (this *LinesGraphicItem) DuplicatePoints(list IBasicArrayList) {
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var point IGPoint = any(list.Get(index)).(IGPoint); _ = point
            if point != nil {
                this.Points.GetPoints().Add(PointFactory_GetInstance().CreateXY(point.GetX(), point.GetY()))
                this.PointTreeNodeVector.Add(NewDefaultMutableTreeNode(fmt.Sprint(point)))
                this.TreeNode.Add(any(this.PointTreeNodeVector.Get(this.PointTreeNodeVector.Size() - 1)).(*DefaultMutableTreeNode))
            }
        }
}


// throws Exception
func (this *LinesGraphicItem) Duplicate() GraphicItemInterface {
        var linesGraphicItem *LinesGraphicItem = NewLinesGraphicItem(); _ = linesGraphicItem
        linesGraphicItem.DuplicatePoints(this.Points.GetPoints())
        linesGraphicItem.Deactivate()
        return linesGraphicItem
}


func (this *LinesGraphicItem) GetAngle() float64 {
        return (float64(this.Theta) * 180 / math.Pi)
}


func (this *LinesGraphicItem) SetFulcrumPoint(point IGPoint) {
        this.FulcrumPoint = point
}


func (this *LinesGraphicItem) Paint(g IGraphics, canvasAngle float64, dimension *IntegerDimension, x int, y int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "mouseMoved", e)
                }
            }()
            var graphics *Graphics2D = any(g).(*Graphics2D); _ = graphics
            graphics.SetColor(this.GetColor())
            graphics.SetStroke(NewBasicStroke(x))
            if this.IsActive() && this.CurrentMousePoint != nil {
                this.Points.GetPoints().Add(this.CurrentMousePoint)
            }
            var tempPointVector IBasicArrayList = PointsUtil_GetInstance().DoTransforms(this.Points.GetPoints(), canvasAngle, PointFactory_GetInstance().CreateXY(dimension.GetWidth(), dimension.GetHeight())); _ = tempPointVector
            if this.IsActive() && this.CurrentMousePoint != nil {
                this.CurrentMousePoint = any(this.Points.GetPoints().RemoveAt(this.Points.GetSize() - 1)).(IGPoint)
            }
            var size int = tempPointVector.Size(); _ = size
            var firstPoint IGPoint = nil; _ = firstPoint
            if size > 0 {
                firstPoint = any(tempPointVector.Get(0)).(IGPoint)
            }
            for index := int(1); index < size; index++ {
                var secondPoint IGPoint = any(tempPointVector.Get(index)).(IGPoint); _ = secondPoint
                graphics.DrawLine((firstPoint.GetX() * x) - (x / 4), (firstPoint.GetY() * y) - (y / 4), (secondPoint.GetX() * x) - (x / 4), (secondPoint.GetY() * y) - (y / 4))
                firstPoint = secondPoint
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *LinesGraphicItem) ToDom(canvasDom *CanvasDom) *Node {
        return PointsDomUtil_GetInstance().ToDom(canvasDom, this.Points.GetPoints())
}


func (this *LinesGraphicItem) IsValid() bool {
        return this.Points.IsValid()
}


func (this *LinesGraphicItem) GetPointsInterface() *Points {
        return this.Points
}


func (this *LinesGraphicItem) SetPointsInterface(points *Points) {
        this.Points = points
}


func (this *LinesGraphicItem) GetColor() *Color {
        return this.Color
}


func (this *LinesGraphicItem) SetColor(color *Color) {
        this.Color = color
}


func (this *LinesGraphicItem) MouseClicked(mouseEvent IMouseEvent, x int, y int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "mouseMoved", e)
                }
            }()
            var mousePoint IGPoint = PointFactory_GetInstance().CreateXY(mouseEvent.GetPoint().X, mouseEvent.GetPoint().Y); _ = mousePoint
            if (mouseEvent.GetModifiers() & mouseEvent.MouseEvent_().BUTTON1_MASK) == mouseEvent.MouseEvent_().BUTTON1_MASK {
                var point IGPoint = PointFactory_GetInstance().CreateXY(mousePoint.GetX() / x, mousePoint.GetY() / y); _ = point
                StatusFactory_GetInstance().SetStatus("Line Point Added: " + fmt.Sprint(point))
                this.AddPoint(point)
            } else if (mouseEvent.GetModifiers() & mouseEvent.MouseEvent_().BUTTON3_MASK) == mouseEvent.MouseEvent_().BUTTON3_MASK {
                StatusFactory_GetInstance().SetStatus("Line Point Removed")
                this.RemovePoint()
            }
            this.CurrentMousePoint = mousePoint
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *LinesGraphicItem) MouseEntered(mouseEvent IMouseEvent) {
}


func (this *LinesGraphicItem) MouseExited(mouseEvent IMouseEvent) {
}


func (this *LinesGraphicItem) MousePressed(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int) {
}


func (this *LinesGraphicItem) MouseReleased(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int) {
}


func (this *LinesGraphicItem) MouseDragged(mouseEvent IMouseEvent, xPixelsPerCell int, yPixelsPerCell int) {
}


func (this *LinesGraphicItem) MouseMoved(mouseEvent IMouseEvent, x int, y int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "mouseMoved", e)
                }
            }()
            var mousePoint IGPoint = PointFactory_GetInstance().CreateXY(mouseEvent.GetPoint().X, mouseEvent.GetPoint().Y); _ = mousePoint
            var point IGPoint = PointFactory_GetInstance().CreateXY(mousePoint.GetX() / x, mousePoint.GetY() / y); _ = point
            this.CurrentMousePoint = point
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *LinesGraphicItem) KeyPressed(keyEvent IKeyEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.KEY_PRESSED, e)
                }
            }()
            var keyCode int = keyEvent.GetKeyCode(); _ = keyCode
            if keyCode == keyEvent.KeyEvent_().VK_ESCAPE {
                StatusFactory_GetInstance().SetStatus("Deactivated")
                this.Deactivate()
            } else if keyCode == keyEvent.KeyEvent_().VK_UP {
                this.Translate(0, -1)
            } else if keyCode == keyEvent.KeyEvent_().VK_DOWN {
                this.Translate(0, 1)
            } else if keyCode == keyEvent.KeyEvent_().VK_LEFT {
                this.Translate(-1, 0)
            } else if keyCode == keyEvent.KeyEvent_().VK_RIGHT {
                this.Translate(1, 0)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *LinesGraphicItem) KeyReleased(keyEvent IKeyEvent) {
}


func (this *LinesGraphicItem) KeyTyped(keyEvent IKeyEvent) {
}

