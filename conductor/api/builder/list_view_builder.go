package builder

import "conductor/common/model"

type ListViewBuilder struct {
	model.ListView
}

func ListView() *ListViewBuilder {
	return &ListViewBuilder{}
}

func (lb *ListViewBuilder) Build() model.ListView {
	return lb.ListView
}

func (lb *ListViewBuilder) Title(title string) *ListViewBuilder {
	lb.ListView.Title = title
	return lb
}

func (lb *ListViewBuilder) ShowStatus(showStatus bool) *ListViewBuilder {
	lb.ListView.ShowStatus = showStatus
	return lb
}

func (lb *ListViewBuilder) Items(items []model.ListViewItem) *ListViewBuilder {
	lb.ListView.Items = append(lb.ListView.Items, items...)
	return lb
}

func (lb *ListViewBuilder) Icon(icon string) *ListViewBuilder {
	lb.ListView.Icon = icon
	return lb
}

func (lb *ListViewBuilder) Item(item model.ListViewItem) *ListViewBuilder {
	lb.ListView.Items = append(lb.ListView.Items, item)
	return lb
}

func (lb *ListViewBuilder) AdditionalInfo(info ...model.ListViewInfo) *ListViewBuilder {
	lb.ListView.AdditionalInfo = append(lb.ListView.AdditionalInfo, info...)
	return lb
}
