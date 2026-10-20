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

type CanvasDom struct {
    GraphicsStrings *GraphicsStrings
    GraphicItemHashMap *HashMap
    Angle float64
    Dimension *IntegerDimension
    Grid *Grid
}

var CanvasDom_FRAME string = "frame"
var CanvasDom_ROTATE string = "rotate"
var CanvasDom_SIZE string = "size"
var CanvasDom_X string = "X"
var CanvasDom_Y string = "Y"
var CanvasDom_ZOOM string = "zoom"
var CanvasDom_GRID string = "grid"
var CanvasDom_ENABLED string = "enabled"
var CanvasDom_POSSIBLE string = "possible"
var CanvasDom_GRAPHICITEMS string = "graphicItems"
var CanvasDom_REAL_SIZE string = "realSize"

func (this *CanvasDom) initFields() {
    this.GraphicsStrings = GraphicsStrings_GetInstance()
    
    
    
    
}




// Java overloads of new CanvasDom
func NewCanvasDom(args ...any) *CanvasDom {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewCanvasDom_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*CanvasJPanel](args[0], pass) {
            return NewCanvasDom_2(Java_Arg[*CanvasJPanel](args[0]))
        }
    }
    panic("No Java overload of new CanvasDom for the arguments")
}
func NewCanvasDom_1(canvasNode *Node) *CanvasDom {
    this := &CanvasDom{}
    this.initFields()
        this.Grid = NewGrid()
        this.Grid.Grid = PointFactory_GetInstance().CreateXY(0, 0)
        this.Dimension = NewIntegerDimension(0, 0)
        var name string = canvasNode.GetNodeName(); _ = name
        if strings.Compare(name, CanvasDom_FRAME) == 0 {
            this.GraphicItemHashMap = StdUtil_GetInstance().CreateHashMap()
            var angleNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_ROTATE, canvasNode.GetChildNodes()); _ = angleNode
            var angleTextNode *Node = angleNode.GetFirstChild(); _ = angleTextNode
            var angle float64 = float64(float64(angleTextNode.GetNodeValue())); _ = angle
            this.SetAngle(angle)
            var sizeNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_SIZE, canvasNode.GetChildNodes()); _ = sizeNode
            var xNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_X, sizeNode.GetChildNodes()); _ = xNode
            var xTextNode *Node = xNode.GetFirstChild(); _ = xTextNode
            var x int = int(int(xTextNode.GetNodeValue())); _ = x
            var yNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_Y, sizeNode.GetChildNodes()); _ = yNode
            var yTextNode *Node = yNode.GetFirstChild(); _ = yTextNode
            var y int = int(int(yTextNode.GetNodeValue())); _ = y
            this.SetWorkArea(x, y)
            var zoomNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_ZOOM, canvasNode.GetChildNodes()); _ = zoomNode
            var zoomTextNode *Node = zoomNode.GetFirstChild(); _ = zoomTextNode
            this.Grid.SetZoom(int(int(zoomTextNode.GetNodeValue())))
            var gridNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_GRID, canvasNode.GetChildNodes()); _ = gridNode
            var gridSizeNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_SIZE, gridNode.GetChildNodes()); _ = gridSizeNode
            var gridXNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_X, gridSizeNode.GetChildNodes()); _ = gridXNode
            var gridXTextNode *Node = gridXNode.GetFirstChild(); _ = gridXTextNode
            var gridX int = int(int(gridXTextNode.GetNodeValue())); _ = gridX
            var gridYNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_Y, gridSizeNode.GetChildNodes()); _ = gridYNode
            var gridYTextNode *Node = gridYNode.GetFirstChild(); _ = gridYTextNode
            var gridY int = int(int(gridYTextNode.GetNodeValue())); _ = gridY
            this.SetGrid(gridX, gridY)
            var enableNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_ENABLED, gridNode.GetChildNodes()); _ = enableNode
            var enableTextNode *Node = enableNode.GetFirstChild(); _ = enableTextNode
            this.Grid.ShowGrid(enableTextNode.GetNodeValue())
            var possibleNode *Node = DomHelper_GetInstance().SearchNodeList(CanvasDom_POSSIBLE, gridNode.GetChildNodes()); _ = possibleNode
            var possibleTextNode *Node = possibleNode.GetFirstChild(); _ = possibleTextNode
            this.Grid.IsGridPossible = possibleTextNode.GetNodeValue()
            var graphicItemNodeList IBasicArrayList = DomHelper_GetInstance().GetChildrenWithoutTextNodes(CanvasDom_GRAPHICITEMS, canvasNode.GetChildNodes()); _ = graphicItemNodeList
            if graphicItemNodeList != nil {
                this.GraphicItemHashMap = GraphicItemFactory_GetInstance().GetInstance(graphicItemNodeList).(GraphicsItemInterfaceFactoryInterface)
            }
        } else {
            panic(errors.New(fmt.Sprint("Frame Element Not Found but Found: " + name)))
        }
    return this
}


