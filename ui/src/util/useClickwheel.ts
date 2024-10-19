import {useCallback, useEffect, useState} from "react";
import {fetchClickWheelData} from "./clickwheelListeners";
import {usePrevious} from "@chakra-ui/react";
import {useLocation} from "react-router-dom";
import {useMockClickwheel} from "./useMockClickwheel";

export type ClickWheelProps = {
    socket: WebSocket,
    maxClickWheelValue: number,
    onSelectButton?: (currentIndex: number) => void,
    onSelectButtonLongPress?: (currentIndex: number) => void,
    onMenuButton?: (currentIndex: number) => void,
}

export type ClickWheelResponse = {
    isPressed: boolean;
    setOnSelectButton:  React.Dispatch<React.SetStateAction<(...args: any[]) => void>>;
    setOnMenuButton:  React.Dispatch<React.SetStateAction<(...args: any[]) => void>>;
    selectedIndex: number;
    setOnSelectButtonLongPress:  React.Dispatch<React.SetStateAction<(...args: any[]) => void>>
}

export const useClickwheel = ({
                                  socket,
                                  onSelectButton,
                                  onSelectButtonLongPress,
                                  maxClickWheelValue,
                                  onMenuButton
                              }: ClickWheelProps,
): ClickWheelResponse => {
    const LONG_PRESS_THRESHOLD = 1000;
    const FALLBACK_FUNCTION = () => {};

    const [isPressed, setIsPressed] = useState(false);
    const [startPosition, setStartPosition] = useState(0);
    const [currentPosition, setCurrentPosition] = useState(0);
    const [selectedIndex, setSelectedIndex] = useState(0);
    const [longPressStart, setLongPressStart] = useState(0);
    const previousPosition = usePrevious(currentPosition);
    const {key} = useLocation();

    const [onSelectCallback, setOnSelectCallback] = useState(() => onSelectButton || FALLBACK_FUNCTION);
    const [onMenuCallback, setOnMenuCallback] = useState(() => onMenuButton || FALLBACK_FUNCTION);
    const [onLongSelectCallback, setOnLongSelectCallback] = useState(() => onSelectButtonLongPress || FALLBACK_FUNCTION);

    useMockClickwheel({
        onSelectButton: onSelectCallback,
        maxClickWheelValue: maxClickWheelValue,
        selectedIndex: selectedIndex,
        setSelectedIndex: setSelectedIndex,
        onMenuButton: onMenuCallback,
        onSelectButtonLongPress: onLongSelectCallback,
        key: key
    })

    useEffect(() => {
        if (isPressed) {
            setStartPosition(currentPosition);
        }
    }, [isPressed]);

    useEffect(() => {
        if (currentPosition <= 46 && currentPosition >= 2) {
            const isClockwise = currentPosition > previousPosition;

            if (isClockwise && selectedIndex + 1 >= maxClickWheelValue) {
                return
            } else if (!isClockwise && selectedIndex - 1 < 0) {
                return
            } else if (Math.abs(currentPosition - startPosition) >= 5) {
                setStartPosition(currentPosition);
                setSelectedIndex(selectedIndex + (isClockwise ? 1 : -1));
            }
        }
    }, [currentPosition]);

    const onMessageHandler = useCallback((e: MessageEvent) => {
        const clickWheelData = fetchClickWheelData(e);
        const isOnKeyUp = !clickWheelData.isClickWheelPressed;

        if (clickWheelData.button === "ClickWheel") {
            if (clickWheelData.isClickWheelPressed !== isPressed) {
                setIsPressed(clickWheelData.isClickWheelPressed);
            }
            setCurrentPosition(clickWheelData.clickWheelPosition);
        } else if (clickWheelData.button === "Select") {
            if (isOnKeyUp) {
                const timeSincePress = Date.now() - longPressStart
                if (timeSincePress > LONG_PRESS_THRESHOLD) {
                    if (onSelectButtonLongPress) {
                        onLongSelectCallback(selectedIndex);
                    }
                } else if (onSelectButton) {
                    onSelectCallback(selectedIndex);
                }
            } else {
                setLongPressStart(Date.now());
            }
        }
    }, [selectedIndex]);

    useEffect(() => {
        socket.addEventListener("message", onMessageHandler);

        setSelectedIndex(() => {
            return parseInt(localStorage.getItem(key + "-selectedIndex") || "0");
        })

        return () => {
            socket.removeEventListener("message", onMessageHandler);
        }
    }, [key]);

    useEffect(() => {
        localStorage.setItem(key + "-selectedIndex", selectedIndex.toString());
    }, [selectedIndex]);

    return {
        selectedIndex: selectedIndex,
        isPressed: isPressed,
        setOnSelectButton: setOnSelectCallback,
        setOnMenuButton: setOnMenuCallback,
        setOnSelectButtonLongPress: setOnLongSelectCallback
    };
}