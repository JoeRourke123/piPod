import {useLocation, useParams} from "react-router-dom";
import {getApiUrl} from "../../../util/functions";

export const useApiUrl = (): string | undefined => {
    const params = useParams();
    const location = useLocation();
    return getApiUrl(location.pathname, params);
}