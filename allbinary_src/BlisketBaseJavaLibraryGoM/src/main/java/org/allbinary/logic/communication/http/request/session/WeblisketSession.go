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

type WeblisketSession struct {
    Session *HttpSession
}

// Implements WeblisketSessionInterface


func (this *WeblisketSession) initFields() {
    
}




// Java overloads of new WeblisketSession
func NewWeblisketSession(args ...any) *WeblisketSession {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*PageContext](args[1], pass) {
            return NewWeblisketSession_1(Java_Arg[*HashMap](args[0]), Java_Arg[*PageContext](args[1]))
        }
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*HttpServletRequest](args[1], pass) {
            return NewWeblisketSession_2(Java_Arg[*HashMap](args[0]), Java_Arg[*HttpServletRequest](args[1]))
        }
    }
    panic("No Java overload of new WeblisketSession for the arguments")
}
func NewWeblisketSession_1(hashMap *HashMap, pageContext *PageContext) *WeblisketSession {
    this := &WeblisketSession{}
        *this = *NewWeblisketSession(hashMap, any(pageContext.GetRequest()).(*HttpServletRequest))
    return this
}


func NewWeblisketSession_2(hashMap *HashMap, httpServletRequest *HttpServletRequest) *WeblisketSession {
    this := &WeblisketSession{}
    this.initFields()
        var propertiesHashMap *HashMap = hashMap; _ = propertiesHashMap
        var request *HttpServletRequest = httpServletRequest; _ = request
        this.Session = request.GetSession(true)
        if this.GetStoreName() == nil {
            this.SetStoreName(any(propertiesHashMap.Get(StoreFrontData_GetInstance().NAME)).(string))
        }
    return this
}


func (this *WeblisketSession) GetId() string {
        return this.Session.GetId()
}


func (this *WeblisketSession) SetPassword(password string) {
        this.Session.SetAttribute(UserData_PASSWORD, password)
}


func (this *WeblisketSession) SetUserName(userName string) {
        this.Session.SetAttribute(UserData_USERNAME, userName)
}


func (this *WeblisketSession) SetPaymentMethod(value string) {
        this.Session.SetAttribute(fmt.Sprint(PaymentGatewayData_NAME), value)
}


func (this *WeblisketSession) SetAuthenticated_1() {
        this.Session.SetAttribute(WeblisketSessionData_AUTHENTICATED, BooleanFactory_GetInstance().TRUE_STRING)
}

// Java overloads of WeblisketSession.setAuthenticated
func (this *WeblisketSession) SetAuthenticated(args ...any) {
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
    panic("No Java overload of WeblisketSession.setAuthenticated for the arguments")
}


func (this *WeblisketSession) SetRole(aRole IUserRole) {
        this.Session.SetAttribute(UserRoleData_NAME, aRole)
}


func (this *WeblisketSession) SetAttempts(value int) {
        this.Session.SetAttribute(WeblisketSessionData_ATTEMPTS, value)
}


func (this *WeblisketSession) SetTimeout(value string) {
        this.Session.SetAttribute(WeblisketSessionData_TIMEOUT, value)
}


func (this *WeblisketSession) SetStoreName(value string) {
        this.Session.SetAttribute(StoreFrontData_GetInstance().NAME, value)
}


func (this *WeblisketSession) GetWebAppPath() string {
        return any(this.Session.GetAttribute(org_Allbinary.Globals.GLOBALS2.WEBAPPPATH)).(string)
}


func (this *WeblisketSession) GetUserName() string {
        return any(this.Session.GetAttribute(UserData_USERNAME)).(string)
}


func (this *WeblisketSession) GetPassword() string {
        return any(this.Session.GetAttribute(UserData_PASSWORD)).(string)
}


func (this *WeblisketSession) GetPaymentMethod() string {
        return any(this.Session.GetAttribute(fmt.Sprint(PaymentGatewayData_NAME))).(string)
}


func (this *WeblisketSession) GetAuthentication() string {
        return any(this.Session.GetAttribute(WeblisketSessionData_AUTHENTICATED)).(string)
}


func (this *WeblisketSession) GetAttempts() int {
        return any(this.Session.GetAttribute(WeblisketSessionData_ATTEMPTS)).(int)
}


func (this *WeblisketSession) GetTimeout() string {
        return any(this.Session.GetAttribute(WeblisketSessionData_TIMEOUT)).(string)
}


func (this *WeblisketSession) GetStoreName() string {
        return any(this.Session.GetAttribute(StoreFrontData_GetInstance().NAME)).(string)
}


// throws Exception
func (this *WeblisketSession) GetRole() IUserRole {
        var userRole IUserRole = any(this.Session.GetAttribute(UserRoleData_NAME)).(IUserRole); _ = userRole
        if userRole != nil {
            return UserRoleB_GetRole(userRole.GetBasicUserRole().GetRole())
        } else {
            return nil
        }
}


// throws Exception
func (this *WeblisketSession) GetOrder() OrderInterface {
        var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetStoreName()); _ = storeFrontInterface
        var basketName string = storeFrontInterface.GetBasketName(); _ = basketName
        if !StringValidationUtil_GetInstance().IsEmpty(basketName) {
            var orderInterface OrderInterface = any(this.Session.GetAttribute(basketName)).(OrderInterface); _ = orderInterface
            if orderInterface == nil {
                orderInterface = any(NewOrder(NewBasket())).(OrderInterface)
            }
            this.Session.SetAttribute(basketName, orderInterface)
            return orderInterface
        } else {
            panic(errors.New(fmt.Sprint("Basket Name Not Found")))
        }
}


func (this *WeblisketSession) SetAuthenticated_2(value bool) {
        this.Session.SetAttribute(WeblisketSessionData_AUTHENTICATED, BooleanFactory_GetInstance().FALSE_STRING)
}


func (this *WeblisketSession) Clear() {
        this.Session.RemoveAttribute(WeblisketSessionData_AUTHENTICATED)
        this.Session.RemoveAttribute(WeblisketSessionData_TIMEOUT)
        this.Session.RemoveAttribute(UserRoleData_NAME)
        this.Session.RemoveAttribute(UserData_USERNAME)
}


func (this *WeblisketSession) GetCreationTime() int64 {
        return this.Session.GetCreationTime()
}


func (this *WeblisketSession) GetLastAccessedTime() int64 {
        return this.Session.GetLastAccessedTime()
}


func (this *WeblisketSession) RemoveBasket() {
        var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetStoreName()); _ = storeFrontInterface
        var basketName string = storeFrontInterface.GetBasketName(); _ = basketName
        if !StringValidationUtil_GetInstance().IsEmpty(basketName) {
            this.Session.RemoveAttribute(basketName)
        }
}

