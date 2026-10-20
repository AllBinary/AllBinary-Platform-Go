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

type ImageToolJFrame struct {
    JFrame
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    FileJDialog *FileJDialog
    ImageProcessorInput *ImageProcessorInput
    AnalysisActionJDialog *JDialog
    AnalyzeJMenuItem *JMenuItem
    BlueJLabel *JLabel
    ColorRangeJLabel *JLabel
    CreateStripImageJMenuItem *JMenuItem
    FileJMenu *JMenu
    GenerateSpriteSplitterMenuItem *JMenuItem
    GenerateRotationMenuItem *JMenuItem
    GenerateMirrorSpriteMenuItem *JMenuItem
    GenerateRotationSpriteJMenuItem *JMenuItem
    GreenJLabel *JLabel
    JPanel1 *JPanel
    MainJMenuBar *JMenuBar
    MaxBlueJTextField *JTextField
    MaxGreenJTextField *JTextField
    MaxJLabel *JLabel
    MaxRedJTextField *JTextField
    MinBlueJTextField *JTextField
    MinGreenJTextField *JTextField
    MinJLabel *JLabel
    MinRedJTextField *JTextField
    MirrorJMenuItem *JMenuItem
    OkJButton *JButton
    OpenImageJMenuItem *JMenuItem
    ProcessingJMenu *JMenu
    RedJLabel *JLabel
    ScaleImagesJMenuItem *JMenuItem
    CanvasImagesJMenuItem *JMenuItem
    CreateIconImagesJMenuItem *JMenuItem
    TitleJLabel *JLabel
}

// Implements FileJDialogListenerInterface

var ImageToolJFrame_ImageProcessorInputCompositeInterface ImageProcessorInputCompositeInterface

func (this *ImageToolJFrame) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewImageToolJFrame() *ImageToolJFrame {
    this := &ImageToolJFrame{}
    this.JFrame = *NewJFrame()
    this.initFields()
        this.InitComponents()
        var commonPhoneStrings *CommonPhoneStrings = CommonPhoneStrings_GetInstance(); _ = commonPhoneStrings
        this.SetSize(640, 480)
        this.GetMaxRedJTextField().SetText("255")
        this.GetMaxGreenJTextField().SetText("255")
        this.GetMaxBlueJTextField().SetText("255")
        this.GetMinRedJTextField().SetText(commonPhoneStrings.ZERO)
        this.GetMinGreenJTextField().SetText(commonPhoneStrings.ZERO)
        this.GetMinBlueJTextField().SetText(commonPhoneStrings.ZERO)
        this.FileJDialog = NewFileJDialog()
        this.FileJDialog.AddFinishedListener(this)
    return this
}


