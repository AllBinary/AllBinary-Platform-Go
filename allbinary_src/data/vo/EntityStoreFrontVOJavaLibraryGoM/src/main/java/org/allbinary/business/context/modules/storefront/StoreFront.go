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

type StoreFront struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Directory *Directory
    StringUtil *StringUtil
    AbPathUtil *AbPathUtil
    Name string
    BasketName string
    HomeHostName string
    HomeHostNamePath *AbPath
    HostName string
    HostNamePath *AbPath
    TestHomeHostName string
    TestHomeHostNamePath *AbPath
    TestHostName string
    TestHostNamePath *AbPath
    ImagePath *AbPath
    StaticPath *AbPath
    CategoryPath *AbPath
    InventoryControl string
    ContextConfigurationInterface ContextConfigurationInterface
    SubStores string
    TagLocation string
    PackageLocation string
    Ftp string
    FtpPath *AbPath
    FtpUserName string
    FtpPassword string
    TestFtp string
    TestFtpPath *AbPath
    TestFtpUserName string
    TestFtpPassword string
    TimeCreated string
    LastModified string
    MINCHAR int
    MINSTORENAMELENGTH int
    MAXCHAR int
    EMPTY_STRING string
    ONE_EMPTY_STRING_ARRAY []string
}

// Implements StoreFrontInterface

var StoreFront_RESOURCES string = " Resources" + fmt.Sprint(AbPathData_GetInstance().SEPARATOR)

func (this *StoreFront) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.Directory = Directory_GetInstance()
    this.StringUtil = StringUtil_GetInstance()
    this.AbPathUtil = AbPathUtil_GetInstance()
    this.Name = this.StringUtil.EMPTY_STRING
    this.BasketName = this.StringUtil.EMPTY_STRING
    this.HomeHostName = this.StringUtil.EMPTY_STRING
    this.HomeHostNamePath = this.AbPathUtil.NO_ABPATH
    this.HostName = this.StringUtil.EMPTY_STRING
    this.HostNamePath = this.AbPathUtil.NO_ABPATH
    this.TestHomeHostName = this.StringUtil.EMPTY_STRING
    this.TestHomeHostNamePath = this.AbPathUtil.NO_ABPATH
    this.TestHostName = this.StringUtil.EMPTY_STRING
    this.TestHostNamePath = this.AbPathUtil.NO_ABPATH
    this.ImagePath = this.AbPathUtil.NO_ABPATH
    this.StaticPath = this.AbPathUtil.NO_ABPATH
    this.CategoryPath = this.AbPathUtil.NO_ABPATH
    this.InventoryControl = this.StringUtil.EMPTY_STRING
    
    this.SubStores = this.StringUtil.EMPTY_STRING
    this.TagLocation = this.StringUtil.EMPTY_STRING
    this.PackageLocation = this.StringUtil.EMPTY_STRING
    this.Ftp = this.StringUtil.EMPTY_STRING
    this.FtpPath = this.AbPathUtil.NO_ABPATH
    this.FtpUserName = this.StringUtil.EMPTY_STRING
    this.FtpPassword = this.StringUtil.EMPTY_STRING
    this.TestFtp = this.StringUtil.EMPTY_STRING
    this.TestFtpPath = this.AbPathUtil.NO_ABPATH
    this.TestFtpUserName = this.StringUtil.EMPTY_STRING
    this.TestFtpPassword = this.StringUtil.EMPTY_STRING
    this.TimeCreated = this.StringUtil.EMPTY_STRING
    this.LastModified = this.StringUtil.EMPTY_STRING
    this.MINCHAR = 0
    this.MINSTORENAMELENGTH = 1
    this.MAXCHAR = 255
    this.EMPTY_STRING = StringUtil_GetInstance().EMPTY_STRING
    this.ONE_EMPTY_STRING_ARRAY = StringUtil_GetInstance().ONE_EMPTY_STRING_ARRAY
}


