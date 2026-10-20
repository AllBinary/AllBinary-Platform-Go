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

type BasicStoreFront struct {
    Name string
    UserName string
    BasketName string
    HomeHostName string
    HomeHostNamePath string
    HostName string
    HostNamePath string
    TestHomeHostName string
    TestHomeHostNamePath string
    TestHostName string
    TestHostNamePath string
    ImagePath string
    StaticPath string
    CategoryPath string
    InventoryControl string
    SubStores string
    TagLocation string
    PackageLocation string
    Ftp string
    FtpPath string
    FtpUserName string
    FtpPassword string
    TestFtp string
    TestFtpPath string
    TestFtpUserName string
    TestFtpPassword string
    TimeCreated string
    LastModified string
}

// Implements BasicStoreFrontInterface


func (this *BasicStoreFront) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewBasicStoreFront() *BasicStoreFront {
    this := &BasicStoreFront{}
    this.initFields()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewBasicStoreFront(storeHashMap *HashMap) *BasicStoreFront {
    this := &BasicStoreFront{}
    this.initFields()
        var storeFrontData *StoreFrontData = StoreFrontData_GetInstance(); _ = storeFrontData
        this.Name = any(storeHashMap.Get(storeFrontData.NAME)).(string)
        this.UserName = any(storeHashMap.Get(UserData_USERNAME)).(string)
        this.BasketName = any(storeHashMap.Get(storeFrontData.NAME)).(string)
        this.HomeHostName = any(storeHashMap.Get(storeFrontData.HOMEHOSTNAME)).(string)
        this.HomeHostNamePath = any(storeHashMap.Get(storeFrontData.HOMEHOSTNAMEPATH)).(string)
        this.HostName = any(storeHashMap.Get(storeFrontData.HOSTNAME)).(string)
        this.HostNamePath = any(storeHashMap.Get(storeFrontData.HOSTNAMEPATH)).(string)
        this.TestHomeHostName = any(storeHashMap.Get(storeFrontData.TESTHOMEHOSTNAME)).(string)
        this.TestHomeHostNamePath = any(storeHashMap.Get(storeFrontData.TESTHOMEHOSTNAMEPATH)).(string)
        this.TestHostName = any(storeHashMap.Get(storeFrontData.TESTHOSTNAME)).(string)
        this.TestHostNamePath = any(storeHashMap.Get(storeFrontData.TESTHOSTNAMEPATH)).(string)
        this.ImagePath = any(storeHashMap.Get(storeFrontData.IMAGEPATH)).(string)
        this.StaticPath = any(storeHashMap.Get(storeFrontData.STATICPATH)).(string)
        this.CategoryPath = any(storeHashMap.Get(storeFrontData.CATEGORYPATH)).(string)
        this.InventoryControl = any(storeHashMap.Get(storeFrontData.INVENTORYCONTROL)).(string)
        this.SubStores = any(storeHashMap.Get(storeFrontData.SUBSTORES)).(string)
        this.TagLocation = any(storeHashMap.Get(storeFrontData.TAGLOCATION)).(string)
        this.PackageLocation = any(storeHashMap.Get(storeFrontData.PACKAGELOCATION)).(string)
        this.Ftp = any(storeHashMap.Get(storeFrontData.FTP)).(string)
        this.FtpPath = any(storeHashMap.Get(storeFrontData.FTPPATH)).(string)
        this.FtpUserName = any(storeHashMap.Get(storeFrontData.FTPUSERNAME)).(string)
        this.FtpPassword = any(storeHashMap.Get(storeFrontData.FTPPASSWORD)).(string)
        this.TestFtp = any(storeHashMap.Get(storeFrontData.TESTFTP)).(string)
        this.TestFtpPath = any(storeHashMap.Get(storeFrontData.TESTFTPPATH)).(string)
        this.TestFtpUserName = any(storeHashMap.Get(storeFrontData.TESTFTPUSERNAME)).(string)
        this.TestFtpPassword = any(storeHashMap.Get(storeFrontData.TESTFTPPASSWORD)).(string)
        this.TimeCreated = any(storeHashMap.Get(EntryData_GetInstance().TIMECREATED)).(string)
        this.LastModified = any(storeHashMap.Get(EntryData_GetInstance().LASTMODIFIED)).(string)
    return this
}


func (this *BasicStoreFront) GetTestHtmlPath() string {
        return org_Allbinary.Globals.URLGLOBALS.GetTestHtmlPath() + this.GetCurrentHostNamePath()
}


func (this *BasicStoreFront) GetCurrentHostName() string {
        var location string = nil; _ = location
        if org_Allbinary.Globals.URLGLOBALS.IsTestingMode() {
            location = this.GetTestHostName()
        } else {
            location = this.GetHostName()
        }
        return location
}


func (this *BasicStoreFront) GetCurrentHostNamePath() string {
        var location string = nil; _ = location
        if org_Allbinary.Globals.URLGLOBALS.IsTestingMode() {
            location = this.GetTestHostNamePath()
        } else {
            location = this.GetHostNamePath()
        }
        return location
}


func (this *BasicStoreFront) GetCurrentHomeHostName() string {
        var location string = nil; _ = location
        if org_Allbinary.Globals.URLGLOBALS.IsTestingMode() {
            location = this.GetTestHomeHostName()
        } else {
            location = this.GetHomeHostName()
        }
        return location
}


func (this *BasicStoreFront) GetCurrentHomeHostNamePath() string {
        var location string = nil; _ = location
        if org_Allbinary.Globals.URLGLOBALS.IsTestingMode() {
            location = this.GetTestHomeHostNamePath()
        } else {
            location = this.GetHomeHostNamePath()
        }
        return location
}


func (this *BasicStoreFront) GetName() string {
        return this.Name
}


func (this *BasicStoreFront) GetUserName() string {
        return this.UserName
}


func (this *BasicStoreFront) GetBasketName() string {
        return this.BasketName
}


func (this *BasicStoreFront) GetHomeHostName() string {
        return this.HomeHostName
}


func (this *BasicStoreFront) GetHomeHostNamePath() string {
        return this.HomeHostNamePath
}


func (this *BasicStoreFront) GetHostName() string {
        return this.HostName
}


func (this *BasicStoreFront) GetHostNamePath() string {
        return this.HostNamePath
}


func (this *BasicStoreFront) GetTestHomeHostName() string {
        return this.TestHomeHostName
}


func (this *BasicStoreFront) GetTestHomeHostNamePath() string {
        return this.TestHomeHostNamePath
}


func (this *BasicStoreFront) GetTestHostName() string {
        return this.TestHostName
}


func (this *BasicStoreFront) GetTestHostNamePath() string {
        return this.TestHostNamePath
}


func (this *BasicStoreFront) GetStaticPath() string {
        return this.StaticPath
}


func (this *BasicStoreFront) GetCategoryPath() string {
        return this.CategoryPath
}


func (this *BasicStoreFront) GetInventoryControl() string {
        return this.InventoryControl
}


// throws Exception
func (this *BasicStoreFront) GetSubStores() *BasicArrayList {
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


func (this *BasicStoreFront) GetTagLocation() string {
        return this.TagLocation
}


func (this *BasicStoreFront) GetPackageLocation() string {
        return this.PackageLocation
}


func (this *BasicStoreFront) GetFtp() string {
        return this.Ftp
}


func (this *BasicStoreFront) GetFtpUserName() string {
        return this.FtpUserName
}


func (this *BasicStoreFront) GetFtpPassword() string {
        return this.FtpPassword
}


func (this *BasicStoreFront) GetTestFtp() string {
        return this.TestFtp
}


func (this *BasicStoreFront) GetTestFtpUserName() string {
        return this.TestFtpUserName
}


func (this *BasicStoreFront) GetTestFtpPassword() string {
        return this.TestFtpPassword
}


func (this *BasicStoreFront) GetFtpPath() string {
        return this.FtpPath
}


func (this *BasicStoreFront) GetTestFtpPath() string {
        return this.TestFtpPath
}


func (this *BasicStoreFront) GetTimeCreated() string {
        return this.TimeCreated
}


func (this *BasicStoreFront) GetLastModified() string {
        return this.LastModified
}

