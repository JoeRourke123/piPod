import {useToast} from "@chakra-ui/toast";

const useErrorToast = () => {
    const toast = useToast();

    return () => toast({
        title: "An error occurred",
        position: "bottom-left",
        status: "error",
        isClosable: false,
        duration: 1000,
    });
}

const useSuccessToast = () => {
    const toast = useToast();

    return (toastMessage?: string) => toast({
        title: toastMessage || "Done!",
        position: "bottom-left",
        isClosable: false,
        duration: 1000
    });
}

export const useItemActionToasts = () => [useErrorToast(), useSuccessToast()];