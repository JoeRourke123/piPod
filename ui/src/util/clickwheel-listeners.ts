export type ClickWheelData = {
    button: string
    isClickWheelPressed: boolean
    clickWheelPosition: number
}

export const fetchClickWheelData = (e: MessageEvent): ClickWheelData | undefined => {
    if (e.data && e.data != "null") {
        const dataJson = JSON.parse(e.data);

        if (dataJson["button"]) {
            return {
                button: dataJson["button"],
                isClickWheelPressed: dataJson["is_click_wheel_pressed"],
                clickWheelPosition: dataJson["click_wheel_position"]
            }
        }
    }
}