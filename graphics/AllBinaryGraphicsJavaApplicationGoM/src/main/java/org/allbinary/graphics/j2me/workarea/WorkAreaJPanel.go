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

type WorkAreaJPanel struct {
    JPanel
    LogUtil *LogUtil
    CanvasJPanelList *LinkedList
    CanvasHolderJPanel *JPanel
    Name string
    SelectedFrame int
    PropertiesJPanel *PropertiesJPanel
    WorkAreaJTreeJPanel *WorkAreaJTreeJPanel
    IsPlaying bool
    InnerJPanel *JPanel
}

// Implements WorkAreaJPanelInterface, MyCanvasEventListener, *EventListener, Runnable

var WorkAreaJPanel_Document *Document

func (this *WorkAreaJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    this.IsPlaying = false
    
}




// Java overloads of new WorkAreaJPanel
func NewWorkAreaJPanel(args ...any) *WorkAreaJPanel {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[string](args[0], pass) && Java_IsArg[*Dimension](args[1], pass) && Java_IsArg[int](args[2], pass) && Java_IsArg[int](args[3], pass) {
            return NewWorkAreaJPanel_1(Java_Arg[string](args[0]), Java_Arg[*Dimension](args[1]), Java_Arg[int](args[2]), Java_Arg[int](args[3]))
        }
        if len(args) == 2 && Java_IsArg[*WorkAreaDom](args[0], pass) && Java_IsArg[*Dimension](args[1], pass) {
            return NewWorkAreaJPanel_2(Java_Arg[*WorkAreaDom](args[0]), Java_Arg[*Dimension](args[1]))
        }
    }
    panic("No Java overload of new WorkAreaJPanel for the arguments")
}
func NewWorkAreaJPanel_1(name string, dimension *Dimension, x int, y int) *WorkAreaJPanel {
    this := &WorkAreaJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            this.InitComponents()
            this.Name = name
            this.InitMyComponents(dimension)
            this.WorkAreaJTreeJPanel = NewWorkAreaJTreeJPanel(this.Name)
            var canvasJPanel *CanvasJPanel = NewCanvasJPanel(this.WorkAreaJTreeJPanel, this.GetSize(), x, y); _ = canvasJPanel
            this.CanvasJPanelList.Add(canvasJPanel)
            this.SelectedFrame = 0
            this.PropertiesJPanel.Set(this.WorkAreaJTreeJPanel)
            this.CanvasHolderJPanel.Add(any(canvasJPanel).(*Component))
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func NewWorkAreaJPanel_2(workAreaDom *WorkAreaDom, dimension *Dimension) *WorkAreaJPanel {
    this := &WorkAreaJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if JPanel_LOGGING.Contains(JPanel_LOGGING.LOADINGERROR) {
                this.LogUtil.Put("Constructor Error", this, "WorkAreaJPanel", e)
            }
            panic(e)
                }
            }()
            this.InitComponents()
            this.Name = workAreaDom.GetName()
            this.InitMyComponents(dimension)
            var canvasNodeList IBasicArrayList = workAreaDom.GetCanvasNodes(); _ = canvasNodeList
            var numberOfFrames int = canvasNodeList.Size(); _ = numberOfFrames
            for index := int(0); index < numberOfFrames; index++ {
                var node *Node = any(canvasNodeList.Get(index)).(*Node); _ = node
                var canvasJPanel *CanvasJPanel = NewCanvasJPanel(this.WorkAreaJTreeJPanel, this.GetSize(), NewCanvasDom(node)); _ = canvasJPanel
                this.CanvasJPanelList.Add(canvasJPanel)
            }
            this.SelectedFrame = 0
            this.CanvasHolderJPanel.Add(any(this.GetCurrentFrame()).(*Component))
            this.PropertiesJPanel.Set(this.WorkAreaJTreeJPanel)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *WorkAreaJPanel) InitMyComponents(dimension *Dimension) {
        this.WorkAreaJTreeJPanel = NewWorkAreaJTreeJPanel(this.Name)
        this.CanvasHolderJPanel = NewJPanel()
        this.CanvasHolderJPanel.SetSize(dimension)
        this.CanvasHolderJPanel.SetLayout(NewGridLayout(1, 1))
        this.SetSize(dimension)
        this.CanvasJPanelList = NewLinkedList()
        var gridBagLayout *GridBagLayout = NewGridBagLayout(); _ = gridBagLayout
        var gridBagConstraints *GridBagConstraints = NewGridBagConstraints(); _ = gridBagConstraints
        this.InnerJPanel.SetLayout(gridBagLayout)
        gridBagConstraints.Gridx = 0
        gridBagConstraints.Gridy = 0
        gridBagConstraints.Gridwidth = 1
        gridBagConstraints.Fill = java_Awt.GridBagConstraints.BOTH
        gridBagConstraints.Weightx = .05
        gridBagConstraints.Weighty = .05
        var toolJPanel *ToolJPanel = NewToolJPanel(); _ = toolJPanel
        gridBagLayout.SetConstraints(toolJPanel, gridBagConstraints)
        this.InnerJPanel.Add(toolJPanel)
        gridBagConstraints = NewGridBagConstraints()
        gridBagConstraints.Gridx = 1
        gridBagConstraints.Gridy = 0
        gridBagConstraints.Gridwidth = 6
        gridBagConstraints.Fill = java_Awt.GridBagConstraints.BOTH
        gridBagConstraints.Weightx = 1
        gridBagConstraints.Weighty = 1
        gridBagLayout.SetConstraints(this.CanvasHolderJPanel, gridBagConstraints)
        this.InnerJPanel.Add(this.CanvasHolderJPanel)
        gridBagConstraints = NewGridBagConstraints()
        gridBagConstraints.Gridx = 8
        gridBagConstraints.Gridy = 0
        gridBagConstraints.Gridwidth = java_Awt.GridBagConstraints.REMAINDER
        gridBagConstraints.Fill = java_Awt.GridBagConstraints.BOTH
        gridBagConstraints.Weightx = .1
        gridBagConstraints.Weighty = .1
        this.PropertiesJPanel = NewPropertiesJPanel()
        gridBagLayout.SetConstraints(this.PropertiesJPanel, gridBagConstraints)
        this.InnerJPanel.Add(this.PropertiesJPanel)
        MyCanvasEventService_AddListener(any(this).(MyCanvasEventListener))
}


