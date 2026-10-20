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

type CanvasJPanel struct {
    JPanel
    LogUtil *LogUtil
    GuiLog *GuiLog
    GameInputStrings *GameInputStrings
    SelectedTool GraphicItemInterface
    GraphicItemHashMap *HashMap
    CanvasDimension *IntegerDimension
    Angle float64
    CanvasTreeNode *DefaultMutableTreeNode
    FrameLabel *CanvasTreeLabel
    WorkAreaJTreeJPanel *WorkAreaJTreeJPanel
    Grid *Grid
}

// Implements MouseListener, *MouseMotionListener, KeyListener, MyGraphicItemEventListener

var CanvasJPanel_Frame int = 0
var CanvasJPanel_GridColor *Color = NewColor(BasicColorFactory_GetInstance().WHITE.IntValue())
var CanvasJPanel_BackgroundColor *Color = NewColor(BasicColorFactory_GetInstance().CLEAR_COLOR.IntValue())

func (this *CanvasJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GuiLog = GuiLog_GetInstance()
    this.GameInputStrings = GameInputStrings_GetInstance()
    this.SelectedTool = nil
    
    
    this.Angle = 0.0
    
    
    
    
}




// Java overloads of new CanvasJPanel
func NewCanvasJPanel(args ...any) *CanvasJPanel {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[*WorkAreaJTreeJPanel](args[0], pass) && Java_IsArg[*Dimension](args[1], pass) && Java_IsArg[int](args[2], pass) && Java_IsArg[int](args[3], pass) {
            return NewCanvasJPanel_1(Java_Arg[*WorkAreaJTreeJPanel](args[0]), Java_Arg[*Dimension](args[1]), Java_Arg[int](args[2]), Java_Arg[int](args[3]))
        }
        if len(args) == 3 && Java_IsArg[*WorkAreaJTreeJPanel](args[0], pass) && Java_IsArg[*Dimension](args[1], pass) && Java_IsArg[*CanvasDom](args[2], pass) {
            return NewCanvasJPanel_2(Java_Arg[*WorkAreaJTreeJPanel](args[0]), Java_Arg[*Dimension](args[1]), Java_Arg[*CanvasDom](args[2]))
        }
    }
    panic("No Java overload of new CanvasJPanel for the arguments")
}
func NewCanvasJPanel_1(workAreaJTreeJPanel *WorkAreaJTreeJPanel, dimension *Dimension, x int, y int) *CanvasJPanel {
    this := &CanvasJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.Grid = NewGrid()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if JPanel_LOGGING.Contains(JPanel_LOGGING.GRAPHICSCREATION) {
                this.GuiLog.Put("Constructor Error", this, "contructor", e)
            }
                }
            }()
            this.WorkAreaJTreeJPanel = workAreaJTreeJPanel
            this.InitComponents()
            this.InitMyComponents()
            this.SetSize(dimension)
            this.SetCanvasSize(x, y)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func NewCanvasJPanel_2(workAreaJTreeJPanel *WorkAreaJTreeJPanel, dimension *Dimension, canvasDom *CanvasDom) *CanvasJPanel {
    this := &CanvasJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.Grid = NewGrid(canvasDom.GetGrid())
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if JPanel_LOGGING.Contains(JPanel_LOGGING.GRAPHICSCREATION) {
                this.GuiLog.Put("Constructor Error", this, "contructor", e)
            }
                }
            }()
            this.WorkAreaJTreeJPanel = workAreaJTreeJPanel
            this.InitComponents()
            this.InitMyComponents()
            this.SetSize(dimension)
            this.SetCanvasSize(canvasDom.GetDimension().GetWidth(), canvasDom.GetDimension().GetHeight())
            this.GraphicItemHashMap = canvasDom.GetGraphicItemHashMap()
            var graphicItemArray []any = this.GraphicItemHashMap.KeySet().ToArray(); _ = graphicItemArray
            var size int = len(graphicItemArray); _ = size
            for index := int(0); index < size; index++ {
                var graphicItemTreeNode *MutableTreeNode = any(graphicItemArray[index]).(*MutableTreeNode); _ = graphicItemTreeNode
                this.CanvasTreeNode.Add(graphicItemTreeNode)
            }
            this.Repaint()
            this.SetAngle(canvasDom.GetAngle())
            this.WorkAreaJTreeJPanel.UpdateTree()
            this.WorkAreaJTreeJPanel.Expand()
            this.WorkAreaJTreeJPanel.Repaint()
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


