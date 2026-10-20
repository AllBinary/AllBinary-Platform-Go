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

type SVGAsJSONSimpleLoader struct {
    GraphicsStrings *GraphicsStrings
    ShapeTypeFactory *ShapeTypeFactory
    KEY_SVG string
    KEY_X1 string
    KEY_Y1 string
    KEY_X2 string
    KEY_Y2 string
    KEY_CX string
    KEY_CY string
    KEY_R string
    KEY_X string
    KEY_Y string
    ERROR_PARSE_SVG_JSON string
    ERROR_MISSING_KEY_PREFIX string
    ERROR_INVALID_NUMBER_PREFIX string
    ERROR_EXPECTED_OBJECT_PREFIX string
    EMPTY_POINTS [][]int
    SHAPE_SEPARATOR []int
    CIRCLE_SEGMENTS int
}

var SVGAsJSONSimpleLoader_Instance *SVGAsJSONSimpleLoader = NewSVGAsJSONSimpleLoader()

func (this *SVGAsJSONSimpleLoader) initFields() {
    this.GraphicsStrings = GraphicsStrings_GetInstance()
    this.ShapeTypeFactory = ShapeTypeFactory_GetInstance()
    this.KEY_SVG = "svg"
    this.KEY_X1 = "x1"
    this.KEY_Y1 = "y1"
    this.KEY_X2 = "x2"
    this.KEY_Y2 = "y2"
    this.KEY_CX = "cx"
    this.KEY_CY = "cy"
    this.KEY_R = "r"
    this.KEY_X = PositionStrings_GetInstance().X
    this.KEY_Y = PositionStrings_GetInstance().Y
    this.ERROR_PARSE_SVG_JSON = "Unable to parse SVG JSON"
    this.ERROR_MISSING_KEY_PREFIX = "Missing key: "
    this.ERROR_INVALID_NUMBER_PREFIX = "Invalid number for key: "
    this.ERROR_EXPECTED_OBJECT_PREFIX = "Expected object for: "
    this.EMPTY_POINTS = PrimitiveIntUtil_GetTwoDimensionalArrayInstance()
    this.SHAPE_SEPARATOR = []int{1000, 1000}
    this.CIRCLE_SEGMENTS = 16
}

func NewSVGAsJSONSimpleLoader() *SVGAsJSONSimpleLoader {
    this := &SVGAsJSONSimpleLoader{}
    this.initFields()
    return this
}


func SVGAsJSONSimpleLoader_GetInstance() *SVGAsJSONSimpleLoader {
        return SVGAsJSONSimpleLoader_Instance
}


// throws Exception
func (this *SVGAsJSONSimpleLoader) Load(svgAsJson string) [][]int {
        if svgAsJson == "" {
            return this.EMPTY_POINTS
        }
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 [][]int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(errors.New(fmt.Sprint(this.ERROR_PARSE_SVG_JSON, e)))
                }
            }()
            var root *JSONObject = NewJSONObject(svgAsJson); _ = root
            var svgValue any = root.Opt(this.KEY_SVG); _ = svgValue
            var svg *JSONObject = func() any { if svgValue == nil { return root }; return this.AsObject(svgValue, this.KEY_SVG) }(); _ = svg
            var pointVector IBasicArrayList = NewBasicArrayListD(); _ = pointVector
            var keys Enumeration = svg.Keys(); _ = keys
            for keys.HasMoreElements() {
                var key string = any(keys.NextElement()).(string); _ = key
                var value any = svg.Opt(key); _ = value
                if (this.ShapeTypeFactory.KEY_LINE == key) {
                    this.AddShapes(pointVector, value, this.ShapeTypeFactory.LINE)
                } else if Java_Equals(this.ShapeTypeFactory, key) {
                    this.AddShapes(pointVector, value, this.ShapeTypeFactory.CIRCLE)
                } else if (this.ShapeTypeFactory.KEY_RECT == key) || (this.ShapeTypeFactory.KEY_RECTANGLE == key) {
                    this.AddShapes(pointVector, value, this.ShapeTypeFactory.RECTANGLE)
                }
            }
            tryControl1, tryResult1 = 1, this.ToPoints(pointVector); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *SVGAsJSONSimpleLoader) AddShapes(pointVector IBasicArrayList, value any, shapeType *ShapeType) {
        if func() bool { _, ok := any(value).(*JSONArray); return ok }() {
            var array *JSONArray = any(value).(*JSONArray); _ = array
            var size int = array.Length(); _ = size
            for index := int(0); index < size; index++ {
                this.AddShape(pointVector, array.Opt(index), shapeType)
            }
        } else {
            this.AddShape(pointVector, value, shapeType)
        }
}


