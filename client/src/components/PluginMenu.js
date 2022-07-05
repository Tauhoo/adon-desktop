import { Button, Collapse, Divider } from 'antd';
import { MoreOutlined } from '@ant-design/icons'
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

const MenuButtonContainer = styled.div`
padding: 10px 30px;
border-style: solid;
border-color:rgb(235, 237, 240) ;
border-width: 0px 0px 1px 0px;
`

const NameContainer = styled.div`
width: 244px;
overflow: hidden;
white-space: nowrap;
overflow: hidden;
text-overflow: ellipsis;
`

function PluginMenu({ onClickFunction, onClickVariable, onDeletePlugin }) {
    const value = usePluginMenu()
    return <Container>
        <MenuButtonContainer>
            <CreatePluginPanelModal></CreatePluginPanelModal>
        </MenuButtonContainer>
        <Collapse ghost >
            {value.map(name => {
                return <Panel header={<NameContainer>{name}</NameContainer>} key={name}>
                    <Button type='dashed' onClick={() => onDeletePlugin(name)}>Delete Plugin</Button>
                    <br /><br />
                    <Button type='text' style={{ width: "100%", textAlign: "start" }} onClick={() => { if (onClickVariable) onClickVariable(name) }}>Variable</Button>
                    <Divider style={{ margin: "5px" }} />
                    <FunctionMenu pluginName={name} onClick={(funcName) => { if (onClickFunction) onClickFunction(name, funcName) }}></FunctionMenu>
                </Panel>
            })}
        </Collapse>
    </Container>
}

export default PluginMenu