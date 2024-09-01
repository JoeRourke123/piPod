import {ListViewItemProps} from "./ListViewTypes";

export const handleKeyUp = (setSelectedIndex: any, listViewItems: ListViewItemProps[], navigate: any) => {
    return (e: KeyboardEvent) => {
        const key = e.key;
        console.log(key);

        if (key === "ArrowUp") {
            setSelectedIndex((index: number) => Math.max(0, index - 1));
        } else if (key === "ArrowDown") {
            setSelectedIndex((index: number) => Math.min((index + 1), listViewItems.length - 1));
        } else if (key === "Enter") {
            setSelectedIndex((index: number) => {
                const selectedPath = listViewItems[index].path;

                if (selectedPath) {
                    navigate(selectedPath);
                    return 0;
                }

                return index;
            });
        } else if (key === "w") {
            navigate(-1);
        }
    };
}