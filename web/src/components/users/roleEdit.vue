<template>
  <el-dialog  :title="metaForm.title" v-model="metaForm.visible" destroy-on-close width="40%">

    <el-form class="custom_form" ref="formSubmit" :rules="rules" :model="formData" 
        :label-width="labelWidth">

        <template v-for="(item, index) in dataBase">
          <el-form-item :key="index"  
              v-if="item.level < 4"
              :label="item.label" :prop="item.name">
             
            <template v-if="item.name == 'password'">
              <el-input v-model="formData[item.name]" type="password"></el-input>
            </template>
            
            <template v-else-if="item.name == 'auth_ids'">
              <!-- 
                :default-expanded-keys="[0,1,2]"  默认展开的key
                @check-change="handleCheckChange"  点击check时触发的事件，返回 handleCheckChange(data, checked, indeterminate)
                @check="handleCheck"  返回 handleCheck(data, checked, indeterminate)
                check-on-click-node  点击即选中
               -->
              <el-tree
                :data="authListTree"
                ref="treeBox"
                node-key="id"
                show-checkbox
                highlight-current
                :expand-on-click-node="false"

                default-expand-all
                :default-checked-keys="formData.auth_ids"
                :check-on-click-node="true"
                >
                <template #default="{ data }">
                  <span class="custom-tree-node">
                    <span>{{ data.name}}({{ data.id}}) </span>
                  </span>
                </template>
              </el-tree>
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

import {mapState} from 'vuex'
export default {
    props: ["metaForm","dataBase","rules"],
    emits: ["closeDialog"],
    data() {
      return {
        labelWidth: '110px',
        formData: {
          id: 0,
          name: '',
          auth_ids: [],
          detail: '',
          status: true,
        },
        
      };
    },
    computed: {
        ...mapState({
            authListTree: state => state.layout.authListTree,
        })
    },
    mounted() {  // 当在DOM挂载组件时执行
      this.tableCoverForm()
      
    },

    methods: {
      
      tableCoverForm(){

        if (this.metaForm.tableInfo.id) { 
            this.formData.id = this.metaForm.tableInfo.id
            this.formData.name = this.metaForm.tableInfo.name
            this.formData.detail = this.metaForm.tableInfo.detail
            this.formData.status = this.metaForm.tableInfo.status
            // 字符串 转 数组
            const ids = this.metaForm.tableInfo.auth_ids.split(",")
            for (var i=0; i<ids.length; i++ ) {
              if (ids[i]) {
                this.formData.auth_ids.push(parseInt(ids[i]))
              }
            }
          // 数据更新时，即下次 dom 加载时执行
          this.$nextTick(() => {
            this.$refs.treeBox.setCheckedKeys(this.formData.auth_ids)
          })  
        }
      },
     
      changeSubmit() {
        this.$refs.formSubmit.validate((data) => {
          if (data) {
              this.formData.auth_ids = this.$refs.treeBox.getCheckedKeys(true)

              if (this.metaForm.tableInfo.id) {
                this.editHttp(this.formData)
              } else {
                this.addHttp(this.formData)
              }
          } 
        })
      },

      addHttp(params){
        this.$http.post('http://localhost:8080/api/user/role', params)
            .then(() => {
                this.$message({
                  type: "sucess",
                  message: "添加成功",
                })
                this.closeSubmit('save')
            })
            .catch((err) => {
                this.$message({
                  type: "error",
                  message: err,
                })
                this.closeSubmit()
            })
      },
      editHttp(params){
        this.$http.put('http://localhost:8080/api/user/role', params)
            .then(() => {
                this.$message({
                  type: "sucess",
                  message: "编辑程功",
                })
                // 更新完角色权限，需要重新获取当前用户的权限列表，写到 store 中
                
                this.closeSubmit('save')
            })
            .catch((err) => {
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

.custom_form .el-form {
  width: 85%;
}
.custom_form .el-form .el-select {
  width: 100%;
}
.custom_form .el-input {
  width: 80%;
}
</style>

<style >

.custom_form .el-tree {
  width: 80%;
  border: 1px solid #DCDFE6;
}

</style>
