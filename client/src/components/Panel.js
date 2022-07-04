import { Tabs, Tag } from "antd"
import { CloseOutlined } from '@ant-design/icons'
import styled from "styled-components"
import { PageType } from "../hook/pageHook"
import FunctionPanel from "./FunctionPanel"
import VariablePanel from "./VariablePanel"

const { TabPane } = Tabs

const Container = styled.div`
padding: 20px 30px;
width: calc(100vw - 300px);
`

function getTabFromPage(page) {
    const name = getTabNameFromPage(page)
    switch (page.type) {
        case PageType.FUNCTION:
            return <><Tag color="blue">Func</Tag> {name}</>
        case PageType.VARIABLE:
            return <><Tag color="orange">Var</Tag> {name}</>
        default:
            return ""
    }
}

function getTabNameFromPage(page) {
    switch (page.type) {
        case PageType.FUNCTION:
            return `${page.functionName} - ${page.pluginName}`
        case PageType.VARIABLE:
            return `${page.pluginName}`
        default:
            return ""
    }
}

function getTabKeyFromPage(page) {
    switch (page.type) {
        case PageType.FUNCTION:
            return `${PageType.FUNCTION}/${page.pluginName}/${page.functionName}`
        case PageType.VARIABLE:
            return `${PageType.VARIABLE}/${page.pluginName}`
        default:
            return ""
    }
}

function PluginPanel({ activePage, pages, selectFuctionPage, selectVariablePage, onRemoveFunctionPage, onRemoveVariablePage }) {
    if (activePage === null) {
        return null
    }

    const onSelect = key => {
        const keyData = key.split("/")
        switch (keyData[0]) {
            case PageType.FUNCTION:
                return selectFuctionPage(keyData[1], keyData[2])
            case PageType.VARIABLE:
                return selectVariablePage(keyData[1])
            default:
                return
        }
    }

    const onEdit = (key, action) => {
        if (action !== "remove") return
        const keyData = key.split("/")
        switch (keyData[0]) {
            case PageType.FUNCTION:
                return onRemoveFunctionPage(keyData[1], keyData[2])
            case PageType.VARIABLE:
                return onRemoveVariablePage(keyData[1])
            default:
                return
        }
    }

    const isPagesAvailable = pages.length !== 0
    if (!isPagesAvailable) {
        return <Container />
    }

    return <Container>
        <Tabs hideAdd={true} activeKey={getTabKeyFromPage(activePage)} onTabClick={onSelect} type="editable-card" onEdit={onEdit}>
            {pages.map((page) => {
                const tab = getTabFromPage(page)
                const key = getTabKeyFromPage(page)
                switch (page.type) {
                    case PageType.FUNCTION:
                        return <TabPane tab={tab} key={key} closable={true}>
                            <FunctionPanel functionName={page.functionName} pluginName={page.pluginName} />
                        </TabPane>
                    case PageType.VARIABLE:
                        return <TabPane tab={tab} key={key} closable={true}>
                            <VariablePanel pluginName={page.pluginName} />
                        </TabPane>
                    default:
                        return null
                }
            })}
        </Tabs>
    </Container>
}

export default PluginPanel