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

type InputAutomationJFrame struct {
    JFrame
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    InputAutomationModuleConfigurations *InputAutomationModuleConfigurations
    InputAutomationModuleFactory *InputAutomationModuleFactoryFactory
    InputAutomationModuleInterface IInputAutomationModuleFactoryInterface
    RunnableInterface RunnableInterface
    Thread IThread
    HelpSet *HelpSet
    Uri *URI
    AboutJMenuItem *JMenuItem
    AutomationModuleConfigurationJPanel *JPanel
    GameRobotJTabbedPane *JTabbedPane
    HelpJMenuItem *JMenuItem
    InputAutomationModuleJList *JList
    InputAutomationModuleJPanel *JPanel
    JScrollPane1 *JScrollPane
    MainJMenuBar *JMenuBar
    ModulesJMenuItem *JMenuItem
    NoModuleSelectedJDialog *JDialog
    NoModuleSelectedJLabel *JLabel
    OptionsJMenu *JMenu
    ProcessingJMenu *JMenu
    StartJMenuItem *JMenuItem
    StopJMenuItem *JMenuItem
    StopOnFocusJCheckBoxMenuItem *JCheckBoxMenuItem
    SubscriptionJMenuItem *JMenuItem
    ToolsJMenu *JMenu
    UpdatesJMenuItem *JMenuItem
}

// Implements InputAutomationConfigurationModuleChangeListener, InputAutomationRobotChangeListener, *HelpSetListener

var InputAutomationJFrame_INPUTAUTOMATION_JFRAME *InputAutomationJFrame

func (this *InputAutomationJFrame) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.InputAutomationModuleConfigurations = NewInputAutomationModuleConfigurations()
    
    
    
    
    
    this.Uri = NewURI("http://geocities.com/allbinary/")
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewInputAutomationJFrame() *InputAutomationJFrame {
    this := &InputAutomationJFrame{}
    this.JFrame = *NewJFrame()
    this.initFields()
        this.InitComponents()
        var url *URL = reflect.TypeOf(this).GetResource("/help/Help.hs"); _ = url
        this.LogUtil.PutF("URL: " + fmt.Sprint(url), this, this.CommonStrings.CONSTRUCTOR)
        this.HelpSet = JavaHelpUtil_GetInstance().GetHelpSet(url)
        url = reflect.TypeOf(this).GetResource("/resources/allbinaryicon8bit.jpg")
        var imageIcon *ImageIcon = NewImageIcon(url); _ = imageIcon
        this.SetIconImage(imageIcon.GetImage())
        this.Init()
    return this
}


func (this *InputAutomationJFrame) HelpSetAdded(helpSetEvent *HelpSetEvent) {
        this.HelpSet.Add(helpSetEvent.GetHelpSet())
}


func (this *InputAutomationJFrame) HelpSetRemoved(helpSetEvent *HelpSetEvent) {
        this.HelpSet.Remove(helpSetEvent.GetHelpSet())
}


// throws Exception
func (this *InputAutomationJFrame) Init() {
        var inputAutomationConfigurationFactory *InputAutomationConfigurationFactory = InputAutomationConfigurationFactory_GetInstance(); _ = inputAutomationConfigurationFactory
        inputAutomationConfigurationFactory.Init(InputAutomationClientInformationFactory_GetInstance())
        var inputAutomationModuleConfigurationList IList[*InputAutomationModuleConfiguration] = inputAutomationConfigurationFactory.InputAutomationConfiguration.GetInputAutomationModuleConfigurationList(); _ = inputAutomationModuleConfigurationList
        if inputAutomationModuleConfigurationList == nil {
        } else {
            this.InputAutomationModuleConfigurations = NewInputAutomationModuleConfigurations(inputAutomationModuleConfigurationList)
        }
        this.InputAutomationModuleFactory = NewInputAutomationModuleFactoryFactory(this.InputAutomationModuleConfigurations, this)
        (NewJListSwingWorker(this.InputAutomationModuleJList, this.InputAutomationModuleFactory.GetListModel())).Execute()
        this.GameRobotJTabbedPane.SetEnabledAt(1, false)
        this.GameRobotJTabbedPane.SetSelectedIndex(0)
}


