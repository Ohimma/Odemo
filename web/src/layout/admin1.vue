<template>
  <div class="layout">
    <!-- <admin-left></admin-left> -->
    <div :class="['left', {'left-hide': isCollapse}]">
      <div class="logo">
        <span v-show="!isCollapse">后台管理页面</span>
        <span>
            <li class="el-icon-view" v-show="isCollapse"></li>
        </span>
        <!--  -->
        <!-- <svg-icon v-show="isCollapse" icon-class="vpn_logo"></svg-icon> -->
      </div>

      <el-menu :default-active="leftOnRoutes()"  @select="leftSelectMenu"
        mode="vertical" :collapse="isCollapse"  collapse-transition>
        
        <el-menu-item index="/home" @click="updateNavList(item)">
          <li class="el-icon-s-home"></li>
          <span>首页</span>
        </el-menu-item>

        <template v-for="(item, index) in authListTree">
            <el-menu-item v-if="!item.children" :key="index" :index="item.url" @click="updateNavList(item)">
              <li :class="item.icon"></li>
              <span>{{ item.name }}</span>
            </el-menu-item>
        
            <el-submenu v-if="item.children" :key="index" :index="item.url">
                <template #title>
                  <li :class="item.icon"></li>
                  <span>{{ item.name }}</span>
                </template>

                <template  v-for="(itemChild,i) in item.children" :key="i">
                  <el-menu-item  :index="itemChild.url" :path="itemChild.url" @click="updateNavList(itemChild)">
                    <!-- <svg-icon icon-class="user"></svg-icon> -->
                    <span>{{ itemChild.name }}</span>
                  </el-menu-item>
                </template>
            </el-submenu>
        </template>
      </el-menu>
    </div>
    
    <div :class="['right', {'right-hide': isCollapse}]">
        <div class="view-header">
          <div class="left-fold">            
            <i @click="updateisCollapse" :class="isCollapse ? 'el-icon-s-fold':'el-icon-s-unfold'"></i>
          </div>
          <div class="user">
            <!-- <div class="user-icon"><img src="@/assets/img/user.png" alt=""></div> -->
            <div class="user-icon"><img src="https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80" alt=""></div>
           
            <el-dropdown @command="handleTags">
              <span class="el-dropdown-link">
                {{ user_name }}<i class="el-icon-arrow-down el-icon--right"></i>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item command="a">黄金糕</el-dropdown-item>
                  <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </div>

        <div class="view-nav">
          <el-breadcrumb separator-class="el-icon-arrow-right">
            <transition-group name="nav">

              <el-breadcrumb-item v-for="(item, i) in navList" :key="i"
                :to="{ path: item.url }" :replace="false">
                {{ item.name }}
              </el-breadcrumb-item>

            </transition-group>
          </el-breadcrumb>
          
          <div class="remove" @click="initNavList">  
            <i class="el-icon-delete"></i>
          </div>
        </div>

        <div class="view-content">
          <router-view v-slot="{ Component }">
            <transition name="fade" mode="out-in">
              <component :is="Component"/>
            </transition>
          </router-view>
        </div>
        
    </div>
  </div>
</template>

<script>  

import {mapState} from 'vuex'

export default {

  data() {
    return {
      isCollapse: false,
      navList: [],
    }
  },
  created() {
    console.log("entrey admin created ......")
    this.initNavList()
    this.setAuthList()
    this.setUserAuth()
    console.log("leavel admin created ......")
  },
  mounted(){
    console.log("entrey admin mounted .......")
    console.log("leavel admin mounted .......")
  },
  // watch: {
  //   $route() {
  //     console.log("enter admin watch ....", this.$route)
  //   },
    
  // },
  
  computed: {
      ...mapState({
        tabList: state => state.layout.tabList,
        user_id: state => state.layout.user_id,
        user_name: state => state.layout.user_name,
        user_role_ids: state => state.layout.user_role_ids,
        authListTree: state => state.layout.authListTree,
      }),
  },
  methods: {
    leftOnRoutes: function () {
      return this.$route.path
    },
    leftSelectMenu: function (index) {
      this.$router.push(index)
    },
    updateisCollapse() {
        this.isCollapse = !this.isCollapse
    },

    initNavList(){
      this.navList = [{
        name: "/",
        url: this.$route.path.path,
      }]
    },

    updateNavList(item){
      this.navList.push({
          name: item.name,
          url: item.url
      })
    },

    setAuthList(){
      this.$http.get(process.env.VUE_APP_BASEURL + 'api/user/auth')
        .then((res) => {
          this.$store.dispatch('layout/AUTH_LIST_ACTION', res.data.result.list)
          this.$store.dispatch('layout/AUTH_LIST_TREE_ACTION', this.$Conver.convTotree(res.data.result.list, 0))
        })
        .catch((err) => {
          this.$message({
              type: "error",
              message: err,
            })

        })
    },

    setUserAuth(){
      let params = {
        role_ids: this.user_role_ids,
      }
      this.$http.get(process.env.VUE_APP_BASEURL + 'api/user/role', params)
        .then((res) => {
          this.$store.dispatch('layout/USER_AUTHS_ACTION', res.data.result.user_auths)
        })
        .catch((err) => {
          this.$message({
              type: "error",
              message: err,
            })
        })
    },
    

    handleTags: function (command) {
      switch (command) {
        case 'all':
          this.closeAll()
          break
        case 'logout':
          this.$store.dispatch('layout/USER_LOGOUT_ACTION')
          this.$router.push({path: '/login'})
          break
      }
    },
    
    
  }
}
</script>

