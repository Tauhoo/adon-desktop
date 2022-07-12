import styled from "styled-components"
import useVariable from "../hook/variableHook"
import { Breadcrumb, Typography, Card, Tag, Button } from "antd"
import VariableInput from "./VariableInput"
import plugin from "../api/plugin"
import { useState } from "react"

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
height: calc(100vh - 64px - 40.5px - 40px);
`

const { Title } = Typography

function VariablePanel({ pluginName }) {
    const { variableList, setVariable } = useVariable(pluginName)
    const [loading, setLoading] = useState(false)
    const onSave = async () => {
        setLoading(true)

        let variableMap = {}
        for (const variable of variableList) {
            variableMap = { ...variableMap, [variable.name]: variable.value }
        }

        try {
            const res = await plugin.setVariables(pluginName, variableMap)
        } catch (error) {
            console.log(error);
        }

        setLoading(false)
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
        <Button type="primary" style={{ width: "100%" }} onClick={onSave} loading={loading}>save</Button>
    </Container>
}

export default VariablePanel