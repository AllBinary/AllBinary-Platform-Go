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

type GeographicMapCellHistory struct {
    LogUtil *LogUtil
    BooleanFactory *BooleanFactory
    List IBasicArrayList
    VisitedList IBasicArrayList
    MISSED_INFO string
    Animation IAnimation
    MyFontProcessor IMyFontProcessor
    TotalVisited int
    HalfWidth int
    FontHeight int
    RED int
}

// Implements UpdateMyFontInterface

var GeographicMapCellHistory_NULL_GEOGRPAHIC_MAP_HISTORY_ARRAY []*GeographicMapCellHistory = make([]*GeographicMapCellHistory, 0)

func (this *GeographicMapCellHistory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BooleanFactory = BooleanFactory_GetInstance()
    
    
    this.MISSED_INFO = "Missed"
    this.Animation = NewTextAnimation(this.MISSED_INFO, AnimationBehavior_GetInstance())
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    
    this.HalfWidth = 0
    this.FontHeight = 0
    this.RED = BasicColorFactory_GetInstance().RED.IntValue()
}


func NewGeographicMapCellHistory() *GeographicMapCellHistory {
    this := &GeographicMapCellHistory{}
    this.initFields()
        this.List = NewBasicArrayListD()
        this.VisitedList = NewBasicArrayListD()
        this.Init()
    return this
}


func (this *GeographicMapCellHistory) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.FontHeight = font.GetHeight()
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *GeographicMapCellHistory) GetTotalVisited() int {
        return this.TotalVisited
}


func (this *GeographicMapCellHistory) GetTotalNotVisited() int {
        return this.GetSize() - this.TotalVisited
}


func (this *GeographicMapCellHistory) GetSize() int {
        return this.List.Size()
}


func (this *GeographicMapCellHistory) TrackAll(list IBasicArrayList) {
        var size int = list.Size(); _ = size
        this.List.EnsureCapacity(size)
        this.VisitedList.EnsureCapacity(size)
        var geographicMapCellPosition IGeographicMapCellPosition; _ = geographicMapCellPosition
        for index := int(0); index < size; index++ {
            geographicMapCellPosition = any(list.Get(index)).(IGeographicMapCellPosition)
            this.Track(geographicMapCellPosition)
        }
}


func (this *GeographicMapCellHistory) Track(geographicMapCellPosition IGeographicMapCellPosition) {
        if !this.List.Contains(geographicMapCellPosition) {
            this.List.Add(geographicMapCellPosition)
            this.VisitedList.Add(this.BooleanFactory.FALSE)
        }
}


func (this *GeographicMapCellHistory) GetTracked() IBasicArrayList {
        return this.List
}


func (this *GeographicMapCellHistory) GetVisited() IBasicArrayList {
        return this.VisitedList
}