func (this *ImageToolJFrame) InitComponents() {
        this.AnalysisActionJDialog = NewJDialog()
        this.TitleJLabel = NewJLabel()
        this.ColorRangeJLabel = NewJLabel()
        this.RedJLabel = NewJLabel()
        this.MinJLabel = NewJLabel()
        this.MaxJLabel = NewJLabel()
        this.GreenJLabel = NewJLabel()
        this.BlueJLabel = NewJLabel()
        this.MinRedJTextField = NewJTextField()
        this.MaxRedJTextField = NewJTextField()
        this.MaxGreenJTextField = NewJTextField()
        this.MinBlueJTextField = NewJTextField()
        this.MaxBlueJTextField = NewJTextField()
        this.OkJButton = NewJButton()
        this.MinGreenJTextField = NewJTextField()
        this.JPanel1 = NewJPanel()
        this.MainJMenuBar = NewJMenuBar()
        this.FileJMenu = NewJMenu()
        this.OpenImageJMenuItem = NewJMenuItem()
        this.ProcessingJMenu = NewJMenu()
        this.AnalyzeJMenuItem = NewJMenuItem()
        this.GenerateSpriteSplitterMenuItem = NewJMenuItem()
        this.GenerateRotationMenuItem = NewJMenuItem()
        this.GenerateMirrorSpriteMenuItem = NewJMenuItem()
        this.GenerateRotationSpriteJMenuItem = NewJMenuItem()
        this.CreateStripImageJMenuItem = NewJMenuItem()
        this.MirrorJMenuItem = NewJMenuItem()
        this.ScaleImagesJMenuItem = NewJMenuItem()
        this.CanvasImagesJMenuItem = NewJMenuItem()
        this.CreateIconImagesJMenuItem = NewJMenuItem()
        this.AnalysisActionJDialog.SetMinimumSize(NewDimension(350, 250))
        this.TitleJLabel.SetText("              Color At Action Options            ")
        this.ColorRangeJLabel.SetText("Color Range")
        this.RedJLabel.SetText("Red:")
        this.MinJLabel.SetText("Minimum")
        this.MaxJLabel.SetText("Maximum")
        this.GreenJLabel.SetText("Green:")
        this.BlueJLabel.SetText("Blue:")
        this.OkJButton.SetText("OK")
        this.OkJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.OkJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        var analysisActionJDialogLayout *GroupLayout = NewGroupLayout(this.AnalysisActionJDialog.GetContentPane()); _ = analysisActionJDialogLayout
        this.AnalysisActionJDialog.GetContentPane().SetLayout(analysisActionJDialogLayout)
        analysisActionJDialogLayout.SetHorizontalGroup(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(analysisActionJDialogLayout.CreateSequentialGroup().AddContainerGap().Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(analysisActionJDialogLayout.CreateSequentialGroup().Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(this.BlueJLabel).Add(this.GreenJLabel).Add(this.RedJLabel).Add(this.ColorRangeJLabel)).Add(29, 29, 29).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(org_Jdesktop.Layout.GroupLayout.TRAILING, this.MinJLabel, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, 58, Short_MAX_VALUE).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.TRAILING, false).Add(org_Jdesktop.Layout.GroupLayout.LEADING, this.MinBlueJTextField).Add(org_Jdesktop.Layout.GroupLayout.LEADING, this.MinGreenJTextField).Add(org_Jdesktop.Layout.GroupLayout.LEADING, this.MinRedJTextField, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, 42, Short_MAX_VALUE).Add(org_Jdesktop.Layout.GroupLayout.LEADING, this.OkJButton))).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.TRAILING).Add(org_Jdesktop.Layout.GroupLayout.LEADING, analysisActionJDialogLayout.CreateSequentialGroup().Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.TRAILING, false).Add(org_Jdesktop.Layout.GroupLayout.LEADING, this.MaxRedJTextField).Add(org_Jdesktop.Layout.GroupLayout.LEADING, analysisActionJDialogLayout.CreateSequentialGroup().Add(1, 1, 1).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(this.MaxBlueJTextField).Add(this.MaxGreenJTextField, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, 43, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE)))).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED, 17, Short_MAX_VALUE)).Add(org_Jdesktop.Layout.GroupLayout.LEADING, this.MaxJLabel, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, 61, Short_MAX_VALUE)).AddContainerGap()).Add(org_Jdesktop.Layout.GroupLayout.TRAILING, analysisActionJDialogLayout.CreateSequentialGroup().Add(this.TitleJLabel).Add(25, 25, 25)))))
        analysisActionJDialogLayout.SetVerticalGroup(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(analysisActionJDialogLayout.CreateSequentialGroup().Add(this.TitleJLabel).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.BASELINE).Add(this.MinJLabel).Add(this.ColorRangeJLabel).Add(this.MaxJLabel)).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.TRAILING).Add(this.RedJLabel).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.BASELINE).Add(this.MinRedJTextField, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).Add(this.MaxRedJTextField, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE))).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.TRAILING).Add(this.GreenJLabel).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.BASELINE).Add(this.MinGreenJTextField, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).Add(this.MaxGreenJTextField, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE))).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.TRAILING).Add(this.BlueJLabel).Add(analysisActionJDialogLayout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.BASELINE).Add(this.MinBlueJTextField, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).Add(this.MaxBlueJTextField, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE))).Add(16, 16, 16).Add(this.OkJButton).AddContainerGap(org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
        this.SetDefaultCloseOperation(javax_Swing.WindowConstants.EXIT_ON_CLOSE)
        this.GetContentPane().SetLayout(NewGridLayout(1, 0))
        this.JPanel1.SetLayout(NewGridLayout(1, 0))
        this.GetContentPane().Add(this.JPanel1)
        this.FileJMenu.SetText("File")
        this.FileJMenu.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.FileJMenuActionPerformed(evt)
        }
*/ NewActionListener())
        this.OpenImageJMenuItem.SetText("Open")
        this.OpenImageJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.OpenImageJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.FileJMenu.Add(this.OpenImageJMenuItem)
        this.MainJMenuBar.Add(this.FileJMenu)
        this.ProcessingJMenu.SetText("Processing")
        this.ProcessingJMenu.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.ProcessingJMenuActionPerformed(evt)
        }
