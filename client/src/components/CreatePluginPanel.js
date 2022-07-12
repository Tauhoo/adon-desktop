import { Modal } from "antd"
import { Button, Cascader, Input, Alert } from "antd"
import { useEffect, useState } from "react"
import styled from "styled-components"
import { LoadingOutlined } from '@ant-design/icons'

import plugin from "../api/plugin"
const Container = styled.div`
display: grid;
grid-template-rows: max-content max-content;
gap: 20px;
`

const FolderSelectorContainer = styled.div`
display: grid;
grid-template-columns:  1fr max-content;
width: 100%;
gap: 10px;
`

const FilePathDisplay = styled.div`
height: 100%;
width: 100%;
padding: 0px 20px;
border-radius:3px ;
background-color: #F6F6F6;
display: flex;
align-items: center;
`

function CreatePluginPanel({ onBuildSucess }) {
    const [isBuilding, setIsBuilding] = useState(false)
    const [goBinPath, setGoBinPath] = useState(null)
    const [projectPath, setProjectPath] = useState(null)
    const [goBinPathChoices, setGoBinPathChoices] = useState([])
    const [warning, setWarning] = useState(null)
    const [name, setName] = useState("")

    const initGoBinPathChoices = async () => {
        try {
            const result = await plugin.getAllGoBinPath()
            const choices = result.data.map(value => ({ label: value, value }))
            setGoBinPathChoices(choices)
            if (choices.length > 0) setGoBinPath(choices[0].value)
        } catch (error) {
            console.log(error);
        }
    }

    useEffect(() => {
        initGoBinPathChoices()
    }, [])

    const onChangeFile = (event) => {
        const files = event.target.files;
        const filterFiles = [...files].filter(value => value.name === "go.mod")
        if (filterFiles.length != 1) {
            setProjectPath(null)
            return
        }

        const [goModFile] = filterFiles
        const filePathLength = goModFile.name.length + 1
        const newProjectPath = goModFile.path.slice(0, goModFile.path.length - filePathLength)
        setProjectPath(newProjectPath)
    }

    const onBuild = async () => {
        if (name === "") {
            return setWarning("name is not defined")
        }

        if (!/^[a-zA-Z_]*$/.test(name)) {
            return setWarning("name must consist of a-z, A-Z and _ only")
        }

        if (goBinPath === null) {
            return setWarning("go binary is not selected")
        }

        if (projectPath === null) {
            return setWarning("project is not selected")
        }

        setWarning(null)
        setIsBuilding(true)

        try {
            await plugin.addNewPlugin(projectPath, goBinPath, name)
            setName("")
            if (onBuildSucess) onBuildSucess()
        } catch (error) {
            setWarning(`build plugin fail code: ${error.code}, info: ${error.data}`)
        }

        setIsBuilding(false)
    }

    return <Container>
        <Input value={name} placeholder="Name" onChange={e => setName(e.target.value)} disabled={isBuilding}></Input>
        <Cascader
            placeholder="Select Go binary"
            style={{ width: '100%' }}
            options={goBinPathChoices}
            value={goBinPath}
            onChange={(value) => setGoBinPath(value ? value[0] : null)}
        />
        <FolderSelectorContainer>
            <FilePathDisplay>{projectPath === null ? "No project is selected" : projectPath}</FilePathDisplay>
            <Button disabled={isBuilding}>
                <label htmlFor="file-upload">
                    Select
                </label>
            </Button>
        </FolderSelectorContainer>
        <input style={{ display: "none" }} id="file-upload" type="file" webkitdirectory="true" directory="true" onChange={onChangeFile} />
        {!isBuilding && <div><Button onClick={onBuild}>Build project</Button></div>}
        {isBuilding && <div>Plugin is building <LoadingOutlined /></div>}
        {warning !== null && <Alert type="error" description={warning}></Alert>}
    </Container >
}

export default CreatePluginPanel

export function CreatePluginPanelModal() {
    const [visible, setVisible] = useState(false)
    return <>
        <div style={{ height: "40px", display: "flex", justifyContent: "center", alignItems: "center", width: "100%" }}>
            <Button style={{ width: "100%" }} type="primary" onClick={() => setVisible(true)}>Create plugin</Button>
        </div>
        <Modal
            visible={visible}
            onCancel={() => setVisible(false)}
            footer={null}
            title="Create plugin"
            width="800px"
        >
            <CreatePluginPanel onBuildSucess={() => setVisible(false)}></CreatePluginPanel>
        </Modal>
    </>
} 