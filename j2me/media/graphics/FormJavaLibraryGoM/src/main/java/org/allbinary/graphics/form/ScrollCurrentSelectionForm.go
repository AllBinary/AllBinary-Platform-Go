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

type ScrollCurrentSelectionForm struct {
    ScrollSelectionForm
    MoveForSmallScreen bool
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    Processor IProcessor
    PreItemIndexDx IItemIndexDx
    Dx int
    Dy int
    MaxWidth int
}

// Implements UpdateMyFontInterface

// The Java class ScrollCurrentSelectionForm has subclasses so its variables are this interface
type IScrollCurrentSelectionForm interface {
    IScrollSelectionForm
    ScrollCurrentSelectionForm_() *ScrollCurrentSelectionForm
    Init(rectangle *Rectangle, formType *FormType)
    UpdateMeasurement(graphics IGraphics)
    ProcessTempHorizontalForm()
    ProcessHorizontalForm()
    ProcessVerticalForm()
    GetXTempHorizontalForm(index int, item IABCustomItem, dx int, dy int) int
    GetXHorizontalForm(index int, item IABCustomItem, dx int, dy int) int
    GetXVerticalForm(index int, item IABCustomItem, dx int, dy int) int
    GetDiffX(item IABCustomItem) int
    GetDx() int
    GetDy() int
    Paint(graphics IGraphics)
}

func (this *ScrollCurrentSelectionForm) ScrollCurrentSelectionForm_() *ScrollCurrentSelectionForm {
    return this
}


func (this *ScrollCurrentSelectionForm) initFields() {
    
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    this.Processor = Processor_GetInstance()
    this.PreItemIndexDx = ItemIndexDx_GetInstance()
    
    
    this.MaxWidth = 0
}