// throws Exception
func (this *SVGAsJSONSimpleLoader) AddShape(pointVector IBasicArrayList, value any, shapeType *ShapeType) {
        var shape *JSONObject = this.AsObject(value, shapeType.GetName()); _ = shape
        if shapeType == this.ShapeTypeFactory.LINE {
            this.AddLine(pointVector, shape)
        } else if shapeType == this.ShapeTypeFactory.CIRCLE {
            this.AddCircle(pointVector, shape)
        } else {
            this.AddRectangle(pointVector, shape)
        }
}


// throws Exception
func (this *SVGAsJSONSimpleLoader) AddLine(pointVector IBasicArrayList, line *JSONObject) {
        this.AddPoint(pointVector, this.ReadInt(line, this.KEY_X1), this.ReadInt(line, this.KEY_Y1))
        this.AddPoint(pointVector, this.ReadInt(line, this.KEY_X2), this.ReadInt(line, this.KEY_Y2))
        this.AddSeparator(pointVector)
}


// throws Exception
func (this *SVGAsJSONSimpleLoader) AddCircle(pointVector IBasicArrayList, circle *JSONObject) {
        var cx int = this.ReadInt(circle, this.KEY_CX); _ = cx
        var cy int = this.ReadInt(circle, this.KEY_CY); _ = cy
        var radius int = this.ReadInt(circle, this.KEY_R); _ = radius
        for index := int(0); index <= this.CIRCLE_SEGMENTS; index++ {
            var angle float64 = (Math_PI * 2.0 * index) / this.CIRCLE_SEGMENTS; _ = angle
            var x int = int(int64(math.Floor(float64(cx + (math.Cos(angle) * radius)) + 0.5))); _ = x
            var y int = int(int64(math.Floor(float64(cy + (math.Sin(angle) * radius)) + 0.5))); _ = y
            this.AddPoint(pointVector, x, y)
        }
        this.AddSeparator(pointVector)
}


// throws Exception
func (this *SVGAsJSONSimpleLoader) AddRectangle(pointVector IBasicArrayList, rectangle *JSONObject) {
        var x int = this.ReadInt(rectangle, this.KEY_X, 0); _ = x
        var y int = this.ReadInt(rectangle, this.KEY_Y, 0); _ = y
        var width int = this.ReadInt(rectangle, this.GraphicsStrings.WIDTH); _ = width
        var height int = this.ReadInt(rectangle, this.GraphicsStrings.HEIGHT); _ = height
        this.AddPoint(pointVector, x, y)
        this.AddPoint(pointVector, x + width, y)
        this.AddPoint(pointVector, x + width, y + height)
        this.AddPoint(pointVector, x, y + height)
        this.AddPoint(pointVector, x, y)
        this.AddSeparator(pointVector)
}


func (this *SVGAsJSONSimpleLoader) AddPoint(pointVector IBasicArrayList, x int, y int) {
        pointVector.Add([]int{x, y})
}


func (this *SVGAsJSONSimpleLoader) AddSeparator(pointVector IBasicArrayList) {
        pointVector.Add([]int{this.SHAPE_SEPARATOR[0], this.SHAPE_SEPARATOR[1]})
}


