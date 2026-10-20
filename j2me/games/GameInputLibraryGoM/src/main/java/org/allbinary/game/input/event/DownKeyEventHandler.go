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

type DownKeyEventHandler struct {
    DownKeyEventHandlerBase
    PlayerInputIdFactory *LocalPlayerInputIdFactory
}

var DownKeyEventHandler_Instance *DownKeyEventHandler = NewDownKeyEventHandler()
var DownKeyEventHandler_InstanceArray []IDownKeyEventHandlerBase = []IDownKeyEventHandlerBase{NewDownKeyEventHandlerBase(), NewDownKeyEventHandlerBase(), NewDownKeyEventHandlerBase(), NewDownKeyEventHandlerBase(), NewDownKeyEventHandlerBase(), NewDownKeyEventHandlerBase()}

func (this *DownKeyEventHandler) initFields() {
    this.PlayerInputIdFactory = LocalPlayerInputIdFactory_GetInstance()
}


func DownKeyEventHandler_GetInstance() *DownKeyEventHandler {
        return DownKeyEventHandler_Instance
}


func (this *DownKeyEventHandler) GetInstanceForDevice(deviceId int) IDownKeyEventHandlerBase {
        var playerInputId int = this.PlayerInputIdFactory.GetPlayerForDevice(deviceId); _ = playerInputId
        return DownKeyEventHandler_InstanceArray[playerInputId]
}


func (this *DownKeyEventHandler) GetInstanceForPlayer(playerInputId int) IDownKeyEventHandlerBase {
        return DownKeyEventHandler_InstanceArray[playerInputId]
}


func NewDownKeyEventHandler() *DownKeyEventHandler {
    this := &DownKeyEventHandler{}
    this.DownKeyEventHandlerBase = *NewDownKeyEventHandlerBase()
    this.initFields()
    return this
}


func (this *DownKeyEventHandler) RemoveAllListeners() {
        this.DownKeyEventHandlerBase.RemoveAllListeners()
        for index := int(len(DownKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            DownKeyEventHandler_InstanceArray[index].RemoveAllListeners()
        }
}


func (this *DownKeyEventHandler) RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        this.DownKeyEventHandlerBase.RemoveListenerSingleThreaded(eventListenerInterface)
        for index := int(len(DownKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            DownKeyEventHandler_InstanceArray[index].RemoveListenerSingleThreaded(eventListenerInterface)
        }
}


// synchronized
func (this *DownKeyEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.DownKeyEventHandlerBase.RemoveListener(eventListenerInterface)
        for index := int(len(DownKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            DownKeyEventHandler_InstanceArray[index].RemoveListener(eventListenerInterface)
        }
}

