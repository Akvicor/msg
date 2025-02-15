package mail

import (
	"fmt"
	"github.com/wneessen/go-mail"
	"strings"
)

type HtmlTable struct {
	title  string
	header []*HtmlTableHeader
	row    [][]*HtmlTableRow
}

type HtmlTableHeader struct {
	Width uint
	Black bool
	Title string
}

type HtmlTableRow struct {
	Black   bool
	Colour  string
	Content string
}

func NewHtmlTable() *HtmlTable {
	return &HtmlTable{
		title:  "",
		header: make([]*HtmlTableHeader, 0),
		row:    make([][]*HtmlTableRow, 0),
	}
}

func (h *HtmlTable) ContentType() mail.ContentType {
	return TypeTextHtml
}

func (h *HtmlTable) GetTitle() string {
	return h.title
}

func (h *HtmlTable) SetTitle(title string) {
	h.title = title
}

func (h *HtmlTable) SetHeader(header []*HtmlTableHeader) {
	h.header = header
}

func (h *HtmlTable) AddHeader(width uint, black bool, title string) {
	h.header = append(h.header, &HtmlTableHeader{
		Width: width,
		Black: black,
		Title: title,
	})
}

func (h *HtmlTable) AddRow(row []*HtmlTableRow) {
	h.row = append(h.row, row)
}

func (h *HtmlTable) Content() string {
	content := strings.Builder{}
	content.WriteString(`<html><head><style type="text/css">.table {width: 100%;max-width: 100%;margin-bottom: 20px;border-collapse: collapse;background-color: transparent} td {padding: 8px;line-height: 1.42857143;vertical-align: top;border: 1px solid #ddd;border-top: 1px solid #ddd}.table-bordered {border: 1px solid #ddd} .colour-tag {display: inline-block; color: #fff; padding: 4px 8px; border-radius: 7px; background-color: #007EB3;}</style></head>`)
	title := ""
	if len(h.title) != 0 {
		title = fmt.Sprintf("<h3>%s</h3>", h.title)
	}
	content.WriteString(fmt.Sprintf(`<body>%s<table class="table table-bordered">`, title))
	{ // head
		content.WriteString(`<tr>`)
		for _, v := range h.header {
			t := v.Title
			if v.Black {
				t = fmt.Sprintf("<b>%s</b>", v.Title)
			}
			content.WriteString(fmt.Sprintf(`<td width="%d%%">%s</td>`, v.Width, t))
		}
		content.WriteString(`</tr>`)
	}
	{ // body
		for _, rows := range h.row {
			content.WriteString(`<tr>`)
			for _, row := range rows {
				c := row.Content
				if row.Black {
					c = fmt.Sprintf("<b>%s</b>", c)
				}
				if row.Colour != "" {
					c = fmt.Sprintf(`<a class="colour-tag" style="background-color: %s;">%s</a>`, row.Colour, c)
				}
				content.WriteString(fmt.Sprintf(`<td>%s</td>`, c))
			}
			content.WriteString(`</tr>`)
		}
	}
	content.WriteString(`</table></body></html>`)

	return content.String()
}
