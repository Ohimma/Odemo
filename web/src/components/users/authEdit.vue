<template>
  <el-dialog  :title="metaForm.title" v-model="metaForm.visible" destroy-on-close width="40%" >

    <el-form ref="formSubmit" :rules="rules" :model="formData" 
        :label-width="labelWidth">

        <template v-for="(item, index) in dataBase">
          <el-form-item :key="index"  v-if="item.level < 4"
              :label="item.label" :prop="item.name">
             

            <template v-if="item.name == 'method'">
              <el-select v-model="formData[item.name]">
                <el-option 
                  v-for="(item, index) in methods_options" :key="index"
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

            <template v-else-if="item.name == 'icon'">
              <el-input v-model="formData[item.name]" ></el-input>
              <a href="https://element-plus.gitee.io/#/zh-CN/component/icon" target="_blank">图标参考(目前只开放了一级菜单)</a>
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
        methods_options: [
          { id: 'GET',name: 'GET' }, 
          { id: 'PUT',name: 'PUT' }, 
          { id: 'POST',name: 'POST' }, 
          { id: 'DELETE',name: 'DELETE' }, 
          { id: 'OPTION',name: 'OPTION' }, 
        ],
        Ids : [],
        formData: {
          id: 0,
          pid: 0,
          name: '',
          method: 'GET',
          level: 1,
          url: '',
          icon: '',
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
        if (this.metaForm.title == "编辑") { 
            this.formData.id = this.metaForm.tableInfo.id
            this.formData.pid = this.metaForm.tableInfo.pid
            this.formData.name = this.metaForm.tableInfo.name
            this.formData.method = this.metaForm.tableInfo.method
            this.formData.icon = this.metaForm.tableInfo.icon
            this.formData.url = this.metaForm.tableInfo.url
            this.formData.detail = this.metaForm.tableInfo.detail
            this.formData.status = this.metaForm.tableInfo.status
        } else {
          this.formData.pid = this.metaForm.tableInfo.id
          this.formData.level = this.metaForm.tableInfo.level + 1
        }
      },
     
      changeSubmit() {
        this.$refs.formSubmit.validate((data) => {
          if (data) {
              // 对表单参数进行过滤和处理
              // 组装参数发送post请求
              console.log("xxxxxx", this.formData)
              if (this.metaForm.title == "编辑") {
                this.editHttp(this.formData)
              } else {
                this.addHttp(this.formData)
              }
          } 
        })
      },

      addHttp(params){
        console.log("entry user methods addHttp", params)
        this.$http.post('http://localhost:8080/api/user/auth', params)
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
        this.$http.put('http://localhost:8080/api/user/auth', params)
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