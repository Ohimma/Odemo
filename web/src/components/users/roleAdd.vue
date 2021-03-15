<template>
<el-dialog :title="info.title" v-model="info.visible" >

  <el-form :model="form" :rules="rules" ref="baseForm" :label-width="formLabelWidth">
    <el-form-item label="角色名称" prop="title" >
      <el-input v-model="form.title" autocomplete="off" >

      </el-input>
    </el-form-item>
    <el-form-item label="其他" >
      <el-input v-model="form.description">
      </el-input>
    </el-form-item>
    <el-form-item label="资源" prop="acls" >
      <el-select v-model="form.acls" multiple>
        <el-option label="区域一" value="shanghai"></el-option>
        <el-option label="区域二" value="beijing"></el-option>
        <el-option label="区域三" value="aaa"></el-option>
        <el-option label="区域四" value="bbb"></el-option>
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
    props: ["info","configs"],
    data() {
      return {
        form: {
          title: '',
          acls: '',
          description: '',
        },
        rules: {
            title: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
            acls: [{ required: true, message: '请选择资源', trigger: 'blur' }]
        },
        formLabelWidth: '100px'
      };
    },
    methods: {
        changeSubmit() {
          this.$refs.baseForm.validate((data) => {
            if (data) {
                // 对表单参数进行过滤和处理

                // 组装参数发送post请求
                let params = {
                    title: this.form.title,
                    acls: typeof this.form.acls === 'object' ? this.form.acls : this.form.acls.join(','),
                    description: this.form.description
                }
                console.log("xxxxxxx", params)
                // this.$http.post('http://localhost:8080/api/console/role', params)
                //     .then(() => {
                //         this.closeSubmit('save')
                //     })
                //     .catch((err) => {
                //         console.log(err, 'err')
                //     })
            } else {
                console.log("roleadd vertify failed")
            }
          })
        },

        closeSubmit(type) {
          this.$emit('closeDialog1', type)
        },
    }
}

</script>
