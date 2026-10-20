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

type GenericProfileActionJPanel struct {
    JPanel
    LogUtil *LogUtil
    GenericProfileAction *GenericProfileAction
    GenericProfileActionScriptJPanel *GenericProfileActionScriptJPanel
    BlankProfileActionScriptJPanel *JPanel
    GenericProfileActionScriptJScrollPane *JScrollPane
    IfJLabel *JLabel
    NewAlwaysJButton *JButton
    NewColorAtJButton *JButton
    NewTimeIntervalJButton *JButton
}


func (this *GenericProfileActionJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    
}


func NewGenericProfileActionJPanel() *GenericProfileActionJPanel {
    this := &GenericProfileActionJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
    return this
}


func (this *GenericProfileActionJPanel) UpdateProfileActionUI() {
        this.BlankProfileActionScriptJPanel.RemoveAll()
        var genericProfileActionScriptJPanel *GenericProfileActionScriptJPanel = NewGenericProfileActionScriptJPanel(); _ = genericProfileActionScriptJPanel
        JTreeInterfaceFactory_GetInstance().Set(genericProfileActionScriptJPanel)
        this.SetGenericProfileActionScriptJPanel(genericProfileActionScriptJPanel)
        this.GetGenericProfileActionScriptJPanel().SetGenericProfileActionScript(this.GetGenericProfileAction().GetGenericProfileActionScript())
        this.GetGenericProfileActionScriptJPanel().UpdateJTree()
        var layout *GroupLayout = any(this.BlankProfileActionScriptJPanel.GetLayout()).(*GroupLayout); _ = layout
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 395, Short_MAX_VALUE).AddComponent(this.GetGenericProfileActionScriptJPanel(), javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 256, Short_MAX_VALUE).AddComponent(this.GetGenericProfileActionScriptJPanel(), javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))
        this.LogUtil.PutF("Updated UI", this, "updateUI")
}


func (this *GenericProfileActionJPanel) GetGenericProfileAction() *GenericProfileAction {
        return this.GenericProfileAction
}


func (this *GenericProfileActionJPanel) SetGenericProfileAction(genericProfileAction *GenericProfileAction) {
        this.GenericProfileAction = genericProfileAction
}


func (this *GenericProfileActionJPanel) GetGenericProfileActionScriptJPanel() *GenericProfileActionScriptJPanel {
        return this.GenericProfileActionScriptJPanel
}


func (this *GenericProfileActionJPanel) SetGenericProfileActionScriptJPanel(genericProfileActionScriptJPanel *GenericProfileActionScriptJPanel) {
        this.GenericProfileActionScriptJPanel = genericProfileActionScriptJPanel
}


func (this *GenericProfileActionJPanel) InitComponents() {
        this.GenericProfileActionScriptJScrollPane = NewJScrollPane()
        this.BlankProfileActionScriptJPanel = NewJPanel()
        this.IfJLabel = NewJLabel()
        this.NewColorAtJButton = NewJButton()
        this.NewTimeIntervalJButton = NewJButton()
        this.NewAlwaysJButton = NewJButton()
        this.SetMinimumSize(NewDimension(440, 275))
        this.SetPreferredSize(NewDimension(440, 275))
        this.GenericProfileActionScriptJScrollPane.SetPreferredSize(NewDimension(474, 270))
        this.BlankProfileActionScriptJPanel.SetMinimumSize(NewDimension(400, 200))
        var blankProfileActionScriptJPanelLayout *GroupLayout = NewGroupLayout(this.BlankProfileActionScriptJPanel); _ = blankProfileActionScriptJPanelLayout
        this.BlankProfileActionScriptJPanel.SetLayout(blankProfileActionScriptJPanelLayout)
        blankProfileActionScriptJPanelLayout.SetHorizontalGroup(blankProfileActionScriptJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 9996, Short_MAX_VALUE))
        blankProfileActionScriptJPanelLayout.SetVerticalGroup(blankProfileActionScriptJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 9967, Short_MAX_VALUE))
        this.GenericProfileActionScriptJScrollPane.SetViewportView(this.BlankProfileActionScriptJPanel)
        this.IfJLabel.SetText("Conditions:")
        this.NewColorAtJButton.SetText("Color Range At")
        this.NewColorAtJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericProfileActionJPanel) ActionPerformed(evt *ActionEvent) {
                this.NewColorAtJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.NewTimeIntervalJButton.SetText("Time Interval")
        this.NewTimeIntervalJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericProfileActionJPanel) ActionPerformed(evt *ActionEvent) {
                this.NewTimeIntervalJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.NewAlwaysJButton.SetText("Always On/Off")
        this.NewAlwaysJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericProfileActionJPanel) ActionPerformed(evt *ActionEvent) {
                this.NewAlwaysJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.IfJLabel).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.NewColorAtJButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.NewTimeIntervalJButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.NewAlwaysJButton).AddContainerGap(61, Short_MAX_VALUE)).AddComponent(this.GenericProfileActionScriptJScrollPane, javax_Swing.GroupLayout.DEFAULT_SIZE, 440, Short_MAX_VALUE))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.IfJLabel, javax_Swing.GroupLayout.PREFERRED_SIZE, 14, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.NewColorAtJButton).AddComponent(this.NewTimeIntervalJButton).AddComponent(this.NewAlwaysJButton)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.GenericProfileActionScriptJScrollPane, javax_Swing.GroupLayout.DEFAULT_SIZE, 246, Short_MAX_VALUE)))
}


func (this *GenericProfileActionJPanel) NewTimeIntervalJButtonActionPerformed(evt *ActionEvent) {
        this.LogUtil.PutF("New Time Interval Action Condition", this, "newTimeIntervalJButtonActionPerformed")
        this.GetGenericProfileAction().GetGenericProfileActionScript().AddCondition(any(NewTimeIntervalActionScriptCondition()).(ProfileActionScriptConditionInterface))
        this.UpdateProfileActionUI()
}


func (this *GenericProfileActionJPanel) NewAlwaysJButtonActionPerformed(evt *ActionEvent) {
        this.LogUtil.PutF("New EveryTime Action Condition", this, "newEveryTimeJButtonActionPerformed")
        this.GetGenericProfileAction().GetGenericProfileActionScript().AddCondition(any(NewAlwaysActionScriptCondition()).(ProfileActionScriptConditionInterface))
        this.UpdateProfileActionUI()
}


func (this *GenericProfileActionJPanel) NewColorAtJButtonActionPerformed(evt *ActionEvent) {
        this.LogUtil.PutF("New Color At Action Condition", this, "newColorAtJButtonActionPerformed")
        this.GetGenericProfileAction().GetGenericProfileActionScript().AddCondition(any(NewColorAtActionScriptCondition()).(ProfileActionScriptConditionInterface))
        this.UpdateProfileActionUI()
}

