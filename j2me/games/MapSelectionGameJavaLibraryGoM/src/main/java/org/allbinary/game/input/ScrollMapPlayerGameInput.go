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

type ScrollMapPlayerGameInput struct {
    PlayerGameInput
    InputProcessorArray []IGameInputProcessor
    Border int
    InputList IBasicArrayList
    IsSingleKeyProcessing bool
    GeographicMapInterface IBasicGeographicMap
    MaxBottom int
    MaxRight int
    ScrollMapEvent *ScrollMapEvent
}


func (this *ScrollMapPlayerGameInput) initFields() {
    this.InputProcessorArray = make([]IGameInputProcessor, InputFactory_GetInstance().MAX)
    this.Border = 150
    
    this.IsSingleKeyProcessing = Features_GetInstance().IsFeature(InputFeatureFactory_GetInstance().SINGLE_KEY_REPEAT_PRESS) || Features_GetInstance().IsFeature(InputFeatureFactory_GetInstance().SINGLE_KEY_PRESS)
    
    
    
    this.ScrollMapEvent = NewScrollMapEvent(this)
}


func NewScrollMapPlayerGameInput(geographicMapInterface IBasicGeographicMap, list IBasicArrayList, playerInputId int) *ScrollMapPlayerGameInput {
    this := &ScrollMapPlayerGameInput{}
        this.PlayerGameInput = *NewPlayerGameInput(list, NewBasicArrayListD(), playerInputId)
        this.initFields()
        this.InitInputProcessors()
        this.GeographicMapInterface = geographicMapInterface
        this.InputList = list
        this.Init()
    return this
}


// throws Exception
func (this *ScrollMapPlayerGameInput) Move(dx int, dy int) {
        var terrainTiledLayer IAllBinaryTiledLayer = this.GeographicMapInterface.GetAllBinaryTiledLayer(); _ = terrainTiledLayer
        terrainTiledLayer.MoveDXY(-dx, -dy)
        this.ScrollMapEvent.SetDxDy(-dx, -dy)
        ScrollMapEventHandler_GetInstance().FireEvent(this.ScrollMapEvent)
}


func (this *ScrollMapPlayerGameInput) Init() {
        this.MaxBottom = this.GetSpecialHeight()
        this.MaxRight = this.GetSpecialWidth()
}


// throws Exception
func (this *ScrollMapPlayerGameInput) ScrollMiddleX() {
        var terrainTiledLayer IAllBinaryTiledLayer = this.GeographicMapInterface.GetAllBinaryTiledLayer(); _ = terrainTiledLayer
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        this.Move(-terrainTiledLayer.GetXP(), 0)
        if terrainTiledLayer.GetWidth() < displayInfo.GetLastWidth() {
            var diffX int = ((displayInfo.GetLastWidth() - terrainTiledLayer.GetWidth()) >> 1); _ = diffX
            this.Move(-diffX, 0)
        } else {
            var diffX int = ((terrainTiledLayer.GetWidth() - displayInfo.GetLastWidth()) >> 1); _ = diffX
            this.Move(diffX, 0)
        }
}


// throws Exception
func (this *ScrollMapPlayerGameInput) ScrollY(anchor int) {
        var terrainTiledLayer IAllBinaryTiledLayer = this.GeographicMapInterface.GetAllBinaryTiledLayer(); _ = terrainTiledLayer
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        if anchor == Graphics_TOP {
            this.LogUtil.PutF("Top", this, "scrollY")
        } else if anchor == Graphics_BOTTOM {
            this.LogUtil.PutF("Bottom", this, "scrollY")
            var diffY int = (terrainTiledLayer.GetHeight() - displayInfo.GetLastHeight()); _ = diffY
            this.Move(0, diffY)
        } else if anchor == Graphics_VCENTER {
            this.LogUtil.PutF("Center", this, "scrollY")
            var diffY int = (terrainTiledLayer.GetHeight() - displayInfo.GetLastHeight()) / 2; _ = diffY
            this.Move(0, diffY)
        } else {
            panic(errors.New(fmt.Sprint("No Such Anchor Supported")))
        }
}


// throws Exception
func (this *ScrollMapPlayerGameInput) Up() {
        var y int = this.GeographicMapInterface.GetAllBinaryTiledLayer().GetYP() - 10; _ = y
        if this.KeepOnMapMinY(y) {
            this.Move(0, -10)
        }
}


// throws Exception
func (this *ScrollMapPlayerGameInput) Down() {
        var y int = -this.GeographicMapInterface.GetAllBinaryTiledLayer().GetYP() + 10; _ = y
        if this.KeepOnMapMaxY(y) {
            this.Move(0, 10)
        }
}