// throws Exception
func (this *SVGAsJSONSimpleLoader) ReadInt_1(anyType *JSONObject, key string) int {
        var value any = anyType.Get(key); _ = value
        if value == nil {
            panic(errors.New(fmt.Sprint(this.ERROR_MISSING_KEY_PREFIX + key)))
        }
        return this.ToInt(value, key)
}

// Java overloads of SVGAsJSONSimpleLoader.readInt
func (this *SVGAsJSONSimpleLoader) ReadInt(args ...any) int {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*JSONObject](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.ReadInt_1(Java_Arg[*JSONObject](args[0]), Java_Arg[string](args[1]))
        }
        if len(args) == 3 && Java_IsArg[*JSONObject](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[int](args[2], pass) {
            return this.ReadInt_2(Java_Arg[*JSONObject](args[0]), Java_Arg[string](args[1]), Java_Arg[int](args[2]))
        }
    }
    panic("No Java overload of SVGAsJSONSimpleLoader.readInt for the arguments")
}


func (this *SVGAsJSONSimpleLoader) ReadInt_2(anyType *JSONObject, key string, defaultValue int) int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
                }
            }()
            if anyType.Has(key) {
                var value any = anyType.Get(key); _ = value
                tryControl1, tryResult1 = 1, this.ToInt(value, key); return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return defaultValue
}


func (this *SVGAsJSONSimpleLoader) ToInt(value any, key string) int {
        if func() bool { _, ok := any(value).(*Number); return ok }() {
            return int(int64(math.Floor(float64(float64((any(value).(*Number)))) + 0.5)))
        }
        if func() bool { _, ok := any(value).(string); return ok }() {
            return int(int64(math.Floor(float64(func() float64 { value, _ := strconv.ParseFloat(any(value).(string), 64); return value }()) + 0.5)))
        }
        panic(errors.New(fmt.Sprint(this.ERROR_INVALID_NUMBER_PREFIX + key)))
}


func (this *SVGAsJSONSimpleLoader) AsObject(value any, name string) *JSONObject {
        if func() bool { _, ok := any(value).(*JSONObject); return ok }() {
            return any(value).(*JSONObject)
        }
        panic(errors.New(fmt.Sprint(this.ERROR_EXPECTED_OBJECT_PREFIX + name)))
}


func (this *SVGAsJSONSimpleLoader) ToPoints(pointVector IBasicArrayList) [][]int {
        if pointVector.Size() == 0 {
            return this.EMPTY_POINTS
        }
        var points [][]int = func() [][]int { array := make([][]int, pointVector.Size()); for index := range array { array[index] = make([]int, 2) }; return array }(); _ = points
        for index := int(0); index < pointVector.Size(); index++ {
            var point []int = any(pointVector.Get(index)).([]int); _ = point
            points[index] = point
        }
        return points
}


// throws Exception
func SVGAsJSONSimpleLoader_Main(args []string) {
        var svgAsJson string = "{\"svg\": {\n" + "    \"rect\": {\n" + "        \"width\": 20,\n" + "        \"fill\": \"#008c00\",\n" + "        \"height\": 20\n" + "    },\n" + "    \"xmlns\": \"http://www.w3.org/2000/svg\",\n" + "    \"width\": 20,\n" + "    \"height\": 20\n" + "}}\n"; _ = svgAsJson
        var points [][]int = SVGAsJSONSimpleLoader_GetInstance().Load(svgAsJson); _ = points
        var stringBuilder *StringMaker = NewStringMaker(); _ = stringBuilder
        var size int = len(points); _ = size
        for index := int(0); index < size; index++ {
            var size2 int = len(points[index]); _ = size2
            for index2 := int(0); index2 < size2; index2++ {
                stringBuilder.Appendint(int(points[index][index2]))
                stringBuilder.Append(CommonSeps_GetInstance().COMMA)
            }
        }
        fmt.Println(fmt.Sprint(stringBuilder))
}

