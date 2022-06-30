import { Menu } from "antd"
import { InboxOutlined } from '@ant-design/icons';
import { usePluginMenu } from "../hook/pluginHook";

function PluginMenu() {
    const value = usePluginMenu()
    const items = value === null ? [] : value.map(name => {
        return {
            key: name,
            icon: <InboxOutlined />,
            label: name,
            type: 'menu',
        }
    })

    return <>
        <Menu
            style={{ width: 300 }}
            defaultSelectedKeys={['1']}
            defaultOpenKeys={['sub1']}
            mode="inline"
            items={items}
        />
    </>
}

export default PluginMenu