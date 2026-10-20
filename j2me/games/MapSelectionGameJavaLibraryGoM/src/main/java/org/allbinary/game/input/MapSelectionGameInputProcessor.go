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

type MapSelectionGameInputProcessor struct {
    Processor
    LogUtil *LogUtil
    GameCanvas IAllBinaryGameCanvas
    RaceTrackGeographicMapInterfaceFactoryInterface MultiLevelRaceTrackGeographicMapInterfaceFactoryInterface
    PreGameSelectorPaintable *PreGameSelectorPaintable
    InputTimeHelper ITimeDelayHelper
    PlayerGameInput IPlayerGameInput
    LockedIndex int
    AbRunnable IABRunnable
}

// Implements GameInputProcessorInterface, PlayerGameInputCompositeInterface


func (this *MapSelectionGameInputProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    this.InputTimeHelper = NewTimeDelayHelper(650)
    
    
    this.AbRunnable = /* Java anonymous class body: implement it as a named type
        

    func (this *MapSelectionGameInputProcessor) Run() {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.SetRunning(false)
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                    }
                }()
                this.SetRunning(true)
                SecondaryPlayerQueueFactory_GetInstance().Add(SelectSound_GetInstance())
                var track int = this.PreGameSelectorPaintable.GetPreGameSelectionForm().GetSelectedIndex() + 1; _ = track
                var wave int = this.RaceTrackGeographicMapInterfaceFactoryInterface.GetFirstWaveWithTrack(track); _ = wave
                this.GameCanvas.GetLayerManager().GetGameInfo().SetCurrentLevel(wave)
                this.GameCanvas.BuildGameInit(false)
                GameKeyEventHandler_GetInstance().RemoveListener(this.GetPlayerGameInput())
                this.SetRunning(false)
                return
            }(); tryControl1 == 1 {
                return
            }
    }
*/ NewABRunnable()
}


func NewMapSelectionGameInputProcessor(gameCanvas IAllBinaryGameCanvas, raceTrackGeographicMapInterfaceFactoryInterface MultiLevelRaceTrackGeographicMapInterfaceFactoryInterface, mapSelectorPaintable *PreGameSelectorPaintable, lockedIndex int) *MapSelectionGameInputProcessor {
    this := &MapSelectionGameInputProcessor{}
    this.Processor = *NewProcessor()
    this.initFields()
        this.GameCanvas = gameCanvas
        this.RaceTrackGeographicMapInterfaceFactoryInterface = raceTrackGeographicMapInterfaceFactoryInterface
        this.PreGameSelectorPaintable = mapSelectorPaintable
        this.PlayerGameInput = NewGameInputProcessorComposite("Map Selection", this)
        this.LockedIndex = lockedIndex
    return this
}


// throws Exception
func (this *MapSelectionGameInputProcessor) Process() {
        this.GetPlayerGameInput().Update()
}


// throws Exception
func (this *MapSelectionGameInputProcessor) OnInput(list IBasicArrayList) {
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var anyType any = list.Get(index); _ = anyType
            var key int = GameKeyEventUtil_GetKey(anyType); _ = key
            if key == Canvas_LEFT || key == Canvas_RIGHT || key == Canvas_UP || key == Canvas_DOWN {
                if this.InputTimeHelper.IsTimeTNT() {
                    SecondaryPlayerQueueFactory_GetInstance().Add(SelectSound_GetInstance())
                    this.PreGameSelectorPaintable.GetPreGameSelectionForm().ProcessInputKey(key)
                    break
                }
            } else if key == Canvas_KEY_NUM0 {
                var selectedIndex int = this.PreGameSelectorPaintable.GetPreGameSelectionForm().GetSelectedIndex(); _ = selectedIndex
                if selectedIndex < this.LockedIndex || !LockedUtil_GetInstance().IsLocked() {
                    if !this.AbRunnable.IsRunning() {
                        this.AbRunnable.SetRunning(true)
                        ThreadObjectUtil_GetInstance().ProcessThread(this.AbRunnable)
                    }
                    break
                } else {
                    LockedFeatureNotificationUtil_GetInstance().Fire()
                }
            }
            list.Clear()
        }
}


// throws Exception
func (this *MapSelectionGameInputProcessor) ProcessInput(layerManager IAllBinaryLayerManager) {
}


func (this *MapSelectionGameInputProcessor) InitInputProcessors() {
}


func (this *MapSelectionGameInputProcessor) GetName() string {
        return this.ToString()
}


func (this *MapSelectionGameInputProcessor) GetPlayerGameInput() IPlayerGameInput {
        return this.PlayerGameInput
}

