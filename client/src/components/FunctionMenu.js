import { Button, Tag } from "antd"
import styled from "styled-components"
import { useFunctionMenu } from "../hook/pluginHook"

const Container = styled.div``
const FunctionContainer = styled.div``

function FunctionMenu({ pluginName, onClick }) {
    const functionMenu = useFunctionMenu(pluginName)
    return <Container>
        {functionMenu.map(({ name }) => {
            return <FunctionContainer key={name} onClick={() => {
                if (onClick) onClick(name)
            }}>
                <Button type="text" style={{ width: "267px", textAlign: "start", overflow: "hidden" }}> {name} </Button>
            </FunctionContainer>
        })}
    </Container>
}

export default FunctionMenu