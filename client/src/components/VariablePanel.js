import { PageHeader, Breadcrumb } from "antd"

function VariablePanel({ pluginName }) {
    return <>
        <Breadcrumb>
            <Breadcrumb.Item>{pluginName}</Breadcrumb.Item>
            <Breadcrumb.Item>Variable</Breadcrumb.Item>
        </Breadcrumb>
        <PageHeader
            title="Variable"
            subTitle="This is a subtitle"
        ></PageHeader>
    </>
}

export default VariablePanel