// throws Exception
func (this *CanvasJPanel) InitMyComponents() {
        this.FrameLabel = NewCanvasTreeLabel("FrameLabel" + fmt.Sprint(int(CanvasJPanel_Frame)))
        this.CanvasTreeNode = NewDefaultMutableTreeNode(this.FrameLabel)
        CanvasJPanel_Frame++
        this.WorkAreaJTreeJPanel.Add(this.GetTreeNode())
        this.GraphicItemHashMap = StdUtil_GetInstance().CreateHashMap()
        this.Grid.Grid = PointFactory_GetInstance().CreateXY(0, 0)
        this.SetCanvasDimension(NewIntegerDimension(0, 0))
        this.SelectedTool = nil
        this.AddKeyListener(this)
        this.AddMouseListener(this)
        this.AddMouseMotionListener(this)
        MyGraphicItemEventService_AddListener(this)
}


func (this *CanvasJPanel) GetTreeNode() *MutableTreeNode {
        return this.CanvasTreeNode
}


func (this *CanvasJPanel) InitComponents() {
        this.SetLayout(NewGridLayout(1, 1))
}


// throws Exception
func (this *CanvasJPanel) SetCanvasSize(x int, y int) {
        this.Grid.IsGridPossible = true
        if x <= 0 {
            this.Grid.IsGridPossible = false
            panic(errors.New(fmt.Sprint("X Size Error: " + fmt.Sprint(x), this, "setWorkAreaSize")))
        }
        if y <= 0 {
            this.Grid.IsGridPossible = false
            panic(errors.New(fmt.Sprint("Y Size Error: " + fmt.Sprint(y), this, "setWorkAreaSize")))
        }
        if x > this.GetWidth() {
            this.Grid.IsGridPossible = false
        }
        if y > this.GetHeight() {
            this.Grid.IsGridPossible = false
        }
        this.SetCanvasDimension(NewIntegerDimension(x, y))
        this.Grid.IsChanged = true
}


func (this *CanvasJPanel) SetGrid_1(point IGPoint) {
        this.SetGrid(point.GetX(), point.GetY())
}

// Java overloads of CanvasJPanel.setGrid
func (this *CanvasJPanel) SetGrid(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[IGPoint](args[0], pass) {
            this.SetGrid_1(Java_Arg[IGPoint](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) {
            this.SetGrid_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]))
            return
        }
    }
    panic("No Java overload of CanvasJPanel.setGrid for the arguments")
}


func (this *CanvasJPanel) SetGrid_2(xSize int, ySize int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            if xSize >= this.GetCanvasDimension().GetWidth() / 2 || xSize < 1 {
                panic(errors.New(fmt.Sprint("X Size Error: " + fmt.Sprint(xSize) + " WorkArea X: " + fmt.Sprint(this.GetCanvasDimension().GetWidth()), this, "setGrid")))
            }
            if ySize >= this.GetCanvasDimension().GetHeight() / 2 || ySize < 1 {
                panic(errors.New(fmt.Sprint("Y Size Error" + fmt.Sprint(ySize) + " WorkArea Y: " + fmt.Sprint(this.GetCanvasDimension().GetHeight()), this, "setGrid")))
            }
            this.Grid.Grid = PointFactory_GetInstance().CreateXY(xSize, ySize)
            this.Grid.IsChanged = true
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *CanvasJPanel) SetDefaultGrid() {
        var xPixelsPerWorkAreaPixel int = this.GetWidth() / this.GetCanvasDimension().GetWidth(); _ = xPixelsPerWorkAreaPixel
        var yPixelsPerWorkAreaPixel int = this.GetHeight() / this.GetCanvasDimension().GetHeight(); _ = yPixelsPerWorkAreaPixel
        this.SetGrid(1, 1)
        this.Grid.IsChanged = true
}


func (this *CanvasJPanel) GetAngle() float64 {
        return this.Angle
}


