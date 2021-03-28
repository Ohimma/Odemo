<template>
  <div class="user" v-loading="loading">

    <div class="header">

      <!-- <el-form size="mini" inline :model="query" ref="query2">
          <el-form-item prop="title" :rules="[{ required: true, message: '字段不能为空'},]">
            <el-input v-model="query.title" placeholder="角色名称"></el-input>
          </el-form-item>

          <el-form-item prop="status">
            <el-select v-model="query.status" placeholder="角色状态">
              <el-option label="使用中" value="1"></el-option>
              <el-option label="禁用中" value="2"></el-option>
              <el-option label="已删除" value="3"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="submitForm('query2')">查询</el-button>
            <el-button @click="resetForm('query2')">重置</el-button>
          </el-form-item>
          </el-form> -->
    </div>

    <div class="nav">
       <div class="add">

           <template v-for="(item, index) in tableData.auths" :key="index">

               <template v-if="user_auths.indexOf(item.id) > 0 || user_id == 1 "> 
                    <el-button v-if="item.method == 'POST'"  type="primary" size="mini" @click="handle2('add')">
                        <span class="el-icon-circle-plus"> 添加({{item.id}})  </span>
                    </el-button>
               </template>
               <template v-else> 
                    <el-button v-if="item.method == 'POST'"  type="primary" size="mini" @click="handle2('add')">
                        <span class="el-icon-circle-plus"> 添加({{item.id}})  </span>
                    </el-button>
               </template>
               
            </template>
          <!-- ref 和 点击事件名字关联，prop名字不匹配无法重置 -->
        </div>
    </div>

    <div class="content" >
        <el-table style="width: 100%" :stripe="true" :border="true" :highlight-current-row="true"
            :default-expand-all="true" size="small" 
            :data="tableData.data.filter(data => !search || data.title.toLowerCase().includes(search.toLowerCase()))" >

            <template v-for="(item,index) in dataBase">
                <el-table-column :key="index" v-if="item.level < 5 && item.name != 'password'" :label="item.label"
                :sortable="item.sortable" :prop="item.name"  
                :min-width="item.width" show-overflow-tooltip>
                    <template #default="scope">
                        <span>{{ scope.row[item.name] }}</span>
                    </template>
                </el-table-column> 
            </template>

            <el-table-column align="right" :fit="true" min-width="150px" >
                <template #header>
                    <el-input size="mini" v-model="search" placeholder="在本页搜索角色名称"/>
                </template>

                <template #default="scope" >
                    <el-button size="mini" @click="handle2('detail', scope.row)"> Detail </el-button>
                    <template v-for="(item, index) in tableData.auths" :key="index">
                        <!-- <template  v-if="item.method == 'GET' && user_auths.indexOf(item.id) > 0">
                            <el-button size="mini" @click="handle2('detail', scope.row)"> Detail</el-button>
                            <el-button size="mini" @click="handle2('edit', scope.row)"> Edit</el-button>
                            <el-button size="mini" @click="handle2('delete', scope.row)" type="danger"> Delete</el-button>
                        </template>
                        <template  v-if="item.method == 'GET' && user_auths.indexOf(item.id) < 0 && user_id == 1">
                            <el-button size="mini" @click="handle2('detail', scope.row)"> Detail</el-button>
                            <el-button size="mini" @click="handle2('edit', scope.row)"> Edit</el-button>
                            <el-button size="mini" @click="handle2('delete', scope.row)" type="danger"> Delete</el-button>
                        </template>
                         <template  v-if="item.method == 'GET' && user_auths.indexOf(item.id) < 0 && user_id != 1">
                            <el-button disabled size="mini" @click="handle2('detail', scope.row)"> Detail</el-button>
                            <el-button disabled size="mini" @click="handle2('edit', scope.row)"> Edit</el-button>
                            <el-button disabled size="mini" @click="handle2('delete', scope.row)" type="danger"> Delete</el-button>
                        </template> -->
                        
                        <template  v-if="user_auths.indexOf(item.id) > 0 || user_id == 1 ">
                            <el-button v-if="item.method == 'PUT'" size="mini" @click="handle2('edit', scope.row)"> Edit </el-button>
                            <el-button v-else-if="item.method == 'DELETE'" size="mini" @click="handle2('delete', scope.row)" type="danger"> Delete </el-button>
                        </template>
                       
                         <template  v-else>
                            <el-button v-if="item.method == 'PUT'"  disabled size="mini" @click="handle2('edit', scope.row)"> Edit </el-button>
                            <el-button v-else-if="item.method == 'DELETE'"  disabled size="mini" @click="handle2('delete', scope.row)" type="danger"> Delete </el-button>
                        </template>
                        
                    </template>
                </template>
            </el-table-column>
        </el-table>
    </div>

    <detail 
        v-if="metaForm.detailVisible"
        :metaForm="metaForm"
        :dataBase="dataBase"
    />
    <edit 
        v-if="metaForm.visible"
        :metaForm="metaForm"
        :rules="rules"
        :dataBase="dataBase"
        @closeDialog="closeDialog"
    />

  </div>
