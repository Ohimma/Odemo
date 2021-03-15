<template>
    <el-dialog class="roleDetail"  :title="metaForm.title" v-model="metaForm.detailVisible">

      <el-form >
        <template v-for="(item, index) in dataBase">
        <el-form-item  :key="index"  v-if="item.name !== 'acls'"  :label="item.title + ': '">
            <span >{{ metaForm.tableInfo[item.name] }}</span>
        </el-form-item>
        <el-form-item :key="index" v-if="item.name == 'acls'" :label="item.title + ': '">
            <span v-for="(itemAcls,i) in metaForm.tableInfo.acls" :key="i">
              {{ itemAcls.title }} | 
            </span>
        </el-form-item>
        </template>
        
      </el-form>

    </el-dialog>
</template>
<script>
export default {
    props: ["metaForm", "dataBase"],
    // emits:["closeDialog3"],

    // methods: {
    //     closeSubmit(type) {
    //         this.$emit('closeDialog3', type)
    //     },
    // },
}
</script>

<style scoped>

.el-form .el-form-item {
  margin-bottom: 0px;
}
/* .el-form .el-form-item .el-input {
  width: 70%;
} */

</style>