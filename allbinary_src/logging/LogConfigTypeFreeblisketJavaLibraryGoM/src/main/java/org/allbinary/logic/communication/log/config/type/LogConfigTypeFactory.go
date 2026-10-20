/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type LogConfigTypeFactory struct {
    NO_DESCRIPTION string
    AUTHENTICATION *LogConfigType
    AUTHENTICATIONERROR *LogConfigType
    CATEGORY *LogConfigType
    CATEGORYERROR *LogConfigType
    CRYPT *LogConfigType
    CRYPTERROR *LogConfigType
    EMAILLOGGING *LogConfigType
    EMAILLOGGINGERROR *LogConfigType
    ENTITYFACTORYERROR *LogConfigType
    NETBEANS_MODULE *LogConfigType
    NETBEANS_MODULE_ERROR *LogConfigType
    FACTORYERROR *LogConfigType
    FILE *LogConfigType
    FILEERROR *LogConfigType
    HTTP *LogConfigType
    HTTPERROR *LogConfigType
    HTTPREQUEST *LogConfigType
    HTTPREQUESTERROR *LogConfigType
    IDLOGGING *LogConfigType
    JSPEXTRAOUTPUT *LogConfigType
    JSPTAGEXTRAOUTPUT *LogConfigType
    JSPTAG *LogConfigType
    JSPTAGERROR *LogConfigType
    LICENSING *LogConfigType
    LICENSINGERROR *LogConfigType
    LOADER *LogConfigType
    LOADERERROR *LogConfigType
    OS *LogConfigType
    OSERROR *LogConfigType
    PAYMENT *LogConfigType
    PAYMENTERROR *LogConfigType
    PRELOADER *LogConfigType
    PRELOADERERROR *LogConfigType
    PRODUCTSEARCHLOGGING *LogConfigType
    PRODUCTSEARCHLOGGINGERROR *LogConfigType
    SERVLET *LogConfigType
    SERVLETERROR *LogConfigType
    SHIPPINGERROR *LogConfigType
    SQLLOGGINGERROR *LogConfigType
    SQLLOGGING *LogConfigType
    SQLTAGS *LogConfigType
    SQLTAGSERROR *LogConfigType
    STYLE *LogConfigType
    STYLEERROR *LogConfigType
    TAGHELPER *LogConfigType
    TAGHELPERERROR *LogConfigType
    TAGHELPERFACTORY *LogConfigType
    TAGHELPERFACTORYERROR *LogConfigType
    TAX *LogConfigType
    TAXERROR *LogConfigType
    VALIDATION *LogConfigType
    VALIDATIONERROR *LogConfigType
    VIEW *LogConfigType
    VIEWERROR *LogConfigType
    VIEWFACTORYERROR *LogConfigType
    WORKFLOW *LogConfigType
    WORKFLOWERROR *LogConfigType
    XMLLOGGING *LogConfigType
    XMLLOGGINGERROR *LogConfigType
    XSLLOGGING *LogConfigType
    XSLLOGGINGERROR *LogConfigType
    REPLACE *LogConfigType
    REPLACE_INFO *LogConfigType
    REPLACEERROR *LogConfigType
    SQLLOGGINGPOOL *LogConfigType
    SQLLOGGINGPOOLERROR *LogConfigType
    STATICPAGEGENERATIONLOGGING *LogConfigType
    STATICPAGEGENERATIONLOGGINGERROR *LogConfigType
}

var LogConfigTypeFactory_Instance *LogConfigTypeFactory = NewLogConfigTypeFactory()

