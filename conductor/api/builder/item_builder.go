package builder

import "conductor/common/model"

type ItemBuilder struct {
	model.ListViewItem
}

func ListViewItem() *ItemBuilder {
	return &ItemBuilder{}
}

func (ib *ItemBuilder) Build() model.ListViewItem {
	return ib.ListViewItem
}

func (ib *ItemBuilder) Title(title string) *ItemBuilder {
	ib.ListViewItem.Title = title
	return ib
}

func (ib *ItemBuilder) Path(path string) *ItemBuilder {
	ib.ListViewItem.Path = path
	return ib
}

func (ib *ItemBuilder) BackgroundImage(backgroundImage string) *ItemBuilder {
	ib.ListViewItem.BackgroundImage = backgroundImage
	return ib
}

func (ib *ItemBuilder) Icon(icon string) *ItemBuilder {
	ib.ListViewItem.Icon = icon
	return ib
}

func (ib *ItemBuilder) Subtitle(subtitle string) *ItemBuilder {
	ib.ListViewItem.Subtitle = subtitle
	return ib
}

func (ib *ItemBuilder) Disabled(disabled bool) *ItemBuilder {
	ib.ListViewItem.Disabled = disabled
	return ib
}

func (ib *ItemBuilder) Action(action model.ListViewItem) *ItemBuilder {
	ib.ListViewItem.Actions = append(ib.ListViewItem.Actions, action)
	return ib
}

func (ib *ItemBuilder) Actions(actions []model.ListViewItem) *ItemBuilder {
	ib.ListViewItem.Actions = append(ib.ListViewItem.Actions, actions...)
	return ib
}

func (ib *ItemBuilder) ActionType(actionType string) *ItemBuilder {
	ib.ListViewItem.ActionType = actionType
	return ib
}

func (ib *ItemBuilder) RequestUrl(requestUrl string) *ItemBuilder {
	ib.ListViewItem.RequestUrl = requestUrl
	return ib
}

func (ib *ItemBuilder) ToastMessage(toastMessage string) *ItemBuilder {
	ib.ListViewItem.ToastMessage = toastMessage
	return ib
}
