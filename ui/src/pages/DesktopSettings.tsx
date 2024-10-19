import {PageProps} from "../state/PageProps";
import {Box, Button, Container, Flex, FormControl, FormLabel, Heading, Input, Select, VStack} from "@chakra-ui/react";

export const DesktopSettings = (props: PageProps) => {
    return <VStack p="6">
        <Container w="full">
            <Heading my="6">Settings</Heading>
            <Box my="6" borderWidth='1px' borderRadius='lg' p="4">
                <FormControl>
                    <FormLabel htmlFor='country'>WiFi</FormLabel>
                    <Flex mt="4" gap="2" flexDirection="column">
                        <Select id='wifi' placeholder='Select network'>
                            <option>United Arab Emirates</option>
                            <option>Nigeria</option>
                        </Select>
                        <Input type="password" placeholder='Enter password' />
                        <Button>Connect</Button>
                    </Flex>
                </FormControl>
            </Box>

            <Box my="6" borderWidth='1px' borderRadius='lg' p="4">
                <FormControl>
                    <FormLabel htmlFor='country'>Bluetooth</FormLabel>
                    <Flex mt="4" gap="2" flexDirection="column">
                        <Select id='wifi' placeholder='Select network'>
                            <option>United Arab Emirates</option>
                            <option>Nigeria</option>
                        </Select>
                        <Button>Connect</Button>
                    </Flex>
                </FormControl>
            </Box>

            <Box my="6" borderWidth='1px' borderRadius='lg' p="4">
                <FormControl>
                    <FormLabel htmlFor='country'>Power</FormLabel>
                    <Flex mt="4" gap="2" flexDirection="column">
                        <Button>Restart PiPod</Button>
                        <Button>Update Software</Button>
                    </Flex>
                </FormControl>
            </Box>
        </Container>
    </VStack>
}