func (this *CanvasJPanel) SetAngle(angle float64) {
        this.Angle = angle
}


func (this *CanvasJPanel) AddAngle(angle float64) {
        this.SetAngle(this.Angle + angle)
}


func (this *CanvasJPanel) ExplodeAll() {
        var newPoints IBasicArrayList = NewBasicArrayListD(); _ = newPoints
        var graphicItemArray []any = this.GetGraphicItemHashMap().KeySet().ToArray(); _ = graphicItemArray
        var size int = len(graphicItemArray); _ = size
        for index := int(0); index < size; index++ {
            var graphicItemNode *MutableTreeNode = any(graphicItemArray[index]).(*MutableTreeNode); _ = graphicItemNode
            var graphicItem GraphicItemInterface = any(this.GetGraphicItemHashMap().Get(graphicItemNode)).(GraphicItemInterface); _ = graphicItem
            if graphicItem.GetName() == LinesGraphicItem_GetStaticName() {
                for graphicItem.GetPointsInterface().GetSize() > 2 {
                    var newGraphicItem GraphicItemInterface = GraphicItemFactory_GetInstance().GetInstance(graphicItem.GetName()).(GraphicsItemInterfaceFactoryInterface).GetInstance(this); _ = newGraphicItem
                    var pointOne IGPoint = GPoint_GetInstance(graphicItem.RemovePoint()); _ = pointOne
                    var pointTwo IGPoint = GPoint_GetInstance(graphicItem.RemovePoint()); _ = pointTwo
                    if pointOne != nil && pointTwo != nil {
                        newGraphicItem.AddPoint(pointOne)
                        newGraphicItem.AddPoint(pointTwo)
                    }
                    newPoints.Add(newGraphicItem)
                }
            }
        }
        var size2 int = newPoints.Size(); _ = size2
        for index := int(0); index < size2; index++ {
            var newGraphicItem GraphicItemInterface = any(newPoints.Get(index)).(GraphicItemInterface); _ = newGraphicItem
            this.CanvasTreeNode.Add(newGraphicItem.GetTreeNode())
            this.GetGraphicItemHashMap().Put(newGraphicItem.GetTreeNode(), newGraphicItem)
        }
        this.WorkAreaJTreeJPanel.UpdateTree()
        this.WorkAreaJTreeJPanel.Expand()
        this.WorkAreaJTreeJPanel.Repaint()
}


// throws Exception
func (this *CanvasJPanel) Explode(howMuch int, explosionType int) {
        var graphicItemArray []any = this.GetGraphicItemHashMap().KeySet().ToArray(); _ = graphicItemArray
        var size int = len(graphicItemArray); _ = size
        for index := int(0); index < size; index++ {
            var item GraphicItemInterface = any(this.GetGraphicItemHashMap().Get(graphicItemArray[index])).(GraphicItemInterface); _ = item
            var basicArrayList IBasicArrayList = VectorExplosionGenerator_GetInstance().CreateList(item.GetPointsInterface().GetPoints(), howMuch, VectorExplosionGenerator_GetInstance().RANDOM); _ = basicArrayList
            var newPoints *Points = NewPoints(); _ = newPoints
            newPoints.AddPoints(basicArrayList)
            item.SetPointsInterface(newPoints)
            var angleDelta int = RandomRotationFactory_GetInstance().GetNextRandomAngle(howMuch); _ = angleDelta
            item.SetAngle(angleDelta)
        }
        this.Repaint()
}


// throws Exception
func (this *CanvasJPanel) Mirror() {
        var width int = this.GetCanvasDimension().GetWidth(); _ = width
        var graphicItemArray []any = this.GetGraphicItemHashMap().KeySet().ToArray(); _ = graphicItemArray
        var size int = len(graphicItemArray); _ = size
        for index := int(0); index < size; index++ {
            var item GraphicItemInterface = any(this.GetGraphicItemHashMap().Get(graphicItemArray[index])).(GraphicItemInterface); _ = item
            var basicArrayList IBasicArrayList = VectorMirrorGenerator_GetInstance().GetInstance(item.GetPointsInterface().GetPoints(), width); _ = basicArrayList
            var newPoints *Points = NewPoints(); _ = newPoints
            newPoints.AddPoints(basicArrayList)
            item.SetPointsInterface(newPoints)
        }
        this.Repaint()
}


