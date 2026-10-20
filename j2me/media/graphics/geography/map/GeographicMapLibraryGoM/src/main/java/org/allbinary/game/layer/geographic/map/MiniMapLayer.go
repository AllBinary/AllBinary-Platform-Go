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

type MiniMapLayer struct {
    AllBinaryLayer
    AllBinaryTiledLayer IAllBinaryTiledLayer
    GeographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory
    List IBasicArrayList
    BasicColorList IBasicArrayList
    PositionList IBasicArrayList
    AnimationInterface *CircleFilledAnimation
}

// Implements GeographicMapCellPositionEventListenerInterface

// The Java class MiniMapLayer has subclasses so its variables are this interface
type IMiniMapLayer interface {
    IAllBinaryLayer
    MiniMapLayer_() *MiniMapLayer
    Init()
    OnEvent(eventObject IAllBinaryEventObject)
    OnRemoveGeographicMapCellPositionEvent(geographicMapCellPositionEvent *GeographicMapCellPositionEvent)
    OnGeographicMapCellPositionEvent(geographicMapCellPositionEvent *GeographicMapCellPositionEvent)
    PaintDots(graphics IGraphics)
    Paint(graphics IGraphics)
}

func (this *MiniMapLayer) MiniMapLayer_() *MiniMapLayer {
    return this
}


func (this *MiniMapLayer) initFields() {
    
    
    
    
    
    
}


func NewMiniMapLayer(geographicMapInterface IBasicGeographicMap, viewPosition IViewPositionBase) *MiniMapLayer {
    this := &MiniMapLayer{}
        this.AllBinaryLayer = *NewAllBinaryLayer(StringUtil_GetInstance().EMPTY_STRING, NewRectangle(PointFactory_GetInstance().CreateXY(0, viewPosition.GetY()), geographicMapInterface.GetAllBinaryTiledLayer().GetWidth(), geographicMapInterface.GetAllBinaryTiledLayer().GetHeight()), viewPosition)
        this.initFields()
        this.AllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer()
        this.GeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory()
        this.AnimationInterface = NewCircleFilledAnimation(this.AllBinaryTiledLayer.GetCellWidth(), this.AllBinaryTiledLayer.GetCellHeight(), BasicColorFactory_GetInstance().WHITE)
        this.List = NewBasicArrayListD()
        this.BasicColorList = NewBasicArrayListD()
        this.PositionList = NewBasicArrayListD()
        GeographicMapCellPositionEventHandler_GetInstance().AddListener(this)
        this.Init()
    return this
}


// throws Exception
func (this *MiniMapLayer) Init() {
        this.AllBinaryTiledLayer.SetPosition(this.X, this.Y, this.Z)
}


func (this *MiniMapLayer) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
// synchronized
func (this *MiniMapLayer) OnRemoveGeographicMapCellPositionEvent(geographicMapCellPositionEvent *GeographicMapCellPositionEvent) {
        var colorCompositeInterface ColorCompositeInterface = any(geographicMapCellPositionEvent.GetSource()).(ColorCompositeInterface); _ = colorCompositeInterface
        var index int = this.List.IndexOf(colorCompositeInterface); _ = index
        if index >= 0 {
            this.List.RemoveAt(index)
            this.BasicColorList.RemoveAt(index)
            this.PositionList.RemoveAt(index)
        }
}


// throws Exception
// synchronized
func (this *MiniMapLayer) OnGeographicMapCellPositionEvent(geographicMapCellPositionEvent *GeographicMapCellPositionEvent) {
        var colorCompositeInterface ColorCompositeInterface = any(geographicMapCellPositionEvent.GetSource()).(ColorCompositeInterface); _ = colorCompositeInterface
        var layerIndex int = -1; _ = layerIndex
        var size int = this.List.Size(); _ = size
        var nextColorCompositeInterface ColorCompositeInterface; _ = nextColorCompositeInterface
        for index := int(0); index < size; index++ {
            nextColorCompositeInterface = any(this.List.Get(index)).(ColorCompositeInterface)
            if nextColorCompositeInterface == colorCompositeInterface {
                layerIndex = index
                break
            }
        }
        var geographicMapCellPosition IGeographicMapCellPosition = this.GeographicMapCellPositionFactory.GetInstance(geographicMapCellPositionEvent.GetGeographicMapCellPosition()); _ = geographicMapCellPosition
        var point IGPoint = geographicMapCellPosition.GetPoint(); _ = point
        if layerIndex == -1 {
            this.List.Add(colorCompositeInterface)
            var basicColor *BasicColor = colorCompositeInterface.GetBasicColorP(); _ = basicColor
            this.BasicColorList.Add(basicColor)
            this.PositionList.Add(point)
            layerIndex = this.BasicColorList.Size() - 1
        } else {
            this.PositionList.Set(layerIndex, point)
        }
}


func (this *MiniMapLayer) PaintDots(graphics IGraphics) {
        var localPositionList IBasicArrayList = this.PositionList; _ = localPositionList
        var localBasicColorList IBasicArrayList = this.BasicColorList; _ = localBasicColorList
        var length int = localPositionList.Size(); _ = length
        var point IGPoint; _ = point
        var basicColor *BasicColor; _ = basicColor
        for index := int(length); Java_PreDecrement(&index) >= 0;  {
            point = any(localPositionList.Get(index)).(IGPoint)
            basicColor = any(localBasicColorList.Get(index)).(*BasicColor)
            graphics.SetColor(basicColor.IntValue())
            this.AnimationInterface.SetBasicColorP(basicColor)
            this.AnimationInterface.PaintXY(graphics, point.GetX() + this.X, point.GetY() + this.Y)
        }
}


func (this *MiniMapLayer) Paint(graphics IGraphics) {
        this.AllBinaryTiledLayer.Paint(graphics)
        this.PaintDots(graphics)
}

