import { useEffect, useState } from "react";
import plugin from "../api/plugin";

export function usePluginMenu() {
    const [pluginMenu, setPluginMenu] = useState([])
    const initPluginMenu = async () => {
        try {
            const result = await plugin.getPluginNameList()
            setPluginMenu(result.data)
        } catch (error) {
            console.log(error);
        }
    }
    useEffect(() => {
        initPluginMenu()
    }, [])

    return pluginMenu
}
