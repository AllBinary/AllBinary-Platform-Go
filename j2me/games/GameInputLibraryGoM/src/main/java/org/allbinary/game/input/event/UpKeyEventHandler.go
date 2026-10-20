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

type UpKeyEventHandler struct {
    UpKeyEventHandlerBase
    PlayerInputIdFactory *LocalPlayerInputIdFactory
}

var UpKeyEventHandler_Instance *UpKeyEventHandler = NewUpKeyEventHandler()
var UpKeyEventHandler_InstanceArray []IUpKeyEventHandlerBase = []IUpKeyEventHandlerBase{NewUpKeyEventHandlerBase(), NewUpKeyEventHandlerBase(), NewUpKeyEventHandlerBase(), NewUpKeyEventHandlerBase(), NewUpKeyEventHandlerBase(), NewUpKeyEventHandlerBase()}

func (this *UpKeyEventHandler) initFields() {
    this.PlayerInputIdFactory = LocalPlayerInputIdFactory_GetInstance()
}


func UpKeyEventHandler_GetInstance() *UpKeyEventHandler {
        return UpKeyEventHandler_Instance
}


func (this *UpKeyEventHandler) GetInstanceForDevice(deviceId int) IUpKeyEventHandlerBase {
        var playerInputId int = this.PlayerInputIdFactory.GetPlayerForDevice(deviceId); _ = playerInputId
        return UpKeyEventHandler_InstanceArray[playerInputId]
}


func (this *UpKeyEventHandler) GetInstanceForPlayer(playerInputId int) IUpKeyEventHandlerBase {
        return UpKeyEventHandler_InstanceArray[playerInputId]
}


func NewUpKeyEventHandler() *UpKeyEventHandler {
    this := &UpKeyEventHandler{}
    this.UpKeyEventHandlerBase = *NewUpKeyEventHandlerBase()
    this.initFields()
    return this
}


func (this *UpKeyEventHandler) RemoveAllListeners() {
        this.UpKeyEventHandlerBase.RemoveAllListeners()
        for index := int(len(UpKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            UpKeyEventHandler_InstanceArray[index].RemoveAllListeners()
        }
}


func (this *UpKeyEventHandler) RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        this.UpKeyEventHandlerBase.RemoveListenerSingleThreaded(eventListenerInterface)
        for index := int(len(UpKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            UpKeyEventHandler_InstanceArray[index].RemoveListenerSingleThreaded(eventListenerInterface)
        }
}


func (this *UpKeyEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.UpKeyEventHandlerBase.RemoveListener(eventListenerInterface)
        for index := int(len(UpKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            UpKeyEventHandler_InstanceArray[index].RemoveListener(eventListenerInterface)
        }
}

