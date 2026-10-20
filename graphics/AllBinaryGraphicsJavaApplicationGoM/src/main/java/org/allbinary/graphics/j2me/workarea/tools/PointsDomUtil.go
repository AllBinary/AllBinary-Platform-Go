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

type PointsDomUtil struct {
    LINES string
    POINTONE string
    POINTTWO string
    ROTATEDPOINTONE string
    ROTATEDPOINTTWO string
    LINE string
}

var PointsDomUtil_Instance *PointsDomUtil = NewPointsDomUtil()

func (this *PointsDomUtil) initFields() {
    this.LINES = "lines"
    this.POINTONE = "pointOne"
    this.POINTTWO = "pointTwo"
    this.ROTATEDPOINTONE = "rotatedPointOne"
    this.ROTATEDPOINTTWO = "rotatedPointTwo"
    this.LINE = "line"
}

func NewPointsDomUtil() *PointsDomUtil {
    this := &PointsDomUtil{}
    this.initFields()
    return this
}


func PointsDomUtil_GetInstance() *PointsDomUtil {
        return PointsDomUtil_Instance
}


// throws Exception
func (this *PointsDomUtil) ToDom(canvasDom *CanvasDom, pointVector IBasicArrayList) *Node {
        var tempPointVector IBasicArrayList = PointsUtil_GetInstance().DoTransforms(pointVector, float64(canvasDom.GetAngle()), PointFactory_GetInstance().CreateXY(canvasDom.GetDimension().GetWidth(), canvasDom.GetDimension().GetHeight())); _ = tempPointVector
        var document *Document = WorkAreaJPanel_GetDocument(); _ = document
        var linesNode *Node = any(document.CreateElement(this.LINES)).(*Node); _ = linesNode
        var size2 int = tempPointVector.Size(); _ = size2
        var size int = pointVector.Size(); _ = size
        var firstPoint IGPoint = nil; _ = firstPoint
        var rotatedFirstPoint IGPoint = nil; _ = rotatedFirstPoint
        var index int = 0; _ = index
        if index < size && index < size2 {
            rotatedFirstPoint = any(tempPointVector.Get(index)).(IGPoint)
            firstPoint = any(pointVector.Get(index)).(IGPoint)
            index++
        }
        for index < size && index < size2 {
            var lineNode *Node = any(document.CreateElement(this.LINE)).(*Node); _ = lineNode
            var rotatedSecondPoint IGPoint = any(tempPointVector.Get(index)).(IGPoint); _ = rotatedSecondPoint
            var secondPoint IGPoint = any(pointVector.Get(index)).(IGPoint); _ = secondPoint
            var pointOneNode *Node = this.GetPointNode(firstPoint, this.POINTONE); _ = pointOneNode
            var pointTwoNode *Node = this.GetPointNode(secondPoint, this.POINTTWO); _ = pointTwoNode
            var rotatedPointOneNode *Node = this.GetPointNode(rotatedFirstPoint, this.ROTATEDPOINTONE); _ = rotatedPointOneNode
            var rotatedPointTwoNode *Node = this.GetPointNode(rotatedSecondPoint, this.ROTATEDPOINTTWO); _ = rotatedPointTwoNode
            lineNode.AppendChild(pointOneNode)
            lineNode.AppendChild(pointTwoNode)
            lineNode.AppendChild(rotatedPointOneNode)
            lineNode.AppendChild(rotatedPointTwoNode)
            firstPoint = secondPoint
            rotatedFirstPoint = rotatedSecondPoint
            linesNode.AppendChild(lineNode)
            index++
        }
        return linesNode
}


func (this *PointsDomUtil) GetPointNode(point IGPoint, pointName string) *Node {
        var document *Document = WorkAreaJPanel_GetDocument(); _ = document
        var pointNode *Node = any(document.CreateElement(pointName)).(*Node); _ = pointNode
        var xNode *Node = any(document.CreateElement(PositionStrings_GetInstance().X)).(*Node); _ = xNode
        var xTextNode *Node = any(document.CreateTextNode(fmt.Sprint(point.GetX()))).(*Node); _ = xTextNode
        xNode.AppendChild(xTextNode)
        var yNode *Node = any(document.CreateElement(PositionStrings_GetInstance().Y)).(*Node); _ = yNode
        var yTextNode *Node = any(document.CreateTextNode(fmt.Sprint(point.GetY()))).(*Node); _ = yTextNode
        yNode.AppendChild(yTextNode)
        pointNode.AppendChild(xNode)
        pointNode.AppendChild(yNode)
        return pointNode
}

