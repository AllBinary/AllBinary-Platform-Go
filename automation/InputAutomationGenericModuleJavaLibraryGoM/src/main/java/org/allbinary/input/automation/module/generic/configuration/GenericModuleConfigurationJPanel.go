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

type GenericModuleConfigurationJPanel struct {
    JPanel
    LogUtil *LogUtil
    StringUtil *StringUtil
    CommonStrings *CommonStrings
    GenericProfiles *GenericProfiles
    SelectedGenericProfile *GenericProfile
    BlankGenericProfileActionsJPanel *JPanel
    ButtonGroup2 *ButtonGroup
    CaptureComparisonJCheckBox *JCheckBox
    CaptureMotionAnalysisJCheckBox *JCheckBox
    DeleteProfileJButton *JButton
    GenericJTabbedPane *JTabbedPane
    GenericProfilesJList *JList
    GenericProfilesJPanel *JPanel
    GenericProfilesJScrollPane *JScrollPane
    JLabel1 *JLabel
    JLabel2 *JLabel
    JLabel3 *JLabel
    JPanel1 *JPanel
    NewProfileJButton *JButton
    NewProfileJDialog *JDialog
    NewProfileNameJLabel *JLabel
    NewProfileNameJTextField *JTextField
    OkNewProfileJButton *JButton
    SavedCaptureJRadioButton *JRadioButton
    ScreenCaptureJRadioButton *JRadioButton
}

// Implements GenericModuleConfigurationInterface


