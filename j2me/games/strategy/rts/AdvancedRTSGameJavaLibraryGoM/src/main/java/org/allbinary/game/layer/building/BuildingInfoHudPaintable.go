/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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
    . "org/allbinary/game/ag"
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

type BuildingInfoHudPaintable struct {
    SelectionHudPaintable
    LogUtil *LogUtil
    HEALTH string
    ProductivityHud INumberStringHud
    EfficiencyHud INumberStringHud
    HealthHud INumberStringHud
    MaxHealthHud INumberStringHud
    RtsLayer PathFindingLayerInterface
    CurrentHealth int
}

var BuildingInfoHudPaintable_Instance *BuildingInfoHudPaintable = NewBuildingInfoHudPaintable()

func (this *BuildingInfoHudPaintable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.HEALTH = "Health:"
    
    
    
    
    this.RtsLayer = NullPathFindingLayer_NULL_PATH_FINDING_LAYER
    this.CurrentHealth = 0
}


func BuildingInfoHudPaintable_GetInstance() *BuildingInfoHudPaintable {
        return BuildingInfoHudPaintable_Instance
}


func NewBuildingInfoHudPaintable() *BuildingInfoHudPaintable {
    this := &BuildingInfoHudPaintable{}
    this.SelectionHudPaintable = *NewSelectionHudPaintable()
    this.initFields()
        var numberStringHudFactory *NumberStringHudFactory = NumberStringHudFactory_GetInstance(); _ = numberStringHudFactory
        var productivityHud INumberStringHud = numberStringHudFactory.NULL_NUMBER_STRING_HUD; _ = productivityHud
        var efficiencyHud INumberStringHud = numberStringHudFactory.NULL_NUMBER_STRING_HUD; _ = efficiencyHud
        var healthHud INumberStringHud = numberStringHudFactory.NULL_NUMBER_STRING_HUD; _ = healthHud
        var maxHealthHud INumberStringHud = numberStringHudFactory.NULL_NUMBER_STRING_HUD; _ = maxHealthHud
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
                }
            }()
            var index int = 0; _ = index
            var basicHudFactory *BasicHudFactory = BasicHudFactory_GetInstance(); _ = basicHudFactory
            var basicColor *BasicColor = this.GetBasicColorP(); _ = basicColor
            var textX int = this.TextX; _ = textX
            var y int = this.Y; _ = y
            var firstIndex int = index; _ = firstIndex
            /* Java local class: move it out of the function
type ProductivityNumberStringHud struct {
    NumberStringHud
}

            

            func NewBuildingInfoHudPaintable() *BuildingInfoHudPaintable {
                this := &BuildingInfoHudPaintable{}
                    this.SelectionHudPaintable = *NewSelectionHudPaintable("Productivity:", 999, this.BasicHudFactory.ABSOLUTE, this.BasicHudFactory.HORIZONTAL, 0, this.BasicColor)
                return this
            }
            

            func (this *BuildingInfoHudPaintable) UpdateMeasurement(graphics IGraphics) {
                    var font Font = graphics.GetFont(); _ = font
                    this.UpdateMaxWidth = this.TextX
                    this.UpdateMaxHeight = this.Y + ((this.FirstIndex + 1) * font.GetHeight())
                    this.SelectionHudPaintable.UpdateMeasurement(graphics)
            }
*/
            productivityHud = NewProductivityNumberStringHud()
            index++
            var secondIndex int = index; _ = secondIndex
            /* Java local class: move it out of the function
type EfficiencyNumberStringHud struct {
    NumberStringHud
}

            

            func NewBuildingInfoHudPaintable() *BuildingInfoHudPaintable {
                this := &BuildingInfoHudPaintable{}
                    this.SelectionHudPaintable = *NewSelectionHudPaintable("Efficiency:", 999, this.BasicHudFactory.ABSOLUTE, this.BasicHudFactory.HORIZONTAL, 0, this.BasicColor)
                return this
            }
            

            func (this *BuildingInfoHudPaintable) UpdateMeasurement(graphics IGraphics) {
                    this.SelectionHudPaintable.UpdateMeasurement(graphics)
                    var font Font = graphics.GetFont(); _ = font
                    this.UpdateMaxWidth = this.TextX
                    this.UpdateMaxHeight = this.Y + ((this.SecondIndex + 1) * font.GetHeight())
            }
*/
            efficiencyHud = NewEfficiencyNumberStringHud()
            index++
            var HEALTH string = this.HEALTH; _ = HEALTH
            var thirdIndex int = index; _ = thirdIndex
            /* Java local class: move it out of the function
type HealthNumberStringHud struct {
    NumberStringHud
}

            

            func NewBuildingInfoHudPaintable() *BuildingInfoHudPaintable {
                this := &BuildingInfoHudPaintable{}
                    this.SelectionHudPaintable = *NewSelectionHudPaintable(this.HEALTH, 99999, this.BasicHudFactory.ABSOLUTE, this.BasicHudFactory.HORIZONTAL, 0, this.BasicColor)
                return this
            }
            

            func (this *BuildingInfoHudPaintable) UpdateMeasurement(graphics IGraphics) {
                    var font Font = graphics.GetFont(); _ = font
                    this.UpdateMaxWidth = this.TextX
                    this.UpdateMaxHeight = this.Y + ((this.ThirdIndex + 1) * font.GetHeight())
                    this.SelectionHudPaintable.UpdateMeasurement(graphics)
            }
*/
            var totalLength int = len(HEALTH) + 1; _ = totalLength
            healthHud = NewHealthNumberStringHud()
            var fourthIndex int = index; _ = fourthIndex
            /* Java local class: move it out of the function
type MaxHealthNumberStringHud struct {
    NumberStringHud
}

            

            func NewBuildingInfoHudPaintable() *BuildingInfoHudPaintable {
                this := &BuildingInfoHudPaintable{}
                    this.SelectionHudPaintable = *NewSelectionHudPaintable("/ ", 99999, this.BasicHudFactory.ABSOLUTE, this.BasicHudFactory.HORIZONTAL, 0, this.BasicColor)
                return this
            }
            

            func (this *BuildingInfoHudPaintable) UpdateMeasurement(graphics IGraphics) {
                    this.SelectionHudPaintable.UpdateMeasurement(graphics)
                    var font Font = graphics.GetFont(); _ = font
                    this.UpdateMaxWidth = this.TextX + (this.TotalLength * font.GetHeight())
                    this.UpdateMaxHeight = this.Y + ((this.FourthIndex + 1) * font.GetHeight())
            }
*/
            maxHealthHud = NewMaxHealthNumberStringHud()
            return
        }(); tryControl1 == 1 {
            return
        }
        this.ProductivityHud = productivityHud
        this.EfficiencyHud = efficiencyHud
        this.HealthHud = healthHud
        this.MaxHealthHud = maxHealthHud
    return this
}


