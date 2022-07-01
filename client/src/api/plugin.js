import api from "./api"

class plugin {
    addNewPlugin(projectPath, goPath, pluginName) {
        return api.send("service/add-new-plugin", { project_path: projectPath, go_path: goPath, plugin_name: pluginName })
    }

    getFunctionList(pluginName) {
        return api.send("service/get-function-list", pluginName)
    }

    getVariableList(pluginName) {
        return api.send("service/get-variable-list", pluginName)
    }

    getPluginNameList() {
        return api.send("service/get-plugin-name-list", null)
    }

    getAllGoBinPath() {
        return api.send("service/get-all-go-bin-path", null)
    }
}

export default new plugin()