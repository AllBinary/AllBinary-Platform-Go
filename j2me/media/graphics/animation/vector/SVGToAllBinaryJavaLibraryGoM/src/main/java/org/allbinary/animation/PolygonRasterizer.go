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

type PolygonRasterizer struct {
}

var PolygonRasterizer_Instance *PolygonRasterizer = NewPolygonRasterizer()

func NewPolygonRasterizer() *PolygonRasterizer {
    this := &PolygonRasterizer{}
    return this
}


func PolygonRasterizer_GetInstance() *PolygonRasterizer {
        return PolygonRasterizer_Instance
}


func (this *PolygonRasterizer) GetPixelBuffer(svgPath string, pathX []int, pathY []int, pixelBuffer []int, width int, height int) []int {
        var pointCount int = PathParser_ParsePath(svgPath, pathX, pathY, len(pathX)); _ = pointCount
        for i := int(0); i < len(pixelBuffer); i++ {
            pixelBuffer[i] = 0xFFFFFFFF
        }
        PolygonRasterizer_GetInstance().FillPolygon(pathX, pathY, pointCount, pixelBuffer, width, height, 0xFFFF00FF)
        return pixelBuffer
}


func (this *PolygonRasterizer) FillPolygon(polyX []int, polyY []int, numPoints int, buffer []int, width int, height int, colorRGB int) {
        if numPoints < 3 {
            return
        }
        var fixedPoint *FixedPoint = FixedPoint_GetInstance(); _ = fixedPoint
        var x []int = make([]int, numPoints); _ = x
        var y []int = make([]int, numPoints); _ = y
        var minY int = height; _ = minY
        var maxY int = 0; _ = maxY
        for i := int(0); i < numPoints; i++ {
            x[i] = fixedPoint.ToInt(polyX[i])
            y[i] = fixedPoint.ToInt(polyY[i])
            if y[i] < minY {
                minY = y[i]
            }
            if y[i] > maxY {
                maxY = y[i]
            }
        }
        if minY < 0 {
            minY = 0
        }
        if maxY >= height {
            maxY = height - 1
        }
        var scanlineX []int = make([]int, numPoints); _ = scanlineX
        for scanY := int(minY); scanY <= maxY; scanY++ {
            var intersectCount int = 0; _ = intersectCount
            for i := int(0); i < numPoints; i++ {
                var next int = (i + 1) % numPoints; _ = next
                var x1 int = x[i]; _ = x1
                var y1 int = y[i]; _ = y1
                var x2 int = x[next]; _ = x2
                var y2 int = y[next]; _ = y2
                if (y1 < scanY && y2 >= scanY) || (y2 < scanY && y1 >= scanY) {
                    if y2 != y1 {
                        var interX int = x1 + (scanY - y1) * (x2 - x1) / (y2 - y1); _ = interX
                        scanlineX[Java_PostIncrement(&intersectCount)] = interX
                    }
                }
            }
            for i := int(0); i < intersectCount - 1; i++ {
                for j := int(i + 1); j < intersectCount; j++ {
                    if scanlineX[i] > scanlineX[j] {
                        var temp int = scanlineX[i]; _ = temp
                        scanlineX[i] = scanlineX[j]
                        scanlineX[j] = temp
                    }
                }
            }
            for i := int(0); i < intersectCount; i += 2 {
                if i + 1 >= intersectCount {
                    break
                }
                var startX int = scanlineX[i]; _ = startX
                var endX int = scanlineX[i + 1]; _ = endX
                if startX < 0 {
                    startX = 0
                }
                if endX >= width {
                    endX = width - 1
                }
                var rowOffset int = scanY * width; _ = rowOffset
                for pixelX := int(startX); pixelX <= endX; pixelX++ {
                    buffer[rowOffset + pixelX] = colorRGB
                }
            }
        }
}


func (this *PolygonRasterizer) Main(args []string) {
        var width int = 100; _ = width
        var height int = 100; _ = height
        var pathX []int = make([]int, 200); _ = pathX
        var pathY []int = make([]int, 200); _ = pathY
        var pixelBuffer []int = make([]int, width * height); _ = pixelBuffer
        var svgPath string = "M 32.139812,81.554575 14.450128,74.294372 V 50.000034 25.705697 l 17.73231,-7.268329 C 41.935209,14.439788 49.953112,11.168898 50,11.168723 c 0.04689,-1.74e-4 8.064791,3.270238 17.817562,7.267583 l 17.73231,7.267899 V 50 74.295795 l -17.73231,7.2679 c -9.752771,3.997344 -17.789855,7.264116 -17.860188,7.259492 -0.07033,-0.0046 -8.088235,-3.275499 -17.817562,-7.268612 z"; _ = svgPath
        PolygonRasterizer_GetInstance().GetPixelBuffer(svgPath, pathX, pathY, pixelBuffer, width, height)
}

