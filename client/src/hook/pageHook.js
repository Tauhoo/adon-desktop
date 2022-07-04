import { useState } from "react"
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
    }

    const onRemoveVariablePage = (pluginName) => {
        const newPage = pages.filter(page => !(page.type === PageType.VARIABLE && page.pluginName === pluginName))
        setPages(newPage)
    }


    return { activePage, pages, selectFuctionPage, selectVariablePage, onRemoveFunctionPage, onRemoveVariablePage }
}

export default usePageHook