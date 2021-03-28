<template>
    <div class="auth">
        <div class="header"></div>
        <div class="content">
            <div class="left">
                <!-- 
                show-checkbox 节点是否可被选择 false
                node-key 唯一标识
                default-expand-all 是否默认展开所有节点 false
                default-expanded-keys 是否默认展开的数组
                expand-on-click-node 点击节点时展开或者收缩true 
                highlight-current 是否高亮显示选中节点
                filter-node-method 对树节点筛选时执行的方法
                node-click 点击节点是的回调 默认回传 data, node, elem @node-click="handleTreeNodeClick"
                -->
                <el-input
                    placeholder="输入关键字进行过滤"
                    v-model="filterText">
                </el-input>
                    <el-tree
                        :data="treeData"
                        node-key="id"
                        :default-expanded-keys="[0,1,14,17]"
                        highlight-current
                        :expand-on-click-node="false"
                        :filter-node-method="filterNode"
                        ref="treeBox"
                        >
                        <template #default="{node, data }">
                            <span class="custom-tree-node">
                                <span>{{ node.label }}({{ data.id }})</span>
                                <span v-if="user_auths.indexOf(data.id) >= 0">
                                    <a v-if="data.level < 3" @click="handle2('add', data)"> Add </a>
                                    <a v-if="data.level > 0" @click="handle2('edit', data)"> Edit </a>
                                    <a v-if="data.level > 0" @click="handle2('delete', data)"> Del </a>
                                </span>
                                <span v-if="user_auths.indexOf(data.id) < 0 && user_id == 1">
                                    <a v-if="data.level < 3" @click="handle2('add', data)"> Add </a>
                                    <a v-if="data.level > 0" @click="handle2('edit', data)"> Edit </a>
                                    <a v-if="data.level > 0" @click="handle2('delete', data)"> Del </a>
                                </span>
                                <span v-if="user_auths.indexOf(data.id) < 0 && user_id != 1">
                                    <a class="disabled" v-if="data.level < 3" > Add </a>
                                    <a class="disabled" v-if="data.level > 0" > Edit </a>
                                    <a class="disabled" v-if="data.level > 0" > Del </a>
                                </span>
                            </span>
                        </template>
                    </el-tree>
                </div>

            <div class="right">
                <div class="tips">
                    <span>注意：xxxxx</span>
                </div>
                 <el-table style="width: 100%" :stripe="true" :border="true" :highlight-current-row="true"
                    :header-cell-style="{background:'#edf3ff'}"
                    :default-expand-all="true" size="small" 
                    @row-click="handleDetail"
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

                </el-table>


                <!-- 点击弹窗位置 -->
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

        </div>
    </div>
</template>

