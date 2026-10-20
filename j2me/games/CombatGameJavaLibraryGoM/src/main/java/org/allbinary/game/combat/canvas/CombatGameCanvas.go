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

type CombatGameCanvas struct {
    AllBinaryGameCanvas
}

// The Java class CombatGameCanvas has subclasses so its variables are this interface
type ICombatGameCanvas interface {
    IAllBinaryGameCanvas
    CombatGameCanvas_() *CombatGameCanvas
    CombatGameCanvas_init()
    InitApp(abeClientInformation AbeClientInformationInterface)
    InitConfigurablePortion(portion int)
    ProcessPlayingGame()
    CleanupGame()
    CleanupManager()
}

func (this *CombatGameCanvas) CombatGameCanvas_() *CombatGameCanvas {
    return this
}

var CombatGameCanvas_BasicLayerProcessor []IBasicLayerProcessor = make([]IBasicLayerProcessor, 0)


func NewCombatGameCanvas(cmdListener CommandListener, gameLayerManager IAllBinaryGameLayerManager, highScoresFactoryInterface HighScoresFactoryInterface, gameInitializationInterfaceFactoryInterface *BasicBuildGameInitializerFactory, buffered bool) *CombatGameCanvas {
    this := &CombatGameCanvas{}
        this.AllBinaryGameCanvas = *NewAllBinaryGameCanvas(cmdListener, gameLayerManager, highScoresFactoryInterface, gameInitializationInterfaceFactoryInterface, buffered)
    return this
}


func (this *CombatGameCanvas) CombatGameCanvas_init() {
        DestroyEventCircularStaticPool_GetInstance().Init(this)
        ProgressCanvasFactory_GetInstance().AddNormalPortion(50, "Destroy Events")
}


// throws Exception
func (this *CombatGameCanvas) InitApp(abeClientInformation AbeClientInformationInterface) {
        this.CombatGameCanvas_init()
        this.AllBinaryGameCanvas.InitApp(abeClientInformation)
}


func (this *CombatGameCanvas) InitConfigurablePortion(portion int) {
        DestroyedLayerProcessor_Init()
        ProgressCanvasFactory_GetInstance().AddNormalPortion(portion, "Basic Processors")
        var features *Features = Features_GetInstance(); _ = features
        var gameFeatureFactory *GameFeatureFactory = GameFeatureFactory_GetInstance(); _ = gameFeatureFactory
        if features.IsFeature(gameFeatureFactory.DROPPED_ITEMS) && features.IsFeature(gameFeatureFactory.DROPPED_ITEMS_FROM_DEATH) {
            CombatGameCanvas_BasicLayerProcessor = make([]IBasicLayerProcessor, 2)
            CombatGameCanvas_BasicLayerProcessor[0] = DestroyedLayerProcessor_GetInstance()
            CombatGameCanvas_BasicLayerProcessor[1] = DropLayerProcessor_GetInstance()
        } else {
            CombatGameCanvas_BasicLayerProcessor = make([]IBasicLayerProcessor, 1)
            CombatGameCanvas_BasicLayerProcessor[0] = DestroyedLayerProcessor_GetInstance()
        }
}


// throws Exception
func (this *CombatGameCanvas) ProcessPlayingGame() {
        this.AllBinaryGameCanvas.ProcessPlayingGame()
        for index := int(len(CombatGameCanvas_BasicLayerProcessor)); Java_PreDecrement(&index) >= 0;  {
            CombatGameCanvas_BasicLayerProcessor[index].Process(this.GameLayerManager)
        }
}


// throws Exception
func (this *CombatGameCanvas) CleanupGame() {
        this.AllBinaryGameCanvas.CleanupGame()
        for index := int(len(CombatGameCanvas_BasicLayerProcessor)); Java_PreDecrement(&index) >= 0;  {
            CombatGameCanvas_BasicLayerProcessor[index].GetList().Clear()
        }
        GroupLayerManagerListener_GetInstance().Clear()
        GroupLayerManagerListener_GetInstance().Log()
        DestroyedEventHandler_GetInstance().RemoveAllListeners()
        this.CleanupManager()
}


// throws Exception
func (this *CombatGameCanvas) CleanupManager() {
        this.GameLayerManager.Cleanup()
}

