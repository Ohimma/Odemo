<template>
  <div class="login">
    <div class="login-msg">
      <div class="title">
        <!-- <img src="../assets/img/login-logo.png" alt=""> -->
        <img src="https://gimg2.baidu.com/image_search/src=http%3A%2F%2Finews.gtimg.com%2Fnewsapp_bt%2F0%2F8851510617%2F640.jpg&refer=http%3A%2F%2Finews.gtimg.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1618458458&t=5bdab2e5e4d75abf5595723d479ca382" alt="">
        <span>管理后台</span>
      </div>
      <!-- 
        :model指定表单使用的数据 :rules进行表单校验规则 ref是提交的对象名字
        prop表单校验的属性
        input 绑定的元素必须是model的直接属性 
        prefix 左边 suffix 右边 clearable 删除所有输入 
        :loading  true时会转圈加载 false时正常
        status-icon 添加校验结果反馈显示图标
      -->
      <el-form :model="loginData" :rules="loginDataRules" ref="loginForm">
        <el-form-item prop="name">
          <el-input type="text" 
            v-model="loginData.name" 
            prefix-icon="el-icon-user" 
            placeholder="请输入用户名"
            tabindex="1"
            auto-complete="on">
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="loginData.password"
            type="password"
            placeholder="请输入密码"  
            prefix-icon="el-icon-unlock"
            clearable show-password
            
            @keyup.enter="loginBtn"
            auto-complete="on">
          </el-input>
        </el-form-item>

        <el-form-item>
          <el-button 
            type="primary"
            :loading="loading" 
            @click.prevent="loginBtn">登&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      loginData: {
        name: '',
        password: ''
      },
      loginDataRules: {
        name: {required: true, message: '请输入用户名', trigger: 'change'},
        password: [{required: true, message: '请输入密码', trigger: 'change'}]
      },
      loading: false,
    }
  },
  created(){
    console.log("entrey login create ......")
  },
  mounted(){
    console.log("entrey login mounted .......")
    // this.loginData.name = localStorage.getItem('name')
  },
  methods: {
    loginBtn () {
      this.loading = true
      // validate 对整个表单验证，参数为回调函数，不传入函数返回promise
      this.$refs.loginForm.validate(valid => {
        if (valid) {  // vaild 是 true/false
          this.$http.post(process.env.VUE_APP_BASEURL + 'api/login', {
            name: this.loginData.name,
            password: this.loginData.password,
          }, false)
            .then((res) => {
              this.$message({
                type: 'success',
                message: '登陆成功'
              })
              this.$store.dispatch('layout/USER_INFO_ACTION', res.data.result)
              this.$store.dispatch('layout/USER_TOKEN_ACTION', res.data.result.token)
              // this.$store.dispatch('layout/USER_NAME_ACTION', this.loginData.name)
              // 登陆成功后，判断是否有redirect，没有的话跳转到/home
              // decodeURIComponent对encodeURLComponent的url乱码解码，
              let redirect = decodeURIComponent(this.$route.query.redirect || '/home')
              this.$router.push({
                path: redirect
              })
              this.loading = false
            })
            .catch(() => {
              this.loading = false
              this.$message({
                type: 'error', 
                message: "认证失败"
              })
              
            })
        } else {
          this.loading = false
        }
      })
    }
  }
}
</script>

<style scoped>
.login {
  height: 100%;
  width: 100%;
  /* display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  background: #11255A url('../assets/img/login-bg.png') no-repeat center bottom;
   */

  background: #090d09 url('https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fwww.yyxt.com%2Fuploads%2Fallimg%2F160318%2F11-16031Q11350.jpg&refer=http%3A%2F%2Fwww.yyxt.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1618457703&t=6d3f1fd89769961b75784d049e106d73') 
              no-repeat right bottom;
  background-size: 90% 100%;

  transition: all 1s;
}


.login .login-msg {
    position: relative;
    top: 23%;
    left: 10%;
    width: 400px;
    height: 400px;
    padding: 20px;
    box-sizing: border-box;
}
.login .login-msg .title {
    width: 100%;
    height: 25%;

    display: flex;
    justify-content: center;
    align-items: center;

    margin-bottom: 20px;
}
.login .login-msg .title img {
    width: 70px;
    height: 40px;
}
.login .login-msg .title span {
  font-size: 20px;
  line-height: 50px;
  color: #FFFFFF;
  margin: 0px 0px;
}

.el-button {
  width: 100%;
  height: 40px;
  margin-top: 25px;
  
  background: #204CC0;
  border: 0;
  border-radius: 4px;

  color: #fff;
  font-size: 20px;
}
.el-button:hover,
.el-button.is-fouse {
  background: #4d73da;
} 


</style>

<style >
.login .el-form .el-input__inner {
  padding-left: 40px;
  letter-spacing: 2px;
  font-size: 16px;
  font-family: "MingLiu";
}
</style>