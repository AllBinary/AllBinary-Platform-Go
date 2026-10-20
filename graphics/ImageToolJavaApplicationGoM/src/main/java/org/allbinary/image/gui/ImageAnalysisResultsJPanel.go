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

type ImageAnalysisResultsJPanel struct {
    JPanel
    AvgBlueJTextField *JTextField
    AvgGreenJTextField *JTextField
    AvgRedJTextField *JTextField
    JComboBox1 *JComboBox
    JLabel1 *JLabel
    JLabel10 *JLabel
    JLabel11 *JLabel
    JLabel12 *JLabel
    JLabel13 *JLabel
    JLabel2 *JLabel
    JLabel3 *JLabel
    JLabel4 *JLabel
    JLabel5 *JLabel
    JLabel6 *JLabel
    JLabel7 *JLabel
    JLabel8 *JLabel
    JLabel9 *JLabel
    MatchingPixelsJTextField *JTextField
    MaxBlueJTextField *JTextField
    MaxGreenJTextField *JTextField
    MaxRedJTextField *JTextField
    MinBlueJTextField *JTextField
    MinGreenJTextField *JTextField
    MinRedJTextField *JTextField
    PercentMatchingJTextField *JTextField
    TotalPixelsJTextField *JTextField
}


func (this *ImageAnalysisResultsJPanel) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewImageAnalysisResultsJPanel(imageAnalysisResultsArray []*ImageAnalysisResults) *ImageAnalysisResultsJPanel {
    this := &ImageAnalysisResultsJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
        var imageAnalysisResults *ImageAnalysisResults = imageAnalysisResultsArray[0]; _ = imageAnalysisResults
        this.UpdateImageColorRangeResults(imageAnalysisResults)
    return this
}


func (this *ImageAnalysisResultsJPanel) InitComponents() {
        this.JLabel1 = NewJLabel()
        this.JLabel2 = NewJLabel()
        this.JLabel3 = NewJLabel()
        this.PercentMatchingJTextField = NewJTextField()
        this.TotalPixelsJTextField = NewJTextField()
        this.MatchingPixelsJTextField = NewJTextField()
        this.JComboBox1 = NewJComboBox()
        this.JLabel4 = NewJLabel()
        this.JLabel5 = NewJLabel()
        this.JLabel6 = NewJLabel()
        this.JLabel7 = NewJLabel()
        this.AvgRedJTextField = NewJTextField()
        this.AvgGreenJTextField = NewJTextField()
        this.AvgBlueJTextField = NewJTextField()
        this.MinGreenJTextField = NewJTextField()
        this.JLabel8 = NewJLabel()
        this.JLabel9 = NewJLabel()
        this.JLabel10 = NewJLabel()
        this.MinRedJTextField = NewJTextField()
        this.MinBlueJTextField = NewJTextField()
        this.JLabel11 = NewJLabel()
        this.JLabel12 = NewJLabel()
        this.JLabel13 = NewJLabel()
        this.MaxRedJTextField = NewJTextField()
        this.MaxGreenJTextField = NewJTextField()
        this.MaxBlueJTextField = NewJTextField()
        this.JLabel1.SetText("Matching Pixels:")
        this.JLabel2.SetText("Total Pixels:")
        this.JLabel3.SetText("Percent Matching:")
        this.PercentMatchingJTextField.SetEditable(false)
        this.TotalPixelsJTextField.SetEditable(false)
        this.MatchingPixelsJTextField.SetEditable(false)
        this.JComboBox1.SetModel(NewDefaultComboBoxModel([]string{CommonPhoneStrings_GetInstance().ONE}))
        this.JComboBox1.SetEnabled(false)
        this.JLabel4.SetText("Image:")
        this.JLabel5.SetText("Average Red:")
        this.JLabel6.SetText("Average Green:")
        this.JLabel7.SetText("Average Blue:")
        this.AvgRedJTextField.SetEditable(false)
        this.AvgGreenJTextField.SetEditable(false)
        this.AvgBlueJTextField.SetEditable(false)
        this.MinGreenJTextField.SetEditable(false)
        this.JLabel8.SetText("Min Red:")
        this.JLabel9.SetText("Min Green:")
        this.JLabel10.SetText("Min Blue:")
        this.MinRedJTextField.SetEditable(false)
        this.MinBlueJTextField.SetEditable(false)
        this.JLabel11.SetText("Max Red:")
        this.JLabel12.SetText("Max Green:")
        this.JLabel13.SetText("Max Blue:")
        this.MaxRedJTextField.SetEditable(false)
        this.MaxGreenJTextField.SetEditable(false)
        this.MaxBlueJTextField.SetEditable(false)
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddContainerGap().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JLabel3).AddComponent(this.JLabel2).AddComponent(this.JLabel1).AddComponent(this.JLabel5).AddComponent(this.JLabel6).AddComponent(this.JLabel7).AddComponent(this.JLabel4)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JComboBox1, javax_Swing.GroupLayout.PREFERRED_SIZE, 81, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.AvgBlueJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 81, Short_MAX_VALUE).AddComponent(this.AvgGreenJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 81, Short_MAX_VALUE).AddComponent(this.AvgRedJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 81, Short_MAX_VALUE).AddComponent(this.TotalPixelsJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 81, Short_MAX_VALUE).AddComponent(this.PercentMatchingJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 81, Short_MAX_VALUE).AddComponent(this.MatchingPixelsJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 81, Short_MAX_VALUE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JLabel12).AddComponent(this.JLabel13).AddComponent(this.JLabel8).AddComponent(this.JLabel9).AddComponent(this.JLabel10).AddComponent(this.JLabel11)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.MinRedJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 82, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.MaxGreenJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 82, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.MinGreenJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 82, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.MinBlueJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 82, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.MaxRedJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 82, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.MaxBlueJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 82, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddContainerGap(18, Short_MAX_VALUE)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel4).AddComponent(this.JComboBox1, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.MatchingPixelsJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel8).AddComponent(this.MinRedJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel2).AddComponent(this.TotalPixelsJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel9).AddComponent(this.MinGreenJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel3).AddComponent(this.PercentMatchingJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel10).AddComponent(this.MinBlueJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel5).AddComponent(this.AvgRedJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel11).AddComponent(this.MaxRedJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel6).AddComponent(this.AvgGreenJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.MaxGreenJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel12)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel7).AddComponent(this.AvgBlueJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel13).AddComponent(this.MaxBlueJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))))
}


