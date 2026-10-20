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

type AllBinaryGameLayer struct {
    AllBinaryLayer
    CommonStrings *CommonStrings
    BasicSetColorUtil IBasicColorSetUtil
    RED *BasicColor
    Paintable IPaintable
    GameKeyEventList IBasicArrayList
}

// Implements ViewPositionEventListenerInterface, OpenGLSurfaceChangedInterface

// The Java class AllBinaryGameLayer has subclasses so its variables are this interface
type IAllBinaryGameLayer interface {
    IAllBinaryLayer
    AllBinaryGameLayer_() *AllBinaryGameLayer
    Set(gl GL)
    Move()
    GetGameKeyEventList() IBasicArrayList
    ProcessInput(allBinaryLayerManager IAllBinaryLayerManager)
    ProcessTick(allBinaryLayerManager IAllBinaryLayerManager)
    OnEvent(eventObject IAllBinaryEventObject)
    OnViewPositionChangeEvent()
    OnChangeEvent(layerManagerEvent *ViewPositionEvent)
    PaintFirst(graphics IGraphics)
    PaintDebug(graphics IGraphics)
    SetWidth(width int)
    SetHeight(height int)
}

func (this *AllBinaryGameLayer) AllBinaryGameLayer_() *AllBinaryGameLayer {
    return this
}


func (this *AllBinaryGameLayer) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.BasicSetColorUtil = BasicColorSetUtil_GetInstance()
    this.RED = BasicColorFactory_GetInstance().RED
    this.Paintable = func() any { if SWTUtil_IsSWT { return /* Java anonymous class body: implement it as a named type
        this.BasicSetColorUtil = BasicColorSetUtil_GetInstance()
        this.BLACK = BasicColorFactory_GetInstance().BLACK
        

    func (this *AllBinaryGameLayer) Paint(graphics IGraphics) {
            this.BasicSetColorUtil.SetBasicColorP(graphics, this.BLACK)
    }
*/ NewPaintable() }; return NullPaintable_GetInstance() }()
    this.GameKeyEventList = NewBasicArrayListD()
}


func NewAllBinaryGameLayer(name string, layerInfo *Rectangle, viewPosition IViewPositionBase) *AllBinaryGameLayer {
    this := &AllBinaryGameLayer{}
        this.AllBinaryLayer = *NewAllBinaryLayer(name, layerInfo, viewPosition)
        this.initFields()
    return this
}


// throws Exception
func (this *AllBinaryGameLayer) Set(gl GL) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *AllBinaryGameLayer) Move() {
}


func (this *AllBinaryGameLayer) GetGameKeyEventList() IBasicArrayList {
        return this.GameKeyEventList
}


// throws Exception
func (this *AllBinaryGameLayer) ProcessInput(allBinaryLayerManager IAllBinaryLayerManager) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *AllBinaryGameLayer) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *AllBinaryGameLayer) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
func (this *AllBinaryGameLayer) OnViewPositionChangeEvent() {
        this.OnChangeEvent(this.ViewPositionEvent)
}


// throws Exception
func (this *AllBinaryGameLayer) OnChangeEvent(layerManagerEvent *ViewPositionEvent) {
        if GameLayerUtil_IsOnScreen(this) {
            this.SetVisible(true)
        } else {
            this.SetVisible(false)
        }
}


func (this *AllBinaryGameLayer) PaintFirst(graphics IGraphics) {
        this.Paintable.Paint(graphics)
}


func (this *AllBinaryGameLayer) PaintDebug(graphics IGraphics) {
        var viewPosition IViewPositionBase = this.GetViewPosition(); _ = viewPosition
        var viewX int = viewPosition.GetX(); _ = viewX
        var viewY int = viewPosition.GetY(); _ = viewY
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.RED)
        graphics.DrawRect(viewX, viewY, this.GetWidth(), this.GetHeight())
}


func (this *AllBinaryGameLayer) SetWidth(width int) {
        this.AllBinaryLayer.SetLayerWidth(width)
        this.SetHalfWidth(width >> 1)
}


func (this *AllBinaryGameLayer) SetHeight(height int) {
        this.AllBinaryLayer.SetLayerHeight(height)
        this.SetHalfHeight(height >> 1)
}

