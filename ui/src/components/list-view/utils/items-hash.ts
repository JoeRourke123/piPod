import {ListViewItemDetails} from "../../../util/ListViewTypes";
import {useState} from "react";
import {stringHash} from "../../../util/functions";

export const useItemsHash = (items: ListViewItemDetails[]): [string, (items: ListViewItemDetails[]) => void] => {
    const generateHash = (items: ListViewItemDetails[]): string => {
        return stringHash(items.map(item => item.title).join(","));
    }

    const [hash, setHash] = useState<string>(generateHash(items));

    return [hash, (items: ListViewItemDetails[]) => {
        setHash(generateHash(items));
    }];
}