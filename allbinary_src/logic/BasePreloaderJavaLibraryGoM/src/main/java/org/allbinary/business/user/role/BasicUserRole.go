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

type BasicUserRole struct {
    M_int_RoleId int
    Role string
    DisplayValue string
    SessionTimeout int64
    SessionInactivityTimeout int64
}

// Implements *Serializable

var BasicUserRole_RoleVector IBasicArrayList = NewBasicArrayListD()

func (this *BasicUserRole) initFields() {
    
    
    
    
    
}




// Java overloads of new BasicUserRole
func NewBasicUserRole(args ...any) *BasicUserRole {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[string](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[int64](args[2], pass) && Java_IsArg[int64](args[3], pass) {
            return NewBasicUserRole_1(Java_Arg[string](args[0]), Java_Arg[int](args[1]), Java_Arg[int64](args[2]), Java_Arg[int64](args[3]))
        }
        if len(args) == 5 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[int](args[2], pass) && Java_IsArg[int64](args[3], pass) && Java_IsArg[int64](args[4], pass) {
            return NewBasicUserRole_2(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[int](args[2]), Java_Arg[int64](args[3]), Java_Arg[int64](args[4]))
        }
    }
    panic("No Java overload of new BasicUserRole for the arguments")
}
func NewBasicUserRole_1(aRole string, a_int_RoleId int, aSessionTimeout int64, aSessionInactivityTimeout int64) *BasicUserRole {
    this := &BasicUserRole{}
        *this = *NewBasicUserRole(aRole, aRole, a_int_RoleId, aSessionTimeout, aSessionInactivityTimeout)
    return this
}


func NewBasicUserRole_2(aRole string, aDisplayValue string, a_int_RoleId int, aSessionTimeout int64, aSessionInactivityTimeout int64) *BasicUserRole {
    this := &BasicUserRole{}
    this.initFields()
        this.M_int_RoleId = a_int_RoleId
        this.Role = aRole
        this.DisplayValue = aDisplayValue
        this.SessionTimeout = aSessionTimeout
        this.SessionInactivityTimeout = aSessionInactivityTimeout
        BasicUserRole_RoleVector.Add(this)
    return this
}


func (this *BasicUserRole) GetRoleId() int {
        return this.M_int_RoleId
}


func (this *BasicUserRole) GetRole() string {
        return this.Role
}


func (this *BasicUserRole) GetSessionTimeout() int64 {
        return this.SessionTimeout
}


func (this *BasicUserRole) GetSessionInactivityTimeout() int64 {
        return this.SessionInactivityTimeout
}


func (this *BasicUserRole) Equals(basicUserRole *BasicUserRole) bool {
        if this.GetRoleId() == basicUserRole.GetRoleId() {
            return true
        }
        return false
}


func (this *BasicUserRole) ToString() string {
        return this.GetRole()
}

func (this *BasicUserRole) String() string {
    return this.ToString()
}


// throws Exception
// synchronized
func BasicUserRole_GetRole(role string) *BasicUserRole {
        var roleVector IBasicArrayList = BasicUserRole_GetAll(); _ = roleVector
        var size int = roleVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var userRole *BasicUserRole = any(roleVector.Get(index)).(*BasicUserRole); _ = userRole
            if strings.Compare(role, fmt.Sprint(userRole)) == 0 {
                return userRole
            }
        }
        panic(errors.New(fmt.Sprint("Unable to get role - Unknown Role")))
}


func BasicUserRole_GetAll() IBasicArrayList {
        return BasicUserRole_RoleVector
}

