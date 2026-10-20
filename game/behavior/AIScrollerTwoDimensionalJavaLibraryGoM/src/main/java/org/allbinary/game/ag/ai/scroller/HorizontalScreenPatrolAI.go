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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type HorizontalScreenPatrolAI struct {
    BasicAI
    KeyDirection int
    DisplayInfoSingleton *DisplayInfoSingleton
    FiringX int
    FiredIndex int
    MAX_FIRE int
    TimeDelayHelper ITimeDelayHelper
    SecondaryPlayerQueue *PlayerQueue
    Sound ISound
}

var HorizontalScreenPatrolAI_SOUND int = SmallIntegerSingletonFactory_GetInstance().GetAt(433)

func (this *HorizontalScreenPatrolAI) initFields() {
    this.KeyDirection = Canvas_LEFT
    this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
    this.FiringX = this.DisplayInfoSingleton.GetLastHalfWidth()
    this.FiredIndex = 0
    this.MAX_FIRE = 28
    this.TimeDelayHelper = NewTimeDelayHelper(240)
    this.SecondaryPlayerQueue = SecondaryPlayerQueueFactory_GetInstance()
    
}


func NewHorizontalScreenPatrolAI(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *HorizontalScreenPatrolAI {
    this := &HorizontalScreenPatrolAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        this.Sound = any(hashtable.Get(HorizontalScreenPatrolAI_SOUND)).(ISound)
        if this.Sound == nil {
            panic(errors.New(fmt.Sprint("No Sound Provided!!")))
        }
    return this
}


// throws Exception
func (this *HorizontalScreenPatrolAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        this.BasicAI.ProcessKeyAI(this.KeyDirection)
        if this.TimeDelayHelper.IsTimeTNT() {
            var layerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = layerInterface
            var currentOwnerLayerX int = layerInterface.GetXP(); _ = currentOwnerLayerX
            if currentOwnerLayerX < -layerInterface.GetWidth() {
                var y int = MyRandomFactory_GetInstance().GetAbsoluteNextIntAllowZero(50); _ = y
                layerInterface.SetPosition(this.DisplayInfoSingleton.GetLastWidth() + layerInterface.GetWidth() + 50, y, layerInterface.GetZP())
                this.FiredIndex = 0
                this.FiringX = this.DisplayInfoSingleton.GetLastHalfWidth() + MyRandomFactory_GetInstance().GetNextInt(this.DisplayInfoSingleton.GetLastHalfWidth())
                this.SecondaryPlayerQueue.Add(this.Sound)
            }
            if currentOwnerLayerX < this.FiringX && currentOwnerLayerX > -layerInterface.GetWidth() {
                this.BasicAI.ProcessKeyAI(Canvas_KEY_NUM1)
                var gameLayerManager IAllBinaryGameLayerManager = any(allBinaryLayerManager).(IAllBinaryGameLayerManager); _ = gameLayerManager
                var FIRE int = (gameLayerManager.GetGameInfo().GetCurrentLevel() + 12) >> 2; _ = FIRE
                if FIRE > this.MAX_FIRE {
                    FIRE = this.MAX_FIRE
                }
                if this.FiredIndex > FIRE {
                    this.FiredIndex = 0
                    this.FiringX = Integer_MIN_VALUE
                }
                this.FiredIndex++
            }
        }
}


func (this *HorizontalScreenPatrolAI) NextDirection() {
}

