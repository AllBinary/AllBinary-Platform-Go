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

type MultipleScrollSelectionHorizontalForm struct {
    ScrollSelectionForm
    RectPaintable IPaintable
    MultipleScrollSelectionHorizontalFormTypeItemIndexPaintable IItemIndexPaintable
    BackgroundColor int
    MyFontProcessor IMyFontProcessor
    Logged bool
    FontHeight int
}

// Implements UpdateMyFontInterface


func (this *MultipleScrollSelectionHorizontalForm) initFields() {
    this.RectPaintable = NullPaintable_GetInstance()
    this.MultipleScrollSelectionHorizontalFormTypeItemIndexPaintable = ItemIndexPaintable_GetInstance()
    this.BackgroundColor = BasicColorFactory_GetInstance().TRANSPARENT_GREY.IntValue()
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.Logged = false
    this.FontHeight = 0
}


// throws Exception
func MultipleScrollSelectionHorizontalForm_CreateForm(title string, items []IABCustomItem, rectangle *Rectangle, formType *FormType, border int, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) IScrollSelectionForm {
        var scrollSelectionForm IScrollSelectionForm = NewMultipleScrollSelectionHorizontalForm(title, items, ItemPaintableFactory_GetInstance(), border, backgroundBasicColor, foregroundBasicColor); _ = scrollSelectionForm
        scrollSelectionForm.Init(rectangle, formType)
        return scrollSelectionForm
}


func NewMultipleScrollSelectionHorizontalForm(title string, items []IABCustomItem, formPaintableFactory IItemPaintableFactory, border int, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *MultipleScrollSelectionHorizontalForm {
    this := &MultipleScrollSelectionHorizontalForm{}
        this.ScrollSelectionForm = *NewScrollSelectionForm(title, items, formPaintableFactory, border, 3, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
        if J2MEUtil_IsJ2ME() || J2MEUtil_IsHTML() {
        } else {
            /* Java local class: move it out of the function
type MPaintable struct {
    Paintable
}


func NewMPaintable() *MPaintable {
    this := &MPaintable{}
    this.Paintable = *NewPaintable()
    return this
}
            

            func (this *MultipleScrollSelectionHorizontalForm) Paint(graphics IGraphics) {
                    this.FillRect(graphics)
            }
*/
            this.RectPaintable = NewMPaintable()
        }
    return this
}


func (this *MultipleScrollSelectionHorizontalForm) FillRect(graphics IGraphics) {
        graphics.SetColor(this.BackgroundColor)
        graphics.FillRect(this.X, this.Y, this.Rectangle.GetWidth(), this.Rectangle.GetHeight())
}


// throws Exception
func (this *MultipleScrollSelectionHorizontalForm) Init(rectangle *Rectangle, formType *FormType) {
        this.ScrollSelectionForm.Init(rectangle, formType)
        var formTypeFactory *FormTypeFactory = FormTypeFactory_GetInstance(); _ = formTypeFactory
        if formType == formTypeFactory.HORIZONTAL_FORM {
            this.MultipleScrollSelectionHorizontalFormTypeItemIndexPaintable = NewMultipleScrollSelectionFormHorizontalPaintable(this)
        } else if formType == formTypeFactory.VERTICAL_CENTER_FORM {
            this.MultipleScrollSelectionHorizontalFormTypeItemIndexPaintable = NewMultipleScrollSelectionFormVerticalPaintable(this)
        } else if formType == formTypeFactory.NULL_FORM_TYPE {
        } else {
            panic(errors.New(fmt.Sprint(formTypeFactory.UNK)))
        }
}


func (this *MultipleScrollSelectionHorizontalForm) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.FontHeight = font.GetHeight()
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


// throws Exception
func (this *MultipleScrollSelectionHorizontalForm) PaintHorizontal(graphics IGraphics, index int, item IABCustomItem, dx int, dy int) int {
        var dx2 int = this.PaintItem(graphics, index, item, dx, dy) + this.Border; _ = dx2
        this.Paintable.Paint(graphics, index, dx, dy)
        dx = dx2
        return dx
}


// throws Exception
func (this *MultipleScrollSelectionHorizontalForm) PaintVertical(graphics IGraphics, index int, item IABCustomItem, dx int, dy int) int {
        this.PaintItem(graphics, index, item, dx, dy)
        this.Paintable.Paint(graphics, index, dx, dy)
        return dx
}


func (this *MultipleScrollSelectionHorizontalForm) GetStartIndex() int {
        return this.GetSelectedIndex()
}


func (this *MultipleScrollSelectionHorizontalForm) Paint(graphics IGraphics) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_Put(this.CommonStrings.EXCEPTION, this, this.CanvasStrings.PAINT)
                }
            }()
            this.MyFontProcessor.Process(graphics)
            var start int = this.GetStartIndex(); _ = start
            var size int = this.Size(); _ = size
            var dx int = this.X; _ = dx
            var dy int = this.Y; _ = dy
            this.RectPaintable.Paint(graphics)
            graphics.DrawString(this.GetTitle(), this.X, this.Y - this.FontHeight, 0)
            var item IABCustomItem; _ = item
            for index := int(start); index < size; index++ {
                item = this.Get(index)
                if dx >= this.Rectangle.GetMaxX() {
                    if this.Logged {
                    } else {
                        this.Logged = true
                        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("painting beyond maxx: ").Appendint(this.Rectangle.GetMaxX())), this, this.CanvasStrings.PAINT)
                    }
                }
                dx = this.MultipleScrollSelectionHorizontalFormTypeItemIndexPaintable.Paint(graphics, index, item, dx, dy)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

type MultipleScrollSelectionFormHorizontalPaintable struct {
    ItemIndexPaintable
    MultipleScrollSelectionForm *MultipleScrollSelectionHorizontalForm
}


func (this *MultipleScrollSelectionFormHorizontalPaintable) initFields() {
        
}
    

    func NewMultipleScrollSelectionHorizontalForm(multipleScrollSelectionForm *MultipleScrollSelectionHorizontalForm) *MultipleScrollSelectionHorizontalForm {
        this := &MultipleScrollSelectionHorizontalForm{}
        this.ItemIndexPaintable = *NewItemIndexPaintable()
        this.initFields()
            this.MultipleScrollSelectionForm = multipleScrollSelectionForm
        return this
    }
    

    // throws Exception
    func (this *MultipleScrollSelectionHorizontalForm) Paint(graphics IGraphics, index int, item IABCustomItem, dx int, dy int) int {
            return this.MultipleScrollSelectionForm.PaintHorizontal(graphics, index, item, dx, dy)
    }

type MultipleScrollSelectionFormVerticalPaintable struct {
    ItemIndexPaintable
    MultipleScrollSelectionForm *MultipleScrollSelectionHorizontalForm
}


func (this *MultipleScrollSelectionFormVerticalPaintable) initFields() {
        
}
    

    func NewMultipleScrollSelectionHorizontalForm(multipleScrollSelectionForm *MultipleScrollSelectionHorizontalForm) *MultipleScrollSelectionHorizontalForm {
        this := &MultipleScrollSelectionHorizontalForm{}
        this.ItemIndexPaintable = *NewItemIndexPaintable()
        this.initFields()
            this.MultipleScrollSelectionForm = multipleScrollSelectionForm
        return this
    }
    

    // throws Exception
    func (this *MultipleScrollSelectionHorizontalForm) Paint(graphics IGraphics, index int, item IABCustomItem, dx int, dy int) int {
            return this.MultipleScrollSelectionForm.PaintVertical(graphics, index, item, dx, dy)
    }

