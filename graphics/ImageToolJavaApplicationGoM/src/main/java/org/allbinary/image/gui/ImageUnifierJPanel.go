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

type ImageUnifierJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageStrings *ImageStrings
    ImagesRatioUtil *ImagesRatioUtil
    ImageProcessorInput *ImageProcessorInput
    Icon []*ImageIconUnique
    ImageUnifierProperties *ImageUnifierProperties
    Result *BufferedImage
    IsImageFillIn bool
    AvgRatioJTextField *JTextField
    CellHeightJTextField *JTextField
    CellRatioJTextField *JTextField
    CellWidthJTextField *JTextField
    ColumnsJTextField *JTextField
    DownJButton *JButton
    FudgeItJButton *JButton
    ImageJList *JList
    JButton1 *JButton
    JButton2 *JButton
    JButtonOrder *JButton
    JLabel1 *JLabel
    JLabel2 *JLabel
    JLabel3 *JLabel
    JLabel4 *JLabel
    JLabel5 *JLabel
    JLabel6 *JLabel
    JPanel1 *JPanel
    JScrollPane1 *JScrollPane
    JScrollPane2 *JScrollPane
    RowsJTextField *JTextField
    UpJButton *JButton
}

// Implements ImageProcessorInputCompositeInterface


func (this *ImageUnifierJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ImageStrings = ImageStrings_GetInstance()
    this.ImagesRatioUtil = ImagesRatioUtil_GetInstance()
    
    
    
    
    this.IsImageFillIn = true
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewImageUnifierJPanel(imageProcessorInput *ImageProcessorInput) *ImageUnifierJPanel {
    this := &ImageUnifierJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
        this.SetImageProcessorInput(imageProcessorInput)
    return this
}


func (this *ImageUnifierJPanel) Init() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e)
                }
            }()
            var imageProcessorInput *ImageProcessorInput = this.GetImageProcessorInput(); _ = imageProcessorInput
            var bufferedImageArray []*BufferedImage = imageProcessorInput.GetBufferedImageArray(); _ = bufferedImageArray
            this.Icon = make([]*ImageIconUnique, len(bufferedImageArray))
            var defaultListModel *DefaultListModel = NewDefaultListModel(); _ = defaultListModel
            for index := int(0); index < len(bufferedImageArray); index++ {
                var width float64 = bufferedImageArray[index].GetWidth(); _ = width
                var height float64 = bufferedImageArray[index].GetHeight(); _ = height
                var newWidth int = 52; _ = newWidth
                var oldRatio float64 = width / height; _ = oldRatio
                var newHeight int = int((float64(newWidth) / oldRatio)); _ = newHeight
                var iconBufferedImage *BufferedImage = ImageUtil_GetInstance().CreateBufferedImage(bufferedImageArray[index], newWidth, newHeight, true); _ = iconBufferedImage
                this.Icon[index] = NewImageIconUnique(iconBufferedImage, index)
                defaultListModel.AddElement(this.Icon[index])
            }
            this.ImageJList.SetModel(defaultListModel)
            this.UpdateOnPropertiesChange()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageUnifierJPanel) UpdateImage_1() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "updateImage", e)
                }
            }()
            var bufferedImageArray []*BufferedImage = this.GetImageProcessorInput().GetBufferedImageArray(); _ = bufferedImageArray
            var tempBufferedImageArray []*BufferedImage = make([]*BufferedImage, len(bufferedImageArray)); _ = tempBufferedImageArray
            for index := int(0); index < len(bufferedImageArray); index++ {
                var indexedImageIcon *ImageIconUnique = any(this.ImageJList.GetModel().GetElementAt(index)).(*ImageIconUnique); _ = indexedImageIcon
                var bufferedImage *BufferedImage = ImageUtil_GetInstance().CreateBufferedImage(bufferedImageArray[indexedImageIcon.GetId()], this.ImageUnifierProperties.GetImageUnifierCell().GetWidth(), this.ImageUnifierProperties.GetImageUnifierCell().GetHeight(), false); _ = bufferedImage
                tempBufferedImageArray[index] = bufferedImage
            }
            var totalImages int = this.ImageUnifierProperties.GetRows() * this.ImageUnifierProperties.GetColumns(); _ = totalImages
            var averageRatio float64 = this.ImagesRatioUtil.GetAverage(tempBufferedImageArray, totalImages); _ = averageRatio
            var averageRatioString string = fmt.Sprint(float64(averageRatio)); _ = averageRatioString
            if len(averageRatioString) > 4 {
                averageRatioString = averageRatioString[0:4]
            }
            this.AvgRatioJTextField.SetText(averageRatioString)
            this.FudgeItJButton.SetEnabled(false)
            if !this.ImagesRatioUtil.IsEqual(tempBufferedImageArray, totalImages) {
                if this.IsImageFillIn {
                    this.FudgeItJButton.SetEnabled(true)
                }
            }
            this.UpdateImage(tempBufferedImageArray)
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of ImageUnifierJPanel.updateImage
func (this *ImageUnifierJPanel) UpdateImage(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            this.UpdateImage_1()
            return
        }
        if len(args) == 1 && Java_IsArg[[]*BufferedImage](args[0], pass) {
            this.UpdateImage_2(Java_Arg[[]*BufferedImage](args[0]))
            return
        }
    }
    panic("No Java overload of ImageUnifierJPanel.updateImage for the arguments")
}


