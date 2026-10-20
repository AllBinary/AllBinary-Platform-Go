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

type RotationImageJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageStrings *ImageStrings
    ImageProcessorInput *ImageProcessorInput
    Result *BufferedImage
    GenerateJButton *JButton
    JLabel1 *JLabel
    JLabel2 *JLabel
    JLabel3 *JLabel
    JPanel1 *JPanel
    JPanel2 *JPanel
    JPanel3 *JPanel
    TotalAngleJComboBox *JComboBox
    WriteOverOriginalJCheckBox *JCheckBox
}

// Implements ImageProcessedVisitor


func (this *RotationImageJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ImageStrings = ImageStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
}


func NewRotationImageJPanel(imageProcessorInput *ImageProcessorInput) *RotationImageJPanel {
    this := &RotationImageJPanel{}
        this.JPanel = *NewJPanel()
        this.initFields()
        this.InitComponents()
        this.ImageProcessorInput = imageProcessorInput
    return this
}


func (this *RotationImageJPanel) Process() {
        /* Java anonymous class body: implement it as a named type
            

        func (this *RotationImageJPanel) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                        }
                    }()
                    var angleAsString string = any(this.TotalAngleJComboBox.GetSelectedItem()).(string); _ = angleAsString
                    ImageArrayRotationUtil_GetInstance().Process(this.GetImageProcessorInput(), angleAsString, this)
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewThread().Start()
}


func (this *RotationImageJPanel) InitComponents() {
        this.JPanel1 = /* Java anonymous class body: implement it as a named type
            

        func (this *RotationImageJPanel) Paint(graphics IGraphics) {
                if this.Result != nil {
                    graphics.DrawImage(this.Result, 0, 0, this.Result.GetWidth(nil), this.Result.GetHeight(nil), nil)
                }
        }
*/ NewJPanel()
        this.JPanel2 = NewJPanel()
        this.JLabel1 = NewJLabel()
        this.GenerateJButton = NewJButton()
        this.TotalAngleJComboBox = NewJComboBox()
        this.JLabel3 = NewJLabel()
        this.WriteOverOriginalJCheckBox = NewJCheckBox()
        this.JPanel3 = NewJPanel()
        this.JLabel2 = NewJLabel()
        var jPanel1Layout *GroupLayout = NewGroupLayout(this.JPanel1); _ = jPanel1Layout
        this.JPanel1.SetLayout(jPanel1Layout)
        jPanel1Layout.SetHorizontalGroup(jPanel1Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(0, 436, Short_MAX_VALUE))
        jPanel1Layout.SetVerticalGroup(jPanel1Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(0, 227, Short_MAX_VALUE))
        this.GenerateJButton.SetText("Generate")
        this.GenerateJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *RotationImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.GenerateJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.TotalAngleJComboBox.SetModel(NewDefaultComboBoxModel([]string{this.CommonStrings.UP, this.CommonStrings.DOWN, "-90", "90"}))
        this.TotalAngleJComboBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *RotationImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.TotalAngleJComboBoxActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel3.SetText("Total Angle:")
        this.WriteOverOriginalJCheckBox.SetText("Write Over Original")
        var jPanel2Layout *GroupLayout = NewGroupLayout(this.JPanel2); _ = jPanel2Layout
        this.JPanel2.SetLayout(jPanel2Layout)
        jPanel2Layout.SetHorizontalGroup(jPanel2Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(jPanel2Layout.CreateSequentialGroup().Add(8, 8, 8).Add(this.JLabel1).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.JLabel3).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.TotalAngleJComboBox, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.WriteOverOriginalJCheckBox).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.GenerateJButton).AddContainerGap(org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
        jPanel2Layout.SetVerticalGroup(jPanel2Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(jPanel2Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.BASELINE).Add(this.JLabel1).Add(this.GenerateJButton).Add(this.TotalAngleJComboBox, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).Add(this.JLabel3).Add(this.WriteOverOriginalJCheckBox)))
        this.JLabel2.SetText("Results:")
        var jPanel3Layout *GroupLayout = NewGroupLayout(this.JPanel3); _ = jPanel3Layout
        this.JPanel3.SetLayout(jPanel3Layout)
        jPanel3Layout.SetHorizontalGroup(jPanel3Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(jPanel3Layout.CreateSequentialGroup().Add(this.JLabel2).AddContainerGap(377, Short_MAX_VALUE)))
        jPanel3Layout.SetVerticalGroup(jPanel3Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(this.JLabel2))
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(org_Jdesktop.Layout.GroupLayout.TRAILING, this.JPanel2, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).Add(layout.CreateSequentialGroup().AddContainerGap().Add(this.JPanel3, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddContainerGap()).Add(this.JPanel1, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE))
        layout.SetVerticalGroup(layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(layout.CreateSequentialGroup().Add(this.JPanel2, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.UNRELATED).Add(this.JPanel3, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.JPanel1, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
}


func (this *RotationImageJPanel) GenerateJButtonActionPerformed(evt *ActionEvent) {
        this.Process()
}


func (this *RotationImageJPanel) TotalAngleJComboBoxActionPerformed(evt *ActionEvent) {
}


func (this *RotationImageJPanel) GetImageProcessorInput() *ImageProcessorInput {
        return this.ImageProcessorInput
}


func (this *RotationImageJPanel) SetImageProcessorInput(imageProcessorInput *ImageProcessorInput) {
        this.ImageProcessorInput = imageProcessorInput
}


// throws IOException
func (this *RotationImageJPanel) Visit(generatedBufferedImage *BufferedImage, name string, index int) {
        this.Result = generatedBufferedImage
        this.GetParent().Repaint()
        var file *File = this.ImageProcessorInput.GetFiles()[index]; _ = file
        if !this.WriteOverOriginalJCheckBox.IsSelected() {
            var filePath string = file.GetAbsolutePath(); _ = filePath
            var extensionIndex int = strings.Index(filePath, string(this.ImageStrings.PNG_EXTENSION)); _ = extensionIndex
            filePath = fmt.Sprint(NewStringMaker().Append(filePath[0:extensionIndex]).Append(CommonSeps_GetInstance().UNDERSCORE).Append(name).Append(this.ImageStrings.PNG_EXTENSION))
            file = NewFile(filePath)
        }
        var isWritten bool = ImageIO_Write(any(this.Result).(*RenderedImage), this.ImageStrings.PNG, file); _ = isWritten
        this.LogUtil.PutF("File: " + fmt.Sprint(file) + " Wrote: " + fmt.Sprint(isWritten), this, this.CommonStrings.RUN)
}

