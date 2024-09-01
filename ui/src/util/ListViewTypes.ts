import React from "react";

export type ListViewProps = {
    title?: string
    customTitle?: React.ReactElement
    items: ListViewItemProps[]
    showStatus: boolean
    fallbackIcon? : (color: string) => React.ReactElement
    pageLoader?: (current_offset: number) => Promise<ListViewItemProps[]>
};

export type ListViewItemProps = {
    title: string,
    icon?: (color: string) => React.ReactElement
    path?: string
}