// throws Exception
func (this *CanvasJPanel) Center() {
        NewVectorCenterGenerator().Transform(this.GetGraphicItemHashMap())
        this.Repaint()
}


func (this *CanvasJPanel) SetNewTool() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            StatusFactory_GetInstance().SetStatus("Tool Selected")
            this.SelectedTool = ToolFactory_GetInstance().GetSelectedToolFactory().GetInstance(this)
            var newNode *MutableTreeNode = this.SelectedTool.GetTreeNode(); _ = newNode
            this.GetGraphicItemHashMap().Put(newNode, this.SelectedTool)
            this.CanvasTreeNode.Add(newNode)
            if this.GetSelectedTool() == nil {
                panic(errors.New(fmt.Sprint("No Tool Selected", this, "setSelectedTool")))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *CanvasJPanel) DuplicateGraphicItem(graphicItem GraphicItemInterface) {
        var graphicItemClone GraphicItemInterface = graphicItem.Duplicate(); _ = graphicItemClone
        this.GetGraphicItemHashMap().Put(graphicItemClone.GetTreeNode(), graphicItemClone)
        this.CanvasTreeNode.Add(graphicItemClone.GetTreeNode())
}


// throws Exception
func (this *CanvasJPanel) DuplicateGraphicItemHashMap(hashMap *HashMap) {
        var mutableTreeNodeArray []any = hashMap.KeySet().ToArray(); _ = mutableTreeNodeArray
        var size int = len(mutableTreeNodeArray); _ = size
        for index := int(0); index < size; index++ {
            var treeNode *MutableTreeNode = any(mutableTreeNodeArray[index]).(*MutableTreeNode); _ = treeNode
            var graphicItem GraphicItemInterface = any(hashMap.Get(treeNode)).(GraphicItemInterface); _ = graphicItem
            this.DuplicateGraphicItem(graphicItem)
        }
}


func (this *CanvasJPanel) GetGraphicItemHashMap() *HashMap {
        return this.GraphicItemHashMap
}


func (this *CanvasJPanel) GetXPixelsPerWorkAreaPixel() int {
        var x int = this.GetGrid().Grid.GetX() * (this.GetWidth() - 10) / this.GetCanvasDimension().GetWidth(); _ = x
        if this.GetGrid().GetZoom() == 1 {
            return 1
        } else {
            return x
        }
}


func (this *CanvasJPanel) GetYPixelsPerWorkAreaPixel() int {
        var y int = this.GetGrid().Grid.GetY() * (this.GetHeight() - 10) / this.GetCanvasDimension().GetHeight(); _ = y
        if this.GetGrid().GetZoom() == 1 {
            return 1
        } else {
            return y
        }
}


func (this *CanvasJPanel) GetSelectedTool() GraphicItemInterface {
        return this.SelectedTool
}


func (this *CanvasJPanel) DrawItems(graphics IGraphics) {
        var angleDouble float64 = float64(this.Angle); _ = angleDouble
        var graphicItemArray []any = this.GetGraphicItemHashMap().KeySet().ToArray(); _ = graphicItemArray
        var size int = len(graphicItemArray); _ = size
        for index := int(0); index < size; index++ {
            var item GraphicItemInterface = any(this.GetGraphicItemHashMap().Get(graphicItemArray[index])).(GraphicItemInterface); _ = item
            item.Paint(graphics, angleDouble, this.GetCanvasDimension(), this.GetXPixelsPerWorkAreaPixel(), this.GetYPixelsPerWorkAreaPixel())
        }
}


