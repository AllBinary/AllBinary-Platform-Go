/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type SpriteSplitterImageJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CommonSeps *CommonSeps
    ImageStrings *ImageStrings
    ImageProcessorInput *ImageProcessorInput
    Result *BufferedImage
    GenerateJButton *JButton
    HeightReductionTextField *JTextField
    IncreaseHeightTextField *JTextField
    IncreaseWidthTextField *JTextField
    JLabel1 *JLabel
    JLabel2 *JLabel
    JLabel3 *JLabel
    JLabel4 *JLabel
    JLabel5 *JLabel
    JLabel6 *JLabel
    JLabel7 *JLabel
    JLabel8 *JLabel
    JPanel1 *JPanel
    JPanel2 *JPanel
    SpriteTypeJComboBox *JComboBox[string]
    TotalAnimationsJComboBox *JComboBox[string]
    TotalFramesJComboBox *JComboBox[string]
    WidthReductionTextField *JTextField
    WriteOverOriginalJCheckBox *JCheckBox
}

// Implements ImageProcessedVisitor


func (this *SpriteSplitterImageJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.ImageStrings = ImageStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewSpriteSplitterImageJPanel(imageProcessorInput *ImageProcessorInput) *SpriteSplitterImageJPanel {
    this := &SpriteSplitterImageJPanel{}
        this.JPanel = *NewJPanel()
        this.initFields()
        this.InitComponents()
        this.ImageProcessorInput = imageProcessorInput
    return this
}


func (this *SpriteSplitterImageJPanel) Process() {
        /* Java anonymous class body: implement it as a named type
            

        func (this *SpriteSplitterImageJPanel) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                        }
                    }()
                    var spriteSplitterUtil *SpriteSplitterUtil = SpriteSplitterUtil_GetInstance(); _ = spriteSplitterUtil
                    var spriteType string = any(this.SpriteTypeJComboBox.GetSelectedItem()).(string); _ = spriteType
                    var totalFrames int = fmt.Sprint(any(this.TotalFramesJComboBox.GetSelectedItem()).(string)); _ = totalFrames
                    var totalAnimations int = fmt.Sprint(any(this.TotalAnimationsJComboBox.GetSelectedItem()).(string)); _ = totalAnimations
                    var widthReduction int = fmt.Sprint(this.WidthReductionTextField.GetText()); _ = widthReduction
                    var heightReduction int = fmt.Sprint(this.HeightReductionTextField.GetText()); _ = heightReduction
                    var increaseWidth int = fmt.Sprint(this.IncreaseWidthTextField.GetText()); _ = increaseWidth
                    var increaseHeight int = fmt.Sprint(this.IncreaseHeightTextField.GetText()); _ = increaseHeight
                    spriteSplitterUtil.Process(this.GetImageProcessorInput(), totalFrames, totalAnimations, widthReduction, heightReduction, increaseWidth, increaseHeight, spriteType, this)
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewThread().Start()
}


func (this *SpriteSplitterImageJPanel) InitComponents() {
        this.WriteOverOriginalJCheckBox = NewJCheckBox()
        this.WidthReductionTextField = NewJTextField()
        this.JLabel5 = NewJLabel()
        this.JPanel1 = /* Java anonymous class body: implement it as a named type
            

        func (this *SpriteSplitterImageJPanel) Paint(graphics IGraphics) {
                if this.Result != nil {
                    graphics.DrawImage(this.Result, 0, 0, this.Result.GetWidth(nil), this.Result.GetHeight(nil), nil)
                }
        }
*/ NewJPanel()
        this.JLabel4 = NewJLabel()
        this.JLabel6 = NewJLabel()
        this.HeightReductionTextField = NewJTextField()
        this.GenerateJButton = NewJButton()
        this.JPanel2 = NewJPanel()
        this.JLabel1 = NewJLabel()
        this.TotalFramesJComboBox = NewJComboBox()
        var size int = 70; _ = size
        var animationStringArray2 []string = make([]string, size - 1); _ = animationStringArray2
        for index := int(1); index < size; index++ {
            animationStringArray2[index - 1] = fmt.Sprint(index)
        }
        this.TotalFramesJComboBox.SetModel(NewDefaultComboBoxModel(animationStringArray2))
        this.JLabel2 = NewJLabel()
        this.TotalAnimationsJComboBox = NewJComboBox()
        var animationStringArray []string = make([]string, size - 1); _ = animationStringArray
        for index := int(1); index < size; index++ {
            animationStringArray[index - 1] = fmt.Sprint(index)
        }
        this.TotalAnimationsJComboBox.SetModel(NewDefaultComboBoxModel(animationStringArray))
        this.JLabel3 = NewJLabel()
        this.SpriteTypeJComboBox = NewJComboBox()
        this.SpriteTypeJComboBox.SetModel(NewDefaultComboBoxModel([]string{org_Allbinary.Media.Image.SpriteSplitterUtil.GetInstance().DIRECTIONAL_ANIMATIONS, org_Allbinary.Media.Image.SpriteSplitterUtil.GetInstance().HORIZONTAL_ANIMATIONS, org_Allbinary.Media.Image.SpriteSplitterUtil.GetInstance().HORIZONTAL_SPRITE}))
        this.JLabel7 = NewJLabel()
        this.IncreaseWidthTextField = NewJTextField()
        this.JLabel8 = NewJLabel()
        this.IncreaseHeightTextField = NewJTextField()
        this.WriteOverOriginalJCheckBox.SetText("Write Over Original")
        this.WriteOverOriginalJCheckBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *SpriteSplitterImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.WriteOverOriginalJCheckBoxActionPerformed(evt)
        }
*/ NewActionListener())
        this.WidthReductionTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.WidthReductionTextField.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *SpriteSplitterImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.WidthReductionTextFieldActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel5.SetText("Less Width:")
        this.JLabel4.SetText("Results:")
        var jPanel1Layout *GroupLayout = NewGroupLayout(this.JPanel1); _ = jPanel1Layout
        this.JPanel1.SetLayout(jPanel1Layout)
        jPanel1Layout.SetHorizontalGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(jPanel1Layout.CreateSequentialGroup().AddComponent(this.JLabel4).AddGap(0, 0, Short_MAX_VALUE)))
        jPanel1Layout.SetVerticalGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(jPanel1Layout.CreateSequentialGroup().AddComponent(this.JLabel4).AddGap(0, 184, Short_MAX_VALUE)))
        this.JLabel6.SetText("Less Height:")
        this.HeightReductionTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.GenerateJButton.SetText("Generate")
        this.GenerateJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *SpriteSplitterImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.GenerateJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel1.SetText("Row Total:")
        this.TotalFramesJComboBox.SetSelectedIndex(1)
        this.TotalFramesJComboBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *SpriteSplitterImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.TotalFramesJComboBoxActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel2.SetText("Column Total :")
        this.TotalAnimationsJComboBox.SetSelectedIndex(3)
        this.TotalAnimationsJComboBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *SpriteSplitterImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.TotalAnimationsJComboBoxActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel3.SetText("Type:")
        this.SpriteTypeJComboBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *SpriteSplitterImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.SpriteTypeJComboBoxActionPerformed(evt)
        }
