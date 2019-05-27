package controllers

import (
	"beego-blog/models"
	"encoding/csv"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"os"
	"strconv"
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
	/*	var file *xlsx.File
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
		cell.Value = "文章数量"*/
	categorys := models.GetAllCategory()
	/*
		for _,value := range categorys{
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = strconv.Itoa(value.Id)
			cell = row.AddCell()
			cell.Value = value.Title
			cell = row.AddCell()
			cell.Value = value.Created.Format("2006-01-02 15:04:05")
			cell = row.AddCell()
			cell.Value = strconv.FormatInt(value.Views,10)
			cell = row.AddCell()
			cell.Value = strconv.FormatInt(value.TopicCount,10)
		}

		err = file.Save("static/uploads/MyXLSXFile.xlsx")
		if err != nil {
			fmt.Printf(err.Error())
		}

		defer os.Remove("static/uploads/MyXLSXFile.xlsx")

		this.Ctx.Output.Download("static/uploads/MyXLSXFile.xlsx")*/

	columns := []string{"Id", "分类名称", "创建时间", "浏览次数", "文章数量"}
	writeToCSV("static/uploads/test.csv", columns, categorys)

	this.Ctx.Output.Download("static/uploads/test.csv")
}

func writeToCSV(file string, columns []string, totalValues []models.Category) {
	f, err := os.Create(file)
	// fmt.Println(columns)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	//f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	for i, value := range totalValues {
		row := []string{
			strconv.Itoa(value.Id),
			value.Title,
			value.Created.Format("2006-01-02 15:04:05"),
			strconv.FormatInt(value.Views, 10),
			strconv.FormatInt(value.TopicCount, 10),
		}
		//第一次写列名+第一行数据
		if i == 0 {
			w.Write(columns)

			w.Write(row)
		} else {
			w.Write(row)
		}
	}

	w.Flush()
	fmt.Println("处理完毕：", file)

}