func (this *ImageUnifierJPanel) UpdateImageWithFudgedImages() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "updateImageWithFudgedImages", e)
                }
            }()
            var bufferedImageArray []*BufferedImage = this.GetImageProcessorInput().GetBufferedImageArray(); _ = bufferedImageArray
            var tempBufferedImageArray []*BufferedImage = make([]*BufferedImage, len(bufferedImageArray)); _ = tempBufferedImageArray
            for index := int(0); index < len(bufferedImageArray); index++ {
                var indexedImageIcon *ImageIconUnique = any(this.ImageJList.GetModel().GetElementAt(index)).(*ImageIconUnique); _ = indexedImageIcon
                var bufferedImage *BufferedImage = ImageUtil_GetInstance().CreateBufferedImage(bufferedImageArray[indexedImageIcon.GetId()], this.ImageUnifierProperties.GetImageUnifierCell().GetWidth(), this.ImageUnifierProperties.GetImageUnifierCell().GetHeight(), false); _ = bufferedImage
                tempBufferedImageArray[index] = bufferedImage
            }
            var totalImages int = this.ImageUnifierProperties.GetRows() * this.ImageUnifierProperties.GetColumns(); _ = totalImages
            var averageRatio float64 = this.ImagesRatioUtil.GetAverage(tempBufferedImageArray, totalImages); _ = averageRatio
            var fudgedBufferedImageArray []*BufferedImage = this.ImagesRatioUtil.Fudge(tempBufferedImageArray, totalImages, averageRatio); _ = fudgedBufferedImageArray
            this.UpdateImage(fudgedBufferedImageArray)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageUnifierJPanel) UpdateImage_2(tempBufferedImageArray []*BufferedImage) {
        this.Result = ImageUnifierUtil_GetInstance().GetImage(tempBufferedImageArray, this.ImageUnifierProperties)
        var image *Icon = NewImageIcon(this.Result); _ = image
        var label *JLabel = NewJLabel(image); _ = label
        this.JScrollPane2.GetViewport().Add(label)
        this.UpdateUI()
}


func (this *ImageUnifierJPanel) UpdateOnPropertiesChange() {
        if this.RowsJTextField.GetText().Length() > 0 && this.ColumnsJTextField.GetText().Length() > 0 && this.CellWidthJTextField.GetText().Length() > 0 && this.CellHeightJTextField.GetText().Length() > 0 {
            this.ImageUnifierProperties = NewImageUnifierProperties()
            this.ImageUnifierProperties.SetRows(int(this.RowsJTextField.GetText()))
            this.ImageUnifierProperties.SetColumns(int(this.ColumnsJTextField.GetText()))
            var imageUnifierCell *ImageUnifierCell = NewImageUnifierCell(int(this.CellWidthJTextField.GetText()), int(this.CellHeightJTextField.GetText())); _ = imageUnifierCell
            this.ImageUnifierProperties.SetImageUnifierCell(imageUnifierCell)
            var cellRatio float64 = (float64(imageUnifierCell.GetWidth()) / imageUnifierCell.GetHeight()); _ = cellRatio
            var cellRatioString string = fmt.Sprint(float64(cellRatio)); _ = cellRatioString
            if len(cellRatioString) > 4 {
                cellRatioString = cellRatioString[0:4]
            }
            this.CellRatioJTextField.SetText(cellRatioString)
            this.UpdateImage()
        }
}


