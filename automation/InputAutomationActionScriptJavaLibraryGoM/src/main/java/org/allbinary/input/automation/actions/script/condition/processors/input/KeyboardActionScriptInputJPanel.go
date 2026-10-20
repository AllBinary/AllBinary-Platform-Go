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

type KeyboardActionScriptInputJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    KeyActionScriptInputInterface KeyboardActionScriptInputInterface
    AddKeyJLabel *JLabel
    DelayBetweenJLabel *JLabel
    DelayJTextField *JTextField
    HoldTimeJLabel *JLabel
    InputAutomationTypeJComboBox *JComboBox
    InputTypeJLabel *JLabel
    JLabel1 *JLabel
    JPanel1 *JPanel
    KeyActionJButton *JButton
    KeyActionJDialog *JDialog
    KeyIntegerJComboBox *JComboBox
    KeyNameJComboBox *JComboBox
    NormalJRadioButton *JRadioButton
    OkJButton *JButton
    PressJRadioButton *JRadioButton
    ReleaseJRadioButton *JRadioButton
    SequenceJRadioButton *JRadioButton
    SimultaneousJRadioButton *JRadioButton
    TestJLabel *JLabel
    TextJTextField *JTextField
    TimeJTextField *JTextField
    TitleJLabel *JLabel
}


func (this *KeyboardActionScriptInputJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewKeyboardActionScriptInputJPanel(keyActionScriptInputInterface KeyboardActionScriptInputInterface) *KeyboardActionScriptInputJPanel {
    this := &KeyboardActionScriptInputJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.InitComponents()
        this.KeyActionScriptInputInterface = keyActionScriptInputInterface
        var defaultComboBoxModel *DefaultComboBoxModel = NewDefaultComboBoxModel(); _ = defaultComboBoxModel
        defaultComboBoxModel.AddElement(TempInputRobotNames_DX_NAME)
        defaultComboBoxModel.AddElement(InputRobot_NAME)
        this.InputAutomationTypeJComboBox.SetModel(defaultComboBoxModel)
        this.KeyNameJComboBox.SetModel(KeyNameDefaultComboBoxModelFactory_GetInstance())
        this.KeyIntegerJComboBox.SetModel(KeyCharDefaultComboBoxModelFactory_GetInstance())
        if this.KeyActionScriptInputInterface.IsNormal() {
            this.NormalJRadioButton.SetSelected(true)
            this.PressJRadioButton.SetSelected(false)
            this.ReleaseJRadioButton.SetSelected(false)
        } else if this.KeyActionScriptInputInterface.IsPress() {
            this.NormalJRadioButton.SetSelected(false)
            this.PressJRadioButton.SetSelected(true)
            this.ReleaseJRadioButton.SetSelected(false)
        } else if this.KeyActionScriptInputInterface.IsRelease() {
            this.NormalJRadioButton.SetSelected(false)
            this.PressJRadioButton.SetSelected(false)
            this.ReleaseJRadioButton.SetSelected(true)
        }
        this.Set()
    return this
}


func (this *KeyboardActionScriptInputJPanel) Set() {
        var inputTypeComboBoxModel *ComboBoxModel = this.InputAutomationTypeJComboBox.GetModel(); _ = inputTypeComboBoxModel
        inputTypeComboBoxModel.SetSelectedItem(this.KeyActionScriptInputInterface.GetInputRobotInterface().GetName())
        this.TextJTextField.SetText(this.KeyActionScriptInputInterface.GetText())
        this.DelayJTextField.SetText(fmt.Sprint(this.KeyActionScriptInputInterface.GetDelayBetweenKeys()))
        this.TimeJTextField.SetText(fmt.Sprint(this.KeyActionScriptInputInterface.GetTime()))
        this.KeyActionScriptInputInterface.Log()
}


// throws Exception
func (this *KeyboardActionScriptInputJPanel) UpdateInputType() {
        var comboBoxModel *ComboBoxModel = this.InputAutomationTypeJComboBox.GetModel(); _ = comboBoxModel
        var selectedItem string = any(comboBoxModel.GetSelectedItem()).(string); _ = selectedItem
        if !StringValidationUtil_GetInstance().IsEmpty(selectedItem) {
            this.KeyActionScriptInputInterface.SetInputRobotInterface(InputRobotFactory_GetInstance().Get(selectedItem).(InputRobotInterface))
        }
}


func (this *KeyboardActionScriptInputJPanel) GetSelectedKey() int {
        var comboBoxModel *ComboBoxModel = this.KeyNameJComboBox.GetModel(); _ = comboBoxModel
        var selectedItem string = any(comboBoxModel.GetSelectedItem()).(string); _ = selectedItem
        if !StringValidationUtil_GetInstance().IsEmpty(selectedItem) {
            return any(KeySingletonFactory_GetHashtable().Get(selectedItem)).(int)
        }
        return nil
}


func (this *KeyboardActionScriptInputJPanel) UpdateKeys() {
        var text string = this.TextJTextField.GetText(); _ = text
        this.KeyActionScriptInputInterface.SetText(text)
}


// throws Exception
func (this *KeyboardActionScriptInputJPanel) Update() {
        this.UpdateInputType()
        this.UpdateKeys()
        this.KeyActionScriptInputInterface.SetTime(int(fmt.Sprint(this.TimeJTextField.GetText())))
        this.KeyActionScriptInputInterface.Log()
}


func (this *KeyboardActionScriptInputJPanel) GetKeyActionJDialog() *JDialog {
        return this.KeyActionJDialog
}


func (this *KeyboardActionScriptInputJPanel) SetKeyActionJDialog(keyActionJDialog *JDialog) {
        this.KeyActionJDialog = keyActionJDialog
}


func (this *KeyboardActionScriptInputJPanel) InitComponents() {
        this.KeyActionJDialog = NewJDialog()
        this.TitleJLabel = NewJLabel()
        this.AddKeyJLabel = NewJLabel()
        this.KeyNameJComboBox = NewJComboBox()
        this.TimeJTextField = NewJTextField()
        this.HoldTimeJLabel = NewJLabel()
        this.InputAutomationTypeJComboBox = NewJComboBox()
        this.InputTypeJLabel = NewJLabel()
        this.DelayJTextField = NewJTextField()
        this.DelayBetweenJLabel = NewJLabel()
        this.TestJLabel = NewJLabel()
        this.TextJTextField = NewJTextField()
        this.PressJRadioButton = NewJRadioButton()
        this.ReleaseJRadioButton = NewJRadioButton()
        this.NormalJRadioButton = NewJRadioButton()
        this.SimultaneousJRadioButton = NewJRadioButton()
        this.SequenceJRadioButton = NewJRadioButton()
        this.JPanel1 = NewJPanel()
        this.OkJButton = NewJButton()
        this.KeyIntegerJComboBox = NewJComboBox()
        this.KeyActionJButton = NewJButton()
        this.JLabel1 = NewJLabel()
        this.KeyActionJDialog.SetMinimumSize(NewDimension(255, 280))
        this.TitleJLabel.SetHorizontalAlignment(javax_Swing.SwingConstants.CENTER)
        this.TitleJLabel.SetText("Keyboard Input Options")
        this.AddKeyJLabel.SetText("Add Key:")
        this.KeyNameJComboBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *KeyboardActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.KeyNameJComboBoxActionPerformed(evt)
        }
*/ NewActionListener())
        this.TimeJTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.HoldTimeJLabel.SetText("Hold Time (ms):")
        this.InputTypeJLabel.SetText("Input Type:")
        this.DelayJTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.DelayBetweenJLabel.SetText("Delay Between Keys (ms):")
        this.TestJLabel.SetText("Text:")
        this.TextJTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.PressJRadioButton.SetText("Press")
        this.PressJRadioButton.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.PressJRadioButton.SetMargin(NewInsets(0, 0, 0, 0))
        this.PressJRadioButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *KeyboardActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.PressJRadioButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.ReleaseJRadioButton.SetText("Release")
        this.ReleaseJRadioButton.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.ReleaseJRadioButton.SetMargin(NewInsets(0, 0, 0, 0))
        this.ReleaseJRadioButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *KeyboardActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.ReleaseJRadioButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.NormalJRadioButton.SetSelected(true)
        this.NormalJRadioButton.SetText("Normal")
        this.NormalJRadioButton.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.NormalJRadioButton.SetMargin(NewInsets(0, 0, 0, 0))
        this.NormalJRadioButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *KeyboardActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.NormalJRadioButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.SimultaneousJRadioButton.SetSelected(true)
        this.SimultaneousJRadioButton.SetText("Simultaneous")
        this.SimultaneousJRadioButton.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.SimultaneousJRadioButton.SetMargin(NewInsets(0, 0, 0, 0))
        this.SimultaneousJRadioButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *KeyboardActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.SimultaneousJRadioButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.SequenceJRadioButton.SetText("Sequence")
        this.SequenceJRadioButton.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.SequenceJRadioButton.SetMargin(NewInsets(0, 0, 0, 0))
        this.SequenceJRadioButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *KeyboardActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.SequenceJRadioButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.OkJButton.SetText("OK")
        this.OkJButton.SetHorizontalTextPosition(javax_Swing.SwingConstants.CENTER)
        this.OkJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *KeyboardActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.OkJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JPanel1.Add(this.OkJButton)
        this.KeyIntegerJComboBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *KeyboardActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.KeyIntegerJComboBoxActionPerformed(evt)
        }
