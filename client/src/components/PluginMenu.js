import { Menu } from "antd"
import { MailOutlined } from '@ant-design/icons';
const items = [
    {
        key: 'sub1',
        icon: <MailOutlined />,
        children: [
            {
                key: 'sub1',
                icon: <MailOutlined />,
                children: null,
                label: "Navigation One 1",
                type: 'item',
            },
            {
                key: 'sub1',
                icon: <MailOutlined />,
                children: null,
                label: "Navigation One 2",
                type: 'item',
            }
        ],
        label: "Navigation One",
        type: 'submenu',
    },
    {
        key: 'sub2',
        icon: <MailOutlined />,
        children: [
            {
                key: 'sub1',
                icon: <MailOutlined />,
                children: null,
                label: "Navigation One 1",
                type: 'item',
            },
            {
                key: 'sub1',
                icon: <MailOutlined />,
                children: null,
                label: "Navigation One 2",
                type: 'item',
            }
        ],
        label: "Navigation One",
        type: 'submenu',
    }
]

function PluginMenu() {
    const onClick = () => {
        console.log("hello");
    }
    return <>
        <Menu
            onClick={onClick}
            style={{ width: 300 }}
            defaultSelectedKeys={['1']}
            defaultOpenKeys={['sub1']}
            mode="inline"
            items={items}
        />
    </>
}

export default PluginMenu