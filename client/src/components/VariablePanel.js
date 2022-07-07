import styled from "styled-components"
import useVariable from "../hook/variableHook"
import { Breadcrumb, Typography, Card, Tag, Button } from "antd"
import VariableInput from "./VariableInput"
import plugin from "../api/plugin"

const ArgumentContainer = styled.div`
display: grid;
grid-auto-rows: max-content;
gap: 20px;
`

const PanelContainer = styled.div`
    height: 550px;
    overflow-y: scroll;
`

const Container = styled.div`
height: 100%;
`

const { Title } = Typography

function VariablePanel({ pluginName }) {
    const { variableList, setVariable } = useVariable(pluginName)
    const onSave = async () => {
        let variableMap = {}
        for (const variable of variableList) {
            variableMap = { ...variableMap, [variable.name]: variable.value }

        }

        try {
            const res = await plugin.setVariables(pluginName, variableMap)
            console.log(res);
        } catch (error) {
            console.log(error);
        }

    }
    return <Container>
        <Breadcrumb>
            <Breadcrumb.Item>{pluginName}</Breadcrumb.Item>
            <Breadcrumb.Item>Variable</Breadcrumb.Item>
        </Breadcrumb>
        <br />
        <Title level={2}>Variable</Title>
        <PanelContainer>
            <ArgumentContainer>
                {variableList.map(({ name, value, type }) => {
                    return <Card key={name}>
                        <Title level={4}>{name} <Tag color="green">{type}</Tag></Title>
                        <VariableInput type={type} value={value} onChange={value => setVariable(name, value)} />
                    </Card>
                })}
            </ArgumentContainer>
        </PanelContainer>
        <br />
        <Button type="primary" style={{ width: "100%" }} onClick={onSave}>save</Button>
    </Container>
}

export default VariablePanel