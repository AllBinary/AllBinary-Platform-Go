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

type CreateIconImagesJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageProcessorInput *ImageProcessorInput
    AboveJButton *JButton
}

// Implements ImageProcessorInputCompositeInterface


func (this *CreateIconImagesJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewCreateIconImagesJPanel(imageProcessorInput *ImageProcessorInput) *CreateIconImagesJPanel {
    this := &CreateIconImagesJPanel{}
        this.JPanel = *NewJPanel()
        this.initFields()
        this.InitComponents()
        this.ImageProcessorInput = imageProcessorInput
        var numberStringArray []string = make([]string, 101); _ = numberStringArray
        numberStringArray[0] = fmt.Sprint(-1)
        for index := int(1); index < 100; index++ {
            numberStringArray[index] = fmt.Sprint(index)
        }
        var araster *Raster = this.ImageProcessorInput.GetBufferedImageArray()[0].GetAlphaRaster(); _ = araster
        if araster == nil {
            fmt.Println("there is no Alpha channel!!!!!!!!!")
        } else {
            fmt.Println("Alpha channel found !")
        }
    return this
}


func (this *CreateIconImagesJPanel) Process() {
        /* Java anonymous class body: implement it as a named type
            

        func (this *CreateIconImagesJPanel) Run() {
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
                    var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
                    var imageProcessorInput *ImageProcessorInput = this.GetImageProcessorInput(); _ = imageProcessorInput
                    var files []*File = imageProcessorInput.GetFiles(); _ = files
                    var iconWidthList IBasicArrayList = NewBasicArrayListD(); _ = iconWidthList
                    var iconHeightList IBasicArrayList = NewBasicArrayListD(); _ = iconHeightList
                    iconWidthList.Add(864)
                    iconHeightList.Add(864)
                    var ANDROID string = "android-icon-"; _ = ANDROID
                    var IOS string = "ios-icon-"; _ = IOS
                    var size int = iconWidthList.Size(); _ = size
                    var width int = 0; _ = width
                    var height int = 0; _ = height
                    for index2 := int(0); index2 < size; index2++ {
                        width = int((any(iconWidthList.Get(index2)).(int)))
                        height = int((any(iconHeightList.Get(index2)).(int)))
                        var generatedBufferedImageArray []*BufferedImage = imageUtil.CreateBufferedImage3(imageProcessorInput.GetBufferedImageArray(), width, height, true); _ = generatedBufferedImageArray
                        var imagePersistanceUtil *ImagePersistanceUtil = ImagePersistanceUtil_GetInstance(); _ = imagePersistanceUtil
                        for index := int(0); index < len(generatedBufferedImageArray); index++ {
                            var filePath string = files[index].GetAbsolutePath(); _ = filePath
                            var endIndex int = strings.LastIndex(filePath, string('\\')); _ = endIndex
                            var path string = filePath[0:endIndex + 1]; _ = path
                            stringMaker.Delete(0, stringMaker.Length())
                            stringMaker.Append(path).Append(ANDROID).Appendint(width).Append(CanvasStrings_GetInstance()._PNG)
                            var newFilePath string = fmt.Sprint(stringMaker); _ = newFilePath
                            imagePersistanceUtil.SaveWithBatik(AbFile_CreateAbFile(newFilePath), generatedBufferedImageArray[index])
                            stringMaker.Delete(0, stringMaker.Length())
                            stringMaker.Append(path).Append(IOS).Appendint(width).Append(CanvasStrings_GetInstance()._PNG)
                            newFilePath = fmt.Sprint(stringMaker)
                            imagePersistanceUtil.SaveWithBatik(AbFile_CreateAbFile(newFilePath), generatedBufferedImageArray[index])
                        }
                    }
                    this.GetParent().Repaint()
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewThread().Start()
}


func (this *CreateIconImagesJPanel) GetImageProcessorInput() *ImageProcessorInput {
        return this.ImageProcessorInput
}


func (this *CreateIconImagesJPanel) SetImageProcessorInput(imageProcessorInput *ImageProcessorInput) {
        this.ImageProcessorInput = imageProcessorInput
}


func (this *CreateIconImagesJPanel) InitComponents() {
        this.AboveJButton = NewJButton()
        this.AboveJButton.SetText("Process")
        this.AboveJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *CreateIconImagesJPanel) ActionPerformed(evt *ActionEvent) {
                this.AboveJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.AboveJButton).AddGap(0, 335, Short_MAX_VALUE)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddContainerGap().AddComponent(this.AboveJButton).AddContainerGap(286, Short_MAX_VALUE)))
}


func (this *CreateIconImagesJPanel) AboveJButtonActionPerformed(evt *ActionEvent) {
        this.Process()
}

