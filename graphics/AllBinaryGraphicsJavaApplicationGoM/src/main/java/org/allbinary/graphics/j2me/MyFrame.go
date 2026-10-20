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

type MyFrame struct {
    JFrame
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    GuiLog *GuiLog
    XslHelper *XslHelper
    JAVA_TEMPLATE string
    SAVEEXTENSION string
    JAVAEXTENSION string
    FileJDialog *FileJDialog
    RotateJDialog *RotateJDialog
    AnimateJButton *JButton
    AutoDamageJMenuItem *JMenuItem
    AutoExplodeJMenuItem *JMenuItem
    AutoGrowJMenuItem *JMenuItem
    AutoRotateGenerationJMenuItem *JMenuItem
    AutoSheildMenuItem *JMenuItem
    AutoShrinkJMenuItem *JMenuItem
    CenterJMenuItem *JMenuItem
    CopyJMenuItem1 *JMenuItem
    ExitJMenuItem *JMenuItem
    ExplodeAllJMenuItem *JMenuItem
    ExportJMenuItem1 *JMenuItem
    FileJMenu *JMenu
    FrameJMenu *JMenu
    GenerateJMenu *JMenu
    InsertAfterJMenuItem *JMenuItem
    InsertBeforeJMenuItem *JMenuItem
    JButton1 *JButton
    JButton2 *JButton
    JButton3 *JButton
    JButton4 *JButton
    JButton41 *JButton
    JButton5 *JButton
    JMenu1 *JMenu
    JMenuItem1 *JMenuItem
    JPanel1 *JPanel
    JTabbedPane *JTabbedPane
    JToolBar1 *JToolBar
    MainJPanel *JPanel
    MenuJMenuBar *JMenuBar
    NewJMenuItem *JMenuItem
    NextFrameJMenuItem *JMenuItem
    ObjectJMenu *JMenu
    OpenJButton *JButton
    OpenJMenuItem *JMenuItem
    PlayJButton *JButton
    PreviousFrameJMenuItem *JMenuItem
    RemoveJMenuItem *JMenuItem
    SaveAsJMenuItem *JMenuItem
    SaveJButton *JButton
    SaveJMenuItem *JMenuItem
    StatusJPanel *JPanel
    StopJjButton *JButton
    ViewJMenu *JMenu
    ZoomInJMenuItem *JMenuItem
    ZoomOutJMenuItem *JMenuItem
}

// Implements FileJDialogListenerInterface, HashMapJDialogListenerInterface, KeyListener

var MyFrame_MAXWORKAREAS int = 10
var MyFrame_NumberOfWorkAreas int = 0
var MyFrame_PATH string = "G:/mnt/bc/mydev/work/graphics/AllBinaryGraphicsJavaApplication"
var MyFrame_SAVE_PATH string = fmt.Sprint(MyFrame_PATH) + "/save"
var MyFrame_TEMPLATE_EXPORT_PATH string = fmt.Sprint(MyFrame_PATH) + "/templates/export"
var MyFrame_StatusJLabel *JLabel

func (this *MyFrame) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.GuiLog = GuiLog_GetInstance()
    this.XslHelper = XslHelper_GetInstance()
    this.JAVA_TEMPLATE = "svgExport.xsl"
    this.SAVEEXTENSION = ".gxml"
    this.JAVAEXTENSION = ".javax"
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewMyFrame() *MyFrame {
    this := &MyFrame{}
    this.JFrame = *NewJFrame()
    this.initFields()
        this.InitComponents()
        this.SetSize(NewDimension(550, 400))
        this.MainJPanel.SetSize(NewDimension(550, 400))
        this.JTabbedPane.SetSize(NewDimension(550, 400))
        this.FileJDialog = NewFileJDialog()
        this.FileJDialog.AddFinishedListener(this)
        this.RotateJDialog = NewRotateJDialog(this, false)
    return this
}