func (this *WorkAreaJPanel) GetCanvasHashMap() *HashMap {
        var canvasJPanelHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = canvasJPanelHashMap
        var size int = this.CanvasJPanelList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var canvasJPanel *CanvasJPanel = any(this.CanvasJPanelList.Get(index)).(*CanvasJPanel); _ = canvasJPanel
            canvasJPanelHashMap.Put(canvasJPanel.GetTreeNode(), int(index))
        }
        return canvasJPanelHashMap
}


// synchronized
func (this *WorkAreaJPanel) Play() {
        this.IsPlaying = true
}


// synchronized
func (this *WorkAreaJPanel) Stop() {
        this.IsPlaying = false
}


// synchronized
func (this *WorkAreaJPanel) IsPlaying() bool {
        return this.IsPlaying
}


func (this *WorkAreaJPanel) GetName() string {
        return this.Name
}


func (this *WorkAreaJPanel) SetName(value string) {
        this.Name = value
}


func (this *WorkAreaJPanel) DisableCanvas() {
        this.CanvasHolderJPanel.RemoveAll()
}


func (this *WorkAreaJPanel) EnableCanvas() {
        this.CanvasHolderJPanel.Add(any(this.GetCurrentFrame()).(*Component))
}


func (this *WorkAreaJPanel) Deselect() {
        this.JPanel.SetVisible(false)
}


func (this *WorkAreaJPanel) Select_1() {
        this.JPanel.SetVisible(true)
        if this.CanvasHolderJPanel.GetComponentCount() == 1 {
            var component *Component = this.CanvasHolderJPanel.GetComponent(0); _ = component
            if component != nil {
                component.Repaint()
            }
        }
}

