export const handleKeyUp = (
    maxClickwheelValue: number,
    setSelectedIndex: any,
    onSelectButton: any,
    onMenuButton: any,
    onSelectButtonLongPress: any,
    selectedIndex: number
) => {
    return (e: KeyboardEvent) => {
        const key = e.key;
        console.log(key);

        if (key === "ArrowUp") {
            setSelectedIndex((index: number) => Math.max(0, index - 1));
        } else if (key === "ArrowDown") {
            setSelectedIndex((index: number) => Math.min((index + 1), maxClickwheelValue - 1));
        } else if (key === "Enter") {
            onSelectButton(selectedIndex);
        } else if (key === "w") {
            onMenuButton(selectedIndex);
        } else if (key === "q") {
            onSelectButtonLongPress(selectedIndex);
        }
    };
}