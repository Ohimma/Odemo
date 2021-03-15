<template>
  <div :class="['left', {'left-hide': isCollapse}]">
      <div class="logo">
        <span v-show="!isCollapse">后台管理页面</span>
        <span>
            <li class="el-icon-view" v-show="isCollapse"></li>
        </span>
        <!--  -->
        <!-- <svg-icon v-show="isCollapse" icon-class="vpn_logo"></svg-icon> -->
      </div>

      <el-menu :default-active="onRoutes()"  @select="handleSelectMenu"
        mode="vertical" :collapse="isCollapse"  collapse-transition>
        <template v-for="(item) in menuList">
            <el-menu-item v-if="!item.child" :key="item.index" :index="item.index" @click="updateTabList(item)">
              <li class="el-icon-s-home"></li>
              <span>{{ item.title }}</span>
            </el-menu-item>
        
            <el-submenu v-if="item.child" :key="item.index" :index="item.index">
              <template #title>
                <li :class="item.icon"></li>
                <span>{{ item.title }}</span>
              </template>

              <template  v-for="(itemChild,i) in item.child" :key="i">
                <el-menu-item  :index="itemChild.index" :path="itemChild.path" @click="updateTabList(itemChild)">
                  <!-- <svg-icon icon-class="user"></svg-icon> -->
                  <span>{{ itemChild.title }}</span>
                </el-menu-item>
              </template>
            </el-submenu>
        </template>
      </el-menu>
  </div>
</template>

<script>
import {mapState} from 'vuex'

export default {
  name: 'adminLeft',
  computed: {
    ...mapState({
      menuList: state => state.layout.menuList,
      tabList: state => state.layout.tabList,
      isCollapse: state => state.layout.isCollapse,
    })
  },
  methods: {
    onRoutes: function () {
      return this.$route.path
    },
    handleSelectMenu: function (index) {
      this.$router.push(index)
    },

    updateTabList(item){
      // console.log("enter updateTab start")
      // item = [{
      //   title: this.$route.meta.title,
      //   path: this.$route.path,
      // }]
      // this.$store.commit('layout/initTabList', item)
      this.$store.commit('layout/updateTabList', item)
    },
    
    // goToGroup() {
    //   this.$router.push('/')
    // }
  }
}
</script>

<style scoped>
.left {
  width: 200px;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  background-color: #192349;
  transition: all 0.5s;
}
.left-hide {
  width: 70px;
  transition: all 1s;
}

.left .logo {
  height: 10%;
  cursor: pointer;
  box-sizing: border-box;
  padding: 12% 0;
  text-align: center;
}
.left-hide .logo {
  padding: 30% 0;
  height: 6%;
}


.left .logo span{
  color: #00b8fd;
  font-size: 20px;
  font-weight: 600;
  /* position: absolute;
  vertical-align: middle; 
  top: 50%;
  left: 50%;
  transform: translate(-60%,-50%); */
}
</style>

<style >

.left .el-menu {
  background-color: transparent;
  border: none;
}

.left .el-menu .el-menu-item,
.left .el-menu .el-submenu .el-submenu__title {
  border-left: 3px solid #192349;
  height: 55px;
  color: #c7cedc;
}

/* .left .el-menu .el-menu-item:focus, */
.left .el-menu .el-menu-item:hover,
.left .el-menu .el-submenu__title:hover{
  border-left: 3px solid #00b8fd;
  /* width: 100%; */
  box-sizing: border-box;
  background-color: #143e69;
  color: #00b8fd;
}



.left .el-menu .el-submenu__title.is-active, 
.left .el-menu .el-menu-item.is-active {
  border-left: 3px solid #00b8fd;
  background-color: transparent;
  color: #00b8fd ;
}



/* 可更改右侧折叠 icon */
.left .el-menu .el-submenu__title:hover .el-icon-arrow-down:before,
.left .el-menu .el-menu-item:focus .el-icon-arrow-down:before,
.left .el-menu .el-menu-item:hover .el-icon-arrow-down:before {
  color: red;
  right: 20px;
}

</style>