func (this *ImageAnalysisResultsJPanel) UpdateImageColorRangeResults(imageAnalysisResults *ImageAnalysisResults) {
        var imageColorRangeResults *ImageColorRangeResults = imageAnalysisResults.GetImageColorRangeResults(); _ = imageColorRangeResults
        this.MatchingPixelsJTextField.SetText(fmt.Sprint(int64(imageColorRangeResults.GetMatchingPixelsChecked())))
        this.TotalPixelsJTextField.SetText(fmt.Sprint(int64(imageColorRangeResults.GetTotalPixelsChecked())))
        this.PercentMatchingJTextField.SetText(fmt.Sprint(float32(imageColorRangeResults.GetPercent())))
        var imageColorResults *ImageColorResults = imageAnalysisResults.GetImageColorResults(); _ = imageColorResults
        this.AvgRedJTextField.SetText(fmt.Sprint(float32(imageColorResults.GetColorAverage().GetAvgRed())))
        this.AvgGreenJTextField.SetText(fmt.Sprint(float32(imageColorResults.GetColorAverage().GetAvgGreen())))
        this.AvgBlueJTextField.SetText(fmt.Sprint(float32(imageColorResults.GetColorAverage().GetAvgBlue())))
        this.MinRedJTextField.SetText(fmt.Sprint(int64(imageColorResults.GetColorRange().GetMinRed())))
        this.MinGreenJTextField.SetText(fmt.Sprint(int64(imageColorResults.GetColorRange().GetMinGreen())))
        this.MinBlueJTextField.SetText(fmt.Sprint(int64(imageColorResults.GetColorRange().GetMinBlue())))
        this.MaxRedJTextField.SetText(fmt.Sprint(int64(imageColorResults.GetColorRange().GetMaxRed())))
        this.MaxGreenJTextField.SetText(fmt.Sprint(int64(imageColorResults.GetColorRange().GetMaxGreen())))
        this.MaxBlueJTextField.SetText(fmt.Sprint(int64(imageColorResults.GetColorRange().GetMaxBlue())))
}