*/ NewActionListener())
        this.AnalyzeJMenuItem.SetText("Image Analyze")
        this.AnalyzeJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.AnalyzeJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.AnalyzeJMenuItem)
        this.GenerateMirrorSpriteMenuItem.SetText("Mirror Sprite Generator")
        this.GenerateMirrorSpriteMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.GenerateMirrorSpriteMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.GenerateMirrorSpriteMenuItem)
        this.GenerateRotationSpriteJMenuItem.SetText("Rotation Sprite Generator")
        this.GenerateRotationSpriteJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.GenerateRotationSpriteJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.GenerateRotationSpriteJMenuItem)
        this.GenerateSpriteSplitterMenuItem.SetText("Sprite Splitter Generator")
        this.GenerateSpriteSplitterMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.GenerateSpriteSplitterMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.GenerateSpriteSplitterMenuItem)
        this.GenerateRotationMenuItem.SetText("Rotate Image Generator")
        this.GenerateRotationMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.GenerateRotationMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.GenerateRotationMenuItem)
        this.CreateStripImageJMenuItem.SetText("Multi Image Unifier")
        this.CreateStripImageJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.CreateStripImageJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.CreateStripImageJMenuItem)
        this.MirrorJMenuItem.SetText("Mirror Image(s)")
        this.MirrorJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.MirrorJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.MirrorJMenuItem)
        this.ScaleImagesJMenuItem.SetText("Scale Image(s)")
        this.ScaleImagesJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.ScaleImagesJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.ScaleImagesJMenuItem)
        this.CanvasImagesJMenuItem.SetText("Canvas Image(s)")
        this.CanvasImagesJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.CanvasImagesJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.CanvasImagesJMenuItem)
        this.CreateIconImagesJMenuItem.SetText("Create Icon Image(s)")
        this.CreateIconImagesJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageToolJFrame) ActionPerformed(evt *ActionEvent) {
                this.CreateIconsJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.CreateIconImagesJMenuItem)
        this.MainJMenuBar.Add(this.ProcessingJMenu)
        this.SetJMenuBar(this.MainJMenuBar)
        this.Pack()
}


func (this *ImageToolJFrame) CreateStripImageJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "createStripImageJMenuItemActionPerformed", e)
                }
            }()
            var imageUnifierJPanel *ImageUnifierJPanel = NewImageUnifierJPanel(this.ImageProcessorInput); _ = imageUnifierJPanel
            this.JPanel1.RemoveAll()
            this.JPanel1.Add(imageUnifierJPanel)
            this.JPanel1.UpdateUI()
            ImageToolJFrame_ImageProcessorInputCompositeInterface = imageUnifierJPanel
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) OkJButtonActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "okJButtonActionPerformed", e)
                }
            }()
            this.LogUtil.PutF("Starting", this, "okJButtonActionPerformed")
            var imageAnalysisResultsArray []*ImageAnalysisResults = ImageAnalysis_GetInstance().Process(this.ImageProcessorInput.GetBufferedImageArray(), this.GetColorRangeInterface()); _ = imageAnalysisResultsArray
            this.JPanel1.RemoveAll()
            this.JPanel1.Add(NewImageAnalysisResultsJPanel(imageAnalysisResultsArray))
            this.JPanel1.UpdateUI()
            this.AnalysisActionJDialog.SetVisible(false)
            ImageToolJFrame_ImageProcessorInputCompositeInterface = nil
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) ProcessingJMenuActionPerformed(evt *ActionEvent) {
}