func (this *MyFrame) InitComponents() {
        this.JToolBar1 = NewJToolBar()
        this.JButton2 = NewJButton()
        this.SaveJButton = NewJButton()
        this.JButton1 = NewJButton()
        this.OpenJButton = NewJButton()
        this.JButton3 = NewJButton()
        this.JButton4 = NewJButton()
        this.JButton41 = NewJButton()
        this.JButton5 = NewJButton()
        this.AnimateJButton = NewJButton()
        this.PlayJButton = NewJButton()
        this.StopJjButton = NewJButton()
        this.JPanel1 = NewJPanel()
        this.MainJPanel = NewJPanel()
        this.JTabbedPane = NewJTabbedPane()
        this.StatusJPanel = NewJPanel()
        MyFrame_StatusJLabel = NewJLabel()
        this.MenuJMenuBar = NewJMenuBar()
        this.FileJMenu = NewJMenu()
        this.NewJMenuItem = NewJMenuItem()
        this.OpenJMenuItem = NewJMenuItem()
        this.SaveJMenuItem = NewJMenuItem()
        this.SaveAsJMenuItem = NewJMenuItem()
        this.ExportJMenuItem1 = NewJMenuItem()
        this.ExitJMenuItem = NewJMenuItem()
        this.ViewJMenu = NewJMenu()
        this.ZoomInJMenuItem = NewJMenuItem()
        this.ZoomOutJMenuItem = NewJMenuItem()
        this.ObjectJMenu = NewJMenu()
        this.ExplodeAllJMenuItem = NewJMenuItem()
        this.FrameJMenu = NewJMenu()
        this.PreviousFrameJMenuItem = NewJMenuItem()
        this.NextFrameJMenuItem = NewJMenuItem()
        this.RemoveJMenuItem = NewJMenuItem()
        this.CopyJMenuItem1 = NewJMenuItem()
        this.InsertAfterJMenuItem = NewJMenuItem()
        this.InsertBeforeJMenuItem = NewJMenuItem()
        this.GenerateJMenu = NewJMenu()
        this.CenterJMenuItem = NewJMenuItem()
        this.AutoDamageJMenuItem = NewJMenuItem()
        this.AutoExplodeJMenuItem = NewJMenuItem()
        this.AutoGrowJMenuItem = NewJMenuItem()
        this.JMenuItem1 = NewJMenuItem()
        this.AutoRotateGenerationJMenuItem = NewJMenuItem()
        this.AutoSheildMenuItem = NewJMenuItem()
        this.AutoShrinkJMenuItem = NewJMenuItem()
        this.JMenu1 = NewJMenu()
        this.AddWindowListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) WindowClosing(evt *WindowEvent) {
                this.ExitForm(evt)
        }
*/ NewWindowAdapter())
        this.JButton2.SetText("New")
        this.JButton2.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.JButton2ActionPerformed(evt)
        }
*/ NewActionListener())
        this.JToolBar1.Add(this.JButton2)
        this.SaveJButton.SetText("Save")
        this.SaveJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.SaveJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JToolBar1.Add(this.SaveJButton)
        this.JButton1.SetText("Export")
        this.JButton1.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.JButton1ActionPerformed(evt)
        }
*/ NewActionListener())
        this.JToolBar1.Add(this.JButton1)
        this.OpenJButton.SetText("Open")
        this.OpenJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.OpenJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JToolBar1.Add(this.OpenJButton)
        this.JButton3.SetText("Copy")
        this.JToolBar1.Add(this.JButton3)
        this.JButton4.SetText("Insert Before")
        this.JToolBar1.Add(this.JButton4)
        this.JButton41.SetText("Insert After")
        this.JToolBar1.Add(this.JButton41)
        this.JButton5.SetText("Delete")
        this.JToolBar1.Add(this.JButton5)
        this.AnimateJButton.SetText("AnimateEdit")
        this.JToolBar1.Add(this.AnimateJButton)
        this.PlayJButton.SetText("Play")
        this.PlayJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) MousePressed(evt IMouseEvent) {
                this.PlayJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.JToolBar1.Add(this.PlayJButton)
        this.StopJjButton.SetText("Stop")
        this.StopJjButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) MousePressed(evt IMouseEvent) {
                this.StopJjButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.JToolBar1.Add(this.StopJjButton)
        this.GetContentPane().Add(this.JToolBar1, java_Awt.BorderLayout.NORTH)
        this.JPanel1.SetLayout(NewGridLayout(1, 0))
        this.MainJPanel.SetLayout(NewGridLayout(1, 0))
        this.JTabbedPane.SetMinimumSize(NewDimension(550, 400))
        this.JTabbedPane.AddChangeListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) StateChanged(evt *ChangeEvent) {
                this.JTabbedPaneStateChanged(evt)
        }