func (this *InputAutomationJFrame) InitComponents() {
        this.NoModuleSelectedJDialog = NewJDialog()
        this.NoModuleSelectedJLabel = NewJLabel()
        this.GameRobotJTabbedPane = NewJTabbedPane()
        this.InputAutomationModuleJPanel = NewJPanel()
        this.JScrollPane1 = NewJScrollPane()
        this.InputAutomationModuleJList = NewJList()
        this.AutomationModuleConfigurationJPanel = NewJPanel()
        this.MainJMenuBar = NewJMenuBar()
        this.ProcessingJMenu = NewJMenu()
        this.StartJMenuItem = NewJMenuItem()
        this.StopJMenuItem = NewJMenuItem()
        this.OptionsJMenu = NewJMenu()
        this.StopOnFocusJCheckBoxMenuItem = NewJCheckBoxMenuItem()
        this.ToolsJMenu = NewJMenu()
        this.HelpJMenuItem = NewJMenuItem()
        this.UpdatesJMenuItem = NewJMenuItem()
        this.SubscriptionJMenuItem = NewJMenuItem()
        this.ModulesJMenuItem = NewJMenuItem()
        this.AboutJMenuItem = NewJMenuItem()
        this.NoModuleSelectedJDialog.SetMinimumSize(NewDimension(200, 100))
        this.NoModuleSelectedJLabel.SetText("Please Select A Module")
        var noModuleSelectedJDialogLayout *GroupLayout = NewGroupLayout(this.NoModuleSelectedJDialog.GetContentPane()); _ = noModuleSelectedJDialogLayout
        this.NoModuleSelectedJDialog.GetContentPane().SetLayout(noModuleSelectedJDialogLayout)
        noModuleSelectedJDialogLayout.SetHorizontalGroup(noModuleSelectedJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(noModuleSelectedJDialogLayout.CreateSequentialGroup().AddContainerGap().AddComponent(this.NoModuleSelectedJLabel).AddContainerGap(javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
        noModuleSelectedJDialogLayout.SetVerticalGroup(noModuleSelectedJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.NoModuleSelectedJLabel, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.DEFAULT_SIZE, 48, Short_MAX_VALUE))
        this.SetDefaultCloseOperation(javax_Swing.WindowConstants.EXIT_ON_CLOSE)
        this.SetMinimumSize(NewDimension(325, 215))
        this.SetResizable(false)
        this.AddWindowFocusListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) WindowGainedFocus(evt *WindowEvent) {
                this.FormWindowGainedFocus(evt)
        }
            

        func (this *InputAutomationJFrame) WindowLostFocus(evt *WindowEvent) {
                this.FormWindowLostFocus(evt)
        }
*/ NewWindowFocusListener())
        this.AddFocusListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) FocusGained(evt *FocusEvent) {
                this.FormFocusGained(evt)
        }
            

        func (this *InputAutomationJFrame) FocusLost(evt *FocusEvent) {
                this.FormFocusLost(evt)
        }
*/ NewFocusAdapter())
        this.GameRobotJTabbedPane.SetMinimumSize(NewDimension(320, 200))
        this.InputAutomationModuleJList.SetMinimumSize(NewDimension(200, 0))
        this.InputAutomationModuleJList.AddListSelectionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) ValueChanged(evt *ListSelectionEvent) {
                this.InputAutomationModuleJListValueChanged(evt)
        }
