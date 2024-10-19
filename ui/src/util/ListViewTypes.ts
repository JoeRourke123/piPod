import React from "react";

export type CategoryListProps = {
    data?: any
}

export type ListViewProps = {
    title: string
    items: ListViewItemDetails[]
    showStatus: boolean
    fallbackIcon? : (color: string) => React.ReactElement
    pageLoader?: (current_offset: number) => Promise<ListViewItemDetails[]>
    onSelectButton?: ((currentlySelected: (ListViewItemDetails)) => void),
};

export type ListViewItemDetails = {
    title: string,
    actionType?: "REDIRECT" | "POST" | "GET",

    actions?: ListViewItemDetails[],

    path?: string,
    requestUrl?: string,
    toastMessage?: string
}