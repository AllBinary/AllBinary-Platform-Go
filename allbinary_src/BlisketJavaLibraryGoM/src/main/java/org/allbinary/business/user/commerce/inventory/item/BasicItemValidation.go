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

type BasicItemValidation struct {
    Validation
    LogUtil *LogUtil
    MAXWEIGHT int
    MAXPRICE int
    ItemInterface ItemInterface
}


func (this *BasicItemValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.MAXWEIGHT = 14
    this.MAXPRICE = 14
    
}


func NewBasicItemValidation(itemInterface ItemInterface) *BasicItemValidation {
    this := &BasicItemValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.ItemInterface = itemInterface
    return this
}


func (this *BasicItemValidation) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to validate form", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if this.ItemInterface.GetId() == nil || len(this.ItemInterface.GetId()) < 1 || len(this.ItemInterface.GetId()) > AbSqlData_MAXNUM {
                tryControl1, tryResult1 = 1, false; return
            }
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            if this.ItemInterface.GetId() != nil && !stringValidationUtil.IsNumber(this.ItemInterface.GetId()) {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetNumber() == nil || len(this.ItemInterface.GetNumber()) < 1 || len(this.ItemInterface.GetNumber()) > AbSqlData_MAXNUM {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetNumber() != nil && !stringValidationUtil.IsNumber(this.ItemInterface.GetNumber()) {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetInBaskets() == nil || len(this.ItemInterface.GetInBaskets()) < 1 || len(this.ItemInterface.GetInBaskets()) > AbSqlData_MAXNUM {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetInBaskets() != nil && !stringValidationUtil.IsNumber(this.ItemInterface.GetInBaskets()) {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetWeight() == nil || len(this.ItemInterface.GetWeight()) < 1 || len(this.ItemInterface.GetWeight()) > this.MAXWEIGHT {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetWeight() != nil && !stringValidationUtil.IsNumber(this.ItemInterface.GetWeight()) {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetEnabled() == nil || len(this.ItemInterface.GetEnabled()) < 1 || len(this.ItemInterface.GetEnabled()) > AbSqlData_MAXSTRING {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetNewOrUsed() == nil || len(this.ItemInterface.GetNewOrUsed()) < 1 || len(this.ItemInterface.GetNewOrUsed()) > AbSqlData_MAXSTRING {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetSummary() == nil || len(this.ItemInterface.GetSummary()) < 1 || len(this.ItemInterface.GetSummary()) > AbSqlData_MAXSTRING {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetDescription() == nil || len(this.ItemInterface.GetDescription()) < 1 || len(this.ItemInterface.GetDescription()) > AbSqlData_MAXBLOB {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetKeywords() == nil || len(this.ItemInterface.GetKeywords()) < 1 || len(this.ItemInterface.GetKeywords()) > AbSqlData_MAXSTRING {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetCategory() == nil || len(this.ItemInterface.GetCategory()) < 1 || len(this.ItemInterface.GetCategory()) > AbSqlData_MAXSTRING {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetType() == nil || len(this.ItemInterface.GetType()) < 1 || len(this.ItemInterface.GetType()) > AbSqlData_MAXSTRING {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetPrice() == nil || !this.ItemInterface.GetPrice().IsValid() {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetCustoms() == nil || len(this.ItemInterface.GetCustoms()) < 1 || len(this.ItemInterface.GetCustoms()) > AbSqlData_MAXNUM {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetDownloads() == nil || len(this.ItemInterface.GetDownloads()) < 1 || len(this.ItemInterface.GetDownloads()) > AbSqlData_MAXNUM {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetGroups() == nil || len(this.ItemInterface.GetGroups()) < 1 || len(this.ItemInterface.GetGroups()) > AbSqlData_MAXNUM {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetOptions() == nil || len(this.ItemInterface.GetOptions()) < 1 || len(this.ItemInterface.GetOptions()) > AbSqlData_MAXNUM {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetPermissions() == nil || len(this.ItemInterface.GetPermissions()) < 1 || len(this.ItemInterface.GetPermissions()) > AbSqlData_MAXNUM {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.ItemInterface.GetSpecials() == nil || len(this.ItemInterface.GetSpecials()) < 1 || len(this.ItemInterface.GetSpecials()) > AbSqlData_MAXNUM {
                tryControl1, tryResult1 = 1, false; return
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasicItemValidation) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Validating Form"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if this.ItemInterface.GetId() == nil || len(this.ItemInterface.GetId()) < 1 || len(this.ItemInterface.GetId()) > AbSqlData_MAXNUM {
                stringBuffer.Append("Id is invalid. Must be number < " + fmt.Sprint(AbSqlData_MAXNUM) + " and > 0 digits.<br />")
            }
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            if this.ItemInterface.GetId() != nil && !stringValidationUtil.IsNumber(this.ItemInterface.GetId()) {
                stringBuffer.Append("Id is invalid. Must be > 0.<br />")
            }
            if this.ItemInterface.GetNumber() == nil || len(this.ItemInterface.GetNumber()) < 1 || len(this.ItemInterface.GetNumber()) > AbSqlData_MAXNUM {
                stringBuffer.Append("Number of items is invalid. Must be number < " + fmt.Sprint(AbSqlData_MAXNUM) + " and > 0 digits.<br />")
            }
            if this.ItemInterface.GetNumber() != nil && !stringValidationUtil.IsNumber(this.ItemInterface.GetNumber()) {
                stringBuffer.Append("Number of items value is not a valid number.<br />")
            }
            if this.ItemInterface.GetInBaskets() == nil || len(this.ItemInterface.GetInBaskets()) < 1 || len(this.ItemInterface.GetInBaskets()) > AbSqlData_MAXNUM {
                stringBuffer.Append("Number in baskets is invalid. Must be < " + fmt.Sprint(AbSqlData_MAXNUM) + " and > 0 digits.<br />")
            }
            if this.ItemInterface.GetInBaskets() != nil && !stringValidationUtil.IsNumber(this.ItemInterface.GetInBaskets()) {
                stringBuffer.Append("In Baskets value is not a valid number.<br />")
            }
            if this.ItemInterface.GetWeight() == nil || len(this.ItemInterface.GetWeight()) < 1 || len(this.ItemInterface.GetWeight()) > this.MAXWEIGHT {
                stringBuffer.Append("Weight value is invalid. Must be < " + fmt.Sprint(this.MAXWEIGHT) + " and > 0 digits.<br />")
            }
            if this.ItemInterface.GetWeight() != nil && !stringValidationUtil.IsNumber(this.ItemInterface.GetWeight()) {
                stringBuffer.Append("Weight value is not a valid number.<br />")
            }
            if this.ItemInterface.GetEnabled() == nil || len(this.ItemInterface.GetEnabled()) < 1 || len(this.ItemInterface.GetEnabled()) > AbSqlData_MAXSTRING {
                stringBuffer.Append("For Sale value is invalid. Must be < " + fmt.Sprint(AbSqlData_MAXSTRING) + " and > 0 characters long.<br />")
            }
            if this.ItemInterface.GetNewOrUsed() == nil || len(this.ItemInterface.GetNewOrUsed()) < 1 || len(this.ItemInterface.GetNewOrUsed()) > AbSqlData_MAXSTRING {
                stringBuffer.Append("New Or Used value is invalid. Must be < " + fmt.Sprint(AbSqlData_MAXSTRING) + " and > 0 characters long.<br />")
            }
            if this.ItemInterface.GetSummary() == nil || len(this.ItemInterface.GetSummary()) < 1 || len(this.ItemInterface.GetSummary()) > AbSqlData_MAXSTRING {
                stringBuffer.Append("Summary value is invalid. Must be < " + fmt.Sprint(AbSqlData_MAXSTRING) + " and > 0 characters long.<br />")
            }
            if this.ItemInterface.GetDescription() == nil || len(this.ItemInterface.GetDescription()) < 1 || len(this.ItemInterface.GetDescription()) > AbSqlData_MAXBLOB {
                stringBuffer.Append("Description value is invalid. Must be < " + fmt.Sprint(AbSqlData_MAXBLOB) + " and > 0 characters long.<br />")
            }
            if this.ItemInterface.GetKeywords() == nil || len(this.ItemInterface.GetKeywords()) < 1 || len(this.ItemInterface.GetKeywords()) > AbSqlData_MAXSTRING {
                stringBuffer.Append("Keywords value is invalid. Must be < " + fmt.Sprint(AbSqlData_MAXSTRING) + " and > 0 characters long.<br />")
            }
            if this.ItemInterface.GetCategory() == nil || len(this.ItemInterface.GetCategory()) < 1 || len(this.ItemInterface.GetCategory()) > AbSqlData_MAXSTRING {
                stringBuffer.Append("Category value is invalid. Must be < " + fmt.Sprint(AbSqlData_MAXSTRING) + " and > 0 characters long.<br />")
            }
            if this.ItemInterface.GetType() == nil || len(this.ItemInterface.GetType()) < 1 || len(this.ItemInterface.GetType()) > AbSqlData_MAXSTRING {
                stringBuffer.Append("Type value is invalid. Must be < " + fmt.Sprint(AbSqlData_MAXSTRING) + " and > 0 characters long.<br />")
            }
            if this.ItemInterface.GetPrice() == nil || !this.ItemInterface.GetPrice().IsValid() {
                stringBuffer.Append("Price value is invalid. Must be < " + fmt.Sprint(this.MAXPRICE) + " and > 0 in length.<br />")
            }
            if this.ItemInterface.GetCustoms() == nil || len(this.ItemInterface.GetCustoms()) < 1 || len(this.ItemInterface.GetCustoms()) > AbSqlData_MAXNUM {
                stringBuffer.Append("Custom items value is invalid. Must be number < " + fmt.Sprint(AbSqlData_MAXNUM) + " and > 0 digits.<br />")
            }
            if this.ItemInterface.GetDownloads() == nil || len(this.ItemInterface.GetDownloads()) < 1 || len(this.ItemInterface.GetDownloads()) > AbSqlData_MAXNUM {
                stringBuffer.Append("Downloadable item value is invalid. Must be number < " + fmt.Sprint(AbSqlData_MAXNUM) + " and > 0 digits.<br />")
            }
            if this.ItemInterface.GetGroups() == nil || len(this.ItemInterface.GetGroups()) < 1 || len(this.ItemInterface.GetGroups()) > AbSqlData_MAXNUM {
                stringBuffer.Append("Group item value is invalid. Must be number < " + fmt.Sprint(AbSqlData_MAXNUM) + " and > 0 digits.<br />")
            }
            if this.ItemInterface.GetOptions() == nil || len(this.ItemInterface.GetOptions()) < 1 || len(this.ItemInterface.GetOptions()) > AbSqlData_MAXNUM {
                stringBuffer.Append("Option item value is invalid. Must be number < " + fmt.Sprint(AbSqlData_MAXNUM) + " and > 0 digits.<br />")
            }
            if this.ItemInterface.GetPermissions() == nil || len(this.ItemInterface.GetPermissions()) < 1 || len(this.ItemInterface.GetPermissions()) > AbSqlData_MAXNUM {
                stringBuffer.Append("Permission item value is invalid. Must be number < " + fmt.Sprint(AbSqlData_MAXNUM) + " and > 0 digits.<br />")
            }
            if this.ItemInterface.GetSpecials() == nil || len(this.ItemInterface.GetSpecials()) < 1 || len(this.ItemInterface.GetSpecials()) > AbSqlData_MAXNUM {
                stringBuffer.Append("Special item value is invalid. Must be number < " + fmt.Sprint(AbSqlData_MAXNUM) + " and > 0 digits.<br />")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasicItemValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *BasicItemValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}

