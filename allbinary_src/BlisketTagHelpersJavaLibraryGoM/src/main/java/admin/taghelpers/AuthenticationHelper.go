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

type AuthenticationHelper struct {
    TagHelper
    LogUtil *LogUtil
    WeblisketSession *WeblisketSession
    Authenticated string
    SessionUserName string
    Role IUserRole
    Timeout string
    AttemptsInteger int
}


func (this *AuthenticationHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
}




// Java overloads of new AuthenticationHelper
func NewAuthenticationHelper(args ...any) *AuthenticationHelper {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*PageContext](args[1], pass) {
            return NewAuthenticationHelper_1(Java_Arg[*HashMap](args[0]), Java_Arg[*PageContext](args[1]))
        }
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*HttpServletRequest](args[1], pass) {
            return NewAuthenticationHelper_2(Java_Arg[*HashMap](args[0]), Java_Arg[*HttpServletRequest](args[1]))
        }
    }
    panic("No Java overload of new AuthenticationHelper for the arguments")
}
func NewAuthenticationHelper_1(hashMap *HashMap, pageContext *PageContext) *AuthenticationHelper {
    this := &AuthenticationHelper{}
        *this = *NewAuthenticationHelper(hashMap, any(pageContext.GetRequest()).(*HttpServletRequest))
    return this
}


func NewAuthenticationHelper_2(hashMap *HashMap, httpServletRequest *HttpServletRequest) *AuthenticationHelper {
    this := &AuthenticationHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.WeblisketSession = NewWeblisketSession(hashMap, httpServletRequest)
        this.GetFormData()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATION) {
            this.LogUtil.PutF(this.OutputSessionInfo(), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


// throws Exception
func (this *AuthenticationHelper) GetFormData() {
        this.Authenticated = this.WeblisketSession.GetAuthentication()
        this.SessionUserName = this.WeblisketSession.GetUserName()
        this.Role = this.WeblisketSession.GetRole()
        this.AttemptsInteger = this.WeblisketSession.GetAttempts()
        this.Timeout = this.WeblisketSession.GetTimeout()
}


// throws Exception
func (this *AuthenticationHelper) InvalidateSession() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATIONERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "invalidateSession()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            this.WeblisketSession.Clear()
            tryControl1, tryResult1 = 1, int(TagSupport_EVAL_BODY_INCLUDE); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AuthenticationHelper) IsSessionOld() bool {
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
                return true
            }
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATION) {
                this.LogUtil.PutF("Role Is Not In Session", this, "isSessionOld()")
            }
        }
        return this.IsRarelyUsedSession()
}


func (this *AuthenticationHelper) IsRarelyUsedSession() bool {
        if this.Role != nil {
            var timeCreated int64 = this.WeblisketSession.GetCreationTime(); _ = timeCreated
            var lastAccess int64 = this.WeblisketSession.GetLastAccessedTime(); _ = lastAccess
            var inactivityAllowed int64 = this.Role.GetSessionInactivityTimeout(); _ = inactivityAllowed
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var timeFirst int64 = calendar.GetTimeInMillis() - lastAccess; _ = timeFirst
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATION) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Session Is Rarely Used - Timeout: ")
                stringBuffer.Appendlong(lastAccess)
                stringBuffer.Append(" Eval: ")
                stringBuffer.Appendlong(timeFirst)
                stringBuffer.Append(" > ")
                stringBuffer.Appendlong(inactivityAllowed)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "isRarelyUsedSession()")
            }
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATION) {
                this.LogUtil.PutF("Role Is Not In Session", this, "isRarelyUsedSession()")
            }
        }
        return false
}


func (this *AuthenticationHelper) ValidRole() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to set valid role"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATIONERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "validRole()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Trying New login<p/>")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AuthenticationHelper) InvalidRole() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to set role invalid"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATIONERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "invalidRole()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            this.WeblisketSession.SetAuthenticated(false)
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Sorry your username and password is invalid on this page. ")
            stringBuffer.Append("Trying New login<p/>")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *AuthenticationHelper) IsAuthenticated() bool {
        if BooleanUtil_GetInstance().GetFromString(this.Authenticated) {
            return true
        } else {
            return false
        }
}


