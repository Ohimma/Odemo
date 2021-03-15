<template>
  <el-dialog  :title="metaForm.title" v-model="metaForm.visible" destroy-on-close width="35%">

    <el-form ref="formSubmit" :rules="rules" :model="formData" 
        :label-width="labelWidth">

        <el-form-item label="角色名称" prop="title" >
          <el-input v-model="formData.title"></el-input>
        </el-form-item>

        <el-form-item label="描述" >
          <el-input v-model="formData.description">
          </el-input>
        </el-form-item>

        <el-form-item label="资源" prop="acls" >
          <el-select v-model="formData.acls" multiple>
            <el-option 
              v-for="(item, index) in aclsList" :key="index"
              :label="item.title" :value="item.id"  
            >
            </el-option>
          </el-select>
        </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="closeSubmit">取 消</el-button>
        <el-button type="primary" @click="changeSubmit">确 定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>

export default {
    props: ["metaForm","dataBase","rules"],
    emits: ["closeDialog"],
    data() {
      return {
        labelWidth: '100px',
        aclsList : [],
        formData: {
          title: '',
          acls: [],
          description: '',
        },
      };
    },

    mounted() {  // 当在DOM挂载组件时执行
      this.tableCoverForm()
    },

    methods: {
      tableCoverForm(){
        console.log("entry roleForm mount tableCoverForm = ", this.metaForm.title)
        if (this.metaForm.tableInfo.id) { 
            this.formData.title = this.metaForm.tableInfo.title
            this.formData.description = this.metaForm.tableInfo.description
            const acls = this.metaForm.tableInfo.acls
            for (var i=0, l=acls.length; i<l; i++ ) {
              this.formData.acls.push(acls[i].id)
            }
        }
        this.getAclsList()
      },
      getAclsList () {  // 获取全部 acl 列表
        console.log("entry roleForm methods getAliasList")
        this.$http.get('http://localhost:8080/api/console/acl',)
          .then((res) => {
              this.aclsList = res.data
          })
          .catch((err) => {
            console.log(err, 'err')
          })
      },


      changeSubmit() {
        this.$refs.formSubmit.validate((data) => {
          if (data) {
              // 对表单参数进行过滤和处理

              // 组装参数发送post请求
              let params = {
                  title: this.formData.title,
                  acls: typeof this.formData.acls === 'object' ? this.formData.acls : this.formData.acls.join(','),
                  description: this.formData.description
              }
              if (this.metaForm.tableInfo.id) {
                this.editHttp(params)
              } else {
                this.addHttp(params)
              }
          } else {
              this.$message({
                type: 'error',
                message: "请输入表单数据"
              })
          }
        })
      },

      addHttp(params){
        console.log("entry roleEdit methods addHttp", params)
        this.$message({
          type: "sucess",
          message: "添加成功",
        })
        this.closeSubmit('save')
        // this.$http.post('http://localhost:8080/api/console/role', params)
        //     .then((res) => {
        //         this.closeSubmit('save')
        //     })
        //     .catch((err) => {
        //         console.log(err, 'err')
        //     })
      },
      editHttp(params){
        console.log("entry roleedit methods editHttp", params)
        this.$message({
          type: "sucess",
          message: "编辑程功",
        })
        this.closeSubmit('save')
        // this.$http.post('http://localhost:8080/api/console/role' + this.info.roleinfo.row.id, params)
        //     .then((res) => {
        //         this.closeSubmit('save')
        //     })
        //     .catch((err) => {
        //         console.log(err, 'err')
        //     })
      },
      closeSubmit(action) {
        this.$emit('closeDialog', action)
      },
    }
}

</script>

<style scoped>

.el-form {
  width: 85%;
}
.el-form .el-select {
  width: 100%;
}

</style>