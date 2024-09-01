import {PageProps} from "../state/PageProps";
import {Center} from "@chakra-ui/react";
import {QR} from "react-qr-rounded";
import {useLocation, useNavigate} from "react-router-dom";
import {useEffect} from "react";
import {fetchAuthStatus} from "../util/service";

export const Auth = (props: PageProps) => {
    const location = useLocation();
    const navigate = useNavigate();

    useEffect(() => {
        setInterval(async () => {
            const authStatus = await fetchAuthStatus();

            if (authStatus["has_token"]) {
                navigate("/", {
                    replace: true
                });
            }
        }, 1000);
    }, []);

    return (
        <Center width="100%" textAlign="center" height="100vh">
            <QR color="#0BC5EA" width="180px" height="180px" rounding={100}>{ location.state?.authUrl || "http://localhost:9091/go" }</QR>
        </Center>
    )
}