func (this *GenericModuleConfigurationJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.StringUtil = StringUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewGenericModuleConfigurationJPanel() *GenericModuleConfigurationJPanel {
    this := &GenericModuleConfigurationJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
        this.SetGenericProfiles(NewGenericProfiles(GenericProfiles_DEFAULT_FILE))
        this.SelectFirstElement()
    return this
}


func (this *GenericModuleConfigurationJPanel) InitComponents() {
        this.NewProfileJDialog = NewJDialog()
        this.OkNewProfileJButton = NewJButton()
        this.NewProfileNameJLabel = NewJLabel()
        this.NewProfileNameJTextField = NewJTextField()
        this.ButtonGroup2 = NewButtonGroup()
        this.GenericJTabbedPane = NewJTabbedPane()
        this.GenericProfilesJPanel = NewJPanel()
        this.GenericProfilesJScrollPane = NewJScrollPane()
        this.GenericProfilesJList = NewJList()
        this.NewProfileJButton = NewJButton()
        this.DeleteProfileJButton = NewJButton()
        this.JPanel1 = NewJPanel()
        this.ScreenCaptureJRadioButton = NewJRadioButton()
        this.JLabel1 = NewJLabel()
        this.JLabel2 = NewJLabel()
        this.SavedCaptureJRadioButton = NewJRadioButton()
        this.CaptureMotionAnalysisJCheckBox = NewJCheckBox()
        this.CaptureComparisonJCheckBox = NewJCheckBox()
        this.JLabel3 = NewJLabel()
        this.BlankGenericProfileActionsJPanel = NewJPanel()
        this.NewProfileJDialog.SetMinimumSize(NewDimension(250, 150))
        this.NewProfileJDialog.SetResizable(false)
        this.OkNewProfileJButton.SetText("Ok")
        this.OkNewProfileJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericModuleConfigurationJPanel) ActionPerformed(evt *ActionEvent) {
                this.OkNewProfileJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.NewProfileNameJLabel.SetText("New Profile Name")
        var newProfileJDialogLayout *GroupLayout = NewGroupLayout(this.NewProfileJDialog.GetContentPane()); _ = newProfileJDialogLayout
        this.NewProfileJDialog.GetContentPane().SetLayout(newProfileJDialogLayout)
        newProfileJDialogLayout.SetHorizontalGroup(newProfileJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(newProfileJDialogLayout.CreateSequentialGroup().AddGroup(newProfileJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(newProfileJDialogLayout.CreateSequentialGroup().AddGap(80, 80, 80).AddComponent(this.NewProfileNameJLabel)).AddGroup(newProfileJDialogLayout.CreateSequentialGroup().AddGap(30, 30, 30).AddComponent(this.NewProfileNameJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 180, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddGroup(newProfileJDialogLayout.CreateSequentialGroup().AddGap(100, 100, 100).AddComponent(this.OkNewProfileJButton))).AddContainerGap(35, Short_MAX_VALUE)))
        newProfileJDialogLayout.SetVerticalGroup(newProfileJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(newProfileJDialogLayout.CreateSequentialGroup().AddGap(10, 10, 10).AddComponent(this.NewProfileNameJLabel).AddGap(16, 16, 16).AddComponent(this.NewProfileNameJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddGap(20, 20, 20).AddComponent(this.OkNewProfileJButton)))
        this.SetLayout(NewGridLayout(1, 0))
        this.SetMinimumSize(NewDimension(400, 300))
        this.SetPreferredSize(NewDimension(700, 700))
        this.GenericJTabbedPane.SetToolTipText(this.StringUtil.EMPTY_STRING)
        this.GenericJTabbedPane.SetMinimumSize(NewDimension(800, 600))
        this.GenericProfilesJPanel.SetMinimumSize(NewDimension(400, 300))
        this.GenericProfilesJPanel.SetPreferredSize(NewDimension(0, 0))
        this.GenericProfilesJList.AddListSelectionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericModuleConfigurationJPanel) ValueChanged(evt *ListSelectionEvent) {
                this.GenericProfilesJListValueChanged(evt)
        }
*/ NewListSelectionListener())
        this.GenericProfilesJScrollPane.SetViewportView(this.GenericProfilesJList)
        this.NewProfileJButton.SetText("New")
        this.NewProfileJButton.SetToolTipText("Creates a new profile")
        this.NewProfileJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericModuleConfigurationJPanel) ActionPerformed(evt *ActionEvent) {
                this.NewProfileJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.DeleteProfileJButton.SetText("Delete")
        this.DeleteProfileJButton.SetToolTipText("Deletes the selected profile")
        this.DeleteProfileJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericModuleConfigurationJPanel) ActionPerformed(evt *ActionEvent) {
                this.DeleteProfileJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JPanel1.SetPreferredSize(NewDimension(1000, 100))
        this.ScreenCaptureJRadioButton.SetText("Screen Captures")
        this.ScreenCaptureJRadioButton.SetToolTipText("Records the screen into memory for logical processing")
        this.ScreenCaptureJRadioButton.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.ScreenCaptureJRadioButton.SetMargin(NewInsets(0, 0, 0, 0))
        this.ScreenCaptureJRadioButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericModuleConfigurationJPanel) ActionPerformed(evt *ActionEvent) {
                this.ScreenCaptureJRadioButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel1.SetText("Profile Options")
        this.JLabel2.SetText("Process From:")
        this.SavedCaptureJRadioButton.SetText("Saved Images")
        this.SavedCaptureJRadioButton.SetToolTipText("Loads images from files for logical processing.  Usually for debugging your profile actions")
        this.SavedCaptureJRadioButton.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.SavedCaptureJRadioButton.SetMargin(NewInsets(0, 0, 0, 0))
        this.SavedCaptureJRadioButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericModuleConfigurationJPanel) ActionPerformed(evt *ActionEvent) {
                this.SavedCaptureJRadioButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.CaptureMotionAnalysisJCheckBox.SetText("Motion Analysis")
        this.CaptureMotionAnalysisJCheckBox.SetToolTipText("Uses image comparison and contraints to locate moving objects")
        this.CaptureMotionAnalysisJCheckBox.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.CaptureMotionAnalysisJCheckBox.SetMargin(NewInsets(0, 0, 0, 0))
        this.CaptureMotionAnalysisJCheckBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericModuleConfigurationJPanel) ActionPerformed(evt *ActionEvent) {
                this.CaptureMotionAnalysisJCheckBoxActionPerformed(evt)
        }
*/ NewActionListener())
        this.CaptureComparisonJCheckBox.SetText("Comparison")
        this.CaptureComparisonJCheckBox.SetToolTipText("Compares the last two input images for multi image logic")
        this.CaptureComparisonJCheckBox.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.CaptureComparisonJCheckBox.SetMargin(NewInsets(0, 0, 0, 0))
        this.CaptureComparisonJCheckBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericModuleConfigurationJPanel) ActionPerformed(evt *ActionEvent) {
                this.CaptureComparisonJCheckBoxActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel3.SetText("Image Processing:")
        this.JLabel3.SetToolTipText("Turns on/off various extended image processing")
        var jPanel1Layout *GroupLayout = NewGroupLayout(this.JPanel1); _ = jPanel1Layout
        this.JPanel1.SetLayout(jPanel1Layout)
        jPanel1Layout.SetHorizontalGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(jPanel1Layout.CreateSequentialGroup().AddGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.ScreenCaptureJRadioButton).AddComponent(this.JLabel1).AddComponent(this.JLabel2).AddComponent(this.SavedCaptureJRadioButton).AddComponent(this.CaptureMotionAnalysisJCheckBox).AddComponent(this.CaptureComparisonJCheckBox).AddComponent(this.JLabel3)).AddContainerGap(22, Short_MAX_VALUE)))
        jPanel1Layout.SetVerticalGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(jPanel1Layout.CreateSequentialGroup().AddComponent(this.JLabel1).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel2).AddGap(7, 7, 7).AddComponent(this.ScreenCaptureJRadioButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.SavedCaptureJRadioButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel3).AddGap(7, 7, 7).AddComponent(this.CaptureComparisonJCheckBox).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.CaptureMotionAnalysisJCheckBox).AddContainerGap(106, Short_MAX_VALUE)))
        var genericProfilesJPanelLayout *GroupLayout = NewGroupLayout(this.GenericProfilesJPanel); _ = genericProfilesJPanelLayout
        this.GenericProfilesJPanel.SetLayout(genericProfilesJPanelLayout)
        genericProfilesJPanelLayout.SetHorizontalGroup(genericProfilesJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(genericProfilesJPanelLayout.CreateSequentialGroup().AddGroup(genericProfilesJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(genericProfilesJPanelLayout.CreateSequentialGroup().AddComponent(this.NewProfileJButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.DeleteProfileJButton)).AddGroup(genericProfilesJPanelLayout.CreateSequentialGroup().AddComponent(this.GenericProfilesJScrollPane, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JPanel1, javax_Swing.GroupLayout.DEFAULT_SIZE, 119, Short_MAX_VALUE))).AddContainerGap()))
        genericProfilesJPanelLayout.SetVerticalGroup(genericProfilesJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(genericProfilesJPanelLayout.CreateSequentialGroup().AddGroup(genericProfilesJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.NewProfileJButton).AddComponent(this.DeleteProfileJButton)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(genericProfilesJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JPanel1, javax_Swing.GroupLayout.DEFAULT_SIZE, 246, Short_MAX_VALUE).AddComponent(this.GenericProfilesJScrollPane, javax_Swing.GroupLayout.DEFAULT_SIZE, 246, Short_MAX_VALUE))))
        this.GenericJTabbedPane.AddTab("Profiles", this.GenericProfilesJPanel)
        var blankGenericProfileActionsJPanelLayout *GroupLayout = NewGroupLayout(this.BlankGenericProfileActionsJPanel); _ = blankGenericProfileActionsJPanelLayout
        this.BlankGenericProfileActionsJPanel.SetLayout(blankGenericProfileActionsJPanelLayout)
        blankGenericProfileActionsJPanelLayout.SetHorizontalGroup(blankGenericProfileActionsJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 395, Short_MAX_VALUE))
        blankGenericProfileActionsJPanelLayout.SetVerticalGroup(blankGenericProfileActionsJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 275, Short_MAX_VALUE))
        this.GenericJTabbedPane.AddTab("Actions", this.BlankGenericProfileActionsJPanel)
        this.Add(this.GenericJTabbedPane)
}