// throws Exception
func (this *ScrollMapPlayerGameInput) Right() {
        var x int = this.GeographicMapInterface.GetAllBinaryTiledLayer().GetYP() + 10; _ = x
        if this.KeepOnMapMinX(x) {
            this.Move(10, 0)
        }
}


// throws Exception
func (this *ScrollMapPlayerGameInput) Left() {
        var x int = -this.GeographicMapInterface.GetAllBinaryTiledLayer().GetYP() - 10; _ = x
        if this.KeepOnMapMaxX(x) {
            this.Move(-10, 0)
        }
}


func (this *ScrollMapPlayerGameInput) InitInputProcessors() {
        this.InputProcessorArray[Canvas_UP] = NewScrollMapUpGameInputProcessor(this)
        this.InputProcessorArray[Canvas_DOWN] = NewScrollMapDownGameInputProcessor(this)
        this.InputProcessorArray[Canvas_KEY_NUM9] = NewScrollMapRightGameInputProcessor(this)
        this.InputProcessorArray[Canvas_KEY_NUM7] = NewScrollMapLeftGameInputProcessor(this)
        GameInputProcessorUtil_Init(this.InputProcessorArray)
}


// throws Exception
func (this *ScrollMapPlayerGameInput) ProcessInputKey(key int) {
        this.InputProcessorArray[key].ProcessEvent(AllBinaryGameLayerManager_GetNullInstance(), GameKeyEvent_NONE)
}


// throws Exception
func (this *ScrollMapPlayerGameInput) ProcessInput(layerManager IAllBinaryLayerManager) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.PROCESS_INPUT, e)
                }
            }()
            var size int = this.InputList.Size(); _ = size
            var key int = 0; _ = key
            for index := int(0); index < size; index++ {
                var gameKeyEvent *GameKeyEvent = any(this.InputList.Get(index)).(*GameKeyEvent); _ = gameKeyEvent
                key = gameKeyEvent.GetKey()
                this.ProcessInputKey(key)
            }
            if this.IsSingleKeyProcessing {
                this.Clear()
            } else {
                this.Update()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ScrollMapPlayerGameInput) GetSpecialHeight() int {
        var terrainTiledLayer IAllBinaryTiledLayer = this.GeographicMapInterface.GetAllBinaryTiledLayer(); _ = terrainTiledLayer
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        if terrainTiledLayer.GetHeight() > displayInfo.GetLastHeight() {
            return terrainTiledLayer.GetHeight() - displayInfo.GetLastHeight()
        } else {
            return displayInfo.GetLastHeight() - terrainTiledLayer.GetHeight() + terrainTiledLayer.GetCellHeight()
        }
}


func (this *ScrollMapPlayerGameInput) GetSpecialWidth() int {
        var terrainTiledLayer IAllBinaryTiledLayer = this.GeographicMapInterface.GetAllBinaryTiledLayer(); _ = terrainTiledLayer
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        if terrainTiledLayer.GetWidth() > displayInfo.GetLastWidth() {
            return terrainTiledLayer.GetWidth() - displayInfo.GetLastWidth()
        } else {
            return displayInfo.GetLastWidth() - terrainTiledLayer.GetWidth() + terrainTiledLayer.GetCellWidth()
        }
}


func (this *ScrollMapPlayerGameInput) KeepOnMapMaxY(newY int) bool {
        var y int = this.GeographicMapInterface.GetAllBinaryTiledLayer().GetYP(); _ = y
        if y > this.MaxBottom + this.Border {
            return false
        } else {
            return true
        }
}


func (this *ScrollMapPlayerGameInput) KeepOnMapMinY(newY int) bool {
        var y int = this.GeographicMapInterface.GetAllBinaryTiledLayer().GetYP(); _ = y
        if y < -this.Border {
            return false
        } else {
            return true
        }
}


func (this *ScrollMapPlayerGameInput) KeepOnMapMaxX(newX int) bool {
        var x int = this.GeographicMapInterface.GetAllBinaryTiledLayer().GetXP(); _ = x
        if x < -this.MaxRight - this.Border {
            return false
        } else {
            return true
        }
}


func (this *ScrollMapPlayerGameInput) KeepOnMapMinX(newX int) bool {
        var x int = this.GeographicMapInterface.GetAllBinaryTiledLayer().GetXP(); _ = x
        if x > this.Border {
            return false
        } else {
            return true
        }
}