func (this *LogConfigTypeFactory) initFields() {
    this.NO_DESCRIPTION = "No Description"
    this.AUTHENTICATION = NewLogConfigType("Authentication", this.NO_DESCRIPTION)
    this.AUTHENTICATIONERROR = NewLogConfigType("Authentication Error", this.NO_DESCRIPTION)
    this.CATEGORY = NewLogConfigType("Category", this.NO_DESCRIPTION)
    this.CATEGORYERROR = NewLogConfigType("Category Error", this.NO_DESCRIPTION)
    this.CRYPT = NewLogConfigType("Crypt", this.NO_DESCRIPTION)
    this.CRYPTERROR = NewLogConfigType("Crypt Error", this.NO_DESCRIPTION)
    this.EMAILLOGGING = NewLogConfigType("Email", this.NO_DESCRIPTION)
    this.EMAILLOGGINGERROR = NewLogConfigType("Email Error", this.NO_DESCRIPTION)
    this.ENTITYFACTORYERROR = NewLogConfigType("Entity Factory Error", this.NO_DESCRIPTION)
    this.NETBEANS_MODULE = NewLogConfigType("NetBeans Module", this.NO_DESCRIPTION)
    this.NETBEANS_MODULE_ERROR = NewLogConfigType("NetBeans Module Error", this.NO_DESCRIPTION)
    this.FACTORYERROR = NewLogConfigType("General Factory Error", this.NO_DESCRIPTION)
    this.FILE = NewLogConfigType("File", this.NO_DESCRIPTION)
    this.FILEERROR = NewLogConfigType("File Error", this.NO_DESCRIPTION)
    this.HTTP = NewLogConfigType("HTTP", this.NO_DESCRIPTION)
    this.HTTPERROR = NewLogConfigType("HTTP Error", this.NO_DESCRIPTION)
    this.HTTPREQUEST = NewLogConfigType("Http Request", this.NO_DESCRIPTION)
    this.HTTPREQUESTERROR = NewLogConfigType("Http Request Error", this.NO_DESCRIPTION)
    this.IDLOGGING = NewLogConfigType("ID Logging", this.NO_DESCRIPTION)
    this.JSPEXTRAOUTPUT = NewLogConfigType("JSP Extra Output", this.NO_DESCRIPTION)
    this.JSPTAGEXTRAOUTPUT = NewLogConfigType("JSP Tag Extra Output", this.NO_DESCRIPTION)
    this.JSPTAG = NewLogConfigType("JSP Tag", this.NO_DESCRIPTION)
    this.JSPTAGERROR = NewLogConfigType("JSP Tag Error", this.NO_DESCRIPTION)
    this.LICENSING = NewLogConfigType("Licensing", this.NO_DESCRIPTION)
    this.LICENSINGERROR = NewLogConfigType("Licensing Error", this.NO_DESCRIPTION)
    this.LOADER = NewLogConfigType("Loader", this.NO_DESCRIPTION)
    this.LOADERERROR = NewLogConfigType("Loader Error", this.NO_DESCRIPTION)
    this.OS = NewLogConfigType("OS", this.NO_DESCRIPTION)
    this.OSERROR = NewLogConfigType("OS Error", this.NO_DESCRIPTION)
    this.PAYMENT = NewLogConfigType("Payment", this.NO_DESCRIPTION)
    this.PAYMENTERROR = NewLogConfigType("Payment Error", this.NO_DESCRIPTION)
    this.PRELOADER = NewLogConfigType("PreLoader", this.NO_DESCRIPTION)
    this.PRELOADERERROR = NewLogConfigType("PreLoader Error", this.NO_DESCRIPTION)
    this.PRODUCTSEARCHLOGGING = NewLogConfigType("Product Search Logging", this.NO_DESCRIPTION)
    this.PRODUCTSEARCHLOGGINGERROR = NewLogConfigType("Product Search Logging Error", this.NO_DESCRIPTION)
    this.SERVLET = NewLogConfigType("Servlet", this.NO_DESCRIPTION)
    this.SERVLETERROR = NewLogConfigType("Servlet Error", this.NO_DESCRIPTION)
    this.SHIPPINGERROR = NewLogConfigType("Shipping Error", this.NO_DESCRIPTION)
    this.SQLLOGGINGERROR = NewLogConfigType("SQL Error", this.NO_DESCRIPTION)
    this.SQLLOGGING = NewLogConfigType("SQL Logging", this.NO_DESCRIPTION)
    this.SQLTAGS = NewLogConfigType("SQL Tag Logging", this.NO_DESCRIPTION)
    this.SQLTAGSERROR = NewLogConfigType("SQL Tag Error", this.NO_DESCRIPTION)
    this.STYLE = NewLogConfigType("Style", this.NO_DESCRIPTION)
    this.STYLEERROR = NewLogConfigType("Style Error", this.NO_DESCRIPTION)
    this.TAGHELPER = NewLogConfigType("Tag Helper", this.NO_DESCRIPTION)
    this.TAGHELPERERROR = NewLogConfigType("Tag Helper Error", this.NO_DESCRIPTION)
    this.TAGHELPERFACTORY = NewLogConfigType("Tag Helper Factory", this.NO_DESCRIPTION)
    this.TAGHELPERFACTORYERROR = NewLogConfigType("Tag Helper Factory Error", this.NO_DESCRIPTION)
    this.TAX = NewLogConfigType("Tax", this.NO_DESCRIPTION)
    this.TAXERROR = NewLogConfigType("Tax Error", this.NO_DESCRIPTION)
    this.VALIDATION = NewLogConfigType("Validation", this.NO_DESCRIPTION)
    this.VALIDATIONERROR = NewLogConfigType("Validation Error", this.NO_DESCRIPTION)
    this.VIEW = NewLogConfigType("View", this.NO_DESCRIPTION)
    this.VIEWERROR = NewLogConfigType("View Error", this.NO_DESCRIPTION)
    this.VIEWFACTORYERROR = NewLogConfigType("View Factory Error", this.NO_DESCRIPTION)
    this.WORKFLOW = NewLogConfigType("Workflow", this.NO_DESCRIPTION)
    this.WORKFLOWERROR = NewLogConfigType("Workflow Error", this.NO_DESCRIPTION)
    this.XMLLOGGING = NewLogConfigType("XML Logging", this.NO_DESCRIPTION)
    this.XMLLOGGINGERROR = NewLogConfigType("XML Logging Error", this.NO_DESCRIPTION)
    this.XSLLOGGING = NewLogConfigType("XSL Logging", this.NO_DESCRIPTION)
    this.XSLLOGGINGERROR = NewLogConfigType("XSL Logging Error", this.NO_DESCRIPTION)
    this.REPLACE = NewLogConfigType("Replace", this.NO_DESCRIPTION)
    this.REPLACE_INFO = NewLogConfigType("Replace Info", this.NO_DESCRIPTION)
    this.REPLACEERROR = NewLogConfigType("Replace Error", this.NO_DESCRIPTION)
    this.SQLLOGGINGPOOL = NewLogConfigType("SQL Pool", this.NO_DESCRIPTION)
    this.SQLLOGGINGPOOLERROR = NewLogConfigType("SQL Pool Error", this.NO_DESCRIPTION)
    this.STATICPAGEGENERATIONLOGGING = NewLogConfigType("Static Page Generation Logging", this.NO_DESCRIPTION)
    this.STATICPAGEGENERATIONLOGGINGERROR = NewLogConfigType("Static Page Generation Logging Error", this.NO_DESCRIPTION)
}


func LogConfigTypeFactory_GetInstance() *LogConfigTypeFactory {
        return LogConfigTypeFactory_Instance
}


func NewLogConfigTypeFactory() *LogConfigTypeFactory {
    this := &LogConfigTypeFactory{}
    this.initFields()
        PreLogUtil_Put("FreeBlisket", this, CommonStrings_GetInstance().CONSTRUCTOR)
    return this
}

