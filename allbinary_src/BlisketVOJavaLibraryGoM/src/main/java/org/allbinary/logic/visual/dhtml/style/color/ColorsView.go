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

type ColorsView struct {
    CommonPhoneStrings *CommonPhoneStrings
    HEXRADIX int
    MAX int
    Delta int
}

// Implements DomNodeInterface


func (this *ColorsView) initFields() {
    this.CommonPhoneStrings = CommonPhoneStrings_GetInstance()
    this.HEXRADIX = 16
    this.MAX = 255
    this.Delta = 20
}


func NewColorsView() *ColorsView {
    this := &ColorsView{}
    this.initFields()
    return this
}


// throws Exception
func (this *ColorsView) GetColorNode(document *Document, redIndex int, greenIndex int, blueIndex int) *Node {
        var hexColorStringBuffer *StringMaker = NewStringMaker(); _ = hexColorStringBuffer
        if redIndex < this.HEXRADIX {
            hexColorStringBuffer.Append(this.CommonPhoneStrings.ZERO)
        }
        hexColorStringBuffer.Append(Integer_ToString(redIndex, this.HEXRADIX))
        if greenIndex < this.HEXRADIX {
            hexColorStringBuffer.Append(this.CommonPhoneStrings.ZERO)
        }
        hexColorStringBuffer.Append(Integer_ToString(greenIndex, this.HEXRADIX))
        if blueIndex < this.HEXRADIX {
            hexColorStringBuffer.Append(this.CommonPhoneStrings.ZERO)
        }
        hexColorStringBuffer.Append(Integer_ToString(blueIndex, this.HEXRADIX))
        var colorNode *Node = ModDomHelper_CreateNameValueNodes(document, ColorData_GetInstance().NAME, fmt.Sprint(hexColorStringBuffer)); _ = colorNode
        return colorNode
}


// throws Exception
func (this *ColorsView) ToXmlNode(document *Document) *Node {
        var node *Node = ModDomHelper_CreateNameValueNodes(document, ColorsData_GetInstance().NAME, "RGB Colors In Hex"); _ = node
        var blackNode *Node = ModDomHelper_CreateNameValueNodes(document, ColorData_GetInstance().NAME, "000000"); _ = blackNode
        node.AppendChild(blackNode)
        for index := int(0); index < this.MAX; index += this.Delta {
            var hexColorStringBuffer *StringMaker = NewStringMaker(); _ = hexColorStringBuffer
            if index < this.HEXRADIX {
                hexColorStringBuffer.Append(this.CommonPhoneStrings.ZERO)
            }
            hexColorStringBuffer.Append(Integer_ToString(index, this.HEXRADIX))
            if index < this.HEXRADIX {
                hexColorStringBuffer.Append(this.CommonPhoneStrings.ZERO)
            }
            hexColorStringBuffer.Append(Integer_ToString(index, this.HEXRADIX))
            if index < this.HEXRADIX {
                hexColorStringBuffer.Append(this.CommonPhoneStrings.ZERO)
            }
            hexColorStringBuffer.Append(Integer_ToString(index, this.HEXRADIX))
            var colorNode *Node = ModDomHelper_CreateNameValueNodes(document, ColorData_GetInstance().NAME, fmt.Sprint(hexColorStringBuffer)); _ = colorNode
            node.AppendChild(colorNode)
        }
        var whiteNode *Node = ModDomHelper_CreateNameValueNodes(document, ColorData_GetInstance().NAME, "FFFFFF"); _ = whiteNode
        node.AppendChild(whiteNode)
        for blueIndex := int(this.MAX / 2); blueIndex < this.MAX; blueIndex += this.Delta {
            for greenIndex := int(0); greenIndex < this.MAX / 2; greenIndex += this.Delta {
                for redIndex := int(0); redIndex < this.MAX / 2; redIndex += this.Delta {
                    node.AppendChild(this.GetColorNode(document, redIndex, greenIndex, blueIndex))
                }
            }
        }
        for redIndex := int(this.MAX / 2); redIndex < this.MAX; redIndex += this.Delta {
            for greenIndex := int(0); greenIndex < this.MAX / 2; greenIndex += this.Delta {
                for blueIndex := int(0); blueIndex < this.MAX / 2; blueIndex += this.Delta {
                    node.AppendChild(this.GetColorNode(document, redIndex, greenIndex, blueIndex))
                }
            }
        }
        for greenIndex := int(this.MAX / 2); greenIndex < this.MAX; greenIndex += this.Delta {
            for redIndex := int(0); redIndex < this.MAX / 2; redIndex += this.Delta {
                for blueIndex := int(0); blueIndex < this.MAX / 2; blueIndex += this.Delta {
                    node.AppendChild(this.GetColorNode(document, redIndex, greenIndex, blueIndex))
                }
            }
        }
        for greenIndex := int(0); greenIndex < this.MAX; greenIndex += 2 * this.Delta {
            for redIndex := int(0); redIndex < this.MAX; redIndex += 2 * this.Delta {
                for blueIndex := int(0); blueIndex < this.MAX; blueIndex += 2 * this.Delta {
                    node.AppendChild(this.GetColorNode(document, redIndex, greenIndex, blueIndex))
                }
            }
        }
        return node
}