*/ NewListSelectionListener())
        this.JScrollPane1.SetViewportView(this.InputAutomationModuleJList)
        var inputAutomationModuleJPanelLayout *GroupLayout = NewGroupLayout(this.InputAutomationModuleJPanel); _ = inputAutomationModuleJPanelLayout
        this.InputAutomationModuleJPanel.SetLayout(inputAutomationModuleJPanelLayout)
        inputAutomationModuleJPanelLayout.SetHorizontalGroup(inputAutomationModuleJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(inputAutomationModuleJPanelLayout.CreateSequentialGroup().AddComponent(this.JScrollPane1, javax_Swing.GroupLayout.PREFERRED_SIZE, 199, javax_Swing.GroupLayout.PREFERRED_SIZE).AddContainerGap(367, Short_MAX_VALUE)))
        inputAutomationModuleJPanelLayout.SetVerticalGroup(inputAutomationModuleJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(inputAutomationModuleJPanelLayout.CreateSequentialGroup().AddComponent(this.JScrollPane1, javax_Swing.GroupLayout.PREFERRED_SIZE, 330, javax_Swing.GroupLayout.PREFERRED_SIZE).AddContainerGap(javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
        this.GameRobotJTabbedPane.AddTab("Modules", this.InputAutomationModuleJPanel)
        var automationModuleConfigurationJPanelLayout *GroupLayout = NewGroupLayout(this.AutomationModuleConfigurationJPanel); _ = automationModuleConfigurationJPanelLayout
        this.AutomationModuleConfigurationJPanel.SetLayout(automationModuleConfigurationJPanelLayout)
        automationModuleConfigurationJPanelLayout.SetHorizontalGroup(automationModuleConfigurationJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 566, Short_MAX_VALUE))
        automationModuleConfigurationJPanelLayout.SetVerticalGroup(automationModuleConfigurationJPanelLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 334, Short_MAX_VALUE))
        this.GameRobotJTabbedPane.AddTab("Configuration", this.AutomationModuleConfigurationJPanel)
        this.ProcessingJMenu.SetText("Processing")
        this.ProcessingJMenu.SetToolTipText("Set module processing state")
        this.StartJMenuItem.SetText(this.CommonStrings.START)
        this.StartJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) ActionPerformed(evt *ActionEvent) {
                this.StartJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.StartJMenuItem)
        this.StopJMenuItem.SetText("Stop")
        this.StopJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) ActionPerformed(evt *ActionEvent) {
                this.StopJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ProcessingJMenu.Add(this.StopJMenuItem)
        this.MainJMenuBar.Add(this.ProcessingJMenu)
        this.OptionsJMenu.SetText("Options")
        this.OptionsJMenu.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) ActionPerformed(evt *ActionEvent) {
                this.OptionsJMenuActionPerformed(evt)
        }
*/ NewActionListener())
        this.StopOnFocusJCheckBoxMenuItem.SetSelected(true)
        this.StopOnFocusJCheckBoxMenuItem.SetEnabled(false)
        this.StopOnFocusJCheckBoxMenuItem.SetLabel("Stop Module On Focus")
        this.StopOnFocusJCheckBoxMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) ActionPerformed(evt *ActionEvent) {
                this.StopOnFocusJCheckBoxMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.OptionsJMenu.Add(this.StopOnFocusJCheckBoxMenuItem)
        this.MainJMenuBar.Add(this.OptionsJMenu)
        this.ToolsJMenu.SetText("Help")
        this.HelpJMenuItem.SetText("Help")
        this.HelpJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) ActionPerformed(evt *ActionEvent) {
                this.HelpJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ToolsJMenu.Add(this.HelpJMenuItem)
        this.UpdatesJMenuItem.SetText("Updates")
        this.UpdatesJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) ActionPerformed(evt *ActionEvent) {
                this.UpdatesJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ToolsJMenu.Add(this.UpdatesJMenuItem)
        this.SubscriptionJMenuItem.SetText("Subscription")
        this.SubscriptionJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) ActionPerformed(evt *ActionEvent) {
                this.SubscriptionJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ToolsJMenu.Add(this.SubscriptionJMenuItem)
        this.ModulesJMenuItem.SetText("Module Manager")
        this.ModulesJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) ActionPerformed(evt *ActionEvent) {
                this.ModulesJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ToolsJMenu.Add(this.ModulesJMenuItem)
        this.AboutJMenuItem.SetText("About")
        this.AboutJMenuItem.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) ActionPerformed(evt *ActionEvent) {
                this.AboutJMenuItemActionPerformed(evt)
        }