func (this *GenericModuleConfigurationJPanel) SavedCaptureJRadioButtonActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "savedCaptureJRadioButtonActionPerformed", e)
                }
            }()
            if this.SavedCaptureJRadioButton.IsSelected() {
                this.LogUtil.PutF("Selected", this, "savedCaptureJRadioButtonActionPerformed")
                this.GetSelectedGenericProfile().Add(GenericProfileDataWorkerType_SAVED_CAPTURE)
                this.GetSelectedGenericProfile().Remove(GenericProfileDataWorkerType_SCREEN_CAPTURE)
                this.ScreenCaptureJRadioButton.SetSelected(false)
                this.GetGenericProfiles().Save()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GenericModuleConfigurationJPanel) ScreenCaptureJRadioButtonActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "screenCaptureJRadioButtonActionPerformed", e)
                }
            }()
            if this.ScreenCaptureJRadioButton.IsSelected() {
                this.LogUtil.PutF("Selected", this, "screenCaptureJRadioButtonActionPerformed")
                this.GetSelectedGenericProfile().Add(GenericProfileDataWorkerType_SCREEN_CAPTURE)
                this.GetSelectedGenericProfile().Remove(GenericProfileDataWorkerType_SAVED_CAPTURE)
                this.SavedCaptureJRadioButton.SetSelected(false)
                this.GetGenericProfiles().Save()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GenericModuleConfigurationJPanel) CaptureMotionAnalysisJCheckBoxActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "captureComparisonJCheckBoxActionPerformed", e)
                }
            }()
            if this.CaptureMotionAnalysisJCheckBox.IsSelected() {
                this.GetSelectedGenericProfile().Add(GenericProfileDataWorkerType_MOTION)
            } else {
                this.GetSelectedGenericProfile().Remove(GenericProfileDataWorkerType_MOTION)
            }
            this.GetGenericProfiles().Save()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GenericModuleConfigurationJPanel) CaptureComparisonJCheckBoxActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "captureComparisonJCheckBoxActionPerformed", e)
                }
            }()
            if this.CaptureComparisonJCheckBox.IsSelected() {
                this.GetSelectedGenericProfile().Add(GenericProfileDataWorkerType_COMPARISON)
            } else {
                this.GetSelectedGenericProfile().Remove(GenericProfileDataWorkerType_COMPARISON)
            }
            this.GetGenericProfiles().Save()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *GenericModuleConfigurationJPanel) UpdateProfileUI(profileNameString string) {
        this.GetBlankGenericProfileActionsJPanel().RemoveAll()
        if profileNameString != "" {
            var genericProfileActionsJPanel *GenericProfileActionsJPanel = NewGenericProfileActionsJPanel(); _ = genericProfileActionsJPanel
            this.SetSelectedGenericProfile(this.GetGenericProfiles().Get(profileNameString))
            this.GetSelectedGenericProfile().SetGenericProfileActions(NewGenericProfileActions(genericProfileActionsJPanel, profileNameString))
            genericProfileActionsJPanel.SetGenericProfileActions(this.GetSelectedGenericProfile().GetGenericProfileActions())
            genericProfileActionsJPanel.SelectFirstElement()
            var layout *GroupLayout = any(this.GetBlankGenericProfileActionsJPanel().GetLayout()).(*GroupLayout); _ = layout
            layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 395, Short_MAX_VALUE).AddComponent(genericProfileActionsJPanel, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))
            layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 256, Short_MAX_VALUE).AddComponent(genericProfileActionsJPanel, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))
            this.UpdateProfileOptions()
            this.LogUtil.PutF("Setting Configuration Profile JPanel", this, "updateProfileUI")
        } else {
            this.LogUtil.PutF("Null Configuration Profile", this, "updateProfileUI")
        }
}