*/ NewActionListener())
        var jPanel2Layout *GroupLayout = NewGroupLayout(this.JPanel2); _ = jPanel2Layout
        this.JPanel2.SetLayout(jPanel2Layout)
        jPanel2Layout.SetHorizontalGroup(jPanel2Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, jPanel2Layout.CreateSequentialGroup().AddContainerGap().AddComponent(this.JLabel1).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.TotalFramesJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, 48, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel2).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.TotalAnimationsJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, 39, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel3).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.SpriteTypeJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, 172, javax_Swing.GroupLayout.PREFERRED_SIZE).AddContainerGap(javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
        jPanel2Layout.SetVerticalGroup(jPanel2Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(jPanel2Layout.CreateSequentialGroup().AddContainerGap(javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddGroup(jPanel2Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.TotalFramesJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel2).AddComponent(this.TotalAnimationsJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel3).AddComponent(this.SpriteTypeJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))))
        this.JLabel7.SetText("More Width:")
        this.IncreaseWidthTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.IncreaseWidthTextField.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *SpriteSplitterImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.IncreaseWidthTextFieldActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel8.SetText("More Height:")
        this.IncreaseHeightTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JPanel2, javax_Swing.GroupLayout.PREFERRED_SIZE, 415, Short_MAX_VALUE).AddGap(78, 78, 78)).AddGroup(layout.CreateSequentialGroup().AddContainerGap().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JPanel1, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddGroup(layout.CreateSequentialGroup().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.WriteOverOriginalJCheckBox).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.GenerateJButton)).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JLabel5).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.WidthReductionTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 41, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel6).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.HeightReductionTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 46, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel7).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.IncreaseWidthTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 41, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel8).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.IncreaseHeightTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 46, javax_Swing.GroupLayout.PREFERRED_SIZE))).AddGap(0, 0, Short_MAX_VALUE))).AddContainerGap()))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JPanel2, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel7).AddComponent(this.IncreaseWidthTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 22, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel8).AddComponent(this.IncreaseHeightTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel5).AddComponent(this.WidthReductionTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 22, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel6).AddComponent(this.HeightReductionTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.WriteOverOriginalJCheckBox).AddComponent(this.GenerateJButton)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JPanel1, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
}