*/ NewActionListener())
        this.ToolsJMenu.Add(this.AboutJMenuItem)
        this.MainJMenuBar.Add(this.ToolsJMenu)
        this.SetJMenuBar(this.MainJMenuBar)
        var layout *GroupLayout = NewGroupLayout(this.GetContentPane()); _ = layout
        this.GetContentPane().SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.GameRobotJTabbedPane, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.DEFAULT_SIZE, 571, Short_MAX_VALUE))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.GameRobotJTabbedPane, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.DEFAULT_SIZE, 359, Short_MAX_VALUE))
        this.Pack()
}


func (this *InputAutomationJFrame) SubscriptionJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "subscriptionJMenuItemActionPerformed", e)
                }
            }()
            Desktop_GetDesktop().Browse(this.Uri)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InputAutomationJFrame) UpdatesJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "updatesJMenuItemActionPerformed", e)
                }
            }()
            Desktop_GetDesktop().Browse(this.Uri)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InputAutomationJFrame) AboutJMenuItemActionPerformed(evt *ActionEvent) {
        var textJDialog *JDialog = NewBasicTextJDialog("Copyright (c) 2002-2007 AllBinary"); _ = textJDialog
        textJDialog.SetVisible(true)
        textJDialog.SetLocationRelativeTo(this)
}


func (this *InputAutomationJFrame) ModulesJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "modulesJMenuItemActionPerformed", e)
                }
            }()
            NewDesktopBundle().Start()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InputAutomationJFrame) HelpJMenuItemActionPerformed(evt *ActionEvent) {
        JavaHelpUtil_GetInstance().Show(this.HelpSet)
}


func (this *InputAutomationJFrame) StopOnFocusJCheckBoxMenuItemActionPerformed(evt *ActionEvent) {
}


func (this *InputAutomationJFrame) OptionsJMenuActionPerformed(evt *ActionEvent) {
}


func (this *InputAutomationJFrame) FormWindowLostFocus(evt *WindowEvent) {
        this.LogUtil.PutF("Nothing", this, "focusLost")
}


func (this *InputAutomationJFrame) FormWindowGainedFocus(evt *WindowEvent) {
        if this.StopOnFocusJCheckBoxMenuItem.IsSelected() {
            this.LogUtil.PutF("Stopping", this, "focusGained")
            if this.RunnableInterface != nil {
                this.RunnableInterface.SetRunning(false)
            }
        }
}


func (this *InputAutomationJFrame) FormFocusLost(evt *FocusEvent) {
}


func (this *InputAutomationJFrame) FormFocusGained(evt *FocusEvent) {
}


