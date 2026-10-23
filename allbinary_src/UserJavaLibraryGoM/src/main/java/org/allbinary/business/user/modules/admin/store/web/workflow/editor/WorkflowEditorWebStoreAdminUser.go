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

type WorkflowEditorWebStoreAdminUser struct {
    User
}





// Java overloads of new WorkflowEditorWebStoreAdminUser
func NewWorkflowEditorWebStoreAdminUser(args ...any) *WorkflowEditorWebStoreAdminUser {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewWorkflowEditorWebStoreAdminUser_1()
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewWorkflowEditorWebStoreAdminUser_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new WorkflowEditorWebStoreAdminUser for the arguments")
}
func NewWorkflowEditorWebStoreAdminUser_1() *WorkflowEditorWebStoreAdminUser {
    this := &WorkflowEditorWebStoreAdminUser{}
        this.User = *NewUser()
    return this
}


func NewWorkflowEditorWebStoreAdminUser_2(userHashMap *HashMap) *WorkflowEditorWebStoreAdminUser {
    this := &WorkflowEditorWebStoreAdminUser{}
        this.User = *NewUser(userHashMap)
    return this
}


func (this *WorkflowEditorWebStoreAdminUser) ValidateSession(weblisketSession WeblisketSessionInterface) {
        this.User.ValidateSession(weblisketSession)
        this.UpdateSession(weblisketSession)
}


func (this *WorkflowEditorWebStoreAdminUser) UpdateSession(weblisketSession WeblisketSessionInterface) {
        weblisketSession.SetStoreName(this.GetPermissions())
}

