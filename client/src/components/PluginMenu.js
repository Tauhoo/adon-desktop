import { Menu } from "antd"
import { InboxOutlined } from '@ant-design/icons';
import { usePluginMenu } from "../hook/pluginHook";
import styled from "styled-components";
import { CreatePluginPanelModal } from "./CreatePluginPanel"


const Container = styled.div`
width: 300px;
height: 100%;
display: grid;
grid-template-rows: max-content 1fr;
border-style: solid;
border-color:rgb(235, 237, 240) ;
border-width: 0px 1px 0px 0px;
overflow: hidden;

`

function PluginMenu() {
    const value = usePluginMenu()
    const items = value.map(name => {
        return {
            key: name,
            icon: <InboxOutlined />,
            label: name,
            type: 'menu',
        }
    })

    return <Container>
        <CreatePluginPanelModal></CreatePluginPanelModal>
        <Menu
            items={[{
                key: "plugin-list",
                children: items,
                label: "Plugin list",
                type: 'group',
            }]}
            onClick={console.log}
            style={{ border: "none", width: 300 }}
        />
    </Container>
}

export default PluginMenu