func (this *InputAutomationJFrame) InputAutomationModuleJListValueChanged(evt *ListSelectionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "gameAutomationRobotModuleJListValueChanged", e)
                }
            }()
            var gameAutomationRobotModuleNameString string = any(this.InputAutomationModuleJList.GetSelectedValue()).(string); _ = gameAutomationRobotModuleNameString
            if gameAutomationRobotModuleNameString != "" {
                this.InputAutomationModuleInterface = this.InputAutomationModuleFactory.GetInstance(gameAutomationRobotModuleNameString)
                this.GetAutomationModuleConfigurationJPanel().RemoveAll()
                var layout *GroupLayout = any(this.GetAutomationModuleConfigurationJPanel().GetLayout()).(*GroupLayout); _ = layout
                layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 395, Short_MAX_VALUE).AddComponent(this.InputAutomationModuleInterface.GetConfigurationJPanel(), javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))
                layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 256, Short_MAX_VALUE).AddComponent(this.InputAutomationModuleInterface.GetConfigurationJPanel(), javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))
                this.GameRobotJTabbedPane.SetEnabledAt(1, true)
                this.LogUtil.PutF("Setting Module: " + gameAutomationRobotModuleNameString + " and Configuration JPanel", this, "gameAutomationRobotModuleJListValueChanged")
            } else {
                this.LogUtil.PutF("Module name was null", this, "gameAutomationRobotModuleJListValueChanged")
                this.GameRobotJTabbedPane.SetEnabledAt(1, false)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InputAutomationJFrame) StopJMenuItemActionPerformed(evt *ActionEvent) {
        if this.RunnableInterface != nil {
            this.RunnableInterface.SetRunning(false)
        } else {
            this.NoModuleSelectedJDialog.SetVisible(true)
            this.NoModuleSelectedJDialog.SetLocationRelativeTo(this)
        }
}


