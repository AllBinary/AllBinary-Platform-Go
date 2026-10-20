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

type AllBinaryLayer struct {
    Layer
    ViewPositionEvent *ViewPositionEvent
    HalfWidth int
    HalfHeight int
    Name string
    ViewPosition IViewPositionBase
}

// Implements LayerInterface

// The Java class AllBinaryLayer has subclasses so its variables are this interface
type IAllBinaryLayer interface {
    ILayer
    AllBinaryLayer_() *AllBinaryLayer
    OnChangeEvent(layerManagerEvent *ViewPositionEvent)
    PaintThreed(graphics IGraphics)
    GetX2() int
    GetY2() int
    GetZ2() int
    GetName() string
    SetHalfWidth(halfWidth int)
    GetHalfWidth() int
    SetHalfHeight(halfHeight int)
    GetHalfHeight() int
    GetHalfDepth() int
    GetViewPosition() IViewPositionBase
    SetViewPosition(viewPosition IViewPositionBase)
    GetGroupInterface() []*Group
    ImplmentsTickableInterface() bool
    ImplmentsGameInputInterface() bool
    ImplmentsArtificialIntelligenceCompositeInterface() bool
    ImplmentsCollidableInterface() bool
    ImplementsThreedInterface() bool
    GetType() int
    GetMultiPlayerType() int
    ToStringAppend(stringBuffer *StringMaker)
    ToString() string
}

func (this *AllBinaryLayer) AllBinaryLayer_() *AllBinaryLayer {
    return this
}

var AllBinaryLayer_NULL_ALLBINARY_LAYER IAllBinaryLayer = NewAllBinaryLayer(StringUtil_GetInstance().EMPTY_STRING, RectangleFactory_SINGLETON, ViewPositionBase_NULL_VIEW_POSITION)
var AllBinaryLayer_TYPE string = "Type: "

func (this *AllBinaryLayer) initFields() {
    this.ViewPositionEvent = NewViewPositionEvent(this)
    
    
    
    this.ViewPosition = ViewPositionBase_NULL_VIEW_POSITION
}


func NewAllBinaryLayer(name string, rectangle *Rectangle, viewPosition IViewPositionBase) *AllBinaryLayer {
    this := &AllBinaryLayer{}
        this.Layer = *NewLayer(rectangle.GetWidth(), rectangle.GetHeight())
        this.initFields()
        var localName string; _ = localName
        if name == "" {
            localName = reflect.TypeOf(this).String()
        } else {
            localName = fmt.Sprint(NewStringMaker().Append(name).Append(CommonSeps_GetInstance().COLON).Append(strconv.FormatInt(int64(TsUtil_GetInstance().HashCode(this)), 16)))
        }
        this.Name = localName
        var point IGPoint = rectangle.GetPoint(); _ = point
        this.SetPosition(point.GetX(), point.GetY(), point.GetZ())
        this.SetHalfWidth((this.GetWidth() >> 1))
        this.SetHalfHeight((this.GetHeight() >> 1))
        this.ViewPosition = viewPosition
        this.ViewPosition.SetAllbinaryLayer(this)
    return this
}


// throws Exception
func (this *AllBinaryLayer) OnChangeEvent(layerManagerEvent *ViewPositionEvent) {
}


func (this *AllBinaryLayer) PaintThreed(graphics IGraphics) {
}


func (this *AllBinaryLayer) GetX2() int {
        return this.GetXP() + this.GetWidth()
}


func (this *AllBinaryLayer) GetY2() int {
        return this.GetYP() + this.GetHeight()
}


func (this *AllBinaryLayer) GetZ2() int {
        return int((int64(this.GetZP()) + this.GetDepth()))
}


func (this *AllBinaryLayer) GetName() string {
        return this.Name
}


func (this *AllBinaryLayer) SetHalfWidth(halfWidth int) {
        this.HalfWidth = halfWidth
}


func (this *AllBinaryLayer) GetHalfWidth() int {
        return this.HalfWidth
}


func (this *AllBinaryLayer) SetHalfHeight(halfHeight int) {
        this.HalfHeight = halfHeight
}


func (this *AllBinaryLayer) GetHalfHeight() int {
        return this.HalfHeight
}


func (this *AllBinaryLayer) GetHalfDepth() int {
        return int(this.GetDepth()) / 2
}


func (this *AllBinaryLayer) GetViewPosition() IViewPositionBase {
        return this.ViewPosition
}


func (this *AllBinaryLayer) SetViewPosition(viewPosition IViewPositionBase) {
        this.ViewPosition = viewPosition
}


func (this *AllBinaryLayer) GetGroupInterface() []*Group {
        return BasicGroupFactory_GetInstance().NONE_ARRAY
}


func (this *AllBinaryLayer) ImplmentsTickableInterface() bool {
        return false
}


func (this *AllBinaryLayer) ImplmentsGameInputInterface() bool {
        return false
}


func (this *AllBinaryLayer) ImplmentsArtificialIntelligenceCompositeInterface() bool {
        return false
}


func (this *AllBinaryLayer) ImplmentsCollidableInterface() bool {
        return false
}


func (this *AllBinaryLayer) ImplementsThreedInterface() bool {
        return OpenGLFeatureUtil_GetInstance().IsAnyThreed()
}


func (this *AllBinaryLayer) GetType() int {
        return -1
}


func (this *AllBinaryLayer) GetMultiPlayerType() int {
        return -1
}


func (this *AllBinaryLayer) ToStringAppend(stringBuffer *StringMaker) {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var positionStrings *PositionStrings = PositionStrings_GetInstance(); _ = positionStrings
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        stringBuffer.Append(this.GetName())
        stringBuffer.Append(commonSeps.COLON)
        stringBuffer.Append(AllBinaryLayer_TYPE)
        stringBuffer.Appendint(this.GetType())
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(positionStrings.X_LABEL)
        stringBuffer.Appendint(this.GetXP())
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(positionStrings.Y_LABEL)
        stringBuffer.Appendint(this.GetYP())
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(positionStrings.Z_LABEL)
        stringBuffer.Appendint(this.GetZP())
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(commonLabels.WIDTH_LABEL)
        stringBuffer.Appendint(this.GetWidth())
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(commonLabels.HEIGHT_LABEL)
        stringBuffer.Appendint(this.GetHeight())
}


func (this *AllBinaryLayer) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        this.ToStringAppend(stringBuffer)
        return fmt.Sprint(stringBuffer)
}

func (this *AllBinaryLayer) String() string {
    return this.ToString()
}