func (this *CanvasJPanel) DrawGrid(graphics IGraphics) {
        var xAdjust int = (this.GetWidth() % this.GetXPixelsPerWorkAreaPixel()) / 2; _ = xAdjust
        var yAdjust int = (this.GetHeight() % this.GetYPixelsPerWorkAreaPixel()) / 2; _ = yAdjust
        var xLeft int = xAdjust; _ = xLeft
        var yUpper int = yAdjust; _ = yUpper
        var xRight int = this.GetWidth() - xAdjust; _ = xRight
        var yLower int = this.GetHeight() - yAdjust; _ = yLower
        graphics.SetColor(CanvasJPanel_GridColor)
        for value := int(yAdjust); value <= this.GetHeight() - yAdjust; value += this.GetYPixelsPerWorkAreaPixel() {
            graphics.DrawLine(xLeft, value, xRight, value)
        }
        for value := int(xAdjust); value <= this.GetWidth() - xAdjust; value += this.GetYPixelsPerWorkAreaPixel() {
            graphics.DrawLine(value, yUpper, value, yLower)
        }
}


func (this *CanvasJPanel) Paint(graphics IGraphics) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            if this.GetGrid().IsChanged {
                graphics.SetColor(CanvasJPanel_BackgroundColor)
                graphics.FillRect(0, 0, this.GetWidth(), this.GetHeight())
                if this.GetGrid().GetZoom() > 2 && this.GetGrid().IsGridOn && this.GetGrid().IsGridPossible {
                    this.SetDefaultGrid()
                    this.DrawGrid(graphics)
                }
                this.DrawItems(graphics)
                graphics.DrawString(fmt.Sprint(float64(this.GetAngle())), this.GetWidth() - 75, this.GetHeight() - 15)
                this.Grid.IsChanged = false
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *CanvasJPanel) Repaint() {
        if this.Grid != nil {
            this.Grid.IsChanged = true
        }
        this.JPanel.Repaint()
}


// throws Exception
func (this *CanvasJPanel) ToDom() *Node {
        return NewCanvasDom(this).ToDom()
}


func (this *CanvasJPanel) Changed() {
        this.Grid.IsChanged = true
}


func (this *CanvasJPanel) MouseClicked(mouseEvent IMouseEvent) {
        StatusFactory_GetInstance().SetStatusNoLog(MouseStrings_GetInstance().MOUSE_CLICKED_LABEL + this.LogMouseEvent(mouseEvent))
        if this.GetSelectedTool() != nil && this.GetSelectedTool().IsActive() {
            var graphicItem GraphicItemInterface = this.GetSelectedTool(); _ = graphicItem
            graphicItem.MouseClicked(mouseEvent, this.GetXPixelsPerWorkAreaPixel(), this.GetYPixelsPerWorkAreaPixel())
        } else {
            this.SetNewTool()
            var graphicItem GraphicItemInterface = this.GetSelectedTool(); _ = graphicItem
            if graphicItem != nil {
                graphicItem.MouseClicked(mouseEvent, this.GetXPixelsPerWorkAreaPixel(), this.GetYPixelsPerWorkAreaPixel())
            }
        }
        this.Grid.IsChanged = true
        this.Repaint()
        this.WorkAreaJTreeJPanel.UpdateTree()
        this.WorkAreaJTreeJPanel.Expand()
        this.WorkAreaJTreeJPanel.Repaint()
}


func (this *CanvasJPanel) MouseEntered(mouseEvent IMouseEvent) {
        StatusFactory_GetInstance().SetStatusNoLog(this.LogMouseEvent(mouseEvent))
}


func (this *CanvasJPanel) MouseExited(mouseEvent IMouseEvent) {
        StatusFactory_GetInstance().SetStatusNoLog(this.LogMouseEvent(mouseEvent))
}


func (this *CanvasJPanel) MousePressed(mouseEvent IMouseEvent) {
        StatusFactory_GetInstance().SetStatusNoLog(MouseStrings_GetInstance().MOUSE_PRESSED_LABEL + this.LogMouseEvent(mouseEvent))
        if this.GetSelectedTool() != nil {
            var graphicItem GraphicItemInterface = this.GetSelectedTool(); _ = graphicItem
            graphicItem.MousePressed(mouseEvent, this.GetXPixelsPerWorkAreaPixel(), this.GetYPixelsPerWorkAreaPixel())
        }
        this.RequestFocus()
        this.Grid.IsChanged = true
        this.Repaint()
}


