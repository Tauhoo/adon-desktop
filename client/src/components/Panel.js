import { PageHeader } from "antd"
import styled from "styled-components"

const Container = styled.div`
padding: 20px 60px;
`

const funcInfo = {
    dataDataStatus: "LOADING|ERROR",
    name: "",
    type: "function",
    params: [{
        type: "",
        value: ""
    }],
    returns: [{
        type: "",
        value: ""
    }],
    functionStatus: "DONE|RUNNING",
}

const variableInfo = {
    dataDataStatus: "LOADING|ERROR",
    type: "variable",
    variable: [{
        type: "string",
        value: "",
        name: "",
    }],
}

function PluginPanel() {
    return <Container>
        <PageHeader
            title="Function name"
            subTitle="This is a subtitle"
        ></PageHeader>
    </Container>
}

export default PluginPanel