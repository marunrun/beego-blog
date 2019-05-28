package controllers

import (
	"beego-blog/libs"
	"beego-blog/models"
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type ApiController struct {
	beego.Controller
}

func (this *ApiController) URLMapping() {
	this.Mapping("post", this.Upload)
}

// @router /api/upload [post]
func (this *ApiController) Upload() {
	f, h, err := this.GetFile("file")
	if err != nil {
		log.Fatal("getfile err ", err)
		panic(err)
	}
	defer f.Close()

	// 文件路径
	fullPath := "static/uploads/" + h.Filename
	// 保存文件
	err = this.SaveToFile("file", fullPath)
	if err != nil {
		log.Fatal("savefile err ", err)
		panic(err)
	}

	// 获取文件后缀
	ext := filepath.Ext(fullPath)
	// 判断文件是否是需要的
	arr := []interface{}{".xlsx", ".xls", ".csv"}
	if !libs.In_array(ext, arr) {
		os.Remove(fullPath)
		this.Data["json"] = res{400, "请上传xlsx，xls或csv文件"}
		this.ServeJSON()
		return
	}

	switch ext {
	case ".csv":
		this.ReadCsv(fullPath)
	case ".xlsx",".xls":
		//this.ReadExcel(fullPath)
		this.ReadXslx(fullPath)
	default:
		panic("文件有误")
	}

	this.ServeJSON()
}

// @router /down [get]
func (this *ApiController) Down() {
	this.writeToCSV()

	defer os.Remove("static/uploads/test.csv")
	this.Ctx.Output.Download("static/uploads/test.csv")

	//defer os.Remove("static/uploads/MyXLSXFile.xlsx")
	//this.Ctx.Output.Download("static/uploads/MyXLSXFile.xlsx")
}

// 使用excelize导出excel
func (thi *ApiController) WriteToXslx() {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")

	columns := []string{"Id", "分类名称", "创建时间", "浏览次数", "文章数量"}
	axis := []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1"}
	// 设置头
	for key, value := range columns {
		f.SetCellValue("Sheet1", axis[key], value)
	}
	categorys := models.GetAllCategory()
	// 给其他的设置数据
	for key, value := range categorys {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(key+2), value.Id)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(key+2), value.Title)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(key+2), value.Created.Format("2006-01-02 15:04:05"))
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(key+2), value.Views)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(key+2), value.TopicCount)
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	// 根据指定路径保存文件
	err := f.SaveAs("static/uploads/MyXLSXFile.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time.Now().String(), "处理完成")
}
// 使用tealeg/xlsx导出excel
func (this *ApiController) WriteToExcel() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")

	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "Id"
	cell = row.AddCell()
	cell.Value = "分类名称"
	cell = row.AddCell()
	cell.Value = "创建时间"
	cell = row.AddCell()
	cell.Value = "浏览次数"
	cell = row.AddCell()
	cell.Value = "文章数量"
	categorys := models.GetAllCategory()

	for _, value := range categorys {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = strconv.Itoa(value.Id)
		cell = row.AddCell()
		cell.Value = value.Title
		cell = row.AddCell()
		cell.Value = value.Created.Format("2006-01-02 15:04:05")
		cell = row.AddCell()
		cell.Value = strconv.FormatInt(value.Views, 10)
		cell = row.AddCell()
		cell.Value = strconv.FormatInt(value.TopicCount, 10)
	}

	err = file.Save("static/uploads/MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(time.Now().String(), "处理完成")

	//columns := []string{"Id", "分类名称", "创建时间", "浏览次数", "文章数量"}
}
// 导出csv
func (this *ApiController) writeToCSV() {
	columns := []string{"Id", "分类名称", "创建时间", "浏览次数", "文章数量"}

	f, err := os.Create("static/uploads/test.csv")
	// fmt.Println(columns)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	totalValues := models.GetAllCategory()
	for i, value := range totalValues {

		//第一次写列名+第一行数据
		if i == 0 {
			w.Write(columns)
		}
		w.Write([]string{
			strconv.Itoa(value.Id),
			value.Title,
			value.Created.Format("2006-01-02 15:04:05"),
			strconv.FormatInt(value.Views, 10),
			strconv.FormatInt(value.TopicCount, 10),
		})

	}

	w.Flush()
	fmt.Println("处理完毕：", time.Now().String())
}

// 使用excelize读取excel
func (this *ApiController) ReadXslx(filePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		panic(err)
	}
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

}
// 使用tealeg/xlsx读取excel
func (this *ApiController) ReadExcel(filePath string) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		panic(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
// 读取csv文件
func (this *ApiController) ReadCsv(filePath string) {
	conent, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	r2 := csv.NewReader(strings.NewReader(string(conent)))
	ss,_ := r2.ReadAll()
	fmt.Println(ss)
}