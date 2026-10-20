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

type ScrollSelectionForm struct {
    PaintableForm
    RectangleCollisionUtil *RectangleCollisionUtil
    Border int
    HalfBorder int
    ButtonBasicColor *BasicColor
    ScrollSelectionFormFormTypeItemIndexPaintable IItemIndexPaintable
    FormTypeItemIndexDx IItemIndexDx
    Paintable IItemPaintable
    AdjustedExtraBorder int
}

// The Java class ScrollSelectionForm has subclasses so its variables are this interface
type IScrollSelectionForm interface {
    IPaintableForm
    ScrollSelectionForm_() *ScrollSelectionForm
    Init(rectangle *Rectangle, formType *FormType)
    PaintItemHorizontal(graphics IGraphics, index int, item IABCustomItem, x int, y int) int
    PaintItemVertical(graphics IGraphics, index int, item IABCustomItem, x int, y int) int
    PaintItemTempHorizontal(graphics IGraphics, index int, item IABCustomItem, x int, y int) int
    GetSelectedIndexForPointHorizontalDx(index int, item IABCustomItem, dx int, dy int) int
    GetSelectedIndexForPointVerticalDx(index int, item IABCustomItem, dx int, dy int) int
    GetSelectedIndexForPointTempHorizontalDx(index int, item IABCustomItem, dx int, dy int) int
    Append(item IABCustomItem) int
    GetSelectedItem(point IGPoint) IABCustomItem
    GetItemIndex(item IABCustomItem) int
    GetStartIndex() int
    GetSelectedIndexForPoint(point IGPoint) int
    ProcessInputKey(gameKeyCode int) int
    IsInForm(point IGPoint) bool
    PaintItem(graphics IGraphics, index int, item IABCustomItem, x int, y int) int
    PaintUnselectedItem(graphics IGraphics, index int, item IABCustomItem, x int, y int) int
    GetDiffX(item IABCustomItem) int
    GetDx() int
    GetDy() int
    SetButtonBasicColor(buttonBasicColor *BasicColor)
    GetButtonBasicColor() *BasicColor
}

func (this *ScrollSelectionForm) ScrollSelectionForm_() *ScrollSelectionForm {
    return this
}

var ScrollSelectionForm_NULL_SCROLL_SELECTION_FORM any = NullUtil_GetInstance().NULL_OBJECT
var ScrollSelectionForm_NULL_SCROLL_SELECTION_HORIZONTAL_FORM any = NullUtil_GetInstance().NULL_OBJECT
var ScrollSelectionForm_INSIDE_FORM string = " inside form"
var ScrollSelectionForm_IS_IN_FORM string = "isInForm"
var ScrollSelectionForm_GET_SELECTED_INDEX string = "getSelectedIndexForPoint"

func (this *ScrollSelectionForm) initFields() {
    this.RectangleCollisionUtil = RectangleCollisionUtil_GetInstance()
    
    
    
    this.ScrollSelectionFormFormTypeItemIndexPaintable = ItemIndexPaintable_GetInstance()
    this.FormTypeItemIndexDx = ItemIndexDx_GetInstance()
    this.Paintable = ItemPaintableFactory_GetInstance()
    
}


