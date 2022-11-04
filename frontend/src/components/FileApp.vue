
<script>
import { defineComponent } from "vue";
import { GetAllDir } from '../../wailsjs/go/application/FileApp'

const data = []
export default defineComponent({
    setup() {
        return {
            getAllDir: () => {
                GetAllDir().then(result => {
                    console.log(result[0])
                    data = result
                })
            }
        };
    }
});

</script>
<template>
    <n-space vertical>
        <n-button @click="getAllDir">
            刷新数据
        </n-button>
        <n-tree block-line :data="data" key-field="Name" label-field="AbsolutePath" virtual-scroll
            style="height: 320px" />
    </n-space>
</template>

<style scoped>
.result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
}

.input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
}

.input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
}

.input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
}
</style>