*/ NewChangeListener())
        this.MainJPanel.Add(this.JTabbedPane)
        this.JPanel1.Add(this.MainJPanel)
        this.GetContentPane().Add(this.JPanel1, java_Awt.BorderLayout.CENTER)
        this.StatusJPanel.SetLayout(NewGridLayout(1, 0))
        MyFrame_StatusJLabel.SetText("Status Bar")
        this.StatusJPanel.Add(MyFrame_StatusJLabel)
        this.GetContentPane().Add(this.StatusJPanel, java_Awt.BorderLayout.SOUTH)
        this.FileJMenu.SetText("File")
        this.FileJMenu.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.FileJMenuActionPerformed(evt)
        }
*/ NewActionListener())
        this.NewJMenuItem.SetText("New")
        this.NewJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.NewJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.FileJMenu.Add(this.NewJMenuItem)
        this.OpenJMenuItem.SetText("Open")
        this.OpenJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.OpenJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.FileJMenu.Add(this.OpenJMenuItem)
        this.SaveJMenuItem.SetText("Save")
        this.SaveJMenuItem.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) MousePressed(evt IMouseEvent) {
                this.SaveJMenuItemMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.FileJMenu.Add(this.SaveJMenuItem)
        this.SaveAsJMenuItem.SetText("SaveAs")
        this.FileJMenu.Add(this.SaveAsJMenuItem)
        this.ExportJMenuItem1.SetText("Export")
        this.ExportJMenuItem1.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.ExportJMenuItem1ActionPerformed(evt)
        }
*/ NewActionListener())
        this.FileJMenu.Add(this.ExportJMenuItem1)
        this.ExitJMenuItem.SetText("Exit")
        this.ExitJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.ExitJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.FileJMenu.Add(this.ExitJMenuItem)
        this.MenuJMenuBar.Add(this.FileJMenu)
        this.ViewJMenu.SetText("View")
        this.ZoomInJMenuItem.SetText("Zoom In")
        this.ZoomInJMenuItem.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) MousePressed(evt IMouseEvent) {
                this.ZoomInJMenuItemMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.ViewJMenu.Add(this.ZoomInJMenuItem)
        this.ZoomOutJMenuItem.SetText("Zoom Out")
        this.ZoomOutJMenuItem.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) MousePressed(evt IMouseEvent) {
                this.ZoomOutJMenuItemMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.ViewJMenu.Add(this.ZoomOutJMenuItem)
        this.MenuJMenuBar.Add(this.ViewJMenu)
        this.ObjectJMenu.SetText("Objects")
        this.ExplodeAllJMenuItem.SetText("Explode All")
        this.ExplodeAllJMenuItem.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) MousePressed(evt IMouseEvent) {
                this.ExplodeAllJMenuItemMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.ObjectJMenu.Add(this.ExplodeAllJMenuItem)
        this.MenuJMenuBar.Add(this.ObjectJMenu)
        this.FrameJMenu.SetText("Frame")
        this.FrameJMenu.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.FrameJMenuActionPerformed(evt)
        }