func (this *ImageUnifierJPanel) InitComponents() {
        var commonPhoneStrings *CommonPhoneStrings = CommonPhoneStrings_GetInstance(); _ = commonPhoneStrings
        this.JScrollPane1 = NewJScrollPane()
        this.ImageJList = NewJList()
        this.JScrollPane2 = NewJScrollPane()
        this.JPanel1 = NewJPanel()
        this.JButton1 = NewJButton()
        this.UpJButton = NewJButton()
        this.DownJButton = NewJButton()
        this.FudgeItJButton = NewJButton()
        this.JLabel2 = NewJLabel()
        this.ColumnsJTextField = NewJTextField()
        this.JLabel1 = NewJLabel()
        this.RowsJTextField = NewJTextField()
        this.JLabel3 = NewJLabel()
        this.CellWidthJTextField = NewJTextField()
        this.JLabel4 = NewJLabel()
        this.CellHeightJTextField = NewJTextField()
        this.JLabel5 = NewJLabel()
        this.CellRatioJTextField = NewJTextField()
        this.JLabel6 = NewJLabel()
        this.AvgRatioJTextField = NewJTextField()
        this.JButton2 = NewJButton()
        this.JButtonOrder = NewJButton()
        this.SetMinimumSize(NewDimension(500, 375))
        this.SetPreferredSize(NewDimension(0, 0))
        this.ImageJList.SetMinimumSize(NewDimension(50, 0))
        this.JScrollPane1.SetViewportView(this.ImageJList)
        this.JButton1.SetText("Save")
        this.JButton1.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) ActionPerformed(evt *ActionEvent) {
                this.JButton1ActionPerformed(evt)
        }
*/ NewActionListener())
        this.UpJButton.SetText("Up")
        this.UpJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) ActionPerformed(evt *ActionEvent) {
                this.UpJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.DownJButton.SetText("Down")
        this.DownJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) ActionPerformed(evt *ActionEvent) {
                this.DownJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.FudgeItJButton.SetText("Fudge It!")
        this.FudgeItJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) ActionPerformed(evt *ActionEvent) {
                this.FudgeItJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel2.SetText("Columns:")
        this.ColumnsJTextField.SetText(commonPhoneStrings.TWO)
        this.ColumnsJTextField.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) ActionPerformed(evt *ActionEvent) {
                this.ColumnsJTextFieldActionPerformed(evt)
        }
*/ NewActionListener())
        this.ColumnsJTextField.AddPropertyChangeListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) PropertyChange(evt *PropertyChangeEvent) {
                this.ColumnsJTextFieldPropertyChange(evt)
        }
*/ NewPropertyChangeListener())
        this.ColumnsJTextField.AddKeyListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) KeyPressed(evt IKeyEvent) {
                this.ColumnsJTextFieldKeyPressed(evt)
        }
            

        func (this *ImageUnifierJPanel) KeyTyped(evt IKeyEvent) {
                this.ColumnsJTextFieldKeyTyped(evt)
        }
*/ NewKeyAdapter())
        this.JLabel1.SetText("Rows:")
        this.RowsJTextField.SetText(commonPhoneStrings.TWO)
        this.RowsJTextField.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) ActionPerformed(evt *ActionEvent) {
                this.RowsJTextFieldActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel3.SetText("Cell Width:")
        this.CellWidthJTextField.SetText("200")
        this.CellWidthJTextField.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) ActionPerformed(evt *ActionEvent) {
                this.CellWidthJTextFieldActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel4.SetText("Cell Height:")
        this.CellHeightJTextField.SetText("150")
        this.CellHeightJTextField.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) ActionPerformed(evt *ActionEvent) {
                this.CellHeightJTextFieldActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel5.SetText("Cell Ratio:")
        this.CellRatioJTextField.SetEditable(false)
        this.JLabel6.SetText("Avg Image Ratio:")
        this.AvgRatioJTextField.SetEditable(false)
        this.JButton2.SetText(this.CommonStrings.UPDATE)
        this.JButton2.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) ActionPerformed(evt *ActionEvent) {
                this.JButton2ActionPerformed(evt)
        }
*/ NewActionListener())
        this.JButtonOrder.SetText("Order")
        this.JButtonOrder.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUnifierJPanel) ActionPerformed(evt *ActionEvent) {
                this.JButtonOrderActionPerformed(evt)
        }
