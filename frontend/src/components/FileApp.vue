<template>
    <div style="padding:10px">
        <el-row style="margin-bottom: 10px;">
            <el-col :span="10">
                <el-input v-model="filterText" placeholder="请输入根目录"></el-input>
            </el-col>
        </el-row>
        <el-tree style="width:100%;" ref="treeRef" :props="props" :data="dataSource"
            :filter-node-method="filterNodeMethod" node-key="AbsolutePath" lazy :load="loadNode">
            <template #default="{ data }" style="width:100%;">
                <el-row style="width:100%;">
                    <el-space wrap :size="10">
                        <el-icon color="#409EFF">
                            <Document v-if="data.IsFile" />
                            <Folder v-else />
                        </el-icon>
                        <span class="ellipsis">{{ data.Name }}</span>
                    </el-space>
                    <el-space wrap :size="10" style="float:right">
                        <el-link v-if="!data.IsFile" type="primary" v-loading="data.opening" :icon="FolderOpened"
                            @click.stop="doOpen(data)">
                        </el-link>
                        <el-tooltip class="box-item" effect="dark" content="复制完整路径" placement="top">
                            <el-link type="primary" :icon="CopyDocument" @click.stop="doCopy(data)"></el-link>
                        </el-tooltip>
                        <el-popover trigger="click" :visible="data.deleteVisible" placement="top" :width="160">
                            <p>是否删除文件夹{{ data.Name }}</p>
                            <div style="text-align: right; margin: 0">
                                <el-button size="small" text @click.stop="data.deleteVisible = false">否
                                </el-button>
                                <el-button size="small" type="primary" @click.stop="doDelete(data)">是
                                </el-button>
                            </div>
                            <template #reference>
                                <el-link type="danger" :icon="Delete"></el-link>
                            </template>
                        </el-popover>
                    </el-space>
                </el-row>
            </template>
        </el-tree>
    </div>
</template>
<script lang="ts" setup>
import { ref, watch } from 'vue'
import { CopyDocument, Delete, Folder, FolderOpened } from '@element-plus/icons-vue'
import { GetByRootpath, OpenfileByPath } from '../../wailsjs/go/application/FileApp'
import { ElButton, ElCol, ElIcon, ElInput, ElLink, ElPopover, ElRow, ElSpace, ElTooltip, ElTree } from 'element-plus'
import { FilterValue, TreeNodeData, FilterNodeMethodFunction } from 'element-plus/es/components/tree/src/tree.type';
import * as TheNode from 'element-plus/es/components/tree/src/model/node';

interface FileTree extends TreeNodeData {
    Name: string
    AbsolutePath: string
    RelativePath: string
    Children?: FileTree[]
    IsFile: boolean
    DeleteVisible: boolean
    opening: boolean
}

const props = {
    label: 'Name',
    children: 'Children',
    isLeaf: 'IsFile',
}
const filterText = ref<string>('')

const treeRef = ref<InstanceType<typeof ElTree>>()

watch(filterText, (val) => {
    doFilter(val)
})

const doFilter = (val: string) => {
    treeRef.value!.filter(val)
}


const dataSource = ref<FileTree[]>()

const filterNodeMethod: FilterNodeMethodFunction = (value: FilterValue, data: TreeNodeData, child: TheNode.default) => {
    return data.AbsolutePath.includes(value)
}

const doOpen = (data: FileTree) => {
    console.log(data.opening)
    if (data.opening) {
        return
    }
    data.opening = true
    OpenfileByPath(data.AbsolutePath)
    setTimeout(() => {
        data.opening = false
    }, 1000)
}

const doCopy = (data: FileTree) => {
    const text = data.AbsolutePath
    navigator.clipboard.writeText(text).then(() => {
        console.log('复制成功')
    }, () => {
        console.log('复制失败')
    })
}

const doPrint = (data: FileTree) => {
    console.log("doPrint=>", data)
}
const doDelete = (data: FileTree) => {
    console.log("doDelete=>", data)
    data.deleteVisible = false
}
const getByRootpath = (rootpath: string) => {
    dataSource.value = []
    GetByRootpath(rootpath).then((result: FileTree[]) => {
        dataSource.value = result
    })
}

const loadNode = (node: any, resolve: any) => {
    if (node.data.Children) {
        resolve(node.data.Children)
        return
    }
    GetByRootpath(node.data.AbsolutePath).then((result: FileTree[]) => {
        if (!result) {
            result = []
        }
        resolve(result)
    })
}

</script>
<style >
.ellipsis {
    text-overflow: ellipsis;
    overflow: hidden;
    text-align: left;
}
</style>