func NewCanvasDom_2(canvasJPanel *CanvasJPanel) *CanvasDom {
    this := &CanvasDom{}
    this.initFields()
        this.Angle = canvasJPanel.GetAngle()
        this.GraphicItemHashMap = canvasJPanel.GetGraphicItemHashMap()
        this.Dimension = canvasJPanel.GetCanvasDimension()
        this.Grid = NewGrid(canvasJPanel.GetGrid())
        this.Angle = canvasJPanel.GetAngle()
    return this
}


func (this *CanvasDom) GetAngle() float64 {
        return this.Angle
}


func (this *CanvasDom) GetGraphicItemHashMap() *HashMap {
        return this.GraphicItemHashMap
}


func (this *CanvasDom) SetAngle(angle float64) {
        this.Angle = angle
}


// throws Exception
func (this *CanvasDom) SetWorkArea(x int, y int) {
        this.SetDimension(NewIntegerDimension(x, y))
}


// throws Exception
func (this *CanvasDom) SetGrid(gridX int, gridY int) {
        this.Grid.Grid = PointFactory_GetInstance().CreateXY(gridX, gridY)
}


func (this *CanvasDom) GetPointNode(point IGPoint, pointName string) *Node {
        var document *Document = WorkAreaJPanel_GetDocument(); _ = document
        var pointNode *Node = any(document.CreateElement(pointName)).(*Node); _ = pointNode
        var xNode *Node = any(document.CreateElement(CanvasDom_X)).(*Node); _ = xNode
        var xTextNode *Node = any(document.CreateTextNode(fmt.Sprint(point.GetX()))).(*Node); _ = xTextNode
        xNode.AppendChild(xTextNode)
        var yNode *Node = any(document.CreateElement(CanvasDom_Y)).(*Node); _ = yNode
        var yTextNode *Node = any(document.CreateTextNode(fmt.Sprint(point.GetY()))).(*Node); _ = yTextNode
        yNode.AppendChild(yTextNode)
        pointNode.AppendChild(xNode)
        pointNode.AppendChild(yNode)
        return pointNode
}


