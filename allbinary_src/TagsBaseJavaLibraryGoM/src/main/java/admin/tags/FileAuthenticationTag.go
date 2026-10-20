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

type FileAuthenticationTag struct {
    CustomTagSupport
    LogUtil *LogUtil
    AttemptsInteger int
    WeblisketSession *BasicWeblisketSession
    Request *HttpServletRequest
    Command string
    UserName string
    Password string
    NewPassword string
    Authenticated string
    SessionUserName string
    Role *BasicUserRole
    Timeout string
    Roles *BasicArrayList
}


func (this *FileAuthenticationTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewFileAuthenticationTag() *FileAuthenticationTag {
    this := &FileAuthenticationTag{}
    this.CustomTagSupport = *NewCustomTagSupport()
    this.initFields()
    return this
}


func (this *FileAuthenticationTag) SetCommand(command string) {
        this.Command = command
}


func (this *FileAuthenticationTag) SetUserName(userName string) {
        this.UserName = userName
}


func (this *FileAuthenticationTag) SetPassword(value string) {
        this.Password = value
}


func (this *FileAuthenticationTag) SetNewPassword(value string) {
        this.NewPassword = value
}


func (this *FileAuthenticationTag) SetRoles(values *BasicArrayList) {
        this.Roles = values
}


func (this *FileAuthenticationTag) MakeSessionValid(storedRole *BasicUserRole, userName string) {
        this.WeblisketSession.SetAuthenticated()
        this.WeblisketSession.SetRole(storedRole)
        this.WeblisketSession.SetUserName(userName)
        this.Request.RemoveAttribute(WeblisketSessionData_REMOVABLEUSERNAME)
        this.Request.RemoveAttribute(WeblisketSessionData_REMOVABLEPASSWORD)
}


func (this *FileAuthenticationTag) GetFormData() {
        this.Authenticated = this.WeblisketSession.GetAuthentication()
        this.SessionUserName = this.WeblisketSession.GetUserName()
        this.Role = this.WeblisketSession.GetRole()
        this.AttemptsInteger = this.WeblisketSession.GetAttempts()
        this.Timeout = this.WeblisketSession.GetTimeout()
}


func (this *FileAuthenticationTag) NextAttempt() {
        if this.AttemptsInteger != 0 && int(this.AttemptsInteger) > 0 {
            this.WeblisketSession.SetAttempts(int(int(this.AttemptsInteger) + 1))
        } else {
            this.WeblisketSession.SetAttempts(int(1))
        }
}


// throws LicensingException
func (this *FileAuthenticationTag) ChangePassword() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var userName string = this.Request.GetParameter(WeblisketSessionData_REMOVABLEUSERNAME); _ = userName
            var password string = this.Request.GetParameter(WeblisketSessionData_REMOVABLEPASSWORD); _ = password
            var newPassword string = this.Request.GetParameter(WeblisketSessionData_REMOVABLENEWPASSWORD); _ = newPassword
            var newReenteredPassword string = this.Request.GetParameter(WeblisketSessionData_REMOVABLEREENTERNEWPASSWORD); _ = newReenteredPassword
            if strings.Compare(newPassword, newReenteredPassword) != 0 {
                tryControl1, tryResult1 = 1, false; return
            }
            var installerInfo *InstallerInfo = NewInstallerInfo(); _ = installerInfo
            if installerInfo.IsValid(userName, password) {
                installerInfo.SetPassword(newPassword)
                installerInfo.Write()
                this.WeblisketSession.SetUserName(userName)
                this.WeblisketSession.SetPassword(newPassword)
                tryControl1, tryResult1 = 1, true; return
            }
            tryControl1, tryResult1 = 1, false; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *FileAuthenticationTag) InvalidateSession() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "invalidateSession()", e)
            }
            tryControl1, tryResult1 = 1, TagSupport_SKIP_BODY; return
                }
            }()
            this.WeblisketSession.Clear()
            tryControl1, tryResult1 = 1, int(int(TagSupport_EVAL_BODY_INCLUDE)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *FileAuthenticationTag) IsSessionOld() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "isSessionOld()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if this.Role != nil {
                var timeCreated int64 = this.WeblisketSession.GetCreationTime(); _ = timeCreated
                var sessionTimout int64 = this.Role.GetSessionTimeout(); _ = sessionTimout
                var calendar *Calendar = Calendar_GetInstance(); _ = calendar
                var timeFirst int64 = calendar.GetTimeInMillis() - sessionTimout; _ = timeFirst
                if timeFirst > timeCreated {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATION) {
                        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                        stringBuffer.Append("Session Is Old - Timeout: ")
                        stringBuffer.Appendlong(sessionTimout)
                        stringBuffer.Append(" Eval: ")
                        stringBuffer.Appendlong(timeFirst)
                        stringBuffer.Append(" > ")
                        stringBuffer.Appendlong(timeCreated)
                        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "isSessionOld()")
                    }
                    tryControl1, tryResult1 = 1, true; return
                }
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATION) {
                    this.LogUtil.PutF("Role Is Not In Session", this, "isSessionOld()")
                }
            }
            tryControl1, tryResult1 = 1, this.IsRarelyUsedSession(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *FileAuthenticationTag) IsRarelyUsedSession() bool {
        if this.Role != nil {
            var timeCreated int64 = this.WeblisketSession.GetCreationTime(); _ = timeCreated
            var lastAccess int64 = this.WeblisketSession.GetLastAccessedTime(); _ = lastAccess
            var timePassed int64 = this.Role.GetSessionInactivityTimeout(); _ = timePassed
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var timeFirst int64 = calendar.GetTimeInMillis() - lastAccess; _ = timeFirst
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATION) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Session Is Rarely Used - Timeout: ")
                stringBuffer.Appendlong(lastAccess)
                stringBuffer.Append(" Eval: ")
                stringBuffer.Appendlong(timeFirst)
                stringBuffer.Append(" > ")
                stringBuffer.Appendlong(timePassed)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "isRarelyUsedSession()")
            }
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATION) {
                this.LogUtil.PutF("Role Is Not In Session", this, "isRarelyUsedSession()")
            }
        }
        return false
}