*/ NewActionListener())
        var jPanel1Layout *GroupLayout = NewGroupLayout(this.JPanel1); _ = jPanel1Layout
        this.JPanel1.SetLayout(jPanel1Layout)
        jPanel1Layout.SetHorizontalGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, jPanel1Layout.CreateSequentialGroup().AddContainerGap().AddGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.TRAILING).AddComponent(this.JButton1, javax_Swing.GroupLayout.DEFAULT_SIZE, 151, Short_MAX_VALUE).AddGroup(jPanel1Layout.CreateSequentialGroup().AddComponent(this.UpJButton, javax_Swing.GroupLayout.PREFERRED_SIZE, 71, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 10, Short_MAX_VALUE).AddComponent(this.DownJButton, javax_Swing.GroupLayout.PREFERRED_SIZE, 70, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddComponent(this.FudgeItJButton, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 151, Short_MAX_VALUE).AddGroup(javax_Swing.GroupLayout.Alignment.LEADING, jPanel1Layout.CreateSequentialGroup().AddComponent(this.JLabel3).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 99, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddGroup(javax_Swing.GroupLayout.Alignment.LEADING, jPanel1Layout.CreateSequentialGroup().AddComponent(this.JLabel4).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 96, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddComponent(this.AvgRatioJTextField, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 151, Short_MAX_VALUE).AddComponent(this.JLabel6, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 151, Short_MAX_VALUE).AddComponent(this.CellRatioJTextField, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 151, Short_MAX_VALUE).AddComponent(this.JLabel5, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 151, Short_MAX_VALUE).AddComponent(this.CellHeightJTextField, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 151, Short_MAX_VALUE).AddComponent(this.CellWidthJTextField, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 151, Short_MAX_VALUE).AddComponent(this.JButton2, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 151, Short_MAX_VALUE).AddGroup(javax_Swing.GroupLayout.Alignment.LEADING, jPanel1Layout.CreateSequentialGroup().AddGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JLabel2).AddComponent(this.JLabel1)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.RowsJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 103, Short_MAX_VALUE).AddGroup(jPanel1Layout.CreateSequentialGroup().AddComponent(this.ColumnsJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 102, Short_MAX_VALUE).AddGap(1, 1, 1)))).AddComponent(this.JButtonOrder, javax_Swing.GroupLayout.DEFAULT_SIZE, 151, Short_MAX_VALUE)).AddContainerGap()))
        jPanel1Layout.SetVerticalGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(jPanel1Layout.CreateSequentialGroup().AddContainerGap().AddComponent(this.JButton1).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.UpJButton).AddComponent(this.DownJButton)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JButtonOrder).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.FudgeItJButton).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel2).AddComponent(this.ColumnsJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.RowsJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel3).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.CellWidthJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel4).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.CellHeightJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JButton2).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel5).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddComponent(this.CellRatioJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddGap(7, 7, 7).AddComponent(this.JLabel6).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.AvgRatioJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddGap(66, 66, 66)))
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JScrollPane1, javax_Swing.GroupLayout.PREFERRED_SIZE, 84, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JPanel1, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JScrollPane2, javax_Swing.GroupLayout.DEFAULT_SIZE, 527, Short_MAX_VALUE)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, layout.CreateSequentialGroup().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.TRAILING).AddComponent(this.JPanel1, javax_Swing.GroupLayout.Alignment.LEADING, 0, 445, Short_MAX_VALUE).AddComponent(this.JScrollPane2, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 445, Short_MAX_VALUE).AddComponent(this.JScrollPane1, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 445, Short_MAX_VALUE)).AddContainerGap()))
}


func (this *ImageUnifierJPanel) JButton2ActionPerformed(evt *ActionEvent) {
        this.UpdateOnPropertiesChange()
}


func (this *ImageUnifierJPanel) ColumnsJTextFieldKeyPressed(evt IKeyEvent) {
}


func (this *ImageUnifierJPanel) ColumnsJTextFieldKeyTyped(evt IKeyEvent) {
}


func (this *ImageUnifierJPanel) ColumnsJTextFieldPropertyChange(evt *PropertyChangeEvent) {
}


