<template>
    <el-drawer :title="metaForm.title" v-model="metaForm.detailVisible" 
                direction="ltr" destroy-on-close show-close size="25%">

      <el-form >
          <el-form-item   v-for="(item, index) in dataBase" :key="index" :label="item.label + ': '">
              <span >{{ metaForm.tableInfo[item.name] }}</span>
          </el-form-item>

        <!-- <el-form-item :key="index" v-if="item.name == 'acls'" :label="item.label + ': '">
            <span v-for="(itemAcls,i) in metaForm.tableInfo.acls" :key="i">
              {{ itemAcls.title }} | 
            </span>
        </el-form-item> -->
      
      </el-form>
    </el-drawer>
</template>

<script>
export default {
    props: ["metaForm", "dataBase"],
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