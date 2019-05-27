package controllers

import (
	"beego-blog/models"
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"strconv"
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
	params := this.Input()
	fmt.Println(params)
	f, h, err := this.GetFile("file")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	this.SaveToFile("file", "static/uploads/"+h.Filename)

}

// @router /down [get]
func (this *ApiController) Down() {

	this.writeToCSV()

	defer os.Remove("static/uploads/test.csv")
	this.Ctx.Output.Download("static/uploads/test.csv")

	//defer os.Remove("static/uploads/MyXLSXFile.xlsx")
	//this.Ctx.Output.Download("static/uploads/MyXLSXFile.xlsx")
}

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
