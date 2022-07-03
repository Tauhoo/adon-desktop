import { Button, Collapse, Divider } from 'antd';
import { usePluginMenu } from "../hook/pluginHook";
import styled from "styled-components";
import { CreatePluginPanelModal } from "./CreatePluginPanel"
import FunctionMenu from './FunctionMenu';

const { Panel } = Collapse;

const Container = styled.div`
width: 300px;
height: calc(100vh - 70px);
display: grid;
grid-template-rows: max-content 1fr;
border-style: solid;
border-color:rgb(235, 237, 240) ;
border-width: 0px 1px 0px 0px;
overflow-y: scroll;
`

function PluginMenu({ onClickFunction, onClickVariable }) {
    const value = usePluginMenu()
    return <Container>
        <CreatePluginPanelModal></CreatePluginPanelModal>
        <div style={{ width: "100%", height: "100%", overflowY: "scroll" }}>
            <Collapse ghost >
                {value.map(name => {
                    return <Panel header={name} key={name}>
                        <Button type='text' style={{ width: "100%", textAlign: "start" }} onClick={() => { if (onClickVariable) onClickVariable(name) }}>Variable</Button>
                        <Divider style={{ margin: "5px" }} />
                        <FunctionMenu pluginName={name} onClick={(funcName) => { if (onClickFunction) onClickFunction(name, funcName) }}></FunctionMenu>
                    </Panel>
                })}
            </Collapse>
        </div>
    </Container>
}

export default PluginMenu