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

type WorkFlowsTag struct {
    StoreValidationTransformTag
}



func NewWorkFlowsTag() *WorkFlowsTag {
    this := &WorkFlowsTag{}
    this.StoreValidationTransformTag = *NewStoreValidationTransformTag()
    return this
}


// throws JspTagException
func (this *WorkFlowsTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, StoreValidationTransformTag_SKIP_BODY; return
                }
            }()
            if this.GetCommand() != nil {
                if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.NEW) == 0 {
                    this.SetName("New WorkFlow View")
                    this.SetObjectFile("views.admin.workflow.NewView")
                } else if strings.Compare(this.GetCommand(), this.CommonStrings.INSERT) == 0 {
                    this.SetName("Add Validation WorkFlow View")
                    this.SetObjectFile("views.admin.workflow.AddValidationView")
                } else if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.VIEW) == 0 {
                    this.SetName("Validation WorkFlow View")
                    this.SetObjectFile("views.admin.workflow.WorkflowsValidationView")
                } else if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.EDIT) == 0 {
                    this.SetName("Validation WorkFlow View")
                    this.SetObjectFile("views.admin.workflow.EditValidationView")
                } else if strings.Compare(this.GetCommand(), this.CommonStrings.UPDATE) == 0 {
                    this.SetName("Update Validation WorkFlow View")
                    this.SetObjectFile("views.admin.workflow.UpdateValidationView")
                } else if strings.Compare(this.GetCommand(), this.CommonStrings.DELETE) == 0 {
                    this.SetName("Delete Validation WorkFlow View")
                    this.SetObjectFile("views.admin.workflow.DeleteValidationView")
                } else if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.BACKUP) == 0 {
                    panic(errors.New(fmt.Sprint("View not implemented")))
                } else if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.RESTORE) == 0 {
                    panic(errors.New(fmt.Sprint("View not implemented")))
                } else if strings.Compare(this.GetCommand(), this.CommonStrings.DROP) == 0 {
                    panic(errors.New(fmt.Sprint("View not implemented")))
                } else if strings.Compare(this.GetCommand(), this.CommonStrings.CREATE) == 0 {
                    panic(errors.New(fmt.Sprint("View not implemented")))
                } else {
                    panic(errors.New(fmt.Sprint("No Such View Command: " + fmt.Sprint(this.GetCommand()))))
                }
                tryControl1, tryResult1 = 1, this.StoreValidationTransformTag.DoStartTag(); return
            }
            panic(errors.New(fmt.Sprint("Command Null")))
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

