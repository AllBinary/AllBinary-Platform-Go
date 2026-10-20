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

type UserTag struct {
    TableTag
    Role string
    Enable string
}


func (this *UserTag) initFields() {
    
    
}


func NewUserTag() *UserTag {
    this := &UserTag{}
    this.TableTag = *NewTableTag()
    this.initFields()
        this.SetTagHelperFactory(NewUserHelperFactory())
        this.SetTagRequestHelperFactory(NewUserHelperFactory())
    return this
}


func (this *UserTag) SetRole(value string) {
        this.Role = value
        this.GetPropertiesHashMap().Put(UserRoleData_NAME, this.Role)
}


func (this *UserTag) SetEnable(value string) {
        this.Enable = value
        this.GetPropertiesHashMap().Put(EntryData_GetInstance().ENABLE, this.Enable)
}


// throws JspTagException
func (this *UserTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, TableTag_SKIP_BODY; return
                }
            }()
            if this.IsEnabled() {
                tryControl1, tryResult1 = 1, this.TableTag.DoStartTag(); return
            }
            tryControl1, tryResult1 = 1, TableTag_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

