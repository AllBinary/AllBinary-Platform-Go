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

type MouseActionScriptInputJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    MouseActionScriptInputInterface MouseActionScriptInputInterface
    Button1JCheckBox *JCheckBox
    Button2JCheckBox *JCheckBox
    InputAutomationTypeJComboBox *JComboBox
    InputTypeJLabel *JLabel
    JLabel1 *JLabel
    JLabel2 *JLabel
    JLabel3 *JLabel
    JLabel4 *JLabel
    JLabel5 *JLabel
    JLabel6 *JLabel
    MouseActionJButton *JButton
    MouseActionJDialog *JDialog
    MouseDelayJTextField *JTextField
    MouseXJTextField *JTextField
    MouseYJTextField *JTextField
    OkJButton *JButton
    TitleJLabel *JLabel
}


func (this *MouseActionScriptInputJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewMouseActionScriptInputJPanel(mouseActionScriptInputInterface MouseActionScriptInputInterface) *MouseActionScriptInputJPanel {
    this := &MouseActionScriptInputJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.LogUtil.PutF(CommonLabels_GetInstance().START + mouseActionScriptInputInterface, this, this.CommonStrings.CONSTRUCTOR)
        this.InitComponents()
        this.MouseActionScriptInputInterface = mouseActionScriptInputInterface
        var defaultComboBoxModel *DefaultComboBoxModel = NewDefaultComboBoxModel(); _ = defaultComboBoxModel
        defaultComboBoxModel.AddElement(TempInputRobotNames_LOW_NAME)
        defaultComboBoxModel.AddElement(TempInputRobotNames_SYS_NAME)
        defaultComboBoxModel.AddElement(InputRobot_NAME)
        this.InputAutomationTypeJComboBox.SetModel(defaultComboBoxModel)
        this.Set()
    return this
}


func (this *MouseActionScriptInputJPanel) Set() {
        var inputTypeComboBoxModel *ComboBoxModel = this.InputAutomationTypeJComboBox.GetModel(); _ = inputTypeComboBoxModel
        inputTypeComboBoxModel.SetSelectedItem(this.MouseActionScriptInputInterface.GetInputRobotInterface().GetName())
        if (this.MouseActionScriptInputInterface.GetButtonClicks() & InputEvent_BUTTON1_MASK) != 0 {
            this.Button1JCheckBox.SetSelected(true)
        } else {
            this.Button1JCheckBox.SetSelected(false)
        }
        if (this.MouseActionScriptInputInterface.GetButtonClicks() & InputEvent_BUTTON2_MASK) != 0 {
            this.Button2JCheckBox.SetSelected(true)
        } else {
            this.Button2JCheckBox.SetSelected(false)
        }
        this.MouseDelayJTextField.SetText(fmt.Sprint(this.MouseActionScriptInputInterface.GetTime()))
        this.MouseXJTextField.SetText(fmt.Sprint(this.MouseActionScriptInputInterface.GetPoint().X))
        this.MouseYJTextField.SetText(fmt.Sprint(this.MouseActionScriptInputInterface.GetPoint().Y))
        this.MouseActionScriptInputInterface.Log()
}


// throws Exception
func (this *MouseActionScriptInputJPanel) UpdateInputType() {
        var comboBoxModel *ComboBoxModel = this.InputAutomationTypeJComboBox.GetModel(); _ = comboBoxModel
        var selectedItem string = any(comboBoxModel.GetSelectedItem()).(string); _ = selectedItem
        if !StringValidationUtil_GetInstance().IsEmpty(selectedItem) {
            this.MouseActionScriptInputInterface.SetInputRobotInterface(InputRobotFactory_GetInstance().Get(selectedItem).(InputRobotInterface))
        }
}


// throws Exception
func (this *MouseActionScriptInputJPanel) Update() {
        this.UpdateInputType()
        this.MouseActionScriptInputInterface.SetTime(fmt.Sprint(this.MouseDelayJTextField.GetText()))
        this.LogUtil.PutF("Button State: " + fmt.Sprint(this.MouseActionScriptInputInterface.GetButtonClicks()), this, "update")
        if this.Button1JCheckBox.IsSelected() {
            if (this.MouseActionScriptInputInterface.GetButtonClicks() & InputEvent_BUTTON1_MASK) == 0 {
                this.MouseActionScriptInputInterface.SetButtonClicks(this.MouseActionScriptInputInterface.GetButtonClicks() | InputEvent_BUTTON1_MASK)
                this.LogUtil.PutF("Button 1 Selected: " + fmt.Sprint(this.MouseActionScriptInputInterface.GetButtonClicks()), this, "update")
            }
        } else {
            if (this.MouseActionScriptInputInterface.GetButtonClicks() & InputEvent_BUTTON1_MASK) != 0 {
                this.MouseActionScriptInputInterface.SetButtonClicks(this.MouseActionScriptInputInterface.GetButtonClicks() ^ InputEvent_BUTTON1_MASK)
                this.LogUtil.PutF("Button 1 Deselected: " + fmt.Sprint(this.MouseActionScriptInputInterface.GetButtonClicks()), this, "update")
            }
        }
        if this.Button2JCheckBox.IsSelected() {
            if (this.MouseActionScriptInputInterface.GetButtonClicks() & InputEvent_BUTTON2_MASK) == 0 {
                this.MouseActionScriptInputInterface.SetButtonClicks(this.MouseActionScriptInputInterface.GetButtonClicks() | InputEvent_BUTTON2_MASK)
                this.LogUtil.PutF("Button 2 Selected: " + fmt.Sprint(this.MouseActionScriptInputInterface.GetButtonClicks()), this, "update")
            }
        } else {
            if (this.MouseActionScriptInputInterface.GetButtonClicks() & InputEvent_BUTTON2_MASK) != 0 {
                this.MouseActionScriptInputInterface.SetButtonClicks(this.MouseActionScriptInputInterface.GetButtonClicks() ^ InputEvent_BUTTON2_MASK)
                this.LogUtil.PutF("Button 2 Deselected: " + fmt.Sprint(this.MouseActionScriptInputInterface.GetButtonClicks()), this, "update")
            }
        }
        var mouseXString string = this.MouseXJTextField.GetText(); _ = mouseXString
        if !StringValidationUtil_GetInstance().IsEmpty(mouseXString) && StringValidationUtil_GetInstance().IsNumber(mouseXString) {
            this.MouseActionScriptInputInterface.GetPoint().X = func() int { value, _ := strconv.Atoi(mouseXString); return value }()
        }
        var mouseYString string = this.MouseYJTextField.GetText(); _ = mouseYString
        if !StringValidationUtil_GetInstance().IsEmpty(mouseYString) && StringValidationUtil_GetInstance().IsNumber(mouseYString) {
            this.MouseActionScriptInputInterface.GetPoint().Y = func() int { value, _ := strconv.Atoi(mouseYString); return value }()
        }
        this.MouseActionScriptInputInterface.Log()
}


func (this *MouseActionScriptInputJPanel) GetMouseActionJDialog() *JDialog {
        return this.MouseActionJDialog
}


func (this *MouseActionScriptInputJPanel) SetMouseActionJDialog(mouseActionJDialog *JDialog) {
        this.MouseActionJDialog = mouseActionJDialog
}


func (this *MouseActionScriptInputJPanel) InitComponents() {
        this.MouseActionJDialog = NewJDialog()
        this.TitleJLabel = NewJLabel()
        this.OkJButton = NewJButton()
        this.JLabel2 = NewJLabel()
        this.JLabel3 = NewJLabel()
        this.Button1JCheckBox = NewJCheckBox()
        this.Button2JCheckBox = NewJCheckBox()
        this.JLabel4 = NewJLabel()
        this.MouseXJTextField = NewJTextField()
        this.JLabel5 = NewJLabel()
        this.MouseYJTextField = NewJTextField()
        this.InputTypeJLabel = NewJLabel()
        this.InputAutomationTypeJComboBox = NewJComboBox()
        this.JLabel6 = NewJLabel()
        this.MouseDelayJTextField = NewJTextField()
        this.MouseActionJButton = NewJButton()
        this.JLabel1 = NewJLabel()
        this.MouseActionJDialog.SetMinimumSize(NewDimension(275, 185))
        this.TitleJLabel.SetText("Mouse Input Options")
        this.OkJButton.SetText("OK")
        this.OkJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MouseActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.OkJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel2.SetText("Buttons:")
        this.JLabel3.SetText("Move To:")
        this.Button1JCheckBox.SetText(CommonPhoneStrings_GetInstance().ONE)
        this.Button1JCheckBox.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.Button1JCheckBox.SetMargin(NewInsets(0, 0, 0, 0))
        this.Button2JCheckBox.SetText(CommonPhoneStrings_GetInstance().TWO)
        this.Button2JCheckBox.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.Button2JCheckBox.SetMargin(NewInsets(0, 0, 0, 0))
        this.Button2JCheckBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MouseActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.Button2JCheckBoxActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel4.SetText("x:")
        this.MouseXJTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.JLabel5.SetText("y:")
        this.MouseYJTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.InputTypeJLabel.SetText("Input Type:")
        this.JLabel6.SetText("Delay:")
        this.MouseDelayJTextField.SetText("100")
        var mouseActionJDialogLayout *GroupLayout = NewGroupLayout(this.MouseActionJDialog.GetContentPane()); _ = mouseActionJDialogLayout
        this.MouseActionJDialog.GetContentPane().SetLayout(mouseActionJDialogLayout)
        mouseActionJDialogLayout.SetHorizontalGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(mouseActionJDialogLayout.CreateSequentialGroup().AddContainerGap().AddGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.InputTypeJLabel).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, mouseActionJDialogLayout.CreateSequentialGroup().AddGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, mouseActionJDialogLayout.CreateSequentialGroup().AddComponent(this.JLabel2).AddGap(25, 25, 25)).AddGroup(mouseActionJDialogLayout.CreateSequentialGroup().AddGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JLabel3).AddComponent(this.JLabel6)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED))).AddGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.MouseDelayJTextField, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.DEFAULT_SIZE, 154, Short_MAX_VALUE).AddGroup(mouseActionJDialogLayout.CreateSequentialGroup().AddComponent(this.TitleJLabel).AddGap(45, 45, 45)).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, mouseActionJDialogLayout.CreateSequentialGroup().AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.InputAutomationTypeJComboBox, 0, 154, Short_MAX_VALUE)).AddGroup(mouseActionJDialogLayout.CreateSequentialGroup().AddComponent(this.Button1JCheckBox).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.Button2JCheckBox)).AddGroup(mouseActionJDialogLayout.CreateSequentialGroup().AddGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.TRAILING).AddGroup(mouseActionJDialogLayout.CreateSequentialGroup().AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.OkJButton)).AddGroup(mouseActionJDialogLayout.CreateSequentialGroup().AddComponent(this.JLabel4).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 5, Short_MAX_VALUE).AddComponent(this.MouseXJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 59, javax_Swing.GroupLayout.PREFERRED_SIZE))).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel5).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.MouseYJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 62, Short_MAX_VALUE))))).AddContainerGap()))
        mouseActionJDialogLayout.SetVerticalGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(mouseActionJDialogLayout.CreateSequentialGroup().AddComponent(this.TitleJLabel).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.InputTypeJLabel).AddComponent(this.InputAutomationTypeJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JLabel2, javax_Swing.GroupLayout.PREFERRED_SIZE, 15, javax_Swing.GroupLayout.PREFERRED_SIZE).AddGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.Button1JCheckBox).AddComponent(this.Button2JCheckBox))).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel6).AddComponent(this.MouseDelayJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(mouseActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel3).AddComponent(this.JLabel4).AddComponent(this.JLabel5).AddComponent(this.MouseYJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.MouseXJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddComponent(this.OkJButton).AddContainerGap()))
        this.SetMinimumSize(NewDimension(100, 0))
        this.MouseActionJButton.SetText("Edit")
        this.MouseActionJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MouseActionScriptInputJPanel) ActionPerformed(evt *ActionEvent) {
                this.MouseActionJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel1.SetText("Mouse:")
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JLabel1).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 36, Short_MAX_VALUE).AddComponent(this.MouseActionJButton)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.MouseActionJButton)))
}


func (this *MouseActionScriptInputJPanel) Button2JCheckBoxActionPerformed(evt *ActionEvent) {
}


func (this *MouseActionScriptInputJPanel) OkJButtonActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "okButtonActionPerformed", e)
                }
            }()
            this.Update()
            this.GetMouseActionJDialog().SetVisible(false)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MouseActionScriptInputJPanel) MouseActionJButtonActionPerformed(evt *ActionEvent) {
        this.GetMouseActionJDialog().SetVisible(true)
}