// throws LicensingException
func (this *FileAuthenticationTag) IsRoleValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "isRoleValid()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.MakeSessionValid(BasicUserRoleFactory_GetInstance().INSTALLER, this.UserName)
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *FileAuthenticationTag) ValidRole() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to set valid role"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "validRole()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Trying New login<p>")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *FileAuthenticationTag) InvalidRole() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to set role invalid"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "invalidRole()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            this.WeblisketSession.SetAuthenticated(false)
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Sorry your username and password is invalid on this page.")
            stringBuffer.Append("Trying New login<p>")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *FileAuthenticationTag) IsAuthenticationSessionValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "isAuthenticationSessionValid()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if BooleanUtil_GetInstance().GetFromString(this.Authenticated) {
                var size int = this.Roles.Size(); _ = size
                for index := int(0); index < size; index++ {
                    var mustBeOfRole *BasicUserRole = any(this.Roles.Get(index)).(*BasicUserRole); _ = mustBeOfRole
                    if this.SessionUserName != "" && Java_Equals(this.Role, mustBeOfRole) {
                        tryControl1, tryResult1 = 1, true; return
                    }
                }
                tryControl1, tryResult1 = 1, false; return
            }
            tryControl1, tryResult1 = 1, false; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *FileAuthenticationTag) ProcessIfNewLogin() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "processIfNewLogin()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.PutF("Starting for User: " + this.UserName + " Attempt: " + fmt.Sprint(this.WeblisketSession.GetAttempts()), this, "processIfNewLogin()")
            }
            if this.UserName != "" && strings.Compare(this.UserName, StringUtil_GetInstance().EMPTY_STRING) != 0 && this.Password != "" && strings.Compare(this.Password, StringUtil_GetInstance().EMPTY_STRING) != 0 {
                var login string; _ = login
                var installerInfo *InstallerInfo = NewInstallerInfo(); _ = installerInfo
                if installerInfo.IsValid(this.UserName, this.Password) {
                    login = org_Allbinary.Globals.GLOBALS2.LOGINSUCCESS
                } else {
                    login = org_Allbinary.Globals.GLOBALS2.LOGINFAILED
                }
                if strings.Compare(login, org_Allbinary.Globals.GLOBALS2.LOGINSUCCESS) == 0 {
                    this.WeblisketSession.SetAttempts(int(0))
                    tryControl1, tryResult1 = 1, true; return
                }
                this.NextAttempt()
                tryControl1, tryResult1 = 1, false; return
            }
            this.NextAttempt()
            tryControl1, tryResult1 = 1, false; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *FileAuthenticationTag) ProcessInvalidation() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "isSetToInvalidate()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if BooleanUtil_GetInstance().GetFromString(this.Timeout) {
                this.InvalidateSession()
                tryControl1, tryResult1 = 1, true; return
            }
            {
                this.WeblisketSession.SetTimeout(BooleanFactory_GetInstance().TRUE_STRING)
                tryControl1, tryResult1 = 1, false; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *FileAuthenticationTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
                }
            }()
            var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
            this.Request = any(this.PageContext.GetRequest()).(*HttpServletRequest)
            this.WeblisketSession = NewBasicWeblisketSession(this.PageContext)
            this.GetFormData()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                this.LogUtil.PutF("Started", this, this.TagStrings.DO_START_TAG)
            }
            if this.Command != "" {
                if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.CHANGEPASSWORD) == 0 {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                        this.LogUtil.PutF("Change Password", this, this.TagStrings.DO_START_TAG)
                    }
                    if !this.ChangePassword() {
                        this.PageContext.GetOut().Print("Password Change Attempt Failed.<p>")
                        tryControl1, tryResult1 = 1, CustomTagSupport_EVAL_BODY_INCLUDE; return
                    } else {
                        this.PageContext.GetOut().Print("Password Changed Successfully.<p>")
                        tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
                    }
                } else if strings.Compare(this.Command, WeblisketSessionData_INVALIDATESESSION) == 0 {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                        this.LogUtil.PutF("INVALIDATESESSION", this, this.TagStrings.DO_START_TAG)
                    }
                    tryControl1, tryResult1 = 1, this.InvalidateSession(); return
                } else {
                    if this.IsSessionOld() {
                        if this.ProcessInvalidation() {
                            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                                this.LogUtil.PutF("Processing Invalidation", this, this.TagStrings.DO_START_TAG)
                            }
                            this.PageContext.GetOut().Print("Please login again.<p>")
                            if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.PROCESSBODYIFAUTHENTICATED) == 0 {
                                tryControl1, tryResult1 = 1, TagSupport_SKIP_BODY; return
                            } else {
                                tryControl1, tryResult1 = 1, TagSupport_EVAL_BODY_INCLUDE; return
                            }
                        } else {
                            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                                this.LogUtil.PutF("Preparing session for invalidation", this, this.TagStrings.DO_START_TAG)
                            }
                            if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.PROCESSBODYIFAUTHENTICATED) == 0 {
                                tryControl1, tryResult1 = 1, TagSupport_SKIP_BODY; return
                            } else {
                                tryControl1, tryResult1 = 1, TagSupport_EVAL_BODY_INCLUDE; return
                            }
                        }
                    }
                    if (this.UserName == "" || strings.Compare(this.UserName, stringUtil.EMPTY_STRING) == 0) && (this.Password == "" || strings.Compare(this.Password, stringUtil.EMPTY_STRING) == 0) && this.Roles != nil && this.IsAuthenticationSessionValid() {
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                            this.LogUtil.PutF("Authenticated session is valid", this, this.TagStrings.DO_START_TAG)
                        }
                        if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.PROCESSBODYIFAUTHENTICATED) == 0 {
                            tryControl1, tryResult1 = 1, TagSupport_EVAL_BODY_INCLUDE; return
                        } else {
                            tryControl1, tryResult1 = 1, TagSupport_SKIP_BODY; return
                        }
                    }
                    if this.UserName != "" && this.Password != "" && this.ProcessIfNewLogin() {
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                            this.LogUtil.PutF("Processing new login.", this, this.TagStrings.DO_START_TAG)
                        }
                        if this.IsRoleValid() {
                            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                                this.LogUtil.PutF("Role is valid - Processing new login", this, this.TagStrings.DO_START_TAG)
                            }
                            this.PageContext.GetOut().Print(this.ValidRole())
                            if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.PROCESSBODYIFAUTHENTICATED) == 0 {
                                tryControl1, tryResult1 = 1, CustomTagSupport_EVAL_BODY_INCLUDE; return
                            } else {
                                tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
                            }
                        } else {
                            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                                this.LogUtil.PutF("Role is not valid - Processing new login", this, this.TagStrings.DO_START_TAG)
                            }
                            this.PageContext.GetOut().Print(this.InvalidRole())
                            if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.PROCESSBODYIFAUTHENTICATED) == 0 {
                                tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
                            } else {
                                tryControl1, tryResult1 = 1, CustomTagSupport_EVAL_BODY_INCLUDE; return
                            }
                        }
                    } else {
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                            this.LogUtil.PutF("Invalid Login", this, this.TagStrings.DO_START_TAG)
                        }
                        if this.UserName != "" && strings.Compare(this.UserName, StringUtil_GetInstance().EMPTY_STRING) != 0 && this.Password != "" && strings.Compare(this.Password, StringUtil_GetInstance().EMPTY_STRING) != 0 {
                            this.PageContext.GetOut().Print("Sorry your username and/or password is invalid.<p>")
                        }
                    }
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                        stringBuffer.Append("Major authentication error - userName: ")
                        stringBuffer.Append(this.UserName)
                        stringBuffer.Append(" Password: ")
                        stringBuffer.Append(this.Password)
                        stringBuffer.Append(" Command: ")
                        stringBuffer.Append(this.Command)
                        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.TagStrings.DO_START_TAG)
                    }
                    if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.PROCESSBODYIFAUTHENTICATED) == 0 {
                        tryControl1, tryResult1 = 1, TagSupport_SKIP_BODY; return
                    } else {
                        tryControl1, tryResult1 = 1, TagSupport_EVAL_BODY_INCLUDE; return
                    }
                }
            }
            tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

