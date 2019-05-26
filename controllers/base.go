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
		// 正则匹配 去掉原先的page=xx
		reg, err := regexp.Compile(`\&?page=[\d]*\&?`)
		if err == nil {
			url = reg.ReplaceAllString(url, "")
		}
		url += "&"
	} else {
		url += "?"
	}

	//算出最大的页码
	maxPage := math.Ceil(float64(total) / float64(limit))

	// 开始拼接分页的HTML字符串
	html = `<nav aria-label="Page navigation">
				<ul class="pagination">`

	// 如果最大页码小于2 也就只有一页了....
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
		// 如果大于两页, 并且当前页是第一页  那么前一页不可以点击
		if page == 1 {
			html += `<li class="disabled">
					<a href="" aria-label="Previous">
						<span aria-hidden="true">&laquo;</span>
					</a>
				</li>`
		} else {
			// 如果不是第一页,那么前一页可以点击
			html += `<li>
					<a href="` + url + `page=` + strconv.Itoa(page-1) + `" aria-label="Previous">
						<span aria-hidden="true">&laquo;</span>
					</a>
				</li>`
		}

		if maxPage > 10 {
			var start int
			var end int

			// 如果总页码大于10 并且当前页是第一页,那么第一页是激活状态. 然后循环拼接10条页面
			if page == 1 {
				html += `<li class="active"><a href="` + url + `page=1">1</a></li>`
				for i := 2; i <= 10; i++ {
					html += `<li><a href="` + url + `page=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
				}
				html += `<li><a href="` + url + `page=` + strconv.Itoa(int(maxPage)) + `">` + strconv.Itoa(int(maxPage)) + `...</a></li>`
			}else if page == int(maxPage) {
				// 如果当前页是最后一页, 并且减去10页之后是第一页,那么就从第二页开始拼接页码
				if page - 10 == 1 {
					start = 2
				}else { // 否则就减去10页再开始拼接页码
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

				// 如果当前页小于等于6 那么就从第二页开始拼接页码,否则就从当前页的前五页开始
				if page <= 6 {
					start = 2
				}else {
					start = page - 5
				}

				// 如果当前页加五条小于最大页码,那么就拼接到最后一页就ok,否则拼接到当前页的后五条
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
			// 如果总页码不到10条 直接循环拼接出来
			for i := 1; i <= int(maxPage); i++ {
				if i == page {
					html += `<li class="active"><a href="` + url + `page=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
					continue
				}
				html += `<li><a href="` + url + `page=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
			}
		}


		// 拼接下一页
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