func (this *SpriteSplitterImageJPanel) WriteOverOriginalJCheckBoxActionPerformed(evt *ActionEvent) {
}


func (this *SpriteSplitterImageJPanel) GenerateJButtonActionPerformed(evt *ActionEvent) {
        this.Process()
}


func (this *SpriteSplitterImageJPanel) SpriteTypeJComboBoxActionPerformed(evt *ActionEvent) {
        var spriteSplitterUtil *SpriteSplitterUtil = SpriteSplitterUtil_GetInstance(); _ = spriteSplitterUtil
        this.JLabel1.SetText(spriteSplitterUtil.ANIMATIONS_LABELS[this.SpriteTypeJComboBox.GetSelectedIndex()])
        this.JLabel2.SetText(spriteSplitterUtil.ANIMATIONS_LABELS[this.SpriteTypeJComboBox.GetSelectedIndex() + 1])
}


func (this *SpriteSplitterImageJPanel) TotalAnimationsJComboBoxActionPerformed(evt *ActionEvent) {
}


func (this *SpriteSplitterImageJPanel) TotalFramesJComboBoxActionPerformed(evt *ActionEvent) {
}


func (this *SpriteSplitterImageJPanel) WidthReductionTextFieldActionPerformed(evt *ActionEvent) {
}


func (this *SpriteSplitterImageJPanel) IncreaseWidthTextFieldActionPerformed(evt *ActionEvent) {
}


func (this *SpriteSplitterImageJPanel) GetImageProcessorInput() *ImageProcessorInput {
        return this.ImageProcessorInput
}


func (this *SpriteSplitterImageJPanel) SetImageProcessorInput(imageProcessorInput *ImageProcessorInput) {
        this.ImageProcessorInput = imageProcessorInput
}


// throws IOException
func (this *SpriteSplitterImageJPanel) Visit(generatedBufferedImage *BufferedImage, name string, index int) {
        this.LogUtil.PutF(this.CommonStrings.VISIT, this, this.CommonStrings.VISIT)
        this.Result = generatedBufferedImage
        this.GetParent().Repaint()
        var file *File = this.ImageProcessorInput.GetFiles()[index]; _ = file
        if !this.WriteOverOriginalJCheckBox.IsSelected() {
            var filePath string = file.GetAbsolutePath(); _ = filePath
            var extensionIndex int = strings.Index(filePath, string(this.ImageStrings.PNG_EXTENSION)); _ = extensionIndex
            filePath = fmt.Sprint(NewStringMaker().Append(filePath[0:extensionIndex]).Append(this.CommonSeps.UNDERSCORE).Append(name).Append(this.ImageStrings.PNG_EXTENSION))
            file = NewFile(filePath)
        }
        var isWritten bool = ImageIO_Write(any(this.Result).(*RenderedImage), this.ImageStrings.PNG, file); _ = isWritten
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("File: ").Append(StringUtil_GetInstance().ToString(file)).Append(" Wrote: ").Appendboolean(isWritten)), this, this.CommonStrings.RUN)
}