func (this *ImageUnifierJPanel) FudgeItJButtonActionPerformed(evt *ActionEvent) {
        this.UpdateImageWithFudgedImages()
}


func (this *ImageUnifierJPanel) CellHeightJTextFieldActionPerformed(evt *ActionEvent) {
        this.UpdateOnPropertiesChange()
}


func (this *ImageUnifierJPanel) CellWidthJTextFieldActionPerformed(evt *ActionEvent) {
        this.UpdateOnPropertiesChange()
}


func (this *ImageUnifierJPanel) ColumnsJTextFieldActionPerformed(evt *ActionEvent) {
        this.UpdateOnPropertiesChange()
}


func (this *ImageUnifierJPanel) RowsJTextFieldActionPerformed(evt *ActionEvent) {
        this.UpdateOnPropertiesChange()
}


func (this *ImageUnifierJPanel) JButton1ActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "jButton1ActionPerformed", e)
                }
            }()
            var fileArray []*File = this.ImageProcessorInput.GetFiles(); _ = fileArray
            var file *File = fileArray[0]; _ = file
            var filePath string = file.GetAbsolutePath(); _ = filePath
            var extensionIndex int = strings.Index(filePath, string(this.ImageStrings.PNG_EXTENSION)); _ = extensionIndex
            filePath = fmt.Sprint(NewStringMaker().Append(filePath[0:extensionIndex]).Append(CommonSeps_GetInstance().UNDERSCORE).Appendint(int(this.ImageUnifierProperties.GetColumns())).Append("_By_").Appendint(int(this.ImageUnifierProperties.GetRows())).Append("_Unified").Append(this.ImageStrings.PNG_EXTENSION))
            this.LogUtil.PutF("New File Path: " + filePath, this, StringUtil_GetInstance().EMPTY_STRING)
            var outputFile *File = NewFile(filePath); _ = outputFile
            ImagePersistanceUtil_GetInstance().SaveWithBatik(FileWrapperUtil_WrapFile(outputFile), this.Result)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageUnifierJPanel) DownJButtonActionPerformed(evt *ActionEvent) {
        var index int = this.ImageJList.GetSelectedIndex(); _ = index
        var defaultListModel *DefaultListModel = any(this.ImageJList.GetModel()).(*DefaultListModel); _ = defaultListModel
        if index >= 0 && index + 1 < defaultListModel.Size() {
            var indexedImageIcon *ImageIconUnique = any(defaultListModel.Remove(index)).(*ImageIconUnique); _ = indexedImageIcon
            defaultListModel.InsertElementAt(indexedImageIcon, index + 1)
            this.ImageJList.SetSelectedIndex(index + 1)
            this.UpdateImage()
        }
}


func (this *ImageUnifierJPanel) UpJButtonActionPerformed(evt *ActionEvent) {
        var index int = this.ImageJList.GetSelectedIndex(); _ = index
        if index > 0 {
            var defaultListModel *DefaultListModel = any(this.ImageJList.GetModel()).(*DefaultListModel); _ = defaultListModel
            var indexedImageIcon *ImageIconUnique = any(defaultListModel.Remove(index)).(*ImageIconUnique); _ = indexedImageIcon
            defaultListModel.InsertElementAt(indexedImageIcon, index - 1)
            this.ImageJList.SetSelectedIndex(index - 1)
            this.UpdateImage()
        }
}


func (this *ImageUnifierJPanel) JButtonOrderActionPerformed(evt *ActionEvent) {
        var defaultListModel *DefaultListModel = any(this.ImageJList.GetModel()).(*DefaultListModel); _ = defaultListModel
        var size int = defaultListModel.Size(); _ = size
        var last int = size - 1; _ = last
        for index := int(last); index >= 0; index-- {
            var indexedImageIcon *ImageIconUnique = any(defaultListModel.Remove(last)).(*ImageIconUnique); _ = indexedImageIcon
            defaultListModel.InsertElementAt(indexedImageIcon, last - index)
        }
        if size > 0 {
            this.ImageJList.SetSelectedIndex(0)
            this.UpdateImage()
        }
}


func (this *ImageUnifierJPanel) GetImageProcessorInput() *ImageProcessorInput {
        return this.ImageProcessorInput
}


func (this *ImageUnifierJPanel) SetImageProcessorInput(imageProcessorInput *ImageProcessorInput) {
        this.ImageProcessorInput = imageProcessorInput
        this.Init()
}