func NewStoreFront() *StoreFront {
    this := &StoreFront{}
    this.initFields()
        this.ContextConfigurationInterface = NewContextConfiguration()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewStoreFront(request *HttpServletRequest) *StoreFront {
    this := &StoreFront{}
    this.initFields()
        this.GetFormData(NewRequestParams(request).ToHashMap())
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF("TWB- needs update for adding and updating context configuration - currently defaults to admin context config", this, "StoreFront")
        }
        this.SetContextConfigurationInterface(ContextConfigurationInterfaceFactory_GetInstance().GetInstance("Admin"))
    return this
}


// Java overloaded constructor: Go needs a different name
func NewStoreFront(hashMap *HashMap) *StoreFront {
    this := &StoreFront{}
    this.initFields()
        this.GetFormData(hashMap)
        var domDocumentString string = any(hashMap.Get(StoreFrontData_GetInstance().CONFIGURATION)).(string); _ = domDocumentString
        if domDocumentString == "" {
            var document *Document = DomDocumentHelper_Create(domDocumentString); _ = document
            this.SetContextConfigurationInterface(ContextConfigurationInterfaceFactory_GetInstance().GetInstance(document))
        }
    return this
}


func (this *StoreFront) IsColumn(columnName string) bool {
        if strings.Compare(columnName, StoreFrontData_GetInstance().NAME) == 0 || strings.Compare(columnName, UserData_MAINEMAIL) == 0 {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *StoreFront) CreatePath(key string, hashMap *HashMap) *AbPath {
        return NewAbPath(this.Create(this.EMPTY_STRING, key, hashMap), StringUtil_GetInstance().EMPTY_STRING)
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *StoreFront) CreatePath(append_ string, key string, hashMap *HashMap) *AbPath {
        return NewAbPath(this.Create(append_, key, hashMap), StringUtil_GetInstance().EMPTY_STRING)
}


func (this *StoreFront) Get(key string, hashMap *HashMap) string {
        return this.Create(this.EMPTY_STRING, key, hashMap)
}


func (this *StoreFront) Create(append_ string, key string, hashMap *HashMap) string {
        var path string = any(hashMap.Get(key)).(string); _ = path
        if StringValidationUtil_GetInstance().IsEmpty(path) {
            return this.EMPTY_STRING
        } else {
            if this.IsColumn(path) {
                return fmt.Sprint(any(hashMap.Get(path)).(string)) + append_
            } else {
                return path
            }
        }
}


// throws Exception
func (this *StoreFront) GetFormData(storeHashMap *HashMap) {
        var storeFrontData *StoreFrontData = StoreFrontData_GetInstance(); _ = storeFrontData
        this.Name = any(storeHashMap.Get(storeFrontData.NAME)).(string)
        this.BasketName = any(storeHashMap.Get(storeFrontData.NAME)).(string)
        this.HomeHostName = any(storeHashMap.Get(storeFrontData.HOMEHOSTNAME)).(string)
        this.HomeHostNamePath = this.CreatePath(storeFrontData.HOMEHOSTNAMEPATH, storeHashMap)
        this.HostName = any(storeHashMap.Get(storeFrontData.HOSTNAME)).(string)
        this.HostNamePath = this.CreatePath(storeFrontData.HOSTNAMEPATH, storeHashMap)
        this.TestHomeHostName = any(storeHashMap.Get(storeFrontData.TESTHOMEHOSTNAME)).(string)
        this.TestHomeHostNamePath = this.CreatePath(storeFrontData.TESTHOMEHOSTNAMEPATH, storeHashMap)
        this.TestHostName = any(storeHashMap.Get(storeFrontData.TESTHOSTNAME)).(string)
        this.TestHostNamePath = this.CreatePath(storeFrontData.TESTHOSTNAMEPATH, storeHashMap)
        this.ImagePath = this.CreatePath(StoreFront_RESOURCES, storeFrontData.IMAGEPATH, storeHashMap)
        this.StaticPath = this.CreatePath(" Products" + fmt.Sprint(AbPathData_GetInstance().SEPARATOR), storeFrontData.STATICPATH, storeHashMap)
        this.CategoryPath = this.CreatePath(StoreFront_RESOURCES, storeFrontData.CATEGORYPATH, storeHashMap)
        this.InventoryControl = any(storeHashMap.Get(storeFrontData.INVENTORYCONTROL)).(string)
        this.SubStores = any(storeHashMap.Get(storeFrontData.SUBSTORES)).(string)
        this.TagLocation = any(storeHashMap.Get(storeFrontData.TAGLOCATION)).(string)
        this.PackageLocation = any(storeHashMap.Get(storeFrontData.PACKAGELOCATION)).(string)
        this.Ftp = any(storeHashMap.Get(storeFrontData.FTP)).(string)
        this.FtpPath = this.CreatePath(storeFrontData.FTPPATH, storeHashMap)
        this.FtpUserName = any(storeHashMap.Get(storeFrontData.FTPUSERNAME)).(string)
        this.FtpPassword = any(storeHashMap.Get(storeFrontData.FTPPASSWORD)).(string)
        this.TestFtp = any(storeHashMap.Get(storeFrontData.TESTFTP)).(string)
        this.TestFtpPath = this.CreatePath(storeFrontData.TESTFTPPATH, storeHashMap)
        this.TestFtpUserName = any(storeHashMap.Get(storeFrontData.TESTFTPUSERNAME)).(string)
        this.TestFtpPassword = any(storeHashMap.Get(storeFrontData.TESTFTPPASSWORD)).(string)
        this.TimeCreated = any(storeHashMap.Get(EntryData_GetInstance().TIMECREATED)).(string)
        this.LastModified = any(storeHashMap.Get(EntryData_GetInstance().LASTMODIFIED)).(string)
}


func (this *StoreFront) IsNameValid() bool {
        if !StringValidationUtil_GetInstance().IsValidRequired(this.Name, this.MINSTORENAMELENGTH, this.MAXCHAR) {
            return false
        } else {
            var isEscapedCharactersContained bool = false; _ = isEscapedCharactersContained
            var storeNameEscaped string = StringEscapeUtils_EscapeHtml3(this.Name); _ = storeNameEscaped
            if strings.Compare(this.Name, storeNameEscaped) != 0 {
                isEscapedCharactersContained = true
            }
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            var isSpacesContained bool = false; _ = isSpacesContained
            if stringValidationUtil.ContainsSpaces(this.Name) {
                isSpacesContained = true
            }
            if isSpacesContained || isEscapedCharactersContained {
                var hashMap *HashMap = SpecialCharacterUtil_GetHashMap(); _ = hashMap
                hashMap.Put([]string{CommonSeps_GetInstance().SPACE}, this.ONE_EMPTY_STRING_ARRAY)
                this.Name = NewReplace(hashMap).All(this.Name)
            }
        }
        return true
}


func (this *StoreFront) IsValid() bool {
        var valid bool = true; _ = valid
        valid = this.IsNameValid()
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if !stringValidationUtil.IsValidRequired(this.BasketName, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidRequired(this.HomeHostName, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(fmt.Sprint(this.HomeHostNamePath), this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidRequired(this.HostName, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidRequired(fmt.Sprint(this.HostNamePath), this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.TestHomeHostName, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(fmt.Sprint(this.TestHomeHostNamePath), this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.TestHostName, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(fmt.Sprint(this.TestHostNamePath), this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidRequired(fmt.Sprint(this.ImagePath), this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidRequired(fmt.Sprint(this.StaticPath), this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidRequired(fmt.Sprint(this.CategoryPath), this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidRequired(this.InventoryControl, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.SubStores, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.TagLocation, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidRequired(this.PackageLocation, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.Ftp, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(fmt.Sprint(this.FtpPath), this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.FtpUserName, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.FtpPassword, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.TestFtp, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(fmt.Sprint(this.TestFtpPath), this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.TestFtpUserName, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.TestFtpPassword, this.MINCHAR, this.MAXCHAR) {
            valid = false
        }
        if valid == true {
            return true
        } else {
            return false
        }
}


func (this *StoreFront) NameValidationInfo() string {
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if !stringValidationUtil.IsValidRequired(this.Name, this.MINCHAR, this.MAXCHAR) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Name is invalid. Must be < ")
            stringBuffer.Appendint(this.MAXCHAR)
            stringBuffer.Append(" and > 0 characters.<br>")
            return fmt.Sprint(stringBuffer)
        }
        return this.EMPTY_STRING
}


func (this *StoreFront) ValidationInfo() string {
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
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.NameValidationInfo())
            if !stringValidationUtil.IsValidRequired(this.BasketName, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Basket name is invalid. Must be < ")
                stringBuffer.Appendint(this.MAXCHAR)
                stringBuffer.Append(" and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidRequired(this.HomeHostName, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Home host name is invalid. Must be < ")
                stringBuffer.Appendint(this.MAXCHAR)
                stringBuffer.Append(" and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(fmt.Sprint(this.HomeHostNamePath), this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Home host name path is invalid. Must be < ")
                stringBuffer.Appendint(this.MAXCHAR)
                stringBuffer.Append(" and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidRequired(this.HostName, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Host name is invalid. Must be < ")
                stringBuffer.Appendint(this.MAXCHAR)
                stringBuffer.Append(" and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidRequired(fmt.Sprint(this.HostNamePath), this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Host name path is invalid. Must be < ")
                stringBuffer.Appendint(this.MAXCHAR)
                stringBuffer.Append(" and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(this.TestHomeHostName, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Test home host name is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(fmt.Sprint(this.TestHomeHostNamePath), this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Test home host name path is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(this.TestHostName, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Test host name is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(fmt.Sprint(this.TestHostNamePath), this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Test host name path is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidRequired(fmt.Sprint(this.ImagePath), this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Image path is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidRequired(fmt.Sprint(this.StaticPath), this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Static path is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidRequired(fmt.Sprint(this.CategoryPath), this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Category path is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidRequired(this.InventoryControl, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Inventory control is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(this.SubStores, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Sub stores is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(this.TagLocation, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Tag location is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidRequired(this.PackageLocation, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Package location is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(this.Ftp, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Ftp is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(fmt.Sprint(this.FtpPath), this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Ftp path is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(this.FtpUserName, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Ftp username is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(this.FtpPassword, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Ftp password is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(this.TestFtp, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Test ftp is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(fmt.Sprint(this.TestFtpPath), this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Test ftp Path is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(this.TestFtpUserName, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Test ftp username is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            if !stringValidationUtil.IsValidNotRequired(this.TestFtpPassword, this.MINCHAR, this.MAXCHAR) {
                stringBuffer.Append("Test ftp password is invalid. Must be < " + fmt.Sprint(this.MAXCHAR) + " and > 0 characters.<br>")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StoreFront) GetTestHtmlPath() string {
        return org_Allbinary.Globals.URLGLOBALS.GetTestHtmlPath() + this.GetCurrentHostNamePath()
}


func (this *StoreFront) GetCurrentHostName() string {
        var location string = this.EMPTY_STRING; _ = location
        if org_Allbinary.Globals.URLGLOBALS.IsTestingMode() {
            location = this.GetTestHostName()
        } else {
            location = this.GetHostName()
        }
        return location
}


func (this *StoreFront) GetCurrentHostNamePath() string {
        var location string = this.EMPTY_STRING; _ = location
        if org_Allbinary.Globals.URLGLOBALS.IsTestingMode() {
            location = this.GetTestHostNamePath()
        } else {
            location = this.GetHostNamePath()
        }
        return location
}


func (this *StoreFront) GetCurrentHomeHostName() string {
        var location string = this.EMPTY_STRING; _ = location
        if org_Allbinary.Globals.URLGLOBALS.IsTestingMode() {
            location = this.GetTestHomeHostName()
        } else {
            location = this.GetHomeHostName()
        }
        return location
}


func (this *StoreFront) GetCurrentHomeHostNamePath() string {
        var location string = this.EMPTY_STRING; _ = location
        if org_Allbinary.Globals.URLGLOBALS.IsTestingMode() {
            location = this.GetTestHomeHostNamePath()
        } else {
            location = this.GetHomeHostNamePath()
        }
        return location
}


func (this *StoreFront) GetName() string {
        return this.Name
}


func (this *StoreFront) GetBasketName() string {
        return this.BasketName
}


func (this *StoreFront) GetHomeHostName() string {
        return this.HomeHostName
}


func (this *StoreFront) GetHomeHostNamePath() string {
        return fmt.Sprint(this.HomeHostNamePath)
}


func (this *StoreFront) GetHostName() string {
        return this.HostName
}


func (this *StoreFront) GetHostNamePath() string {
        return fmt.Sprint(this.HostNamePath)
}


func (this *StoreFront) GetTestHomeHostName() string {
        return this.TestHomeHostName
}


func (this *StoreFront) GetTestHomeHostNamePath() string {
        return fmt.Sprint(this.TestHomeHostNamePath)
}


func (this *StoreFront) GetTestHostName() string {
        return this.TestHostName
}


func (this *StoreFront) GetTestHostNamePath() string {
        return fmt.Sprint(this.TestHostNamePath)
}


func (this *StoreFront) GetStaticPath() string {
        return fmt.Sprint(this.StaticPath)
}


func (this *StoreFront) GetCategoryPath() string {
        return fmt.Sprint(this.CategoryPath)
}


func (this *StoreFront) GetInventoryControl() string {
        return this.InventoryControl
}


// throws Exception
func (this *StoreFront) GetSubStores() *BasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var tokenizer *Tokenizer = NewTokenizer(CommonSeps_GetInstance().SEMICOLON); _ = tokenizer
            var subStoreVector *BasicArrayList = tokenizer.GetTokensFromString(this.SubStores, NewBasicArrayListD()); _ = subStoreVector
            tryControl1, tryResult1 = 1, subStoreVector; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StoreFront) GetTagLocation() string {
        return this.TagLocation
}


func (this *StoreFront) GetPackageLocation() string {
        return this.PackageLocation
}


func (this *StoreFront) GetFtp() string {
        return this.Ftp
}


func (this *StoreFront) GetFtpUserName() string {
        return this.FtpUserName
}


func (this *StoreFront) GetFtpPassword() string {
        return this.FtpPassword
}


func (this *StoreFront) GetTestFtp() string {
        return this.TestFtp
}


func (this *StoreFront) GetTestFtpUserName() string {
        return this.TestFtpUserName
}


func (this *StoreFront) GetTestFtpPassword() string {
        return this.TestFtpPassword
}


func (this *StoreFront) GetFtpPath() string {
        return fmt.Sprint(this.FtpPath)
}


func (this *StoreFront) GetTestFtpPath() string {
        return fmt.Sprint(this.TestFtpPath)
}


func (this *StoreFront) GetTimeCreated() string {
        return this.TimeCreated
}


func (this *StoreFront) GetLastModified() string {
        return this.LastModified
}


func (this *StoreFront) SetName(value string) {
        this.Name = value
}


func (this *StoreFront) SetBasketName(value string) {
        this.BasketName = value
}


func (this *StoreFront) SetHomeHostName(value string) {
        this.HomeHostName = value
}


// throws Exception
func (this *StoreFront) SetHomeHostNamePath(value string) {
        this.HomeHostNamePath = NewAbPath(value, StringUtil_GetInstance().EMPTY_STRING)
}


func (this *StoreFront) SetHostName(value string) {
        this.HostName = value
}


// throws Exception
func (this *StoreFront) SetHostNamePath(value string) {
        this.HostNamePath = NewAbPath(value, StringUtil_GetInstance().EMPTY_STRING)
}


func (this *StoreFront) SetTestHomeHostName(value string) {
        this.TestHomeHostName = value
}


// throws Exception
func (this *StoreFront) SetTestHomeHostNamePath(value string) {
        this.TestHomeHostNamePath = NewAbPath(value, StringUtil_GetInstance().EMPTY_STRING)
}


func (this *StoreFront) SetTestHostName(value string) {
        this.TestHostName = value
}


// throws Exception
func (this *StoreFront) SetTestHostNamePath(value string) {
        this.TestHostNamePath = NewAbPath(value, StringUtil_GetInstance().EMPTY_STRING)
}


// throws Exception
func (this *StoreFront) SetImagePath(value string) {
        this.ImagePath = NewAbPath(value, StringUtil_GetInstance().EMPTY_STRING)
}


// throws Exception
func (this *StoreFront) SetStaticPath(value string) {
        this.StaticPath = NewAbPath(value, StringUtil_GetInstance().EMPTY_STRING)
}


// throws Exception
func (this *StoreFront) SetCategoryPath(value string) {
        this.CategoryPath = NewAbPath(value, StringUtil_GetInstance().EMPTY_STRING)
}


func (this *StoreFront) SetSubStores(value string) {
        this.SubStores = value
}


func (this *StoreFront) SetTagLocation(value string) {
        this.TagLocation = value
}


func (this *StoreFront) SetPackageLocation(value string) {
        this.PackageLocation = value
}


func (this *StoreFront) SetInventoryControl(value string) {
        this.InventoryControl = value
}


// Java overloaded method: Go needs a different name
func (this *StoreFront) GetPackageLocation(value string) {
        this.PackageLocation = value
}


func (this *StoreFront) SetFtp(value string) {
        this.Ftp = value
}


func (this *StoreFront) SetFtpUserName(value string) {
        this.FtpUserName = value
}


func (this *StoreFront) SetFtpPassword(value string) {
        this.FtpPassword = value
}


func (this *StoreFront) SetTestFtp(value string) {
        this.TestFtp = value
}


func (this *StoreFront) SetTestFtpUserName(value string) {
        this.TestFtpUserName = value
}


func (this *StoreFront) SetTestFtpPassword(value string) {
        this.TestFtpPassword = value
}


func (this *StoreFront) SetTimeCreated(value string) {
        this.TimeCreated = value
}


func (this *StoreFront) SetLastModified(value string) {
        this.LastModified = value
}


// throws Exception
func (this *StoreFront) SetFtpPath(value string) {
        this.FtpPath = NewAbPath(value, StringUtil_GetInstance().EMPTY_STRING)
}


// throws Exception
func (this *StoreFront) SetTestFtpPath(value string) {
        this.TestFtpPath = NewAbPath(value, StringUtil_GetInstance().EMPTY_STRING)
}


// throws Exception
func (this *StoreFront) CreateDirectories() bool {
        var storeAbPath *AbPath = NewAbPath(URLGLOBALS.GetWebappPath() + this.GetCurrentHomeHostNamePath(), StringUtil_GetInstance().EMPTY_STRING); _ = storeAbPath
        if !this.Directory.Create(storeAbPath) {
            return false
        }
        if !this.Directory.Create(NewAbPath(storeAbPath + this.GetCategoryPath(), StringUtil_GetInstance().EMPTY_STRING)) {
            return false
        }
        if !this.Directory.Create(NewAbPath(storeAbPath + this.GetStaticPath(), StringUtil_GetInstance().EMPTY_STRING)) {
            return false
        }
        return true
}


// throws Exception
func (this *StoreFront) Install(current int, total int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "install()", e)
            }
            panic(e)
                }
            }()
            if !this.CreateDirectories() {
                panic(errors.New(fmt.Sprint("Unable to create store directories")))
            }
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(URLGLOBALS.GetMainPath())
            stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
            var viewPath string = fmt.Sprint(stringBuffer); _ = viewPath
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(viewPath)
            stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().INSTALLPATH)
            var fromDirectoryAbPath *AbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = fromDirectoryAbPath
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(viewPath)
            stringBuffer.Append("special")
            stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
            var fromSpecialDirectoryAbPath *AbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = fromSpecialDirectoryAbPath
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(viewPath)
            stringBuffer.Append(this.GetName())
            stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
            var toDirectoryAbPath *AbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = toDirectoryAbPath
            if current == 0 {
                if !this.Directory.Create(toDirectoryAbPath) {
                    panic(errors.New(fmt.Sprint("Unable to create store view directory: " + fmt.Sprint(toDirectoryAbPath))))
                }
            }
            var viewTotal int = (total * 85) / 100; _ = viewTotal
            var installTotal int = (total * 93) / 100; _ = installTotal
            if current < viewTotal {
                var file *AbFile = AbFile_CreateAbFileFromAbPath(fromSpecialDirectoryAbPath); _ = file
                if file.IsDirectory() {
                    var halfViewTotal int = viewTotal / 2; _ = halfViewTotal
                    if current < halfViewTotal {
                        this.InstallViews(fromSpecialDirectoryAbPath, toDirectoryAbPath, current, halfViewTotal)
                    } else if current < viewTotal {
                        this.InstallViews(fromDirectoryAbPath, toDirectoryAbPath, current - halfViewTotal, viewTotal - halfViewTotal)
                    }
                } else {
                    this.InstallViews(fromDirectoryAbPath, toDirectoryAbPath, current, viewTotal)
                }
            } else if current < installTotal {
                this.InstallResources(fromDirectoryAbPath, current - viewTotal, installTotal - viewTotal)
            } else {
                this.InstallMedia(fromDirectoryAbPath, current - installTotal, total - installTotal)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *StoreFront) InstallViews(fromDirectoryAbPath *AbPath, toDirectoryAbPath *AbPath, current int, total int) {
        FileUtil_GetInstance().CopyDirectoryPortion(fromDirectoryAbPath, toDirectoryAbPath, true, false, current, total)
}


// throws Exception
func (this *StoreFront) InstallResources(fromDirectoryAbPath *AbPath, current int, total int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(fmt.Sprint(fromDirectoryAbPath))
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        stringBuffer.Append("Resources")
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        var installCategoryAbPath *AbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = installCategoryAbPath
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(URLGLOBALS.GetMainPath())
        stringBuffer.Append(this.GetCurrentHomeHostNamePath())
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        stringBuffer.Append(this.GetCategoryPath())
        var categoryAbPath *AbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = categoryAbPath
        FileUtil_GetInstance().CopyDirectoryPortion(installCategoryAbPath, categoryAbPath, true, false, current, total)
}


// throws Exception
func (this *StoreFront) InstallMedia(fromDirectoryAbPath *AbPath, current int, total int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(fmt.Sprint(fromDirectoryAbPath))
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().TEMPLATEPATH)
        stringBuffer.Append("images")
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        var viewStoreImagesDirectoryAbPath *AbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = viewStoreImagesDirectoryAbPath
        var storeAbPath *AbPath = NewAbPath(URLGLOBALS.GetWebappPath() + this.GetCurrentHomeHostNamePath(), StringUtil_GetInstance().EMPTY_STRING); _ = storeAbPath
        FileUtil_GetInstance().CopyDirectoryPortion(viewStoreImagesDirectoryAbPath, storeAbPath, true, false, current, total)
}


// throws Exception
func (this *StoreFront) ToVector() *BasicArrayList {
        var dataVector *BasicArrayList = NewBasicArrayListD(); _ = dataVector
        dataVector.Add(this.Name)
        dataVector.Add(this.HomeHostName)
        dataVector.Add(fmt.Sprint(this.HomeHostNamePath))
        dataVector.Add(this.HostName)
        dataVector.Add(fmt.Sprint(this.HostNamePath))
        dataVector.Add(this.TestHomeHostName)
        dataVector.Add(fmt.Sprint(this.TestHomeHostNamePath))
        dataVector.Add(this.TestHostName)
        dataVector.Add(fmt.Sprint(this.TestHostNamePath))
        dataVector.Add(fmt.Sprint(this.ImagePath))
        dataVector.Add(fmt.Sprint(this.StaticPath))
        dataVector.Add(fmt.Sprint(this.CategoryPath))
        dataVector.Add(this.InventoryControl)
        var contextConfigurationDomDocumentMapping *ContextConfigurationDomDocumentMapping = NewContextConfigurationDomDocumentMapping(this.GetContextConfigurationInterface()); _ = contextConfigurationDomDocumentMapping
        dataVector.Add(contextConfigurationDomDocumentMapping.ToDomDocumentString())
        dataVector.Add(this.SubStores)
        dataVector.Add(this.TagLocation)
        dataVector.Add(this.PackageLocation)
        dataVector.Add(this.Ftp)
        dataVector.Add(fmt.Sprint(this.FtpPath))
        dataVector.Add(this.FtpUserName)
        dataVector.Add(this.FtpPassword)
        dataVector.Add(this.TestFtp)
        dataVector.Add(fmt.Sprint(this.TestFtpPath))
        dataVector.Add(this.TestFtpUserName)
        dataVector.Add(this.TestFtpPassword)
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        dataVector.Add(time)
        dataVector.Add(time)
        return dataVector
}


// throws Exception
func (this *StoreFront) ToHashMap() *HashMap {
        var storeFrontData *StoreFrontData = StoreFrontData_GetInstance(); _ = storeFrontData
        var dataHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = dataHashMap
        dataHashMap.Put(storeFrontData.NAME, this.Name)
        dataHashMap.Put(storeFrontData.HOMEHOSTNAME, this.HomeHostName)
        dataHashMap.Put(storeFrontData.HOMEHOSTNAMEPATH, fmt.Sprint(this.HomeHostNamePath))
        dataHashMap.Put(storeFrontData.HOSTNAME, this.HostName)
        dataHashMap.Put(storeFrontData.HOSTNAMEPATH, fmt.Sprint(this.HostNamePath))
        dataHashMap.Put(storeFrontData.TESTHOMEHOSTNAME, this.TestHomeHostName)
        dataHashMap.Put(storeFrontData.TESTHOMEHOSTNAMEPATH, fmt.Sprint(this.TestHomeHostNamePath))
        dataHashMap.Put(storeFrontData.TESTHOSTNAME, this.TestHostName)
        dataHashMap.Put(storeFrontData.TESTHOSTNAMEPATH, fmt.Sprint(this.TestHostNamePath))
        dataHashMap.Put(storeFrontData.IMAGEPATH, fmt.Sprint(this.ImagePath))
        dataHashMap.Put(storeFrontData.STATICPATH, fmt.Sprint(this.StaticPath))
        dataHashMap.Put(storeFrontData.CATEGORYPATH, fmt.Sprint(this.CategoryPath))
        dataHashMap.Put(storeFrontData.INVENTORYCONTROL, this.InventoryControl)
        var contextConfigurationDomDocumentMapping *ContextConfigurationDomDocumentMapping = NewContextConfigurationDomDocumentMapping(this.GetContextConfigurationInterface()); _ = contextConfigurationDomDocumentMapping
        dataHashMap.Put(storeFrontData.CONFIGURATION, contextConfigurationDomDocumentMapping.ToDomDocumentString())
        dataHashMap.Put(storeFrontData.SUBSTORES, this.SubStores)
        dataHashMap.Put(storeFrontData.TAGLOCATION, this.TagLocation)
        dataHashMap.Put(storeFrontData.PACKAGELOCATION, this.PackageLocation)
        dataHashMap.Put(storeFrontData.FTP, this.Ftp)
        dataHashMap.Put(storeFrontData.FTPPATH, fmt.Sprint(this.FtpPath))
        dataHashMap.Put(storeFrontData.FTPUSERNAME, this.FtpUserName)
        dataHashMap.Put(storeFrontData.FTPPASSWORD, this.FtpPassword)
        dataHashMap.Put(storeFrontData.TESTFTP, this.TestFtp)
        dataHashMap.Put(storeFrontData.TESTFTPPATH, fmt.Sprint(this.TestFtpPath))
        dataHashMap.Put(storeFrontData.TESTFTPUSERNAME, this.TestFtpUserName)
        dataHashMap.Put(storeFrontData.TESTFTPPASSWORD, this.TestFtpPassword)
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        dataHashMap.Put(EntryData_GetInstance().LASTMODIFIED, time)
        return dataHashMap
}


func (this *StoreFront) GetKey() any {
        return any(this.GetName()).(any)
}


func (this *StoreFront) GetContextConfigurationInterface() ContextConfigurationInterface {
        return this.ContextConfigurationInterface
}


func (this *StoreFront) SetContextConfigurationInterface(contextConfigurationInterface ContextConfigurationInterface) {
        this.ContextConfigurationInterface = contextConfigurationInterface
}