func ScrollSelectionForm_CreateForm(title string, items []IABCustomItem, formPaintableFactory IItemPaintableFactory, border int, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) IPaintableForm {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IPaintableForm) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            logUtil.Put(commonStrings.EXCEPTION, logUtil, commonStrings.EXCEPTION, e)
            tryControl1, tryResult1 = 1, PaintableForm_GetNullPaintableForm(); return
                }
            }()
            if formPaintableFactory == ItemPaintableFactory_GetInstance() {
                tryControl1, tryResult1 = 1, NewScrollSelectionForm(title, items, formPaintableFactory, border, -3, backgroundBasicColor, foregroundBasicColor); return
            } else {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                logUtil.PutF(commonStrings.EXCEPTION, logUtil, commonStrings.EXCEPTION)
                tryControl1, tryResult1 = 1, PaintableForm_GetNullPaintableForm(); return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func ScrollSelectionForm_GetNullScrollSelectionForm() IScrollSelectionForm {
        if ScrollSelectionForm_NULL_SCROLL_SELECTION_FORM == NullUtil_GetInstance().NULL_OBJECT {
            ScrollSelectionForm_NULL_SCROLL_SELECTION_FORM = ScrollSelectionForm_CreateForm(StringUtil_GetInstance().EMPTY_STRING, make([]IABCustomItem, 0), ItemPaintableFactory_GetInstance(), 0, BasicColorFactory_GetInstance().BLACK, BasicColorFactory_GetInstance().WHITE)
        }
        return any(ScrollSelectionForm_NULL_SCROLL_SELECTION_FORM).(IScrollSelectionForm)
}


func ScrollSelectionForm_GetNullScrollSelectionFormHorizontal() IScrollSelectionForm {
        if ScrollSelectionForm_NULL_SCROLL_SELECTION_HORIZONTAL_FORM == NullUtil_GetInstance().NULL_OBJECT {
            ScrollSelectionForm_NULL_SCROLL_SELECTION_HORIZONTAL_FORM = ScrollSelectionForm_CreateForm(StringUtil_GetInstance().EMPTY_STRING, make([]IABCustomItem, 0), ItemPaintableFactory_GetInstance(), 0, BasicColorFactory_GetInstance().BLACK, BasicColorFactory_GetInstance().WHITE)
        }
        return any(ScrollSelectionForm_NULL_SCROLL_SELECTION_HORIZONTAL_FORM).(IScrollSelectionForm)
}


func NewScrollSelectionForm(title string, items []IABCustomItem, formPaintableFactory IItemPaintableFactory, border int, adjustedExtraBorder int, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *ScrollSelectionForm {
    this := &ScrollSelectionForm{}
        this.PaintableForm = *NewPaintableForm(title, items, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
        this.AdjustedExtraBorder = adjustedExtraBorder
        this.ButtonBasicColor = foregroundBasicColor
        this.Border = border
        this.HalfBorder = (border >> 1)
        this.Paintable = formPaintableFactory.GetInstanceItemPaintable(this)
    return this
}


// throws Exception
func (this *ScrollSelectionForm) Init(rectangle *Rectangle, formType *FormType) {
        this.PaintableForm.Init(rectangle, formType)
        var formTypeFactory *FormTypeFactory = FormTypeFactory_GetInstance(); _ = formTypeFactory
        if formType == formTypeFactory.HORIZONTAL_FORM {
            this.ScrollSelectionFormFormTypeItemIndexPaintable = NewScrollSelectionFormHorizontalPaintable(this)
            this.FormTypeItemIndexDx = NewScrollSelectionFormHorizontalDx(this)
        } else if formType == formTypeFactory.VERTICAL_CENTER_FORM {
            this.ScrollSelectionFormFormTypeItemIndexPaintable = NewScrollSelectionFormVerticalPaintable(this)
            this.FormTypeItemIndexDx = NewScrollSelectionFormVericalDx(this)
        } else if formType == formTypeFactory.TEMP_HORIZONTAL_FORM {
            this.ScrollSelectionFormFormTypeItemIndexPaintable = NewScrollSelectionFormTempHorizontalPaintable(this)
            this.FormTypeItemIndexDx = NewScrollSelectionFormTempHorizontalDx(this)
        } else if formType == formTypeFactory.NULL_FORM_TYPE {
        } else {
            panic(errors.New(fmt.Sprint(formTypeFactory.UNK)))
        }
}


func (this *ScrollSelectionForm) PaintItemHorizontal(graphics IGraphics, index int, item IABCustomItem, x int, y int) int {
        var width int = item.GetMinimumWidth(); _ = width
        return x + width + this.Border
}


func (this *ScrollSelectionForm) PaintItemVertical(graphics IGraphics, index int, item IABCustomItem, x int, y int) int {
        var height int = item.GetMinimumHeight(); _ = height
        return y + height + this.Border
}


func (this *ScrollSelectionForm) PaintItemTempHorizontal(graphics IGraphics, index int, item IABCustomItem, x int, y int) int {
        return 0
}


func (this *ScrollSelectionForm) GetSelectedIndexForPointHorizontalDx(index int, item IABCustomItem, dx int, dy int) int {
        return dx - this.HalfBorder
}


func (this *ScrollSelectionForm) GetSelectedIndexForPointVerticalDx(index int, item IABCustomItem, dx int, dy int) int {
        return dx + this.GetDiffX(item)
}


func (this *ScrollSelectionForm) GetSelectedIndexForPointTempHorizontalDx(index int, item IABCustomItem, dx int, dy int) int {
        return dx + this.GetDiffX(item)
}


func (this *ScrollSelectionForm) Append(item IABCustomItem) int {
        return this.PaintableForm.Append(item)
}


// throws Exception
func (this *ScrollSelectionForm) GetSelectedItem(point IGPoint) IABCustomItem {
        var index int = this.GetSelectedIndexForPoint(point); _ = index
        if index != -1 {
            return this.Get(index)
        } else {
            return ABCustomItem_GetNullInstance()
        }
}


// throws Exception
func (this *ScrollSelectionForm) GetItemIndex(item IABCustomItem) int {
        var nextItem IABCustomItem; _ = nextItem
        var size int = this.Size(); _ = size
        for index := int(0); index < size; index++ {
            nextItem = any(this.Get(index)).(IABCustomItem)
            if nextItem == item {
                return index
            }
        }
        return -1
}


func (this *ScrollSelectionForm) GetStartIndex() int {
        return 0
}


// throws Exception
func (this *ScrollSelectionForm) GetSelectedIndexForPoint(point IGPoint) int {
        var start int = this.GetStartIndex(); _ = start
        var size int = this.Size(); _ = size
        var dx int = this.GetDx(); _ = dx
        var dy int = this.GetDy(); _ = dy
        var formTypeFactory *FormTypeFactory = FormTypeFactory_GetInstance(); _ = formTypeFactory
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(commonLabels.START_LABEL)
        stringBuffer.Appendint(start)
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(commonLabels.TOTAL_LABEL)
        stringBuffer.Appendint(size)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, ScrollSelectionForm_GET_SELECTED_INDEX)
        var item IABCustomItem; _ = item
        var width int; _ = width
        var height int; _ = height
        var diffX int; _ = diffX
        for index := int(start); index < size; index++ {
            item = this.Get(index)
            width = item.GetMinimumWidth()
            height = item.GetMinimumHeight()
            diffX = this.FormTypeItemIndexDx.GetDx(index, item, dx, dy)
            if this.RectangleCollisionUtil.IsInside(diffX, dy - this.HalfBorder, diffX + width + this.Border, dy + height + this.HalfBorder + 1, point.GetX(), point.GetY()) {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(commonLabels.ITEM_LABEL)
                stringBuffer.Append(item.GetLabel())
                stringBuffer.Append(CommonSeps_GetInstance().SPACE)
                stringBuffer.Append(commonLabels.INDEX_LABEL)
                stringBuffer.Appendint(index)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, ScrollSelectionForm_GET_SELECTED_INDEX)
                return index
            }
            if this.FormType == formTypeFactory.HORIZONTAL_FORM {
                dx = dx + width + this.Border
                if dx > this.Rectangle.GetMaxX() {
                    break
                }
            } else if this.FormType == formTypeFactory.VERTICAL_CENTER_FORM {
                dy = dy + height + this.Border
                if dy > this.Rectangle.GetMaxY() {
                    break
                }
            } else if this.FormType == formTypeFactory.NULL_FORM_TYPE {
            } else {
                panic(errors.New(fmt.Sprint(formTypeFactory.UNK)))
            }
        }
        return -1
}


// throws Exception
func (this *ScrollSelectionForm) ProcessInputKey(gameKeyCode int) int {
        var formTypeFactory *FormTypeFactory = FormTypeFactory_GetInstance(); _ = formTypeFactory
        if this.Size() == 0 {
            return -1
        }
        var index int = this.GetSelectedIndex(); _ = index
        if this.FormType == formTypeFactory.HORIZONTAL_FORM || this.FormType == formTypeFactory.TEMP_HORIZONTAL_FORM {
            if gameKeyCode == Canvas_RIGHT {
                index++
            } else if gameKeyCode == Canvas_LEFT {
                index--
            }
        } else if this.FormType == formTypeFactory.VERTICAL_CENTER_FORM {
            if gameKeyCode == Canvas_DOWN {
                index++
            } else if gameKeyCode == Canvas_UP {
                index--
            }
        } else if this.FormType == formTypeFactory.NULL_FORM_TYPE {
        } else {
            panic(errors.New(fmt.Sprint(formTypeFactory.UNK)))
        }
        var max_ int = this.Size() - 1; _ = max_
        if index < 0 {
            this.SetSelectedIndex(0)
        } else if index > max_ {
            this.SetSelectedIndex(max_)
        } else {
            this.SetSelectedIndex(index)
        }
        return -1
}


func (this *ScrollSelectionForm) IsInForm(point IGPoint) bool {
        if this.RectangleCollisionUtil.IsInside(this.X, this.Y - this.HalfBorder, this.Rectangle.GetMaxX() + this.Border, this.Rectangle.GetMaxY() + this.Border, point.GetX(), point.GetY()) {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(StringUtil_GetInstance().ToString(point)).Append(ScrollSelectionForm_INSIDE_FORM)), this, ScrollSelectionForm_IS_IN_FORM)
            return true
        }
        return false
}


