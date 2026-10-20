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

type InventoryTag struct {
    TableTag
    StoreName string
}


func (this *InventoryTag) initFields() {
    
}


func NewInventoryTag() *InventoryTag {
    this := &InventoryTag{}
        this.TableTag = *NewTableTag()
        this.initFields()
        this.SetTagHelperFactory(NewInventoryHelperFactory())
        this.SetTagRequestHelperFactory(NewInventoryRequestHelperFactory())
    return this
}


func (this *InventoryTag) SetStoreName(value string) {
        this.StoreName = value
}


// throws JspTagException
func (this *InventoryTag) DoStartTag() int {
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
                if this.GetCommand() != nil {
                    if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.BACKUP) != 0 && strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.RESTORE) != 0 && strings.Compare(this.GetCommand(), this.CommonStrings.CREATE) != 0 && strings.Compare(this.GetCommand(), this.CommonStrings.DROP) != 0 {
                        var parentTag *Tag = this.GetParent(); _ = parentTag
                        ParentInventoryTagHelper_GetInstance().IsValid(this, parentTag)
                        this.GetPropertiesHashMap().Put(AbTagData_PARENT, parentTag)
                        this.GetPropertiesHashMap().Put(StoreFrontData_GetInstance().NAME, this.StoreName)
                    }
                    tryControl1, tryResult1 = 1, this.TableTag.DoStartTag(); return
                }
            }
            tryControl1, tryResult1 = 1, TableTag_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

