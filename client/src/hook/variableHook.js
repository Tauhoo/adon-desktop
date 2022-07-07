import { useEffect, useState } from "react";
import plugin from "../api/plugin";

function useVariable(pluginName) {
    const [variableList, setVariableList] = useState([])

    const setVariable = (name, value) => {
        const newVariableList = []
        for (let variable of variableList) {
            if (variable.name === name) {
                newVariableList.push({ ...variable, value })
            } else {
                newVariableList.push(variable)
            }
        }


        setVariableList(newVariableList.sort((a, b) => {
            const sorted = [a.name, b.name].sort()
            return sorted[0] === a.name ? -1 : 1
        }))
    }

    const initVariableList = async () => {
        try {
            const result = await plugin.getVariableList(pluginName)
            setVariableList(result.data.sort((a, b) => {
                const sorted = [a.name, b.name].sort()
                return sorted[0] === a.name ? -1 : 1
            }))
        } catch (error) {
            console.log(error);
        }
    }

    useEffect(() => {
        initVariableList()
    }, [])

    return { variableList, setVariable }
}

export default useVariable