*/ NewActionListener())
        var keyActionJDialogLayout *GroupLayout = NewGroupLayout(this.KeyActionJDialog.GetContentPane()); _ = keyActionJDialogLayout
        this.KeyActionJDialog.GetContentPane().SetLayout(keyActionJDialogLayout)
        keyActionJDialogLayout.SetHorizontalGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(keyActionJDialogLayout.CreateSequentialGroup().AddContainerGap().AddGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(keyActionJDialogLayout.CreateSequentialGroup().AddGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING, false).AddComponent(this.AddKeyJLabel, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddComponent(this.InputTypeJLabel, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.TRAILING).AddComponent(this.InputAutomationTypeJComboBox, 0, 148, Short_MAX_VALUE).AddComponent(this.TextJTextField, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 148, Short_MAX_VALUE).AddGroup(keyActionJDialogLayout.CreateSequentialGroup().AddComponent(this.KeyNameJComboBox, 0, 101, Short_MAX_VALUE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.KeyIntegerJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, 41, javax_Swing.GroupLayout.PREFERRED_SIZE)))).AddComponent(this.TestJLabel, javax_Swing.GroupLayout.PREFERRED_SIZE, 60, javax_Swing.GroupLayout.PREFERRED_SIZE).AddGroup(keyActionJDialogLayout.CreateSequentialGroup().AddComponent(this.NormalJRadioButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.PressJRadioButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.ReleaseJRadioButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 48, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, keyActionJDialogLayout.CreateSequentialGroup().AddComponent(this.HoldTimeJLabel).AddGap(22, 22, 22).AddComponent(this.TimeJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 113, Short_MAX_VALUE)).AddGroup(keyActionJDialogLayout.CreateSequentialGroup().AddComponent(this.DelayBetweenJLabel).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 15, Short_MAX_VALUE).AddComponent(this.DelayJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 68, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddGroup(keyActionJDialogLayout.CreateSequentialGroup().AddComponent(this.SimultaneousJRadioButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.SequenceJRadioButton))).AddContainerGap()).AddComponent(this.TitleJLabel, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.DEFAULT_SIZE, 229, Short_MAX_VALUE).AddComponent(this.JPanel1, javax_Swing.GroupLayout.DEFAULT_SIZE, 229, Short_MAX_VALUE))
        keyActionJDialogLayout.SetVerticalGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(keyActionJDialogLayout.CreateSequentialGroup().AddComponent(this.TitleJLabel).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.InputTypeJLabel).AddComponent(this.InputAutomationTypeJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.AddKeyJLabel).AddComponent(this.KeyIntegerJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.KeyNameJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.TestJLabel).AddComponent(this.TextJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.PressJRadioButton).AddComponent(this.NormalJRadioButton).AddComponent(this.ReleaseJRadioButton)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.HoldTimeJLabel).AddComponent(this.TimeJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.DelayBetweenJLabel).AddComponent(this.DelayJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(keyActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.SimultaneousJRadioButton).AddComponent(this.SequenceJRadioButton)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JPanel1, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddContainerGap()))
        this.SetMinimumSize(NewDimension(100, 0))
        this.KeyActionJButton.SetText("Edit")
        this.KeyActionJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *KeyboardActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.KeyActionJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel1.SetText("Key:")
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JLabel1).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 49, Short_MAX_VALUE).AddComponent(this.KeyActionJButton)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.KeyActionJButton)))
}


