import styled from "styled-components"
import { Breadcrumb, Card, Typography, Tag, Divider, Button, Alert } from "antd"
import { LoadingOutlined } from '@ant-design/icons'
import useFunction, { ExecuteState } from "../hook/functionHook"
import VariableInput from "./VariableInput"
import plugin from "../api/plugin"
import { useState } from "react"

const Container = styled.div`
width: 100%;
height: calc(100vh - 64px - 40.5px - 40px);
overflow-y: scroll;
`

const ArgumentContainer = styled.div`
display: grid;
grid-auto-rows: max-content;
gap: 20px;
`

const OutputContainer = styled.div`
width: 100%;
border-radius:3px;
padding: 10px 30px;
background-color: #ecf0f1;
margin-bottom: 15px;
white-space: pre-wrap;
`

const { Title } = Typography

function FunctionPanel({ functionName, pluginName }) {
    const { executeState, args, returns, outputs, params, setParam } = useFunction(pluginName, functionName)
    const [errorMsg, setErrorMsg] = useState(null)
    const onExecute = async () => {
        for (const param of params) {
            if (param === null) return setErrorMsg("some parameter is empty")
        }
        setErrorMsg(null)
        await plugin.executeFunction(pluginName, functionName, params)
    }

    return <Container>
        <Breadcrumb>
            <Breadcrumb.Item>{pluginName}</Breadcrumb.Item>
            <Breadcrumb.Item>{functionName}</Breadcrumb.Item>
        </Breadcrumb>
        <br />
        <Title level={2}>{functionName}</Title>
        <Divider></Divider>
        <Title level={3}>Arguments</Title>
        <ArgumentContainer>
            {args.map((type, index) => {
                return <Card key={String(index)}>
                    <Title level={4}>Arg {index} <Tag color="green">{type}</Tag></Title>
                    <VariableInput key={String(index)} type={type} value={params[index]} onChange={(value) =>
                        setParam(index, value)
                    } />
                </Card>
            })}
        </ArgumentContainer>
        <br />
        <div style={{ display: "flex", justifyContent: "center" }}>
            <Button type="primary" style={{ width: "100%" }} onClick={onExecute}>Execute</Button>
        </div>
        <br />
        {errorMsg !== null && <Alert type="error" description={errorMsg} />}
        <br />
        <br />
        <Title level={3}>Return</Title>
        <ArgumentContainer>
            {returns.map((type, index) => {
                const isOutputAvailable = outputs[index] !== null
                const isRunning = executeState === ExecuteState.ExecuteRunning
                return <Card key={String(index)}>
                    <Title level={4}>Return {index} <Tag color="green">{type}</Tag></Title>
                    {isRunning && <LoadingOutlined />}
                    {!isRunning && isOutputAvailable && <OutputContainer>{typeof outputs[index] === "boolean" ? (outputs[index] ? "TRUE" : "FALSE") : outputs[index]}</OutputContainer>}
                    {!isRunning && isOutputAvailable && <Button type="primary" onClick={() => navigator.clipboard.writeText(String(outputs[index]))}>Copy</Button>}
                </Card>
            })}
        </ArgumentContainer>
    </Container>
}

export default FunctionPanel