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

type AuthenticationTag struct {
    CustomTagSupport
    LogUtil *LogUtil
    Command string
    UserName string
    Password string
    NewPassword string
    Roles IBasicArrayList
    PropertiesHashMap *HashMap
}


func (this *AuthenticationTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
}


func NewAuthenticationTag() *AuthenticationTag {
    this := &AuthenticationTag{}
    this.CustomTagSupport = *NewCustomTagSupport()
    this.initFields()
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        this.UserName = stringUtil.EMPTY_STRING
        this.Password = stringUtil.EMPTY_STRING
    return this
}


func (this *AuthenticationTag) SetCommand(command string) {
        this.Command = command
}


func (this *AuthenticationTag) SetUserName(userName string) {
        this.UserName = userName
}


func (this *AuthenticationTag) SetPassword(value string) {
        this.Password = value
}


func (this *AuthenticationTag) SetNewPassword(value string) {
        this.NewPassword = value
}


func (this *AuthenticationTag) SetRoles(values IBasicArrayList) {
        this.Roles = values
}


// throws LicensingException
func (this *AuthenticationTag) ChangePassword() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewAuthenticationRequestHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("changePassword", nil); _ = method
            var result bool = any(method.Invoke(anyType, nil)).(bool); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *AuthenticationTag) NewPassword() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewAuthenticationRequestHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("newPassword", nil); _ = method
            var result bool = any(method.Invoke(anyType, nil)).(bool); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *AuthenticationTag) InvalidateSession() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewAuthenticationHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("invalidateSession", nil); _ = method
            var resultInteger int = any(method.Invoke(anyType, nil)).(int); _ = resultInteger
            tryControl1, tryResult1 = 1, int(resultInteger); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *AuthenticationTag) IsSessionOld() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewAuthenticationHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var helperClass reflect.Type = reflect.TypeOf(anyType); _ = helperClass
            var method *Method = helperClass.GetMethod("isSessionOld", nil); _ = method
            var resultBoolean bool = any(method.Invoke(anyType, nil)).(bool); _ = resultBoolean
            tryControl1, tryResult1 = 1, resultBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *AuthenticationTag) IsRoleValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewAuthenticationRequestHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var methodParams []reflect.Type = []reflect.Type{reflect.TypeOf(this.UserName), reflect.TypeOf(this.Password), reflect.TypeOf(this.Roles)}; _ = methodParams
            var method *Method = addressHelperClass.GetMethod("isRoleValid", methodParams); _ = method
            var methodArgs []any = []any{this.UserName, this.Password, this.Roles}; _ = methodArgs
            var resultBoolean bool = any(method.Invoke(anyType, methodArgs)).(bool); _ = resultBoolean
            tryControl1, tryResult1 = 1, resultBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *AuthenticationTag) ValidRole() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewAuthenticationHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("validRole", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *AuthenticationTag) InvalidRole() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewAuthenticationHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("invalidRole", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *AuthenticationTag) IsAuthenticationSessionValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewAuthenticationHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var methodParams []reflect.Type = []reflect.Type{reflect.TypeOf(this.Roles)}; _ = methodParams
            var method *Method = addressHelperClass.GetMethod("isAuthenticationSessionValid", methodParams); _ = method
            var methodArgs []any = []any{this.Roles}; _ = methodArgs
            var resultBoolean bool = any(method.Invoke(anyType, methodArgs)).(bool); _ = resultBoolean
            tryControl1, tryResult1 = 1, resultBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *AuthenticationTag) ProcessIfNewLogin() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewAuthenticationHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var helperClass reflect.Type = reflect.TypeOf(anyType); _ = helperClass
            var methodParams []reflect.Type = []reflect.Type{reflect.TypeOf(this.UserName), reflect.TypeOf(this.Password)}; _ = methodParams
            var method *Method = helperClass.GetMethod("processIfNewLogin", methodParams); _ = method
            var methodArgs []any = []any{this.UserName, this.Password}; _ = methodArgs
            var resultBoolean bool = any(method.Invoke(anyType, methodArgs)).(bool); _ = resultBoolean
            tryControl1, tryResult1 = 1, resultBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *AuthenticationTag) ProcessInvalidation() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewAuthenticationHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("processInvalidation", nil); _ = method
            var resultBoolean bool = any(method.Invoke(anyType, nil)).(bool); _ = resultBoolean
            tryControl1, tryResult1 = 1, resultBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *AuthenticationTag) DoStartTag() int {
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
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                var request *HttpServletRequest = any(this.PageContext.GetRequest()).(*HttpServletRequest); _ = request
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append(" Request URI: ")
                stringBuffer.Append(request.GetRequestURI())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.TagStrings.DO_START_TAG)
            }
            if this.Command != "" {
                this.PropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
                if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.NEWPASSWORD) == 0 {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                        this.LogUtil.PutF("New Password", this, this.TagStrings.DO_START_TAG)
                    }
                    if !this.NewPassword() {
                        this.PageContext.GetOut().Print("New Password Request Failed.<p/>")
                        tryControl1, tryResult1 = 1, CustomTagSupport_EVAL_BODY_INCLUDE; return
                    } else {
                        this.PageContext.GetOut().Print("New Password Was Sent.<p/>")
                        tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
                    }
                } else if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.CHANGEPASSWORD) == 0 {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                        this.LogUtil.PutF("Change Password", this, this.TagStrings.DO_START_TAG)
                    }
                    if !this.ChangePassword() {
                        this.PageContext.GetOut().Print("Password Change Attempt Failed.<p/>")
                        tryControl1, tryResult1 = 1, CustomTagSupport_EVAL_BODY_INCLUDE; return
                    } else {
                        this.PageContext.GetOut().Print("Password Changed Successfully.<p/>")
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
                        if this.UserName != "" && strings.Compare(this.UserName, stringUtil.EMPTY_STRING) != 0 && this.Password != "" && strings.Compare(this.Password, stringUtil.EMPTY_STRING) != 0 {
                            this.PageContext.GetOut().Print("Sorry your username and/or password is invalid.<p/>")
                        }
                    }
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Major authentication error - userName: ").Append(this.UserName).Append(" Password: ").Append(this.Password).Append(" Command: ").Append(this.Command)), this, this.TagStrings.DO_START_TAG)
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

