<template>
  <div class="role" v-loading="loading">

    <div class="role-header">

      <el-form size="mini" inline :model="query" ref="query2">
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
          </el-form>
    </div>

    <div class="role-nav">
        <div class="add">
          <el-button type="primary" size="mini" @click="handle2('add')">
            <span class="el-icon-circle-plus"> 添加</span>
          </el-button>

          <!-- ref 和 点击事件名字关联，prop名字不匹配无法重置 -->
          
        </div>
        
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"

          background small layout="sizes, prev, pager, next"
          :total="tableData.total" :page-sizes="[5, 10, 20]"
          :pager-count="7"
          :currentPage="tableData.currentPage"
          hide-on-single-page>
        </el-pagination>
    </div>

    <div class="role-content" >
    <!-- data 显示的数据
    stripe 是否显示斑马纹
    border 是否显示纵边框
    highlight-current-row 是否高粱当前行
    sortable 是否显示排序选项
    default-expand-all 是否展开行
    lable 显示的标题
    prop 对应列内容的字段名
    show-overflow-tooltip 当内容过长时显示ip
    height="250" 固定表头

    header-align 表头对齐方式，若没有则使用表格对齐方式 align !search || -->
        <el-table :data="tableData.data.filter(data => !search || data.title.toLowerCase().includes(search.toLowerCase()))"  
            style="width: 100%" :stripe="true" :border="true" :highlight-current-row="true"
            :default-expand-all="true" size="small" 
            @sort-change="handleTableSort" >


            <template v-for="(item, index) in dataBase" >
              <el-table-column :key="index" v-if="item.level < 2" :label="item.title"  
              :sortable="item.sortable" :prop="item.name" :min-width="item.width" show-overflow-tooltip>
                  <template #default="scope">
                      <span>{{ scope.row[item.name] }}</span>
                  </template>
              </el-table-column>
            </template>

            <el-table-column align="right" :fit="true" min-width="200px" >
                <template #header>
                    <el-input size="mini" v-model="search" placeholder="在本页搜索角色名称"/>
                </template>

                <template #default="scope">
                    <el-button size="mini"
                    @click="handle2('edit', scope)">Edit</el-button>

                    <el-button size="mini"
                    @click="handle2('detail', scope)">Detail</el-button>

                    <el-button size="mini" type="danger"
                    @click="handle2('delete', scope)">Delete</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>

    <detailRoleTmpl 
        v-if="metaForm.detailVisible"
        :metaForm="metaForm"
        :dataBase="dataBase"
        @closeDialog3="closeDialog3"
    />
    <roleForm 
        v-if="metaForm.visible"
        :metaForm="metaForm"
        :rules="rules"
        :dataBase="dataBase"
        @closeDialog="closeDialog"
    />



  </div>
</template>

