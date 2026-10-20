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

type ResizeImageJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageProcessorInput *ImageProcessorInput
    AboveJButton *JButton
    AdjustFor16AboveJButton *JButton
    AdjustFor16BelowJButton *JButton
    FloatPercentJTextField *JTextField
    JComboBox1 *JComboBox[string]
    JLabel1 *JLabel
    JLabel2 *JLabel
    JLabel3 *JLabel
    JTextField1 *JTextField
    JTextField2 *JTextField
    UpdateJButton *JButton
}

// Implements ImageProcessorInputCompositeInterface


func (this *ResizeImageJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewResizeImageJPanel(imageProcessorInput *ImageProcessorInput) *ResizeImageJPanel {
    this := &ResizeImageJPanel{}
        this.JPanel = *NewJPanel()
        this.initFields()
        this.InitComponents()
        this.ImageProcessorInput = imageProcessorInput
        var numberStringArray []string = make([]string, 101); _ = numberStringArray
        numberStringArray[0] = fmt.Sprint(-1)
        for index := int(1); index < 100; index++ {
            numberStringArray[index] = fmt.Sprint(index)
        }
        this.JComboBox1.SetModel(NewDefaultComboBoxModel(numberStringArray))
        var araster *Raster = this.ImageProcessorInput.GetBufferedImageArray()[0].GetAlphaRaster(); _ = araster
        if araster == nil {
            fmt.Println("there is no Alpha channel!!!!!!!!!")
        } else {
            fmt.Println("Alpha channel found !")
        }
    return this
}


func (this *ResizeImageJPanel) Process() {
        /* Java anonymous class body: implement it as a named type
            

        func (this *ResizeImageJPanel) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                        }
                    }()
                    var imageUtil *ImageUtil = ImageUtil_GetInstance(); _ = imageUtil
                    var percent int = fmt.Sprint(any(this.JComboBox1.GetSelectedItem()).(string)); _ = percent
                    var percentAsFloat float32 = func() float32 { value, _ := strconv.ParseFloat(this.FloatPercentJTextField.GetText(), 32); return float32(value) }(); _ = percentAsFloat
                    var imageProcessorInput *ImageProcessorInput = this.GetImageProcessorInput(); _ = imageProcessorInput
                    var files []*File = imageProcessorInput.GetFiles(); _ = files
                    var generatedBufferedImageArray []*BufferedImage = nil; _ = generatedBufferedImageArray
                    if int(percentAsFloat) != -1 {
                        generatedBufferedImageArray = imageUtil.CreateBufferedImage2(imageProcessorInput.GetBufferedImageArray(), percentAsFloat, true)
                    } else if int(percent) != -1 {
                        generatedBufferedImageArray = imageUtil.CreateBufferedImageForResize(imageProcessorInput.GetBufferedImageArray(), percent, true)
                    } else {
                        var width int = fmt.Sprint(any(this.JTextField1.GetText()).(string)); _ = width
                        var height int = fmt.Sprint(any(this.JTextField2.GetText()).(string)); _ = height
                        generatedBufferedImageArray = imageUtil.CreateBufferedImage3(imageProcessorInput.GetBufferedImageArray(), width, height, true)
                    }
                    var araster *Raster = generatedBufferedImageArray[0].GetAlphaRaster(); _ = araster
                    if araster == nil {
                        fmt.Println("No Alpha Channel In Result")
                    } else {
                        fmt.Println("Found Alpha Channel In Result")
                    }
                    var imagePersistanceUtil *ImagePersistanceUtil = ImagePersistanceUtil_GetInstance(); _ = imagePersistanceUtil
                    for index := int(0); index < len(generatedBufferedImageArray); index++ {
                        imagePersistanceUtil.SaveWithBatik(FileWrapperUtil_WrapFile(files[index]), generatedBufferedImageArray[index])
                    }
                    this.GetParent().Repaint()
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewThread().Start()
}


func (this *ResizeImageJPanel) Update() {
        /* Java anonymous class body: implement it as a named type
            

        func (this *ResizeImageJPanel) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                        }
                    }()
                    var imageProcessorInput *ImageProcessorInput = this.GetImageProcessorInput(); _ = imageProcessorInput
                    var size int = len(imageProcessorInput.GetBufferedImageArray()); _ = size
                    var bufferedImageArray []*BufferedImage = imageProcessorInput.GetBufferedImageArray(); _ = bufferedImageArray
                    var bufferedImage *BufferedImage; _ = bufferedImage
                    for index := int(0); index < size; index++ {
                        bufferedImage = bufferedImageArray[0]
                        this.JTextField1.SetText(fmt.Sprint(bufferedImage.GetWidth()))
                        this.JTextField2.SetText(fmt.Sprint(bufferedImage.GetHeight()))
                    }
                    this.GetParent().Repaint()
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewThread().Start()
}


func (this *ResizeImageJPanel) UpdateFor16Above() {
        /* Java anonymous class body: implement it as a named type
            

        func (this *ResizeImageJPanel) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                        }
                    }()
                    var imageProcessorInput *ImageProcessorInput = this.GetImageProcessorInput(); _ = imageProcessorInput
                    var size int = len(imageProcessorInput.GetBufferedImageArray()); _ = size
                    var bufferedImageArray []*BufferedImage = imageProcessorInput.GetBufferedImageArray(); _ = bufferedImageArray
                    var bufferedImage *BufferedImage; _ = bufferedImage
                    for index := int(0); index < size; index++ {
                        bufferedImage = bufferedImageArray[0]
                        this.JTextField1.SetText(fmt.Sprint(bufferedImage.GetWidth() / 16 * 16))
                        this.JTextField2.SetText(fmt.Sprint(bufferedImage.GetHeight() / 16 * 16))
                    }
                    this.GetParent().Repaint()
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewThread().Start()
}


func (this *ResizeImageJPanel) UpdateFor16Below() {
        /* Java anonymous class body: implement it as a named type
            

        func (this *ResizeImageJPanel) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                        }
                    }()
                    var imageProcessorInput *ImageProcessorInput = this.GetImageProcessorInput(); _ = imageProcessorInput
                    var size int = len(imageProcessorInput.GetBufferedImageArray()); _ = size
                    var bufferedImageArray []*BufferedImage = imageProcessorInput.GetBufferedImageArray(); _ = bufferedImageArray
                    var bufferedImage *BufferedImage; _ = bufferedImage
                    for index := int(0); index < size; index++ {
                        bufferedImage = bufferedImageArray[0]
                        this.JTextField1.SetText(fmt.Sprint(((bufferedImage.GetWidth() / 16) + 1) * 16))
                        this.JTextField2.SetText(fmt.Sprint(((bufferedImage.GetHeight() / 16) + 1) * 16))
                    }
                    this.GetParent().Repaint()
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewThread().Start()
}


func (this *ResizeImageJPanel) GetImageProcessorInput() *ImageProcessorInput {
        return this.ImageProcessorInput
}


func (this *ResizeImageJPanel) SetImageProcessorInput(imageProcessorInput *ImageProcessorInput) {
        this.ImageProcessorInput = imageProcessorInput
}


func (this *ResizeImageJPanel) InitComponents() {
        this.JLabel1 = NewJLabel()
        this.JComboBox1 = NewJComboBox()
        this.AboveJButton = NewJButton()
        this.JTextField1 = NewJTextField()
        this.JLabel2 = NewJLabel()
        this.JLabel3 = NewJLabel()
        this.JTextField2 = NewJTextField()
        this.UpdateJButton = NewJButton()
        this.FloatPercentJTextField = NewJTextField()
        this.AdjustFor16AboveJButton = NewJButton()
        this.AdjustFor16BelowJButton = NewJButton()
        this.JLabel1.SetText("Percent:")
        this.JComboBox1.SetModel(NewDefaultComboBoxModel([]string{"Item 1", "Item 2", "Item 3", "Item 4"}))
        this.AboveJButton.SetText("Process")
        this.AboveJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ResizeImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.AboveJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel2.SetText("Width:")
        this.JLabel3.SetText("Height:")
        this.UpdateJButton.SetText("Update")
        this.UpdateJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ResizeImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.UpdateJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.FloatPercentJTextField.SetText("-1.000000")
        this.FloatPercentJTextField.SetMinimumSize(NewDimension(120, 22))
        this.AdjustFor16AboveJButton.SetText("Adjust for 16 Above")
        this.AdjustFor16AboveJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ResizeImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.AdjustFor16AboveJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.AdjustFor16BelowJButton.SetText("Adjust for 16 Below")
        this.AdjustFor16BelowJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ResizeImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.AdjustFor16BelowJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddContainerGap().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JLabel1).AddComponent(this.JLabel2).AddComponent(this.JLabel3)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING, false).AddComponent(this.JComboBox1, 0, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddComponent(this.JTextField1).AddComponent(this.JTextField2)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.UNRELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.FloatPercentJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.AdjustFor16AboveJButton).AddComponent(this.AdjustFor16BelowJButton)).AddContainerGap(javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)).AddGroup(layout.CreateSequentialGroup().AddComponent(this.UpdateJButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.AboveJButton).AddGap(0, 250, Short_MAX_VALUE)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddContainerGap().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.JComboBox1, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.FloatPercentJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JTextField1, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel2).AddComponent(this.AdjustFor16AboveJButton)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel3).AddComponent(this.JTextField2, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.AdjustFor16BelowJButton)).AddGap(18, 18, 18).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.AboveJButton).AddComponent(this.UpdateJButton)).AddContainerGap(175, Short_MAX_VALUE)))
}


func (this *ResizeImageJPanel) AboveJButtonActionPerformed(evt *ActionEvent) {
        this.Process()
}


func (this *ResizeImageJPanel) UpdateJButtonActionPerformed(evt *ActionEvent) {
        this.Update()
}


func (this *ResizeImageJPanel) AdjustFor16AboveJButtonActionPerformed(evt *ActionEvent) {
        this.UpdateFor16Above()
}


func (this *ResizeImageJPanel) AdjustFor16BelowJButtonActionPerformed(evt *ActionEvent) {
        this.UpdateFor16Below()
}

