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

type GenericProfileActionsJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    GenericProfileActions *GenericProfileActions
    GenericProfileActionJPanel *GenericProfileActionJPanel
    BlankGenericProfileActionJPanel *JPanel
    DeleteActionJButton *JButton
    GenericProfileActionsJList *JList
    GenericProfileActionsJPanel *JPanel
    GenericProfileActionsJScrollPane *JScrollPane
    NewActionJButton *JButton
    NewProfileActionJDialog *JDialog
    NewProfileActionJLabel *JLabel
    NewProfileActionJTextField *JTextField
    OkNewProfileActionJButton *JButton
    SaveJButton *JButton
}


func (this *GenericProfileActionsJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewGenericProfileActionsJPanel() *GenericProfileActionsJPanel {
    this := &GenericProfileActionsJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
    return this
}


func (this *GenericProfileActionsJPanel) InitComponents() {
        this.NewProfileActionJDialog = NewJDialog()
        this.NewProfileActionJLabel = NewJLabel()
        this.OkNewProfileActionJButton = NewJButton()
        this.NewProfileActionJTextField = NewJTextField()
        this.GenericProfileActionsJPanel = NewJPanel()
        this.GenericProfileActionsJScrollPane = NewJScrollPane()
        this.GenericProfileActionsJList = NewJList()
        this.NewActionJButton = NewJButton()
        this.DeleteActionJButton = NewJButton()
        this.SaveJButton = NewJButton()
        this.BlankGenericProfileActionJPanel = NewJPanel()
        this.NewProfileActionJDialog.SetMinimumSize(NewDimension(250, 150))
        this.NewProfileActionJDialog.SetResizable(false)
        this.NewProfileActionJLabel.SetText("New Action Name")
        this.OkNewProfileActionJButton.SetText("Ok")
        this.OkNewProfileActionJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericProfileActionsJPanel) ActionPerformed(evt *ActionEvent) {
                this.OkNewProfileActionJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        var newProfileActionJDialogLayout *GroupLayout = NewGroupLayout(this.NewProfileActionJDialog.GetContentPane()); _ = newProfileActionJDialogLayout
        this.NewProfileActionJDialog.GetContentPane().SetLayout(newProfileActionJDialogLayout)
        newProfileActionJDialogLayout.SetHorizontalGroup(newProfileActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(newProfileActionJDialogLayout.CreateSequentialGroup().AddGroup(newProfileActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(newProfileActionJDialogLayout.CreateSequentialGroup().AddGap(93, 93, 93).AddComponent(this.OkNewProfileActionJButton)).AddGroup(newProfileActionJDialogLayout.CreateSequentialGroup().AddGap(75, 75, 75).AddComponent(this.NewProfileActionJLabel)).AddGroup(newProfileActionJDialogLayout.CreateSequentialGroup().AddGap(29, 29, 29).AddComponent(this.NewProfileActionJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 178, javax_Swing.GroupLayout.PREFERRED_SIZE))).AddContainerGap(34, Short_MAX_VALUE)))
        newProfileActionJDialogLayout.SetVerticalGroup(newProfileActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, newProfileActionJDialogLayout.CreateSequentialGroup().AddContainerGap().AddComponent(this.NewProfileActionJLabel).AddGap(21, 21, 21).AddComponent(this.NewProfileActionJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 20, Short_MAX_VALUE).AddComponent(this.OkNewProfileActionJButton).AddContainerGap()))
        this.GenericProfileActionsJPanel.SetPreferredSize(NewDimension(10000, 10000))
        this.GenericProfileActionsJList.AddListSelectionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericProfileActionsJPanel) ValueChanged(evt *ListSelectionEvent) {
                this.GenericProfileActionsJListValueChanged(evt)
        }
*/ NewListSelectionListener())
        this.GenericProfileActionsJScrollPane.SetViewportView(this.GenericProfileActionsJList)
        this.NewActionJButton.SetText("New")
        this.NewActionJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericProfileActionsJPanel) ActionPerformed(evt *ActionEvent) {
                this.NewActionJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.DeleteActionJButton.SetText("Delete")
        this.DeleteActionJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericProfileActionsJPanel) ActionPerformed(evt *ActionEvent) {
                this.DeleteActionJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.SaveJButton.SetText("Save")
        this.SaveJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericProfileActionsJPanel) ActionPerformed(evt *ActionEvent) {
                this.SaveJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.BlankGenericProfileActionJPanel.SetMinimumSize(NewDimension(500, 300))
        var blankGenericProfileActionJPanelLayout *GroupLayout = NewGroupLayout(this.BlankGenericProfileActionJPanel); _ = blankGenericProfileActionJPanelLayout
        this.BlankGenericProfileActionJPanel.SetLayout(blankGenericProfileActionJPanelLayout)
        blankGenericProfileActionJPanelLayout.SetHorizontalGroup(blankGenericProfileActionJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 500, Short_MAX_VALUE))
        blankGenericProfileActionJPanelLayout.SetVerticalGroup(blankGenericProfileActionJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 296, Short_MAX_VALUE))
        var genericProfileActionsJPanelLayout *GroupLayout = NewGroupLayout(this.GenericProfileActionsJPanel); _ = genericProfileActionsJPanelLayout
        this.GenericProfileActionsJPanel.SetLayout(genericProfileActionsJPanelLayout)
        genericProfileActionsJPanelLayout.SetHorizontalGroup(genericProfileActionsJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(genericProfileActionsJPanelLayout.CreateSequentialGroup().AddComponent(this.GenericProfileActionsJScrollPane, javax_Swing.GroupLayout.PREFERRED_SIZE, 95, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.BlankGenericProfileActionJPanel, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)).AddGroup(genericProfileActionsJPanelLayout.CreateSequentialGroup().AddComponent(this.NewActionJButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.DeleteActionJButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.SaveJButton).AddContainerGap()))
        genericProfileActionsJPanelLayout.SetVerticalGroup(genericProfileActionsJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(genericProfileActionsJPanelLayout.CreateSequentialGroup().AddGroup(genericProfileActionsJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.NewActionJButton).AddComponent(this.DeleteActionJButton).AddComponent(this.SaveJButton)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(genericProfileActionsJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING, false).AddComponent(this.BlankGenericProfileActionJPanel, 0, 296, Short_MAX_VALUE).AddComponent(this.GenericProfileActionsJScrollPane, javax_Swing.GroupLayout.DEFAULT_SIZE, 296, Short_MAX_VALUE)).AddContainerGap(javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.GenericProfileActionsJPanel, javax_Swing.GroupLayout.PREFERRED_SIZE, 596, javax_Swing.GroupLayout.PREFERRED_SIZE).AddContainerGap(javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.GenericProfileActionsJPanel, javax_Swing.GroupLayout.PREFERRED_SIZE, 336, javax_Swing.GroupLayout.PREFERRED_SIZE))
}


