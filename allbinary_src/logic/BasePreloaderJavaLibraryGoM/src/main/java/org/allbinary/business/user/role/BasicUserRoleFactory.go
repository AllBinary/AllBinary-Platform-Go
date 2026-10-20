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

type BasicUserRoleFactory struct {
    INACTIVITYTIMEOUT int64
    CUSTOMERTIMEOUT int64
    CUSTOMERMAXSESSIONTIME int64
    MAXSESSIONTIME int64
    INSTALLER *BasicUserRole
    ADMINISTRATOR *BasicUserRole
    CUSTOMER *BasicUserRole
    SUBSCRIBERCUSTOMER *BasicUserRole
    WHOLESALECUSTOMER *BasicUserRole
    STOREMANAGER *BasicUserRole
    PRODUCTMANAGER *BasicUserRole
    CUSTOMERMANAGER *BasicUserRole
    PAYMENTMANAGER *BasicUserRole
    ORDERMANAGER *BasicUserRole
    WEBMANAGER *BasicUserRole
    REVIEWER *BasicUserRole
    SHIPPING *BasicUserRole
    ADJUSTER *BasicUserRole
    WORKFLOWEDITOR *BasicUserRole
    WIZARD *BasicUserRole
    VIEWEDITOR *BasicUserRole
}

var BasicUserRoleFactory_Instance *BasicUserRoleFactory = NewBasicUserRoleFactory()

func (this *BasicUserRoleFactory) initFields() {
    this.INACTIVITYTIMEOUT = 360000000
    this.CUSTOMERTIMEOUT = 360000000
    this.CUSTOMERMAXSESSIONTIME = 360000000
    this.MAXSESSIONTIME = 360000000
    this.INSTALLER = NewBasicUserRole("Installer", 16, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.ADMINISTRATOR = NewBasicUserRole("Administrator", 0, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.CUSTOMER = NewBasicUserRole("CustomerUser", "Customer User", 1, this.CUSTOMERTIMEOUT, this.CUSTOMERMAXSESSIONTIME)
    this.SUBSCRIBERCUSTOMER = NewBasicUserRole("SubscriberUser", "Subscriber User", 2, this.CUSTOMERTIMEOUT, this.CUSTOMERMAXSESSIONTIME)
    this.WHOLESALECUSTOMER = NewBasicUserRole("WholesaleUser", "Wholesale User", 3, this.CUSTOMERTIMEOUT, this.CUSTOMERMAXSESSIONTIME)
    this.STOREMANAGER = NewBasicUserRole("StoreManager", "Store Manager", 4, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.PRODUCTMANAGER = NewBasicUserRole("ProductManager", "Product Manager", 5, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.CUSTOMERMANAGER = NewBasicUserRole("UserManager", "User Manager", 6, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.PAYMENTMANAGER = NewBasicUserRole("PaymentManager", "Payment Manager", 7, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.ORDERMANAGER = NewBasicUserRole("OrderManager", "Order Manager", 8, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.WEBMANAGER = NewBasicUserRole("WebManager", "Web Manager", 9, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.REVIEWER = NewBasicUserRole("Reviewer", 10, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.SHIPPING = NewBasicUserRole("Shipping", 11, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.ADJUSTER = NewBasicUserRole("Adjuster", 12, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.WORKFLOWEDITOR = NewBasicUserRole("WorkFlowEditor", "WorkFlow Editor", 13, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.WIZARD = NewBasicUserRole("WizardUser", "Wizard User", 14, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
    this.VIEWEDITOR = NewBasicUserRole("ViewEditor", "View Editor", 15, this.MAXSESSIONTIME, this.INACTIVITYTIMEOUT)
}

func NewBasicUserRoleFactory() *BasicUserRoleFactory {
    this := &BasicUserRoleFactory{}
    this.initFields()
    return this
}


func BasicUserRoleFactory_GetInstance() *BasicUserRoleFactory {
        return BasicUserRoleFactory_Instance
}
