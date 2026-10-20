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

type AuthenticationHelperUtil struct {
    LogUtil *LogUtil
}

var AuthenticationHelperUtil_Instance *AuthenticationHelperUtil = NewAuthenticationHelperUtil()

func (this *AuthenticationHelperUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewAuthenticationHelperUtil() *AuthenticationHelperUtil {
    this := &AuthenticationHelperUtil{}
    this.initFields()
    return this
}


func AuthenticationHelperUtil_GetInstance() *AuthenticationHelperUtil {
        return AuthenticationHelperUtil_Instance
}


// throws Exception
func (this *AuthenticationHelperUtil) IsAuthorized(authenticationHelper *AuthenticationHelper, filePath string) bool {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
            this.LogUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().START) + filePath, this, "isAuthorized()")
        }
        var endIndex int = HttpRequestUtil_GetInstance().GetLastSeparatorIndex(filePath); _ = endIndex
        var fullPath IAbPath = NewAbPath(fmt.Sprint(URLGLOBALS.GetWebappPath()) + filePath[0:endIndex], StringUtil_GetInstance().EMPTY_STRING); _ = fullPath
        var downloadFilePermissions *HttpFilePermissions = NewHttpFilePermissions(fullPath); _ = downloadFilePermissions
        var userRoles IBasicArrayList = downloadFilePermissions.GetUserRoles(); _ = userRoles
        var userRole IUserRole = authenticationHelper.GetRole(); _ = userRole
        var basicUserRole *BasicUserRole = userRole.GetBasicUserRole(); _ = basicUserRole
        if userRoles.Contains(basicUserRole) {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Authorized", this, "isAuthorized()")
            }
            return true
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Not Authorized: " + fmt.Sprint(basicUserRole), this, "isAuthorized()")
            }
            return false
        }
}

