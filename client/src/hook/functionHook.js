import { useEffect, useState } from "react";
import plugin from "../api/plugin";

function useFunction(pluginName, functionName) {
    const [args, setArgs] = useState([])
    const [returns, setReturns] = useState([])
    const [outputs, setOutput] = useState([])

    const [params, setParams] = useState([])

    const initFunction = async () => {
        const func = await plugin.getFunction(pluginName, functionName)
        setArgs(func.data.arg_types)
        setReturns(func.data.return_types)
        setParams(new Array(func.data.arg_types.length))
    }

    useEffect(() => {
        initFunction()
    }, [])

    const setParam = (index, value) => {
        params[index] = value
        setParams([...params])
    }

    return { args, returns, outputs, setParams }
}

export default useFunction