<script>
import detailRoleTmpl from '@/components/users/roleDetail2.vue'
import roleForm from '@/components/users/roleForm.vue'
export default {
    components: {
        detailRoleTmpl, roleForm
    },
    data() {
      return {
        query: {
          title: '',
          status: '',
        },
        loading: false,
        search: '',
        metaForm: {
            visible: false,
            detailVisible: false,
            title: "",
            tableInfo: {},
        },
        rules: {
          title: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
          acls: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
        },
        tableData: {
          total: 0,
          page: 1,
          pageSize: 10,
          order_by: 'id',
          order_type: 'desc',
          currentPage: 1,
          data: [],
          // column: data_base.table_label,
          selection: false,
        },
        dataBase:[
          {
            name: "id", title: "ID", "sortable": true, level: 1, width: "80px"
          },
          {
            name: "title", title: "角色名称", sortable: true, level: 1, width: "120px"
          },
          {
            name: "description", title: "角色描述", sortable: false, level: 1, width: "150px"
          },
          {
            name: "user_sum", title: "绑定用户", sortable: false, level: 3, width: "100px"
          },
          {
            name: "status_name", title: "角色状态名称", sortable: false, level: 2, width: "10px"
          },
          {
            name: "creator_name", title: "创建人员名称", sortable: false, level: 2, width: "100px"
          },
          {
            name: "editor_name", title: "编辑账号名称", sortable: false, level: 3, width: "100px"
          },
          {
            name: "created_at", title: "创建时间", sortable: false, level: 2, width: "200px"
          },
          {
            name: "status", title: "角色状态码", sortable: true, level: 1, width: "80px"
          },
          {
            name: "creator_id", title: "创建人员ID", sortable: false, level: 3, width: "200px"
          },
          {
            name: "editor_id", title: "编辑账号ID", sortable: false, level: 3, width: "200px"
          },
          {
            name: "remover_id", title: "操作删除的账号ID", sortable: false, level: 3, width: "200px"
          },
          {
            name: "remover_name", title: "操作删除的账号名称", sortable: false, level: 3, width: "200px"
          },
          {
            name: "updated_at", title: "编辑时间", sortable: false, level: 3, width: "200px"
          },
          {
            name: "deleted_at", title: "删除时间", sortable: false, level: 3, width: "200px"
          },
          {
            name: "acls", title: "权限列表", sortable: false, level: 1, width: "280px", rules: [{ required: true, message: '请选择权限列表', trigger: 'blur' }]
          },
        ],
      }
    },
 
    mounted () {
        console.log("entry user role mount")
        this.getRoleList()
    },
    methods: {

      submitForm(formName) {
        console.log("submit")
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.getRoleList()
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      
     
     
      handle2(type, scope){
        console.log("entry role handle2 = ", type);
        if (type =="detail") {
            this.metaForm.detailVisible = true
            this.metaForm.title = "详细信息"
            this.metaForm.tableInfo = scope.row
        } 
        else if (type == "add") {
            this.metaForm.tableInfo = ""
            this.metaForm.visible = true
            this.metaForm.title = "添加"
        }
        else if (type =="edit"){
            this.metaForm.visible = true
            this.metaForm.title = "编辑"
            this.metaForm.tableInfo = scope.row
        } 
        else if (type =="delete"){
          this.$confirm('确定删除该角色吗？ (' + scope.row.id + ') ' + scope.row.title, '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          })
          .then(() => {
              this.delRoleList(scope.row.id)
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
   
      delRoleList(id) {
        console.log("entry role methods delRoleList");
        this.$message({
            type: 'sucess',
            message: id + ' 删除成功' ,
        })
        // this.$http.delete('http://localhost:8080/api/console/role/' + id
        // )
        // .then((res) => {
        //     this.getRoleList()
        // }).catch((err) => {
        //     console.log(err, 'err')
        // })
      },

      handleTableSort(column){
          console.log("enter sort", column)
          this.tableData.order_by = column.prop
          if (column.order == "ascending") {
            this.tableData.order_type = 'asc'
          } else {
            this.tableData.order_type = 'desc'
          }
          // this.getRoleList()
      },
      handleSizeChange(val) {
        this.tableData.pageSize = val
        this.getRoleList()
      },

      handleCurrentChange(val) {
        this.tableData.page = val
        this.getRoleList()
      },

      getRoleList() {
        this.loading = true
        console.log("entry role methods getRoleList", this.loading)
          let params = {
            page: this.tableData.page,
            pagesize: this.tableData.pageSize,
            role: this.query.title,
            // status: Number(this.query.status),
            status: this.query.status,
            // order_by: this.tableData.order_by,
            // order_type: this.tableData.order_type,
          }
          console.log("xxxxxxxxxxxxxx=", params)
          this.tableData.total = 45
          this.tableData.data = [{id: 1246, title: "编辑角色", description: "哆-狂!孜`诋~笫|", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}] }
,{id: 1239, title: "snjzb", description: "Вячеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1238, title: "ldibo", description: "十!蚱&悍:诼#嗳%", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1237, title: "nzkod", description: "&*!*!:%\]-", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1236, title: "iwjnc", description: "-92684", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1235, title: "jswnf", description: "60888.18619", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1234, title: "ccfhc", description: "-7613383836", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1220, title: "kwqkq", description: "Мартын", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1219, title: "vcumr", description: "牯|潆)猃:铉=蝽?", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1218, title: "vkfdx", description: "xxx", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1217, title: "ekmvx", description: "-27944", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1216, title: "ddoor", description: "72935.15199", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1215, title: "fiegs", description: "-7093981915", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1202, title: "agsty", description: "Панкратий", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1201, title: "omceh", description: "壕-芮'枭.荀@耳*", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1200, title: "atqde", description: "+!_-_.^+_{", status: 1, status_name: "使用中",acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}] }
,{id: 1199, title: "jpanp", description: "-23659", status: 1, status_name: "使用中",acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}] }
,{id: 1198, title: "jqimp", description: "29958.15405", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1197, title: "vribg", description: "-2097676745", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}, {"id": 1, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 4, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 2, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 1, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}
,{id: 1186, title: "ipsfa", description: "Болеслав", status: 3, status_name: "使用中", acls: [{ "id": 2, "parent_id": 0, "title": "boomsh"}]}],
          // this.$http.get('http://localhost:8080/api/console/role', params)
          // .then((res) => {
          //   this.tableData.data = res.data.list
          //   this.tableData.total =  res.data.total
          //   this.loading = false
          // })
          // .catch((err) => {
          //   this.loading = false
          //   console.log(err, 'err')
          // })
        this.loading = false
        console.log("entry role methods getRoleList", this.loading)
      },
      
      closeDialog(action) {
        if (action == 'save') {
          this.tableData.page = 1
          this.getRoleList()
        }
        this.metaForm.visible = false
        this.metaForm.title = ""
        this.metaForm.tableInfo = ""
      },

      
  },
}
</script>


<style scoped>
.role {
    height: 100%;
    background-color: #fff;
}

.role-header {
    height: 10%;
    padding: 18px 10px;
}

.role-header .el-form .el-form-item {
    margin-bottom: 0;
}

.role-nav {
    height: 9%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 3px 10px;
}


.role-nav .el-pagination.el-pagination--small {
  display: inline-block;
}

.role-nav .el-pagination .el-input--mini .el-input__inner {
   height: 20px;
}
.role-nav .el-pagination .el-select .el-input {
  width: 82px;
  margin: 0;
}

.role-content {
    height: 75%;
    margin-top: 5px; 
    display: flex;
    justify-content: flex-start;
    /* background-color: yellow; */
}


</style>

