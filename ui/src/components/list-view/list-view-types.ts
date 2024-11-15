export type ListViewProps = {
    title: string;
    items: ListViewItemDetails[];
    showStatus: boolean;
    icon?: string;
    additionalInfo: AdditionalListViewInfo[];
};

export type ListViewItemDetails = {
    title: string;
    subtitle?: string;
    icon?: string;
    backgroundImage?: string;
    disabled?: boolean;

    actions?: ListViewItemDetails[];

    actionType?: "REDIRECT" | "POST" | "GET";
    path?: string;
    requestUrl?: string;
    toastMessage?: string;
}

export type AdditionalListViewInfo = {
    text: string;
    icon: string;
    bold: boolean;
}