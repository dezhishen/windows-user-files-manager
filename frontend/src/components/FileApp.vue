<script>
import { GetByRootpath } from '../../wailsjs/go/application/FileApp'
export default {
    name: "FileApp",
    data() {
        return {
            data: [],
            defaultProps: {
                children: 'Children',
                label: 'Name',
                isLeaf: "IsFile"
            }
        }
    },
    methods: {
        getAllDir() {
            GetByRootpath("").then(res => {
                this.data = res
            })
        },
        loadNode: (node, resolve) => {
            GetByRootpath(node.data.AbsolutePath).then(res => {
                if (!res) {
                    res = []
                }
                resolve(res)
            }).catch(err => {
                console.log("err", err)
            })
        },
        handleNodeClick: (node) => {
            console.log(node)
        }
    }
}
</script>
<template>
    <div>
        <el-button type="primary" @click="getAllDir">Click me</el-button>
        <el-tree :data="data" lazy :load="loadNode" :props="defaultProps" />
    </div>
</template>