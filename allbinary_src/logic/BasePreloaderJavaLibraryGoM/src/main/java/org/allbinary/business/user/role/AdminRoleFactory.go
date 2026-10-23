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

type AdminRoleFactory struct {
    AdminRoles IBasicArrayList
    StoreAdminRoles IBasicArrayList
    StoreManagementRoles IBasicArrayList
    ProductManagementRoles IBasicArrayList
    ReviewManagementRoles IBasicArrayList
    CustomerManagementRoles IBasicArrayList
    WorkflowManagementRoles IBasicArrayList
    OrderManagementRoles IBasicArrayList
    AdjusterManagementRoles IBasicArrayList
    ShippingManagementRoles IBasicArrayList
    WebManagementRoles IBasicArrayList
}

var AdminRoleFactory_Instance *AdminRoleFactory = NewAdminRoleFactory()

func (this *AdminRoleFactory) initFields() {
    this.AdminRoles = NewBasicArrayListD()
    this.StoreAdminRoles = NewBasicArrayListD()
    this.StoreManagementRoles = NewBasicArrayListD()
    this.ProductManagementRoles = NewBasicArrayListD()
    this.ReviewManagementRoles = NewBasicArrayListD()
    this.CustomerManagementRoles = NewBasicArrayListD()
    this.WorkflowManagementRoles = NewBasicArrayListD()
    this.OrderManagementRoles = NewBasicArrayListD()
    this.AdjusterManagementRoles = NewBasicArrayListD()
    this.ShippingManagementRoles = NewBasicArrayListD()
    this.WebManagementRoles = NewBasicArrayListD()
}


func AdminRoleFactory_GetInstance() *AdminRoleFactory {
        return AdminRoleFactory_Instance
}


func NewAdminRoleFactory() *AdminRoleFactory {
    this := &AdminRoleFactory{}
    this.initFields()
        var basicUserRoleFactory *BasicUserRoleFactory = BasicUserRoleFactory_GetInstance(); _ = basicUserRoleFactory
        this.AdminRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.StoreAdminRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.StoreAdminRoles.Add(basicUserRoleFactory.STOREMANAGER)
        this.StoreManagementRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.StoreManagementRoles.Add(basicUserRoleFactory.STOREMANAGER)
        this.StoreManagementRoles.Add(basicUserRoleFactory.CUSTOMERMANAGER)
        this.StoreManagementRoles.Add(basicUserRoleFactory.ADJUSTER)
        this.StoreManagementRoles.Add(basicUserRoleFactory.PRODUCTMANAGER)
        this.StoreManagementRoles.Add(basicUserRoleFactory.REVIEWER)
        this.StoreManagementRoles.Add(basicUserRoleFactory.SHIPPING)
        this.ProductManagementRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.ProductManagementRoles.Add(basicUserRoleFactory.STOREMANAGER)
        this.ProductManagementRoles.Add(basicUserRoleFactory.PRODUCTMANAGER)
        this.ReviewManagementRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.ReviewManagementRoles.Add(basicUserRoleFactory.STOREMANAGER)
        this.ReviewManagementRoles.Add(basicUserRoleFactory.REVIEWER)
        this.CustomerManagementRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.CustomerManagementRoles.Add(basicUserRoleFactory.STOREMANAGER)
        this.CustomerManagementRoles.Add(basicUserRoleFactory.CUSTOMERMANAGER)
        this.WorkflowManagementRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.WorkflowManagementRoles.Add(basicUserRoleFactory.STOREMANAGER)
        this.WorkflowManagementRoles.Add(basicUserRoleFactory.WORKFLOWEDITOR)
        this.OrderManagementRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.OrderManagementRoles.Add(basicUserRoleFactory.STOREMANAGER)
        this.AdjusterManagementRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.AdjusterManagementRoles.Add(basicUserRoleFactory.STOREMANAGER)
        this.AdjusterManagementRoles.Add(basicUserRoleFactory.ADJUSTER)
        this.ShippingManagementRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.ShippingManagementRoles.Add(basicUserRoleFactory.STOREMANAGER)
        this.ShippingManagementRoles.Add(basicUserRoleFactory.SHIPPING)
        this.WebManagementRoles.Add(basicUserRoleFactory.ADMINISTRATOR)
        this.WebManagementRoles.Add(basicUserRoleFactory.STOREMANAGER)
        this.WebManagementRoles.Add(basicUserRoleFactory.WEBMANAGER)
        this.WebManagementRoles.Add(basicUserRoleFactory.WORKFLOWEDITOR)
        this.WebManagementRoles.Add(basicUserRoleFactory.VIEWEDITOR)
    return this
}


func (this *AdminRoleFactory) GetProductManagementRoles() IBasicArrayList {
        return this.ProductManagementRoles
}


func (this *AdminRoleFactory) GetAdminRoles() IBasicArrayList {
        return this.AdminRoles
}


func (this *AdminRoleFactory) GetStoreAdminRoles() IBasicArrayList {
        return this.StoreAdminRoles
}


func (this *AdminRoleFactory) GetStoreManagementRoles() IBasicArrayList {
        return this.StoreManagementRoles
}


func (this *AdminRoleFactory) GetReviewManagementRoles() IBasicArrayList {
        return this.ReviewManagementRoles
}


func (this *AdminRoleFactory) GetCustomerManagementRoles() IBasicArrayList {
        return this.CustomerManagementRoles
}


func (this *AdminRoleFactory) GetWorkflowManagementRoles() IBasicArrayList {
        return this.WorkflowManagementRoles
}


func (this *AdminRoleFactory) GetWebManagementRoles() IBasicArrayList {
        return this.WebManagementRoles
}


func (this *AdminRoleFactory) GetShippingManagementRoles() IBasicArrayList {
        return this.ShippingManagementRoles
}


func (this *AdminRoleFactory) GetAdjusterManagementRoles() IBasicArrayList {
        return this.AdjusterManagementRoles
}


func (this *AdminRoleFactory) GetOrderManagementRoles() IBasicArrayList {
        return this.OrderManagementRoles
}

