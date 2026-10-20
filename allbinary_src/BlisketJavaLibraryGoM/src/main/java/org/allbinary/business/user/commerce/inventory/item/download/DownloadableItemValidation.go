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

type DownloadableItemValidation struct {
    Validation
    LogUtil *LogUtil
    DownloadableItem *DownloadableItem
}


func (this *DownloadableItemValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewDownloadableItemValidation(downloadableItem *DownloadableItem) *DownloadableItemValidation {
    this := &DownloadableItemValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.DownloadableItem = downloadableItem
    return this
}


func (this *DownloadableItemValidation) IsValid() bool {
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
            if this.DownloadableItem.GetId() == nil || len(this.DownloadableItem.GetId()) < 1 || len(this.DownloadableItem.GetId()) > AbSqlData_MAXNUM {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.DownloadableItem.GetId() != nil && !StringValidationUtil_GetInstance().IsNumber(this.DownloadableItem.GetId()) {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.DownloadableItem.GetSystem() == nil || len(this.DownloadableItem.GetSystem()) < 1 || len(this.DownloadableItem.GetSystem()) > AbSqlData_MAXSTRING {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.DownloadableItem.GetPlatform() == nil || len(this.DownloadableItem.GetPlatform()) < 1 || len(this.DownloadableItem.GetPlatform()) > AbSqlData_MAXSTRING {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.DownloadableItem.GetSpecialName() == nil || len(this.DownloadableItem.GetSpecialName()) < 0 || len(this.DownloadableItem.GetSpecialName()) > AbSqlData_MAXSTRING {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.DownloadableItem.GetVersion() == nil || len(this.DownloadableItem.GetVersion()) < 1 || len(this.DownloadableItem.GetVersion()) > AbSqlData_MAXSTRING {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.DownloadableItem.GetChanges() == nil || len(this.DownloadableItem.GetChanges()) < 1 || len(this.DownloadableItem.GetChanges()) > AbSqlData_MAXBLOB {
                tryControl1, tryResult1 = 1, false; return
            }
            if this.DownloadableItem.GetSize() == nil {
                tryControl1, tryResult1 = 1, false; return
            } else {
                var size string = fmt.Sprint(this.DownloadableItem.GetSize()); _ = size
                if len(size) < 1 || len(size) > AbSqlData_MAXNUM {
                    tryControl1, tryResult1 = 1, false; return
                }
            }
            if this.DownloadableItem.GetValidTime() == nil {
                tryControl1, tryResult1 = 1, false; return
            } else {
                var size string = fmt.Sprint(this.DownloadableItem.GetValidTime()); _ = size
                if len(size) < 1 || len(size) > AbSqlData_MAXNUM {
                    tryControl1, tryResult1 = 1, false; return
                }
            }
            if this.DownloadableItem.GetRetries() == nil {
                tryControl1, tryResult1 = 1, false; return
            } else {
                var size string = this.DownloadableItem.GetRetries(); _ = size
                if len(size) < 1 || len(size) > AbSqlData_MAXNUM {
                    tryControl1, tryResult1 = 1, false; return
                }
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *DownloadableItemValidation) ValidationInfo() string {
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
            if this.DownloadableItem.GetId() == nil || len(this.DownloadableItem.GetId()) < 1 || len(this.DownloadableItem.GetId()) > AbSqlData_MAXNUM {
                stringBuffer.Append("Id is invalid. Must be number < ")
                stringBuffer.Appendint(AbSqlData_MAXNUM)
                stringBuffer.Append(" and > 0 digits.<br/>")
            }
            if this.DownloadableItem.GetId() != nil && !StringValidationUtil_GetInstance().IsNumber(this.DownloadableItem.GetId()) {
                stringBuffer.Append("Id is invalid. Must be > 0.<br/>")
            }
            if this.DownloadableItem.GetSystem() == nil || len(this.DownloadableItem.GetSystem()) < 1 || len(this.DownloadableItem.GetSystem()) > AbSqlData_MAXSTRING {
                stringBuffer.Append("System value is invalid. Must be < ")
                stringBuffer.Appendint(AbSqlData_MAXSTRING)
                stringBuffer.Append(" and > 0 characters long.<br/>")
            }
            if this.DownloadableItem.GetPlatform() == nil || len(this.DownloadableItem.GetPlatform()) < 1 || len(this.DownloadableItem.GetPlatform()) > AbSqlData_MAXSTRING {
                stringBuffer.Append("Platform value is invalid. Must be < ")
                stringBuffer.Appendint(AbSqlData_MAXSTRING)
                stringBuffer.Append(" and > 0 characters long.<br/>")
            }
            if this.DownloadableItem.GetSpecialName() == nil || len(this.DownloadableItem.GetSpecialName()) < 0 || len(this.DownloadableItem.GetSpecialName()) > AbSqlData_MAXSTRING {
                stringBuffer.Append("Special Name value is invalid. Must be < ")
                stringBuffer.Appendint(AbSqlData_MAXSTRING)
                stringBuffer.Append(" and > 0 characters long.<br/>")
            }
            if this.DownloadableItem.GetVersion() == nil || len(this.DownloadableItem.GetVersion()) < 1 || len(this.DownloadableItem.GetVersion()) > AbSqlData_MAXSTRING {
                stringBuffer.Append("Version value is invalid. Must be < ")
                stringBuffer.Appendint(AbSqlData_MAXSTRING)
                stringBuffer.Append(" and > 0 characters long.<br/>")
            }
            if this.DownloadableItem.GetChanges() == nil || len(this.DownloadableItem.GetChanges()) < 0 || len(this.DownloadableItem.GetChanges()) > AbSqlData_MAXBLOB {
                stringBuffer.Append("Changes value is invalid. Must be < ")
                stringBuffer.Appendint(AbSqlData_MAXBLOB)
                stringBuffer.Append(" and >= 0 characters long.<br/>")
            }
            if this.DownloadableItem.GetSize() == nil {
                stringBuffer.Append("File Size is null")
            } else {
                var size string = fmt.Sprint(this.DownloadableItem.GetSize()); _ = size
                if len(size) < 1 || len(size) > AbSqlData_MAXNUM {
                    stringBuffer.Append("File Size is invalid. Must be number < ")
                    stringBuffer.Appendint(AbSqlData_MAXNUM)
                    stringBuffer.Append(" and > 0 digits.<br/>")
                }
            }
            if this.DownloadableItem.GetValidTime() == nil {
                stringBuffer.Append("Valid Time is null")
            } else {
                var size string = fmt.Sprint(this.DownloadableItem.GetValidTime()); _ = size
                if len(size) < 1 || len(size) > AbSqlData_MAXNUM {
                    stringBuffer.Append("Valid Time is invalid. Must be number < ")
                    stringBuffer.Appendint(AbSqlData_MAXNUM)
                    stringBuffer.Append(" and > 0 digits.<br/>")
                }
            }
            if this.DownloadableItem.GetRetries() == nil {
                stringBuffer.Append("Retries is null")
            } else {
                var size string = this.DownloadableItem.GetRetries(); _ = size
                if len(size) < 1 || len(size) > AbSqlData_MAXNUM {
                    stringBuffer.Append("Retries is invalid. Must be number < ")
                    stringBuffer.Appendint(AbSqlData_MAXNUM)
                    stringBuffer.Append(" and > 0 digits.<br/>")
                }
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *DownloadableItemValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *DownloadableItemValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}