*/ NewActionListener())
        this.PreviousFrameJMenuItem.SetText("Previous")
        this.FrameJMenu.Add(this.PreviousFrameJMenuItem)
        this.NextFrameJMenuItem.SetText(CommonLabels_GetInstance().NEXT)
        this.FrameJMenu.Add(this.NextFrameJMenuItem)
        this.RemoveJMenuItem.SetText("Remove")
        this.FrameJMenu.Add(this.RemoveJMenuItem)
        this.CopyJMenuItem1.SetText("Copy")
        this.FrameJMenu.Add(this.CopyJMenuItem1)
        this.InsertAfterJMenuItem.SetText("Ins After")
        this.FrameJMenu.Add(this.InsertAfterJMenuItem)
        this.InsertBeforeJMenuItem.SetText("Ins Before")
        this.FrameJMenu.Add(this.InsertBeforeJMenuItem)
        this.MenuJMenuBar.Add(this.FrameJMenu)
        this.GenerateJMenu.SetText("Generate")
        this.CenterJMenuItem.SetText("Center")
        this.CenterJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.CenterJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.GenerateJMenu.Add(this.CenterJMenuItem)
        this.AutoDamageJMenuItem.SetText("Damage")
        this.GenerateJMenu.Add(this.AutoDamageJMenuItem)
        this.AutoExplodeJMenuItem.SetText("Explode")
        this.AutoExplodeJMenuItem.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) MousePressed(evt IMouseEvent) {
                this.AutoExplodeJMenuItemMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.GenerateJMenu.Add(this.AutoExplodeJMenuItem)
        this.AutoGrowJMenuItem.SetText("Grow")
        this.GenerateJMenu.Add(this.AutoGrowJMenuItem)
        this.JMenuItem1.SetText("Mirror")
        this.JMenuItem1.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) ActionPerformed(evt *ActionEvent) {
                this.JMenuItem1ActionPerformed(evt)
        }
*/ NewActionListener())
        this.GenerateJMenu.Add(this.JMenuItem1)
        this.AutoRotateGenerationJMenuItem.SetText("Rotate")
        this.AutoRotateGenerationJMenuItem.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MyFrame) MousePressed(evt IMouseEvent) {
                this.AutoRotateGenerationJMenuItemMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.GenerateJMenu.Add(this.AutoRotateGenerationJMenuItem)
        this.AutoSheildMenuItem.SetText("Shield")
        this.GenerateJMenu.Add(this.AutoSheildMenuItem)
        this.AutoShrinkJMenuItem.SetText("Shrink")
        this.GenerateJMenu.Add(this.AutoShrinkJMenuItem)
        this.MenuJMenuBar.Add(this.GenerateJMenu)
        this.JMenu1.SetText("Help")
        this.MenuJMenuBar.Add(this.JMenu1)
        this.SetJMenuBar(this.MenuJMenuBar)
        this.Pack()
}


func (this *MyFrame) ExplodeAllJMenuItemMousePressed(evt IMouseEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "explodeAllJMenuItemMousePressed", e)
                }
            }()
            MyCanvasEventService_Fire(NewMyCanvasEvent(any(NewMyCanvasEventSource(MyCanvasEventService_EXPLODE, nil)).(any)))
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MyFrame) AutoExplodeJMenuItemMousePressed(evt IMouseEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "autoExplodeJMenuItemMousePressed", e)
                }
            }()
            MyCanvasEventService_Fire(NewMyCanvasEvent(any(NewMyCanvasEventSource(MyCanvasEventService_AUTOEXPLODE, nil)).(any)))
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MyFrame) StopJjButtonMousePressed(evt IMouseEvent) {
        this.GetSelectedWorkArea().Stop()
}


func (this *MyFrame) PlayJButtonMousePressed(evt IMouseEvent) {
        this.GetSelectedWorkArea().Play()
}


func (this *MyFrame) AutoRotateGenerationJMenuItemMousePressed(evt IMouseEvent) {
        this.RotateJDialog.SetVisible(true)
}


func (this *MyFrame) ZoomOutJMenuItemMousePressed(evt IMouseEvent) {
        this.GetSelectedWorkArea().ChangeZoom(-1)
}


func (this *MyFrame) ZoomInJMenuItemMousePressed(evt IMouseEvent) {
        this.GetSelectedWorkArea().ChangeZoom(1)
}


func (this *MyFrame) OpenJMenuItemActionPerformed(evt *ActionEvent) {
        this.OpenGraphics()
}


