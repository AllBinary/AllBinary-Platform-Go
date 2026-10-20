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

type ImageActionScriptOutputJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CapturedImageActionScriptOutputInterface ImageActionScriptOutputInterface
    CapturedImageActionJButton *JButton
    CapturedImageActionJDialog *JDialog
    DisplayJCheckBox *JCheckBox
    ItemTypesJList *JList
    JLabel1 *JLabel
    JScrollPane1 *JScrollPane
    OkJButton *JButton
    SaveJCheckBox *JCheckBox
    TitleJLabel *JLabel
}


func (this *ImageActionScriptOutputJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
}


func NewImageActionScriptOutputJPanel(capturedImageActionScriptOutputInterface ImageActionScriptOutputInterface) *ImageActionScriptOutputJPanel {
    this := &ImageActionScriptOutputJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.InitComponents()
        this.CapturedImageActionScriptOutputInterface = capturedImageActionScriptOutputInterface
        this.Set()
    return this
}


func (this *ImageActionScriptOutputJPanel) Set() {
        this.SetImageTypeJlistModel()
        var imageTypes *ImageTypes = this.CapturedImageActionScriptOutputInterface.GetImageTypes(); _ = imageTypes
        var vector IBasicArrayList = imageTypes.GetVector(); _ = vector
        var selectedArray []int = make([]int, vector.Size()); _ = selectedArray
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var imageType *InputImageType = any(vector.Get(index)).(*InputImageType); _ = imageType
            selectedArray[Java_PostIncrement(&index)] = imageType.GetIndex()
        }
        this.ItemTypesJList.SetSelectedIndices(selectedArray)
        if this.CapturedImageActionScriptOutputInterface.IsSaved() {
            this.SaveJCheckBox.SetSelected(true)
        }
        if this.CapturedImageActionScriptOutputInterface.IsSaved() {
            this.DisplayJCheckBox.SetSelected(true)
        }
        this.CapturedImageActionScriptOutputInterface.Log()
}


func (this *ImageActionScriptOutputJPanel) Update() {
        var anyType []any = this.ItemTypesJList.GetSelectedValues(); _ = anyType
        var imageTypes *ImageTypes = ImageTypes_ValueOf(anyType); _ = imageTypes
        this.CapturedImageActionScriptOutputInterface.SetImageTypes(imageTypes)
        if this.SaveJCheckBox.IsSelected() {
            this.CapturedImageActionScriptOutputInterface.SetSaved(true)
        } else {
            this.CapturedImageActionScriptOutputInterface.SetSaved(false)
        }
        if this.DisplayJCheckBox.IsSelected() {
            this.CapturedImageActionScriptOutputInterface.SetDisplay(true)
        } else {
            this.CapturedImageActionScriptOutputInterface.SetDisplay(false)
        }
        this.CapturedImageActionScriptOutputInterface.Log()
}


func (this *ImageActionScriptOutputJPanel) SetImageTypeJlistModel() {
        var defaultListModel *DefaultListModel = NewDefaultListModel(); _ = defaultListModel
        var vector IBasicArrayList = InputImageType_GetAllAsVector(); _ = vector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var imageType *InputImageType = any(vector.Get(index)).(*InputImageType); _ = imageType
            defaultListModel.AddElement(imageType.GetName())
        }
        this.ItemTypesJList.SetModel(defaultListModel)
}


func (this *ImageActionScriptOutputJPanel) GetCapturedImageActionJDialog() *JDialog {
        return this.CapturedImageActionJDialog
}


func (this *ImageActionScriptOutputJPanel) SetCapturedImageActionJDialog(capturedImageActionJDialog *JDialog) {
        this.CapturedImageActionJDialog = capturedImageActionJDialog
}


