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

type UserRoleFactory struct {
    BasicUserRoleFactory *BasicUserRoleFactory
    INSTALLER IUserRole
    ADMINISTRATOR IUserRole
    CUSTOMER IUserRole
    SUBSCRIBERCUSTOMER IUserRole
    WHOLESALECUSTOMER IUserRole
    STOREMANAGER IUserRole
    PRODUCTMANAGER IUserRole
    CUSTOMERMANAGER IUserRole
    PAYMENTMANAGER IUserRole
    ORDERMANAGER IUserRole
    WEBMANAGER IUserRole
    REVIEWER IUserRole
    SHIPPING IUserRole
    ADJUSTER IUserRole
    WORKFLOWEDITOR IUserRole
    WIZARD IUserRole
    VIEWEDITOR IUserRole
}

var UserRoleFactory_Instance *UserRoleFactory = NewUserRoleFactory()

func (this *UserRoleFactory) initFields() {
    this.BasicUserRoleFactory = BasicUserRoleFactory_GetInstance()
    this.INSTALLER = NewUserRoleB(this.BasicUserRoleFactory.INSTALLER, nil)
    this.ADMINISTRATOR = NewUserRoleB(this.BasicUserRoleFactory.ADMINISTRATOR, NewAdminUserFactory())
    this.CUSTOMER = NewUserRoleB(this.BasicUserRoleFactory.CUSTOMER, NewUserFactory())
    this.SUBSCRIBERCUSTOMER = NewUserRoleB(this.BasicUserRoleFactory.SUBSCRIBERCUSTOMER, NewSubscriberUserFactory())
    this.WHOLESALECUSTOMER = NewUserRoleB(this.BasicUserRoleFactory.WHOLESALECUSTOMER, NewWholesaleUserFactory())
    this.STOREMANAGER = NewUserRoleB(this.BasicUserRoleFactory.STOREMANAGER, NewStoreAdminUserFactory())
    this.PRODUCTMANAGER = NewUserRoleB(this.BasicUserRoleFactory.PRODUCTMANAGER, NewProductStoreAdminUserFactory())
    this.CUSTOMERMANAGER = NewUserRoleB(this.BasicUserRoleFactory.CUSTOMERMANAGER, NewCustomerStoreAdminUserFactory())
    this.PAYMENTMANAGER = NewUserRoleB(this.BasicUserRoleFactory.PAYMENTMANAGER, NewPaymentStoreAdminUserFactory())
    this.ORDERMANAGER = NewUserRoleB(this.BasicUserRoleFactory.ORDERMANAGER, NewOrderStoreAdminUserFactory())
    this.WEBMANAGER = NewUserRoleB(this.BasicUserRoleFactory.WEBMANAGER, NewWebStoreAdminUserFactory())
    this.REVIEWER = NewUserRoleB(this.BasicUserRoleFactory.REVIEWER, NewReviewOrderStoreAdminUserFactory())
    this.SHIPPING = NewUserRoleB(this.BasicUserRoleFactory.SHIPPING, NewShipOrderStoreAdminUserFactory())
    this.ADJUSTER = NewUserRoleB(this.BasicUserRoleFactory.ADJUSTER, NewAdjustOrderStoreAdminUserFactory())
    this.WORKFLOWEDITOR = NewUserRoleB(this.BasicUserRoleFactory.WORKFLOWEDITOR, NewWorkflowEditorWebStoreAdminUserFactory())
    this.WIZARD = NewUserRoleB(this.BasicUserRoleFactory.WIZARD, NewWizardWebStoreAdminUserFactory())
    this.VIEWEDITOR = NewUserRoleB(this.BasicUserRoleFactory.VIEWEDITOR, NewViewEditorWebStoreAdminUserFactory())
}

func NewUserRoleFactory() *UserRoleFactory {
    this := &UserRoleFactory{}
    this.initFields()
    return this
}


func UserRoleFactory_GetInstance() *UserRoleFactory {
        return UserRoleFactory_Instance
}

