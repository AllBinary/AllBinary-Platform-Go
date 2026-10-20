/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type PathAnimation struct {
    Animation
    LogUtil *LogUtil
    LinePathAnimation ILinePathAnimation
    PathFindingLayer PathFindingLayerInterface
    AllBinaryGameLayerManagerP IAllBinaryGameLayerManager
    StartColor *BasicColor
    MiddleColor *BasicColor
    EndColor *BasicColor
}


func (this *PathAnimation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    this.AllBinaryGameLayerManagerP = AllBinaryGameLayerManager_GetNullInstance()
    this.StartColor = BasicColorFactory_GetInstance().RED
    this.MiddleColor = BasicColorFactory_GetInstance().BLUE
    this.EndColor = BasicColorFactory_GetInstance().GREEN
}


func NewPathAnimation(pathFindingLayer PathFindingLayerInterface, linePathAnimation ILinePathAnimation) *PathAnimation {
    this := &PathAnimation{}
    this.Animation = *NewAnimation()
    this.initFields()
        this.PathFindingLayer = pathFindingLayer
        this.LinePathAnimation = linePathAnimation
    return this
}


// throws Exception
func (this *PathAnimation) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.AllBinaryGameLayerManagerP = allBinaryGameLayerManager
}


// throws Exception
func (this *PathAnimation) NextFrame() {
}


func (this *PathAnimation) PaintXY(graphics IGraphics, x int, y int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF("Cleared Target", this, CanvasStrings_GetInstance().PAINT)
                }
            }()
            var waypointBehaviorBase IWaypointBehaviorBase = this.PathFindingLayer.GetWaypointBehavior(); _ = waypointBehaviorBase
            if waypointBehaviorBase != nil {
                var geographicMapCellHistory *GeographicMapCellHistory = waypointBehaviorBase.GetCurrentGeographicMapCellHistory(); _ = geographicMapCellHistory
                var list IBasicArrayList = geographicMapCellHistory.GetTracked(); _ = list
                var size int = list.Size(); _ = size
                if size > 0 {
                    var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
                    var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
                    var tiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
                    var geographicMapCellPosition IGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION; _ = geographicMapCellPosition
                    var nextPoint IGPoint; _ = nextPoint
                    var point IGPoint; _ = point
                    var start int = size - 2; _ = start
                    for index := int(0); index < size - 1; index++ {
                        if index == start {
                            this.SetBasicColorP(this.StartColor)
                            this.BasicSetColorUtil.SetBasicColorP3(graphics, this.GetBasicColorP(), this.GetColor())
                        } else if geographicMapCellHistory.IsVisited(geographicMapCellPosition) {
                            this.SetBasicColorP(this.EndColor)
                            this.BasicSetColorUtil.SetBasicColorP3(graphics, this.GetBasicColorP(), this.GetColor())
                        } else {
                            this.SetBasicColorP(this.MiddleColor)
                            this.BasicSetColorUtil.SetBasicColorP3(graphics, this.GetBasicColorP(), this.GetColor())
                        }
                        geographicMapCellPosition = any(list.Get(index)).(IGeographicMapCellPosition)
                        point = geographicMapCellPosition.GetMidPoint()
                        geographicMapCellPosition = any(list.Get(index + 1)).(IGeographicMapCellPosition)
                        nextPoint = geographicMapCellPosition.GetMidPoint()
                        this.LinePathAnimation.Paint(graphics, point, nextPoint, tiledLayer)
                    }
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

