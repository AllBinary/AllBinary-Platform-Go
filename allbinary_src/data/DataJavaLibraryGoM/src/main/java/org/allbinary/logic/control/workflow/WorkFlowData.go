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

type WorkFlowData struct {
    WORKFLOWS string
    WORKFLOW string
    NAME string
    DATA string
}

var WorkFlowData_Instance *WorkFlowData = NewWorkFlowData()

func (this *WorkFlowData) initFields() {
    this.WORKFLOWS = "WORKFLOWS"
    this.WORKFLOW = "WORKFLOW"
    this.NAME = "WORKFLOW_NAME"
    this.DATA = "WORKFLOW_DATA"
}


func WorkFlowData_GetInstance() *WorkFlowData {
        return WorkFlowData_Instance
}


func NewWorkFlowData() *WorkFlowData {
    this := &WorkFlowData{}
    this.initFields()
    return this
}

