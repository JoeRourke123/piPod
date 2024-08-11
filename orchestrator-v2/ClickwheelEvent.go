package main

type ClickWheelEvent struct {
	Button              string `json:"button"`
	IsClickWheelPressed bool   `json:"is_click_wheel_pressed"`
	ClickwheelPosition  int    `json:"click_wheel_position"`
}

func BuildClickWheelEvent(previous *ClickWheelEvent, buttonId int, clickWheelFlag int, clickwheelPosition int) *ClickWheelEvent {
	buttonMap := map[int]string{
		29:  "ClickWheel",
		7:   "Select",
		8:   "Skip",
		9:   "Back",
		10:  "Play",
		11:  "Menu",
		255: previous.Button,
	}

	button := buttonMap[buttonId]

	var isClickWheelPressed bool
	if clickWheelFlag > 1 {
		isClickWheelPressed = previous.IsClickWheelPressed
	} else {
		isClickWheelPressed = clickWheelFlag == 1
	}

	return &ClickWheelEvent{
		Button:              button,
		IsClickWheelPressed: isClickWheelPressed,
		ClickwheelPosition:  clickwheelPosition,
	}
}
