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

type TimeIntervalActionScriptConditionJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TimeIntervalActionScriptConditionInterface TimeIntervalActionScriptConditionInterface
    JLabel1 *JLabel
    JLabel2 *JLabel
    OkJButton *JButton
    TimeIntervalActionJButton *JButton
    TimeIntervalActionJDialog *JDialog
    TimeIntervalJTextField *JTextField
    TitleJLabel *JLabel
}


func (this *TimeIntervalActionScriptConditionJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
}


func NewTimeIntervalActionScriptConditionJPanel(timeIntervalActionScriptConditionInterface TimeIntervalActionScriptConditionInterface) *TimeIntervalActionScriptConditionJPanel {
    this := &TimeIntervalActionScriptConditionJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.InitComponents()
        this.TimeIntervalActionScriptConditionInterface = timeIntervalActionScriptConditionInterface
        this.Set()
    return this
}


func (this *TimeIntervalActionScriptConditionJPanel) Set() {
        this.GetTimeIntervalJTextField().SetText(fmt.Sprint(this.TimeIntervalActionScriptConditionInterface.GetTimeDelayHelper().TimeDelayHelper_().Delay))
        this.TimeIntervalActionScriptConditionInterface.Log()
}


func (this *TimeIntervalActionScriptConditionJPanel) Update() {
        this.TimeIntervalActionScriptConditionInterface.GetTimeDelayHelper().TimeDelayHelper_().Delay = fmt.Sprint(this.GetTimeIntervalJTextField().GetText())
        this.TimeIntervalActionScriptConditionInterface.Log()
}


func (this *TimeIntervalActionScriptConditionJPanel) GetTimeIntervalJTextField() *JTextField {
        return this.TimeIntervalJTextField
}


func (this *TimeIntervalActionScriptConditionJPanel) SetTimeIntervalJTextField(timeIntervalJTextField *JTextField) {
        this.TimeIntervalJTextField = timeIntervalJTextField
}


func (this *TimeIntervalActionScriptConditionJPanel) GetTimeIntervalActionJDialog() *JDialog {
        return this.TimeIntervalActionJDialog
}


func (this *TimeIntervalActionScriptConditionJPanel) SetTimeIntervalActionJDialog(timeIntervalActionJDialog *JDialog) {
        this.TimeIntervalActionJDialog = timeIntervalActionJDialog
}


func (this *TimeIntervalActionScriptConditionJPanel) InitComponents() {
        this.TimeIntervalActionJDialog = NewJDialog()
        this.TitleJLabel = NewJLabel()
        this.OkJButton = NewJButton()
        this.JLabel2 = NewJLabel()
        this.TimeIntervalJTextField = NewJTextField()
        this.TimeIntervalActionJButton = NewJButton()
        this.JLabel1 = NewJLabel()
        this.TimeIntervalActionJDialog.SetMinimumSize(NewDimension(250, 150))
        this.TitleJLabel.SetText("Time Interval Condition Options")
        this.OkJButton.SetText("OK")
        this.OkJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *TimeIntervalActionScriptConditionJPanel) ActionPerformed(evt *ActionEvent) {
                this.OkJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel2.SetText("Time (ms):")
        var timeIntervalActionJDialogLayout *GroupLayout = NewGroupLayout(this.TimeIntervalActionJDialog.GetContentPane()); _ = timeIntervalActionJDialogLayout
        this.TimeIntervalActionJDialog.GetContentPane().SetLayout(timeIntervalActionJDialogLayout)
        timeIntervalActionJDialogLayout.SetHorizontalGroup(timeIntervalActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(timeIntervalActionJDialogLayout.CreateSequentialGroup().AddGroup(timeIntervalActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.TRAILING).AddGroup(javax_Swing.GroupLayout.Alignment.LEADING, timeIntervalActionJDialogLayout.CreateSequentialGroup().AddGap(82, 82, 82).AddComponent(this.OkJButton)).AddGroup(javax_Swing.GroupLayout.Alignment.LEADING, timeIntervalActionJDialogLayout.CreateSequentialGroup().AddGap(32, 32, 32).AddGroup(timeIntervalActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(timeIntervalActionJDialogLayout.CreateSequentialGroup().AddComponent(this.JLabel2).AddGap(16, 16, 16).AddComponent(this.TimeIntervalJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 85, Short_MAX_VALUE)).AddComponent(this.TitleJLabel)))).AddContainerGap(36, Short_MAX_VALUE)))
        timeIntervalActionJDialogLayout.SetVerticalGroup(timeIntervalActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(timeIntervalActionJDialogLayout.CreateSequentialGroup().AddContainerGap().AddComponent(this.TitleJLabel).AddGap(20, 20, 20).AddGroup(timeIntervalActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel2).AddComponent(this.TimeIntervalJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddGap(15, 15, 15).AddComponent(this.OkJButton).AddContainerGap()))
        this.SetMinimumSize(NewDimension(100, 0))
        this.TimeIntervalActionJButton.SetText("Edit")
        this.TimeIntervalActionJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *TimeIntervalActionScriptConditionJPanel) ActionPerformed(evt *ActionEvent) {
                this.TimeIntervalActionJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel1.SetText("Time Interval:")
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JLabel1).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddComponent(this.TimeIntervalActionJButton)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.TimeIntervalActionJButton)))
}


func (this *TimeIntervalActionScriptConditionJPanel) OkJButtonActionPerformed(evt *ActionEvent) {
        this.Update()
        this.GetTimeIntervalActionJDialog().SetVisible(false)
}


func (this *TimeIntervalActionScriptConditionJPanel) TimeIntervalActionJButtonActionPerformed(evt *ActionEvent) {
        this.GetTimeIntervalActionJDialog().SetVisible(true)
}

