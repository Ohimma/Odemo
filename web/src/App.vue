<template>

  <router-view class="router-view" v-slot="{ Component }">
    <transition :name="transitionName" mode="out-in">
      <component :is="Component" />
    </transition>
  </router-view>
  
</template>


<script>
export default {
  data () {
    return {
      transitionName: ''
    }
  },
  watch: {
    $route (to, from) {
      if (from.path == "/login" && to.path == "/home" ){
        this.transitionName = 'login'
      } 
      else if  (from.path == "/home" && to.path == "/login" ) {
        this.transitionName = 'logout'
      } 
    }
  }
}
</script>

<style scoped>
.login-enter-active,
.login-leave-active,
.logout-enter-active,
.logout-leave-active {
    height: 100%;
    transition: all 10ms cubic-bezier(.55,0,.1,1);
    /* position: absolute; */
    /* backface-visibility: hidden; */
}

.login-enter-active{
    opacity: 0;
    transform: translate3d(100%, 0, 0);
}
.login-enter-to {
  opacity: 1;
  transform: translate3d(0, 0, 0);
}
.logout-enter-active{
    opacity: 0;
    transform: translate3d(-100%, 0, 0);
}
.logout-enter-to {
  opacity: 1;
  transform: translate3d(0, 0, 0);
}


.login-leave-to {
  opacity: 0;
  transform: translate3d(-100%, 100%, 0);
}

.logout-leave-to {
  opacity: 0;
  transform: translate3d(100%, 100%, 0);
}



</style>>
