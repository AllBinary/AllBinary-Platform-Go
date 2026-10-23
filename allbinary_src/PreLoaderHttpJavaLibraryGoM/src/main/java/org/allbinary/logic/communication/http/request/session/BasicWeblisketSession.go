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

type BasicWeblisketSession struct {
    Session *HttpSession
    PageContext *PageContext
}


func (this *BasicWeblisketSession) initFields() {
    
    
}


func NewBasicWeblisketSession(pageContext *PageContext) *BasicWeblisketSession {
    this := &BasicWeblisketSession{}
    this.initFields()
        this.PageContext = pageContext
        var request *HttpServletRequest = any(pageContext.GetRequest()).(*HttpServletRequest); _ = request
        this.Session = request.GetSession(true)
    return this
}


func (this *BasicWeblisketSession) SetRole(aRole *BasicUserRole) {
        this.Session.SetAttribute(BasicUserRoleData_NAME, aRole)
}


func (this *BasicWeblisketSession) GetRole() *BasicUserRole {
        return any(this.Session.GetAttribute(BasicUserRoleData_NAME)).(*BasicUserRole)
}


func (this *BasicWeblisketSession) SetPassword(password string) {
        this.Session.SetAttribute(UserData_PASSWORD, password)
}


func (this *BasicWeblisketSession) SetUserName(userName string) {
        this.Session.SetAttribute(UserData_USERNAME, userName)
}


func (this *BasicWeblisketSession) SetPaymentMethod(value string) {
        this.Session.SetAttribute(fmt.Sprint(PaymentGatewayData_NAME), value)
}


func (this *BasicWeblisketSession) SetAuthenticated_1() {
        this.Session.SetAttribute(WeblisketSessionData_AUTHENTICATED, BooleanFactory_GetInstance().TRUE_STRING)
}

// Java overloads of BasicWeblisketSession.setAuthenticated
func (this *BasicWeblisketSession) SetAuthenticated(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            this.SetAuthenticated_1()
            return
        }
        if len(args) == 1 && Java_IsArg[bool](args[0], pass) {
            this.SetAuthenticated_2(Java_Arg[bool](args[0]))
            return
        }
    }
    panic("No Java overload of BasicWeblisketSession.setAuthenticated for the arguments")
}


func (this *BasicWeblisketSession) SetAttempts(value int) {
        this.Session.SetAttribute(WeblisketSessionData_ATTEMPTS, value)
}


func (this *BasicWeblisketSession) SetTimeout(value string) {
        this.Session.SetAttribute(WeblisketSessionData_TIMEOUT, value)
}


func (this *BasicWeblisketSession) SetStoreName(value string) {
        this.Session.SetAttribute(StoreFrontData_GetInstance().NAME, value)
}


func (this *BasicWeblisketSession) GetWebAppPath() string {
        return any(this.Session.GetAttribute(org_Allbinary.Globals.GLOBALS2.WEBAPPPATH)).(string)
}


func (this *BasicWeblisketSession) GetUserName() string {
        return any(this.Session.GetAttribute(UserData_USERNAME)).(string)
}


func (this *BasicWeblisketSession) GetPassword() string {
        return any(this.Session.GetAttribute(UserData_PASSWORD)).(string)
}


func (this *BasicWeblisketSession) GetPaymentMethod() string {
        return any(this.Session.GetAttribute(fmt.Sprint(PaymentGatewayData_NAME))).(string)
}


func (this *BasicWeblisketSession) GetAuthentication() string {
        return any(this.Session.GetAttribute(WeblisketSessionData_AUTHENTICATED)).(string)
}


func (this *BasicWeblisketSession) GetAttempts() int {
        return any(this.Session.GetAttribute(WeblisketSessionData_ATTEMPTS)).(int)
}


func (this *BasicWeblisketSession) GetTimeout() string {
        return any(this.Session.GetAttribute(WeblisketSessionData_TIMEOUT)).(string)
}


func (this *BasicWeblisketSession) GetStoreName() string {
        return any(this.Session.GetAttribute(StoreFrontData_GetInstance().NAME)).(string)
}


func (this *BasicWeblisketSession) SetAuthenticated_2(value bool) {
        this.Session.SetAttribute(WeblisketSessionData_AUTHENTICATED, BooleanFactory_GetInstance().FALSE_STRING)
}


func (this *BasicWeblisketSession) Clear() {
        this.Session.RemoveAttribute(WeblisketSessionData_AUTHENTICATED)
        this.Session.RemoveAttribute(WeblisketSessionData_TIMEOUT)
        this.Session.RemoveAttribute(BasicUserRoleData_NAME)
        this.Session.RemoveAttribute(UserData_USERNAME)
}


func (this *BasicWeblisketSession) GetCreationTime() int64 {
        return this.Session.GetCreationTime()
}


func (this *BasicWeblisketSession) GetLastAccessedTime() int64 {
        var lastAccessedTime int64 = this.Session.GetLastAccessedTime(); _ = lastAccessedTime
        return lastAccessedTime
}

