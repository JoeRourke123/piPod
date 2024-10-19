export type ListViewProps = {
    title: string
    items: ListViewItemDetails[]
    showStatus: boolean
};

export type ListViewItemDetails = {
    title: string,
    actionType?: "REDIRECT" | "POST" | "GET",

    actions?: ListViewItemDetails[],

    path?: string,
    requestUrl?: string,
    toastMessage?: string
}