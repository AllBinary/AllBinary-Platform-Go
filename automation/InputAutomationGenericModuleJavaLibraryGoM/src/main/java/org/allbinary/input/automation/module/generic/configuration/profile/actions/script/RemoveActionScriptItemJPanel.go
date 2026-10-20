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

type RemoveActionScriptItemJPanel struct {
    Panel
    LogUtil *LogUtil
    GenericProfileActionJPanel *GenericProfileActionJPanel
    ProfileActionScriptConditionInterface ProfileActionScriptConditionInterface
    RemoveJButton *JButton
}


func (this *RemoveActionScriptItemJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewRemoveActionScriptItemJPanel(genericProfileActionJPanel *GenericProfileActionJPanel, profileActionScriptConditionInterface ProfileActionScriptConditionInterface) *RemoveActionScriptItemJPanel {
    this := &RemoveActionScriptItemJPanel{}
    this.Panel = *NewPanel()
    this.initFields()
        this.InitComponents()
        this.GenericProfileActionJPanel = genericProfileActionJPanel
        this.ProfileActionScriptConditionInterface = profileActionScriptConditionInterface
    return this
}


func (this *RemoveActionScriptItemJPanel) InitComponents() {
        this.RemoveJButton = NewJButton()
        this.SetLayout(NewGridLayout())
        this.RemoveJButton.SetText("Remove")
        this.RemoveJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *RemoveActionScriptItemJPanel) ActionPerformed(evt *ActionEvent) {
                this.RemoveJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.Add(this.RemoveJButton)
}


func (this *RemoveActionScriptItemJPanel) RemoveJButtonActionPerformed(evt *ActionEvent) {
        var genericProfileAction *GenericProfileAction = this.GenericProfileActionJPanel.GetGenericProfileAction(); _ = genericProfileAction
        if genericProfileAction == nil {
            this.LogUtil.PutF("GenericProfileAction was null", this, "removeJButtonActionPerformed")
            return
        }
        genericProfileAction.GetGenericProfileActionScript().RemoveCondition(this.ProfileActionScriptConditionInterface)
        this.GenericProfileActionJPanel.UpdateProfileActionUI()
}

