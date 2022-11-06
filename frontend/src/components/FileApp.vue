<template>
    <el-tree :props="props" :data="dataSource" node-key="Name" :load="loadNode" lazy>
        <template #default="{ data }">
            <span class="tree-node">
                <el-icon color="#409EFF" v-if="data.IsFile">
                    <Document />
                </el-icon>
                <el-icon color="#409EFF" v-else>
                    <Folder />
                </el-icon>
                <span>{{ data.Name }}</span>
                <span style="padding-right: 8px;">
                    <el-button type="text">1</el-button>
                </span>
            </span>
        </template>
    </el-tree>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
import { GetByRootpath } from '../../wailsjs/go/application/FileApp'
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

const dataSource = ref<FileInfo[]>()

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
        console.log("result=>", result)
        if (!result) {
            data.Children = []
        } else {
            data.Children = result
        }
        resolve(data.Children)
    })
}
getByRootpath("")
</script>
<style>
.tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    padding-right: 8px;
    width: 100%;
}
</style>