func (this *CanvasJPanel) MouseReleased(mouseEvent IMouseEvent) {
        StatusFactory_GetInstance().SetStatusNoLog(MouseStrings_GetInstance().MOUSE_RELEASED_LABEL + this.LogMouseEvent(mouseEvent))
        var graphicItem GraphicItemInterface = this.GetSelectedTool(); _ = graphicItem
        if graphicItem != nil {
            graphicItem.MouseReleased(mouseEvent, this.GetXPixelsPerWorkAreaPixel(), this.GetYPixelsPerWorkAreaPixel())
        }
        this.RequestFocus()
        this.Grid.IsChanged = true
        this.Repaint()
}


func (this *CanvasJPanel) MouseDragged(mouseEvent IMouseEvent) {
        StatusFactory_GetInstance().SetStatusNoLog(MouseStrings_GetInstance().MOUSE_DRAGGED_LABEL + this.LogMouseEvent(mouseEvent))
        var graphicItem GraphicItemInterface = this.GetSelectedTool(); _ = graphicItem
        if graphicItem != nil {
            graphicItem.MouseDragged(mouseEvent, this.GetXPixelsPerWorkAreaPixel(), this.GetYPixelsPerWorkAreaPixel())
        }
        this.RequestFocus()
        this.Grid.IsChanged = true
        this.Repaint()
}


func (this *CanvasJPanel) MouseMoved(mouseEvent IMouseEvent) {
        StatusFactory_GetInstance().SetStatusNoLog(MouseStrings_GetInstance().MOUSE_MOVED_LABEL + this.LogMouseEvent(mouseEvent))
        if this.GetSelectedTool() != nil && this.GetSelectedTool().IsActive() {
            var graphicItem GraphicItemInterface = this.GetSelectedTool(); _ = graphicItem
            graphicItem.MouseMoved(mouseEvent, this.GetXPixelsPerWorkAreaPixel(), this.GetYPixelsPerWorkAreaPixel())
        }
        this.RequestFocus()
        this.Grid.IsChanged = true
        this.Repaint()
}


func (this *CanvasJPanel) LogMouseEvent(mouseEvent IMouseEvent) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(PositionStrings_GetInstance().X_LABEL)
        stringBuffer.Appendint(int(mouseEvent.GetX() / this.GetXPixelsPerWorkAreaPixel()))
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(PositionStrings_GetInstance().Y_LABEL)
        stringBuffer.Appendint(int(mouseEvent.GetY() / this.GetYPixelsPerWorkAreaPixel()))
        return fmt.Sprint(stringBuffer)
}


