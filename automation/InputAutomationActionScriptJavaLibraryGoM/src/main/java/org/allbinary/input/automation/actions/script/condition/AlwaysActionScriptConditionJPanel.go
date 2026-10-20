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

type AlwaysActionScriptConditionJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AlwaysActionScriptConditionInterface AlwaysActionScriptConditionInterface
    AlwaysActionJButton *JButton
    AlwaysActionJDialog *JDialog
    AlwaysJComboBox *JComboBox
    JLabel1 *JLabel
    JLabel2 *JLabel
    OkJButton *JButton
    TitleJLabel *JLabel
}


func (this *AlwaysActionScriptConditionJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
}


func NewAlwaysActionScriptConditionJPanel(alwaysActionScriptConditionInterface AlwaysActionScriptConditionInterface) *AlwaysActionScriptConditionJPanel {
    this := &AlwaysActionScriptConditionJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.InitComponents()
        this.AlwaysActionScriptConditionInterface = alwaysActionScriptConditionInterface
        this.Set()
    return this
}


func (this *AlwaysActionScriptConditionJPanel) Set() {
        if this.AlwaysActionScriptConditionInterface.IsIsOn() {
            this.AlwaysJComboBox.SetSelectedIndex(0)
        } else {
            this.AlwaysJComboBox.SetSelectedIndex(1)
        }
        this.AlwaysActionScriptConditionInterface.Log()
}


func (this *AlwaysActionScriptConditionJPanel) Update() {
        if this.AlwaysJComboBox.GetSelectedIndex() == 0 {
            this.AlwaysActionScriptConditionInterface.SetIsOn(true)
        } else {
            this.AlwaysActionScriptConditionInterface.SetIsOn(false)
        }
        this.AlwaysActionScriptConditionInterface.Log()
}


func (this *AlwaysActionScriptConditionJPanel) GetAlwaysActionJDialog() *JDialog {
        return this.AlwaysActionJDialog
}


func (this *AlwaysActionScriptConditionJPanel) SetAlwaysActionJDialog(alwaysActionJDialog *JDialog) {
        this.AlwaysActionJDialog = alwaysActionJDialog
}


func (this *AlwaysActionScriptConditionJPanel) InitComponents() {
        this.AlwaysActionJDialog = NewJDialog()
        this.TitleJLabel = NewJLabel()
        this.OkJButton = NewJButton()
        this.JLabel2 = NewJLabel()
        this.AlwaysJComboBox = NewJComboBox()
        this.AlwaysActionJButton = NewJButton()
        this.JLabel1 = NewJLabel()
        this.AlwaysActionJDialog.SetMinimumSize(NewDimension(250, 150))
        this.TitleJLabel.SetText("Always Condition Options")
        this.OkJButton.SetText("OK")
        this.OkJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *AlwaysActionScriptConditionJPanel) ActionPerformed(evt *ActionEvent) {
                this.OkJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel2.SetText("On:")
        this.AlwaysJComboBox.SetModel(NewDefaultComboBoxModel([]string{"true", "false"}))
        var alwaysActionJDialogLayout *GroupLayout = NewGroupLayout(this.AlwaysActionJDialog.GetContentPane()); _ = alwaysActionJDialogLayout
        this.AlwaysActionJDialog.GetContentPane().SetLayout(alwaysActionJDialogLayout)
        alwaysActionJDialogLayout.SetHorizontalGroup(alwaysActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(alwaysActionJDialogLayout.CreateSequentialGroup().AddGroup(alwaysActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(alwaysActionJDialogLayout.CreateSequentialGroup().AddGap(47, 47, 47).AddGroup(alwaysActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING, false).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, alwaysActionJDialogLayout.CreateSequentialGroup().AddComponent(this.JLabel2).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddComponent(this.AlwaysJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddComponent(this.TitleJLabel))).AddGroup(alwaysActionJDialogLayout.CreateSequentialGroup().AddGap(79, 79, 79).AddComponent(this.OkJButton))).AddContainerGap(45, Short_MAX_VALUE)))
        alwaysActionJDialogLayout.SetVerticalGroup(alwaysActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(alwaysActionJDialogLayout.CreateSequentialGroup().AddContainerGap().AddComponent(this.TitleJLabel).AddGap(20, 20, 20).AddGroup(alwaysActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.AlwaysJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel2)).AddGap(15, 15, 15).AddComponent(this.OkJButton).AddContainerGap()))
        this.SetMinimumSize(NewDimension(100, 0))
        this.AlwaysActionJButton.SetText("Edit")
        this.AlwaysActionJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *AlwaysActionScriptConditionJPanel) ActionPerformed(evt *ActionEvent) {
                this.AlwaysActionJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel1.SetText("Always:")
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JLabel1).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 33, Short_MAX_VALUE).AddComponent(this.AlwaysActionJButton)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.AlwaysActionJButton)))
}


func (this *AlwaysActionScriptConditionJPanel) OkJButtonActionPerformed(evt *ActionEvent) {
        this.Update()
        this.GetAlwaysActionJDialog().SetVisible(false)
}


func (this *AlwaysActionScriptConditionJPanel) AlwaysActionJButtonActionPerformed(evt *ActionEvent) {
        this.GetAlwaysActionJDialog().SetVisible(true)
}

