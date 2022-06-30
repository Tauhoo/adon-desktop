import api from "./api"

class plugin {
    addNewPlugin(projectPath, goPath, pluginName) {
        return api.send("service/add-new-plugin", { projectPath, goPath, pluginName })
    }

    getFunctionList(pluginName) {
        return api.send("service/get-function-list", pluginName)
    }

    getVariableList(pluginName) {
        return api.send("service/get-variable-list", pluginName)
    }

    GetPluginNameList() {
        return api.send("service/get-plugin-name-list", null)
    }
}

export default new plugin()