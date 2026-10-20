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

type ActionScriptItemJPanel struct {
    JPanel
    ChildActionScriptConditionJPanel *JPanel
    NewActionScriptConditionJPanel *JPanel
    NewActionScriptInputJPanel *JPanel
    RemoveJPanel *JPanel
    ScriptItemJPanel *JPanel
}


func (this *ActionScriptItemJPanel) initFields() {
    
    
    
    
    
}


func NewActionScriptItemJPanel() *ActionScriptItemJPanel {
    this := &ActionScriptItemJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
    return this
}


func (this *ActionScriptItemJPanel) GetNewActionScriptInputJPanel() *JPanel {
        return this.NewActionScriptInputJPanel
}


func (this *ActionScriptItemJPanel) SetNewActionScriptInputJPanel(newActionScriptInputJPanel *JPanel) {
        this.NewActionScriptInputJPanel = newActionScriptInputJPanel
}


func (this *ActionScriptItemJPanel) GetRemoveJPanel() *JPanel {
        return this.RemoveJPanel
}


func (this *ActionScriptItemJPanel) SetRemoveJPanel(removeJPanel *JPanel) {
        this.RemoveJPanel = removeJPanel
}


func (this *ActionScriptItemJPanel) GetScriptItemJPanel() *JPanel {
        return this.ScriptItemJPanel
}


func (this *ActionScriptItemJPanel) SetScriptItemJPanel(scriptItemJPanel *JPanel) {
        this.ScriptItemJPanel = scriptItemJPanel
}


func (this *ActionScriptItemJPanel) GetNewActionScriptConditionJPanel() *JPanel {
        return this.NewActionScriptConditionJPanel
}


func (this *ActionScriptItemJPanel) SetNewActionScriptConditionJPanel(newActionScriptConditionJPanel *JPanel) {
        this.NewActionScriptConditionJPanel = newActionScriptConditionJPanel
}


func (this *ActionScriptItemJPanel) GetChildActionScriptConditionJPanel() *JPanel {
        return this.ChildActionScriptConditionJPanel
}


func (this *ActionScriptItemJPanel) SetChildActionScriptConditionJPanel(childActionScriptConditionJPanel *JPanel) {
        this.ChildActionScriptConditionJPanel = childActionScriptConditionJPanel
}


func (this *ActionScriptItemJPanel) InitComponents() {
        this.NewActionScriptInputJPanel = NewJPanel()
        this.ChildActionScriptConditionJPanel = NewJPanel()
        this.NewActionScriptConditionJPanel = NewJPanel()
        this.ScriptItemJPanel = NewJPanel()
        this.RemoveJPanel = NewJPanel()
        this.NewActionScriptInputJPanel.SetLayout(NewGridLayout(1, 0))
        var childActionScriptConditionJPanelLayout *GroupLayout = NewGroupLayout(this.ChildActionScriptConditionJPanel); _ = childActionScriptConditionJPanelLayout
        this.ChildActionScriptConditionJPanel.SetLayout(childActionScriptConditionJPanelLayout)
        childActionScriptConditionJPanelLayout.SetHorizontalGroup(childActionScriptConditionJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 229, Short_MAX_VALUE))
        childActionScriptConditionJPanelLayout.SetVerticalGroup(childActionScriptConditionJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 8, Short_MAX_VALUE))
        this.NewActionScriptConditionJPanel.SetLayout(NewGridLayout(1, 0))
        this.ScriptItemJPanel.SetLayout(NewGridLayout(1, 1))
        this.RemoveJPanel.SetLayout(NewGridLayout(1, 1))
        this.RemoveJPanel.SetMinimumSize(NewDimension(30, 20))
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.ScriptItemJPanel, javax_Swing.GroupLayout.PREFERRED_SIZE, 167, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.RemoveJPanel, javax_Swing.GroupLayout.DEFAULT_SIZE, 56, Short_MAX_VALUE)).AddComponent(this.NewActionScriptInputJPanel, javax_Swing.GroupLayout.DEFAULT_SIZE, 229, Short_MAX_VALUE).AddComponent(this.NewActionScriptConditionJPanel, javax_Swing.GroupLayout.DEFAULT_SIZE, 229, Short_MAX_VALUE).AddComponent(this.ChildActionScriptConditionJPanel, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.ScriptItemJPanel, javax_Swing.GroupLayout.DEFAULT_SIZE, 31, Short_MAX_VALUE).AddComponent(this.RemoveJPanel, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.DEFAULT_SIZE, 31, Short_MAX_VALUE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.ChildActionScriptConditionJPanel, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.NewActionScriptConditionJPanel, javax_Swing.GroupLayout.PREFERRED_SIZE, 28, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.NewActionScriptInputJPanel, javax_Swing.GroupLayout.PREFERRED_SIZE, 30, javax_Swing.GroupLayout.PREFERRED_SIZE)))
}