</template>


<script>

import detail from "@/components/users/Detail.vue"
import edit from "@/components/users/userEdit.vue"
import {mapState} from 'vuex'

export default {
    components: {
        detail,edit
    },
    data() {
        return {
            search: '',
            loading: false,
            queryParams: {
                name: "",
                phone: "",
                page: 1,
                pageSize: 10,
                order_by: 'id',
                order_type: 'desc',
                currentPage: 1,
            },
            tableData: {
                total: 0,
                data: [],
                auths: [],
            },
            rules: {
                name: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
                password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
                phone: [{ required: true, message: '请输入手机号', trigger: 'blur' }],
            },
            dataBase: [
                { name: "id", label: "Id", "sortable": true, level: 4, width: "30px" },
                { name: "name", label: "用户名", sortable: true, level: 1, width: "70px" },
                { name: "phone", label: "手机", sortable: false, level: 1, width: "100px" },
                { name: "password", label: "密码", sortable: false, level: 3, width: "50px" },
                { name: "role_ids", label: "角色列表", sortable: false, level: 3, width: "50px" },
                { name: "create_id", label: "创建者id", sortable: false, level: 6, width: "30px" },
                { name: "update_id", label: "更新者id", sortable: false, level: 6, width: "30px" },
                { name: "create_at", label: "创建时间", sortable: false, level: 6, width: "80px" },
                { name: "update_at", label: "更新时间", sortable: false, level: 4, width: "80px" },
                { name: "login_ip", label: "最后登录ip", sortable: false, level: 6, width: "80px" },
                { name: "login_time", label: "最后登录时间", sortable: false, level: 6, width: "80px" },
                { name: "detail", label: "备注", sortable: false, level: 3, width: "80px" },
                { name: "status", label: "状态", sortable: false, level: 2, width: "50px" },
            ],
             metaForm: {
                visible: false,
                detailVisible: false,
                title: "",
                tableInfo: {},
            },
        }
    },
    mounted() {
        console.log("entry user mount")
        this.getList()
        this.getAuthSelf()
    },
    computed: {
        ...mapState({
            user_id: state => state.layout.user_id,
            user_auths: state => state.layout.user_auths,
        }),
    },
    methods: {
        handle2(type, row){
            console.log("entry role handle2 = ", type);
            if (type == "add") {
                this.metaForm.tableInfo = ""
                this.metaForm.visible = true
                this.metaForm.title = "添加"
            } else if (type =="edit"){
                this.metaForm.visible = true
                this.metaForm.title = "编辑"
                this.metaForm.tableInfo = row
            } else if (type =="detail") {
                this.metaForm.detailVisible = true
                this.metaForm.title = "详细信息"
                this.metaForm.tableInfo = row
            } else if (type =="delete"){
                this.$confirm('确定删除该角色吗？ (' + row.id + ') ' + row.name, '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning',
                })
                .then(() => {
                    this.deleteId(row)
                })
                .catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除'
                    })
                })

            } else {
                console.log('err')
            }
        },

        getAuthSelf(){
            // this.$nextTick(() => {
                const auth_list = this.$store.state.layout.authListInit
                for (var i in auth_list) {
                    if (this.$route.path == auth_list[i].url) {
                        this.tableData.auths.push(auth_list[i])
                    } 
                }
            // }) 
        },

        getList() {
            this.loading = true
            console.log("entry user methods getList", this.loading)
            let params = {
                name: this.queryParams.name,
                // status: Number(this.query.status),
                // status: this.query.status,
                offset: (this.queryParams.page -1) * this.queryParams.pageSize,
                limit: this.queryParams.pageSize,
                order_by: this.queryParams.order_by,
                order_type: this.queryParams.order_type,
            }
            this.$http.get('http://localhost:8080/api/user/user', params)
            .then((res) => {
              this.tableData.data = res.data.result.list
              this.tableData.total =  res.data.result.total
              this.loading = false
            })
            .catch((err) => {
              this.loading = false
              console.log(err, 'err')
            })
            this.loading = false
        },
        deleteId(row) {
            let params = {
                id: row.id,
                name: row.name,
            }
            this.$http.delete('http://localhost:8080/api/user/user', params)
            .then((res) => {
                this.$message({
                    type: 'sucess',
                    message: row.id + ' 删除成功' ,
                })
                this.getList()
            }).catch((err) => {
                console.log(err, 'err')
            })
        },
        closeDialog(action) {
            if (action == 'save') {
                this.queryParams.page = 1
                this.queryParams.pageSize = 10
                this.getList()
            }
            this.metaForm.visible = false
            this.metaForm.title = ""
            this.metaForm.tableInfo = ""
        },

    },
}
</script>


<style scoped>
/* .user {
    height: 100%;
} */

.user .header {
    height: 2%;
}

.user .nav {
    height: 7%;
    display: flex;
    justify-content: space-between;
    align-items: center;
}


.user .content {
    margin-top: 15px; 
}

</style>


