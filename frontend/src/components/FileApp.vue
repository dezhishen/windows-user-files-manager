<template>
    <div style="padding:10px">
        <el-row style="margin-bottom: 10px;">
            <el-col :span="20">
                <el-input v-model="filterText" placeholder="请输入根目录"></el-input>
            </el-col>
            <el-col :span="4">
                <el-button style="float:left" :icon="Search" circle @click="doFilter" />
            </el-col>
        </el-row>
        <el-tree style="width:100%;" ref="treeRef" :props="props" :data="dataSource" :filter-node-method="filterNode"
            node-key="AbsolutePath" :load="loadNode" lazy>
            <template #default="{ data }" style="width:100%;">
                <el-row style="width:100%;">
                    <el-space wrap :size="10">
                        <el-icon color="#409EFF">
                            <Document v-if="data.IsFile" />
                            <Folder v-else />
                        </el-icon>
                        <el-tooltip class="box-item" effect="dark" :content="data.AbsolutePath" placement="top">
                            <span class="ellipsis">{{ data.Name }}</span>
                        </el-tooltip>
                    </el-space>
                    <el-space wrap :size="10" style="float:right">
                        <el-link type="primary" :icon="CopyDocument" @click.stop="doCopy(data)"></el-link>
                        <el-link type="danger" :icon="Delete" @click.stop="doDelete(data)"></el-link>
                    </el-space>
                </el-row>
            </template>
        </el-tree>
    </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
import { Search, CopyDocument, Delete } from '@element-plus/icons-vue'
import { GetByRootpath } from '../../wailsjs/go/application/FileApp'

import { ElTree } from 'element-plus'
interface FileInfo {
    Name: string
    AbsolutePath: string
    RelativePath: string
    FileInfo: {}
    Children?: FileInfo[]
    IsFile: boolean
}

const props = {
    label: 'Name',
    children: 'Children',
    isLeaf: 'IsFile',
}
const filterText = ref<string>('')

const treeRef = ref<InstanceType<typeof ElTree>>()

const doFilter = () => {
    treeRef.value!.filter(filterText.value)
}

const dataSource = ref<FileInfo[]>()

const filterNode = (value: string, data: FileInfo) => {
    if (!value) return true
    return data.Name === value
}

const doCopy = (data: FileInfo) => {
    const text = data.AbsolutePath;
    // 复制到剪贴板
    navigator.clipboard.writeText(text).then(() => {
        console.log('复制成功')
    }, () => {
        console.log('复制失败')
    })
}

const doPrint = (data: FileInfo) => {
    console.log("doPrint=>", data)
}
const doDelete = (data: FileInfo) => {
    console.log("doDelete=>", data)
}
const getByRootpath = (rootpath: string) => {
    GetByRootpath(rootpath).then((result: FileInfo[]) => {
        dataSource.value = result
    })
}

const loadNode = (node: any, resolve: any) => {
    const data = node.data
    if (data.Children) {
        return resolve(data.Children)
    }
    GetByRootpath(data.AbsolutePath).then((result: FileInfo[]) => {
        if (!result) {
            result = []
        }
        resolve(result)
    })
}
getByRootpath("")
</script>
<style >
.ellipsis {
    text-overflow: ellipsis;
    overflow: hidden;
    text-align: left;
}
</style>