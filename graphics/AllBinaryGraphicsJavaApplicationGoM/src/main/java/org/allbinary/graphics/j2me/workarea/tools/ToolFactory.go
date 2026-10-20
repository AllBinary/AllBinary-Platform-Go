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

type ToolFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    SelectedTool GraphicsItemInterfaceFactoryInterface
    Tool string
    SELECTED_LABEL string
}

var ToolFactory_Instance *ToolFactory = NewToolFactory()

func (this *ToolFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.SelectedTool = nil
    this.Tool = nil
    this.SELECTED_LABEL = "Selected: "
}

func NewToolFactory() *ToolFactory {
    this := &ToolFactory{}
    this.initFields()
    return this
}


func ToolFactory_GetInstance() *ToolFactory {
        return ToolFactory_Instance
}


func (this *ToolFactory) SetSelectedTool(tool string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setSelectionTool", e)
                }
            }()
            StatusFactory_GetInstance().SetStatus(this.SELECTED_LABEL + tool)
            this.Tool = tool
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ToolFactory) GetSelectedToolFactory() GraphicsItemInterfaceFactoryInterface {
        StatusFactory_GetInstance().SetStatus("Tool In Use: " + this.Tool)
        if this.Tool != "" {
            this.SelectedTool = GraphicItemFactory_GetInstance().GetInstance(this.Tool).(GraphicsItemInterfaceFactoryInterface)
            this.LogUtil.PutF("New Tool: " + fmt.Sprint(this.SelectedTool), this, "getSelectedToolFactory")
            return this.SelectedTool
        } else {
            return nil
        }
}

