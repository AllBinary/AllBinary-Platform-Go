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

type UpGameKeyEventHandler struct {
    UpGameKeyEventHandlerBase
    PlayerInputIdFactory *LocalPlayerInputIdFactory
}

var UpGameKeyEventHandler_Instance *UpGameKeyEventHandler = NewUpGameKeyEventHandler()
var UpGameKeyEventHandler_InstanceArray []IUpGameKeyEventHandlerBase = []IUpGameKeyEventHandlerBase{NewUpGameKeyEventHandlerBase(), NewUpGameKeyEventHandlerBase(), NewUpGameKeyEventHandlerBase(), NewUpGameKeyEventHandlerBase(), NewUpGameKeyEventHandlerBase(), NewUpGameKeyEventHandlerBase()}

func (this *UpGameKeyEventHandler) initFields() {
    this.PlayerInputIdFactory = LocalPlayerInputIdFactory_GetInstance()
}


func UpGameKeyEventHandler_GetInstance() *UpGameKeyEventHandler {
        return UpGameKeyEventHandler_Instance
}


func (this *UpGameKeyEventHandler) GetInstanceForDevice(deviceId int) IUpGameKeyEventHandlerBase {
        var playerInputId int = this.PlayerInputIdFactory.GetPlayerForDevice(deviceId); _ = playerInputId
        return UpGameKeyEventHandler_InstanceArray[playerInputId]
}


func (this *UpGameKeyEventHandler) GetInstanceForPlayer(playerInputId int) IUpGameKeyEventHandlerBase {
        return UpGameKeyEventHandler_InstanceArray[playerInputId]
}


func NewUpGameKeyEventHandler() *UpGameKeyEventHandler {
    this := &UpGameKeyEventHandler{}
    this.UpGameKeyEventHandlerBase = *NewUpGameKeyEventHandlerBase()
    this.initFields()
    return this
}


func (this *UpGameKeyEventHandler) RemoveAllListeners() {
        this.UpGameKeyEventHandlerBase.RemoveAllListeners()
        for index := int(len(UpGameKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            UpGameKeyEventHandler_InstanceArray[index].RemoveAllListeners()
        }
}


func (this *UpGameKeyEventHandler) RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        this.UpGameKeyEventHandlerBase.RemoveListenerSingleThreaded(eventListenerInterface)
        for index := int(len(UpGameKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            UpGameKeyEventHandler_InstanceArray[index].RemoveListenerSingleThreaded(eventListenerInterface)
        }
}


func (this *UpGameKeyEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.UpGameKeyEventHandlerBase.RemoveListener(eventListenerInterface)
        for index := int(len(UpGameKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            UpGameKeyEventHandler_InstanceArray[index].RemoveListener(eventListenerInterface)
        }
}

