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

type ResolveCompleteMotionGestureListener struct {
    LogUtil *LogUtil
}

// Implements CompleteMotionGestureListenerInterface


func (this *ResolveCompleteMotionGestureListener) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewResolveCompleteMotionGestureListener() *ResolveCompleteMotionGestureListener {
    this := &ResolveCompleteMotionGestureListener{}
    this.initFields()
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF("MotionGesture to CompleteMotionGesture Reciever", this, commonStrings.CONSTRUCTOR)
        CompleteMotionGestureInputEventHandler_GetInstance().AddListener(NewGameKeyFromMotionGestureEventListener())
    return this
}


// throws Exception
func (this *ResolveCompleteMotionGestureListener) OnMotionGestureCompleted(list IBasicArrayList) {
        var configuration *MotionGestureConfiguration = MotionGestureConfigurationFactory_GetInstance(); _ = configuration
        var commandActionsList IBasicArrayList = configuration.GetAssociateCommandActionsList(); _ = commandActionsList
        var size int = commandActionsList.Size(); _ = size
        for index := int(size - 1); index >= 0; index-- {
            var association *MotionGestureToMotionGestureActionAssociation = any(commandActionsList.BasicArrayList_().ObjectArray[index]).(*MotionGestureToMotionGestureActionAssociation); _ = association
            if association.IsMotionGestureArrayEquals(list) == true {
                var completeMotionGestureInputEvent IGameKeyCompleteMotionGestureInputEvent = any(association.GetCommandAction()).(IGameKeyCompleteMotionGestureInputEvent); _ = completeMotionGestureInputEvent
                CompleteMotionGestureInputEventHandler_GetInstance().FireEvent(completeMotionGestureInputEvent)
                break
            }
        }
}

