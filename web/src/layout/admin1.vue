<template>
  <div class="layout">
    <admin-left></admin-left>
    
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
                {{ userName }}<i class="el-icon-arrow-down el-icon--right"></i>
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
              <el-breadcrumb-item v-for="(item,i) in tabList" :key="i"
                :to="{ path: item.path }" :replace="false">
                {{ item.title }}
              </el-breadcrumb-item>
            </transition-group>
          </el-breadcrumb>
          
          <div class="remove" @click="initTabList">  
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
import adminLeft from '@/layout/admin1Left.vue'

import {mapState, mapMutations} from 'vuex'

export default {
  data() {
    return {
    }
  },
  created() {
    this.initTabList()
  },
  // watch: {
  //   $route() {
  //     console.log("enter watch")
  //     this.tabList = this.$utils.getItem('tabList')
  //     console.log("this.tablist = ", this.tabList )
  //   },
  //   tabList(val,oldval){
  //     console.log("enter watch tabList")
  //     console.log(val,oldval )
  //   }
  // },
  
  components: {
      adminLeft
  },
  computed: {
      ...mapState({
        isCollapse: state => state.layout.isCollapse,
        tabList: state => state.layout.tabList,
        userName: state => state.login.userName
      })
  },
  methods: {
    ...mapMutations(['layout/updateisCollapse']),
    updateisCollapse() {
        this.$store.commit('layout/updateisCollapse')
    },

    handleTags: function (command) {
      switch (command) {
        case 'all':
          this.closeAll()
          break
        case 'logout':
          console.log("handleTags logout")
          // this.$message("logout")
          this.$store.dispatch('USER_LOGOUT_ACTION')
          this.$router.push({path: '/login'})
          break
      }
    },
    
    // 初始化 tablist，再路由信息中获取
    initTabList(){
      const item = [{
        title: this.$route.meta.title,
        path: this.$route.path,
      }]
      this.$store.commit('layout/initTabList', item)
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
  background-color: #eceef5;
}


.layout .right {
  height: 100%;
  position: absolute;
  left: 208px;
  top: 0;
  right: 0;
  transition: all 0.5s;
}
.layout .right-hide {
  left: 80px;
  transition: all 1s;
}

.right .view-header {
  height: 55px;
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

.right .view-content {
  width: 100%;
  position: absolute;
  top: 90px;
  left: 0;
  right:0;
  bottom: 0;
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
    transform: translate(100%, 100%);
}

.fade-enter-active {
    opacity: 0;
    transform: translate(-40%, -40%);
}
.fade-enter-to{
    opacity: 1;
    transform: translate(0, 0);
}

</style>