func (this *MyFrame) ExportJMenuItem1ActionPerformed(evt *ActionEvent) {
        this.Export()
}


func (this *MyFrame) SaveJMenuItemMousePressed(evt IMouseEvent) {
        this.Save()
}


func (this *MyFrame) JTabbedPaneStateChanged(evt *ChangeEvent) {
        this.UpdateSelectedWorkArea()
}


func (this *MyFrame) FrameJMenuActionPerformed(evt *ActionEvent) {
}


func (this *MyFrame) ExitJMenuItemActionPerformed(evt *ActionEvent) {
        os.Exit(0)
}


func (this *MyFrame) NewJMenuItemActionPerformed(evt *ActionEvent) {
        this.NewGraphics()
}


func (this *MyFrame) FileJMenuActionPerformed(evt *ActionEvent) {
}


func (this *MyFrame) ExitForm(evt *WindowEvent) {
        os.Exit(0)
}


func (this *MyFrame) OpenJButtonActionPerformed(evt *ActionEvent) {
        this.OpenGraphics()
}


func (this *MyFrame) JButton2ActionPerformed(evt *ActionEvent) {
        this.NewGraphics()
}


func (this *MyFrame) CenterJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "centerJMenuItemActionPerformed", e)
                }
            }()
            MyCanvasEventService_Fire(NewMyCanvasEvent(any(NewMyCanvasEventSource(MyCanvasEventService_CENTER, nil)).(any)))
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MyFrame) SaveJButtonActionPerformed(evt *ActionEvent) {
        this.Save()
}


func (this *MyFrame) JButton1ActionPerformed(evt *ActionEvent) {
        this.Export()
}


func (this *MyFrame) JMenuItem1ActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "centerJMenuItemActionPerformed", e)
                }
            }()
            this.GetSelectedWorkArea().AutoMirror()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MyFrame) OnFiles(files []*File) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onFiles", e)
                }
            }()
            this.LogUtil.PutF("Reading: " + fmt.Sprint(files), this, "onFiles")
            for index := int(0); index < len(files); index++ {
                var document *Document = DomHelper_GetInstance().CreateDocument(files[index]); _ = document
                var workAreaJPanel *WorkAreaJPanel = NewWorkAreaJPanel(NewWorkAreaDom(document), this.GetSize()); _ = workAreaJPanel
                workAreaJPanel.SetLayout(NewGridLayout(1, 1))
                this.Add(any(workAreaJPanel).(WorkAreaJPanelInterface))
            }
            this.FileJDialog.SetVisible(false)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MyFrame) NewGraphics() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if JFrame_LOGGING.Contains(JFrame_LOGGING.MENUEVENT) {
                this.GuiLog.Put(this.CommonStrings.EXCEPTION, this, "newJMenuItemActionPerformed", e)
            }
                }
            }()
            var newName string = "graphic" + fmt.Sprint(MyFrame_NumberOfWorkAreas); _ = newName
            var dim *Dimension = this.GetSize(); _ = dim
            var newDialog *NewWorkAreaJDialog = NewNewWorkAreaJDialog(this, false, dim, newName); _ = newDialog
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MyFrame) OpenGraphics() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if JFrame_LOGGING.Contains(JFrame_LOGGING.MENUEVENT) {
                this.GuiLog.Put(this.CommonStrings.EXCEPTION, this, "open", e)
            }
                }
            }()
            this.LogUtil.PutF("Starting", this, "openImageJMenuItemActionPerformed")
            this.FileJDialog.SetVisible(true)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MyFrame) Export() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if JFrame_LOGGING.Contains(JFrame_LOGGING.MENUEVENT) {
                this.GuiLog.Put(this.CommonStrings.EXCEPTION, this, "export", e)
            }
                }
            }()
            var workAreaJPanel WorkAreaJPanelInterface = any(this.JTabbedPane.GetSelectedComponent()).(WorkAreaJPanelInterface); _ = workAreaJPanel
            var file *File = NewFile(MyFrame_SAVE_PATH + "/" + workAreaJPanel.GetName() + this.JAVAEXTENSION); _ = file
            var document *Document = workAreaJPanel.ToDocument(); _ = document
            this.LogUtil.PutF("Export: " + DomHelper_GetInstance().ToString(document), this, "exportJMenuItem1ActionPerformed")
            this.XslHelper.Export(file, fmt.Sprint(MyFrame_TEMPLATE_EXPORT_PATH) + "/" + this.JAVA_TEMPLATE, document)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MyFrame) Save() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if JFrame_LOGGING.Contains(JFrame_LOGGING.MENUEVENT) {
                this.GuiLog.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.SAVE, e)
            }
                }
            }()
            var workAreaJPanel WorkAreaJPanelInterface = any(this.JTabbedPane.GetSelectedComponent()).(WorkAreaJPanelInterface); _ = workAreaJPanel
            DomHelper_GetInstance().Save(NewFile(MyFrame_SAVE_PATH + "/" + workAreaJPanel.GetName() + this.SAVEEXTENSION), workAreaJPanel.ToDocument())
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *MyFrame) OnHashMap(hashMap *HashMap) {
        var angleIncrementInteger int = any(hashMap.Get(fmt.Sprint(0))).(int); _ = angleIncrementInteger
        var totalRotationInteger int = any(hashMap.Get(fmt.Sprint(1))).(int); _ = totalRotationInteger
        this.GetSelectedWorkArea().AutoRotate(int(angleIncrementInteger), int(totalRotationInteger))
}


