<template>
    <el-dialog v-model="settingDialogVisible" title="设置" width="60%">
        <el-form :model="settings" label-width="200px">
            <el-form-item label="是否显示隐藏文件">
                <el-switch v-model="settings.ShowHidden" active-color="#13ce66" inactive-color="#ff4949">
                </el-switch>
            </el-form-item>
            <el-form-item label="删除到回收站">
                <el-switch v-model="settings.DeleteToRecycleBin" active-color="#13ce66" inactive-color="#ff4949">
                </el-switch>
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="settingDialogVisible = false">Cancel</el-button>
                <el-button type="primary" @click="saveSettings">
                    Confirm
                </el-button>
            </span>
        </template>
    </el-dialog>
    <el-dialog v-model="newRootpathdialogVisible" title="新增一个路径" width="30%">
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
                <el-button @click="newRootpathdialogVisible = false">Cancel</el-button>
                <el-button type="primary" @click="doAdd">
                    Confirm
                </el-button>
            </span>
        </template>
    </el-dialog>
    <el-container style="padding:10px">
        <el-main>
            <el-tabs editable @edit="handleTabsEdit" v-loading="loading" type="card" v-model="activeName">
                <el-tab-pane v-for="item in fileTabs" :label="item.Label" :name="item.Name">
                    <FileApp :ref="setItemRef" :root-path="item.Rootpath"></FileApp>
                </el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { ElTabPane, ElDialog, ElTabs, TabPaneName } from 'element-plus'
import FileApp from './FileApp.vue'
import { OpenDirectoryDialog, GetByRootpaths } from '../../wailsjs/go/application/FileApp'
import { GetAllValues, PutAllValues } from '../../wailsjs/go/application/ConfigApp'
import { Settings } from '../types/settings/settings'
import { EventsOn } from '../../wailsjs/runtime/runtime'
interface Rootpath {
    Name: string
    Label: string
    Rootpath: string
}
const itemRefs = ref<any>([]);

const setItemRef = (el: any) => {
    if (el) {
        itemRefs.value.push(el);
    }
};

const settingDialogVisible = ref(false)

const settings = ref<Settings>({
    ShowHidden: true,
    DeleteToRecycleBin: true,
})

const initSettings = async () => {
    GetAllValues().then((result) => {
        settings.value!.ShowHidden = result.ShowHidden == 'true'
        settings.value!.DeleteToRecycleBin = result.DeleteToRecycleBin == 'true'
    })
}

initSettings()

const openSettingDialog = () => {
    GetAllValues().then((result) => {
        settings.value!.ShowHidden = result.ShowHidden == 'true'
        settings.value!.DeleteToRecycleBin = result.DeleteToRecycleBin == 'true'
        settingDialogVisible.value = true
    })
}

EventsOn(
    "openSettingDialog",
    () => {
        openSettingDialog()
    }
)

EventsOn(
    "openNewRootpathdialog",
    () => {
        openNewRootpathdialog()
    }
)



const saveSettings = () => {
    let values = {
        ShowHidden: settings.value.ShowHidden + "",
        DeleteToRecycleBin: settings.value.DeleteToRecycleBin + "",
    }
    PutAllValues(values).then(() => {
        settingDialogVisible.value = false
        itemRefs.value.forEach((item: any) => {
            item.refresh()
        })
    })
}

const newRootpath = ref<Rootpath>({
    Name: '',
    Label: '',
    Rootpath: ''
})

const newRootpathdialogVisible = ref(false)
const activeName = ref<string>()
const fileTabs = ref<Rootpath[]>([])
const loading = ref<boolean>(true)
const getRootPaths = () => {
    loading.value = true
    GetByRootpaths().then((data) => {
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
        newRootpath.value.Label = data
    })
}
const doAdd = () => {
    if (newRootpath.value.Rootpath == '') {
        return
    }
    fileTabs.value.push(newRootpath.value)
    activeName.value = newRootpath.value.Rootpath
    newRootpathdialogVisible.value = false
}

const openNewRootpathdialog = () => {
    newRootpath.value = {
        Name: '',
        Label: '',
        Rootpath: ''
    }
    newRootpathdialogVisible.value = true
}

const handleTabsEdit = (targetName: TabPaneName | undefined, action: 'remove' | 'add') => {
    if (action === 'add') {
        openNewRootpathdialog()
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