func (this *InputAutomationJFrame) StartJMenuItemActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "startJMenuItemActionPerformed", e)
                }
            }()
            this.LogUtil.PutF("Starting", this, "startJMenuItemActionPerformed")
            if this.InputAutomationModuleInterface != nil {
                this.RunnableInterface = this.InputAutomationModuleInterface.GetInstance()
                if this.RunnableInterface != nil {
                    this.Thread = NewThread(this.RunnableInterface)
                    this.Thread.SetPriority(Thread_MIN_PRIORITY + 1)
                    this.Thread.Start()
                } else {
                    this.LogUtil.PutF("Runnnable was null", this, "startJMenuItemActionPerformed")
                    this.NoModuleSelectedJDialog.SetVisible(true)
                    this.NoModuleSelectedJDialog.SetLocationRelativeTo(this)
                }
            } else {
                this.LogUtil.PutF("InputAutomationModuleInterface was null", this, "startJMenuItemActionPerformed")
                this.NoModuleSelectedJDialog.SetVisible(true)
                this.NoModuleSelectedJDialog.SetLocationRelativeTo(this)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func InputAutomationJFrame_GetInstance() *InputAutomationJFrame {
        return InputAutomationJFrame_INPUTAUTOMATION_JFRAME
}


func InputAutomationJFrame_Destroy() {
        InputAutomationJFrame_INPUTAUTOMATION_JFRAME.SetVisible(false)
}


func InputAutomationJFrame_Create(inputAutomationBundleActivatorListenerInterface InputAutomationBundleActivatorListenerInterface) {
        LogConfigTypeFactory_GetInstance()
        var appUrlGlobals *AppUrlGlobals = NewAppUrlGlobals(); _ = appUrlGlobals
        appUrlGlobals.SetWebappPath("G:/mnt/bc/mydev/work_automation/automation/run/")
        JFrame_URLGLOBALS.Init(appUrlGlobals)
        WebappClassLoaderInfo_SetLoader(reflect.TypeOf(appUrlGlobals).GetClassLoader())
        java_Awt.EventQueue.InvokeLater(/* Java anonymous class body: implement it as a named type
            

        func (this *InputAutomationJFrame) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                        }
                    }()
                    logUtil.PutF("Running", this, commonStrings.RUN)
                    InputAutomationJFrame_INPUTAUTOMATION_JFRAME = NewInputAutomationJFrame()
                    InputRobotFactory_GetInstance().AddListener(InputAutomationJFrame_GetInstance())
                    if inputAutomationBundleActivatorListenerInterface != nil {
                        inputAutomationBundleActivatorListenerInterface.RegisterAsService()
                        inputAutomationBundleActivatorListenerInterface.UseServices()
                    }
                    InputAutomationJFrame_INPUTAUTOMATION_JFRAME.SetVisible(true)
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewARunnable())
}


// throws Exception
func InputAutomationJFrame_Main(args []string) {
        InputAutomationJFrame_Create(nil)
}


func (this *InputAutomationJFrame) GetAutomationModuleConfigurationJPanel() *JPanel {
        return this.AutomationModuleConfigurationJPanel
}


func (this *InputAutomationJFrame) SetAutomationModuleConfigurationJPanel(automationModuleConfigurationJPanel *JPanel) {
        this.AutomationModuleConfigurationJPanel = automationModuleConfigurationJPanel
}


func (this *InputAutomationJFrame) OnAdd_1(inputAutomationRobotChangeEvent *InputAutomationRobotChangeEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onAdd", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "onAdd")
            InputRobotFactory_GetInstance().Add(inputAutomationRobotChangeEvent.GetInputAutomationRobotInterfaceWrapper().GetInputRobotInterface())
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of InputAutomationJFrame.onAdd
func (this *InputAutomationJFrame) OnAdd(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*InputAutomationRobotChangeEvent](args[0], pass) {
            this.OnAdd_1(Java_Arg[*InputAutomationRobotChangeEvent](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[*InputAutomationConfigurationModuleChangeEvent](args[0], pass) {
            this.OnAdd_2(Java_Arg[*InputAutomationConfigurationModuleChangeEvent](args[0]))
            return
        }
    }
    panic("No Java overload of InputAutomationJFrame.onAdd for the arguments")
}


func (this *InputAutomationJFrame) OnRemove_1(inputAutomationRobotChangeEvent *InputAutomationRobotChangeEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onRemove", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "onRemove")
            InputRobotFactory_GetInstance().Add(inputAutomationRobotChangeEvent.GetInputAutomationRobotInterfaceWrapper().GetInputRobotInterface())
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of InputAutomationJFrame.onRemove
func (this *InputAutomationJFrame) OnRemove(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*InputAutomationRobotChangeEvent](args[0], pass) {
            this.OnRemove_1(Java_Arg[*InputAutomationRobotChangeEvent](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[*InputAutomationConfigurationModuleChangeEvent](args[0], pass) {
            this.OnRemove_2(Java_Arg[*InputAutomationConfigurationModuleChangeEvent](args[0]))
            return
        }
    }
    panic("No Java overload of InputAutomationJFrame.onRemove for the arguments")
}


func (this *InputAutomationJFrame) OnAdd_2(inputAutomationConfigurationChangeEvent *InputAutomationConfigurationModuleChangeEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onAdd", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "onAdd")
            this.InputAutomationModuleConfigurations.Add(inputAutomationConfigurationChangeEvent.GetInputAutomationModuleConfiguration())
            this.Init()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InputAutomationJFrame) OnRemove_2(inputAutomationConfigurationChangeEvent *InputAutomationConfigurationModuleChangeEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onRemove", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "onRemove")
            var inputAutomationConfiguration *InputAutomationConfiguration = InputAutomationConfigurationFactory_GetInstance().InputAutomationConfiguration; _ = inputAutomationConfiguration
            var inputAutomationModuleConfigurationList IList[*InputAutomationModuleConfiguration] = inputAutomationConfiguration.GetInputAutomationModuleConfigurationList(); _ = inputAutomationModuleConfigurationList
            if inputAutomationModuleConfigurationList == nil {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                logUtil.PutF("inputAutomationModuleConfigurationList: " + fmt.Sprint(inputAutomationModuleConfigurationList), this, commonStrings.INIT)
                panic(errors.New(fmt.Sprint()))
            }
            var inputAutomationModuleConfigurations *InputAutomationModuleConfigurations = NewInputAutomationModuleConfigurations(inputAutomationConfiguration.GetInputAutomationModuleConfigurationList()); _ = inputAutomationModuleConfigurations
            inputAutomationModuleConfigurations.Remove(inputAutomationConfigurationChangeEvent.GetInputAutomationModuleConfiguration())
            this.Init()
            return
        }(); tryControl1 == 1 {
            return
        }
}

