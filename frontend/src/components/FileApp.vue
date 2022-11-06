<template>
    <el-tree :data="dataSource" show-checkbox node-key="id" default-expand-all :expand-on-click-node="false">
        <template #default="{ node, data }">
            <span class="custom-tree-node">
                <span>{{ node.label }}</span>
                <span>
                    操作对象
                </span>
            </span>
        </template>
    </el-tree>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
interface FileInfo {
    Name: string
    AbsolutePath: string
    RelativePath: string
    FileInfo: {}
    Children?: FileInfo[]
    IsFile: boolean
}
const dataSource = ref<FileInfo[]>()
</script>