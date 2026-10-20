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

type LockedFeatureNotificationUtil struct {
    GameNotificationEvent *GameNotificationEvent
}

var LockedFeatureNotificationUtil_Instance *LockedFeatureNotificationUtil = NewLockedFeatureNotificationUtil()

func (this *LockedFeatureNotificationUtil) initFields() {
    this.GameNotificationEvent = NewGameNotificationEvent(this, LicenseStrings_GetInstance().LOCKED, SmallIntegerSingletonFactory_GetInstance().GetAt(3), BasicColorFactory_GetInstance().RED, BooleanFactory_GetInstance().FALSE)
}

func NewLockedFeatureNotificationUtil() *LockedFeatureNotificationUtil {
    this := &LockedFeatureNotificationUtil{}
    this.initFields()
    return this
}


func LockedFeatureNotificationUtil_GetInstance() *LockedFeatureNotificationUtil {
        return LockedFeatureNotificationUtil_Instance
}


// throws Exception
func (this *LockedFeatureNotificationUtil) Fire() {
        PrimaryPlayerQueueFactory_GetInstance().Add(ErrorSound_GetInstance())
        GameNotificationEventHandler_GetInstance().FireEvent(this.GameNotificationEvent)
}


// throws Exception
func (this *LockedFeatureNotificationUtil) FireWithColor(basicColor *BasicColor) {
        PrimaryPlayerQueueFactory_GetInstance().Add(ErrorSound_GetInstance())
        this.GameNotificationEvent.SetBasicColorP(basicColor)
        GameNotificationEventHandler_GetInstance().FireEvent(this.GameNotificationEvent)
}