func (this *ImageActionScriptOutputJPanel) InitComponents() {
        this.CapturedImageActionJDialog = NewJDialog()
        this.TitleJLabel = NewJLabel()
        this.OkJButton = NewJButton()
        this.SaveJCheckBox = NewJCheckBox()
        this.DisplayJCheckBox = NewJCheckBox()
        this.JScrollPane1 = NewJScrollPane()
        this.ItemTypesJList = NewJList()
        this.CapturedImageActionJButton = NewJButton()
        this.JLabel1 = NewJLabel()
        this.CapturedImageActionJDialog.SetMinimumSize(NewDimension(250, 250))
        this.TitleJLabel.SetText("Image Output Options")
        this.OkJButton.SetText("OK")
        this.OkJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageActionScriptOutputJPanel) ActionPerformed(evt *ActionEvent) {
                this.OkJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.SaveJCheckBox.SetText("Save To File")
        this.SaveJCheckBox.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.SaveJCheckBox.SetMargin(NewInsets(0, 0, 0, 0))
        this.DisplayJCheckBox.SetText("Display")
        this.DisplayJCheckBox.SetBorder(javax_Swing.BorderFactory.CreateEmptyBorder(0, 0, 0, 0))
        this.DisplayJCheckBox.SetMargin(NewInsets(0, 0, 0, 0))
        this.DisplayJCheckBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageActionScriptOutputJPanel) ActionPerformed(evt *ActionEvent) {
                this.DisplayJCheckBoxActionPerformed(evt)
        }
*/ NewActionListener())
        this.ItemTypesJList.SetModel(/* Java anonymous class body: implement it as a named type
            this.Strings = []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"}
            

        func (this *ImageActionScriptOutputJPanel) GetSize() int {
                return len(strings)
        }
            

        func (this *ImageActionScriptOutputJPanel) GetElementAt(i int) any {
                return strings[i]
        }
*/ NewAbstractListModel())
        this.JScrollPane1.SetViewportView(this.ItemTypesJList)
        var capturedImageActionJDialogLayout *GroupLayout = NewGroupLayout(this.CapturedImageActionJDialog.GetContentPane()); _ = capturedImageActionJDialogLayout
        this.CapturedImageActionJDialog.GetContentPane().SetLayout(capturedImageActionJDialogLayout)
        capturedImageActionJDialogLayout.SetHorizontalGroup(capturedImageActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(capturedImageActionJDialogLayout.CreateSequentialGroup().AddGroup(capturedImageActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(capturedImageActionJDialogLayout.CreateSequentialGroup().AddGap(61, 61, 61).AddComponent(this.TitleJLabel)).AddGroup(capturedImageActionJDialogLayout.CreateSequentialGroup().AddContainerGap().AddGroup(capturedImageActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(capturedImageActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.DisplayJCheckBox).AddComponent(this.SaveJCheckBox, javax_Swing.GroupLayout.Alignment.TRAILING)).AddGroup(capturedImageActionJDialogLayout.CreateSequentialGroup().AddGap(84, 84, 84).AddComponent(this.OkJButton)))).AddGroup(capturedImageActionJDialogLayout.CreateSequentialGroup().AddContainerGap().AddComponent(this.JScrollPane1, javax_Swing.GroupLayout.DEFAULT_SIZE, 216, Short_MAX_VALUE))).AddContainerGap()))
        capturedImageActionJDialogLayout.SetVerticalGroup(capturedImageActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(capturedImageActionJDialogLayout.CreateSequentialGroup().AddComponent(this.TitleJLabel).AddGap(5, 5, 5).AddComponent(this.JScrollPane1, javax_Swing.GroupLayout.DEFAULT_SIZE, 105, Short_MAX_VALUE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.SaveJCheckBox).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.DisplayJCheckBox).AddGap(9, 9, 9).AddComponent(this.OkJButton).AddContainerGap()))
        this.SetMinimumSize(NewDimension(100, 0))
        this.CapturedImageActionJButton.SetText("Edit")
        this.CapturedImageActionJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageActionScriptOutputJPanel) ActionPerformed(evt *ActionEvent) {
                this.CapturedImageActionJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel1.SetText("Capture Image:")
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JLabel1).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddComponent(this.CapturedImageActionJButton)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.CapturedImageActionJButton)))
}


func (this *ImageActionScriptOutputJPanel) DisplayJCheckBoxActionPerformed(evt *ActionEvent) {
}


func (this *ImageActionScriptOutputJPanel) OkJButtonActionPerformed(evt *ActionEvent) {
        this.Update()
        this.GetCapturedImageActionJDialog().SetVisible(false)
}


func (this *ImageActionScriptOutputJPanel) CapturedImageActionJButtonActionPerformed(evt *ActionEvent) {
        this.GetCapturedImageActionJDialog().SetVisible(true)
}

