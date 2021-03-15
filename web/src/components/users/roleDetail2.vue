<template>
    <el-drawer :title="metaForm.title" v-model="metaForm.detailVisible" 
                direction="ltr" destroy-on-close show-close size="25%">

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

    </el-drawer>
</template>
<script>
export default {
    prop: ["metaForm", "dataBase"],
}
</script>

<style scoped>

.el-form {
  padding-left: 25px;
}
.el-form .el-form-item {
  margin-bottom: 0px;
}


</style>