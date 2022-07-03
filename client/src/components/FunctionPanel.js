import { PageHeader, Breadcrumb } from "antd"
import useFunction from "../hook/functionHook"

function FunctionPanel({ functionName, pluginName }) {
    const { args, returns, outputs } = useFunction(pluginName, functionName)
    return <>
        <Breadcrumb>
            <Breadcrumb.Item>{pluginName}</Breadcrumb.Item>
            <Breadcrumb.Item>{functionName}</Breadcrumb.Item>
        </Breadcrumb>
        <PageHeader
            title={functionName}
            subTitle="This is a subtitle"
        ></PageHeader>
        {JSON.stringify(args)}
        {JSON.stringify(returns)}
        {JSON.stringify(outputs)}
    </>
}

export default FunctionPanel