func (this *BuildingInfoHudPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        var totalLength int = len(this.HEALTH) + MathUtil_GetInstance().GetTotalDigits(this.CurrentHealth); _ = totalLength
        this.MaxHealthHud.SetX(this.TextX + MyFontProcessor_DefaultStringWidth(font, totalLength))
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *BuildingInfoHudPaintable) SetBasicColorP(basicColor *BasicColor) {
        this.SelectionHudPaintable.SetBasicColorP(basicColor)
        this.ProductivityHud.SetBasicColorP(basicColor)
        this.EfficiencyHud.SetBasicColorP(basicColor)
        this.HealthHud.SetBasicColorP(basicColor)
        this.MaxHealthHud.SetBasicColorP(basicColor)
}


func (this *BuildingInfoHudPaintable) Paint(graphics IGraphics) {
        this.SelectionHudPaintable.Paint(graphics)
        this.MyFontProcessor.Process(graphics)
        this.ProductivityHud.Paint(graphics)
        this.EfficiencyHud.Paint(graphics)
        this.HealthHud.Paint(graphics)
        this.MaxHealthHud.Paint(graphics)
        this.GetAnimationInterface().PaintXY(graphics, this.ImageX, this.Y)
}


func (this *BuildingInfoHudPaintable) UpdateSelectionInfo() {
        var buildingLayer *BuildingLayer = any(this.RtsLayer).(*BuildingLayer); _ = buildingLayer
        this.SetName(buildingLayer.GetName())
        this.SetAnimationInterface(buildingLayer.GetVerticleBuildAnimationInterface())
        this.ProductivityHud.Set(buildingLayer.GetProductivity())
        this.EfficiencyHud.Set(buildingLayer.GetEfficiency() / 100)
        this.CurrentHealth = buildingLayer.GetHealthInterface().GetHealth()
        this.HealthHud.Set(this.CurrentHealth)
        this.MyFontProcessor = this.UpdateMyFontProcessor
        this.MaxHealthHud.Set(buildingLayer.GetHealthInterface().GetMaxHealth())
}


func (this *BuildingInfoHudPaintable) SetRtsLayer(rtsLayer PathFindingLayerInterface) {
        this.RtsLayer = rtsLayer
}

