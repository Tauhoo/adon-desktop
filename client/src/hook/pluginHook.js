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

    useEffect(() => {
        initPluginMenu()
        return plugin.onPluginAdded(onPluginAdded)
    }, [])

    return pluginMenu
}
