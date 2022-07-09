import { useEffect, useState } from "react";
import plugin from "../api/plugin";

export const ExecuteState = {
    ExecuteIdle: "IDLE",
    ExecuteError: "ERROR",
    ExecuteRunning: "RUNNING",
    ExecuteDone: "DONE",
}

function useFunction(pluginName, functionName) {
    const [args, setArgs] = useState([])
    const [returns, setReturns] = useState([])
    const [outputs, setOutput] = useState([])
    const [executeState, setExecuteState] = useState(ExecuteState.ExecuteIdle)

    const [params, setParams] = useState([])

    const initFunction = async () => {
        const func = await plugin.getFunction(pluginName, functionName)
        setArgs(func.data.arg_types)
        setReturns(func.data.return_types)
        setParams(new Array(func.data.arg_types.length).fill(null))
        setOutput(new Array(func.data.return_types.length).fill(null))
    }

    useEffect(() => {
        initFunction()
        return plugin.onExecuteStateChange(({ data }) => {
            if (data.state != ExecuteState.ExecuteDone || data.plugin_name !== pluginName || data.function_name !== functionName) return
            setExecuteState(data.state)
            setOutput(data.info)
        })
    }, [])

    const setParam = (index, value) => {
        params[index] = value
        setParams([...params])
    }

    return { executeState, args, returns, outputs, params, setParam }
}

export default useFunction