// throws Exception
func (this *ScrollSelectionForm) PaintItem(graphics IGraphics, index int, item IABCustomItem, x int, y int) int {
        item.PaintXY(graphics, x, y)
        var width int = item.GetMinimumWidth(); _ = width
        var height int = item.GetMinimumHeight(); _ = height
        graphics.SetColor(this.GetButtonBasicColor().IntValue())
        var adjustedBorder int = 3; _ = adjustedBorder
        graphics.DrawRect(x - this.HalfBorder - adjustedBorder, y - this.HalfBorder - adjustedBorder, width + this.Border + this.AdjustedExtraBorder, height + this.Border + this.AdjustedExtraBorder)
        return this.ScrollSelectionFormFormTypeItemIndexPaintable.Paint(graphics, index, item, x, y)
}


// throws Exception
func (this *ScrollSelectionForm) PaintUnselectedItem(graphics IGraphics, index int, item IABCustomItem, x int, y int) int {
        graphics.SetColor(this.GetButtonBasicColor().IntValue())
        item.PaintUnselected(graphics, x, y)
        return this.ScrollSelectionFormFormTypeItemIndexPaintable.Paint(graphics, index, item, x, y)
}


func (this *ScrollSelectionForm) GetDiffX(item IABCustomItem) int {
        return 0
}


