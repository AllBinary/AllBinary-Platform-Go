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

type StoreFrontData struct {
    CANCEL string
    SELECT string
    INSTALL string
    INSTALL_COMPLETE string
    NAME string
    CURRENTHOSTNAME string
    CURRENTHOSTNAMEPATH string
    CURRENTHOMEHOSTNAME string
    CURRENTHOMEHOSTNAMEPATH string
    HOMEHOSTNAME string
    HOMEHOSTNAMEPATH string
    HOSTNAME string
    HOSTNAMEPATH string
    TESTHOMEHOSTNAME string
    TESTHOMEHOSTNAMEPATH string
    TESTHOSTNAME string
    TESTHOSTNAMEPATH string
    IMAGEPATH string
    STATICPATH string
    CATEGORYPATH string
    INVENTORYCONTROL string
    CONFIGURATION string
    SUBSTORES string
    TAGLOCATION string
    PACKAGELOCATION string
    FTP string
    FTPPATH string
    FTPUSERNAME string
    FTPPASSWORD string
    TESTFTP string
    TESTFTPPATH string
    TESTFTPUSERNAME string
    TESTFTPPASSWORD string
    SELECTSTORENAME string
}

var StoreFrontData_Instance *StoreFrontData = NewStoreFrontData()

func (this *StoreFrontData) initFields() {
    this.CANCEL = "Cancel Store Front"
    this.SELECT = "Select Store Front"
    this.INSTALL = "Install Store Front"
    this.INSTALL_COMPLETE = "Install Store Front Complete"
    this.NAME = "STOREFRONT_NAME"
    this.CURRENTHOSTNAME = "STOREFRONT_CURRENTHOSTNAME"
    this.CURRENTHOSTNAMEPATH = "STOREFRONT_CURRENTHOSTNAMEPATH"
    this.CURRENTHOMEHOSTNAME = "STOREFRONT_CURRENTHOMEHOSTNAME"
    this.CURRENTHOMEHOSTNAMEPATH = "STOREFRONT_CURRENTHOMEHOSTNAMEPATH"
    this.HOMEHOSTNAME = "STOREFRONT_HOMEHOSTNAME"
    this.HOMEHOSTNAMEPATH = "STOREFRONT_HOMEHOSTNAMEPATH"
    this.HOSTNAME = "STOREFRONT_HOSTNAME"
    this.HOSTNAMEPATH = "STOREFRONT_HOSTNAMEPATH"
    this.TESTHOMEHOSTNAME = "STOREFRONT_TESTHOMEHOSTNAME"
    this.TESTHOMEHOSTNAMEPATH = "STOREFRONT_TESTHOMEHOSTNAMEPATH"
    this.TESTHOSTNAME = "STOREFRONT_TESTHOSTNAME"
    this.TESTHOSTNAMEPATH = "STOREFRONT_TESTHOSTNAMEPATH"
    this.IMAGEPATH = "STOREFRONT_IMAGEPATH"
    this.STATICPATH = "STOREFRONT_STATICPATH"
    this.CATEGORYPATH = "STOREFRONT_CATEGORYPATH"
    this.INVENTORYCONTROL = "STOREFRONT_INVENTORYCONTROL"
    this.CONFIGURATION = "STOREFRONT_CONFIGURATION"
    this.SUBSTORES = "STOREFRONT_SUBSTORES"
    this.TAGLOCATION = "STOREFRONT_TAGLOCATION"
    this.PACKAGELOCATION = "STOREFRONT_PACKAGELOCATION"
    this.FTP = "STOREFRONT_FTP"
    this.FTPPATH = "STOREFRONT_FTPPATH"
    this.FTPUSERNAME = "STOREFRONT_FTPUSERNAME"
    this.FTPPASSWORD = "STOREFRONT_FTPPASSWORD"
    this.TESTFTP = "STOREFRONT_TESTFTP"
    this.TESTFTPPATH = "STOREFRONT_TESTFTPPATH"
    this.TESTFTPUSERNAME = "STOREFRONT_TESTFTPUSERNAME"
    this.TESTFTPPASSWORD = "STOREFRONT_TESTFTPPASSWORD"
    this.SELECTSTORENAME = "STOREFRONT_SELECTSTORENAME"
}

func NewStoreFrontData() *StoreFrontData {
    this := &StoreFrontData{}
    this.initFields()
    return this
}


func StoreFrontData_GetInstance() *StoreFrontData {
        return StoreFrontData_Instance
}

