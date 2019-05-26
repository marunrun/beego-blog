package controllers

import (
	"github.com/astaxie/beego"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
}

var IsLogin bool = false

func (this *BaseController) Prepare() {

	if err := this.GetSession("uid"); err != nil {
		IsLogin = true
		// 管理员名称
		adminName := this.GetSession("name")
		this.Data["Admin"] = adminName
	}
	this.Data["IsLogin"] = IsLogin
}

// 返回分页的html
func (this *BaseController) paginate(page int, total int64, limit int64) (html string) {

	// 获取当前url
	url := this.Ctx.Request.RequestURI
	// 如果当前url带了参数就加个&连接符
	if strings.Contains(url, "?") {
		reg, err := regexp.Compile(`\&?page=[\d]*\&?`)
		if err == nil {
			url = reg.ReplaceAllString(url, "")
		}
		url += "&"
	} else {
		url += "?"
	}

	maxPage := math.Ceil(float64(total) / float64(limit))

	html = `<nav aria-label="Page navigation">
				<ul class="pagination">`

	if maxPage < 2 {
		html += `<li class="disabled">
					<a href="" aria-label="Previous">
						<span aria-hidden="true">&laquo;</span>
					</a>
				</li>
			<li class="active"><a href="` + url + `page=1">1</a></li>
			<li class="disabled">
				<a href="#" aria-label="Next">
					<span aria-hidden="true">&raquo;</span>
				</a>
			</li>
		</ul>
	</nav>`
	} else {
		if page == 1 {
			html += `<li class="disabled">
					<a href="" aria-label="Previous">
						<span aria-hidden="true">&laquo;</span>
					</a>
				</li>`
		} else {
			html += `<li>
					<a href="` + url + `page=` + strconv.Itoa(page-1) + `" aria-label="Previous">
						<span aria-hidden="true">&laquo;</span>
					</a>
				</li>`
		}

		if maxPage > 10 {
			var start int
			var end int

			if page == 1 {
				html += `<li class="active"><a href="` + url + `page=1">1</a></li>`
				for i := 2; i <= 10; i++ {
					html += `<li><a href="` + url + `page=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
				}
				html += `<li><a href="` + url + `page=` + strconv.Itoa(int(maxPage)) + `">` + strconv.Itoa(int(maxPage)) + `...</a></li>`
			}else if page == int(maxPage) {

				if page - 10 == 1 {
					start = 2
				}else {
					start = page - 10
				}

				html += `<li><a href="` + url + `page=1">1...</a></li>`
				for i := start; i <= int(maxPage); i++ {
					if i == page {
						html += `<li class="active"><a href="` + url + `page=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
						continue
					}
					html += `<li><a href="` + url + `page=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
				}


			}else {

				html += `<li><a href="` + url + `page=1">1...</a></li>`

				if page <= 6 {
					start = 2
				}else {
					start = page - 5
				}

				if page + 5 > int(maxPage) {
					end = int(maxPage)
				}else {
					end = page +5
				}

				for i := start; i <= end; i++ {
					if i == page {
						html += `<li class="active"><a href="` + url + `page=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
						continue
					}
					html += `<li><a href="` + url + `page=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
				}
			}
		} else {
			for i := 1; i <= int(maxPage); i++ {
				if i == page {
					html += `<li class="active"><a href="` + url + `page=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
					continue
				}
				html += `<li><a href="` + url + `page=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
			}
		}

		if page >= int(maxPage) {
			html +=
				`<li class="disabled">
      				<a href="#" aria-label="Next">
        				<span aria-hidden="true">&raquo;</span>
      				</a>
    			</li>`
		} else {
			html +=
				`<li>
      				<a href="` + url + `page=` + strconv.Itoa(page+1) + `" aria-label="Next">
        				<span aria-hidden="true">&raquo;</span>
      				</a>
    			</li>`
		}

	}

	return html
}
