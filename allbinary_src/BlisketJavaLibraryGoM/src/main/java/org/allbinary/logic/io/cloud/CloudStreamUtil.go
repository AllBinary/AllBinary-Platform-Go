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

type CloudStreamUtil struct {
    LogUtil *LogUtil
}

var CloudStreamUtil_Instance *CloudStreamUtil = NewCloudStreamUtil()

func (this *CloudStreamUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewCloudStreamUtil() *CloudStreamUtil {
    this := &CloudStreamUtil{}
    this.initFields()
    return this
}


func CloudStreamUtil_GetInstance() *CloudStreamUtil {
        return CloudStreamUtil_Instance
}


// throws Exception
func (this *CloudStreamUtil) GetFile(file *AbFile) IInputStream {
        var inputStream IInputStream = NewAbFileInputStream(file); _ = inputStream
        var available int = inputStream.Available(); _ = available
        if available > 0 {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Cloud File Bytes: " + fmt.Sprint(available), this, "processRequest()")
            }
        }
        return inputStream
}


// throws Exception
func (this *CloudStreamUtil) GetFileAnyWhere(file *AbFile) IInputStream {
        var inputStream IInputStream = NewAbFileLocalInputStream(file); _ = inputStream
        var available int = inputStream.Available(); _ = available
        if available > 0 {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Local File Bytes: " + fmt.Sprint(available), this, "processRequest()")
            }
        } else {
            inputStream = NewAbFileInputStream(file)
            available = inputStream.Available()
            if available > 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                    this.LogUtil.PutF("Cloud File Bytes: " + fmt.Sprint(available), this, "processRequest()")
                }
            }
        }
        return inputStream
}


// throws Exception
func (this *CloudStreamUtil) GetFileLocal(file *AbFile) IInputStream {
        var inputStream IInputStream = NewAbFileLocalInputStream(file); _ = inputStream
        var available int = inputStream.Available(); _ = available
        if available > 0 {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Local File Bytes: " + fmt.Sprint(available), this, "processRequest()")
            }
        }
        return inputStream
}