// Java overloads of WorkAreaJPanel.select
func (this *WorkAreaJPanel) Select(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            this.Select_1()
            return
        }
        if len(args) == 1 && Java_IsArg[int](args[0], pass) {
            this.Select_2(Java_Arg[int](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[*MutableTreeNode](args[0], pass) {
            this.Select_3(Java_Arg[*MutableTreeNode](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[*MyCanvasEvent](args[0], pass) {
            this.Select_4(Java_Arg[*MyCanvasEvent](args[0]))
            return
        }
    }
    panic("No Java overload of WorkAreaJPanel.select for the arguments")
}


func (this *WorkAreaJPanel) ChangeZoom(factor int) {
        var size int = this.CanvasJPanelList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var canvasJPanel *CanvasJPanel = any(this.CanvasJPanelList.Get(index)).(*CanvasJPanel); _ = canvasJPanel
            canvasJPanel.GetGrid().SetZoom(canvasJPanel.GetGrid().GetZoom() + factor)
        }
}


func (this *WorkAreaJPanel) GetCurrentFrame() *CanvasJPanel {
        return any(this.CanvasJPanelList.Get(this.SelectedFrame)).(*CanvasJPanel)
}


func (this *WorkAreaJPanel) InitDuplicate(newCanvasJPanel *CanvasJPanel) {
        if newCanvasJPanel != nil {
            this.CanvasJPanelList.Add(newCanvasJPanel)
            this.Select(this.CanvasJPanelList.Size() - 1)
        }
}


// throws Exception
func (this *WorkAreaJPanel) Duplicate_1(treeNode *MutableTreeNode) {
        var canvasJPanelHashMap *HashMap = this.GetCanvasHashMap(); _ = canvasJPanelHashMap
        if canvasJPanelHashMap.ContainsKey(treeNode) {
            var frameInteger int = any(canvasJPanelHashMap.Get(treeNode)).(int); _ = frameInteger
            var canvasJPanel *CanvasJPanel = any(this.CanvasJPanelList.Get(int(frameInteger))).(*CanvasJPanel); _ = canvasJPanel
            this.InitDuplicate(canvasJPanel.Duplicate().(*CanvasJPanel))
        }
}

// Java overloads of WorkAreaJPanel.duplicate
func (this *WorkAreaJPanel) Duplicate(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*MutableTreeNode](args[0], pass) {
            this.Duplicate_1(Java_Arg[*MutableTreeNode](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[*MyCanvasEvent](args[0], pass) {
            this.Duplicate_2(Java_Arg[*MyCanvasEvent](args[0]))
            return
        }
    }
    panic("No Java overload of WorkAreaJPanel.duplicate for the arguments")
}


// throws Exception
func (this *WorkAreaJPanel) AutoRotate(increments int, totalAngle int) {
        var incrementAngle int = (totalAngle / increments); _ = incrementAngle
        for index := int(1); index < increments; index++ {
            var canvasJPanel *CanvasJPanel = any(this.GetCurrentFrame()).(*CanvasJPanel); _ = canvasJPanel
            var newCanvasJPanel *CanvasJPanel = canvasJPanel.Duplicate().(*CanvasJPanel); _ = newCanvasJPanel
            newCanvasJPanel.SetAngle(incrementAngle * index)
            this.InitDuplicate(newCanvasJPanel)
        }
}


func (this *WorkAreaJPanel) ExplodeAll() {
        var canvasJPanel *CanvasJPanel = any(this.GetCurrentFrame()).(*CanvasJPanel); _ = canvasJPanel
        canvasJPanel.ExplodeAll()
}


// throws Exception
func (this *WorkAreaJPanel) AutoExplode_1(numberOfFrames int, explodeType int) {
        for index := int(1); index < numberOfFrames; index++ {
            var canvasJPanel *CanvasJPanel = any(this.GetCurrentFrame()).(*CanvasJPanel); _ = canvasJPanel
            var newCanvasJPanel *CanvasJPanel = canvasJPanel.Duplicate().(*CanvasJPanel); _ = newCanvasJPanel
            newCanvasJPanel.Explode(index + 1, explodeType)
            this.InitDuplicate(newCanvasJPanel)
        }
}

// Java overloads of WorkAreaJPanel.autoExplode
func (this *WorkAreaJPanel) AutoExplode(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) {
            this.AutoExplode_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]))
            return
        }
        if len(args) == 1 && Java_IsArg[*MyCanvasEvent](args[0], pass) {
            this.AutoExplode_2(Java_Arg[*MyCanvasEvent](args[0]))
            return
        }
    }
    panic("No Java overload of WorkAreaJPanel.autoExplode for the arguments")
}


