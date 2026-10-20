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

type NoGameNotificationHudSingleton struct {
    LogUtil *LogUtil
    GameNotificationHud IGameNotificationHud
}

var NoGameNotificationHudSingleton_Instance *NoGameNotificationHudSingleton = NewNoGameNotificationHudSingleton()

func (this *NoGameNotificationHudSingleton) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GameNotificationHud = GameNotificationHud_NULL_GAME_NOTIFICATION
}


func NoGameNotificationHudSingleton_GetInstance() *NoGameNotificationHudSingleton {
        return NoGameNotificationHudSingleton_Instance
}


func NewNoGameNotificationHudSingleton() *NoGameNotificationHudSingleton {
    this := &NoGameNotificationHudSingleton{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
                }
            }()
            var basicHudFactory *BasicHudFactory = BasicHudFactory_GetInstance(); _ = basicHudFactory
            /* Java local class: move it out of the function
type NoGameGameNotificationHud struct {
    GameNotificationHud
}

            

            func NewNoGameNotificationHudSingleton() *NoGameNotificationHudSingleton {
                this := &NoGameNotificationHudSingleton{}
                    
                    this.UpdateMaxWidth = 40
                    this.UpdateMaxHeight = 14
                return this
            }
*/
            this.GameNotificationHud = NewNoGameGameNotificationHud()
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *NoGameNotificationHudSingleton) GetGameNotificationHud() IGameNotificationHud {
        return this.GameNotificationHud
}