<style scoped>
.layout {
  height: 100%;
  width: 100%;
  position: relative;
  transition: all 1s;
  /* background-color: #eceef5; */
}

.layout .left {
  width: 200px;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  background-color: #1d1e23;
  transition: all 0.5s;
}
.left-hide {
  width: 70px;
  transition: all 1s;
}

.left .logo {
  height: 70px;
  cursor: pointer;
  box-sizing: border-box;
  padding: 11% 0;
  text-align: center;
}
.left-hide .logo {
  padding: 30% 0;
  height: 6%;
}

.left .logo span{
  color: #dbb68d ;
  font-size: 20px;
  font-weight: 500;
}

.layout .right {
  min-height: 100%;
  position: absolute;
  left: 201px;
  top: 0;
  right: 0;
  transition: all 0.5s;
  background: #eceef5;
}
.layout .right-hide {
  left: 80px;
  transition: all 1s;
}

.right .view-header {
  height: 50px;
  display: flex;
  justify-content: space-between;
  background-color: #fff;
  box-shadow: 5px 6px 5px 0 #E2E8EC;
  padding: 0 10px;
}
.right .view-nav {
  height: 40px;
  padding: 0 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.right .view-nav .remove:hover {
  cursor: pointer;
}

.right .view-content {
  width: 98%;
  box-sizing: border-box;
  margin: 0 12px;
  padding: 15px 15px;
  background-color: #fff;
}

.right .view-header .left-fold {
  padding: 13px 0;
}
.right .view-header .left-fold .el-icon-s-fold,
.right .view-header .left-fold .el-icon-s-unfold {
    font-size: 30px;
    cursor: pointer;
} 

.right .view-header .user {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.right .view-header .user img {
  height: 30px;
  width: 30px;
  margin: 0 10px;
}


.fade-enter-active, .fade-leave-active{
    transition: all 0.3s;
}

.fade-leave{
    opacity: 1;
    transform: translate(0, 0);
}
.fade-leave-to{
    opacity: 0;
    transform: translate(70%, 70%);
}


</style>


<style >


.left .el-menu {
  background-color: transparent;
  border: none;
}

.left .el-menu .el-menu-item,
.left .el-menu .el-submenu .el-submenu__title {
  border-left: 3px solid #101117;
  height: 55px;
  color: #c7cedc;
}

.left .el-menu .el-menu-item:focus,
.left .el-menu .el-menu-item:hover,
.left .el-menu .el-submenu__title:hover{
  /* background-color: #143e69;
  border-left: 3px solid #00b8fd;
  color: #00b8fd; */

  box-sizing: border-box;
  border-left: 3px solid #333439;
  background-color: #101117;
  color: #dbb68d;
}

.left .el-menu .el-submenu__title.is-active, 
.left .el-menu .el-menu-item.is-active {
  background-color: transparent;
  /* border-left: 3px solid #00b8fd;
  color: #00b8fd ; */

  border-left: 3px solid #333439;
  color: #dbb68d ;
  /*  */
}

/* 可更改右侧折叠 icon */
.left .el-menu .el-submenu__title:hover .el-icon-arrow-down:before,
.left .el-menu .el-menu-item:focus .el-icon-arrow-down:before,
.left .el-menu .el-menu-item:hover .el-icon-arrow-down:before {
  color: red;
  right: 20px;
}
</style>