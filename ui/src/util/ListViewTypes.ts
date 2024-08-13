import React from "react";

export type ListViewProps = {
    title: string
    items: ListViewItemProps[]
    fallbackIcon? : (color: string) => React.ReactElement
};

export type ListViewItemProps = {
    title: string,
    icon?: (color: string) => React.ReactElement
    path?: string
}