func (this *ImageToolJFrame) OpenImageJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "openImageJMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF("Starting", this, "openImageJMenuItemActionPerformed")
            this.FileJDialog.SetVisible(true)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) OnFiles(files []*File) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onFiles", e)
                }
            }()
            this.LogUtil.PutF("Reading " + fmt.Sprint(len(files)) + " files.", this, "onFiles")
            Arrays_Sort(files, /* Java anonymous class body: implement it as a named type
                

            func (this *ImageToolJFrame) Compare(file1 *File, file2 *File) int {
                    var fileNameIndexValue1 int = this.ExtractIndexNumberIfAnyFromFileName(file1.GetName()); _ = fileNameIndexValue1
                    var fileNameIndexValue2 int = this.ExtractIndexNumberIfAnyFromFileName(file2.GetName()); _ = fileNameIndexValue2
                    return fileNameIndexValue1 - fileNameIndexValue2
            }
                

            func (this *ImageToolJFrame) ExtractIndexNumberIfAnyFromFileName(name string) int {
                    var i int = 0; _ = i
                    if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
                        defer func() {
                            if recovered := recover(); recovered != nil {
                                e := Java_ToError(recovered)
                                _ = e
                        i = 0
                            }
                        }()
                        var start int = strings.LastIndex(name, string('_')) + 1; _ = start
                        var end int = strings.LastIndex(name, string('.')); _ = end
                        var number string = name[start:end]; _ = number
                        i = func() int { value, _ := strconv.Atoi(number); return value }()
                        return
                    }(); tryControl1 == 1 {
                        return tryResult1
                    }
                    return i
            }
*/ NewComparator())
            var bufferedImageArray []*BufferedImage = make([]*BufferedImage, len(files)); _ = bufferedImageArray
            for index := int(0); index < len(files); index++ {
                bufferedImageArray[index] = ImageIO_Read(files[index])
                if bufferedImageArray[index] == nil {
                    panic(errors.New(fmt.Sprint("Unable to read file: " + fmt.Sprint(files[index].GetPath()))))
                }
            }
            this.ImageProcessorInput = NewImageProcessorInput(files, bufferedImageArray)
            this.FileJDialog.SetVisible(false)
            if ImageToolJFrame_ImageProcessorInputCompositeInterface != nil {
                ImageToolJFrame_ImageProcessorInputCompositeInterface.SetImageProcessorInput(this.ImageProcessorInput)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) FileJMenuActionPerformed(evt *ActionEvent) {
}


func (this *ImageToolJFrame) AnalyzeJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "analyzeJMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF("Starting", this, "analyzeJMenuItemActionPerformed")
            this.AnalysisActionJDialog.SetVisible(true)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) GenerateRotationSpriteJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "generateRotationSpriteJMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF("Starting", this, "generateRotationSpriteJMenuItemActionPerformed")
            var spriteImageJPanel *RotationSpriteImageJPanel = NewRotationSpriteImageJPanel(this.ImageProcessorInput); _ = spriteImageJPanel
            this.JPanel1.RemoveAll()
            this.JPanel1.Add(spriteImageJPanel)
            this.JPanel1.UpdateUI()
            ImageToolJFrame_ImageProcessorInputCompositeInterface = spriteImageJPanel
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) GenerateSpriteSplitterMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "generateRotationMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF("Starting", this, "generateRotationSpriteJMenuItemActionPerformed")
            var imageJPanel *SpriteSplitterImageJPanel = NewSpriteSplitterImageJPanel(this.ImageProcessorInput); _ = imageJPanel
            this.JPanel1.RemoveAll()
            this.JPanel1.Add(imageJPanel)
            this.JPanel1.UpdateUI()
            ImageToolJFrame_ImageProcessorInputCompositeInterface = imageJPanel
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) GenerateRotationMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "generateRotationMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF("Starting", this, "generateRotationSpriteJMenuItemActionPerformed")
            var rotationImageJPanel *RotationImageJPanel = NewRotationImageJPanel(this.ImageProcessorInput); _ = rotationImageJPanel
            this.JPanel1.RemoveAll()
            this.JPanel1.Add(rotationImageJPanel)
            this.JPanel1.UpdateUI()
            ImageToolJFrame_ImageProcessorInputCompositeInterface = rotationImageJPanel
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) ScaleImagesJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "scaleImagesJMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "scaleImagesJMenuItemActionPerformed")
            var resizeImageJPanel *ResizeImageJPanel = NewResizeImageJPanel(this.ImageProcessorInput); _ = resizeImageJPanel
            this.JPanel1.RemoveAll()
            this.JPanel1.Add(resizeImageJPanel)
            this.JPanel1.UpdateUI()
            ImageToolJFrame_ImageProcessorInputCompositeInterface = resizeImageJPanel
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) CreateIconsJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "scaleImagesJMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "createIconsJMenuItemActionPerformed")
            var createIconImagesJPanel *CreateIconImagesJPanel = NewCreateIconImagesJPanel(this.ImageProcessorInput); _ = createIconImagesJPanel
            this.JPanel1.RemoveAll()
            this.JPanel1.Add(createIconImagesJPanel)
            this.JPanel1.UpdateUI()
            ImageToolJFrame_ImageProcessorInputCompositeInterface = createIconImagesJPanel
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) CanvasImagesJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "scaleImagesJMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "canvasImagesJMenuItemActionPerformed")
            var resizeImageJPanel *CanvasImageJPanel = NewCanvasImageJPanel(this.ImageProcessorInput); _ = resizeImageJPanel
            this.JPanel1.RemoveAll()
            this.JPanel1.Add(resizeImageJPanel)
            this.JPanel1.UpdateUI()
            ImageToolJFrame_ImageProcessorInputCompositeInterface = resizeImageJPanel
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) GenerateMirrorSpriteMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "generateMirrorSpriteJMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF("Starting", this, "generateMirrorSpriteJMenuItemActionPerformed")
            var spriteImageJPanel *MirrorSpriteImageJPanel = NewMirrorSpriteImageJPanel(this.ImageProcessorInput); _ = spriteImageJPanel
            this.JPanel1.RemoveAll()
            this.JPanel1.Add(spriteImageJPanel)
            this.JPanel1.UpdateUI()
            ImageToolJFrame_ImageProcessorInputCompositeInterface = spriteImageJPanel
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) MirrorJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "generateMirrorJMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF("Starting", this, "generateMirrorJMenuItemActionPerformed")
            var mirrorImageJPanel *MirrorImageJPanel = NewMirrorImageJPanel(this.ImageProcessorInput); _ = mirrorImageJPanel
            this.JPanel1.RemoveAll()
            this.JPanel1.Add(mirrorImageJPanel)
            this.JPanel1.UpdateUI()
            ImageToolJFrame_ImageProcessorInputCompositeInterface = mirrorImageJPanel
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImageToolJFrame) GetColorRangeInterface() ColorRangeInterface {
        var colorRangeInterface ColorRangeInterface = any(NewColorRange()).(ColorRangeInterface); _ = colorRangeInterface
        colorRangeInterface.SetMinRed(int(int(this.GetMinRedJTextField().GetText())))
        colorRangeInterface.SetMaxRed(int(int(this.GetMaxRedJTextField().GetText())))
        colorRangeInterface.SetMinGreen(int(int(this.GetMinGreenJTextField().GetText())))
        colorRangeInterface.SetMaxGreen(int(int(this.GetMaxGreenJTextField().GetText())))
        colorRangeInterface.SetMinBlue(int(int(this.GetMinBlueJTextField().GetText())))
        colorRangeInterface.SetMaxBlue(int(int(this.GetMaxBlueJTextField().GetText())))
        return colorRangeInterface
}