// throws Exception
func (this *GenericModuleConfigurationJPanel) SelectFirstElement() {
        this.GenericProfilesJList.SetModel(this.GetGenericProfiles().GetDefaultListModelHelper().GetListModel())
        if this.GenericProfilesJList.GetModel().GetSize() > 0 {
            this.GenericProfilesJList.SetSelectedValue(this.GenericProfilesJList.GetModel().GetElementAt(0), true)
            this.UpdateProfileUI(any(this.GenericProfilesJList.GetModel().GetElementAt(0)).(string))
        }
}


// throws Exception
func (this *GenericModuleConfigurationJPanel) UpdateProfileOptions() {
        this.LogUtil.PutF(this.CommonStrings.START, this, "updateProfileOptions")
        var genericProfile *GenericProfile = this.GetSelectedGenericProfile(); _ = genericProfile
        var vector IBasicArrayList = genericProfile.GetGenericProfileDataWorkerTypeVector(); _ = vector
        this.ScreenCaptureJRadioButton.SetSelected(false)
        this.SavedCaptureJRadioButton.SetSelected(false)
        this.CaptureComparisonJCheckBox.SetSelected(false)
        this.CaptureMotionAnalysisJCheckBox.SetSelected(false)
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var genericProfileDataWorkerType IGenericProfileDataWorkerType = any(vector.Get(index)).(IGenericProfileDataWorkerType); _ = genericProfileDataWorkerType
            this.LogUtil.PutF(genericProfile.GetName() + " has GenericProfileDataWorkerType: " + fmt.Sprint(genericProfileDataWorkerType), this, "updateProfileOptions")
            if genericProfileDataWorkerType == GenericProfileDataWorkerType_SCREEN_CAPTURE {
                this.LogUtil.PutF("Screen Capture Selected", this, "updateProfileOptions")
                this.ScreenCaptureJRadioButton.SetSelected(true)
            } else if genericProfileDataWorkerType == GenericProfileDataWorkerType_SAVED_CAPTURE {
                this.LogUtil.PutF("Saved Capture Selected", this, "updateProfileOptions")
                this.SavedCaptureJRadioButton.SetSelected(true)
            } else if genericProfileDataWorkerType == GenericProfileDataWorkerType_COMPARISON {
                this.LogUtil.PutF("Checking Comparison Box", this, "updateProfileOptions")
                this.CaptureComparisonJCheckBox.SetSelected(true)
            } else if genericProfileDataWorkerType == GenericProfileDataWorkerType_MOTION {
                this.LogUtil.PutF("Checking Motion Box", this, "updateProfileOptions")
                this.CaptureMotionAnalysisJCheckBox.SetSelected(true)
            }
        }
        if !this.ScreenCaptureJRadioButton.IsSelected() && !this.SavedCaptureJRadioButton.IsSelected() {
            this.ScreenCaptureJRadioButton.SetSelected(true)
            this.GetSelectedGenericProfile().Add(GenericProfileDataWorkerType_SCREEN_CAPTURE)
        }
        this.LogUtil.PutF(this.CommonStrings.END, this, "updateProfileOptions")
}


