import {useParams} from "react-router-dom";
import {getApiUrl} from "../../../util/functions";

export const useApiUrl = (apiEndpoint?: string): string | undefined => {
    const params = useParams();

    if (!apiEndpoint) {
        return undefined;
    }

    return getApiUrl(apiEndpoint, params);
}