func (this *ImageToolJFrame) GetMaxBlueJTextField() *JTextField {
        return this.MaxBlueJTextField
}


func (this *ImageToolJFrame) SetMaxBlueJTextField(maxBlueJTextField *JTextField) {
        this.MaxBlueJTextField = this.MaxBlueJTextField
}


func (this *ImageToolJFrame) GetMaxGreenJTextField() *JTextField {
        return this.MaxGreenJTextField
}


func (this *ImageToolJFrame) SetMaxGreenJTextField(maxGreenJTextField *JTextField) {
        this.MaxGreenJTextField = this.MaxGreenJTextField
}


func (this *ImageToolJFrame) GetMaxRedJTextField() *JTextField {
        return this.MaxRedJTextField
}


func (this *ImageToolJFrame) SetMaxRedJTextField(maxRedJTextField *JTextField) {
        this.MaxRedJTextField = this.MaxRedJTextField
}


func (this *ImageToolJFrame) GetMinBlueJTextField() *JTextField {
        return this.MinBlueJTextField
}


func (this *ImageToolJFrame) SetMinBlueJTextField(minBlueJTextField *JTextField) {
        this.MinBlueJTextField = this.MinBlueJTextField
}


func (this *ImageToolJFrame) GetMinGreenJTextField() *JTextField {
        return this.MinGreenJTextField
}


func (this *ImageToolJFrame) SetMinGreenJTextField(minGreenJTextField *JTextField) {
        this.MinGreenJTextField = this.MinGreenJTextField
}


func (this *ImageToolJFrame) GetMinRedJTextField() *JTextField {
        return this.MinRedJTextField
}


func (this *ImageToolJFrame) SetMinRedJTextField(minRedJTextField *JTextField) {
        this.MinRedJTextField = this.MinRedJTextField
}