func (this *GenericModuleConfigurationJPanel) GetBlankGenericProfileActionsJPanel() *JPanel {
        return this.BlankGenericProfileActionsJPanel
}


func (this *GenericModuleConfigurationJPanel) SetBlankGenericProfileActionsJPanel(blankGenericProfileActionsJPanel *JPanel) {
        this.BlankGenericProfileActionsJPanel = blankGenericProfileActionsJPanel
}


func (this *GenericModuleConfigurationJPanel) GetGenericProfiles() *GenericProfiles {
        return this.GenericProfiles
}


func (this *GenericModuleConfigurationJPanel) SetGenericProfiles(genericProfiles *GenericProfiles) {
        this.GenericProfiles = genericProfiles
}


func (this *GenericModuleConfigurationJPanel) GetSelectedGenericProfile() *GenericProfile {
        return this.SelectedGenericProfile
}


func (this *GenericModuleConfigurationJPanel) SetSelectedGenericProfile(selectedGenericProfile *GenericProfile) {
        this.SelectedGenericProfile = selectedGenericProfile
}


func (this *GenericModuleConfigurationJPanel) GenericProfilesJListValueChanged(evt *ListSelectionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "genericProfilesJListValueChanged", e)
                }
            }()
            var profileNameString string = any(this.GenericProfilesJList.GetSelectedValue()).(string); _ = profileNameString
            this.UpdateProfileUI(profileNameString)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GenericModuleConfigurationJPanel) DeleteProfileJButtonActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "deleteProfileJButtonActionPerformed", e)
                }
            }()
            var anyType []any = this.GenericProfilesJList.GetSelectedValues(); _ = anyType
            for index := int(0); index < len(anyType); index++ {
                var nextSelectedValueString string = any(anyType[index]).(string); _ = nextSelectedValueString
                this.GetGenericProfiles().Remove(nextSelectedValueString)
            }
            this.SelectFirstElement()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GenericModuleConfigurationJPanel) OkNewProfileJButtonActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "okNewProfileJButtonActionPerformed", e)
                }
            }()
            var newProfileName string = this.NewProfileNameJTextField.GetText(); _ = newProfileName
            this.GetGenericProfiles().Add(newProfileName)
            this.GenericProfilesJList.SetModel(this.GetGenericProfiles().GetDefaultListModelHelper().GetListModel())
            this.GenericProfilesJList.SetSelectedValue(newProfileName, true)
            this.NewProfileJDialog.SetVisible(false)
            this.UpdateProfileUI(newProfileName)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GenericModuleConfigurationJPanel) NewProfileJButtonActionPerformed(evt *ActionEvent) {
        this.NewProfileJDialog.SetVisible(true)
}

