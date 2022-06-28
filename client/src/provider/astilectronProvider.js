import React, { createContext, useContext, useEffect, useState } from "react"

const AstilectronContext = createContext({})

export const useAstilectron = () => useContext(AstilectronContext)

export const AstilectronProvider = ({ children }) => {
    const [astilectron, setAstilectron] = useState(null)

    const onReady = () => {
        setAstilectron(window.astilectron)
    }

    useEffect(() => {
        document.addEventListener('astilectron-ready', onReady)
        return () => document.removeEventListener("astilectron-ready", onReady)
    }, [])

    return (
        <AstilectronContext.Provider value={{ astilectron }}>
            {children}
        </AstilectronContext.Provider>
    )
}