func (this *AuthenticationHelper) IsAuthenticationSessionValid(roles IBasicArrayList) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATIONERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "isAuthenticatedSessionValid()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if this.Role != nil {
                if this.IsAuthenticated() {
                    if this.SessionUserName != "" {
                        var basicUserRole *BasicUserRole = this.Role.GetBasicUserRole(); _ = basicUserRole
                        var basicUserRoleArray []any = roles.ToArray(); _ = basicUserRoleArray
                        var size int = len(basicUserRoleArray); _ = size
                        for index := int(0); index < size; index++ {
                            var mustBeOfRole *BasicUserRole = any(basicUserRoleArray[index]).(*BasicUserRole); _ = mustBeOfRole
                            if Java_Equals(basicUserRole, mustBeOfRole) {
                                if Java_Equals(basicUserRole, UserRoleFactory_GetInstance().CUSTOMER) {
                                    var userEntityInterface UserEntityInterface = UserEntityFactory_GetInstance(); _ = userEntityInterface
                                    var userInterface UserInterface = userEntityInterface.GetUser(this.SessionUserName); _ = userInterface
                                    if !userInterface.IsSessionValid() {
                                        tryControl1, tryResult1 = 1, false; return
                                    }
                                }
                                tryControl1, tryResult1 = 1, true; return
                            }
                        }
                    }
                    tryControl1, tryResult1 = 1, false; return
                }
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATION) {
                    this.LogUtil.PutF("Role Is Not In Session", this, "isAuthenticationSessionValid()")
                }
            }
            tryControl1, tryResult1 = 1, false; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AuthenticationHelper) NextAttempt() {
        if this.AttemptsInteger != 0 && int(this.AttemptsInteger) > 0 {
            this.WeblisketSession.SetAttempts(int(int(this.AttemptsInteger) + 1))
        } else {
            this.WeblisketSession.SetAttempts(int(1))
        }
}


func (this *AuthenticationHelper) ProcessIfNewLogin(userName string, password string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATIONERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "processIfNewLogin()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATIONERROR) {
                this.LogUtil.PutF("Starting: " + userName, this, "processIfNewLogin()")
            }
            if userName != "" && strings.Compare(userName, StringUtil_GetInstance().EMPTY_STRING) != 0 && password != "" && strings.Compare(password, StringUtil_GetInstance().EMPTY_STRING) != 0 {
                var login string; _ = login
                var userEntityInterface UserEntityInterface = UserEntityFactory_GetInstance(); _ = userEntityInterface
                login = userEntityInterface.Login(userName, password)
                if strings.Compare(login, TagHelper_GLOBALS2.LOGINSUCCESS) == 0 {
                    var userInterface UserInterface = userEntityInterface.GetUser(userName); _ = userInterface
                    if userInterface.IsSessionValid() {
                        this.WeblisketSession.SetAttempts(int(0))
                        tryControl1, tryResult1 = 1, true; return
                    }
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


func (this *AuthenticationHelper) ProcessInvalidation() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().AUTHENTICATIONERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "processInvalidation()", e)
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


func (this *AuthenticationHelper) OutputSessionInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, "Error"; return
                }
            }()
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var timeCreated int64 = this.WeblisketSession.GetCreationTime(); _ = timeCreated
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var lastAccess int64 = this.WeblisketSession.GetLastAccessedTime(); _ = lastAccess
            stringBuffer.Append("Time Since Last Access: " + fmt.Sprint((calendar.GetTimeInMillis() - lastAccess)) + "<br/>\n")
            if this.Role != nil {
                stringBuffer.Append("Auth: ")
                stringBuffer.Append(this.Authenticated)
                stringBuffer.Append(" UserName: ")
                stringBuffer.Append(this.SessionUserName)
                stringBuffer.Append(" Role: ")
                stringBuffer.Append(fmt.Sprint(this.Role))
                stringBuffer.Append("<br/>\n")
                stringBuffer.Append("Time Left: ")
                stringBuffer.Appendlong((this.Role.GetSessionTimeout() - (calendar.GetTimeInMillis() - timeCreated)))
                stringBuffer.Append("<br/>\n")
                stringBuffer.Append("Inactivity Time Allowed: ")
                stringBuffer.Appendlong(this.Role.GetSessionTimeout())
                stringBuffer.Append("<br/>\n")
                stringBuffer.Append("Inactivity Time Allowed: ")
                stringBuffer.Appendlong(this.Role.GetSessionInactivityTimeout())
                stringBuffer.Append("<br/>\n")
            } else {
                stringBuffer.Append("Role Is Not Set<br/>\n")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AuthenticationHelper) GetRole() IUserRole {
        return this.Role
}

