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

type AuthenticationRequestHelper struct {
    TagHelper
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    WeblisketSession *WeblisketSession
    Request *HttpServletRequest
}


func (this *AuthenticationRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
    
}




// Java overloads of new AuthenticationRequestHelper
func NewAuthenticationRequestHelper(args ...any) *AuthenticationRequestHelper {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*PageContext](args[1], pass) {
            return NewAuthenticationRequestHelper_1(Java_Arg[*HashMap](args[0]), Java_Arg[*PageContext](args[1]))
        }
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*HttpServletRequest](args[1], pass) {
            return NewAuthenticationRequestHelper_2(Java_Arg[*HashMap](args[0]), Java_Arg[*HttpServletRequest](args[1]))
        }
    }
    panic("No Java overload of new AuthenticationRequestHelper for the arguments")
}
func NewAuthenticationRequestHelper_1(hashMap *HashMap, pageContext *PageContext) *AuthenticationRequestHelper {
    this := &AuthenticationRequestHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.WeblisketSession = NewWeblisketSession(hashMap, pageContext)
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
    return this
}


func NewAuthenticationRequestHelper_2(hashMap *HashMap, httpServletRequest *HttpServletRequest) *AuthenticationRequestHelper {
    this := &AuthenticationRequestHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.WeblisketSession = NewWeblisketSession(hashMap, httpServletRequest)
        this.Request = httpServletRequest
    return this
}


// throws Exception
func (this *AuthenticationRequestHelper) GenerateNewPassword() string {
        if this.WeblisketSession != nil && this.WeblisketSession.GetId() != nil {
            var startIndex int = len(this.WeblisketSession.GetId()); _ = startIndex
            if startIndex >= 8 {
                return this.WeblisketSession.GetId()[startIndex - 8:]
            } else {
                panic(errors.New(fmt.Sprint("Error Generating New Password")))
            }
        } else {
            panic(errors.New(fmt.Sprint("No Session Available For Generating New Password")))
        }
}


func (this *AuthenticationRequestHelper) NewPassword() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "newPassword()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var userName string = this.Request.GetParameter(WeblisketSessionData_REMOVABLEUSERNAME); _ = userName
            var email string = this.Request.GetParameter(UserData_MAINEMAIL); _ = email
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                this.LogUtil.PutF("Generating New Password For: " + userName, this, "newPassword()")
            }
            var userEntityInterface UserEntityInterface = UserEntityFactory_GetInstance(); _ = userEntityInterface
            var userInterface UserInterface = userEntityInterface.GetUser(userName); _ = userInterface
            if strings.Compare(userInterface.GetMainEmail(), email) != 0 {
                if userInterface.GetSecondaryEmail() != nil && strings.Compare(userInterface.GetSecondaryEmail(), email) != 0 {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                        this.LogUtil.PutF("Requested Email Is Not In User Profile", this, "newPassword()")
                    }
                    tryControl1, tryResult1 = 1, false; return
                }
            }
            var newPassword string = this.GenerateNewPassword(); _ = newPassword
            userInterface.SetPassword(newPassword)
            var newPasswordHashMap *HashMap = userInterface.ToPasswordHashMap(); _ = newPasswordHashMap
            UserEntityFactory_GetInstance().Update(userName, newPasswordHashMap)
            NewNewPasswordEmail(this.AbeClientInformation, userInterface, newPassword).Process()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                this.LogUtil.PutF("Generated New Password For: " + userName, this, "newPassword()")
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AuthenticationRequestHelper) ChangePassword() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put("Authentication Failed", this, "changePassword()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var userName string = this.Request.GetParameter(WeblisketSessionData_REMOVABLEUSERNAME); _ = userName
            var passwordString string = this.Request.GetParameter(WeblisketSessionData_REMOVABLEPASSWORD); _ = passwordString
            var newPassword string = this.Request.GetParameter(WeblisketSessionData_REMOVABLENEWPASSWORD); _ = newPassword
            var newReenteredPassword string = this.Request.GetParameter(WeblisketSessionData_REMOVABLEREENTERNEWPASSWORD); _ = newReenteredPassword
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                this.LogUtil.PutF("Changing Password For User: " + userName, this, "changePassword()")
            }
            var password *Password = NewPassword(newPassword); _ = password
            if !password.IsValid() {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                    this.LogUtil.PutF("New Password Is Not Valid", this, "changePassword()")
                }
                tryControl1, tryResult1 = 1, false; return
            }
            if strings.Compare(newPassword, newReenteredPassword) != 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                    this.LogUtil.PutF("New Password Fields Do Not Match", this, "changePassword()")
                }
                tryControl1, tryResult1 = 1, false; return
            }
            var sessionPassword string = this.WeblisketSession.GetPassword(); _ = sessionPassword
            var userEntityInterface UserEntityInterface = UserEntityFactory_GetInstance(); _ = userEntityInterface
            var userInterface UserInterface = userEntityInterface.GetUser(userName); _ = userInterface
            var login string = userEntityInterface.Login(userName, passwordString); _ = login
            if strings.Compare(login, TagHelper_GLOBALS2.LOGINSUCCESS) == 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                    this.LogUtil.PutF("Authentication Successful", this, "changePassword()")
                }
                var newPasswordHashMap *HashMap = password.ToHashMap(nil); _ = newPasswordHashMap
                UserEntityFactory_GetInstance().Update(userName, newPasswordHashMap)
                NewNewPasswordEmail(this.AbeClientInformation, userInterface, newPassword).Process()
                tryControl1, tryResult1 = 1, true; return
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                this.LogUtil.PutF("Authentication Failed", this, "changePassword()")
            }
            tryControl1, tryResult1 = 1, false; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AuthenticationRequestHelper) IsRoleValid(userName string, password string, roles IBasicArrayList) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "isRoleValid()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var userEntityInterface UserEntityInterface = UserEntityFactory_GetInstance(); _ = userEntityInterface
            var userInterface UserInterface = userEntityInterface.GetUser(userName); _ = userInterface
            if userInterface.GetRole() == nil {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                    this.LogUtil.PutF("Role is null: " + fmt.Sprint(userInterface.GetRole()) + " Valid Roles: " + fmt.Sprint(roles), this, "isRoleValid()")
                }
                tryControl1, tryResult1 = 1, false; return
            }
            var basicUserRoleArray []any = roles.ToArray(); _ = basicUserRoleArray
            var size int = len(basicUserRoleArray); _ = size
            for index := int(0); index < size; index++ {
                var nextRole *BasicUserRole = any(basicUserRoleArray[index]).(*BasicUserRole); _ = nextRole
                if Java_Equals(userInterface.GetRole().GetBasicUserRole(), nextRole) {
                    userInterface.ValidateSession(any(this.WeblisketSession).(WeblisketSessionInterface))
                    this.Request.RemoveAttribute(WeblisketSessionData_REMOVABLEUSERNAME)
                    this.Request.RemoveAttribute(WeblisketSessionData_REMOVABLEPASSWORD)
                    tryControl1, tryResult1 = 1, true; return
                }
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Role is not valid: ")
                stringBuffer.Append(fmt.Sprint(userInterface.GetRole()))
                stringBuffer.Append(" Valid Roles: ")
                stringBuffer.Append(fmt.Sprint(roles))
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "isRoleValid()")
            }
            tryControl1, tryResult1 = 1, false; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

