<template>
    <div>
        <el-dialog v-model="dialogVisible" title="新增一个路径" width="30%">
            <el-form :model="newRootpath" label-width="80px">
                <el-form-item label="路径">
                    <el-input @click="choosePath" v-model="newRootpath.Rootpath" placeholder="请输入路径"></el-input>
                </el-form-item>
                <el-form-item label="名称">
                    <el-input v-model="newRootpath.Label" placeholder="请输入名称"></el-input>
                </el-form-item>
            </el-form>

            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogVisible = false">Cancel</el-button>
                    <el-button type="primary" @click="doAdd">
                        Confirm
                    </el-button>
                </span>
            </template>
        </el-dialog>
        <el-tabs editable @edit="handleTabsEdit" v-loading="loading" type="card" v-model="activeName">
            <el-tab-pane v-for="item in fileTabs" :label="item.Label" :name="item.Name">
                <FileApp :root-path="item.Rootpath"></FileApp>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { ElTabPane, ElDialog, ElTabs, TabPaneName } from 'element-plus'
import FileApp from './FileApp.vue'
import { OpenDirectoryDialog, GetByRootpaths } from '../../wailsjs/go/application/FileApp'
interface Rootpath {
    Name: string
    Label: string
    Rootpath: string
}

const newRootpath = ref<Rootpath>({
    Name: '',
    Label: '',
    Rootpath: ''
})
const dialogVisible = ref(false)
const activeName = ref<string>()
const fileTabs = ref<Rootpath[]>([])
const loading = ref<boolean>(true)
const getRootPaths = () => {
    loading.value = true
    GetByRootpaths().then((data) => {
        console.log(data)
        fileTabs.value = data
        activeName.value = data[0].Name
    }).finally(() => {
        loading.value = false
    })
}

const choosePath = () => {
    OpenDirectoryDialog().then((data) => {
        newRootpath.value.Rootpath = data
        newRootpath.value.Name = data
    })
}
const doAdd = () => {
    console.log(newRootpath.value)
    if (newRootpath.value.Rootpath == '') {
        return
    }
    fileTabs.value.push(newRootpath.value)
    dialogVisible.value = false
}

const handleTabsEdit = (targetName: TabPaneName | undefined, action: 'remove' | 'add') => {
    if (action === 'add') {
        newRootpath.value = {
            Name: '',
            Label: '',
            Rootpath: ''
        }
        dialogVisible.value = true
    } else if (action === 'remove') {
        fileTabs.value = fileTabs.value.filter((tab) => tab.Name !== targetName)
        if (activeName.value === targetName) {
            if (fileTabs.value.length > 0) {
                activeName.value = fileTabs.value[0].Name
            } else {
                activeName.value = ''
            }
        }
    }
}
getRootPaths()
</script>