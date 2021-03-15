<template>
<el-dialog :title="info.title" v-model="info.visible" >

  <el-form :model="formData" :rules="rules" ref="formSubmit" :label-width="formLabelWidth">
    <el-form-item label="角色名称" prop="title" >
      <el-input v-model="formData.title" autocomplete="off"></el-input>
    </el-form-item>

    <el-form-item label="描述" >
      <el-input v-model="formData.description">
      </el-input>
    </el-form-item>

    <el-form-item label="资源" prop="acls" >
      <el-select v-model="formData.acls" multiple>
        <el-option 
          v-for="(item, index) in optionList" :key="index"
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
    props: ["info","configs"],
    emits: ["closeDialog2"],
    data() {
      return {
        formData: {
          title: '',
          acls: [],
          description: '',
        },
        rules: {
            title: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
            acls: [{ required: true, message: '请选择资源', trigger: 'blur' }]
        },
        formLabelWidth: '80px',
        optionList : [
          {"id": 1, "title": "aaa", "type": 2},
          {"id": 2, "title": "bbb", "type": 2},
          {"id": 3, "title": "ccc", "type": 2},
          {"id": 4, "title": "fff", "type": 2},
          {"id": 5, "title": "eee", "type": 2},
          {"id": 6, "title": "ggg", "type": 2},
          {"id": 7, "title": "hhh", "type": 2},
        ],
      };
    },
    create(){
      console.log("entry roleedit create")
    },
    mounted() {  // 当在DOM挂载组件时执行
      console.log("entry roleedit mount")
      this.formCover()
    },
    methods: {
      formCover(){
        console.log("entry roleedit methods formCover")
        this.formData.title = this.info.roleinfo.row.title
        this.formData.description = this.info.roleinfo.row.description
        const acls = this.info.roleinfo.row.acls
        for (var i=0, l=acls.length; i<l; i++ ) {
          console.log("xxxxx", acls[i].id)
          this.formData.acls.push(acls[i].id)
        }

        this.getAliasList()
      },
      getAliasList () {  // 获取全部 acl 列表
        console.log("entry roleedit methods getAliasList", this.formData)
        // this.$http.get('http://localhost:8080/api/console/acl',)
        //   .then((res) => {
        //     console.log("xxxxxxxxxxxx", res)
        //     res.data.forEach((item) => {
        //       item.label = item.title
        //       item.value = item.id
        //     })
        //     this.formConfigs.forEach((item) => {
        //       if (item.formItemProp.prop === 'acls') {
        //         item.optionList = res.data
        //       }
        //     })
        //   })
        //   .catch((err) => {
        //     console.log(err, 'err')
        //   })
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
              if (this.info.roleinfo.row.id) {
                this.editHttp(params)
              } else {
                this.addHttp(params)
              }
          } else {
              console.log("roleEdit vertify failed")
          }
        })
      },

      addHttp(params){
        console.log("entry roleEdit methods addHttp", params)
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
      closeSubmit(type) {
        this.$emit('closeDialog2', type)
      },
    }
}

</script>
