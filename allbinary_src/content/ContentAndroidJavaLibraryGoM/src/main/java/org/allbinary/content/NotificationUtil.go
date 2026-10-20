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

type NotificationUtil struct {
    LogUtil *LogUtil
    NotificationManager *NotificationManager
    NotificationBuilder INotificationBuilder
}

var NotificationUtil_SINGLETON *NotificationUtil = NewNotificationUtil()

func (this *NotificationUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.NotificationManager = any(ResourceUtil_GetInstance().GetContext().GetSystemService(Context_NOTIFICATION_SERVICE)).(*NotificationManager)
    
}


func NotificationUtil_GetInstance() *NotificationUtil {
        return NotificationUtil_SINGLETON
}


func NewNotificationUtil() *NotificationUtil {
    this := &NotificationUtil{}
    this.initFields()
        var SDK_VERSION int = AndroidInfoFactory_GetInstance().GetVersion(); _ = SDK_VERSION
        if SDK_VERSION > 22 {
            this.NotificationBuilder = NewNotificationBuilderAPI23()
        } else {
            this.NotificationBuilder = NotificationBuilder_NULL_NOTIFICATION_BUILDER
        }
    return this
}


func (this *NotificationUtil) Notify(command ICommand, resource string, message string) {
        var context *Context = ResourceUtil_GetInstance().GetContext(); _ = context
        var intent *Intent = CommandUriAction_GetInstance().GetIntent(command); _ = intent
        var integer int = ResourceUtil_GetInstance().GetResourceId(resource); _ = integer
        var FLAG_IMMUTABLE int = 1 << 26; _ = FLAG_IMMUTABLE
        var SDK_VERSION int = AndroidInfoFactory_GetInstance().GetVersion(); _ = SDK_VERSION
        var pendingIntent *PendingIntent = PendingIntent_GetActivity(context, 0, intent, func() int { if SDK_VERSION > 22 { return FLAG_IMMUTABLE }; return 0 }()); _ = pendingIntent
        if this.NotificationBuilder == nil {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.EXCEPTION, this, commonStrings.NOT_IMPLEMENTED)
            return
        }
        this.NotificationBuilder.Build(context, command, message, integer, pendingIntent)
}

