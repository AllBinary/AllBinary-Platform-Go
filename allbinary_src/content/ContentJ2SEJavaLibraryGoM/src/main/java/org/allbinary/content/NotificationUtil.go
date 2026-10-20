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
    NotificationManager *NotificationManager
}

var NotificationUtil_SINGLETON *NotificationUtil = NewNotificationUtil()

func (this *NotificationUtil) initFields() {
    this.NotificationManager = any(ResourceUtil_GetInstance().GetContext().GetSystemService(Context_NOTIFICATION_SERVICE)).(*NotificationManager)
}

func NewNotificationUtil() *NotificationUtil {
    this := &NotificationUtil{}
    this.initFields()
    return this
}


func NotificationUtil_GetInstance() *NotificationUtil {
        return NotificationUtil_SINGLETON
}


func (this *NotificationUtil) Notify(command ICommand, resource string, message string) {
        var context *Context = ResourceUtil_GetInstance().GetContext(); _ = context
        var intent *Intent = CommandUriAction_GetInstance().GetIntent(command); _ = intent
        var integer int = ResourceUtil_GetInstance().GetResourceId(resource); _ = integer
        var notification *Notification = NewNotification(int(integer), message, time.Now().UnixMilli()); _ = notification
        var pendingIntent *PendingIntent = PendingIntent_GetActivity(context, 0, intent, 0); _ = pendingIntent
        notification.SetLatestEventInfo(context, command.GetLabel(), message, pendingIntent)
        this.NotificationManager.Notify(TsUtil_GetInstance().HashCode(command), notification)
}

