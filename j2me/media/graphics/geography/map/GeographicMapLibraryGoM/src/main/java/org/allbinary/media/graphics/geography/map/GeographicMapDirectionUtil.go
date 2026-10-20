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

type GeographicMapDirectionUtil struct {
    MESSAGE string
    ERROR string
    EQUAL string
}

var GeographicMapDirectionUtil_Instance *GeographicMapDirectionUtil = NewGeographicMapDirectionUtil()

func (this *GeographicMapDirectionUtil) initFields() {
    this.MESSAGE = "Cell Position should have been cached!!!"
    this.ERROR = "Error: "
    this.EQUAL = " == "
}


func GeographicMapDirectionUtil_GetInstance() *GeographicMapDirectionUtil {
        return GeographicMapDirectionUtil_Instance
}


func NewGeographicMapDirectionUtil() *GeographicMapDirectionUtil {
    this := &GeographicMapDirectionUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *GeographicMapDirectionUtil) GetDirectionFromCellPositionToAdjacentCellPosition(fromGeographicMapCellPosition IGeographicMapCellPosition, toGeographicMapCellPosition IGeographicMapCellPosition) *Direction {
        var fromColumn int = fromGeographicMapCellPosition.GetColumn(); _ = fromColumn
        var fromRow int = fromGeographicMapCellPosition.GetRow(); _ = fromRow
        var goColumn int = toGeographicMapCellPosition.GetColumn(); _ = goColumn
        var goRow int = toGeographicMapCellPosition.GetRow(); _ = goRow
        var directionFactory *DirectionFactory = DirectionFactory_GetInstance(); _ = directionFactory
        if fromColumn - 1 == goColumn {
            return directionFactory.LEFT
        } else if fromColumn + 1 == goColumn {
            return directionFactory.RIGHT
        } else if fromRow - 1 == goRow {
            return directionFactory.UP
        } else if fromRow + 1 == goRow {
            return directionFactory.DOWN
        } else if fromGeographicMapCellPosition == toGeographicMapCellPosition {
            return directionFactory.NO_DIRECTION
        } else if fromRow == goRow && fromColumn == goColumn {
            panic(errors.New(fmt.Sprint(this.MESSAGE)))
        }
        return directionFactory.NOT_BORDERED_WITH
}


// throws Exception
func (this *GeographicMapDirectionUtil) GetEightDirectionFromCellPositionToAdjacentCellPosition(fromGeographicMapCellPosition IGeographicMapCellPosition, toGeographicMapCellPosition IGeographicMapCellPosition) *Direction {
        var fromColumn int = fromGeographicMapCellPosition.GetColumn(); _ = fromColumn
        var fromRow int = fromGeographicMapCellPosition.GetRow(); _ = fromRow
        var goColumn int = toGeographicMapCellPosition.GetColumn(); _ = goColumn
        var goRow int = toGeographicMapCellPosition.GetRow(); _ = goRow
        var directionFactory *DirectionFactory = DirectionFactory_GetInstance(); _ = directionFactory
        if fromRow == goRow {
            if fromColumn - 1 == goColumn {
                return directionFactory.LEFT
            } else if fromColumn + 1 == goColumn {
                return directionFactory.RIGHT
            }
        } else if fromRow - 1 == goRow {
            if fromColumn == goColumn {
                return directionFactory.UP
            }
            if fromColumn - 1 == goColumn {
                return directionFactory.UP_LEFT
            } else if fromColumn + 1 == goColumn {
                return directionFactory.UP_RIGHT
            }
        } else if fromRow + 1 == goRow {
            if fromColumn == goColumn {
                return directionFactory.DOWN
            }
            if fromColumn - 1 == goColumn {
                return directionFactory.DOWN_LEFT
            } else if fromColumn + 1 == goColumn {
                return directionFactory.DOWN_RIGHT
            }
        } else if fromGeographicMapCellPosition == toGeographicMapCellPosition {
            return directionFactory.NO_DIRECTION
        } else if fromRow == goRow && fromColumn == goColumn {
            panic(errors.New(fmt.Sprint(this.MESSAGE)))
        }
        return directionFactory.NOT_BORDERED_WITH
}


// throws Exception
func (this *GeographicMapDirectionUtil) GetDirectionFromCellPositionToCellPosition(fromGeographicMapCellPosition IGeographicMapCellPosition, toGeographicMapCellPosition IGeographicMapCellPosition) *Direction {
        var fromColumn int = fromGeographicMapCellPosition.GetColumn(); _ = fromColumn
        var fromRow int = fromGeographicMapCellPosition.GetRow(); _ = fromRow
        var goColumn int = toGeographicMapCellPosition.GetColumn(); _ = goColumn
        var goRow int = toGeographicMapCellPosition.GetRow(); _ = goRow
        var directionFactory *DirectionFactory = DirectionFactory_GetInstance(); _ = directionFactory
        if fromColumn > goColumn {
            return directionFactory.LEFT
        } else if fromColumn < goColumn {
            return directionFactory.RIGHT
        } else if fromRow > goRow {
            return directionFactory.UP
        } else if fromRow < goRow {
            return directionFactory.DOWN
        }
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var string_ string = fmt.Sprint(stringMaker.Append(CellPosition_ToStringCellPosition(fromGeographicMapCellPosition)).Append(" == ").Append(CellPosition_ToStringCellPosition(toGeographicMapCellPosition))); _ = string_
        stringMaker.Delete(0, stringMaker.Length())
        panic(errors.New(fmt.Sprint(fmt.Sprint(stringMaker.Append("Error: ").Append(string_)))))
}


// throws Exception
func (this *GeographicMapDirectionUtil) GetEightDirectionFromCellPositionToCellPosition(fromGeographicMapCellPosition IGeographicMapCellPosition, toGeographicMapCellPosition IGeographicMapCellPosition) *Direction {
        var fromColumn int = fromGeographicMapCellPosition.GetColumn(); _ = fromColumn
        var fromRow int = fromGeographicMapCellPosition.GetRow(); _ = fromRow
        var goColumn int = toGeographicMapCellPosition.GetColumn(); _ = goColumn
        var goRow int = toGeographicMapCellPosition.GetRow(); _ = goRow
        var directionFactory *DirectionFactory = DirectionFactory_GetInstance(); _ = directionFactory
        if fromGeographicMapCellPosition == toGeographicMapCellPosition {
            return directionFactory.NO_DIRECTION
        } else if fromRow == goRow {
            if fromColumn > goColumn {
                return directionFactory.LEFT
            } else if fromColumn < goColumn {
                return directionFactory.RIGHT
            }
        } else if fromRow > goRow {
            if fromColumn > goColumn {
                return directionFactory.UP_LEFT
            } else if fromColumn < goColumn {
                return directionFactory.UP_RIGHT
            } else if fromColumn == goColumn {
                return directionFactory.UP
            }
        } else if fromRow < goRow {
            if fromColumn > goColumn {
                return directionFactory.DOWN_LEFT
            } else if fromColumn < goColumn {
                return directionFactory.DOWN_RIGHT
            } else if fromColumn == goColumn {
                return directionFactory.DOWN
            }
        }
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.ERROR)
        stringBuffer.Append(CellPosition_ToStringCellPosition(fromGeographicMapCellPosition))
        stringBuffer.Append(this.EQUAL)
        stringBuffer.Append(CellPosition_ToStringCellPosition(toGeographicMapCellPosition))
        panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
}