func (this *KeyboardActionScriptInputJPanel) PrependText(integer int) {
        if integer != 0 {
            var newKey string = "&#" + fmt.Sprint(integer) + ";"; _ = newKey
            this.TextJTextField.SetText(newKey + fmt.Sprint(this.TextJTextField.GetText()))
        }
}


func (this *KeyboardActionScriptInputJPanel) KeyIntegerJComboBoxActionPerformed(evt *ActionEvent) {
        this.KeyIntegerJComboBox.GetSelectedIndex()
        this.KeyNameJComboBox.SetSelectedIndex(this.KeyIntegerJComboBox.GetSelectedIndex())
}


func (this *KeyboardActionScriptInputJPanel) SequenceJRadioButtonActionPerformed(evt *ActionEvent) {
        if this.SequenceJRadioButton.IsSelected() {
            this.DelayJTextField.SetText(CommonPhoneStrings_GetInstance().ONE)
            this.SimultaneousJRadioButton.SetSelected(false)
        }
}


func (this *KeyboardActionScriptInputJPanel) SimultaneousJRadioButtonActionPerformed(evt *ActionEvent) {
        if this.SimultaneousJRadioButton.IsSelected() {
            this.DelayJTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
            this.SequenceJRadioButton.SetSelected(false)
        }
}


func (this *KeyboardActionScriptInputJPanel) KeyNameJComboBoxActionPerformed(evt *ActionEvent) {
        var integer int = this.GetSelectedKey(); _ = integer
        this.PrependText(integer)
        this.KeyIntegerJComboBox.SetSelectedIndex(this.KeyNameJComboBox.GetSelectedIndex())
}


func (this *KeyboardActionScriptInputJPanel) OkJButtonActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "okButtonActionPerformed", e)
                }
            }()
            this.Update()
            this.GetKeyActionJDialog().SetVisible(false)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *KeyboardActionScriptInputJPanel) KeyActionJButtonActionPerformed(evt *ActionEvent) {
        this.GetKeyActionJDialog().SetVisible(true)
}


func (this *KeyboardActionScriptInputJPanel) NormalJRadioButtonActionPerformed(evt *ActionEvent) {
        if this.NormalJRadioButton.IsSelected() {
            this.KeyActionScriptInputInterface.SetNormal()
            this.PressJRadioButton.SetSelected(false)
            this.ReleaseJRadioButton.SetSelected(false)
        }
}


func (this *KeyboardActionScriptInputJPanel) ReleaseJRadioButtonActionPerformed(evt *ActionEvent) {
        if this.ReleaseJRadioButton.IsSelected() {
            this.KeyActionScriptInputInterface.SetRelease(true)
            this.KeyActionScriptInputInterface.SetPress(false)
            this.PressJRadioButton.SetSelected(false)
            this.NormalJRadioButton.SetSelected(false)
        }
}


func (this *KeyboardActionScriptInputJPanel) PressJRadioButtonActionPerformed(evt *ActionEvent) {
        if this.PressJRadioButton.IsSelected() {
            this.KeyActionScriptInputInterface.SetPress(true)
            this.KeyActionScriptInputInterface.SetRelease(false)
            this.ReleaseJRadioButton.SetSelected(false)
            this.NormalJRadioButton.SetSelected(false)
        }
}

