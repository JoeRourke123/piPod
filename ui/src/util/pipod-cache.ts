import {ListViewProps} from "../components/list-view/list-view-types";
import {unmarshallView} from "../components/list-view/utils/view-loader";

export class PipodCache {
    public static selectedIndex = {
        set(k: string, si: number) {
            localStorage.setItem(k + "-selectedIndex", si.toString());
        },
        get(k: string) {
            return parseInt(localStorage.getItem(k + "-selectedIndex") || "0");
        }
    }

    public static view = {
        set(k: string, items: ListViewProps) {
            localStorage.setItem(k + "-items", JSON.stringify(items));
        },
        get(k: string): ListViewProps | undefined {
            const listViewStr = localStorage.getItem(k + "-items")
            if (!listViewStr) {
                return undefined;
            }

            const listViewJson = JSON.parse(listViewStr);
            return unmarshallView(listViewJson);
        }
    }

    public static clear(key: string) {
        localStorage.removeItem(key + "-selectedIndex");
        localStorage.removeItem(key + "-items");
    }
}