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

type UserRole struct {
    BasicUserRole *BasicUserRole
    UserFactoryInterface UserFactoryInterface
}

// Implements *Serializable

// The Java class UserRole has subclasses so its variables are this interface
type IUserRole interface {
    UserRole_() *UserRole
    ToString() string
    GetBasicUserRole() *BasicUserRole
    GetUserFactory() UserFactoryInterface
    GetSessionTimeout() int64
    GetSessionInactivityTimeout() int64
}

func (this *UserRole) UserRole_() *UserRole {
    return this
}


func (this *UserRole) initFields() {
    
    
}


func NewUserRole(aBasicUserRole *BasicUserRole, aUserFactoryInterface UserFactoryInterface) *UserRole {
    this := &UserRole{}
    this.initFields()
        this.BasicUserRole = aBasicUserRole
        this.UserFactoryInterface = aUserFactoryInterface
    return this
}


func (this *UserRole) ToString() string {
        return fmt.Sprint(this.BasicUserRole)
}

func (this *UserRole) String() string {
    return this.ToString()
}


func (this *UserRole) GetBasicUserRole() *BasicUserRole {
        return this.BasicUserRole
}


// throws Exception
func (this *UserRole) GetUserFactory() UserFactoryInterface {
        return this.UserFactoryInterface
}


func (this *UserRole) GetSessionTimeout() int64 {
        return this.BasicUserRole.GetSessionTimeout()
}


func (this *UserRole) GetSessionInactivityTimeout() int64 {
        return this.BasicUserRole.GetSessionInactivityTimeout()
}

