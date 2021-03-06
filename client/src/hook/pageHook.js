import { message } from "antd"
import { useEffect, useState } from "react"
import plugin from "../api/plugin"
export const PageType = {
    FUNCTION: "FUNCTION",
    VARIABLE: "VARIABLE"
}

function usePageHook() {
    const [pages, setPages] = useState([])
    const [activePage, setActivePage] = useState(null)

    const selectFuctionPage = (pluginName, functionName) => {
        for (const page of pages) {
            if (pluginName === page.pluginName && functionName === page.functionName && page.type === PageType.FUNCTION) {
                setActivePage(page)
                return
            }
        }

        const page = {
            type: PageType.FUNCTION,
            pluginName,
            functionName
        }

        setPages([...pages, page])
        setActivePage(page)
    }

    const selectVariablePage = pluginName => {
        for (const page of pages) {
            if (pluginName === page.pluginName && page.type === PageType.VARIABLE) {
                setActivePage(page)
                return
            }
        }

        const page = {
            type: PageType.VARIABLE,
            pluginName,
        }

        setPages([...pages, page])
        setActivePage(page)
    }

    const onRemoveFunctionPage = (pluginName, functionName) => {
        const newPage = pages.filter(page => !(page.type === PageType.FUNCTION && page.pluginName === pluginName && page.functionName === functionName))
        setPages(newPage)
        if (newPage.length === 0) {
            setActivePage(null)
        } else if (activePage.type === PageType.FUNCTION && activePage.pluginName === pluginName && activePage.functionName === functionName) {
            setActivePage(newPage[0])
        }
    }

    const onRemoveVariablePage = (pluginName) => {
        const newPage = pages.filter(page => !(page.type === PageType.VARIABLE && page.pluginName === pluginName))
        setPages(newPage)
        if (newPage.length === 0) {
            setActivePage(null)
        } else if (activePage.type === PageType.VARIABLE && activePage.pluginName === pluginName) {
            setActivePage(newPage[0])
        }
    }

    const onRemovePlugin = message => {
        const pluginName = message.data
        const newPage = pages.filter(page => page.pluginName !== pluginName)
        setPages(newPage)
        if (newPage.length === 0) {
            setActivePage(null)
        } else if (activePage.pluginName === pluginName) {
            setActivePage(newPage[0])
        }
    }

    useEffect(() => {
        return plugin.onPluginDeleted(onRemovePlugin)
    })

    return { activePage, pages, selectFuctionPage, selectVariablePage, onRemoveFunctionPage, onRemoveVariablePage }
}

export default usePageHook