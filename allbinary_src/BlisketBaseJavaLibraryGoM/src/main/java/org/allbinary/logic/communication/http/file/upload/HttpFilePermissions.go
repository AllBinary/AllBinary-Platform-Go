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

type HttpFilePermissions struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    File *AbFile
    UserRoles IBasicArrayList
}


func (this *HttpFilePermissions) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    this.UserRoles = NewBasicArrayListD()
}


func NewHttpFilePermissions(fullPath IAbPath) *HttpFilePermissions {
    this := &HttpFilePermissions{}
    this.initFields()
        this.File = AbFile_CreateAbFileFromAbPath(NewAbPath(fullPath.ToFileSystemString(), FileUploadData_GetInstance().FILE))
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
            this.LogUtil.PutF("Loading Privileges: " + fmt.Sprint(fullPath.ToFileSystemString()), this, this.CommonStrings.CONSTRUCTOR)
        }
        if this.IsFile() {
            this.UserRoles.Add(BasicUserRoleFactory_GetInstance().ADMINISTRATOR)
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
            this.LogUtil.PutF("User Roles: " + fmt.Sprint(this.UserRoles), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func (this *HttpFilePermissions) IsFile() bool {
        return this.File.IsFile()
}


func (this *HttpFilePermissions) GetUserRoles() IBasicArrayList {
        return this.UserRoles
}

