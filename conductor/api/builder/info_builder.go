package builder

import "conductor/common/model"

type ListViewInfoBuilder struct {
	model.ListViewInfo
}

func ListViewInfo() *ListViewInfoBuilder {
	return &ListViewInfoBuilder{}
}

func (lb *ListViewInfoBuilder) Build() model.ListViewInfo {
	return lb.ListViewInfo
}

func (lb *ListViewInfoBuilder) Text(text string) *ListViewInfoBuilder {
	lb.ListViewInfo.Text = text
	return lb
}

func (lb *ListViewInfoBuilder) Icon(icon string) *ListViewInfoBuilder {
	lb.ListViewInfo.Icon = icon
	return lb
}

func (lb *ListViewInfoBuilder) Bold() *ListViewInfoBuilder {
	lb.ListViewInfo.Bold = true
	return lb
}
