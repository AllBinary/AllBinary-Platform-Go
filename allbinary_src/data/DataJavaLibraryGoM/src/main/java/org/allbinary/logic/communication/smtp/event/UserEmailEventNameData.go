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

type UserEmailEventNameData struct {
    EventNameId int
    EventHandlerName string
}

var UserEmailEventNameData_UserNameEvenNameHashMap *HashMap = StdUtil_GetInstance().CreateHashMap()
var UserEmailEventNameData_NextId int = 0
var UserEmailEventNameData_INSTALLER *UserEmailEventNameData = NewUserEmailEventNameData("Installer")
var UserEmailEventNameData_EXCEPTION *UserEmailEventNameData = NewUserEmailEventNameData("Exception")
var UserEmailEventNameData_LOGGEDEXCEPTION *UserEmailEventNameData = NewUserEmailEventNameData("Logged Exception")
var UserEmailEventNameData_ERROR *UserEmailEventNameData = NewUserEmailEventNameData("Error")
var UserEmailEventNameData_OUTOFMEMORY *UserEmailEventNameData = NewUserEmailEventNameData("Out Of Memory Error")
var UserEmailEventNameData_ADMINGENERATINGSTATICPAGES *UserEmailEventNameData = NewUserEmailEventNameData("Admin Generating Static Pages")
var UserEmailEventNameData_STOREERROROCCURRED *UserEmailEventNameData = NewUserEmailEventNameData("Store Error Occurred")
var UserEmailEventNameData_STOREOUTOFMEMORY *UserEmailEventNameData = NewUserEmailEventNameData("Store Out Of Memory Error")
var UserEmailEventNameData_STORECREATED *UserEmailEventNameData = NewUserEmailEventNameData("Store Created")
var UserEmailEventNameData_STOREOPENED *UserEmailEventNameData = NewUserEmailEventNameData("Store Opened")
var UserEmailEventNameData_STORECLOSED *UserEmailEventNameData = NewUserEmailEventNameData("Store Closed")
var UserEmailEventNameData_STOREGENERATINGSTATICPAGES *UserEmailEventNameData = NewUserEmailEventNameData("Store Generating Static Pages")
var UserEmailEventNameData_PRODUCTISOUT *UserEmailEventNameData = NewUserEmailEventNameData("Product Is Out")
var UserEmailEventNameData_PRODUCTISAVAILABLE *UserEmailEventNameData = NewUserEmailEventNameData("Product Is Available")
var UserEmailEventNameData_PRODUCTISOLD *UserEmailEventNameData = NewUserEmailEventNameData("Product Is Old")
var UserEmailEventNameData_PRODUCTISONSALE *UserEmailEventNameData = NewUserEmailEventNameData("Product Is On Sale")
var UserEmailEventNameData_PRODUCTSALEISOVER *UserEmailEventNameData = NewUserEmailEventNameData("Product Sale Is Over")
var UserEmailEventNameData_ORDERPREPROCESSING *UserEmailEventNameData = NewUserEmailEventNameData(OrderHistoryData_PREPROCESSING)
var UserEmailEventNameData_ORDERPROCESSING *UserEmailEventNameData = NewUserEmailEventNameData(OrderHistoryData_PROCESSING)
var UserEmailEventNameData_ORDERCANCELLED *UserEmailEventNameData = NewUserEmailEventNameData(OrderHistoryData_CANCELLED)
var UserEmailEventNameData_ORDERPARTIALLYSHIPPED *UserEmailEventNameData = NewUserEmailEventNameData(OrderHistoryData_PARTIALLYSHIPPED)
var UserEmailEventNameData_ORDERSHIPPED *UserEmailEventNameData = NewUserEmailEventNameData(OrderHistoryData_SHIPPED)
var UserEmailEventNameData_QUOTEREQUEST *UserEmailEventNameData = NewUserEmailEventNameData("Quote Request")
var UserEmailEventNameData_NEWPASSWORD *UserEmailEventNameData = NewUserEmailEventNameData("New Password")

func (this *UserEmailEventNameData) initFields() {
    
    
}


func NewUserEmailEventNameData(eventHandlerName string) *UserEmailEventNameData {
    this := &UserEmailEventNameData{}
    this.initFields()
        this.EventHandlerName = eventHandlerName
        this.InitNextId()
        UserEmailEventNameData_UserNameEvenNameHashMap.Put(this.EventHandlerName, this)
    return this
}


// synchronized
func (this *UserEmailEventNameData) InitNextId() {
        this.EventNameId = UserEmailEventNameData_NextId
        UserEmailEventNameData_NextId++
}


func (this *UserEmailEventNameData) IsEvent(userEmailEventNameString string) bool {
        return UserEmailEventNameData_UserNameEvenNameHashMap.ContainsKey(userEmailEventNameString)
}


// throws Exception
// synchronized
func UserEmailEventNameData_GetInstance(userEmailEventNameString string) *UserEmailEventNameData {
        var userEmailEventNameData *UserEmailEventNameData = any(UserEmailEventNameData_UserNameEvenNameHashMap.Get(userEmailEventNameString)).(*UserEmailEventNameData); _ = userEmailEventNameData
        if userEmailEventNameData == nil {
            panic(errors.New(fmt.Sprint()))
        } else {
            return userEmailEventNameData
        }
}


func (this *UserEmailEventNameData) ToString() string {
        return this.EventHandlerName
}

func (this *UserEmailEventNameData) String() string {
    return this.ToString()
}

