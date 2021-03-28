<template>
  <el-dialog  :title="metaForm.title" v-model="metaForm.visible" destroy-on-close width="40%">

    <el-form ref="formSubmit" :rules="rules" :model="formData" 
        :label-width="labelWidth">

        <template v-for="(item, index) in dataBase">
          <el-form-item :key="index"  v-if="item.level < 4"
              :label="item.label" :prop="item.name">
             
            <template v-if="item.name == 'password'">
              <el-input v-model="formData[item.name]" type="password"></el-input>
            </template>
            
            <template v-else-if="item.name == 'role_ids'">
              <el-select v-model="formData[item.name]" multiple>
                <el-option 
                  v-for="(item, index) in roleIds" :key="index"
                  :label="item.name" :value="item.id"  
                >
                </el-option>
              </el-select>
            </template>

            <template v-else-if="item.name == 'status'">
              <el-switch v-model="formData[item.name]"  
                  active-color="#13ce66" inactive-color="#ff4949">
              </el-switch>
            </template>

            <template v-else>
              <el-input v-model="formData[item.name]" ></el-input>
            </template>
          </el-form-item> 
        </template>
        
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
        labelWidth: '110px',
        roleIds : [],
        formData: {
          id: 0,
          name: '',
          password: '',
          phone: '',
          role_ids: [],
          detail: '',
          status: true,
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
            this.formData.id = this.metaForm.tableInfo.id
            this.formData.name = this.metaForm.tableInfo.name
            this.formData.password = this.metaForm.tableInfo.password
            this.formData.phone = this.metaForm.tableInfo.phone
            this.formData.detail = this.metaForm.tableInfo.detail
            this.formData.status = this.metaForm.tableInfo.status
            // 字符串 转 数组
            const ids = this.metaForm.tableInfo.role_ids.split(",")
            for (var i=0; i<ids.length; i++ ) {
              if (ids[i]) {
                this.formData.role_ids.push(parseInt(ids[i]))
              }
            }
        }
        this.getRoleList()
      },
       getRoleList () {  // 获取全部 acl 列表
        console.log("entry roleForm methods getRoleList")
        this.$http.get('http://localhost:8080/api/user/role',)
          .then((res) => {
              this.roleIds = res.data.result.list
          })
          .catch((err) => {
            console.log(err, 'err')
                this.$message({
                  type: "error",
                  message: err,
                })
            this.closeSubmit()
          })
      },

      changeSubmit() {
        this.$refs.formSubmit.validate((data) => {
          if (data) {
              // 对表单参数进行过滤和处理
              // 组装参数发送post请求
              if (this.metaForm.tableInfo.id) {
                this.editHttp(this.formData)
              } else {
                this.addHttp(this.formData)
              }
          } 
        })
      },

      addHttp(params){
        console.log("entry user methods addHttp", params)
        this.$http.post('http://localhost:8080/api/user/user', params)
            .then((res) => {
                this.$message({
                  type: "sucess",
                  message: "添加成功",
                })
                this.closeSubmit('save')
            })
            .catch((err) => {
                console.log(err, 'err')
                this.$message({
                  type: "error",
                  message: err,
                })
                this.closeSubmit()
            })
      },
      editHttp(params){
        console.log("entry user methods editHttp", params)
        this.$http.put('http://localhost:8080/api/user/user', params)
            .then((res) => {
                this.$message({
                  type: "sucess",
                  message: "编辑程功",
                })
                this.closeSubmit('save')
            })
            .catch((err) => {
                console.log(err, 'err')
                this.$message({
                  type: "error",
                  message: err,
                })
                this.closeSubmit()
            })
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