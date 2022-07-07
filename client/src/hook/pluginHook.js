import { useEffect, useState } from "react";
import plugin from "../api/plugin";

export function usePluginMenu() {
    const [pluginMenu, setPluginMenu] = useState([])

    const initPluginMenu = async () => {
        try {
            const result = await plugin.getPluginNameList()
            setPluginMenu(result.data.sort())
        } catch (error) {
            console.log(error);
        }
    }

    const onPluginAdded = (message) => {
        setPluginMenu([...pluginMenu, message.data].sort())
    }

    const onPluginDeleted = (message) => {
        setPluginMenu(pluginMenu.filter(value => value !== message.data))
    }

    useEffect(() => {
        initPluginMenu()
    }, [])

    useEffect(() => {
        const clearOnPluginAdded = plugin.onPluginAdded(onPluginAdded)
        const clearOnPluginDeleted = plugin.onPluginDeleted(onPluginDeleted)
        return () => {
            clearOnPluginAdded()
            clearOnPluginDeleted()
        }
    }, [pluginMenu])

    return pluginMenu
}

export function useFunctionMenu(name) {
    const [functionMenu, setFunctionMenu] = useState([])

    const initFunctionMenu = async () => {
        try {
            const result = await plugin.getFunctionList(name)
            setFunctionMenu(result.data)
        } catch (error) {
            console.log(error);
        }
    }


    useEffect(() => {
        initFunctionMenu()
    }, [])

    return functionMenu
}