func (this *ScrollSelectionForm) GetDx() int {
        return this.X
}


func (this *ScrollSelectionForm) GetDy() int {
        return this.Y
}


func (this *ScrollSelectionForm) SetButtonBasicColor(buttonBasicColor *BasicColor) {
        this.ButtonBasicColor = buttonBasicColor
}


func (this *ScrollSelectionForm) GetButtonBasicColor() *BasicColor {
        return this.ButtonBasicColor
}

type ScrollSelectionFormHorizontalPaintable struct {
    ItemIndexPaintable
    ScrollSelectionForm IScrollSelectionForm
}


func (this *ScrollSelectionFormHorizontalPaintable) initFields() {
        
}
    

    func NewScrollSelectionForm(scrollSelectionForm IScrollSelectionForm) *ScrollSelectionForm {
        this := &ScrollSelectionForm{}
        this.ItemIndexPaintable = *NewItemIndexPaintable()
        this.initFields()
            this.ScrollSelectionForm = scrollSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollSelectionForm) Paint(graphics IGraphics, index int, item IABCustomItem, dx int, dy int) int {
            return this.ScrollSelectionForm.PaintItemHorizontal(graphics, index, item, dx, dy)
    }

type ScrollSelectionFormVerticalPaintable struct {
    ItemIndexPaintable
    ScrollSelectionForm IScrollSelectionForm
}