// throws Exception
func (this *CanvasDom) ToDom() *Node {
        var document *Document = WorkAreaJPanel_GetDocument(); _ = document
        var frameNode *Node = any(document.CreateElement(CanvasDom_FRAME)).(*Node); _ = frameNode
        var angleNode *Node = any(document.CreateElement(CanvasDom_ROTATE)).(*Node); _ = angleNode
        var angleTextNode *Node = any(document.CreateTextNode(fmt.Sprint(float64(this.GetAngle())))).(*Node); _ = angleTextNode
        angleNode.AppendChild(angleTextNode)
        var sizeNode *Node = any(this.GetPointNode(PointFactory_GetInstance().CreateXY(this.GetDimension().GetWidth(), this.GetDimension().GetHeight()), CanvasDom_SIZE)).(*Node); _ = sizeNode
        var zoomNode *Node = any(document.CreateElement(CanvasDom_ZOOM)).(*Node); _ = zoomNode
        var zoomTextNode *Node = any(document.CreateTextNode(fmt.Sprint(int(this.GetGrid().GetZoom())))).(*Node); _ = zoomTextNode
        zoomNode.AppendChild(zoomTextNode)
        var gridNode *Node = any(document.CreateElement(CanvasDom_GRID)).(*Node); _ = gridNode
        var gridSizeNode *Node = any(this.GetPointNode(this.GetGrid().Grid, CanvasDom_SIZE)).(*Node); _ = gridSizeNode
        gridNode.AppendChild(gridSizeNode)
        var enableNode *Node = any(document.CreateElement(CanvasDom_ENABLED)).(*Node); _ = enableNode
        var enableTextNode *Node = any(document.CreateTextNode(fmt.Sprint(this.GetGrid().IsGridOn()))).(*Node); _ = enableTextNode
        enableNode.AppendChild(enableTextNode)
        gridNode.AppendChild(enableNode)
        var possibleNode *Node = any(document.CreateElement(CanvasDom_POSSIBLE)).(*Node); _ = possibleNode
        var possibleTextNode *Node = any(document.CreateTextNode(fmt.Sprint(this.GetGrid().IsGridPossible()))).(*Node); _ = possibleTextNode
        possibleNode.AppendChild(possibleTextNode)
        gridNode.AppendChild(possibleNode)
        var realSizeNode *Node = any(document.CreateElement(CanvasDom_REAL_SIZE)).(*Node); _ = realSizeNode
        var vectorCenterGenerator *VectorCenterGenerator = NewVectorCenterGenerator(); _ = vectorCenterGenerator
        vectorCenterGenerator.Calculate(this.GetGraphicItemHashMap())
        var widthNode *Node = any(document.CreateElement(this.GraphicsStrings.WIDTH)).(*Node); _ = widthNode
        var widthTextNode *Node = any(document.CreateTextNode(fmt.Sprint(vectorCenterGenerator.GetWidth()))).(*Node); _ = widthTextNode
        widthNode.AppendChild(widthTextNode)
        realSizeNode.AppendChild(widthNode)
        var heightNode *Node = any(document.CreateElement(this.GraphicsStrings.HEIGHT)).(*Node); _ = heightNode
        var heightTextNode *Node = any(document.CreateTextNode(fmt.Sprint(vectorCenterGenerator.GetHeight()))).(*Node); _ = heightTextNode
        heightNode.AppendChild(heightTextNode)
        realSizeNode.AppendChild(heightNode)
        frameNode.AppendChild(angleNode)
        frameNode.AppendChild(sizeNode)
        frameNode.AppendChild(zoomNode)
        frameNode.AppendChild(gridNode)
        frameNode.AppendChild(realSizeNode)
        var graphicItemNode *Node = any(document.CreateElement(CanvasDom_GRAPHICITEMS)).(*Node); _ = graphicItemNode
        var graphicItemArray []any = this.GetGraphicItemHashMap().KeySet().ToArray(); _ = graphicItemArray
        var size int = len(graphicItemArray); _ = size
        for index := int(0); index < size; index++ {
            var item GraphicItemInterface = any(this.GraphicItemHashMap.Get(graphicItemArray[index])).(GraphicItemInterface); _ = item
            if item.GetName() == LinesGraphicItem_GetStaticName() {
                if item.IsValid() {
                    graphicItemNode.AppendChild(item.ToDom(this))
                }
            }
        }
        frameNode.AppendChild(graphicItemNode)
        return frameNode
}


func (this *CanvasDom) GetDimension() *IntegerDimension {
        return this.Dimension
}


func (this *CanvasDom) SetDimension(dimension *IntegerDimension) {
        this.Dimension = dimension
}


func (this *CanvasDom) GetGrid() *Grid {
        return this.Grid
}

