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

type AllBinaryGameLayerManager struct {
    AllBinaryLayerManager
    BackgroundBasicColor *BasicColor
    ForegroundBasicColor *BasicColor
    GameInfo *GameInfo
    PlayerGameInputCompositeInterface PlayerGameInputCompositeInterface
}

// The Java class AllBinaryGameLayerManager has subclasses so its variables are this interface
type IAllBinaryGameLayerManager interface {
    IAllBinaryLayerManager
    AllBinaryGameLayerManager_() *AllBinaryGameLayerManager
    GetLayerManager() IAllBinaryLayerManager
    Init()
    GetBackgroundBasicColor() *BasicColor
    SetBackgroundBasicColor(backgroundBasicColor *BasicColor)
    GetForegroundBasicColor() *BasicColor
    SetForegroundBasicColor(foregroundBasicColor *BasicColor)
    GetGameInfo() *GameInfo
    SetGameInfo(gameInfo *GameInfo)
    SetLayerProcessorList(list IBasicArrayList)
    SetPlayerGameInputCompositeInterface(playerGameInputCompositeInterface PlayerGameInputCompositeInterface)
    GetPlayerGameInputCompositeInterface() PlayerGameInputCompositeInterface
}

func (this *AllBinaryGameLayerManager) AllBinaryGameLayerManager_() *AllBinaryGameLayerManager {
    return this
}

var AllBinaryGameLayerManager_NULL_ALLBINARY_LAYER_MANAGER any = NullUtil_GetInstance().NULL_OBJECT
var AllBinaryGameLayerManager_ID int = SmallIntegerSingletonFactory_GetInstance().GetAt(22)

func (this *AllBinaryGameLayerManager) initFields() {
    
    
    
    this.PlayerGameInputCompositeInterface = NullPlayerGameInputComposite_NULL_PLAYER_GAME_INPUT_COMPOSITE
}


func AllBinaryGameLayerManager_GetNullInstance() IAllBinaryGameLayerManager {
        if AllBinaryGameLayerManager_NULL_ALLBINARY_LAYER_MANAGER == NullUtil_GetInstance().NULL_OBJECT {
            AllBinaryGameLayerManager_NULL_ALLBINARY_LAYER_MANAGER = NewAllBinaryGameLayerManager(BasicColorFactory_GetInstance().BLACK, BasicColorFactory_GetInstance().WHITE, GameInfo_NONE)
        }
        return any(AllBinaryGameLayerManager_NULL_ALLBINARY_LAYER_MANAGER).(IAllBinaryGameLayerManager)
}


func NewAllBinaryGameLayerManager(backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor, gameInfo *GameInfo) *AllBinaryGameLayerManager {
    this := &AllBinaryGameLayerManager{}
    this.AllBinaryLayerManager = *NewAllBinaryLayerManager()
    this.initFields()
        this.BackgroundBasicColor = backgroundBasicColor
        this.ForegroundBasicColor = foregroundBasicColor
        var itemColorFactory *ItemColorFactory = ItemColorFactory_GetInstance(); _ = itemColorFactory
        itemColorFactory.INVERT_PAINT = this.BackgroundBasicColor.IntValue()
        itemColorFactory.PAINT = this.ForegroundBasicColor.IntValue()
        this.GameInfo = gameInfo
    return this
}


func (this *AllBinaryGameLayerManager) GetLayerManager() IAllBinaryLayerManager {
        return this
}


func (this *AllBinaryGameLayerManager) Init() {
}


func (this *AllBinaryGameLayerManager) GetBackgroundBasicColor() *BasicColor {
        return this.BackgroundBasicColor
}


func (this *AllBinaryGameLayerManager) SetBackgroundBasicColor(backgroundBasicColor *BasicColor) {
        this.BackgroundBasicColor = backgroundBasicColor
        ItemColorFactory_GetInstance().INVERT_PAINT = this.BackgroundBasicColor.IntValue()
}


func (this *AllBinaryGameLayerManager) GetForegroundBasicColor() *BasicColor {
        return this.ForegroundBasicColor
}


func (this *AllBinaryGameLayerManager) SetForegroundBasicColor(foregroundBasicColor *BasicColor) {
        this.ForegroundBasicColor = foregroundBasicColor
        ItemColorFactory_GetInstance().PAINT = this.ForegroundBasicColor.IntValue()
}


func (this *AllBinaryGameLayerManager) GetGameInfo() *GameInfo {
        return this.GameInfo
}


func (this *AllBinaryGameLayerManager) SetGameInfo(gameInfo *GameInfo) {
        this.GameInfo = gameInfo
}


func (this *AllBinaryGameLayerManager) SetLayerProcessorList(list IBasicArrayList) {
        var layerProcessorInterfaceArray []ILayerProcessor = make([]ILayerProcessor, list.Size()); _ = layerProcessorInterfaceArray
        var size int = len(layerProcessorInterfaceArray); _ = size
        for index := int(0); index < size; index++ {
            layerProcessorInterfaceArray[index] = any(list.BasicArrayList_().ObjectArray[index]).(ILayerProcessor)
        }
        this.SetLayerProcessorArray(layerProcessorInterfaceArray)
}


func (this *AllBinaryGameLayerManager) SetPlayerGameInputCompositeInterface(playerGameInputCompositeInterface PlayerGameInputCompositeInterface) {
        this.PlayerGameInputCompositeInterface = playerGameInputCompositeInterface
}


func (this *AllBinaryGameLayerManager) GetPlayerGameInputCompositeInterface() PlayerGameInputCompositeInterface {
        return this.PlayerGameInputCompositeInterface
}

