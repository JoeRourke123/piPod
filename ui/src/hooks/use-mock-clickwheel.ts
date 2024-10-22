import {useEffect} from "react";
import {handleKeyUp} from "../util/handle-key-up";
import {useBeforeUnload} from "react-router-dom";

export type MockClickwheelProps = {
    maxClickWheelValue: number,
    setSelectedIndex: any,
    selectedIndex: number,
    onSelectButton?: (currentIndex: number) => void,
    onSelectButtonLongPress?: (currentIndex: number) => void,
    onMenuButton?: (currentIndex: number) => void,
    key: string
}

export const useMockClickwheel = ({
                                      onSelectButton,
                                      maxClickWheelValue,
                                      setSelectedIndex,
                                      selectedIndex,
                                      key,
                                      onMenuButton,
                                      onSelectButtonLongPress
                                  }: MockClickwheelProps) => {
    useEffect(() => {
        document.onkeyup = handleKeyUp(maxClickWheelValue, setSelectedIndex, onSelectButton, onMenuButton, onSelectButtonLongPress, selectedIndex);
    }, [key, onMenuButton, onSelectButton, onSelectButtonLongPress]);

    useBeforeUnload(() => {
        document.onkeyup = null;
    })

    useEffect(() => {
        document.onkeyup = handleKeyUp(maxClickWheelValue, setSelectedIndex, onSelectButton, onMenuButton, onSelectButtonLongPress, selectedIndex);
    }, [selectedIndex]);
}