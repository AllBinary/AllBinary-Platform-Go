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

type InitializerData struct {
    DBUSER string
    DBPASSWORD string
    ADMINJDBCDRIVER string
    ADMINSCHEMA string
    ADMINSERVER string
    ADMINPORT string
    CUSTOMERJDBCDRIVER string
    CUSTOMERNAME string
    CUSTOMERUSERNAME string
    CUSTOMERPASSWORD string
    CUSTOMERSCHEMA string
    CUSTOMERSERVER string
    CUSTOMERPORT string
    HISTORYJDBCDRIVER string
    HISTORYNAME string
    HISTORYUSERNAME string
    HISTORYPASSWORD string
    HISTORYSCHEMA string
    HISTORYSERVER string
    HISTORYPORT string
    LOGJDBCDRIVER string
    LOGNAME string
    LOGUSERNAME string
    LOGPASSWORD string
    LOGSCHEMA string
    LOGSERVER string
    LOGPORT string
    INVENTORYJDBCDRIVER string
    INVENTORYNAME string
    INVENTORYUSERNAME string
    INVENTORYPASSWORD string
    INVENTORYSCHEMA string
    INVENTORYSERVER string
    INVENTORYPORT string
    STATICPAGESJDBCDRIVER string
    STATICPAGESNAME string
    STATICPAGESUSERNAME string
    STATICPAGESPASSWORD string
    STATICPAGESSCHEMA string
    STATICPAGESSERVER string
    STATICPAGESPORT string
}

var InitializerData_Instance *InitializerData = NewInitializerData()

func (this *InitializerData) initFields() {
    this.DBUSER = "DBUSER"
    this.DBPASSWORD = "DBPASSWORD"
    this.ADMINJDBCDRIVER = "ADMIN_JDBCDRIVER"
    this.ADMINSCHEMA = "ADMINSCHEMA"
    this.ADMINSERVER = "ADMINSERVER"
    this.ADMINPORT = "ADMINPORT"
    this.CUSTOMERJDBCDRIVER = "CUSTOMER_JDBCDRIVER"
    this.CUSTOMERNAME = "CUSTOMERNAME"
    this.CUSTOMERUSERNAME = "CUSTOMERUSERNAME"
    this.CUSTOMERPASSWORD = "CUSTOMERPASSWORD"
    this.CUSTOMERSCHEMA = "CUSTOMERSCHEMA"
    this.CUSTOMERSERVER = "CUSTOMERSERVER"
    this.CUSTOMERPORT = "CUSTOMERPORT"
    this.HISTORYJDBCDRIVER = "HISTORY_JDBCDRIVER"
    this.HISTORYNAME = "HISTORYNAME"
    this.HISTORYUSERNAME = "HISTORYUSERNAME"
    this.HISTORYPASSWORD = "HISTORYPASSWORD"
    this.HISTORYSCHEMA = "HISTORYSCHEMA"
    this.HISTORYSERVER = "HISTORYSERVER"
    this.HISTORYPORT = "HISTORYPORT"
    this.LOGJDBCDRIVER = "LOG_JDBCDRIVER"
    this.LOGNAME = "LOGNAME"
    this.LOGUSERNAME = "LOGUSERNAME"
    this.LOGPASSWORD = "LOGPASSWORD"
    this.LOGSCHEMA = "LOGSCHEMA"
    this.LOGSERVER = "LOGSERVER"
    this.LOGPORT = "LOGPORT"
    this.INVENTORYJDBCDRIVER = "INVENTORY_JDBCDRIVER"
    this.INVENTORYNAME = "INVENTORYNAME"
    this.INVENTORYUSERNAME = "INVENTORYUSERNAME"
    this.INVENTORYPASSWORD = "INVENTORYPASSWORD"
    this.INVENTORYSCHEMA = "INVENTORYSCHEMA"
    this.INVENTORYSERVER = "INVENTORYSERVER"
    this.INVENTORYPORT = "INVENTORYPORT"
    this.STATICPAGESJDBCDRIVER = "STATICPAGES_JDBCDRIVER"
    this.STATICPAGESNAME = "STATICPAGESNAME"
    this.STATICPAGESUSERNAME = "STATICPAGESUSERNAME"
    this.STATICPAGESPASSWORD = "STATICPAGESPASSWORD"
    this.STATICPAGESSCHEMA = "STATICPAGESSCHEMA"
    this.STATICPAGESSERVER = "STATICPAGESSERVER"
    this.STATICPAGESPORT = "STATICPAGESPORT"
}

func NewInitializerData() *InitializerData {
    this := &InitializerData{}
    this.initFields()
    return this
}


func InitializerData_GetInstance() *InitializerData {
        return InitializerData_Instance
}


func (this *InitializerData) GetJdbcDriverSolutionInfo() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("The following describes the possible problems and solutions regarding the Jdbc Driver error(s):<p/>")
        stringBuffer.Append("1. A JDBC driver you specified is not in your classpath.<br/>")
        stringBuffer.Append("Solution 1: Move the JDBC driver into any directory specified in the existing classpath. <br/>")
        stringBuffer.Append("Solution 2: Add the directory that contains the JDBC driver to the classpath.<br/>")
        stringBuffer.Append("Solution 3: Add the JDBC driver to the WEB-INF/lib directory where you installed this webapp.<br/>")
        stringBuffer.Append("2. The JDBC driver you specified does not exit.<br/>")
        stringBuffer.Append("Solution: Get a JDBC driver. <br/>")
        stringBuffer.Append("3. The JDBC driver you specified is not valid.<br/>")
        stringBuffer.Append("Solution: Use a valid JDBC driver.<p/>")
        return fmt.Sprint(stringBuffer)
}

