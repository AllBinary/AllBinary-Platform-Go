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

type StoreFrontUsersStatistics struct {
    TotalNumberOfUsers int64
    TotalUsersByRoleHashMap *HashMap
}

// Implements StoreFrontUsersStatisticsInterface


func (this *StoreFrontUsersStatistics) initFields() {
    
    
}


func NewStoreFrontUsersStatistics(storeFrontInterface StoreFrontInterface) *StoreFrontUsersStatistics {
    this := &StoreFrontUsersStatistics{}
    this.initFields()
        this.TotalUsersByRoleHashMap = StdUtil_GetInstance().CreateHashMap()
        var userEntityInterface UserEntityInterface = UserEntityFactory_GetInstance(); _ = userEntityInterface
        var userVector IBasicArrayList = userEntityInterface.GetCustomers(); _ = userVector
        this.TotalNumberOfUsers = int64(userVector.Size())
        var size int = userVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var userInterface UserInterface = any(userVector.Get(index)).(UserInterface); _ = userInterface
            var nextUserRole IUserRole = userInterface.GetRole(); _ = nextUserRole
            var currentNumberOfUsersForRole int64 = this.GetNewTotal(nextUserRole); _ = currentNumberOfUsersForRole
            this.TotalUsersByRoleHashMap.Put(nextUserRole, currentNumberOfUsersForRole)
        }
    return this
}


func (this *StoreFrontUsersStatistics) GetNewTotal(userRole IUserRole) int64 {
        var numberOfUsersForRoleLong int64 = any(this.TotalUsersByRoleHashMap.Get(userRole)).(int64); _ = numberOfUsersForRoleLong
        if numberOfUsersForRoleLong == 0 {
            numberOfUsersForRoleLong = int64(0)
        }
        return int64(int64(numberOfUsersForRoleLong) + 1)
}


func (this *StoreFrontUsersStatistics) GetNumberOfUsers() int64 {
        return this.TotalNumberOfUsers
}


func (this *StoreFrontUsersStatistics) GetNumberOfUsersByRole(role string) int64 {
        var totalForRole int64 = any(this.TotalUsersByRoleHashMap.Get(role)).(int64); _ = totalForRole
        return totalForRole
}


func (this *StoreFrontUsersStatistics) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(StoreFrontUsersStatisticsData_GetInstance().NUMBEROFUSERS, fmt.Sprint(this.GetNumberOfUsers()))
        var setOfUserRoles Set = this.TotalUsersByRoleHashMap.KeySet(); _ = setOfUserRoles
        var nextUserRole IUserRole; _ = nextUserRole
        var userRoleArray []any = setOfUserRoles.ToArray(); _ = userRoleArray
        var size int = len(userRoleArray); _ = size
        for index := int(0); index < size; index++ {
            nextUserRole = any(userRoleArray[index]).(IUserRole)
            var totalForRole int64 = any(this.TotalUsersByRoleHashMap.Get(nextUserRole)).(int64); _ = totalForRole
            hashMap.Put(fmt.Sprint(nextUserRole), fmt.Sprint(totalForRole))
        }
        return hashMap
}


func (this *StoreFrontUsersStatistics) ToVector() IBasicArrayList {
        return nil
}


func (this *StoreFrontUsersStatistics) GetKey() any {
        return nil
}

