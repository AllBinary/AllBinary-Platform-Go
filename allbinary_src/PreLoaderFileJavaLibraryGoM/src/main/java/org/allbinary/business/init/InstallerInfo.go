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

type InstallerInfo struct {
    LogUtil *LogUtil
}

var InstallerInfo_Instance *InstallerInfo = NewInstallerInfo()
var InstallerInfo_INITFILENAME string = "installerdata.dat"
var InstallerInfo_PACKAGE string = PATH_GLOBALS.GetInstance().INIT_PATH
var InstallerInfo_UserName string = StringUtil_GetInstance().EMPTY_STRING
var InstallerInfo_Password string = StringUtil_GetInstance().EMPTY_STRING
var InstallerInfo_HasRead bool = false

func (this *InstallerInfo) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewInstallerInfo() *InstallerInfo {
    this := &InstallerInfo{}
    this.initFields()
    return this
}


// throws Exception
// synchronized
func (this *InstallerInfo) Write() {
        var FILEABPATH IAbPath = NewAbPath(URLGLOBALS.GetMainPath() + InstallerInfo_PACKAGE, InstallerInfo_INITFILENAME); _ = FILEABPATH
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRELOADERERROR) {
                this.LogUtil.PutF("Failed", this, "write")
            }
            panic(e)
                }
            }()
            var newFile *AbFile = AbFile_CreateAbFileFromAbPath(FILEABPATH); _ = newFile
            newFile.CreateNewFile()
            var dataOutputStream *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(newFile); _ = dataOutputStream
            var cryptedUserName []int8 = Java_GetBytes(NewWeakCrypt(1).Encrypt(this.GetUserName())); _ = cryptedUserName
            var cryptedPassword []int8 = Java_GetBytes(NewWeakCrypt(2).Encrypt(this.GetPassword())); _ = cryptedPassword
            dataOutputStream.WriteUTF(DatabaseEncoder_Encode(cryptedUserName))
            dataOutputStream.WriteUTF(DatabaseEncoder_Encode(cryptedPassword))
            InstallerInfo_HasRead = false
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
// synchronized
func (this *InstallerInfo) Read() {
        var FILEABPATH IAbPath = NewAbPath(URLGLOBALS.GetMainPath() + InstallerInfo_PACKAGE, InstallerInfo_INITFILENAME); _ = FILEABPATH
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRELOADERERROR) {
                this.LogUtil.PutF("Failed", this, "read")
            }
                }
            }()
            var file *AbFile = AbFile_CreateAbFileFromAbPath(FILEABPATH); _ = file
            if file.IsFile() {
                var iFile *AbFileLocalInputStream = NewAbFileLocalInputStream(file); _ = iFile
                var iData *AbDataInputStream = NewAbDataInputStream(iFile); _ = iData
                var decryptedUserName string = DatabaseEncoder_Decode(iData.ReadUTF()); _ = decryptedUserName
                var decryptedPassword string = DatabaseEncoder_Decode(iData.ReadUTF()); _ = decryptedPassword
                this.SetUserName(NewWeakCrypt(1).Decrypt(decryptedUserName))
                this.SetPassword(NewWeakCrypt(2).Decrypt(decryptedPassword))
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRELOADER) {
                    this.LogUtil.PutF("Not a File - Failed Loading: " + fmt.Sprint(FILEABPATH), this, "read")
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func InstallerInfo_SetHasRead(value bool) {
        InstallerInfo_HasRead = value
}


func (this *InstallerInfo) SetUserName(userName string) {
        InstallerInfo_UserName = userName
}


func (this *InstallerInfo) SetPassword(password string) {
        InstallerInfo_Password = password
}


// throws Exception
// synchronized
func (this *InstallerInfo) UpdateIfNeeded() {
        if !InstallerInfo_HasRead {
            InstallerInfo_HasRead = true
            this.Read()
            if InstallerInfo_UserName == nil || InstallerInfo_Password == nil {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRELOADER) {
                    this.LogUtil.PutF("Failed", this, "updateIfNeeded")
                }
            }
        }
}


// throws Exception
func (this *InstallerInfo) GetUserName() string {
        this.UpdateIfNeeded()
        return InstallerInfo_UserName
}


// throws Exception
func (this *InstallerInfo) GetPassword() string {
        this.UpdateIfNeeded()
        return InstallerInfo_Password
}


// throws Exception
func (this *InstallerInfo) IsValid(userName string, password string) bool {
        this.UpdateIfNeeded()
        if InstallerInfo_UserName != "" && strings.Compare(InstallerInfo_UserName, userName) == 0 && InstallerInfo_Password != "" && strings.Compare(InstallerInfo_Password, password) == 0 {
            return true
        }
        return false
}