func (this *MyFrame) GetSelectedWorkArea() WorkAreaJPanelInterface {
        return any(this.JTabbedPane.GetSelectedComponent()).(WorkAreaJPanelInterface)
}


func (this *MyFrame) UpdateSelectedWorkArea() {
        var numberOfTabs int = this.JTabbedPane.GetTabCount(); _ = numberOfTabs
        for index := int(0); index < numberOfTabs; index++ {
            var workAreaJPanel WorkAreaJPanelInterface = any(this.JTabbedPane.GetComponentAt(index)).(WorkAreaJPanelInterface); _ = workAreaJPanel
            workAreaJPanel.Deselect()
        }
        var workAreaJPanel WorkAreaJPanelInterface = any(this.JTabbedPane.GetSelectedComponent()).(WorkAreaJPanelInterface); _ = workAreaJPanel
        if workAreaJPanel != nil {
            workAreaJPanel.Select()
        }
}


// throws Exception
func (this *MyFrame) Add(newCanvasInterface WorkAreaJPanelInterface) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if JFrame_LOGGING.Contains(JFrame_LOGGING.GRAPHICSCREATION) {
                this.GuiLog.Put(this.CommonStrings.EXCEPTION, this, "addNewCanvas", e)
            }
                }
            }()
            MyFrame_NumberOfWorkAreas++
            if MyFrame_NumberOfWorkAreas == MyFrame_MAXWORKAREAS {
                panic(errors.New(fmt.Sprint()))
            }
            this.JTabbedPane.AddTab(newCanvasInterface.GetName(), any(newCanvasInterface).(*JPanel))
            this.JTabbedPane.SetSelectedComponent(any(newCanvasInterface).(*JPanel))
            var animationThread IThread = NewThread(any(newCanvasInterface).(Runnable)); _ = animationThread
            animationThread.Start()
            this.UpdateSelectedWorkArea()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MyFrame) KeyPressed(keyEvent IKeyEvent) {
        this.GetSelectedWorkArea().KeyPressed(keyEvent)
}


func (this *MyFrame) KeyReleased(keyEvent IKeyEvent) {
}


func (this *MyFrame) KeyTyped(keyEvent IKeyEvent) {
}


func MyFrame_Main(args []string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            logUtil.Put(CommonStrings_GetInstance().EXCEPTION, nil, "main", e)
                }
            }()
            AngleFactory_GetInstance()
            SmallIntegerSingletonFactory_GetInstance().Init()
            NewMyFrame().Show()
            return
        }(); tryControl1 == 1 {
            return
        }
}

