import api from "./api"

class plugin {
    addNewPlugin(projectPath, goPath, pluginName) {
        return api.send("service/add-new-plugin", { project_path: projectPath, go_path: goPath, plugin_name: pluginName })
    }

    getFunctionList(pluginName) {
        return api.send("service/get-function-list", pluginName)
    }

    getFunction(pluginName, functionName) {
        return api.send("service/get-function", {
            plugin_name: pluginName,
            function_name: functionName
        })
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

    onPluginAdded(callback) {
        return api.listen("route/plugin-added", callback)
    }
}

export default new plugin()