package main

type ClickWheelEvent struct {
	button              string `json:"button"`
	isClickWheelPressed bool   `json:"is_click_wheel_pressed"`
	clickwheelPosition  int    `json:"click_wheel_position"`
}

func BuildClickWheelEvent(previous ClickWheelEvent, buttonId int, clickWheelFlag int, clickwheelPosition int) ClickWheelEvent {
	buttonMap := map[int]string{
		29:  "ClickWheel",
		7:   "Select",
		8:   "Skip",
		9:   "Back",
		10:  "Play",
		11:  "Menu",
		255: previous.button,
	}

	button := buttonMap[buttonId]

	var isClickWheelPressed bool
	if clickWheelFlag > 1 {
		isClickWheelPressed = previous.isClickWheelPressed
	} else {
		isClickWheelPressed = clickWheelFlag == 1
	}

	return ClickWheelEvent{
		button:              button,
		isClickWheelPressed: isClickWheelPressed,
		clickwheelPosition:  clickwheelPosition,
	}
}
