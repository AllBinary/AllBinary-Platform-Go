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

type GameKeyEventHandler struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    PressGameKeyEventHandler *PressGameKeyEventHandler
    UpGameKeyEventHandler *UpGameKeyEventHandler
    DownGameKeyEventHandler *DownGameKeyEventHandler
}

var GameKeyEventHandler_Instance *GameKeyEventHandler = NewGameKeyEventHandler()

func (this *GameKeyEventHandler) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.PressGameKeyEventHandler = PressGameKeyEventHandler_GetInstance()
    this.UpGameKeyEventHandler = UpGameKeyEventHandler_GetInstance()
    this.DownGameKeyEventHandler = DownGameKeyEventHandler_GetInstance()
}

func NewGameKeyEventHandler() *GameKeyEventHandler {
    this := &GameKeyEventHandler{}
    this.initFields()
    return this
}


func GameKeyEventHandler_GetInstance() *GameKeyEventHandler {
        return GameKeyEventHandler_Instance
}


// synchronized
func (this *GameKeyEventHandler) RemoveAllListeners() {
        this.PressGameKeyEventHandler.RemoveAllListeners()
        this.UpGameKeyEventHandler.RemoveAllListeners()
        this.DownGameKeyEventHandler.RemoveAllListeners()
}


func (this *GameKeyEventHandler) AddListenerForPlayer(eventListenerInterface EventListenerInterface, playerInputId int) {
        this.LogUtil.PutF(fmt.Sprint(eventListenerInterface), this, this.CommonStrings.ADD_LISTENER)
        this.PressGameKeyEventHandler.AddListenerSingleThreaded(eventListenerInterface)
        this.UpGameKeyEventHandler.GetInstanceForPlayer(playerInputId).AddListenerSingleThreaded(eventListenerInterface)
        this.DownGameKeyEventHandler.GetInstanceForPlayer(playerInputId).AddListenerSingleThreaded(eventListenerInterface)
}


func (this *GameKeyEventHandler) AddListener(eventListenerInterface EventListenerInterface) {
        this.LogUtil.PutF(fmt.Sprint(eventListenerInterface), this, this.CommonStrings.ADD_LISTENER)
        this.PressGameKeyEventHandler.AddListenerSingleThreaded(eventListenerInterface)
        this.UpGameKeyEventHandler.AddListenerSingleThreaded(eventListenerInterface)
        this.DownGameKeyEventHandler.AddListenerSingleThreaded(eventListenerInterface)
}


func (this *GameKeyEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.LogUtil.PutF(fmt.Sprint(eventListenerInterface), this, this.CommonStrings.REMOVE_LISTENER)
        this.PressGameKeyEventHandler.RemoveListenerSingleThreaded(eventListenerInterface)
        this.UpGameKeyEventHandler.RemoveListenerSingleThreaded(eventListenerInterface)
        this.DownGameKeyEventHandler.RemoveListenerSingleThreaded(eventListenerInterface)
}