func (this *GenericProfileActionsJPanel) SetModel() {
        var defaultListModelHelper *DefaultListModelHelper = this.GetGenericProfileActions().GetDefaultListModelHelper(); _ = defaultListModelHelper
        if defaultListModelHelper.GetListModel().GetSize() > 0 {
            this.GenericProfileActionsJList.SetModel(defaultListModelHelper.GetListModel())
        }
}


func (this *GenericProfileActionsJPanel) SaveJButtonActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "saveJButtonActionPerformed", e)
                }
            }()
            this.GetGenericProfileActions().Save()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GenericProfileActionsJPanel) UpdateProfileActionUI(profileActionNameString string) {
        this.BlankGenericProfileActionJPanel.RemoveAll()
        if profileActionNameString != "" {
            var genericProfileAction *GenericProfileAction = this.GetGenericProfileActions().GetAction(profileActionNameString); _ = genericProfileAction
            this.SetGenericProfileActionJPanel(NewGenericProfileActionJPanel())
            this.GetGenericProfileActionJPanel().SetGenericProfileAction(genericProfileAction)
            this.GetGenericProfileActionJPanel().UpdateProfileActionUI()
            var layout *GroupLayout = any(this.BlankGenericProfileActionJPanel.GetLayout()).(*GroupLayout); _ = layout
            layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 395, Short_MAX_VALUE).AddComponent(this.GetGenericProfileActionJPanel(), javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))
            layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 256, Short_MAX_VALUE).AddComponent(this.GetGenericProfileActionJPanel(), javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))
            this.LogUtil.PutF("Setting Configuration Profile Actions JPanel", this, "genericProfileActionsJListValueChanged")
        }
}


func (this *GenericProfileActionsJPanel) GenericProfileActionsJListValueChanged(evt *ListSelectionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "genericProfileActionsJListValueChanged", e)
                }
            }()
            var profileActionNameString string = any(this.GenericProfileActionsJList.GetSelectedValue()).(string); _ = profileActionNameString
            this.UpdateProfileActionUI(profileActionNameString)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GenericProfileActionsJPanel) OkNewProfileActionJButtonActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "okNewProfileActionJButtonActionPerformed", e)
                }
            }()
            var newActionName string = this.NewProfileActionJTextField.GetText(); _ = newActionName
            this.GetGenericProfileActions().Add(newActionName)
            this.SetModel()
            this.GenericProfileActionsJList.SetSelectedValue(newActionName, true)
            this.NewProfileActionJDialog.SetVisible(false)
            this.UpdateProfileActionUI(newActionName)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GenericProfileActionsJPanel) DeleteActionJButtonActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "deleteActionJButtonActionPerformed", e)
                }
            }()
            var anyType []any = this.GenericProfileActionsJList.GetSelectedValues(); _ = anyType
            for index := int(0); index < len(anyType); index++ {
                var nextSelectedValueString string = any(anyType[index]).(string); _ = nextSelectedValueString
                this.GetGenericProfileActions().Remove(nextSelectedValueString)
            }
            this.SelectFirstElement()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *GenericProfileActionsJPanel) SelectFirstElement() {
        this.SetModel()
        if this.GenericProfileActionsJList.GetModel().GetSize() > 0 {
            this.GenericProfileActionsJList.SetSelectedValue(this.GenericProfileActionsJList.GetModel().GetElementAt(0), true)
            this.UpdateProfileActionUI(any(this.GenericProfileActionsJList.GetModel().GetElementAt(0)).(string))
        } else {
            this.BlankGenericProfileActionJPanel.RemoveAll()
            this.BlankGenericProfileActionJPanel.Repaint()
        }
}


func (this *GenericProfileActionsJPanel) GetGenericProfileActionJPanel() *GenericProfileActionJPanel {
        return this.GenericProfileActionJPanel
}


func (this *GenericProfileActionsJPanel) SetGenericProfileActionJPanel(genericProfileActionJPanel *GenericProfileActionJPanel) {
        this.GenericProfileActionJPanel = genericProfileActionJPanel
}


func (this *GenericProfileActionsJPanel) GetGenericProfileActions() *GenericProfileActions {
        return this.GenericProfileActions
}


func (this *GenericProfileActionsJPanel) SetGenericProfileActions(genericProfileActions *GenericProfileActions) {
        this.GenericProfileActions = genericProfileActions
}


func (this *GenericProfileActionsJPanel) NewActionJButtonActionPerformed(evt *ActionEvent) {
        this.NewProfileActionJDialog.SetVisible(true)
}