func (this *CanvasJPanel) KeyPressed(keyEvent IKeyEvent) {
        StatusFactory_GetInstance().SetStatus("Key Pressed")
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, this.GameInputStrings.KEY_PRESSED, e)
                }
            }()
            var graphicItemArray []any = this.GraphicItemHashMap.KeySet().ToArray(); _ = graphicItemArray
            var size int = len(graphicItemArray); _ = size
            for index := int(0); index < size; index++ {
                var graphicItemInterface GraphicItemInterface = any(this.GraphicItemHashMap.Get(graphicItemArray[index])).(GraphicItemInterface); _ = graphicItemInterface
                var keyCode int = keyEvent.GetKeyCode(); _ = keyCode
                if graphicItemInterface.IsActive() {
                    if keyCode == keyEvent.KeyEvent_().VK_UP {
                        graphicItemInterface.Translate(0, -1)
                    } else if keyCode == keyEvent.KeyEvent_().VK_DOWN {
                        graphicItemInterface.Translate(0, 1)
                    } else if keyCode == keyEvent.KeyEvent_().VK_LEFT {
                        graphicItemInterface.Translate(-1, 0)
                    } else if keyCode == keyEvent.KeyEvent_().VK_RIGHT {
                        graphicItemInterface.Translate(1, 0)
                    }
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
        if this.GetSelectedTool() != nil && this.GetSelectedTool().IsActive() {
            StatusFactory_GetInstance().SetStatus("Key Pressed for Tool")
            var graphicItem GraphicItemInterface = this.GetSelectedTool(); _ = graphicItem
            graphicItem.KeyPressed(keyEvent)
        } else {
            this.SetNewTool()
        }
        this.Grid.IsChanged = true
        this.Repaint()
}


func (this *CanvasJPanel) KeyReleased(keyEvent IKeyEvent) {
}


func (this *CanvasJPanel) KeyTyped(keyEvent IKeyEvent) {
}


func (this *CanvasJPanel) Delete(evt *MyGraphicItemEvent) {
        this.GetGraphicItemHashMap().Remove(evt.GetTreeNode())
        this.WorkAreaJTreeJPanel.UpdateTree()
        this.WorkAreaJTreeJPanel.Expand()
        this.WorkAreaJTreeJPanel.Repaint()
}


func (this *CanvasJPanel) Deselect(evt *MyGraphicItemEvent) {
        var graphicItemInterface GraphicItemInterface = any(this.GetGraphicItemHashMap().Get(evt.GetTreeNode())).(GraphicItemInterface); _ = graphicItemInterface
        if graphicItemInterface != nil {
            graphicItemInterface.SetColor(Color_WHITE)
            graphicItemInterface.Deactivate()
            this.WorkAreaJTreeJPanel.Repaint()
        }
}


func (this *CanvasJPanel) Highlight(evt *MyGraphicItemEvent) {
        var graphicItemInterface GraphicItemInterface = any(this.GetGraphicItemHashMap().Get(evt.GetTreeNode())).(GraphicItemInterface); _ = graphicItemInterface
        if graphicItemInterface != nil {
            graphicItemInterface.SetColor(Color_YELLOW)
            graphicItemInterface.Activate()
            this.WorkAreaJTreeJPanel.Repaint()
        }
}


// throws Exception
func (this *CanvasJPanel) Duplicate_1() *CanvasJPanel {
        var newCanvasJPanel *CanvasJPanel = NewCanvasJPanel(this.WorkAreaJTreeJPanel, this.GetSize(), this.GetCanvasDimension().GetWidth(), this.GetCanvasDimension().GetHeight()); _ = newCanvasJPanel
        newCanvasJPanel.SetGrid(this.GetGrid().GetGrid())
        newCanvasJPanel.SetAngle(this.GetAngle())
        newCanvasJPanel.DuplicateGraphicItemHashMap(this.GetGraphicItemHashMap())
        this.WorkAreaJTreeJPanel.UpdateTree()
        this.WorkAreaJTreeJPanel.Expand()
        this.WorkAreaJTreeJPanel.Repaint()
        return newCanvasJPanel
}

// Java overloads of CanvasJPanel.duplicate
func (this *CanvasJPanel) Duplicate(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.Duplicate_1()
        }
        if len(args) == 1 && Java_IsArg[*MyGraphicItemEvent](args[0], pass) {
            return this.Duplicate_2(Java_Arg[*MyGraphicItemEvent](args[0]))
        }
    }
    panic("No Java overload of CanvasJPanel.duplicate for the arguments")
}


// throws Exception
func (this *CanvasJPanel) Duplicate_2(event *MyGraphicItemEvent) {
        var treeNode *MutableTreeNode = event.GetTreeNode(); _ = treeNode
        var graphicItem GraphicItemInterface = any(this.GetGraphicItemHashMap().Get(treeNode)).(GraphicItemInterface); _ = graphicItem
        if graphicItem != nil {
            this.DuplicateGraphicItem(graphicItem)
            this.WorkAreaJTreeJPanel.UpdateTree()
            this.WorkAreaJTreeJPanel.Expand()
            this.WorkAreaJTreeJPanel.Repaint()
        }
}


func (this *CanvasJPanel) Rotate(event *MyGraphicItemEvent) {
        var treeNode *MutableTreeNode = event.GetTreeNode(); _ = treeNode
        var graphicItem GraphicItemInterface = any(this.GetGraphicItemHashMap().Get(treeNode)).(GraphicItemInterface); _ = graphicItem
        if graphicItem != nil {
            graphicItem.AddAngle(event.GetAngle())
        }
}


func (this *CanvasJPanel) GetCanvasDimension() *IntegerDimension {
        return this.CanvasDimension
}


func (this *CanvasJPanel) SetCanvasDimension(canvasDimension *IntegerDimension) {
        this.CanvasDimension = canvasDimension
}


func (this *CanvasJPanel) GetGrid() *Grid {
        return this.Grid
}