// throws Exception
func (this *WorkAreaJPanel) AutoMirror() {
        var canvasJPanel *CanvasJPanel = any(this.GetCurrentFrame()).(*CanvasJPanel); _ = canvasJPanel
        var newCanvasJPanel *CanvasJPanel = canvasJPanel.Duplicate().(*CanvasJPanel); _ = newCanvasJPanel
        newCanvasJPanel.Mirror()
        this.InitDuplicate(newCanvasJPanel)
}


// synchronized
func (this *WorkAreaJPanel) Select_2(index int) {
        if index < this.CanvasJPanelList.Size() {
            var canvasJPanel *CanvasJPanel = any(this.CanvasJPanelList.Get(index)).(*CanvasJPanel); _ = canvasJPanel
            if canvasJPanel != nil {
                this.CanvasHolderJPanel.RemoveAll()
                this.CanvasHolderJPanel.Add(canvasJPanel)
                this.SelectedFrame = index
                canvasJPanel.Changed()
                this.CanvasHolderJPanel.Repaint()
                canvasJPanel.Repaint()
            }
        }
}


func (this *WorkAreaJPanel) Select_3(treeNode *MutableTreeNode) {
        var canvasJPanelHashMap *HashMap = this.GetCanvasHashMap(); _ = canvasJPanelHashMap
        var frameInteger int = any(canvasJPanelHashMap.Get(treeNode)).(int); _ = frameInteger
        if frameInteger != 0 {
            this.Select(int(frameInteger))
        }
}


func WorkAreaJPanel_GetDocument() *Document {
        return WorkAreaJPanel_Document
}


// throws Exception
func (this *WorkAreaJPanel) ToDocument() *Document {
        WorkAreaJPanel_Document = DomHelper_GetInstance().CreateDocument()
        var workAreaNode *Node = WorkAreaJPanel_Document.CreateElement(WorkAreaDom_WORKAREA); _ = workAreaNode
        var nameNode *Node = WorkAreaJPanel_Document.CreateElement(WorkAreaDom_NAME); _ = nameNode
        var nameTextNode *Node = WorkAreaJPanel_Document.CreateTextNode(this.Name); _ = nameTextNode
        nameNode.AppendChild(nameTextNode)
        var framesNode *Node = WorkAreaJPanel_Document.CreateElement(WorkAreaDom_FRAMES); _ = framesNode
        var size int = this.CanvasJPanelList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var canvasJPanel *CanvasJPanel = any(this.CanvasJPanelList.Get(index)).(*CanvasJPanel); _ = canvasJPanel
            framesNode.AppendChild(canvasJPanel.ToDom())
        }
        workAreaNode.AppendChild(nameNode)
        workAreaNode.AppendChild(framesNode)
        WorkAreaJPanel_Document.AppendChild(workAreaNode)
        return WorkAreaJPanel_Document
}


func (this *WorkAreaJPanel) InitComponents() {
        this.InnerJPanel = NewJPanel()
        this.SetLayout(NewGridLayout(1, 1))
        this.InnerJPanel.AddKeyListener(/* Java anonymous class body: implement it as a named type
            

        func (this *WorkAreaJPanel) KeyTyped(evt IKeyEvent) {
                this.InnerJPanelKeyTyped(evt)
        }
            

        func (this *WorkAreaJPanel) KeyPressed(evt IKeyEvent) {
                this.InnerJPanelKeyPressed(evt)
        }
            

        func (this *WorkAreaJPanel) KeyReleased(evt IKeyEvent) {
                this.InnerJPanelKeyReleased(evt)
        }
*/ NewKeyAdapter())
        this.Add(this.InnerJPanel)
}


func (this *WorkAreaJPanel) InnerJPanelKeyReleased(evt IKeyEvent) {
        StatusFactory_GetInstance().SetStatus("Key Pressed1")
}


