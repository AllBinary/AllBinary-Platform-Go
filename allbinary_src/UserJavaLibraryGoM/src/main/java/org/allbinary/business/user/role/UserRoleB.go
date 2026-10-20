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

type UserRoleB struct {
    UserRole
}

var UserRoleB_RoleVector *BasicArrayList = NewBasicArrayListD()


func NewUserRoleB(aBasicUserRole *BasicUserRole, aUserFactoryInterface UserFactoryInterface) *UserRoleB {
    this := &UserRoleB{}
        this.UserRole = *NewUserRole(aBasicUserRole, aUserFactoryInterface)
        UserRoleB_RoleVector.Add(this)
    return this
}


// throws Exception
// synchronized
func UserRoleB_GetRole(role string) *UserRole {
        UserRoleFactory_GetInstance()
        if !StringValidationUtil_GetInstance().IsEmpty(role) {
            var roleVector *BasicArrayList = UserRoleB_GetAll(); _ = roleVector
            var size int = roleVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var userRole *UserRole = any(roleVector.Get(index)).(*UserRole); _ = userRole
                if strings.Compare(role, fmt.Sprint(userRole)) == 0 {
                    return userRole
                }
            }
        }
        panic(errors.New(fmt.Sprint("Unable to get role - Unknown Role = " + role)))
}


func UserRoleB_GetAll() *BasicArrayList {
        UserRoleFactory_GetInstance()
        return UserRoleB_RoleVector
}

