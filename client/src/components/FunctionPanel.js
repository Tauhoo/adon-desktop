import styled from "styled-components"
import { PageHeader, Breadcrumb, Card, Typography, Tag, Divider } from "antd"
import useFunction from "../hook/functionHook"
import VariableInput from "./VariableInput"

const ArgumentContainer = styled.div`
display: grid;
grid-auto-rows: max-content;
gap: 20px;
`

const { Title } = Typography

function FunctionPanel({ functionName, pluginName }) {
    const { args, returns, outputs, setArg } = useFunction(pluginName, functionName)
    return <>
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
                return <Card>
                    <Title level={4}>Arg {index} <Tag color="green">{type}</Tag></Title>
                    <VariableInput key={String(index)} type={type} onChange={(value) => setArg(index, value)} />
                </Card>
            })}
        </ArgumentContainer>
        <br />
        <Title level={3}>Return</Title>
    </>
}

export default FunctionPanel