func (this *ScrollSelectionFormVerticalPaintable) initFields() {
        
}
    

    func NewScrollSelectionForm(scrollSelectionForm IScrollSelectionForm) *ScrollSelectionForm {
        this := &ScrollSelectionForm{}
        this.ItemIndexPaintable = *NewItemIndexPaintable()
        this.initFields()
            this.ScrollSelectionForm = scrollSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollSelectionForm) Paint(graphics IGraphics, index int, item IABCustomItem, dx int, dy int) int {
            return this.ScrollSelectionForm.PaintItemVertical(graphics, index, item, dx, dy)
    }

type ScrollSelectionFormTempHorizontalPaintable struct {
    ItemIndexPaintable
    ScrollSelectionForm IScrollSelectionForm
}


func (this *ScrollSelectionFormTempHorizontalPaintable) initFields() {
        
}
    

    func NewScrollSelectionForm(scrollSelectionForm IScrollSelectionForm) *ScrollSelectionForm {
        this := &ScrollSelectionForm{}
        this.ItemIndexPaintable = *NewItemIndexPaintable()
        this.initFields()
            this.ScrollSelectionForm = scrollSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollSelectionForm) Paint(graphics IGraphics, index int, item IABCustomItem, dx int, dy int) int {
            return this.ScrollSelectionForm.PaintItemTempHorizontal(graphics, index, item, dx, dy)
    }

type ScrollSelectionFormHorizontalDx struct {
    ItemIndexDx
    ScrollSelectionForm IScrollSelectionForm
}


func (this *ScrollSelectionFormHorizontalDx) initFields() {
        
}
    

    func NewScrollSelectionForm(multipleScrollSelectionForm IScrollSelectionForm) *ScrollSelectionForm {
        this := &ScrollSelectionForm{}
        this.ItemIndexDx = *NewItemIndexDx()
        this.initFields()
            this.ScrollSelectionForm = multipleScrollSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollSelectionForm) GetDx(index int, item IABCustomItem, dx int, dy int) int {
            return this.ScrollSelectionForm.GetSelectedIndexForPointHorizontalDx(index, item, dx, dy)
    }

type ScrollSelectionFormVericalDx struct {
    ItemIndexDx
    ScrollSelectionForm IScrollSelectionForm
}


func (this *ScrollSelectionFormVericalDx) initFields() {
        
}
    

    func NewScrollSelectionForm(multipleScrollSelectionForm IScrollSelectionForm) *ScrollSelectionForm {
        this := &ScrollSelectionForm{}
        this.ItemIndexDx = *NewItemIndexDx()
        this.initFields()
            this.ScrollSelectionForm = multipleScrollSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollSelectionForm) GetDx(index int, item IABCustomItem, dx int, dy int) int {
            return this.ScrollSelectionForm.GetSelectedIndexForPointVerticalDx(index, item, dx, dy)
    }

type ScrollSelectionFormTempHorizontalDx struct {
    ItemIndexDx
    ScrollSelectionForm IScrollSelectionForm
}


func (this *ScrollSelectionFormTempHorizontalDx) initFields() {
        
}
    

    func NewScrollSelectionForm(multipleScrollSelectionForm IScrollSelectionForm) *ScrollSelectionForm {
        this := &ScrollSelectionForm{}
        this.ItemIndexDx = *NewItemIndexDx()
        this.initFields()
            this.ScrollSelectionForm = multipleScrollSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollSelectionForm) GetDx(index int, item IABCustomItem, dx int, dy int) int {
            return this.ScrollSelectionForm.GetSelectedIndexForPointTempHorizontalDx(index, item, dx, dy)
    }