func NewScrollCurrentSelectionForm(title string, items []IABCustomItem, formPaintableFactory IItemPaintableFactory, border int, adjustedExtraBorder int, moveForSmallScreen bool, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *ScrollCurrentSelectionForm {
    this := &ScrollCurrentSelectionForm{}
        this.ScrollSelectionForm = *NewScrollSelectionForm(title, items, formPaintableFactory, border, adjustedExtraBorder, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
        this.MoveForSmallScreen = moveForSmallScreen
    return this
}


// throws Exception
func (this *ScrollCurrentSelectionForm) Init(rectangle *Rectangle, formType *FormType) {
        this.ScrollSelectionForm.Init(rectangle, formType)
        var formTypeFactory *FormTypeFactory = FormTypeFactory_GetInstance(); _ = formTypeFactory
        if formType == formTypeFactory.TEMP_HORIZONTAL_FORM {
            this.Processor = NewTempHorizontalFormProcessor(this)
            this.PreItemIndexDx = NewTempHorizontalItemIndexDx(this)
        } else if formType == formTypeFactory.HORIZONTAL_FORM {
            this.Processor = NewHorizontalFormProcessor(this)
            this.PreItemIndexDx = NewHorizontalItemIndexDx(this)
        } else if formType == formTypeFactory.VERTICAL_CENTER_FORM {
            this.Processor = NewVerticalFormProcessor(this)
            this.PreItemIndexDx = NewVerticalItemIndexDx(this)
        } else if formType == formTypeFactory.NULL_FORM_TYPE {
        } else {
            this.LogUtil.PutF(formTypeFactory.UNK, this, this.CommonStrings.INIT)
        }
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *ScrollCurrentSelectionForm) UpdateMeasurement(graphics IGraphics) {
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *ScrollCurrentSelectionForm) ProcessTempHorizontalForm() {
        this.Dx = this.X - 30 + (this.Rectangle.GetWidth() >> 1)
        this.Dy = this.Y
}


func (this *ScrollCurrentSelectionForm) ProcessHorizontalForm() {
        var size int = this.Size(); _ = size
        var totalWidth int = 0; _ = totalWidth
        var item IABCustomItem; _ = item
        for index := int(0); index < size; index++ {
            item = this.Get(index)
            totalWidth += item.GetMinimumWidth() + this.Border
        }
        this.Dx = this.X + (this.Rectangle.GetWidth() >> 1) - (totalWidth >> 1)
        this.Dy = this.Y + (this.Rectangle.GetHeight() >> 1)
        if this.MoveForSmallScreen {
            var maxTitleHeight int = 175; _ = maxTitleHeight
            if this.Dy < maxTitleHeight {
                this.Dy = maxTitleHeight
            }
        }
}


func (this *ScrollCurrentSelectionForm) ProcessVerticalForm() {
        var totalHeight int = 0; _ = totalHeight
        var size int = this.Size(); _ = size
        var item2 IABCustomItem; _ = item2
        for index := int(0); index < size; index++ {
            item2 = this.Get(index)
            if this.MaxWidth < item2.GetMinimumWidth() {
                this.MaxWidth = item2.GetMinimumWidth()
            }
            totalHeight += item2.GetMinimumHeight() + this.Border
        }
        this.Dx = ((this.Rectangle.GetWidth() - this.MaxWidth) / 2)
        if this.Size() > 0 {
            this.Dy = this.Y + ((this.Rectangle.GetHeight() - totalHeight) / 2)
            if this.MoveForSmallScreen {
                var maxTitleHeight int = 175; _ = maxTitleHeight
                if this.Dy < maxTitleHeight {
                    this.Dy = maxTitleHeight
                }
            }
        } else {
            this.Dy = this.Y
        }
}


func (this *ScrollCurrentSelectionForm) GetXTempHorizontalForm(index int, item IABCustomItem, dx int, dy int) int {
        return this.GetDiffX(item)
}


func (this *ScrollCurrentSelectionForm) GetXHorizontalForm(index int, item IABCustomItem, dx int, dy int) int {
        return 0
}


func (this *ScrollCurrentSelectionForm) GetXVerticalForm(index int, item IABCustomItem, dx int, dy int) int {
        return this.GetDiffX(item) + this.HalfBorder
}


func (this *ScrollCurrentSelectionForm) GetDiffX(item IABCustomItem) int {
        return ((this.MaxWidth - item.GetMinimumWidth()) >> 1)
}


func (this *ScrollCurrentSelectionForm) GetDx() int {
        return this.Dx
}


func (this *ScrollCurrentSelectionForm) GetDy() int {
        return this.Dy
}


func (this *ScrollCurrentSelectionForm) Paint(graphics IGraphics) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CanvasStrings.PAINT, e)
                }
            }()
            this.Processor.Process()
            var delta int = 0; _ = delta
            var deltaX int = this.GetDx(); _ = deltaX
            var deltaY int = this.GetDy(); _ = deltaY
            var size int = this.Size(); _ = size
            var formTypeFactory *FormTypeFactory = FormTypeFactory_GetInstance(); _ = formTypeFactory
            var item IABCustomItem; _ = item
            for index := int(0); index < size; index++ {
                item = this.Get(index)
                var diffX int = this.PreItemIndexDx.GetDx(index, item, this.Dx, this.Dy); _ = diffX
                if index == this.GetSelectedIndex() {
                    delta = this.PaintItem(graphics, index, item, deltaX + diffX, deltaY)
                    this.Paintable.Paint(graphics, index, deltaX + diffX, deltaY)
                } else {
                    delta = this.PaintUnselectedItem(graphics, index, item, deltaX + diffX, deltaY)
                    this.Paintable.Paint(graphics, index, deltaX + diffX, deltaY)
                }
                if this.FormType == formTypeFactory.TEMP_HORIZONTAL_FORM {
                } else if this.FormType == formTypeFactory.HORIZONTAL_FORM {
                    deltaX = delta
                } else if this.FormType == formTypeFactory.VERTICAL_CENTER_FORM {
                    deltaY = delta
                } else {
                    panic(errors.New(fmt.Sprint(formTypeFactory.UNK)))
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

type TempHorizontalFormProcessor struct {
    Processor
    ScrollCurrentSelectionForm IScrollCurrentSelectionForm
}


func (this *TempHorizontalFormProcessor) initFields() {
        
}
    

    func NewScrollCurrentSelectionForm(scrollCurrentSelectionForm IScrollCurrentSelectionForm) *ScrollCurrentSelectionForm {
        this := &ScrollCurrentSelectionForm{}
        this.Processor = *NewProcessor()
        this.initFields()
            this.ScrollCurrentSelectionForm = scrollCurrentSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollCurrentSelectionForm) Process() {
            this.ScrollCurrentSelectionForm.ProcessTempHorizontalForm()
            this.ScrollCurrentSelectionForm.Processor = Processor_GetInstance()
    }

type HorizontalFormProcessor struct {
    Processor
    ScrollCurrentSelectionForm IScrollCurrentSelectionForm
}


func (this *HorizontalFormProcessor) initFields() {
        
}
    

    func NewScrollCurrentSelectionForm(scrollCurrentSelectionForm IScrollCurrentSelectionForm) *ScrollCurrentSelectionForm {
        this := &ScrollCurrentSelectionForm{}
        this.Processor = *NewProcessor()
        this.initFields()
            this.ScrollCurrentSelectionForm = scrollCurrentSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollCurrentSelectionForm) Process() {
            this.ScrollCurrentSelectionForm.ProcessHorizontalForm()
            this.ScrollCurrentSelectionForm.Processor = Processor_GetInstance()
    }

type VerticalFormProcessor struct {
    Processor
    ScrollCurrentSelectionForm IScrollCurrentSelectionForm
}


func (this *VerticalFormProcessor) initFields() {
        
}
    

    func NewScrollCurrentSelectionForm(scrollCurrentSelectionForm IScrollCurrentSelectionForm) *ScrollCurrentSelectionForm {
        this := &ScrollCurrentSelectionForm{}
        this.Processor = *NewProcessor()
        this.initFields()
            this.ScrollCurrentSelectionForm = scrollCurrentSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollCurrentSelectionForm) Process() {
            this.ScrollCurrentSelectionForm.ProcessVerticalForm()
            this.ScrollCurrentSelectionForm.Processor = Processor_GetInstance()
    }

type TempHorizontalItemIndexDx struct {
    ItemIndexDx
    ScrollCurrentSelectionForm IScrollCurrentSelectionForm
}


func (this *TempHorizontalItemIndexDx) initFields() {
        
}
    

    func NewScrollCurrentSelectionForm(scrollCurrentSelectionForm IScrollCurrentSelectionForm) *ScrollCurrentSelectionForm {
        this := &ScrollCurrentSelectionForm{}
        this.ItemIndexDx = *NewItemIndexDx()
        this.initFields()
            this.ScrollCurrentSelectionForm = scrollCurrentSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollCurrentSelectionForm) GetDx(index int, item IABCustomItem, dx int, dy int) int {
            return this.ScrollCurrentSelectionForm.GetXTempHorizontalForm(index, item, dx, dy)
    }

type HorizontalItemIndexDx struct {
    ItemIndexDx
    ScrollCurrentSelectionForm IScrollCurrentSelectionForm
}


func (this *HorizontalItemIndexDx) initFields() {
        
}
    

    func NewScrollCurrentSelectionForm(scrollCurrentSelectionForm IScrollCurrentSelectionForm) *ScrollCurrentSelectionForm {
        this := &ScrollCurrentSelectionForm{}
        this.ItemIndexDx = *NewItemIndexDx()
        this.initFields()
            this.ScrollCurrentSelectionForm = scrollCurrentSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollCurrentSelectionForm) GetDx(index int, item IABCustomItem, dx int, dy int) int {
            return this.ScrollCurrentSelectionForm.GetXHorizontalForm(index, item, dx, dy)
    }

type VerticalItemIndexDx struct {
    ItemIndexDx
    ScrollCurrentSelectionForm IScrollCurrentSelectionForm
}


func (this *VerticalItemIndexDx) initFields() {
        
}
    

    func NewScrollCurrentSelectionForm(scrollCurrentSelectionForm IScrollCurrentSelectionForm) *ScrollCurrentSelectionForm {
        this := &ScrollCurrentSelectionForm{}
        this.ItemIndexDx = *NewItemIndexDx()
        this.initFields()
            this.ScrollCurrentSelectionForm = scrollCurrentSelectionForm
        return this
    }
    

    // throws Exception
    func (this *ScrollCurrentSelectionForm) GetDx(index int, item IABCustomItem, dx int, dy int) int {
            return this.ScrollCurrentSelectionForm.GetXVerticalForm(index, item, dx, dy)
    }

