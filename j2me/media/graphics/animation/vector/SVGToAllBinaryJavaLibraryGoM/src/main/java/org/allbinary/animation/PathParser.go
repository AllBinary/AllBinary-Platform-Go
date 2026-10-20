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

type PathParser struct {
}


func NewPathParser() *PathParser {
    this := &PathParser{}
    return this
}


func PathParser_Tokenize(d string) []string {
        var tokens IBasicArrayList = NewBasicArrayListD(); _ = tokens
        var sb *strings.Builder = &strings.Builder{}; _ = sb
        for i := int(0); i < len(d); i++ {
            var c rune = rune(d[i]); _ = c
            if c == ' ' || c == ',' || c == '\t' || c == '\n' || c == '\r' {
                if sb.Len() > 0 {
                    tokens.Add(sb.String())
                    sb.SetLength(0)
                }
            } else if c == 'M' || c == 'm' || c == 'L' || c == 'l' || c == 'V' || c == 'v' || c == 'C' || c == 'c' || c == 'Z' || c == 'z' {
                if sb.Len() > 0 {
                    tokens.Add(sb.String())
                    sb.SetLength(0)
                }
                tokens.Add(fmt.Sprint(c))
            } else {
                sb.Append(c)
            }
        }
        if sb.Len() > 0 {
            tokens.Add(sb.String())
        }
        var res []string = make([]string, tokens.Size()); _ = res
        tokens.ToArrayType(res)
        return res
}


func PathParser_ParsePath(d string, outX []int, outY []int, maxPoints int) int {
        var tokens []string = PathParser_Tokenize(d); _ = tokens
        var count []int = []int{0}; _ = count
        var fixedPoint *FixedPoint = FixedPoint_GetInstance(); _ = fixedPoint
        var curX int = 0; _ = curX
        var curY int = 0; _ = curY
        var startX int = 0; _ = startX
        var startY int = 0; _ = startY
        var idx int = 0; _ = idx
        for idx < len(tokens) {
            var tok string = tokens[Java_PostIncrement(&idx)]; _ = tok
            var cmd rune = rune(tok[0]); _ = cmd
            if cmd == 'M' || cmd == 'm' {
                var x int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = x
                var y int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = y
                if cmd == 'm' {
                    x += curX
                    y += curY
                }
                curX = x
                curY = y
                startX = x
                startY = y
                outX[count[0]] = curX
                outY[count[0]] = curY
                count[0]++
            } else if cmd == 'L' || cmd == 'l' {
                var x int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = x
                var y int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = y
                if cmd == 'l' {
                    x += curX
                    y += curY
                }
                curX = x
                curY = y
                outX[count[0]] = curX
                outY[count[0]] = curY
                count[0]++
            } else if cmd == 'V' || cmd == 'v' {
                var y int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = y
                if cmd == 'v' {
                    y += curY
                }
                curY = y
                outX[count[0]] = curX
                outY[count[0]] = curY
                count[0]++
            } else if cmd == 'C' || cmd == 'c' {
                var x1 int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = x1
                var y1 int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = y1
                var x2 int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = x2
                var y2 int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = y2
                var x3 int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = x3
                var y3 int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = y3
                if cmd == 'c' {
                    x1 += curX
                    y1 += curY
                    x2 += curX
                    y2 += curY
                    x3 += curX
                    y3 += curY
                }
                PGUtil_GetInstance().FlattenCubicBezier(curX, curY, x1, y1, x2, y2, x3, y3, outX, outY, count, maxPoints)
                curX = x3
                curY = y3
            } else if cmd == 'Z' || cmd == 'z' {
                curX = startX
                curY = startY
                outX[count[0]] = curX
                outY[count[0]] = curY
                count[0]++
            }
            for idx < len(tokens) && !PathParser_IsCommand(tokens[idx]) {
                if cmd == 'M' || cmd == 'm' || cmd == 'L' || cmd == 'l' {
                    var x int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = x
                    var y int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = y
                    if cmd == 'm' || cmd == 'l' {
                        x += curX
                        y += curY
                    }
                    curX = x
                    curY = y
                    outX[count[0]] = curX
                    outY[count[0]] = curY
                    count[0]++
                } else if cmd == 'V' || cmd == 'v' {
                    var y int = fixedPoint.ToFixed(func() float64 { value, _ := strconv.ParseFloat(tokens[Java_PostIncrement(&idx)], 64); return value }()); _ = y
                    if cmd == 'v' {
                        y += curY
                    }
                    curY = y
                    outX[count[0]] = curX
                    outY[count[0]] = curY
                    count[0]++
                }
            }
        }
        return count[0]
}


func PathParser_IsCommand(tok string) bool {
        var c rune = rune(tok[0]); _ = c
        return strings.Index("MmLlVvCcZz", string(c)) >= 0
}