<script>
import detail from "@/components/users/Detail.vue"
import edit from "@/components/users/authEdit.vue"
import {mapState} from 'vuex'

  export default {
    components: {
        edit, detail
    },
    data() {
      return {
        search:'',
        filterText: '',
        treeData: [],
        tableData: {
            data: [],
            total: 0,
        },
        rules: {
            name: [{ required: true, message: '请输入名字', trigger: 'blur' }],
            pid: [{ required: true, message: '请输入父id', trigger: 'blur' }],
            url: [{ required: true, message: '请输入url', trigger: 'blur' }],
            method: [{ required: true, message: '请输入method', trigger: 'blur' }],
        },
        dataBase: [
            { name: "id", label: "Id", "sortable": true, level: 4, width: "30px" },
            { name: "pid", label: "Pid", "sortable": true, level: 4, width: "40px" },
            { name: "level", label: "菜单", "sortable": true, level: 5, width: "40px" },
            { name: "name", label: "名字", sortable: false, level: 1, width: "60px" },
            { name: "url", label: "路径", sortable: true, level: 1, width: "60px" },
            { name: "method", label: "访问方法", sortable: true, level: 1, width: "50px" },
            { name: "icon", label: "图标", sortable: false, level: 3, width: "60px" },
            { name: "create_at", label: "创建时间", sortable: false, level: 6, width: "80px" },
            { ame: "update_at", label: "更新时间", sortable: false, level: 6, width: "80px" },
            { name: "detail", label: "备注", sortable: false, level: 3, width: "60px" },
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
    mounted(){
        console.log("entry auth mounted .....", )   
        this.getList()
    },
    computed: {
        ...mapState({
            authListTree: state => state.layout.authListTree,
            user_auths: state => state.layout.user_auths,
            user_id: state => state.layout.user_id,
        })
    },
    watch: {
      filterText(val) {
        console.log("entry watch mounted .....")
        this.$refs.treeBox.filter(val);
      }
    },
    methods: {
        getList(){
            this.loading = true
            console.log("entry auth methods getList", this.loading)
            this.$http.get('http://localhost:8080/api/user/auth')
                .then((res) => {
                this.$store.dispatch('layout/AUTH_LIST_ACTION', res.data.result.list)
                this.$store.dispatch('layout/AUTH_LIST_TREE_ACTION', this.$Conver.convTotree(res.data.result.list, 0))

                this.tableData.data = res.data.result.list
                this.tableData.total =  res.data.result.total
                this.treeData = [{id: 0, label: "根节点", level: 0, children: this.authListTree}]
                this.loading = false
                })
                .catch((err) => {
                this.loading = false
                console.log(err, 'err')
                })
            this.loading = false
            console.log("entry auth methods getRoleList", this.loading)
        },
      
        filterNode(value, data) {
            if (!value) return true;
            return data.label.indexOf(value) !== -1;
        },
        handleDetail(row){
            this.metaForm.detailVisible = true
            this.metaForm.title = "详细信息"
            this.metaForm.tableInfo = row
        },
        handle2(type, data){
            console.log("entry role handle2 = ", type);
            if (type == "add") {
                this.metaForm.tableInfo = data
                this.metaForm.visible = true
                this.metaForm.title = "添加"
            } else if (type =="edit"){
                this.metaForm.visible = true
                this.metaForm.title = "编辑"
                this.metaForm.tableInfo = data
            } else if (type =="detail") {
                this.handleDetail(data)
            } else if (type =="delete"){
                this.$confirm('确定删除该角色吗？ (' + data.id + ') ' + data.name, '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning',
                })
                .then(() => {
                    this.deleteId(data)
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

        deleteId(row) {
            let params = {
                id: row.id,
                name: row.name,
            }
            this.$http.delete('http://localhost:8080/api/user/auth', params)
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
                this.getList()
            }
            this.metaForm.visible = false
            this.metaForm.title = ""
            this.metaForm.tableInfo = ""
        },
    }
  };
</script>

<style scoped>

.auth {
    color: #606266;
}

.auth .header {
    height: 1%;
}

.auth .content {
    height: 98%;
    display: flex;
    justify-content: space-between;
}
.auth .content .left{
    height: 100%;
    width: 30%;
}

.auth .content .left .el-input{
    width: 90%;
    padding: 10px 10px;
}

.auth .content .left .el-tree{
    height: 88%;
    background: #edf3ff;
    margin: 0 21px 0 10px;
    border-radius: 15px;
}
.auth .content .left  .el-tree-node__content {
    height: 30px;
}

.auth .content .left  .custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 13px;
    padding-right: 8px;
}


.auth .content .right{
    height: 100%;
    width: 70%;
    margin-right: 10px;
}
.auth .content .right .tips{
    height: 60px;
    box-sizing: border-box;
    padding: 20px 0;
}
.auth a {
    color: #e6a23b;
    margin: 0 2px;
}
.auth a:hover {
    color: red;
    cursor: pointer;
}

.auth .disabled {
    color: #c1c4cb;
}
.auth .disabled:hover {
    cursor: help;
    color: #c1c4cb;
}
</style>

<style >
.auth .content .left input.el-input__inner{
    border-radius: 20px;;
    background: #edf3ff;

}
</style>