func (this *GeographicMapCellHistory) GetAfterIfNotLast(geographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition {
        var localList IBasicArrayList = this.List; _ = localList
        var index int = localList.IndexOf(geographicMapCellPosition); _ = index
        if localList.Size() > index + 1 {
            return any(localList.Get(index + 1)).(IGeographicMapCellPosition)
        }
        return geographicMapCellPosition
}


func (this *GeographicMapCellHistory) GetFirstUnvisited() IGeographicMapCellPosition {
        var localList IBasicArrayList = this.List; _ = localList
        var localVisitedList IBasicArrayList = this.VisitedList; _ = localVisitedList
        var size int = localVisitedList.Size(); _ = size
        var value bool; _ = value
        for index := int(0); index < size; index++ {
            value = any(this.VisitedList.Get(index)).(bool)
            if value == this.BooleanFactory.FALSE {
                return any(localList.Get(index)).(IGeographicMapCellPosition)
            }
        }
        return any(localList.Get(0)).(IGeographicMapCellPosition)
}


func (this *GeographicMapCellHistory) GetFirstUnvisitedIndex() int {
        var localVisitedList IBasicArrayList = this.VisitedList; _ = localVisitedList
        var size int = localVisitedList.Size(); _ = size
        var value bool; _ = value
        for index := int(0); index < size; index++ {
            value = any(this.VisitedList.Get(index)).(bool)
            if value == this.BooleanFactory.FALSE {
                return index
            }
        }
        return 0
}


func (this *GeographicMapCellHistory) GetInPathButNotTracked(pathList IBasicArrayList) IBasicArrayList {
        var inPathButNotTrackedList IBasicArrayList = NewBasicArrayListD(); _ = inPathButNotTrackedList
        var localList IBasicArrayList = this.List; _ = localList
        var size int = pathList.Size(); _ = size
        var geographicMapCellPosition IGeographicMapCellPosition; _ = geographicMapCellPosition
        for index := int(0); index < size; index++ {
            geographicMapCellPosition = any(pathList.Get(index)).(IGeographicMapCellPosition)
            if !localList.Contains(geographicMapCellPosition) {
                inPathButNotTrackedList.Add(geographicMapCellPosition)
            }
        }
        return inPathButNotTrackedList
}


func (this *GeographicMapCellHistory) IsVisited(geographicMapCellPosition IGeographicMapCellPosition) bool {
        var index int = this.List.IndexOf(geographicMapCellPosition); _ = index
        if index != -1 {
            var value bool = any(this.VisitedList.Get(index)).(bool); _ = value
            if value == this.BooleanFactory.TRUE {
                return true
            }
        } else {
        }
        return false
}


func (this *GeographicMapCellHistory) Visit(geographicMapCellPosition IGeographicMapCellPosition) bool {
        var index int = this.List.IndexOf(geographicMapCellPosition); _ = index
        var value bool; _ = value
        if index != -1 {
            value = any(this.VisitedList.Get(index)).(bool)
            var TRUE bool = this.BooleanFactory.TRUE; _ = TRUE
            if value != TRUE {
                this.VisitedList.Set(index, TRUE)
                this.TotalVisited++
                return true
            }
        } else {
        }
        return false
}


// throws Exception
func (this *GeographicMapCellHistory) IsAllVisited() bool {
        if this.TotalVisited == this.GetSize() - 1 {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *GeographicMapCellHistory) IsAllVisited2() bool {
        if this.TotalVisited == this.GetSize() {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *GeographicMapCellHistory) IsPortionVisited(basicDecimal *BasicDecimal) bool {
        var size int = this.GetSize(); _ = size
        var numberRequired int = (size << basicDecimal.GetScaledFactor()) / int(basicDecimal.GetUnscaled()); _ = numberRequired
        var numberNotVisited int = this.GetSize() - 1 - this.TotalVisited; _ = numberNotVisited
        if size - numberNotVisited > numberRequired {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *GeographicMapCellHistory) Reset() {
        var localVisitedList IBasicArrayList = this.VisitedList; _ = localVisitedList
        var localFalseBoolean bool = this.BooleanFactory.FALSE; _ = localFalseBoolean
        var size int = localVisitedList.Size(); _ = size
        for index := int(size - 1); index >= 0; index-- {
            localVisitedList.Set(index, localFalseBoolean)
        }
        this.TotalVisited = 0
}


func (this *GeographicMapCellHistory) PaintNotVisitedRelativeToPoint(graphics IGraphics, tiledLayer IAllBinaryTiledLayer, point IGPoint) {
        var x int = point.GetX() - tiledLayer.GetXP(); _ = x
        var y int = point.GetY() - tiledLayer.GetYP(); _ = y
        if this.HalfWidth == 0 {
            this.HalfWidth = (graphics.GetFont().StringWidth(this.MISSED_INFO) >> 1)
        }
        var height int = 2 * this.FontHeight; _ = height
        this.Animation.PaintXY(graphics, x + this.HalfWidth, y + (height))
}


func (this *GeographicMapCellHistory) PaintNotVisited(graphics IGraphics, geographicMapInterface IBasicGeographicMap) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "paintNotVisited", e)
                }
            }()
            this.MyFontProcessor.Process(graphics)
            graphics.SetColor(this.RED)
            var localVisitedList IBasicArrayList = this.VisitedList; _ = localVisitedList
            var size int = localVisitedList.Size(); _ = size
            var geographicMapCellPosition IGeographicMapCellPosition; _ = geographicMapCellPosition
            var isCellVisitedBoolean bool; _ = isCellVisitedBoolean
            for index := int(size); Java_PreDecrement(&index) >= 0;  {
                geographicMapCellPosition = any(this.List.Get(index)).(IGeographicMapCellPosition)
                isCellVisitedBoolean = any(localVisitedList.Get(index)).(bool)
                if !isCellVisitedBoolean {
                    this.PaintNotVisitedRelativeToPoint(graphics, geographicMapInterface.GetAllBinaryTiledLayer(), geographicMapCellPosition.GetPoint())
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GeographicMapCellHistory) Init() {
        this.List.Clear()
        this.VisitedList.Clear()
        this.TotalVisited = 0
}

