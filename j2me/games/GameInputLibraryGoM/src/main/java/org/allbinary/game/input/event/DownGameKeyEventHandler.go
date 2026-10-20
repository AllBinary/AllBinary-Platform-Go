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

type DownGameKeyEventHandler struct {
    DownGameKeyEventHandlerBase
    PlayerInputIdFactory *LocalPlayerInputIdFactory
}

var DownGameKeyEventHandler_Instance *DownGameKeyEventHandler = NewDownGameKeyEventHandler()
var DownGameKeyEventHandler_InstanceArray []IDownGameKeyEventHandlerBase = []IDownGameKeyEventHandlerBase{NewDownGameKeyEventHandlerBase(), NewDownGameKeyEventHandlerBase(), NewDownGameKeyEventHandlerBase(), NewDownGameKeyEventHandlerBase(), NewDownGameKeyEventHandlerBase(), NewDownGameKeyEventHandlerBase()}

func (this *DownGameKeyEventHandler) initFields() {
    this.PlayerInputIdFactory = LocalPlayerInputIdFactory_GetInstance()
}


func DownGameKeyEventHandler_GetInstance() *DownGameKeyEventHandler {
        return DownGameKeyEventHandler_Instance
}


func (this *DownGameKeyEventHandler) GetInstanceForDevice(deviceId int) IDownGameKeyEventHandlerBase {
        var playerInputId int = this.PlayerInputIdFactory.GetPlayerForDevice(deviceId); _ = playerInputId
        return DownGameKeyEventHandler_InstanceArray[playerInputId]
}


func (this *DownGameKeyEventHandler) GetInstanceForPlayer(playerInputId int) IDownGameKeyEventHandlerBase {
        return DownGameKeyEventHandler_InstanceArray[playerInputId]
}


func NewDownGameKeyEventHandler() *DownGameKeyEventHandler {
    this := &DownGameKeyEventHandler{}
    this.DownGameKeyEventHandlerBase = *NewDownGameKeyEventHandlerBase()
    this.initFields()
    return this
}


func (this *DownGameKeyEventHandler) RemoveAllListeners() {
        this.DownGameKeyEventHandlerBase.RemoveAllListeners()
        for index := int(len(DownGameKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            DownGameKeyEventHandler_InstanceArray[index].RemoveAllListeners()
        }
}


func (this *DownGameKeyEventHandler) RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        this.DownGameKeyEventHandlerBase.RemoveListenerSingleThreaded(eventListenerInterface)
        for index := int(len(DownGameKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            DownGameKeyEventHandler_InstanceArray[index].RemoveListenerSingleThreaded(eventListenerInterface)
        }
}


// synchronized
func (this *DownGameKeyEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.DownGameKeyEventHandlerBase.RemoveListener(eventListenerInterface)
        for index := int(len(DownGameKeyEventHandler_InstanceArray) - 1); index >= 0; index-- {
            DownGameKeyEventHandler_InstanceArray[index].RemoveListener(eventListenerInterface)
        }
}

