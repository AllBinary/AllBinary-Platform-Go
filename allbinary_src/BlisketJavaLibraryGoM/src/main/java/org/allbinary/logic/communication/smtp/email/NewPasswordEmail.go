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

type NewPasswordEmail struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AbeClientInformation AbeClientInformationInterface
    UserInterface UserInterface
    NewPassword string
}


func (this *NewPasswordEmail) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}


func NewNewPasswordEmail(abeClientInformation AbeClientInformationInterface, userInterface UserInterface, newPassword string) *NewPasswordEmail {
    this := &NewPasswordEmail{}
    this.initFields()
        this.AbeClientInformation = abeClientInformation
        this.UserInterface = userInterface
        this.NewPassword = newPassword
    return this
}


// throws Exception
func (this *NewPasswordEmail) Process() {
        this.NotifyStoreAdmin()
        this.NotifyUser()
}


// throws Exception
func (this *NewPasswordEmail) NotifyStoreAdmin() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "emailAdmin", e)
            }
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
                this.LogUtil.PutF("notifyStoreAdmin", this, "notifyStoreAdmin")
            }
            var emailSubject string = "New Password For User: " + fmt.Sprint(this.UserInterface.GetUserName()); _ = emailSubject
            var emailBody string = "New Password: " + this.NewPassword; _ = emailBody
            var basicEmailInfo IBasicEmailInfo = any(NewAdminEmailInfo(emailSubject, emailBody)).(IBasicEmailInfo); _ = basicEmailInfo
            var emailInfo *EmailInfo = NewEmailInfo(basicEmailInfo); _ = emailInfo
            var adminUserEmailEventHandler *UserEmailEventHandler = AdminUserEmailEventHandlerSingletons_GetInstance().GetInstance(this.AbeClientInformation, UserEmailEventNameData_NEWPASSWORD); _ = adminUserEmailEventHandler
            adminUserEmailEventHandler.ReceiveEmailInfo(UserEmailEventNameData_NEWPASSWORD, emailInfo)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *NewPasswordEmail) NotifyUser() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "notifyUser", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
                this.LogUtil.PutF("Email User", this, "notifyUser()")
            }
            var subject string = "New Password"; _ = subject
            var body string = "New Password: " + this.NewPassword; _ = body
            var basicEmailInfo IBasicEmailInfo = any(NewAdminEmailInfo(subject, body)).(IBasicEmailInfo); _ = basicEmailInfo
            var emailInfo *EmailInfo = NewEmailInfo(basicEmailInfo); _ = emailInfo
            var userEmailEventHandler *UserEmailEventHandler = UserEmailEventHandlerSingletons_GetInstance().GetInstance(this.AbeClientInformation, UserEmailEventNameData_NEWPASSWORD, this.UserInterface); _ = userEmailEventHandler
            userEmailEventHandler.ReceiveEmailInfo(UserEmailEventNameData_NEWPASSWORD, emailInfo)
            return
        }(); tryControl1 == 1 {
            return
        }
}

