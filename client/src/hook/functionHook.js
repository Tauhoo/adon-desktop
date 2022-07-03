import { useEffect, useState } from "react";
import plugin from "../api/plugin";

function useFunction(pluginName, functionName) {
    const [args, setArgs] = useState([])
    const [returns, setReturns] = useState([])
    const [outputs, setOutput] = useState([])

    const initFunction = async () => {
        const func = await plugin.getFunction(pluginName, functionName)
        setArgs(func.data.arg_types)
        setReturns(func.data.return_types)
    }

    useEffect(() => {
        initFunction()
    }, [])

    return { args, returns, outputs }
}

export default useFunction