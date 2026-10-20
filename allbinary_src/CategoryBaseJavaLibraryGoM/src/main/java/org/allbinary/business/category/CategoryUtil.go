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

type CategoryUtil struct {
}



func NewCategoryUtil() *CategoryUtil {
    this := &CategoryUtil{}
    return this
}


// synchronized
func CategoryUtil_GetPathLevel(categoryPath IAbPath) int {
        var count int = StringUtils_CountMatches(fmt.Sprint(categoryPath), AbPathData_GetInstance().SEPARATOR); _ = count
        return count
}


// synchronized
func CategoryUtil_GetNameFromNode(node *Node) string {
        var attributes *NamedNodeMap = node.GetAttributes(); _ = attributes
        var attrNode *Attr = any(attributes.GetNamedItem(CategoryData_GetInstance().LABEL)).(*Attr); _ = attrNode
        if attrNode != nil {
            return attrNode.GetValue()
        } else {
            var labelNode *Node = DomSearchHelper_GetNodeNoThrow(CategoryData_GetInstance().LABEL, node.GetChildNodes()); _ = labelNode
            if labelNode != nil {
                return DomNodeHelper_GetTextNodeValue(labelNode)
            }
        }
        return StringUtil_GetInstance().EMPTY_STRING
}