func (this *WorkAreaJPanel) InnerJPanelKeyPressed(evt IKeyEvent) {
        StatusFactory_GetInstance().SetStatus("Key Pressed2")
}


func (this *WorkAreaJPanel) InnerJPanelKeyTyped(evt IKeyEvent) {
        StatusFactory_GetInstance().SetStatus("Key Pressed3")
}


func (this *WorkAreaJPanel) KeyPressed(keyEvent IKeyEvent) {
        this.GetCurrentFrame().KeyPressed(keyEvent)
}


func (this *WorkAreaJPanel) KeyReleased(keyEvent IKeyEvent) {
}


func (this *WorkAreaJPanel) KeyTyped(keyEvent IKeyEvent) {
}


func (this *WorkAreaJPanel) Delete(evt *MyCanvasEvent) {
        var canvasJPanelHashMap *HashMap = this.GetCanvasHashMap(); _ = canvasJPanelHashMap
        if canvasJPanelHashMap.ContainsKey(evt.GetTreeNode()) {
            this.CanvasHolderJPanel.RemoveAll()
            var frameInt int = any(canvasJPanelHashMap.Get(evt.GetTreeNode())).(int); _ = frameInt
            this.CanvasJPanelList.Remove(int(frameInt))
            this.Select(int(frameInt))
        }
}


// throws Exception
func (this *WorkAreaJPanel) Duplicate_2(evt *MyCanvasEvent) {
        this.Duplicate(evt.GetTreeNode())
}


func (this *WorkAreaJPanel) Select_4(evt *MyCanvasEvent) {
        this.Select(evt.GetTreeNode())
        StatusFactory_GetInstance().SetStatus("Frame Selected")
}


func (this *WorkAreaJPanel) Rotate(evt *MyCanvasEvent) {
        var canvasJPanelHashMap *HashMap = this.GetCanvasHashMap(); _ = canvasJPanelHashMap
        var treeNode *MutableTreeNode = evt.GetTreeNode(); _ = treeNode
        var frameInt int = any(canvasJPanelHashMap.Get(evt.GetTreeNode())).(int); _ = frameInt
        var canvasJPanel *CanvasJPanel = any(this.CanvasJPanelList.Get(int(frameInt))).(*CanvasJPanel); _ = canvasJPanel
        if canvasJPanel != nil {
            canvasJPanel.AddAngle(evt.GetAngle())
            this.Select(int(frameInt))
        }
        StatusFactory_GetInstance().SetStatus("Canvas Rotated")
}


func (this *WorkAreaJPanel) Explode(evt *MyCanvasEvent) {
        this.ExplodeAll()
}


// throws Exception
func (this *WorkAreaJPanel) AutoExplode_2(evt *MyCanvasEvent) {
        this.AutoExplode(10, 0)
}


// throws Exception
func (this *WorkAreaJPanel) Center_1() {
        var canvasJPanel *CanvasJPanel = any(this.GetCurrentFrame()).(*CanvasJPanel); _ = canvasJPanel
        canvasJPanel.Center()
}

// Java overloads of WorkAreaJPanel.center
func (this *WorkAreaJPanel) Center(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            this.Center_1()
            return
        }
        if len(args) == 1 && Java_IsArg[*MyCanvasEvent](args[0], pass) {
            this.Center_2(Java_Arg[*MyCanvasEvent](args[0]))
            return
        }
    }
    panic("No Java overload of WorkAreaJPanel.center for the arguments")
}


// throws Exception
func (this *WorkAreaJPanel) Center_2(evt *MyCanvasEvent) {
        this.Center()
}


func (this *WorkAreaJPanel) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            for true {
                if this.IsPlaying() {
                    var size int = this.CanvasJPanelList.Size(); _ = size
                    for index := int(0); index < size; index++ {
                        if !this.IsPlaying() {
                            break
                        }
                        this.Select(index)
                        time.Sleep(time.Duration(100) * time.Millisecond)
                    }
                } else {
                    time.Sleep(time.Duration(500) * time.Millisecond)
                }
                time.Sleep(time.Duration(250) * time.Millisecond)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

